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
	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCName)
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
	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCName)
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

	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCName)
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
	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCName)
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

	_, err = vdcManager.AddNewVM(vmNamePrefix, vAppName, catalog,
		templateName, placementPolicyName, computePolicyName, "*", guestCustScript)
	require.NoError(t, err, "unable to create VM [%s]", vmNamePrefix)

	_ = vdcManager.WaitForGuestScriptCompletion(vAppName, vmNamePrefix)

	//err = VcdClient.DeleteVM(vAppName, vmName)
	//assert.NoError(t, err, "unable to delete VM")

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
	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCName)
	assert.NoError(t, err, "there should be no error when creating VDCManager object")

	vApp, err := vdcManager.GetOrCreateVApp(vAppName, vcdConfig.OvdcNetwork)
	assert.NoError(t, err, "unable to create vApp")
	require.NotNil(t, vApp, "vApp created should not be nil")

	// create vm
	vmName := "test-vm-1"

	// TODO: allow these vm params to be user passed through a config
	task, err := vdcManager.AddNewTkgVM(vmName, vAppName, "cse",
		"Ubuntu 20.04 and Kubernetes v1.23.10+vmware.1", "",
		"TKG medium", "Development2")
	assert.NoError(t, err, "unable to create vm [%s]", vmName)

	err = task.WaitTaskCompletion()
	assert.NoError(t, err, "should wait for VM creation task successfully")

	vm, err := vApp.GetVMByName(vmName, true)
	assert.NoError(t, err, "failed to get vm")

	// Set cloud-init config, boot the VM and check if the script has succeeded
	extraConfigMap := map[string]string{
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
