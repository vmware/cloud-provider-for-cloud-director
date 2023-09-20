/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package ccm

import (
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"k8s.io/klog"
	"strings"
	"sync"
	"time"

	"github.com/vmware/go-vcloud-director/v2/govcd"
	v1 "k8s.io/api/core/v1"
	v1helper "k8s.io/cloud-provider/node/helpers"
)

type VmInfo struct {
	vm        *govcd.VM
	OVDC      string
	Zone      string
	UUID      string
	Name      string
	Type      string
	Addresses []v1.NodeAddress
	TimeStamp time.Time
}

// VmInfoCache caches VM details. Ideally we need an LRU cache with ttl-based expiry. But since we have O(1k) nodes
// per cluster, we can ignore limits and expiry.
type VmInfoCache struct {
	rwLock          sync.RWMutex
	expiry          time.Duration
	nameMap         map[string]*VmInfo
	uuidMap         map[string]*VmInfo
	client          *vcdsdk.Client
	clusterVAppName string
}

func newVmInfoCache(client *vcdsdk.Client, clusterVAppName string, expiry time.Duration) *VmInfoCache {
	return &VmInfoCache{
		expiry:          expiry,
		nameMap:         make(map[string]*VmInfo),
		uuidMap:         make(map[string]*VmInfo),
		client:          client,
		clusterVAppName: clusterVAppName,
	}
}

func (vmic *VmInfoCache) vmToVMInfo(vm *govcd.VM, ovdc string, captureTime time.Time) (*VmInfo, error) {

	if vm == nil {
		return nil, fmt.Errorf("vm parameter should not be nil")
	}
	if vm.VM == nil {
		return nil, fmt.Errorf("vm.VM struct should not be nil")
	}

	vmInfo := &VmInfo{
		vm:        vm,
		OVDC:      ovdc,
		UUID:      vm.VM.ID,
		Name:      vm.VM.Name,
		Type:      "",
		TimeStamp: captureTime,
	}

	if vm.VM.NetworkConnectionSection != nil {
		vmAddresses := make([]v1.NodeAddress, 0)
		for _, netConn := range vm.VM.NetworkConnectionSection.NetworkConnection {
			if netConn.NetworkConnectionIndex == vm.VM.NetworkConnectionSection.PrimaryNetworkConnectionIndex {
				v1helper.AddToNodeAddresses(&vmAddresses,
					v1.NodeAddress{
						Type:    v1.NodeInternalIP,
						Address: netConn.IPAddress,
					})
			}
			v1helper.AddToNodeAddresses(&vmAddresses,
				v1.NodeAddress{
					Type:    v1.NodeExternalIP,
					Address: netConn.IPAddress,
				},
				v1.NodeAddress{
					Type:    v1.NodeHostName,
					Address: vm.VM.Name,
				})
		}
		vmInfo.Addresses = vmAddresses
	}

	return vmInfo, nil
}

func CreateVAppNamePrefix(clusterName string, ovdcID string) (string, error) {
	parts := strings.Split(ovdcID, ":")
	if len(parts) != 4 {
		// urn:vcloud:org:<uuid>
		return "", fmt.Errorf("invalid URN format for OVDC: [%s]", ovdcID)
	}

	return fmt.Sprintf("%s_%s", clusterName, parts[3]), nil
}

func (vmic *VmInfoCache) SearchVMAcrossVDCs(vmName string, vmId string) (*govcd.VM, string, error) {
	if err := vmic.client.RefreshBearerToken(); err != nil {
		return nil, "", fmt.Errorf("error while obtaining access token: [%v]", err)
	}

	org, err := vmic.client.VCDClient.GetOrgByName(vmic.client.ClusterOrgName)
	if err != nil {
		return nil, "", fmt.Errorf("unable to get org by name [%s]: [%v]", vmic.client.ClusterOrgName, err)
	}

	ovdcList, err := org.QueryOrgVdcList()
	if err != nil {
		return nil, "", fmt.Errorf("unable to get list of OVDCs from org [%s]: [%v]",
			vmic.client.ClusterOrgName, err)
	}

	for _, ovdcRecordType := range ovdcList {
		klog.Infof("Looking for VM [name:%s],[id:%s] of cluster [%s] in OVDC [%s]",
			vmName, vmId, vmic.clusterVAppName, ovdcRecordType.Name)
		vdc, err := org.GetVDCByName(ovdcRecordType.Name, false)
		if err != nil {
			klog.Infof("unable to query VDC [%s] in Org [%s] by name: [%v]",
				ovdcRecordType.Name, vmic.client.ClusterOrgName, err)
			continue
		}
		vAppNamePrefix, err := CreateVAppNamePrefix(vmic.clusterVAppName, vdc.Vdc.ID)
		if err != nil {
			klog.Infof("Unable to create a vApp name prefix for cluster [%s] in OVDC [%s] with OVDC ID [%s]: [%v]",
				vmic.clusterVAppName, vdc.Vdc.Name, vdc.Vdc.ID, err)
			continue
		}
		klog.Infof("Looking for vApps with a prefix of [%s]", vAppNamePrefix)
		vAppList := vdc.GetVappList()
		// check if the VM exists in any cluster-vApps in this OVDC
		for _, vApp := range vAppList {
			if strings.HasPrefix(vApp.Name, vAppNamePrefix) {
				// check if VM exists
				klog.Infof("Looking for VM [name:%s],[id:%s] in vApp [%s] in OVDC [%s] is a vApp in cluster [%s]",
					vmName, vmId, vApp.Name, vdc.Vdc.Name, vmic.clusterVAppName)
				vdcManager, err := vcdsdk.NewVDCManager(vmic.client, vmic.client.ClusterOrgName, vdc.Vdc.Name)
				if err != nil {
					return nil, "", fmt.Errorf("error creating VDCManager object for VDC [%s]: [%v]",
						vdc.Vdc.Name, err)
				}

				var vm *govcd.VM = nil
				if vmName != "" {
					vm, err = vdcManager.FindVMByName(vApp.Name, vmName)
				} else if vmId != "" {
					vm, err = vdcManager.FindVMByUUID(vApp.Name, vmId)
				} else {
					return nil, "", fmt.Errorf("either vm name [%s] or ID [%s] should be passed", vmName, vmId)
				}
				if err != nil {
					klog.Infof("Could not find VM [name:%s],[id:%s] in vApp [%s] of Cluster [%s] in OVDC [%s]: [%v]",
						vmName, vmId, vApp.Name, vmic.clusterVAppName, vdc.Vdc.Name, err)
					continue
				}

				// If we reach here, we found the VM
				klog.Infof("Found VM [name:%s],[id:%s] in vApp [%s] of Cluster [%s] in OVDC [%s]: [%v]",
					vmName, vmId, vApp.Name, vmic.clusterVAppName, vdc.Vdc.Name, err)
				return vm, vdc.Vdc.Name, nil
			}
		}
		klog.Infof("Could not find VM [name:%s],[id:%s] of cluster [%s] in OVDC [%s]",
			vmName, vmId, vmic.clusterVAppName, ovdcRecordType.Name)
	}

	return nil, "", govcd.ErrorEntityNotFound
}

func (vmic *VmInfoCache) GetByName(vmName string) (*VmInfo, error) {
	vmic.rwLock.Lock()
	defer vmic.rwLock.Unlock()

	now := time.Now()
	if vmInfo, ok := vmic.nameMap[vmName]; ok {
		if now.Sub(vmInfo.TimeStamp) < vmic.expiry {
			klog.Infof("now is [%v], info captured at [%v], hence reusing cached value", now, vmInfo.TimeStamp)
			return vmInfo, nil
		}

		// don't use old values
		delete(vmic.nameMap, vmName)
	}

	vm, ovdcName, err := vmic.SearchVMAcrossVDCs(vmName, "")
	if err != nil {
		return nil, fmt.Errorf("unable to find VM [%s] in org [%s] for cluster [%s]: [%v]",
			vmName, vmic.client.ClusterOrgName, vmic.clusterVAppName, err)
	}

	vmInfo, err := vmic.vmToVMInfo(vm, ovdcName, time.Now())
	if err != nil {
		return nil, fmt.Errorf("unable to convert vm struct [%v] to vmInfo: [%v]", vm, err)
	}

	vmic.nameMap[vmName] = vmInfo
	vmic.uuidMap[vm.VM.ID] = vmInfo

	return vmInfo, nil
}

func (vmic *VmInfoCache) GetByUUID(vmUUID string) (*VmInfo, error) {
	vmic.rwLock.Lock()
	defer vmic.rwLock.Unlock()

	now := time.Now()
	if vmInfo, ok := vmic.uuidMap[vmUUID]; ok {
		if now.Sub(vmInfo.TimeStamp) < vmic.expiry {
			klog.Infof("now is [%v], info captured at [%v], hence reusing cached value", now, vmInfo.TimeStamp)
			return vmInfo, nil
		}

		// don't use old values
		delete(vmic.nameMap, vmUUID)
	}
	vm, ovdcName, err := vmic.SearchVMAcrossVDCs("", vmUUID)
	if err != nil {
		return nil, fmt.Errorf("unable to find VM [%s] in org [%s] for cluster [%s]: [%v]",
			vmUUID, vmic.client.ClusterOrgName, vmic.clusterVAppName, err)
	}

	vmInfo, err := vmic.vmToVMInfo(vm, ovdcName, time.Now())
	if err != nil {
		return nil, fmt.Errorf("unable to convert vm struct [%v] to vmInfo: [%v]", vm, err)
	}

	vmic.nameMap[vmUUID] = vmInfo
	vmic.uuidMap[vm.VM.ID] = vmInfo

	return vmInfo, nil
}
