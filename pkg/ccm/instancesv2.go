package ccm

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	v1 "k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog/v2"
)

type instancesV2 struct {
	vmInfoCache *VmInfoCache
	zoneMap     *vcdsdk.ZoneMap
}

func (vcdCP *VCDCloudProvider) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return vcdCP.instancesV2, true
}

func newInstancesV2(vmInfoCache *VmInfoCache, zoneMap *vcdsdk.ZoneMap) cloudprovider.InstancesV2 {
	return instancesV2{
		vmInfoCache: vmInfoCache,
		zoneMap:     zoneMap,
	}
}

// InstanceExists returns true if the instance for the given node exists according to the cloud provider.
// Use the node.name or node.spec.providerID field to find the node in the cloud provider.
func (i instancesV2) InstanceExists(ctx context.Context, node *v1.Node) (bool, error) {
	vmName := node.Name
	vmId := node.Spec.ProviderID
	klog.Infof("instances.InstanceShutdown called with vmName [%s], vmId [%s]", vmName, vmId)
	switch {
	case vmName != "":
		_, err := i.vmInfoCache.GetByName(vmName)
		if err != nil {
			if err == govcd.ErrorEntityNotFound {
				return false, cloudprovider.InstanceNotFound
			}
			return false, fmt.Errorf("unable to find instance ID from vm name [%s]: [%v]", vmName, err)
		}
		return true, nil

	case vmId != "":
		_, err := i.vmInfoCache.GetByUUID(vmId)
		if err != nil {
			if err == govcd.ErrorEntityNotFound {
				return false, cloudprovider.InstanceNotFound
			}
			return false, fmt.Errorf("unable to find instance ID from vm Id [%s]: [%v]", vmId, err)
		}
		return true, nil

	default:
		return false, fmt.Errorf("both vm name and vm UUID are not specified")
	}
}

// InstanceShutdown returns true if the instance is shutdown according to the cloud provider.
// Use the node.name or node.spec.providerID field to find the node in the cloud provider.
func (i instancesV2) InstanceShutdown(ctx context.Context, node *v1.Node) (bool, error) {
	providerID := node.Spec.ProviderID
	klog.Infof("instances.InstanceShutdown called with providerID [%s]", providerID)

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
		vdcManager, vdcManagerErr := vcdsdk.NewVDCManager(i.vmInfoCache.client, i.vmInfoCache.client.ClusterOrgName,
			i.vmInfoCache.client.ClusterOVDCName)
		if vdcManagerErr != nil {
			return false, fmt.Errorf("error creating VDC manager object: [%v]", err)
		}
		if vdcManager.IsVmNotAvailable(err) {
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

// InstanceMetadata returns the instance's metadata. The values returned in InstanceMetadata are
// translated into specific fields and labels in the Node object on registration.
// Implementations should always check node.spec.providerID first when trying to discover the instance
// for a given node. In cases where node.spec.providerID is empty, implementations can use other
// properties of the node like its name, labels and annotations.
func (i instancesV2) InstanceMetadata(ctx context.Context, node *v1.Node) (*cloudprovider.InstanceMetadata, error) {
	providerID := node.Spec.ProviderID
	klog.Infof("instancesV2.InstanceMetadata called with name [%s], providerID [%s]", node.Name, providerID)

	var vmInfo *VmInfo = nil
	var err error
	if providerID != "" {
		vmUUID := getUUIDFromProviderID(providerID)
		vmInfo, err = i.vmInfoCache.GetByUUID(vmUUID)
		if err != nil {
			return nil, fmt.Errorf("error trying to find VM with name [%s], UUID [%s] by UUID: [%v]",
				node.Name, vmUUID, err)
		}
	} else {
		vmInfo, err = i.vmInfoCache.GetByName(node.Name)
		if err != nil {
			return nil, fmt.Errorf("error trying to fingd VM with name [%s], UUID [%s] by UUID: [%v]",
				node.Name, "", err)
		}
	}

	zone := ""
	if i.zoneMap != nil {
		zone, _ = i.zoneMap.VdcToZoneMap[vmInfo.OVDC]
	}

	instanceMetadata := &cloudprovider.InstanceMetadata{
		ProviderID:    getProviderIDFromUUID(vmInfo.UUID),
		InstanceType:  vmInfo.Type,
		NodeAddresses: vmInfo.Addresses,
		Zone:          zone,
		Region:        "",
	}

	klog.Infof("reporting instanceV2 Metadata for vm [%s] as [%#v]", vmInfo.Name, instanceMetadata)

	return instanceMetadata, nil
}
