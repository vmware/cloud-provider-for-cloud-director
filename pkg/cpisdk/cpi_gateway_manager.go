/*
   Copyright 2022 VMware, Inc.
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

func (cgm *CpiGatewayManager) CreateLoadBalancer(ctx context.Context, virtualServiceNamePrefix string,
	lbPoolNamePrefix string, ips []string, portDetailsList []vcdsdk.PortDetails, oneArm *vcdsdk.OneArm,
	enableVirtualServiceSharedIP bool, portNameToIP map[string]string) (string, error) {

	gm := cgm.VcdGatewayManager
	if gm == nil {
		return "", fmt.Errorf("GatewayManager cannot be nil")
	}

	externalIP, err := gm.CreateLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix,
		ips, portDetailsList, oneArm, enableVirtualServiceSharedIP, portNameToIP, "")
	if err != nil {
		return "", fmt.Errorf(
			"unable to create load balancer with vs prefix [%s], lbpool prefix [%s], ips [%v], ports [%v]: [%v]",
			virtualServiceNamePrefix, lbPoolNamePrefix, ips, portDetailsList, err)
	}

	if externalIP != "" {
		rm := NewRDEManager(gm.Client, cgm.ClusterID)
		err = rm.addVirtualIpToRDE(ctx, externalIP)
		if err != nil {
			klog.Errorf("error when adding virtual IP to RDE: [%v]", err)
		}
	}

	return externalIP, nil
}

func (cgm *CpiGatewayManager) UpdateLoadBalancer(ctx context.Context, lbPoolName string, virtualServiceName string,
	ips []string, internalPort int32, externalPort int32, oneArm *vcdsdk.OneArm, enableVirtualServiceSharedIP bool,
	protocol string) error {

	gm := cgm.VcdGatewayManager
	if gm == nil {
		return fmt.Errorf("GatewayManager cannot be nil")
	}

	if err := gm.UpdateLoadBalancer(ctx, lbPoolName, virtualServiceName, ips, internalPort, externalPort,
		oneArm, enableVirtualServiceSharedIP, protocol); err != nil {
		return fmt.Errorf(
			"unable to create load balancer with vs [%s], lbpool [%s], ips [%v], ports [%d->%d]: [%v]",
			virtualServiceName, lbPoolName, ips, externalPort, internalPort, err)
	}

	return nil
}

func (cgm *CpiGatewayManager) DeleteLoadBalancer(ctx context.Context, virtualServiceNamePrefix string,
	lbPoolNamePrefix string, portDetailsList []vcdsdk.PortDetails, oneArm *vcdsdk.OneArm) error {
	gm := cgm.VcdGatewayManager
	if gm == nil {
		return fmt.Errorf("GatewayManager cannot be nil")
	}

	rdeVIP, err := gm.DeleteLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, oneArm)
	if err != nil {
		return fmt.Errorf(
			"unable to delete load balancer with vs prefix [%s], lbpool prefix [%s], ports [%v]: [%v]",
			virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, err)
	}
	rm := NewRDEManager(gm.Client, cgm.ClusterID)
	err = rm.removeVirtualIpFromRDE(ctx, rdeVIP)
	if err != nil {
		return fmt.Errorf("error when removing vip [%s] from RDE: [%v]", rdeVIP, err)
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
