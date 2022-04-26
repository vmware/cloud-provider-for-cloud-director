/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package vcdsdk

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestVMCreation(t *testing.T) {

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	// get client
	vcdClient, err := GetTestVCDClient(
		cloudConfig,
		map[string]interface{}{
			"getVdcClient": true,
		})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	// create vApp
	vAppName := "test-vapp"
	vdcManager, err := NewVDCManager(vcdClient)
	assert.NoError(t, err, "There should be no error when creating VDCManager")

	vApp, err := vdcManager.GetOrCreateVApp(vAppName, "tenant1_ovdc_nw")
	assert.NoError(t, err, "unable to create vApp")
	require.NotNil(t, vApp, "vApp created should not be nil")

	// create vm
	vmNamePrefix := "test-vm-1"
	vmNum := 1
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
	vmm := NewVMManager(vcdClient, vAppName)
	_, err = vmm.AddNewMultipleVM(vApp, vmNamePrefix, vmNum, "ProviderCatalogs",
		"ubuntu-16.04_k8-1.20_weave-2.6.5_rev2",
		"cse----native", "2core2gb", guestCustScript, true, true)
	require.NoError(t, err, "unable to create [%d] VMs", vmNum)

	_ = vmm.WaitForGuestScriptCompletion(vmNamePrefix, vAppName)

	//err = VcdClient.DeleteVM(vAppName, vmName)
	//assert.NoError(t, err, "unable to delete VM")

	return
}

func TestVMExtraConfig(t *testing.T) {
	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")
	// get client
	vcdClient, err := GetTestVCDClient(
		cloudConfig,
		map[string]interface{}{
			"getVdcClient": true,
		})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	// create vApp
	vAppName := "test-vapp"
	vdcManager, err := NewVDCManager(vcdClient)
	assert.NoError(t, err, "no errors while creating ")

	vApp, err := vdcManager.GetOrCreateVApp(vAppName, cloudConfig.LB.VDCNetwork)
	assert.NoError(t, err, "unable to create vApp")
	require.NotNil(t, vApp, "vApp created should not be nil")

	// create vm
	vmNamePrefix := "test-vm-1"
	vmNum := 1
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

	vmm := NewVMManager(vcdClient, vAppName)
	// TODO: allow these vm params to be user passed through a config
	err = vmm.AddNewVM(vmNamePrefix, vAppName, vmNum, "cse",
		"ubuntu-16.04_k8-1.21_weave-2.8.1_rev1", "cse----native",
		"2core2gb", guestCustScript, true)
	assert.NoError(t, err, "unable to create [%d] VMs", vmNum)

	vms, err := vdcManager.FindAllVMsInVapp(vAppName)
	assert.NoError(t, err, "unable to find VMs in vApp")
	assert.NotNil(t, vms, "some VMs should be returned")
	assert.True(t, len(vms) == vmNum, "vapp should have [%d] vm(s), but has %d vm(s)", vmNum, len(vms))

	vm, err := vApp.GetVMByName(vms[0].Name, true)
	assert.NoError(t, err, "failed to get vm")

	// test extra config
	key := "extraconfig.test"
	value := "test123"
	err = vmm.SetVmExtraConfigKeyValue(vm, key, value, false)
	assert.NoError(t, err, "error setting VM extra config: [%v]", err)

	retrievedValue, err := vmm.GetExtraConfigValue(vm, key)
	assert.NoError(t, err, "failed to get hardware section value for key %s", key)
	assert.Equal(t, value, retrievedValue, "retrieved incorrect value")

	err = vdcManager.DeleteVApp(vAppName)
	assert.NoError(t, err, "unable to delete vApp: [%s]", vAppName)
}

