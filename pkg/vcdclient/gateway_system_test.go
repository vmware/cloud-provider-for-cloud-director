/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package vcdclient

import (
	"context"
	"fmt"
	"net/http"
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

func TestGetUnusedGatewayIP(t *testing.T) {

	vcdClient, err := getTestVCDClient(map[string]interface{}{
		"subnet": "",
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	validSubnet := "10.150.191.253/19"
	externalIP, err := vcdClient.getUnusedExternalIPAddress(ctx, validSubnet)
	assert.NoError(t, err, "should not get an error for this range")
	assert.NotEmpty(t, externalIP, "should get a valid IP address in the range [%s]", validSubnet)

	invalidSubnet := "1.1.1.1/24"
	externalIP, err = vcdClient.getUnusedExternalIPAddress(ctx, invalidSubnet)
	assert.Error(t, err, "should get an error for this range")
	assert.Empty(t, externalIP, "should not get a valid IP address in the range [%s]", invalidSubnet)

	everythingAllowedSubnet := ""
	externalIP, err = vcdClient.getUnusedExternalIPAddress(ctx, everythingAllowedSubnet)
	assert.NoError(t, err, "should not get an error for this range")
	assert.NotEmpty(t, externalIP, "should get a valid IP address in the empty range")

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
	externalIP := "10.11.12.13"
	internalIP := "2.3.4.5"
	vsRef, err := vcdClient.createVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
		internalIP, externalIP, "HTTP", 80, "")
	assert.NoError(t, err, "Unable to create virtual service")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	vsRefObtained, err := vcdClient.getVirtualService(ctx, virtualServiceName)
	assert.NoError(t, err, "Unable to get virtual service ref")
	require.NotNil(t, vsRefObtained, "Virtual service reference should not be nil")
	assert.Equal(t, vsRefObtained.Name, vsRef.Name, "Virtual service reference name should match")
	assert.NotEmpty(t, vsRefObtained.Id, "Virtual service ID should not be empty")

	rdeVips, _, _, err := vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Unable to get RDE vips after virtual service creation")
	assert.True(t, foundStringInSlice(externalIP, rdeVips), "external ip should be found in rde vips")
	assert.False(t, foundStringInSlice(internalIP, rdeVips), "internal ip should not be in rde vips")

	// repeated creation should not fail
	vsRef, err = vcdClient.createVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
		internalIP, externalIP, "HTTP", 80, "")
	assert.NoError(t, err, "Unable to create virtual service for the second time")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, true, externalIP)
	assert.NoError(t, err, "Unable to delete Virtual Service")

	rdeVips, _, _, err = vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Unable to get vips from RDE after virtual service deletion")
	assert.False(t, foundStringInSlice(externalIP, rdeVips), "external ip should not be found in RDE vips")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, true, externalIP)
	assert.Error(t, err, "Should fail when deleting non-existing Virtual Service")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, false, externalIP)
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

	externalIP := "11.12.13.14"
	internalIP := "3.4.5.6"
	virtualServiceName := fmt.Sprintf("test-virtual-service-https-%s", uuid.New().String())
	certName := fmt.Sprintf("%s-cert", vcdClient.ClusterID)
	vsRef, err := vcdClient.createVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
		internalIP, externalIP, "HTTPS", 443, certName)
	assert.NoError(t, err, "Unable to create virtual service")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	vsRefObtained, err := vcdClient.getVirtualService(ctx, virtualServiceName)
	assert.NoError(t, err, "Unable to get virtual service ref")
	require.NotNil(t, vsRefObtained, "Virtual service reference should not be nil")
	assert.Equal(t, vsRefObtained.Name, vsRef.Name, "Virtual service reference name should match")
	assert.NotEmpty(t, vsRefObtained.Id, "Virtual service ID should not be empty")

	rdeVips, _, _, err := vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Unable to get RDE vips after virtual service creation")
	assert.True(t, foundStringInSlice(externalIP, rdeVips), "external ip should be found in rde vips")
	assert.False(t, foundStringInSlice(internalIP, rdeVips), "internal ip should not be found in rde vips")

	// repeated creation should not fail
	vsRef, err = vcdClient.createVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
		internalIP, externalIP, "HTTPS", 443, certName)
	assert.NoError(t, err, "Unable to create virtual service for the second time")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, true, externalIP)
	assert.NoError(t, err, "Unable to delete Virtual Service")

	rdeVips, _, _, err = vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Unable to get vips from RDE after virtual service deletion")
	assert.False(t, foundStringInSlice(externalIP, rdeVips), "external ip should not be found in RDE vips")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, true, externalIP)
	assert.Error(t, err, "Should fail when deleting non-existing Virtual Service")

	err = vcdClient.deleteVirtualService(ctx, virtualServiceName, false, externalIP)
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

func TestUpdateRDEUsingEtag(t *testing.T) {
	vcdClient, err := getTestVCDClient(nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	// get rde Vips
	rdeVips1, etag1, defEnt1, err := vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips on first attempt")
	rdeVips2, etag2, defEnt2, err := vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips on second attempt")
	assert.Equal(t, etag1, etag2, "etags from consecutive RDE GET calls should match")
	origRdeVips := make([]string, len(rdeVips1))
	copy(origRdeVips, rdeVips1)

	// Test successfully updating using first etag
	addIp1 := "1.2.3.4"
	addIp2 := "2.3.4.5"
	updatedRdeVips1 := append(rdeVips1, addIp1)
	httpResponse1, err := vcdClient.updateRDEVirtualIps(ctx, updatedRdeVips1, etag1, defEnt1)
	assert.NoError(t, err, "RDE should be updated")
	assert.Equal(t, http.StatusOK, httpResponse1.StatusCode, "RDE update status code should be 200 (OK)")
	rdeVips3, _, _, err := vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips successfully")
	assert.True(t, foundStringInSlice(addIp1, rdeVips3), "ip [%s] should be found in rde vips", addIp1)

	// Test adding addIp2 with outdated etag
	updatedRdeVips2 := append(rdeVips2, addIp2)
	httpResponse2, err := vcdClient.updateRDEVirtualIps(ctx, updatedRdeVips2, etag2, defEnt2)
	assert.Error(t, err, "Should have an error updating RDE with outdated etag")
	assert.Equal(t, http.StatusPreconditionFailed, httpResponse2.StatusCode, "RDE update status code should be 412 (Precondition failed)")
	rdeVips3, etag3, defEnt3, err := vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips successfully")
	assert.False(t, foundStringInSlice(addIp2, rdeVips3), "ip [%s] should not be found in rde vips", addIp2)

	// Try adding addIp2 with correct etag
	updatedRdeVips3 := append(rdeVips3, addIp2)
	httpResponse3, err := vcdClient.updateRDEVirtualIps(ctx, updatedRdeVips3, etag3, defEnt3)
	assert.NoError(t, err, "RDE should be updated")
	assert.Equal(t, http.StatusOK, httpResponse3.StatusCode, "RDE update status code should be 200 (OK)")
	rdeVips4, etag4, defEnt4, err := vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips successfully")
	assert.True(t, foundStringInSlice(addIp2, rdeVips4), "ip [%s] should be found in rde vips", addIp2)


	// reset RDE vips to original state
	httpResponse4, err := vcdClient.updateRDEVirtualIps(ctx, rdeVips1, etag4, defEnt4)
	assert.NoError(t, err, "RDE should be updated")
	assert.Equal(t, http.StatusOK, httpResponse4.StatusCode, "RDE update status code should be 200 (OK)")
	// no check to ensure ip's removed because they may have been previously present in the RDE vips
	rdeVips5, _, _, err := vcdClient.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips to check added ips are removed")
	assert.False(t, foundStringInSlice(addIp1, rdeVips5), "ip [%s] should not be found in rde vips", addIp1)
	assert.False(t, foundStringInSlice(addIp2, rdeVips5), "ip [%s] should not be found in rde vips", addIp2)
}
