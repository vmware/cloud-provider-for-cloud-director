/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package vcdsdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"testing"
)

func TestGatewayUsingIpSpaces(t *testing.T) {
	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing vcd info file contents.")

	testDataFile := filepath.Join(gitRoot, "testdata/ip_space_test.yaml")
	testDataFileContent, err := os.ReadFile(testDataFile)
	assert.NoError(t, err, "There should be no error reading the ip space test data file contents.")

	var testData ipSpaceDetails
	err = yaml.Unmarshal(testDataFileContent, &testData)
	assert.NoError(t, err, "There should be no error parsing ip space test data file content.")

	vcdClient, err := getTestVCDClient(vcdConfig, map[string]interface{}{
		"user":         testData.Username,
		"refreshToken": testData.RefreshToken,
		"userOrg":      testData.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, testData.OrgvdcNetworkIpSpace, "", testData.Orgvdc)
	assert.NoErrorf(t, err, "Unable to create Gateway Manager for network %s", testData.OrgvdcNetworkIpSpace)

	result, err := gm.IsUsingIpSpaces()
	assert.NoErrorf(t, err, "Error calling method IsUsingIpSpaces for netowrk %s", testData.OrgvdcNetworkIpSpace)
	assert.True(t, result)

	gm2, err := NewGatewayManager(ctx, vcdClient, testData.OrgvdcNetworkNoIpSpace, "", testData.Orgvdc)
	assert.NoErrorf(t, err, "Unable to create Gateway Manager for network %s", testData.OrgvdcNetworkNoIpSpace)

	result, err = gm2.IsUsingIpSpaces()
	assert.NoErrorf(t, err, "Error calling method IsUsingIpSpaces for network %s", testData.OrgvdcNetworkNoIpSpace)
	assert.False(t, result)
}
