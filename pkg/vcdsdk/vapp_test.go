/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/
package vcdsdk

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVApp(t *testing.T) {

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
	vAppName := "manual-vapp"
	vdcManager, err := NewVDCManager(vcdClient, vcdClient.ClusterOrgName, vcdClient.ClusterOVDCName)
	assert.NoError(t, err, "error creating VDCManager")

	vms, err := vdcManager.FindAllVMsInVapp(vAppName)
	assert.NoError(t, err, "unable to find VMs in vApp")
	assert.NotNil(t, vms, "some VMs should be returned")

	return
}

func TestDeleteVapp(t *testing.T) {
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

	vdcManager := VdcManager{
		Client:  vcdClient,
		Vdc:     vcdClient.VDC,
	}
	vapp, err := vdcManager.GetOrCreateVApp("vapp1", "ovdc1_nw")
	assert.NoError(t, err, "unable to find vApp")
	assert.NotNil(t, vapp, "vapp should not be nil")
	err = vdcManager.DeleteVApp("vapp1")
	assert.NoError(t, err, "unable to delete vApp")
}

func TestVdcManager_CacheVdcDetails(t *testing.T) {
	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := GetTestVCDClient(
		cloudConfig,
		map[string]interface{}{
			"getVdcClient": true,
		})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")
	vdcManager := VdcManager{
		Client:  vcdClient,
	}
	err = vdcManager.CacheVdcDetails()
	assert.NoError(t, err, "There should no error while caching VDC details")
}

