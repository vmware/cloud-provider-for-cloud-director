/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

// +build integration

package vcdclient

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCacheGatewayDetails(t *testing.T) {

	vcdClient, err := getTestVCDClient(nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	err = vcdClient.CacheGatewayDetails(ctx)
	assert.NoError(t, err, "Unable to cache gateway details")

	require.NotNil(t, vcdClient.gatewayRef, "Gateway reference should not be nil")
	assert.NotEmpty(t, vcdClient.gatewayRef.Name, "Gateway Name should not be empty")
	assert.NotEmpty(t, vcdClient.gatewayRef.Id, "Gateway Id should not be empty")

	// Missing network name should be reported
	vcdClient, err = getTestVCDClient(map[string]interface{}{
		"network": "",
	})
	assert.Error(t, err, "Should get error for unknown network")
	assert.Nil(t, vcdClient, "Client should be nil when erroring out")

	return
}

func TestDNATRuleCRUDE(t *testing.T) {

	vcdClient, err := getTestVCDClient(map[string]interface{}{
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	dnatRuleName := fmt.Sprintf("test-dnat-rule-%s", uuid.New().String())
	err = vcdClient.createDNATRule(ctx, dnatRuleName, "1.2.3.4", "1.2.3.5", 80)
	assert.NoError(t, err, "Unable to create dnat rule")

	// repeated creation should not fail
	err = vcdClient.createDNATRule(ctx, dnatRuleName, "1.2.3.4", "1.2.3.5", 80)
	assert.NoError(t, err, "Unable to create dnat rule for the second time")

	natRuleRef, err := vcdClient.getNATRuleRef(ctx, dnatRuleName)
	assert.NoError(t, err, "Unable to get dnat rule")
	require.NotNil(t, natRuleRef, "Nat Rule reference should not be nil")
	assert.Equal(t, dnatRuleName, natRuleRef.Name, "Nat Rule name should match")
	assert.NotEmpty(t, natRuleRef.ID, "Nat Rule ID should not be empty")

	err = vcdClient.deleteDNATRule(ctx, dnatRuleName, true)
	assert.NoError(t, err, "Unable to delete dnat rule")

	err = vcdClient.deleteDNATRule(ctx, dnatRuleName, true)
	assert.Error(t, err, "Should fail when deleting non-existing dnat rule")

	err = vcdClient.deleteDNATRule(ctx, dnatRuleName, false)
	assert.NoError(t, err, "Should not fail when deleting non-existing dnat rule")

	natRuleRef, err = vcdClient.getNATRuleRef(ctx, dnatRuleName)
	assert.NoError(t, err, "Get should not fail when nat rule is absent")
	assert.Nil(t, natRuleRef, "Deleted nat rule reference should be nil")

	return
}

func TestLBPoolCRUDE(t *testing.T) {

	vcdClient, err := getTestVCDClient(nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	lbPoolName := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	lbPoolRef, err := vcdClient.createLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234)
	assert.NoError(t, err, "Unable to create lb pool")
	require.NotNil(t, lbPoolRef, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolName, lbPoolRef.Name, "LB Pool name should match")

	// repeated creation should not fail
	lbPoolRef, err = vcdClient.createLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234)
	assert.NoError(t, err, "Unable to create lb pool for the second time")
	require.NotNil(t, lbPoolRef, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolName, lbPoolRef.Name, "LB Pool name should match")

	lbPoolRefObtained, err := vcdClient.getLoadBalancerPool(ctx, lbPoolName)
	assert.NoError(t, err, "Unable to get lbPool ref")
	require.NotNil(t, lbPoolRefObtained, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolRefObtained.Name, lbPoolRef.Name, "LB Pool name should match")
	assert.NotEmpty(t, lbPoolRefObtained.Id, "LB Pool ID should not be empty")

	updatedIps := []string{"5.5.5.5"}
	lbPoolRefUpdated, err := vcdClient.updateLoadBalancerPool(ctx, lbPoolName, updatedIps, 55555)
	assert.NoError(t, err, "No lbPool ref for updated lbPool")
	require.NotNil(t, lbPoolRefUpdated, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolRefUpdated.Name, lbPoolRef.Name, "LB Pool name should match")
	assert.NotEmpty(t, lbPoolRefUpdated.Id, "LB Pool ID should not be empty")

	lbPoolSummaryUpdated, err := vcdClient.getLoadBalancerPoolSummary(ctx, lbPoolName)
	assert.NoError(t, err, "No LB Pool summary for updated pool.")
	require.NotNil(t, lbPoolSummaryUpdated, "LB Pool summary reference should not be nil")
	assert.Equal(t, lbPoolSummaryUpdated.MemberCount, int32(len(updatedIps)), "LB Pool should have updated size %d", len(updatedIps))

	err = vcdClient.deleteLoadBalancerPool(ctx, lbPoolName, true)
	assert.NoError(t, err, "Unable to delete LB Pool")

	err = vcdClient.deleteLoadBalancerPool(ctx, lbPoolName, true)
	assert.Error(t, err, "Should fail when deleting non-existing lb pool")

	err = vcdClient.deleteLoadBalancerPool(ctx, lbPoolName, false)
	assert.NoError(t, err, "Should not fail when deleting non-existing lb pool")

	lbPoolRef, err = vcdClient.getLoadBalancerPool(ctx, lbPoolName)
	assert.NoError(t, err, "Get should not fail when lb pool is absent")
	assert.Nil(t, lbPoolRef, "Deleted lb pool reference should be nil")

	lbPoolRef, err = vcdClient.updateLoadBalancerPool(ctx, lbPoolName, updatedIps, 55555)
	assert.Error(t, err, "Updating deleted lb pool should fail")
	assert.Nil(t, lbPoolRef, "Deleted lb pool reference should be nil")

	return
}

func TestGetLoadBalancerSEG(t *testing.T) {

	vcdClient, err := getTestVCDClient(nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	segRef, err := vcdClient.getLoadBalancerSEG(ctx)
	assert.NoError(t, err, "Unable to get ServiceEngineGroup ref")
	require.NotNil(t, segRef, "ServiceEngineGroup reference should not be nil")
	assert.NotEmpty(t, segRef.Name, "ServiceEngineGroup Name should not be empty")
	assert.NotEmpty(t, segRef.Id, "ServiceEngineGroup ID should not be empty")

	return
}

func TestVirtualServiceHttpCRUDE(t *testing.T) {

	vcdClient, err := getTestVCDClient(nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	lbPoolName := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	lbPoolRef, err := vcdClient.createLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234)
	assert.NoError(t, err, "Unable to create lb pool")

	segRef, err := vcdClient.getLoadBalancerSEG(ctx)
	assert.NoError(t, err, "Unable to get ServiceEngineGroup ref")

	virtualServiceName := fmt.Sprintf("test-virtual-service-%s", uuid.New().String())
	vsRef, err := vcdClient.createVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
		"2.3.4.5", "HTTP", 80, "")
	assert.NoError(t, err, "Unable to create virtual service")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	// repeated creation should not fail
	vsRef, err = vcdClient.createVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
		"2.3.4.5", "HTTP", 80, "")
	assert.NoError(t, err, "Unable to create virtual service for the second time")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	vsRefObtained, err := vcdClient.getVirtualService(ctx, virtualServiceName)
	assert.NoError(t, err, "Unable to get virtual service ref")
	require.NotNil(t, vsRefObtained, "Virtual service reference should not be nil")
	assert.Equal(t, vsRefObtained.Name, vsRef.Name, "Virtual service reference name should match")
	assert.NotEmpty(t, vsRefObtained.Id, "Virtual service ID should not be empty")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, true)
	assert.NoError(t, err, "Unable to delete Virtual Service")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, true)
	assert.Error(t, err, "Should fail when deleting non-existing Virtual Service")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, false)
	assert.NoError(t, err, "Should not fail when deleting non-existing Virtual Service")

	vsRefObtained, err = vcdClient.getVirtualService(ctx, virtualServiceName)
	assert.NoError(t, err, "Get should not fail when Virtual Service is absent")
	assert.Nil(t, vsRefObtained, "Deleted Virtual Service reference should be nil")

	err = vcdClient.deleteLoadBalancerPool(ctx, lbPoolName, true)
	assert.NoError(t, err, "Should not fail when deleting lb pool")

	return
}

func foundStringInSlice(find string, slice []string) bool {
	for _, currElement := range slice {
		if currElement == find {
			return true
		}
	}
	return  false
}

func TestVirtualServiceHttpsCRUDE(t *testing.T) {

	vcdClient, err := getTestVCDClient(nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	lbPoolName := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	lbPoolRef, err := vcdClient.createLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234)
	assert.NoError(t, err, "Unable to create lb pool")

	segRef, err := vcdClient.getLoadBalancerSEG(ctx)
	assert.NoError(t, err, "Unable to get ServiceEngineGroup ref")

	virtualServiceName := fmt.Sprintf("test-virtual-service-https-%s", uuid.New().String())
	vsRef, err := vcdClient.createVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
		"2.3.4.5", "HTTPS", 443, "test")
	assert.NoError(t, err, "Unable to create virtual service")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	vsRefObtained, err := vcdClient.getVirtualService(ctx, virtualServiceName)
	assert.NoError(t, err, "Unable to get virtual service ref")
	require.NotNil(t, vsRefObtained, "Virtual service reference should not be nil")
	assert.Equal(t, vsRefObtained.Name, vsRef.Name, "Virtual service reference name should match")
	assert.NotEmpty(t, vsRefObtained.Id, "Virtual service ID should not be empty")

	rdeVips, err := vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Unable to get RDE vips after virtual service creation")
	assert.Equal(t, true, foundStringInSlice(vsRefObtained.VirtualIpAddress, rdeVips), "virtual service ip should be found in rde vips")

	// repeated creation should not fail
	vsRef, err = vcdClient.createVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
		"2.3.4.5", "HTTPS", 443, "test")
	assert.NoError(t, err, "Unable to create virtual service for the second time")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, true)
	assert.NoError(t, err, "Unable to delete Virtual Service")

	rdeVips, err = vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Unable to get vips from RDE after virtual service deletion")
	assert.Equal(t, false, foundStringInSlice(vsRefObtained.VirtualIpAddress, rdeVips), "virtual service ip should not be found in RDE vips")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, true)
	assert.Error(t, err, "Should fail when deleting non-existing Virtual Service")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, false)
	assert.NoError(t, err, "Should not fail when deleting non-existing Virtual Service")

	vsRefObtained, err = vcdClient.getVirtualService(ctx, virtualServiceName)
	assert.NoError(t, err, "Get should not fail when Virtual Service is absent")
	assert.Nil(t, vsRefObtained, "Deleted Virtual Service reference should be nil")

	err = vcdClient.deleteLoadBalancerPool(ctx, lbPoolName, true)
	assert.NoError(t, err, "Should not fail when deleting lb pool")

	return
}

func TestLoadBalancerCRUDE(t *testing.T) {

	vcdClient, err := getTestVCDClient(nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	virtualServiceNamePrefix := fmt.Sprintf("test-virtual-service-https-%s", uuid.New().String())
	lbPoolNamePrefix := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	freeIP, err := vcdClient.CreateLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix, []string{"1.2.3.4", "1.2.3.5"}, 31234, 31235)
	assert.NoError(t, err, "Load Balancer should be created")
	assert.NotEmpty(t, freeIP, "There should be a non-empty IP returned")

	virtualServiceNameHttp := fmt.Sprintf("%s-http", virtualServiceNamePrefix)
	freeIPObtained, err := vcdClient.GetLoadBalancer(ctx, virtualServiceNameHttp)
	assert.NoError(t, err, "Load Balancer should be found")
	assert.Equal(t, freeIP, freeIPObtained, "The IPs should match")

	virtualServiceNameHttps := fmt.Sprintf("%s-https", virtualServiceNamePrefix)
	freeIPObtained, err = vcdClient.GetLoadBalancer(ctx, virtualServiceNameHttps)
	assert.NoError(t, err, "Load Balancer should be found")
	assert.Equal(t, freeIP, freeIPObtained, "The IPs should match")

	freeIP, err = vcdClient.CreateLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix, []string{"1.2.3.4", "1.2.3.5"}, 31234, 31235)
	assert.NoError(t, err, "Load Balancer should be created even on second attempt")
	assert.NotEmpty(t, freeIP, "There should be a non-empty IP returned")

	updatedIps := []string{"5.5.5.5"}
	updatedInternalPort := int32(55555)
	err = vcdClient.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", updatedIps, updatedInternalPort)
	assert.NoError(t, err, "HTTP Load Balancer should be updated")
	err = vcdClient.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", updatedIps, updatedInternalPort)
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")


	err = vcdClient.DeleteLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix)
	assert.NoError(t, err, "Load Balancer should be deleted")

	freeIPObtained, err = vcdClient.GetLoadBalancer(ctx, virtualServiceNameHttp)
	assert.NoError(t, err, "Load Balancer should not be found")
	assert.Empty(t, freeIPObtained, "The VIP should not be found")

	freeIPObtained, err = vcdClient.GetLoadBalancer(ctx, virtualServiceNameHttps)
	assert.NoError(t, err, "Load Balancer should not be found")
	assert.Empty(t, freeIPObtained, "The VIP should not be found")

	err = vcdClient.DeleteLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix)
	assert.NoError(t, err, "Repeated deletion of Load Balancer should not fail")

	err = vcdClient.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", updatedIps, updatedInternalPort)
	assert.Error(t, err, "updating deleted HTTP Load Balancer should be an error")
	err = vcdClient.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", updatedIps, updatedInternalPort)
	assert.Error(t, err, "updating deleted HTTPS Load Balancer should be an error")

	return
}
