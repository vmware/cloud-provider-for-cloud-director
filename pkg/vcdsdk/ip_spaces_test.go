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
	"github.com/vmware/go-vcloud-director/v2/types/v56"
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
	ipSpaceIds, err := gm.FetchIpSpacesBackingGateway(ctx)
	assert.NoErrorf(t, err, "Error calling method FetchIpSpacesBackingGateway for netowrk %s", testData.OrgvdcNetworkIpSpace)
	assert.Greater(t, len(ipSpaceIds), 0)

	fmt.Println("Calling FilterIpSpacesByType for gateway using Ip Spaces")
	ipSpaces, err := gm.FilterIpSpacesByType(ipSpaceIds, "PUBLIC")
	assert.NoErrorf(t, err, "Error calling method FilterIpSpacesByType for Ids %s", ipSpaceIds)
	assert.Equal(t, 1, len(ipSpaces))

	gm2, err := NewGatewayManager(ctx, vcdClient, testData.OrgvdcNetworkNoIpSpace, "", testData.Orgvdc)
	assert.NoErrorf(t, err, "Unable to create Gateway Manager for network %s", testData.OrgvdcNetworkNoIpSpace)

	fmt.Println("Calling FetchIpSpacesBackingGateway for gateway using Ip blocks")
	_, err = gm2.FetchIpSpacesBackingGateway(ctx)
	assert.Error(t, err, "Expected error calling method IsUsingIpSpaces for network %s", testData.OrgvdcNetworkNoIpSpace)
	assert.True(t, strings.Contains(err.Error(), "403 Forbidden"), "Expected '403 forbidden' message in error found %s", err.Error())
}

func TestIpSpaceOperations(t *testing.T) {
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
	ipSpaceIds, err := gm.FetchIpSpacesBackingGateway(ctx)
	assert.NoErrorf(t, err, "Error calling method FetchIpSpacesBackingGateway for netowrk %s", testData.OrgvdcNetworkIpSpace)
	assert.Greater(t, len(ipSpaceIds), 0)

	fmt.Println("Calling FilterIpSpacesByType for gateway using Ip Spaces")
	ipSpaces, err := gm.FilterIpSpacesByType(ipSpaceIds, "PUBLIC")
	assert.NoErrorf(t, err, "Error calling method FilterIpSpacesByType for Ids %s", ipSpaceIds)
	assert.Equal(t, 1, len(ipSpaces))

	ipSpace := ipSpaces[0]

	fmt.Println("Calling AllocateIpFromIpSpace for gateway using Ip Spaces")
	allocationId, allocatedIp, err := gm.AllocateIpFromIpSpace(ipSpace)
	assert.NoError(t, err, "Error calling method AllocateIpFromIpSpace for IpSpace [%s]", ipSpace.IpSpace.Name)
	assert.NotNil(t, allocationId, "Expected non nil Allocation Id")
	assert.NotNil(t, allocatedIp, "Expected non nil Allocated Ip")

	fmt.Println("Calling FindIpAllocationByIp for gateway using Ip Spaces")
	retrievedIpSpaceAllocation, err := gm.FindIpAllocationByIp(ipSpace, allocatedIp)
	assert.NoError(t, err, "Error calling method FindIpAllocationByIp for IpSpace [%s], Ip [%s]", ipSpace.IpSpace.Name, allocatedIp)
	assert.NotNil(t, retrievedIpSpaceAllocation)
	assert.Equal(t, allocationId, retrievedIpSpaceAllocation.IpSpaceIpAllocation.ID, "")
	assert.Equal(t, allocatedIp, retrievedIpSpaceAllocation.IpSpaceIpAllocation.Value, "")

	dummyMarker := "Dummy Marker"
	fmt.Println("Calling MarkIpAsUsed for gateway using Ip Spaces")
	updatedIpSpaceAllocation, err := gm.MarkIpAsUsed(retrievedIpSpaceAllocation, dummyMarker)
	assert.NoError(t, err, "Error calling method MarkIpAsUsed for Allocation [%s]", retrievedIpSpaceAllocation.IpSpaceIpAllocation.ID)
	assert.NotNil(t, updatedIpSpaceAllocation)
	assert.Equal(t, types.IpSpaceIpAllocationUsedManual, updatedIpSpaceAllocation.IpSpaceIpAllocation.UsageState)
	assert.Equal(t, dummyMarker, updatedIpSpaceAllocation.IpSpaceIpAllocation.Description)

	fmt.Println("Calling FindIpAllocationByMarker for gateway using Ip Spaces")
	retrievedIpSpaceAllocation2, err := gm.FindIpAllocationByMarker(ipSpace, dummyMarker)
	assert.NoError(t, err, "Error calling method MarkIpAsUsed for Allocation [%s]", retrievedIpSpaceAllocation.IpSpaceIpAllocation.ID)
	assert.NotNil(t, retrievedIpSpaceAllocation2)
	assert.Equal(t, retrievedIpSpaceAllocation.IpSpaceIpAllocation.ID, retrievedIpSpaceAllocation2.IpSpaceIpAllocation.ID, "")
	assert.Equal(t, retrievedIpSpaceAllocation.IpSpaceIpAllocation.Value, retrievedIpSpaceAllocation2.IpSpaceIpAllocation.Value, "")

	fmt.Println("Calling MarkIpAsUnused for gateway using Ip Spaces")
	updatedIpSpaceAllocation, err = gm.MarkIpAsUnused(retrievedIpSpaceAllocation2)
	assert.NoError(t, err, "Error calling method MarkIpAsUnused for Allocation [%s]", retrievedIpSpaceAllocation2.IpSpaceIpAllocation.ID)
	assert.NotNil(t, updatedIpSpaceAllocation)
	assert.Equal(t, types.IpSpaceIpAllocationUnused, updatedIpSpaceAllocation.IpSpaceIpAllocation.UsageState)
	assert.Equal(t, "", updatedIpSpaceAllocation.IpSpaceIpAllocation.Description)

	fmt.Println("Calling ReleaseIp for gateway using Ip Spaces")
	err = gm.ReleaseIp(updatedIpSpaceAllocation)
	assert.NoError(t, err, "Error calling method ReleaseIp for Allocation [%s]", updatedIpSpaceAllocation.IpSpaceIpAllocation.ID)
}
