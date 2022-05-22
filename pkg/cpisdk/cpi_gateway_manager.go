/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package cpisdk

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"k8s.io/klog"
)

// CpiGatewayManager represents CPI's custom gateway manager. The separation is because CPI does some set of custom
// operations on RDE after operations like Create/Delete virtual services.
type CpiGatewayManager struct {
	VcdGatewayManager *vcdsdk.GatewayManager
	ClusterID         string
}

// NewCpiGatewayManager is a factory method to generate a new CpiGatewayManager object.
func NewCpiGatewayManager(ctx context.Context, client *vcdsdk.Client, networkName string, ipamSubnet string, clusterID string) (*CpiGatewayManager, error) {
	gatewayManger, err := vcdsdk.NewGatewayManager(ctx, client, networkName, ipamSubnet)
	if err != nil {
		return nil, fmt.Errorf("failed to create GatewayManager: [%v]", err)
	}

	return &CpiGatewayManager{
		VcdGatewayManager: gatewayManger,
		ClusterID:         clusterID,
	}, nil
}

func (cgm *CpiGatewayManager) CreateLoadBalancer(ctx context.Context, virtualServiceNamePrefix string, lbPoolNamePrefix string,
	ips []string, portDetailsList []vcdsdk.PortDetails, oneArm *vcdsdk.OneArm) (string, error) {

	gm := cgm.VcdGatewayManager
	if gm == nil {
		return "", fmt.Errorf("GatewayManager cannot be nil")
	}
	client := gm.Client
	client.RWLock.Lock()
	defer client.RWLock.Unlock()

	if len(portDetailsList) == 0 {
		// nothing to do here
		klog.Infof("There is no port specified. Hence nothing to do.")
		return "", fmt.Errorf("nothing to do since http and https ports are not specified")
	}

	if gm.GatewayRef == nil {
		return "", fmt.Errorf("gateway reference should not be nil")
	}

	// Separately loop through all DNAT rules to see if any exist, so that we can reuse the external IP in case a
	// partial creation of load-balancer is continued and an externalIP was claimed earlier by a dnat rule
	externalIP := ""
	var err error
	if oneArm != nil {
		for _, portDetails := range portDetailsList {
			if portDetails.InternalPort == 0 {
				continue
			}

			virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portDetails.PortSuffix)
			dnatRuleName := vcdsdk.GetDNATRuleName(virtualServiceName)
			dnatRuleRef, err := gm.GetNATRuleRef(ctx, dnatRuleName)
			if err != nil {
				return "", fmt.Errorf("unable to retrieve created dnat rule [%s]: [%v]", dnatRuleName, err)
			}
			if dnatRuleRef == nil {
				continue // ths implies that the rule does not exist
			}

			if externalIP != "" && externalIP != dnatRuleRef.ExternalIP {
				return "", fmt.Errorf("as per dnat there are two external IP rules for the same service: [%s], [%s]",
					externalIP, dnatRuleRef.ExternalIP)
			}

			externalIP = dnatRuleRef.ExternalIP
		}
	}

	if externalIP == "" {
		externalIP, err = gm.GetUnusedExternalIPAddress(ctx, gm.IPAMSubnet)
		if err != nil {
			return "", fmt.Errorf("unable to get unused IP address from subnet [%s]: [%v]",
				gm.IPAMSubnet, err)
		}
	}
	klog.Infof("Using external IP [%s] for virtual service\n", externalIP)

	for _, portDetails := range portDetailsList {
		if portDetails.InternalPort == 0 {
			klog.Infof("No internal port specified for [%s], hence loadbalancer not created\n",
				portDetails.PortSuffix)
			continue
		}

		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portDetails.PortSuffix)
		lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portDetails.PortSuffix)

		vsSummary, err := gm.GetVirtualService(ctx, virtualServiceName)
		if err != nil {
			return "", fmt.Errorf("unexpected error while querying for virtual service [%s]: [%v]",
				virtualServiceName, err)
		}
		if vsSummary != nil {
			if vsSummary.LoadBalancerPoolRef.Name != lbPoolName {
				return "", fmt.Errorf("virtual Service [%s] found with unexpected loadbalancer pool [%s]",
					virtualServiceName, lbPoolName)
			}

			klog.V(3).Infof("LoadBalancer Virtual Service [%s] already exists", virtualServiceName)
			if err = gm.CheckIfVirtualServiceIsPending(ctx, virtualServiceName); err != nil {
				return "", err
			}

			continue
		}

		virtualServiceIP := externalIP
		if oneArm != nil {
			internalIP, err := gm.GetUnusedInternalIPAddress(ctx, oneArm)
			if err != nil {
				return "", fmt.Errorf("unable to get internal IP address for one-arm mode: [%v]", err)
			}

			dnatRuleName := vcdsdk.GetDNATRuleName(virtualServiceName)
			if err = gm.CreateDNATRule(ctx, dnatRuleName, externalIP, internalIP,
				portDetails.ExternalPort, portDetails.InternalPort); err != nil {
				return "", fmt.Errorf("unable to create dnat rule [%s:%d] => [%s:%d]: [%v]",
					externalIP, portDetails.ExternalPort, internalIP, portDetails.InternalPort, err)
			}
			// use the internal IP to create virtual service
			virtualServiceIP = internalIP

			// We get an IP address above and try to get-or-create a DNAT rule from external IP => internal IP.
			// If the rule already existed, the old DNAT rule will remain unchanged. Hence we get the old externalIP
			// from the old rule and use it. What happens to the new externalIP that we selected above? It just remains
			// unused and hence does not get allocated and disappears. Since there is no IPAM based resource
			// _acquisition_, the new externalIP can just be forgotten about.
			dnatRuleRef, err := gm.GetNATRuleRef(ctx, dnatRuleName)
			if err != nil {
				return "", fmt.Errorf("unable to retrieve created dnat rule [%s]: [%v]", dnatRuleName, err)
			}
			if dnatRuleRef == nil {
				return "", fmt.Errorf("retrieved dnat rule ref is nil")
			}

			externalIP = dnatRuleRef.ExternalIP
		}

		segRef, err := gm.GetLoadBalancerSEG(ctx)
		if err != nil {
			return "", fmt.Errorf("unable to get service engine group from edge [%s]: [%v]",
				gm.GatewayRef.Name, err)
		}

		lbPoolRef, err := gm.CreateLoadBalancerPool(ctx, lbPoolName, ips, portDetails.InternalPort)
		if err != nil {
			return "", fmt.Errorf("unable to create load balancer pool [%s]: [%v]", lbPoolName, err)
		}

		virtualServiceRef, err := gm.CreateVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
			virtualServiceIP, portDetails.Protocol, portDetails.ExternalPort,
			portDetails.UseSSL, portDetails.CertAlias)
		if err != nil {
			// return  plain error if vcdsdk.VirtualServicePendingError is returned. Helps the caller recognize that the
			// error is because VirtualService is still in Pending state.
			if _, ok := err.(*vcdsdk.VirtualServicePendingError); ok {
				return "", err
			}
			return "", err
		}

		// TODO: add VIP to RDE
		// update RDE with virtual IP
		rm := NewRDEManager(client, cgm.ClusterID)
		err = rm.addVirtualIpToRDE(ctx, externalIP)
		if err != nil {
			klog.Errorf("error when adding virtual IP to RDE: [%v]", err)
		}

		klog.Infof("Created Load Balancer with virtual service [%v], pool [%v] on gateway [%s]\n",
			virtualServiceRef, lbPoolRef, gm.GatewayRef.Name)
	}

	return externalIP, nil
}

func (cgm *CpiGatewayManager) UpdateLoadBalancer(ctx context.Context, lbPoolName string, virtualServiceName string,
	ips []string, internalPort int32, externalPort int32) error {

	gm := cgm.VcdGatewayManager
	if gm == nil {
		return fmt.Errorf("GatewayManager cannot be nil")
	}

	client := gm.Client
	client.RWLock.Lock()
	defer client.RWLock.Unlock()

	_, err := gm.UpdateLoadBalancerPool(ctx, lbPoolName, ips, internalPort)
	if err != nil {
		if lbPoolBusyErr, ok := err.(*vcdsdk.LoadBalancerPoolBusyError); ok {
			klog.Errorf("update loadbalancer pool failed; loadbalancer pool [%s] is busy: [%v]", lbPoolName, err)
			return lbPoolBusyErr
		}
		return fmt.Errorf("unable to update load balancer pool [%s]: [%v]", lbPoolName, err)
	}
	err = gm.UpdateVirtualServicePort(ctx, virtualServiceName, externalPort)
	if err != nil {
		if vsBusyErr, ok := err.(*vcdsdk.VirtualServiceBusyError); ok {
			klog.Errorf("update virtual service failed; virtual service [%s] is busy: [%v]", virtualServiceName, err)
			return vsBusyErr
		}
		return fmt.Errorf("unable to update virtual service [%s] with port [%d]: [%v]", virtualServiceName, externalPort, err)
	}
	// update app port profile
	dnatRuleName := vcdsdk.GetDNATRuleName(virtualServiceName)
	appPortProfileName := vcdsdk.GetAppPortProfileName(dnatRuleName)
	err = gm.UpdateAppPortProfile(appPortProfileName, externalPort)
	if err != nil {
		return fmt.Errorf("unable to update application port profile [%s] with external port [%d]: [%v]", appPortProfileName, externalPort, err)
	}

	// update DNAT rule
	dnatRuleRef, err := gm.GetNATRuleRef(ctx, dnatRuleName)
	if err != nil {
		return fmt.Errorf("unable to retrieve created dnat rule [%s]: [%v]", dnatRuleName, err)
	}
	err = gm.UpdateDNATRule(ctx, dnatRuleName, dnatRuleRef.ExternalIP, dnatRuleRef.InternalIP, externalPort)
	if err != nil {
		return fmt.Errorf("unable to update DNAT rule [%s]: [%v]", dnatRuleName, err)
	}
	return nil
}

func (cgm *CpiGatewayManager) DeleteLoadBalancer(ctx context.Context, virtualServiceNamePrefix string,
	lbPoolNamePrefix string, portDetailsList []vcdsdk.PortDetails, oneArm *vcdsdk.OneArm) error {
	gm := cgm.VcdGatewayManager
	if gm == nil {
		return fmt.Errorf("GatewayManager cannot be nil")
	}

	client := gm.Client
	client.RWLock.Lock()
	defer client.RWLock.Unlock()

	// TODO: try to continue in case of errors
	var err error

	// Here the principle is to delete what is available; retry in case of failure
	// but do not fail for missing entities, since a retry will always have missing
	// entities.
	for _, portDetails := range portDetailsList {
		if portDetails.InternalPort == 0 {
			klog.Infof("No internal port specified for [%s], hence loadbalancer not created\n",
				portDetails.PortSuffix)
			continue
		}

		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portDetails.PortSuffix)
		lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portDetails.PortSuffix)

		// get external IP
		rdeVIP := ""
		dnatRuleName := ""
		if oneArm != nil {
			dnatRuleName = vcdsdk.GetDNATRuleName(virtualServiceName)
			dnatRuleRef, err := gm.GetNATRuleRef(ctx, dnatRuleName)
			if err != nil {
				return fmt.Errorf("unable to get dnat rule ref for nat rule [%s]: [%v]", dnatRuleName, err)
			}
			if dnatRuleRef != nil {
				rdeVIP = dnatRuleRef.ExternalIP
			}
		} else {
			vsSummary, err := gm.GetVirtualService(ctx, virtualServiceName)
			if err != nil {
				return fmt.Errorf("unable to get summary for LB Virtual Service [%s]: [%v]",
					virtualServiceName, err)
			}
			if vsSummary != nil {
				rdeVIP = vsSummary.VirtualIpAddress
			}
		}

		err = gm.DeleteVirtualService(ctx, virtualServiceName, false)
		if err != nil {
			if vsBusyErr, ok := err.(*vcdsdk.VirtualServiceBusyError); ok {
				klog.Errorf("delete virtual service failed; virtual service [%s] is busy: [%v]", virtualServiceName, err)
				return vsBusyErr
			}
			return fmt.Errorf("unable to delete virtual service [%s]: [%v]", virtualServiceName, err)
		}

		// remove virtual ip from RDE
		rm := NewRDEManager(client, cgm.ClusterID)
		err = rm.removeVirtualIpFromRDE(ctx, rdeVIP)
		if err != nil {
			return fmt.Errorf("error when removing vip from RDE: [%v]", err)
		}

		err = gm.DeleteLoadBalancerPool(ctx, lbPoolName, false)
		if err != nil {
			if lbPoolBusyErr, ok := err.(*vcdsdk.LoadBalancerPoolBusyError); ok {
				klog.Errorf("delete loadbalancer pool failed; loadbalancer pool [%s] is busy: [%v]", lbPoolName, err)
				return lbPoolBusyErr
			}
			return fmt.Errorf("unable to delete load balancer pool [%s]: [%v]", lbPoolName, err)
		}

		if oneArm != nil {
			err = gm.DeleteDNATRule(ctx, dnatRuleName, false)
			if err != nil {
				return fmt.Errorf("unable to delete dnat rule [%s]: [%v]", dnatRuleName, err)
			}
		}
	}

	return nil
}

func (cgm *CpiGatewayManager) GetLoadBalancer(ctx context.Context, virtualServiceName string, oneArm *vcdsdk.OneArm) (string, error) {
	gm := cgm.VcdGatewayManager
	if gm == nil {
		return "", fmt.Errorf("GatewayManager cannot be nil")
	}
	return gm.GetLoadBalancer(ctx, virtualServiceName, oneArm)
}

func (cgm *CpiGatewayManager) GetLoadBalancerPool(ctx context.Context, lbPoolName string) (*swaggerClient.EntityReference, error) {
	gm := cgm.VcdGatewayManager
	if gm == nil {
		return nil, fmt.Errorf("GatewayManager cannot be nil")
	}
	return gm.GetLoadBalancerPool(ctx, lbPoolName)
}
