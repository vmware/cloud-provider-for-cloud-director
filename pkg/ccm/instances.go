/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package ccm

import (
	"context"
	"fmt"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	cloudProvider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"runtime/debug"
	"strings"
)

type instances struct {
	vmInfoCache *VmInfoCache
}

func newInstances(vmInfoCache *VmInfoCache) cloudProvider.Instances {
	return &instances{vmInfoCache}
}

func getUUIDFromProviderID(providerID string) string {
	withoutPrefix := strings.TrimPrefix(providerID, ProviderName + "://")
	return strings.ToLower(strings.TrimSpace(withoutPrefix))
}

// NodeAddresses returns the addresses of the specified instance.
func (i *instances) NodeAddresses(ctx context.Context, nodeName types.NodeName) ([]v1.NodeAddress, error) {
	klog.Info("instances.NodeAddresses() called with ", string(nodeName))

	vmName := string(nodeName)
	vmInfo, err := i.vmInfoCache.GetByName(vmName)
	if err != nil {
		if err == govcd.ErrorEntityNotFound {
			return nil, cloudProvider.InstanceNotFound
		}

		return nil, fmt.Errorf("unable to find node addresses for [%s]: [%v]", vmName, err)
	}

	return vmInfo.Addresses, nil
}

// NodeAddressesByProviderID returns the addresses of the specified instance.
// The instance is specified using the providerID of the node. The
// ProviderID is a unique identifier of the node. This will not be called
// from the node whose node-addresses are being queried. i.e. local metadata
// services cannot be used in this method to obtain node-addresses
func (i *instances) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
	klog.Infof("instances.NodeAddressesByProviderID() called with [%s]", providerID)
	if providerID == "" {
		klog.Infof("Provider is empty: [%s]", string(debug.Stack()))
		return nil, fmt.Errorf("providerID is empty, hence node is uninitialized")
	}

	vmUUID := getUUIDFromProviderID(providerID)
	vmInfo, err := i.vmInfoCache.GetByUUID(vmUUID)
	if err != nil {
		if err == govcd.ErrorEntityNotFound {
			return nil, cloudProvider.InstanceNotFound
		}

		return nil, fmt.Errorf("unable to find node addresses by vm UUID for [%s]: [%v]", vmUUID, err)
	}

	return vmInfo.Addresses, nil
}

// InstanceID returns the cloud provider ID of the node with the specified NodeName.
// Note that if the instance does not exist, we must return ("", cloudprovider.InstanceNotFound)
// cloudprovider.InstanceNotFound should NOT be returned for instances that exist but are stopped/sleeping
func (i *instances) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
	klog.Infof("vcd.InstanceID() called with [%v]", nodeName)

	vmName := string(nodeName)
	vmInfo, err := i.vmInfoCache.GetByName(vmName)
	if err != nil {
		if err == govcd.ErrorEntityNotFound {
			return "", cloudProvider.InstanceNotFound
		}

		return "", fmt.Errorf("unable to find instance ID from vm name [%s]: [%v]", vmName, err)
	}

	return vmInfo.UUID, nil
}

// InstanceType returns the type of the specified instance.
func (i *instances) InstanceType(ctx context.Context, nodeName types.NodeName) (string, error) {
	klog.Infof("vcd.InstanceType() called with name [%v]", nodeName)

	vmName := string(nodeName)
	vmInfo, err := i.vmInfoCache.GetByName(vmName)
	if err != nil {
		if err == govcd.ErrorEntityNotFound {
			return "", cloudProvider.InstanceNotFound
		}

		return "", fmt.Errorf("unable to find instance type from vm name [%s]: [%v]", vmName, err)
	}

	return vmInfo.Type, nil
}

// InstanceTypeByProviderID returns the type of the specified instance.
func (i *instances) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
	klog.Infof("vcd.InstanceTypeByProviderID() called")

	vmUUID := getUUIDFromProviderID(providerID)
	vmInfo, err := i.vmInfoCache.GetByUUID(vmUUID)
	if err != nil {
		if err == govcd.ErrorEntityNotFound {
			return "", cloudProvider.InstanceNotFound
		}

		return "", fmt.Errorf("unable to find instance type from vm uuid [%s]: [%v]", vmUUID, err)
	}

	return vmInfo.Type, nil
}

// AddSSHKeyToAllInstances adds an SSH public key as a legal identity for all instances
// expected format for the key is standard ssh-keygen format: <protocol> <blob>
func (i *instances) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
	klog.Infof("vcd.AddSSHKeyToAllInstances() called")
	return cloudProvider.NotImplemented
}

// CurrentNodeName returns the name of the node we are currently running on
// On most clouds (e.g. GCE) this is the hostname, so we provide the hostname
func (i *instances) CurrentNodeName(ctx context.Context, hostName string) (types.NodeName, error) {
	klog.Infof("instances.CurrentNodeName() called with hostname [%s]", hostName)
	return types.NodeName(hostName), nil
}

// InstanceExistsByProviderID returns true if the instance for the given provider exists.
// If false is returned with no error, the instance will be immediately deleted by the cloud controller manager.
// This method should still return true for instances that exist but are stopped/sleeping.
func (i *instances) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	klog.Infof("instances.InstanceExistsByProviderID() called with provider ID [%s]", providerID)

	vmUUID := getUUIDFromProviderID(providerID)
	_, err := i.vmInfoCache.GetByUUID(vmUUID)
	if err != nil {
		if err == govcd.ErrorEntityNotFound {
			klog.Infof("instances.InstanceExistsByProviderID() for provider ID [%s] is false", providerID)
			return false, nil
		}

		klog.Infof("instances.InstanceExistsByProviderID() for provider ID [%s] is false due to error [%v]",
			providerID, err)
		return false, fmt.Errorf("unable to find instance from vm uuid [%s]: [%v]", vmUUID, err)
	}

	klog.Infof("instances.InstanceExistsByProviderID() for provider ID [%s] is true", providerID)
	return true, nil
}

// InstanceShutdownByProviderID returns true if the instance is shutdown in cloudprovider
func (i *instances) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	klog.Infof("instances.InstanceShutdownByProviderID() called with providerID [%s]", providerID)

	vmUUID := getUUIDFromProviderID(providerID)
	vmInfo, err := i.vmInfoCache.GetByUUID(vmUUID)
	if err != nil {
		if err == govcd.ErrorEntityNotFound {
			return false, nil
		}

		return false, fmt.Errorf("unable to find instance type from vm uuid [%s]: [%v]", vmUUID, err)
	}

	if vmInfo.vm == nil {
		return false, fmt.Errorf("vm struct in vmInfo nil for vm uuid [%s]: [%v]", vmUUID, err)
	}

	status, err := vmInfo.vm.GetStatus()
	if err != nil {
		if i.vmInfoCache.vcdClient.IsVmNotAvailable(err) {
			return true, nil
		}

		return false, fmt.Errorf("unable to get status of vm uuid [%s], name [%s]: [%v]",
			vmUUID, vmInfo.Name, err)
	}
	if status == "POWERED_ON" {
		return false, nil
	}

	klog.Infof("instances.InstanceShutdownByProviderID() for provider ID [%s] is true", providerID)
	return true, nil
}
