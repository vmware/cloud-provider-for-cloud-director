/*
   Copyright 2022 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package cpisdk

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/util"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"github.com/vmware/cloud-provider-for-cloud-director/release"
	"k8s.io/klog"
	"time"
)

const (
	// Errors
	CreateLoadbalancerError  = "CreateLoadbalancerError"
	UpdateLoadbalancerError  = "UpdateLoadbalancerError"
	DeleteLoadbalancerError  = "DeleteLoadbalancerError"
	GetLoadbalancerError     = "GetLoadbalancerError"
	CPIStatusUpgradeRdeError = "CPIStatusUpgradeRdeError"
	RemoveVIPFromRdeError    = "RemoveVirtualIPFromRdeError"
	AddVIPToRdeError         = "AddVirtualIPToRdeError"

	// Events
	ClientAuthenticated  = "ClientAuthenticated"
	CreatedLoadbalancer  = "CreatedLoadbalancer"
	UpdatedLoadbalancer  = "UpdatedLoadbalancer"
	DeletedLoadbalancer  = "DeletedLoadbalancer"
	CPIStatusRDEUpgraded = "CPIStatusRDEUpgraded"
)

// TODO(VCDA-3647): make RDE operations in CreateLoadBalancer, UpdateLoadBalancer, DeleteLoadBalancer and GetLoadBalancer
//   generic and move them to vcdsdk

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

	vcdsdkRdeManager := vcdsdk.NewRDEManager(gm.Client, cgm.ClusterID,
		release.CloudControllerManagerName, release.CpiVersion)
	cpiRdeManager := NewCPIRDEManager(vcdsdkRdeManager)

	resourceAllocationMap := &util.AllocatedResourcesMap{}

	// TODO: In the future, if we want component level errors we will need to pass additional params to below function to determine the caller
	// We could utilize a subErrorType in the additional details to handle different components (virtual service, DNAT, etc)
	externalIP, err := gm.CreateLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix,
		ips, portDetailsList, oneArm, resourceAllocationMap)

	if err != nil {
		addToErrorSetErr := AddToErrorSet(ctx, cpiRdeManager, CreateLoadbalancerError, cgm.ClusterID, err.Error())
		if addToErrorSetErr != nil {
			klog.Errorf("error adding CPI error to RDE: [%s], [%v]", cgm.ClusterID, addToErrorSetErr)
		}
		return "", fmt.Errorf(
			"unable to create load balancer with vs prefix [%s], lbpool prefix [%s], ips [%v], ports [%v]: [%v]",
			virtualServiceNamePrefix, lbPoolNamePrefix, ips, portDetailsList, err)
	}
	klog.Infof("Allocated resources during creation of load balancer with prefix [%s]: [%v]",
		virtualServiceNamePrefix, resourceAllocationMap)

	if externalIP != "" {
		// We can record this error so users can know that the RDE may not contain the external IP even though it exists
		err = cpiRdeManager.addVirtualIpToRDE(ctx, externalIP)
		if err != nil {
			addToErrorSetErr := AddToErrorSet(ctx, cpiRdeManager, AddVIPToRdeError, cgm.ClusterID, err.Error())
			if addToErrorSetErr != nil {
				klog.Errorf("error adding CPI error to RDE: [%s], [%v]", cgm.ClusterID, addToErrorSetErr)
			}
			klog.Errorf("error when adding virtual IP to RDE: [%v]", err)
		}
	}

	// No errors and external IP exists, we can remove the create LB error
	err = cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, CreateLoadbalancerError, cgm.ClusterID)
	if err != nil {
		klog.Errorf("there was an error removing CPI error [%s] from RDE [%s], [%v]", CreateLoadbalancerError, cgm.ClusterID, err)
	}

	err = AddToEventSet(ctx, cpiRdeManager, CreatedLoadbalancer, cgm.ClusterID, fmt.Sprintf("Created loadbalancer successfully for [%s] with external IP: [%s]", cgm.ClusterID, externalIP))
	if err != nil {
		klog.Errorf("error adding CPI event to RDE: [%v]", err)
	}

	return externalIP, nil
}

func (cgm *CpiGatewayManager) UpdateLoadBalancer(ctx context.Context, lbPoolName string, virtualServiceName string,
	ips []string, internalPort int32, externalPort int32) error {

	// TODO: If OneArm is not supplied, CreateLoadBalancer() won't create a DNAT rule and App port profile. Add checks
	//		for oneArm to prevent errors.

	gm := cgm.VcdGatewayManager
	if gm == nil {
		return fmt.Errorf("GatewayManager cannot be nil")
	}

	vcdsdkRdeManager := vcdsdk.NewRDEManager(gm.Client, cgm.ClusterID,
		release.CloudControllerManagerName, release.CpiVersion)
	cpiRdeManager := NewCPIRDEManager(vcdsdkRdeManager)

	if err := gm.UpdateLoadBalancer(ctx, lbPoolName, virtualServiceName, ips, internalPort, externalPort); err != nil {
		addToErrorSetErr := AddToErrorSet(ctx, cpiRdeManager, UpdateLoadbalancerError, virtualServiceName, err.Error())
		if addToErrorSetErr != nil {
			klog.Errorf("error adding CPI error to RDE: [%s], [%v]", cgm.ClusterID, addToErrorSetErr)
		}
		return fmt.Errorf(
			"unable to update load balancer with vs [%s], lbpool [%s], ips [%v], ports [%d->%d]: [%v]",
			virtualServiceName, lbPoolName, ips, externalPort, internalPort, err)
	}

	err := cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, UpdateLoadbalancerError, virtualServiceName)
	if err != nil {
		klog.Errorf("there was an error removing CPI error [%s] from RDE [%s], [%v]", UpdateLoadbalancerError, cgm.ClusterID, err)
	}

	err = AddToEventSet(ctx, cpiRdeManager, UpdatedLoadbalancer, virtualServiceName, fmt.Sprintf("Successfully updated loadbalancer with virtual service name [%s]: ", virtualServiceName))
	if err != nil {
		klog.Errorf("error adding CPI event to RDE: [%s], [%v]", cgm.ClusterID, err)
	}
	return nil
}

func (cgm *CpiGatewayManager) DeleteLoadBalancer(ctx context.Context, virtualServiceNamePrefix string,
	lbPoolNamePrefix string, portDetailsList []vcdsdk.PortDetails, oneArm *vcdsdk.OneArm) error {
	gm := cgm.VcdGatewayManager
	if gm == nil {
		return fmt.Errorf("GatewayManager cannot be nil")
	}

	vcdsdkRdeManager := vcdsdk.NewRDEManager(gm.Client, cgm.ClusterID,
		release.CloudControllerManagerName, release.CpiVersion)
	cpiRdeManager := NewCPIRDEManager(vcdsdkRdeManager)

	rdeVIP, err := gm.DeleteLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, oneArm)
	if err != nil {
		addToErrorSetErr := AddToErrorSet(ctx, cpiRdeManager, DeleteLoadbalancerError, cgm.ClusterID, err.Error())
		if addToErrorSetErr != nil {
			klog.Errorf("error adding CPI error to RDE: [%v]", addToErrorSetErr)
		}

		return fmt.Errorf(
			"unable to delete load balancer with vs prefix [%s], lbpool prefix [%s], ports [%v]: [%v]",
			virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, err)
	}

	err = cpiRdeManager.removeVirtualIpFromRDE(ctx, rdeVIP)
	if err != nil {
		addToErrorSetErr := AddToErrorSet(ctx, cpiRdeManager, RemoveVIPFromRdeError, cgm.ClusterID, err.Error())
		if addToErrorSetErr != nil {
			klog.Errorf("unable to add CPI error to RDE: [%s], [%v]", cgm.ClusterID, addToErrorSetErr)
		}

		return fmt.Errorf("error when removing vip [%s] from RDE: [%v]", rdeVIP, err)
	}

	err = vcdsdkRdeManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, DeleteLoadbalancerError, cgm.ClusterID)
	if err != nil {
		klog.Errorf("there was an error removing CPI error [%s] from RDE [%s], [%v]", DeleteLoadbalancerError, cgm.ClusterID, err)
	}

	err = AddToEventSet(ctx, cpiRdeManager, DeletedLoadbalancer, cgm.ClusterID, fmt.Sprintf("Successfully deleted loadbalancer associated with [%s], deleted external IP [%s]", cgm.ClusterID, rdeVIP))
	if err != nil {
		klog.Errorf("error adding CPI event to RDE: [%v]", err)
	}
	return nil
}

func (cgm *CpiGatewayManager) GetLoadBalancer(ctx context.Context, virtualServiceName string, oneArm *vcdsdk.OneArm) (string, error) {
	gm := cgm.VcdGatewayManager
	if gm == nil {
		return "", fmt.Errorf("GatewayManager cannot be nil")
	}

	rdeManager := vcdsdk.NewRDEManager(gm.Client, cgm.ClusterID,
		release.CloudControllerManagerName, release.CpiVersion)
	cpiRdeManager := NewCPIRDEManager(rdeManager)

	extIP, err := gm.GetLoadBalancer(ctx, virtualServiceName, oneArm)
	if err != nil {
		addToErrorSetErr := AddToErrorSet(ctx, cpiRdeManager, GetLoadbalancerError, virtualServiceName, err.Error())
		if addToErrorSetErr != nil {
			klog.Errorf("unable to add CPI error to RDE [%v]", addToErrorSetErr)
		}
	}

	removeErr := rdeManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, GetLoadbalancerError, virtualServiceName)
	if removeErr != nil {
		klog.Errorf("there was an error removing CPI error [%s] from RDE [%s], [%v]", GetLoadbalancerError, cgm.ClusterID, err)
	}

	return extIP, err
}

func (cgm *CpiGatewayManager) GetLoadBalancerPool(ctx context.Context, lbPoolName string) (*swaggerClient.EntityReference, error) {
	gm := cgm.VcdGatewayManager
	if gm == nil {
		return nil, fmt.Errorf("GatewayManager cannot be nil")
	}
	return gm.GetLoadBalancerPool(ctx, lbPoolName)
}

func AddToErrorSet(ctx context.Context, cpiRdeManager *CPIRDEManager, errorName, vcdResourceId, detailedErrorMessage string) error {
	backendErr := vcdsdk.BackendError{
		Name:              errorName,
		OccurredAt:        time.Now(),
		VcdResourceId:     vcdResourceId,
		AdditionalDetails: map[string]interface{}{"Detailed Error": detailedErrorMessage},
	}
	return cpiRdeManager.RDEManager.AddToErrorSet(ctx, vcdsdk.ComponentCPI, backendErr, vcdsdk.DefaultRollingWindowSize)
}

func AddToEventSet(ctx context.Context, cpiRdeManager *CPIRDEManager, eventName, vcdResourceId, detailedErrorMessage string) error {
	backendEvent := vcdsdk.BackendEvent{
		Name:              eventName,
		OccurredAt:        time.Now(),
		VcdResourceId:     vcdResourceId,
		AdditionalDetails: map[string]interface{}{"Detailed Event": detailedErrorMessage},
	}
	return cpiRdeManager.RDEManager.AddToEventSet(ctx, vcdsdk.ComponentCPI, backendEvent, vcdsdk.DefaultRollingWindowSize)
}
