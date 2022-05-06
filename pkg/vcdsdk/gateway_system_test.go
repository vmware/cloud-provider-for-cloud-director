/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package vcdsdk

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	swagger "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"
)

const BusyRetries = 5

func TestCacheGatewayDetails(t *testing.T) {

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := GetTestVCDClient(cloudConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, cloudConfig.LB.VDCNetwork, cloudConfig.LB.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	require.NotNil(t, gm.GatewayRef, "Gateway reference should not be nil")
	assert.NotEmpty(t, gm.GatewayRef.Name, "Gateway Name should not be empty")
	assert.NotEmpty(t, gm.GatewayRef.Id, "Gateway Id should not be empty")

	// Missing network name should be reported
	vcdClient, err = GetTestVCDClient(cloudConfig, map[string]interface{}{
		"network": "",
	})
	assert.Error(t, err, "Should get error for unknown network")
	assert.Nil(t, vcdClient, "Client should be nil when erroring out")

	return
}

func TestDNATRuleCRUDE(t *testing.T) {

	// NOTE: For the time being, NoRdePrefix prefix will be used in clusterID
	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := GetTestVCDClient(cloudConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, cloudConfig.LB.VDCNetwork, cloudConfig.LB.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	dnatRuleName := fmt.Sprintf("test-dnat-rule-%s", uuid.New().String())

	appPortProfileName := GetAppPortProfileName(dnatRuleName)
	appPortProfile, err := gm.CreateAppPortProfile(appPortProfileName, 80)
	assert.NoError(t, err, "there should be no error creating an app port profile")
	assert.NotNil(t, appPortProfile, "app port profile created should not be nil")
	assert.NotNil(t, appPortProfile.NsxtAppPortProfile, "nsxt app port profile should not be nil")

	err = gm.CreateDNATRule(ctx, dnatRuleName, "1.2.3.4", "1.2.3.5", 80, 36123, appPortProfile)
	assert.NoError(t, err, "Unable to create dnat rule")

	// repeated creation should not fail
	err = gm.CreateDNATRule(ctx, dnatRuleName, "1.2.3.4", "1.2.3.5", 80, 36123, appPortProfile)
	assert.NoError(t, err, "Unable to create dnat rule for the second time")

	natRuleRef, err := gm.GetNATRuleRef(ctx, dnatRuleName)
	assert.NoError(t, err, "Unable to get dnat rule")
	require.NotNil(t, natRuleRef, "Nat Rule reference should not be nil")
	assert.Equal(t, dnatRuleName, natRuleRef.Name, "Nat Rule name should match")
	assert.NotEmpty(t, natRuleRef.ID, "Nat Rule ID should not be empty")

	_, err = gm.UpdateDNATRule(ctx, dnatRuleName, "2.3.4.5", "2.3.4.5", 8080)
	assert.NoError(t, err, "Unable to update dnat rule")

	_, err = gm.UpdateDNATRule(ctx, dnatRuleName, "2.3.4.5", "2.3.4.5", 8080)
	assert.NoError(t, err, "repeated updates to dnat rule should not fail")

	err = gm.DeleteDNATRule(ctx, dnatRuleName, true)
	assert.NoError(t, err, "Unable to delete dnat rule")

	err = gm.DeleteDNATRule(ctx, dnatRuleName, true)
	assert.Error(t, err, "Should fail when deleting non-existing dnat rule")

	err = gm.DeleteDNATRule(ctx, dnatRuleName, false)
	assert.NoError(t, err, "Should not fail when deleting non-existing dnat rule")

	err = gm.DeleteAppPortProfile(appPortProfileName, false)
	assert.NoError(t, err, "there should be no error when app port profile is deleted")

	err = gm.DeleteAppPortProfile(appPortProfileName, false)
	assert.NoError(t, err, "there should be no error when app port profile is deleted again")

	natRuleRef, err = gm.GetNATRuleRef(ctx, dnatRuleName)
	assert.NoError(t, err, "Get should not fail when nat rule is absent")
	assert.Nil(t, natRuleRef, "Deleted nat rule reference should be nil")

	return
}

func TestLBPoolCRUDE(t *testing.T) {

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := GetTestVCDClient(cloudConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	// Avoid RDE updates by using clusterID which has `NoRdePrefix` prefix
	cloudConfig.ClusterID = fmt.Sprintf("%s-%s", uuid.New().String(), "NoRdePrefix")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, cloudConfig.LB.VDCNetwork, cloudConfig.LB.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	lbPoolName := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())

	lbPoolRef, err := gm.CreateLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234)
	assert.NoError(t, err, "Unable to create lb pool")
	require.NotNil(t, lbPoolRef, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolName, lbPoolRef.Name, "LB Pool name should match")

	// repeated creation should not fail
	lbPoolRef, err = gm.CreateLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234)
	assert.NoError(t, err, "Unable to create lb pool for the second time")
	require.NotNil(t, lbPoolRef, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolName, lbPoolRef.Name, "LB Pool name should match")

	lbPoolRefObtained, err := gm.getLoadBalancerPool(ctx, lbPoolName)
	assert.NoError(t, err, "Unable to get lbPool ref")
	require.NotNil(t, lbPoolRefObtained, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolRefObtained.Name, lbPoolRef.Name, "LB Pool name should match")
	assert.NotEmpty(t, lbPoolRefObtained.Id, "LB Pool ID should not be empty")

	updatedIps := []string{"5.5.5.5"}
	lbPoolRefUpdated, err := gm.UpdateLoadBalancerPool(ctx, lbPoolName, updatedIps, 55555)
	assert.NoError(t, err, "No lbPool ref for updated lbPool")
	require.NotNil(t, lbPoolRefUpdated, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolRefUpdated.Name, lbPoolRef.Name, "LB Pool name should match")
	assert.NotEmpty(t, lbPoolRefUpdated.Id, "LB Pool ID should not be empty")

	// repeated update should work
	lbPoolRefUpdated, err = gm.UpdateLoadBalancerPool(ctx, lbPoolName, updatedIps, 55555)
	assert.NoError(t, err, "There should be no error on repeated LB pool update")
	require.NotNil(t, lbPoolRefUpdated, "LB Pool reference should not be nil on repeated update")
	assert.Equal(t, lbPoolRefUpdated.Name, lbPoolRef.Name, "LB Pool name should match on repeated update")
	assert.NotEmpty(t, lbPoolRefUpdated.Id, "LB Pool ID should not be empty on repeated update")

	lbPoolSummaryUpdated, err := gm.getLoadBalancerPoolSummary(ctx, lbPoolName)
	assert.NoError(t, err, "No LB Pool summary for updated pool.")
	require.NotNil(t, lbPoolSummaryUpdated, "LB Pool summary reference should not be nil")
	assert.Equal(t, lbPoolSummaryUpdated.MemberCount, int32(len(updatedIps)), "LB Pool should have updated size %d", len(updatedIps))

	err = gm.DeleteLoadBalancerPool(ctx, lbPoolName, true)
	assert.NoError(t, err, "Unable to delete LB Pool")

	err = gm.DeleteLoadBalancerPool(ctx, lbPoolName, true)
	assert.Error(t, err, "Should fail when deleting non-existing lb pool")

	err = gm.DeleteLoadBalancerPool(ctx, lbPoolName, false)
	assert.NoError(t, err, "Should not fail when deleting non-existing lb pool")

	lbPoolRef, err = gm.getLoadBalancerPool(ctx, lbPoolName)
	assert.NoError(t, err, "Get should not fail when lb pool is absent")
	assert.Nil(t, lbPoolRef, "Deleted lb pool reference should be nil")

	lbPoolRef, err = gm.UpdateLoadBalancerPool(ctx, lbPoolName, updatedIps, 55555)
	assert.Error(t, err, "Updating deleted lb pool should fail")
	assert.Nil(t, lbPoolRef, "Deleted lb pool reference should be nil")

	return
}

func TestGetLoadBalancerSEG(t *testing.T) {

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := GetTestVCDClient(cloudConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, cloudConfig.LB.VDCNetwork, cloudConfig.LB.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	segRef, err := gm.GetLoadBalancerSEG(ctx)
	assert.NoError(t, err, "Unable to get ServiceEngineGroup ref")
	require.NotNil(t, segRef, "ServiceEngineGroup reference should not be nil")
	assert.NotEmpty(t, segRef.Name, "ServiceEngineGroup Name should not be empty")
	assert.NotEmpty(t, segRef.Id, "ServiceEngineGroup ID should not be empty")

	return
}

func TestGetUnusedGatewayIP(t *testing.T) {

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := GetTestVCDClient(cloudConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
		"subnet":       "",
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, cloudConfig.LB.VDCNetwork, cloudConfig.LB.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	validSubnet := cloudConfig.LB.VIPSubnet
	externalIP, err := gm.GetUnusedExternalIPAddress(ctx, validSubnet)
	assert.NoError(t, err, "should not get an error for this range")
	assert.NotEmpty(t, externalIP, "should get a valid IP address in the range [%s]", validSubnet)

	invalidSubnet := "1.1.1.1/24"
	externalIP, err = gm.GetUnusedExternalIPAddress(ctx, invalidSubnet)
	assert.Error(t, err, "should get an error for this range")
	assert.Empty(t, externalIP, "should not get a valid IP address in the range [%s]", invalidSubnet)

	everythingAllowedSubnet := ""
	externalIP, err = gm.GetUnusedExternalIPAddress(ctx, everythingAllowedSubnet)
	assert.NoError(t, err, "should not get an error for this range")
	assert.NotEmpty(t, externalIP, "should get a valid IP address in the empty range")

	return
}

func TestVirtualServiceHttpCRUDE(t *testing.T) {

	// NOTE: Create a dummy virtual service before running this test. Time to create the first virtual service may
	// bottle-neck the test execution.

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := GetTestVCDClient(cloudConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	// Avoid RDE updates by using clusterID which has `NoRdePrefix` prefix
	cloudConfig.ClusterID = fmt.Sprintf("%s-%s", uuid.New().String(), "NoRdePrefix")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, cloudConfig.LB.VDCNetwork, cloudConfig.LB.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	lbPoolName := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	lbPoolRef, err := gm.CreateLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234)
	assert.NoError(t, err, "Unable to create lb pool")

	segRef, err := gm.GetLoadBalancerSEG(ctx)
	assert.NoError(t, err, "Unable to get ServiceEngineGroup ref")

	virtualServiceName := fmt.Sprintf("test-virtual-service-%s", uuid.New().String())
	internalIP := "2.3.4.5"
	var vsRef *swagger.EntityReference
	for i := 0; i < BusyRetries; i++ {
		vsRef, err = gm.CreateVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
			internalIP, "HTTP", 80, false, "")
		if err != nil {
			if _, ok := err.(*VirtualServicePendingError); !ok {
				break
			}
		}
		time.Sleep(3*time.Second)
	}
	assert.NoError(t, err, "Unable to create virtual service")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	vsRefObtained, err := gm.GetVirtualService(ctx, virtualServiceName)
	assert.NoError(t, err, "Unable to get virtual service ref")
	require.NotNil(t, vsRefObtained, "Virtual service reference should not be nil")
	assert.Equal(t, vsRefObtained.Name, vsRef.Name, "Virtual service reference name should match")
	assert.NotEmpty(t, vsRefObtained.Id, "Virtual service ID should not be empty")

	// repeated creation should not fail
	for i := 0; i < BusyRetries; i++ {
		vsRef, err = gm.CreateVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
			internalIP, "HTTP", 80, false, "")
		if err != nil {
			if _, ok := err.(*VirtualServicePendingError); !ok {
				break
			}
		}
	}
	assert.NoError(t, err, "Unable to create virtual service for the second time")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	_, err = gm.UpdateVirtualServicePort(ctx, virtualServiceName, 8080)
	assert.NoError(t, err, "Unable to update external port")

	// repeated update should not fail
	_, err = gm.UpdateVirtualServicePort(ctx, virtualServiceName, 8080)
	assert.NoError(t, err, "Repeated update to external port should not fail")

	_, err = gm.UpdateVirtualServicePort(ctx, virtualServiceName+"-invalid", 8080)
	assert.Error(t, err, "Update virtual service on a non-existent virtual service should fail")

	err = gm.DeleteVirtualService(ctx, virtualServiceName, true)
	assert.NoError(t, err, "Unable to delete Virtual Service")

	err = gm.DeleteVirtualService(ctx, virtualServiceName, true)
	assert.Error(t, err, "Should fail when deleting non-existing Virtual Service")

	err = gm.DeleteVirtualService(ctx, virtualServiceName, false)
	assert.NoError(t, err, "Should not fail when deleting non-existing Virtual Service")

	vsRefObtained, err = gm.GetVirtualService(ctx, virtualServiceName)
	assert.NoError(t, err, "Get should not fail when Virtual Service is absent")
	assert.Nil(t, vsRefObtained, "Deleted Virtual Service reference should be nil")

	err = gm.DeleteLoadBalancerPool(ctx, lbPoolName, true)
	assert.NoError(t, err, "Should not fail when deleting lb pool")

	return
}

func TestVirtualServiceHttpsCRUDE(t *testing.T) {

	// NOTE: Create a dummy virtual service before running this test. Time to create the first virtual service may
	// bottle-neck the test execution.

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := GetTestVCDClient(cloudConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	// Avoid RDE updates by using clusterID which has `NoRdePrefix` prefix
	cloudConfig.ClusterID = fmt.Sprintf("%s-%s", uuid.New().String(), NoRdePrefix)

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, cloudConfig.LB.VDCNetwork, cloudConfig.LB.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	lbPoolName := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())

	lbPoolRef, err := gm.CreateLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234)
	assert.NoError(t, err, "Unable to create lb pool")

	segRef, err := gm.GetLoadBalancerSEG(ctx)
	assert.NoError(t, err, "Unable to get ServiceEngineGroup ref")

	//externalIP := "11.12.13.14"
	internalIP := "3.4.5.6"
	virtualServiceName := fmt.Sprintf("test-virtual-service-https-%s", uuid.New().String())
	certName := cloudConfig.LB.CertificateAlias
	if certName == "" {
		certName = fmt.Sprintf("%s-cert", cloudConfig.ClusterID)
	}

	var vsRef *swagger.EntityReference
	for i := 0; i < BusyRetries; i++ {
		vsRef, err = gm.CreateVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
			internalIP, "HTTP", 80, false, "")
		if err != nil {
			if _, ok := err.(*VirtualServicePendingError); !ok {
				break
			}
		}
	}
	assert.NoError(t, err, "Unable to create virtual service")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	vsRefObtained, err := gm.GetVirtualService(ctx, virtualServiceName)
	assert.NoError(t, err, "Unable to get virtual service ref")
	require.NotNil(t, vsRefObtained, "Virtual service reference should not be nil")
	assert.Equal(t, vsRefObtained.Name, vsRef.Name, "Virtual service reference name should match")
	assert.NotEmpty(t, vsRefObtained.Id, "Virtual service ID should not be empty")

	// repeated creation should not fail
	vsRef, err = gm.CreateVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
		internalIP, "HTTPS", 443, true, certName)
	assert.NoError(t, err, "Unable to create virtual service for the second time")
	require.NotNil(t, vsRef, "VirtualServiceRef should not be nil")
	assert.Equal(t, virtualServiceName, vsRef.Name, "Virtual Service name should match")

	// update and delete calls might error out if virtual services are busy. Retry if the error is caused by the busy virtual services
	_, err = gm.UpdateVirtualServicePort(ctx, virtualServiceName, 8443)
	assert.NoError(t, err, "Unable to update external port")

	// repeated update should not fail
	_, err = gm.UpdateVirtualServicePort(ctx, virtualServiceName, 8443)
	assert.NoError(t, err, "Repeated update to external port should not fail")

	// update of invalid virtual service
	_, err = gm.UpdateVirtualServicePort(ctx, virtualServiceName+"-invalid\n", 8443)
	assert.Error(t, err, "Update virtual service on a non-existent virtual service should fail")

	err = gm.DeleteVirtualService(ctx, virtualServiceName, true)
	assert.NoError(t, err, "Unable to delete Virtual Service")

	err = gm.DeleteVirtualService(ctx, virtualServiceName, true)
	assert.Error(t, err, "Should fail when deleting non-existing Virtual Service")

	err = gm.DeleteVirtualService(ctx, virtualServiceName, false)
	assert.NoError(t, err, "Should not fail when deleting non-existing Virtual Service")

	vsRefObtained, err = gm.GetVirtualService(ctx, virtualServiceName)
	assert.NoError(t, err, "Get should not fail when Virtual Service is absent")
	assert.Nil(t, vsRefObtained, "Deleted Virtual Service reference should be nil")

	err = gm.DeleteLoadBalancerPool(ctx, lbPoolName, true)
	assert.NoError(t, err, "Should not fail when deleting lb pool")

	return
}
