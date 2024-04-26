/*
Copyright 2021 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/
package vcdsdk

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestVApp(t *testing.T) {

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	// get client
	vcdClient, err := getTestVCDClient(vcdConfig, nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	// create vApp
	vAppName := "test-vapp"
	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCIdentifier)
	assert.NoError(t, err, "error creating VDCManager")

	// create VApp
	vappObj, err := vdcManager.GetOrCreateVApp(vAppName, vcdConfig.OvdcNetwork)
	assert.NoError(t, err, "error creating VApp")
	assert.NotNil(t, vappObj, "vApp created should not be nil")

	// create VM
	vmNamePrefix := "test-vm"
	task, err := vdcManager.AddNewTkgVM(vmNamePrefix, vAppName, "cse",
		"Ubuntu 20.04 and Kubernetes v1.21.11+vmware.1", "", "4core4gb",
		"")
	assert.NoError(t, err, "should create vm correctly")

	err = task.WaitTaskCompletion()
	assert.NoError(t, err, "should wait for task successfully")

	vms, err := vdcManager.FindAllVMsInVapp(vAppName)
	assert.NoError(t, err, "unable to find VMs in vApp")
	assert.NotNil(t, vms, "some VMs should be returned")

	return
}

func TestDeleteVapp(t *testing.T) {

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")
	// get client
	vcdClient, err := getTestVCDClient(vcdConfig, nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	vappName := "vapp1"
	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCIdentifier)
	assert.NoError(t, err, "there should be no error in creating VDCManager object")
	vapp, err := vdcManager.GetOrCreateVApp(vappName, vcdConfig.OvdcNetwork)
	assert.NoError(t, err, "unable to find vApp")
	assert.NotNil(t, vapp, "vapp should not be nil")

	// delete VApp
	err = vdcManager.DeleteVApp(vappName)
	assert.NoError(t, err, "unable to delete vApp")
}

func TestVdcManager_CacheVdcDetails(t *testing.T) {
	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := os.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := getTestVCDClient(vcdConfig, nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCIdentifier)
	err = vdcManager.cacheVdcDetails()
	assert.NoError(t, err, "There should no error while caching VDC details")
}

func TestVMCreation(t *testing.T) {

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	// get client
	vcdClient, err := getTestVCDClient(vcdConfig, nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	// create vApp
	vAppName := "test-vapp"
	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCIdentifier)
	assert.NoError(t, err, "there should be no error when creating VDCManager object")

	vApp, err := vdcManager.GetOrCreateVApp(vAppName, vcdConfig.OvdcNetwork)
	assert.NoError(t, err, "unable to create vApp")
	require.NotNil(t, vApp, "vApp created should not be nil")

	// create vm
	vmNamePrefix := "test-vm-1"
	guestCustScript := `
#!/usr/bin/env bash -x
echo "Command called with arguments [$@] at time $(date)" >> /root/output.txt
if [ -f "/.guest-customization-post-reboot-pending" ]
then
	echo "Reboot pending, hence will do nothing."  >> /root/output.txt
elif [ "$1" = "postcustomization" ]
then
	kubeadm init --ttl=0
	join_token = $(kubeadm token create --print-join-command)
	vmtoolsd --cmd "info-set guestinfo.joinToken "${join_token}"
else
	echo "Skipping script since postcustomization is not involved." >> /root/output.txt
fi
exit 0
`

	// TODO: read the following variables from testdata
	catalog := "cse"
	templateName := "ubuntu-16.04_k8-1.20_weave-2.6.5_rev2"
	placementPolicyName := "cse----native"
	computePolicyName := "2core2gb"

	task, err := vdcManager.AddNewVM(vmNamePrefix, vAppName, catalog,
		templateName, placementPolicyName, computePolicyName, "*", guestCustScript)
	require.NoError(t, err, "unable to create VM [%s]", vmNamePrefix)

	err = task.WaitTaskCompletion()
	assert.NoError(t, err, "create VM task failed")

	assert.NoError(t, vApp.Refresh(), "failed to refresh vApp")

	vm, err := vApp.GetVMByName(vmNamePrefix, false)
	assert.NoError(t, err, "failed to get VM by name")

	task, err = vm.PowerOn()
	assert.NoError(t, err, "error occurred while powering on the VM")

	err = task.WaitTaskCompletion()
	assert.NoError(t, err, "error occurred while waiting for the power on task to complete")

	_ = vdcManager.WaitForGuestScriptCompletion(vAppName, vmNamePrefix)

	task, err = vdcManager.DeleteVM(vAppName, vmNamePrefix)
	assert.NoError(t, err, "unable to delete VM")
	assert.NotEmpty(t, task, "obtained an empty task")

	err = task.WaitTaskCompletion()
	assert.NoError(t, err, "task for delete VM failed")

	return
}

func TestSearchVMAcrossVDCs(t *testing.T) {

	vcdConfig, err := getTestVCDConfig()
	require.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vdcName1 := "tenant1_ovdc"
	vdcName2 := "tenant1_ovdc2"

	// get client
	vcdClient, err := getTestVCDClient(vcdConfig, nil)
	require.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	vdcManager1, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vdcName1)
	require.NoError(t, err, "there should be no error when creating VDCManager object")
	vdcManager2, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vdcName2)
	require.NoError(t, err, "there should be no error when creating VDCManager object")

	// create vApps
	vAppNamePrefix := "arun-vapp"
	catalogName := "cse"
	templateName := "Ubuntu 20.04 and Kubernetes v1.25.7+vmware.2"

	vAppName_vdc1 := vAppNamePrefix + "_" + strings.Split(vdcManager1.Vdc.Vdc.ID, ":")[3] + "_nodepool1"
	vApp1, err := vdcManager1.GetOrCreateVApp(vAppName_vdc1, vcdConfig.OvdcNetwork)
	require.NoError(t, err, "unable to create vApp")
	require.NotNil(t, vApp1, "vApp created should not be nil")
	defer func() {
		_ = vdcManager1.DeleteVApp(vAppName_vdc1)
	}()

	vAppName_vdc2 := vAppNamePrefix + "_" + strings.Split(vdcManager2.Vdc.Vdc.ID, ":")[3] + "_nodepool2"
	vApp2, err := vdcManager2.GetOrCreateVApp(vAppName_vdc2, vcdConfig.OvdcNetwork)
	require.NoError(t, err, "unable to create vApp")
	require.NotNil(t, vApp2, "vApp created should not be nil")
	defer func() {
		_ = vdcManager2.DeleteVApp(vAppName_vdc2)
	}()

	// create vm
	vmName_vdc1 := "arun-vm-1"
	// TODO: allow these vm params to be user passed through a config
	task, err := vdcManager1.AddNewTkgVM(vmName_vdc1, vAppName_vdc1, catalogName,
		templateName, "", "", "Development2")
	require.NoError(t, err, "unable to create vm [%s]", vmName_vdc1)
	defer func() {
		_, _ = vdcManager1.DeleteVM(vAppName_vdc1, vmName_vdc1)
	}()

	err = task.WaitTaskCompletion()
	require.NoError(t, err, "should wait for VM creation task successfully")

	vmName_vdc2 := "arun-vm-2"
	// TODO: allow these vm params to be user passed through a config
	task, err = vdcManager2.AddNewTkgVM(vmName_vdc2, vAppName_vdc2, catalogName,
		templateName, "", "", "Development2")
	require.NoError(t, err, "unable to create vm [%s]", vmName_vdc2)
	defer func() {
		_, _ = vdcManager2.DeleteVM(vAppName_vdc2, vmName_vdc2)
	}()

	err = task.WaitTaskCompletion()
	require.NoError(t, err, "should wait for VM creation task successfully")

	orgManager, err := NewOrgManager(vcdClient, vcdClient.ClusterOrgName)
	require.NoError(t, err, "there should be no error when creating OrgManager object")

	vm1, vdcNameReceived, err := orgManager.SearchVMAcrossVDCs(vmName_vdc1, vAppNamePrefix, "", true)
	require.NoError(t, err, "vm named [%s] should be found", vmName_vdc1)
	require.Equal(t, vdcName1, vdcNameReceived)
	require.Equal(t, vm1.VM.Name, vmName_vdc1)

	vAppFound1, err := vm1.GetParentVApp()
	require.NoError(t, err, "parent of vm should be found", "vm", vm1.VM.Name)
	require.Equal(t, vAppFound1.VApp.Name, vAppName_vdc1)

	vmUsingId1, _, err := orgManager.SearchVMAcrossVDCs("", vAppNamePrefix, vm1.VM.ID, true)
	require.NoError(t, err, "vm named [%s] should be found", vmName_vdc1)
	require.Equal(t, vmUsingId1.VM.ID, vm1.VM.ID)

	vm2, vdcNameReceived, err := orgManager.SearchVMAcrossVDCs(vmName_vdc2, vAppNamePrefix, "", true)
	require.NoError(t, err, "vm named [%s] should be found", vmName_vdc2)
	require.Equal(t, vdcName2, vdcNameReceived)
	require.Equal(t, vm2.VM.Name, vmName_vdc2)

	vAppFound2, err := vm2.GetParentVApp()
	require.NoError(t, err, "parent of vm [%s] should be found", vm1.VM.Name)
	require.Equal(t, vAppFound2.VApp.Name, vAppName_vdc2)

	vmNonMultiAZ, vdcNonMultiAZ, err := orgManager.SearchVMAcrossVDCs(vmName_vdc1, vAppName_vdc1, "", false)
	require.NoError(t, err, "vm named [%s] should be found", vmName_vdc1)
	require.Equal(t, vmNonMultiAZ.VM.Name, vmName_vdc1)
	require.Equal(t, vdcNonMultiAZ, vdcName1)

	return
}

func TestVMExtraConfig(t *testing.T) {

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	// get client
	vcdClient, err := getTestVCDClient(vcdConfig, nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	// create vApp
	vAppName := "test-vapp"
	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCIdentifier)
	assert.NoError(t, err, "there should be no error when creating VDCManager object")

	vApp, err := vdcManager.GetOrCreateVApp(vAppName, vcdConfig.OvdcNetwork)
	assert.NoError(t, err, "unable to create vApp")
	require.NotNil(t, vApp, "vApp created should not be nil")

	// create vm
	vmName := "test-vm-1"

	// TODO: allow these vm params to be user passed through a config
	task, err := vdcManager.AddNewTkgVM(vmName, vAppName, "cse",
		"Ubuntu 20.04 and Kubernetes v1.23.10+vmware.1", "",
		"", "Development2")
	assert.NoError(t, err, "unable to create vm [%s]", vmName)

	err = task.WaitTaskCompletion()
	assert.NoError(t, err, "should wait for VM creation task successfully")

	vm, err := vApp.GetVMByName(vmName, true)
	assert.NoError(t, err, "failed to get vm")

	// Set cloud-init config, boot the VM and check if the script has succeeded
	extraConfigMap := map[string]string{
		"guestinfo.sometest": strings.Repeat("aaaaa", 65536),
		"guestinfo.userdata": b64.StdEncoding.EncodeToString([]byte(`
#cloud-config
users:
- name: core
  groups: [ sudo ]
  shell: /bin/bash
  lock_passwd: false
  chpasswd: { expire: False }
  passwd: '$6$rounds=4096$JjhnQS5RRn$pAACOIBbau4c1RkioIP/FqVPTpjBpzXnfeBGnmbMwl0ZciubHTiZVy.jc2R2rwG7MriCWyeRsJIFArPhZoNod1'
write_files:
- path: /root/test.sh
  owner: root
  content: |
    #!/bin/bash
    set -eEx
    ping -c 4 8.8.8.8
    nslookup www.google.com
    ping -c 4 www.google.com
    vmtoolsd --cmd "info-set guestinfo.test test123"
runcmd:
- bash /root/test.sh
`)),
		"guestinfo.userdata.encoding": "base64",
	}

	task, err = vdcManager.SetMultiVmExtraConfigKeyValuePairs(vm, extraConfigMap, false)
	assert.NoError(t, err, "error setting VM extra config [%s]: [%v]", extraConfigMap, err)
	err = task.WaitTaskCompletion()
	assert.NoError(t, err, "should wait for task to set extra config on VM [%s] successfully", vm.VM.Name)

	task, err = vm.PowerOn()
	assert.NoError(t, err, "Unable to power on VM [%s]: [%v]", vm.VM.Name, err)
	err = task.WaitTaskCompletion()
	assert.NoError(t, err, "should wait for task to power on VM [%s] successfully", vm.VM.Name)

	err = vdcManager.WaitForGuestScriptCompletion(vAppName, vmName)
	assert.NoError(t, err, "should wait for guest script completion on VM [%s] successfully", vm.VM.Name)

	// wait until VM is powered on, run the cloud-init script and sets the keys and values
	key := "guestinfo.test"
	value := "test123"
	retrievedValue := ""
	success := false
	endTime := time.Now().Add(10 * time.Minute)
	for currTime := time.Now(); currTime.Before(endTime); currTime = time.Now() {
		err = vm.Refresh()
		assert.NoError(t, err, "unable to refresh vm [%s]", vm.VM.Name)
		retrievedValue, err = vdcManager.GetExtraConfigValue(vm, key)
		assert.NoError(t, err, "failed to get hardware section value for key %s", key)
		fmt.Printf("expected value is [%s], obtained value is [%s]\n", value, retrievedValue)
		if retrievedValue == value {
			success = true
			break
		}
		time.Sleep(30 * time.Second)
	}
	assert.True(t, success, "Retrieved value [%s], expected value [%s]", retrievedValue, value)

	err = vdcManager.DeleteVApp(vAppName)
	assert.NoError(t, err, "unable to delete vApp: [%s]", vAppName)
}
