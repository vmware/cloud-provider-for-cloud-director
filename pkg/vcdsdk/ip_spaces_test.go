/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package vcdsdk

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
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
		"userOrg":      testData.UserOrg,
		"refreshToken": testData.RefreshToken,
		"org":          testData.UserOrg,
		"vdc":          testData.Orgvdc,
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

func TestFetchIpSpacesBackingGateway(t *testing.T) {
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
		"userOrg":      testData.UserOrg,
		"refreshToken": testData.RefreshToken,
		"org":          testData.UserOrg,
		"vdc":          testData.Orgvdc,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, testData.OrgvdcNetworkIpSpace, "", testData.Orgvdc)
	assert.NoErrorf(t, err, "Unable to create Gateway Manager for network %s", testData.OrgvdcNetworkIpSpace)

	fmt.Println("Calling FetchIpSpacesBackingGateway for gateway using Ip Spaces")
	result, err := gm.FetchIpSpacesBackingGateway(ctx)
	assert.NoErrorf(t, err, "Error calling method FetchIpSpacesBackingGateway for netowrk %s", testData.OrgvdcNetworkIpSpace)
	assert.Greater(t, len(result), 0)

	gm2, err := NewGatewayManager(ctx, vcdClient, testData.OrgvdcNetworkNoIpSpace, "", testData.Orgvdc)
	assert.NoErrorf(t, err, "Unable to create Gateway Manager for network %s", testData.OrgvdcNetworkNoIpSpace)

	fmt.Println("Calling FetchIpSpacesBackingGateway for gateway using Ip blocks")
	result, err = gm2.FetchIpSpacesBackingGateway(ctx)
	assert.Error(t, err, "Expected error calling method IsUsingIpSpaces for network %s", testData.OrgvdcNetworkNoIpSpace)
	assert.True(t, strings.Contains(err.Error(), "403 Forbidden"), "Expected '403 forbidden' message in error found %s", err.Error())
}
