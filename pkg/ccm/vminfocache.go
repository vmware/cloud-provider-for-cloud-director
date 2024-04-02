/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package ccm

import (
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"k8s.io/klog"
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
	zm              *vcdsdk.ZoneMap
}

func newVmInfoCache(client *vcdsdk.Client, clusterVAppName string, expiry time.Duration, zm *vcdsdk.ZoneMap) *VmInfoCache {
	return &VmInfoCache{
		expiry:          expiry,
		nameMap:         make(map[string]*VmInfo),
		uuidMap:         make(map[string]*VmInfo),
		client:          client,
		clusterVAppName: clusterVAppName,
		zm:              zm,
	}
}

func (vmic *VmInfoCache) vmToVMInfo(vm *govcd.VM, ovdcIdentifier string, captureTime time.Time) (*VmInfo, error) {

	if vm == nil {
		return nil, fmt.Errorf("vm parameter should not be nil")
	}
	if vm.VM == nil {
		return nil, fmt.Errorf("vm.VM struct should not be nil")
	}

	vmInfo := &VmInfo{
		vm:        vm,
		OVDC:      ovdcIdentifier,
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

func (vmic *VmInfoCache) SearchVMAcrossVDCs(vmName string, vmId string) (*govcd.VM, string, error) {
	if err := vmic.client.RefreshBearerToken(); err != nil {
		return nil, "", fmt.Errorf("error while obtaining access token: [%v]", err)
	}

	var ovdcIdentifierList []string = nil
	var isMultiZoneCluster bool
	if vmic.zm != nil {
		for key, _ := range vmic.zm.VdcToZoneMap {
			ovdcIdentifierList = append(ovdcIdentifierList, key)
		}
		isMultiZoneCluster = true
	} else {
		ovdcIdentifierList = []string{
			vmic.client.ClusterOVDCIdentifier,
		}
		isMultiZoneCluster = false
	}

	orgManager := vcdsdk.OrgManager{
		Client:  vmic.client,
		OrgName: vmic.client.ClusterOrgName,
	}

	return orgManager.SearchVMAcrossVDCs(vmName, vmic.clusterVAppName, vmId, ovdcIdentifierList, isMultiZoneCluster)
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

	vm, ovdcIdentifier, err := vmic.SearchVMAcrossVDCs(vmName, "")
	if err != nil {
		return nil, fmt.Errorf("unable to find VM [%s] in org [%s] for cluster [%s]: [%v]",
			vmName, vmic.client.ClusterOrgName, vmic.clusterVAppName, err)
	}

	vmInfo, err := vmic.vmToVMInfo(vm, ovdcIdentifier, time.Now())
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
	vm, ovdcIdentifier, err := vmic.SearchVMAcrossVDCs("", vmUUID)
	if err != nil {
		return nil, fmt.Errorf("unable to find VM [%s] in org [%s] for cluster [%s]: [%v]",
			vmUUID, vmic.client.ClusterOrgName, vmic.clusterVAppName, err)
	}

	vmInfo, err := vmic.vmToVMInfo(vm, ovdcIdentifier, time.Now())
	if err != nil {
		return nil, fmt.Errorf("unable to convert vm struct [%v] to vmInfo: [%v]", vm, err)
	}

	vmic.nameMap[vmUUID] = vmInfo
	vmic.uuidMap[vm.VM.ID] = vmInfo

	return vmInfo, nil
}
