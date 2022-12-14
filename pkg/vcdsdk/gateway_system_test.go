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
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/util"
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

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := getTestVCDClient(vcdConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, vcdConfig.OvdcNetwork, vcdConfig.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	require.NotNil(t, gm.GatewayRef, "Gateway reference should not be nil")
	assert.NotEmpty(t, gm.GatewayRef.Name, "Gateway Name should not be empty")
	assert.NotEmpty(t, gm.GatewayRef.Id, "Gateway Id should not be empty")

	// Missing network name should be reported
	gatewayManager, err := NewGatewayManager(ctx, vcdClient, "", vcdConfig.VIPSubnet)
	assert.Error(t, err, "Should get error for unknown network")
	assert.Nil(t, gatewayManager, "gateway manager should be nil when erroring out")

	return
}

func TestDNATRuleCRUDE(t *testing.T) {

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := getTestVCDClient(vcdConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, vcdConfig.OvdcNetwork, vcdConfig.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	dnatRuleName := fmt.Sprintf("test-dnat-rule-%s", uuid.New().String())
	appPortProfile, err := gm.CreateAppPortProfile(GetAppPortProfileName(dnatRuleName), 80)
	assert.NoError(t, err, "Unable to create App Port Profile")

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

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := getTestVCDClient(vcdConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, vcdConfig.OvdcNetwork, vcdConfig.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	lbPoolName := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	lbPoolRef, err := gm.CreateLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234, "HTTP")
	assert.NoError(t, err, "Unable to create lb pool")
	require.NotNil(t, lbPoolRef, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolName, lbPoolRef.Name, "LB Pool name should match")

	// repeated creation should not fail
	lbPoolRef, err = gm.CreateLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234, "HTTP")
	assert.NoError(t, err, "Unable to create lb pool for the second time")
	require.NotNil(t, lbPoolRef, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolName, lbPoolRef.Name, "LB Pool name should match")

	lbPoolRefObtained, err := gm.getLoadBalancerPool(ctx, lbPoolName)
	assert.NoError(t, err, "Unable to get lbPool ref")
	require.NotNil(t, lbPoolRefObtained, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolRefObtained.Name, lbPoolRef.Name, "LB Pool name should match")
	assert.NotEmpty(t, lbPoolRefObtained.Id, "LB Pool ID should not be empty")

	updatedIps := []string{"5.5.5.5"}
	lbPoolRefUpdated, err := gm.UpdateLoadBalancerPool(ctx, lbPoolName, updatedIps, 55555, "HTTP")
	assert.NoError(t, err, "No lbPool ref for updated lbPool")
	require.NotNil(t, lbPoolRefUpdated, "LB Pool reference should not be nil")
	assert.Equal(t, lbPoolRefUpdated.Name, lbPoolRef.Name, "LB Pool name should match")
	assert.NotEmpty(t, lbPoolRefUpdated.Id, "LB Pool ID should not be empty")

	// repeated update should work
	lbPoolRefUpdated, err = gm.UpdateLoadBalancerPool(ctx, lbPoolName, updatedIps, 55555, "HTTP")
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

	lbPoolRef, err = gm.UpdateLoadBalancerPool(ctx, lbPoolName, updatedIps, 55555, "HTTP")
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

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := getTestVCDClient(vcdConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, vcdConfig.OvdcNetwork, vcdConfig.VIPSubnet)
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

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := getTestVCDClient(vcdConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
		"subnet":       "",
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, vcdConfig.OvdcNetwork, vcdConfig.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	validSubnet := vcdConfig.VIPSubnet
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

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := getTestVCDClient(vcdConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, vcdConfig.OvdcNetwork, vcdConfig.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	lbPoolName := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	lbPoolRef, err := gm.CreateLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234, "HTTP")
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
		time.Sleep(3 * time.Second)
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


	_, err = gm.UpdateVirtualService(ctx, virtualServiceName, "", 8080, true)
	assert.NoError(t, err, "Unable to update external port")

	// repeated update should not fail
	_, err = gm.UpdateVirtualService(ctx, virtualServiceName, "",8080, true)
	assert.NoError(t, err, "Repeated update to external port should not fail")

	_, err = gm.UpdateVirtualService(ctx, virtualServiceName+"-invalid", "", 8080, true)
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

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := getTestVCDClient(vcdConfig, nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	// Avoid RDE updates by using clusterID which has `NoRdePrefix` prefix
	vcdConfig.ClusterID = fmt.Sprintf("%s-%s", uuid.New().String(), NoRdePrefix)

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, vcdConfig.OvdcNetwork, vcdConfig.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	lbPoolName := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	lbPoolRef, err := gm.CreateLoadBalancerPool(ctx, lbPoolName, []string{"1.2.3.4", "1.2.3.5"}, 31234, "HTTPS")
	assert.NoError(t, err, "Unable to create lb pool")

	segRef, err := gm.GetLoadBalancerSEG(ctx)
	assert.NoError(t, err, "Unable to get ServiceEngineGroup ref")

	//externalIP := "11.12.13.14"
	internalIP := "3.4.5.6"
	virtualServiceName := fmt.Sprintf("test-virtual-service-https-%s", uuid.New().String())
	certName := vcdConfig.CertificateAlias
	if certName == "" {
		certName = fmt.Sprintf("%s-cert", vcdConfig.ClusterID)
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
	_, err = gm.UpdateVirtualService(ctx, virtualServiceName, "", 8443, true)
	assert.NoError(t, err, "Unable to update external port")

	// repeated update should not fail
	_, err = gm.UpdateVirtualService(ctx, virtualServiceName, "", 8443, true)
	assert.NoError(t, err, "Repeated update to external port should not fail")

	// update of invalid virtual service
	_, err = gm.UpdateVirtualService(ctx, virtualServiceName+"-invalid\n", "", 8443, true)
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

func TestLoadBalancerCRUDE(t *testing.T) {

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	testConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := getTestVCDClient(testConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, testConfig.OvdcNetwork, testConfig.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	virtualServiceNamePrefix := fmt.Sprintf("test-virtual-service-https-%s", uuid.New().String())
	lbPoolNamePrefix := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	certName := testConfig.CertificateAlias
	if certName == "" {
		certName = fmt.Sprintf("%s-cert", testConfig.ClusterID)
	}
	portDetailsList := []PortDetails{
		{
			PortSuffix:   `http`,
			ExternalPort: 80,
			InternalPort: 31234,
			Protocol:     "HTTP",
			UseSSL:       false,
		},
		{
			PortSuffix:   `https`,
			ExternalPort: 443,
			InternalPort: 31235,
			Protocol:     "HTTPS",
			UseSSL:       true,
			CertAlias:    certName,
		},
	}
	freeIP := ""

	oneArm := &OneArm{
		StartIP: "192.168.8.2",
		EndIP:   "192.168.8.100",
	}
	freeIP, err = gm.CreateLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix, []string{"1.2.3.4", "1.2.3.5"}, portDetailsList, oneArm, false, nil, "", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Load Balancer should be created")
	assert.NotEmpty(t, freeIP, "There should be a non-empty IP returned")

	virtualServiceNameHttp := fmt.Sprintf("%s-http", virtualServiceNamePrefix)

	// using _ to escape the *util.AllocatedResourcesMap
	lbPoolNameHttp := fmt.Sprintf("%s-http", lbPoolNamePrefix)
	freeIPObtained, _, err := gm.GetLoadBalancer(ctx, virtualServiceNameHttp, lbPoolNameHttp, oneArm)
	assert.NoError(t, err, "Load Balancer should be found")
	assert.Equal(t, freeIP, freeIPObtained, "The IPs should match")

	virtualServiceNameHttps := fmt.Sprintf("%s-https", virtualServiceNamePrefix)
	lbPoolNameHttps := fmt.Sprintf("%s-https", lbPoolNamePrefix)
	freeIPObtained, _, err = gm.GetLoadBalancer(ctx, virtualServiceNameHttps, lbPoolNameHttps, oneArm)
	assert.NoError(t, err, "Load Balancer should be found")
	assert.Equal(t, freeIP, freeIPObtained, "The IPs should match")

	freeIP, err = gm.CreateLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix, []string{"1.2.3.4", "1.2.3.5"}, portDetailsList, oneArm, false, nil, "", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Load Balancer should be created even on second attempt")
	assert.NotEmpty(t, freeIP, "There should be a non-empty IP returned")

	updatedIps := []string{"5.5.5.5"}
	updatedInternalPort := int32(55555)
	// update IPs and internal port
	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, "", updatedInternalPort, 80, nil, false, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, "", updatedInternalPort, 443, nil, false, "HTTPS", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	// update external port only
	updatedExternalPortHttp := int32(8080)
	updatedExternalPortHttps := int32(8443)

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, "", updatedInternalPort, updatedExternalPortHttp, nil, false, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, "", updatedInternalPort, updatedExternalPortHttps, nil, false, "HTTPS", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	// No error on repeated update
	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, "", updatedInternalPort, updatedExternalPortHttp, nil, false, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, "", updatedInternalPort, updatedExternalPortHttps, nil, false, "HTTPS", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	_, err = gm.DeleteLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, oneArm, &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Load Balancer should be deleted")

	freeIPObtained, _, err = gm.GetLoadBalancer(ctx, virtualServiceNameHttp, lbPoolNameHttp, oneArm)
	assert.NoError(t, err, "Load Balancer should not be found")
	assert.Empty(t, freeIPObtained, "The VIP should not be found")

	freeIPObtained, _, err = gm.GetLoadBalancer(ctx, virtualServiceNameHttps, lbPoolNameHttps, oneArm)
	assert.NoError(t, err, "Load Balancer should not be found")
	assert.Empty(t, freeIPObtained, "The VIP should not be found")

	_, err = gm.DeleteLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, oneArm, &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Repeated deletion of Load Balancer should not fail")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, "", updatedInternalPort, 80, nil, false, "HTTP", &util.AllocatedResourcesMap{})
	assert.Error(t, err, "updating deleted HTTP Load Balancer should be an error")
	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"https", updatedIps, "", updatedInternalPort, 43, nil, false, "HTTPS", &util.AllocatedResourcesMap{})
	assert.Error(t, err, "updating deleted HTTPS Load Balancer should be an error")

	return
}

func TestLoadBalancer_ExplicitLBIP_OneArmDisabled_CRUDE(t *testing.T) {

	// NOTE: make sure to specify a free external IP address in testdata/vcd_info.yaml
	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	testConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")
	assert.NotEmpty(t, testConfig.FreeLoadBalancerIP, "An external IP address should be specified in th test config.")

	vcdClient, err := getTestVCDClient(testConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, testConfig.OvdcNetwork, testConfig.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	virtualServiceNamePrefix := fmt.Sprintf("test-virtual-service-https-%s", uuid.New().String())
	lbPoolNamePrefix := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	certName := testConfig.CertificateAlias
	if certName == "" {
		certName = fmt.Sprintf("%s-cert", testConfig.ClusterID)
	}
	portDetailsList := []PortDetails{
		{
			PortSuffix:   `http`,
			ExternalPort: 80,
			InternalPort: 31234,
			Protocol:     "HTTP",
			UseSSL:       false,
		},
		{
			PortSuffix:   `https`,
			ExternalPort: 443,
			InternalPort: 31235,
			Protocol:     "HTTPS",
			UseSSL:       true,
			CertAlias:    certName,
		},
	}
	freeIP := testConfig.FreeLoadBalancerIP

	var oneArm *OneArm
	freeIP, err = gm.CreateLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix, []string{"1.2.3.4", "1.2.3.5"}, portDetailsList, oneArm, true, nil, testConfig.FreeLoadBalancerIP, &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Load Balancer should be created")
	assert.NotEmpty(t, freeIP, "There should be a non-empty IP returned")
	assert.Equal(t, freeIP, testConfig.FreeLoadBalancerIP, "the provided external IP address should be the same as the load balancer IP address.")

	virtualServiceNameHttp := fmt.Sprintf("%s-http", virtualServiceNamePrefix)

	// using _ to escape the *util.AllocatedResourcesMap
	lbPoolNameHttp := fmt.Sprintf("%s-http", lbPoolNamePrefix)
	freeIPObtained, _, err := gm.GetLoadBalancer(ctx, virtualServiceNameHttp, lbPoolNameHttp, oneArm)
	assert.NoError(t, err, "Load Balancer should be found")
	assert.Equal(t, freeIP, freeIPObtained, "The IPs should match")

	virtualServiceNameHttps := fmt.Sprintf("%s-https", virtualServiceNamePrefix)
	lbPoolNameHttps := fmt.Sprintf("%s-https", lbPoolNamePrefix)
	freeIPObtained, _, err = gm.GetLoadBalancer(ctx, virtualServiceNameHttps, lbPoolNameHttps, oneArm)
	assert.NoError(t, err, "Load Balancer should be found")
	assert.Equal(t, freeIP, freeIPObtained, "The IPs should match")

	freeIP, err = gm.CreateLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix, []string{"1.2.3.4", "1.2.3.5"}, portDetailsList, oneArm, true, nil, testConfig.FreeLoadBalancerIP, &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Load Balancer should be created even on second attempt")
	assert.NotEmpty(t, freeIP, "There should be a non-empty IP returned")
	assert.Equal(t, freeIP, testConfig.FreeLoadBalancerIP, "the provided external IP address should be the same as the load balancer IP address.")

	updatedIps := []string{"5.5.5.5"}
	updatedInternalPort := int32(55555)
	// update IPs and internal port
	freeIP, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, 80, nil, true, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")
	assert.Equal(t, freeIP, testConfig.FreeLoadBalancerIP, "IPs should match")

	freeIP, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, 443, nil, true, "HTTPS", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")
	assert.Equal(t, freeIP, testConfig.FreeLoadBalancerIP, "IPs should match")

	// update external port only
	updatedExternalPortHttp := int32(8080)
	updatedExternalPortHttps := int32(8443)

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, updatedExternalPortHttp, nil, true, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, updatedExternalPortHttps, nil, true, "HTTPS", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	// No error on repeated update
	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, updatedExternalPortHttp, nil, true, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, updatedExternalPortHttps, nil, true, "HTTPS", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	// Update LB IP address of the Load Balancer
	newLBIP := "192.168.100.20"
	lbIP, err := gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, newLBIP, updatedInternalPort, updatedExternalPortHttp, nil, true, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")
	assert.Equal(t, lbIP, newLBIP, "updated external IP address should match the value specified")

	_, err = gm.DeleteLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, oneArm, &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Load Balancer should be deleted")

	freeIPObtained, _, err = gm.GetLoadBalancer(ctx, virtualServiceNameHttp, lbPoolNameHttp, oneArm)
	assert.NoError(t, err, "Load Balancer should not be found")
	assert.Empty(t, freeIPObtained, "The VIP should not be found")

	freeIPObtained, _, err = gm.GetLoadBalancer(ctx, virtualServiceNameHttps, lbPoolNameHttps, oneArm)
	assert.NoError(t, err, "Load Balancer should not be found")
	assert.Empty(t, freeIPObtained, "The VIP should not be found")

	_, err = gm.DeleteLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, oneArm, &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Repeated deletion of Load Balancer should not fail")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, 80, nil, true, "HTTP", &util.AllocatedResourcesMap{})
	assert.Error(t, err, "updating deleted HTTP Load Balancer should be an error")
	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"https", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, 43, nil, true, "HTTPS", &util.AllocatedResourcesMap{})
	assert.Error(t, err, "updating deleted HTTPS Load Balancer should be an error")

	return
}

func TestLoadBalancer_ExplicitLBIP_OneArmEnabled_CRUDE(t *testing.T) {

	// NOTE: make sure to specify a free external IP address in testdata/vcd_info.yaml
	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	testConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")
	assert.NotEmpty(t, testConfig.FreeLoadBalancerIP, "An external IP address should be specified in th test config.")

	vcdClient, err := getTestVCDClient(testConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	gm, err := NewGatewayManager(ctx, vcdClient, testConfig.OvdcNetwork, testConfig.VIPSubnet)
	assert.NoError(t, err, "gateway manager should be created without error")

	virtualServiceNamePrefix := fmt.Sprintf("test-virtual-service-https-%s", uuid.New().String())
	lbPoolNamePrefix := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	certName := testConfig.CertificateAlias
	if certName == "" {
		certName = fmt.Sprintf("%s-cert", testConfig.ClusterID)
	}
	portDetailsList := []PortDetails{
		{
			PortSuffix:   `http`,
			ExternalPort: 80,
			InternalPort: 31234,
			Protocol:     "HTTP",
			UseSSL:       false,
		},
		{
			PortSuffix:   `https`,
			ExternalPort: 443,
			InternalPort: 31235,
			Protocol:     "HTTPS",
			UseSSL:       true,
			CertAlias:    certName,
		},
	}
	freeIP := ""

	oneArm := &OneArm{
		StartIP: "192.168.8.2",
		EndIP:   "192.168.8.100",
	}

	freeIP, err = gm.CreateLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix, []string{"1.2.3.4", "1.2.3.5"}, portDetailsList, oneArm, true, nil, testConfig.FreeLoadBalancerIP, &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Load Balancer should be created")
	assert.NotEmpty(t, freeIP, "There should be a non-empty IP returned")
	assert.Equal(t, freeIP, testConfig.FreeLoadBalancerIP, "the provided external IP address should be the same as the load balancer IP address.")

	virtualServiceNameHttp := fmt.Sprintf("%s-http", virtualServiceNamePrefix)

	// using _ to escape the *util.AllocatedResourcesMap
	lbPoolNameHttp := fmt.Sprintf("%s-http", lbPoolNamePrefix)
	freeIPObtained, _, err := gm.GetLoadBalancer(ctx, virtualServiceNameHttp, lbPoolNameHttp, oneArm)
	assert.NoError(t, err, "Load Balancer should be found")
	assert.Equal(t, freeIP, freeIPObtained, "The IPs should match")

	virtualServiceNameHttps := fmt.Sprintf("%s-https", virtualServiceNamePrefix)
	lbPoolNameHttps := fmt.Sprintf("%s-https", lbPoolNamePrefix)
	freeIPObtained, _, err = gm.GetLoadBalancer(ctx, virtualServiceNameHttps, lbPoolNameHttps, oneArm)
	assert.NoError(t, err, "Load Balancer should be found")
	assert.Equal(t, freeIP, freeIPObtained, "The IPs should match")

	freeIP, err = gm.CreateLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix, []string{"1.2.3.4", "1.2.3.5"}, portDetailsList, oneArm, true, nil, testConfig.FreeLoadBalancerIP, &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Load Balancer should be created even on second attempt")
	assert.NotEmpty(t, freeIP, "There should be a non-empty IP returned")

	updatedIps := []string{"5.5.5.5"}
	updatedInternalPort := int32(55555)
	// update IPs and internal port
	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, 80, oneArm, true, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, 443, oneArm, true, "HTTPS", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	// update external port only
	updatedExternalPortHttp := int32(8080)
	updatedExternalPortHttps := int32(8443)

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, updatedExternalPortHttp, oneArm, true, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, updatedExternalPortHttps, oneArm, true, "HTTPS", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	// No error on repeated update
	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, updatedExternalPortHttp, oneArm, true, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, updatedExternalPortHttps, oneArm, true, "HTTPS", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	// No error on updating the loadbalancer IP
	newLBIP := "192.168.100.20"
	lbIP, err := gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, newLBIP, updatedInternalPort, updatedExternalPortHttp, oneArm, true, "HTTP", &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "HTTP Load Balancer should be updated")
	assert.Equal(t, newLBIP, lbIP, "The external IP for the load balancer should be updated")

	_, err = gm.DeleteLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, oneArm, &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Load Balancer should be deleted")

	freeIPObtained, _, err = gm.GetLoadBalancer(ctx, virtualServiceNameHttp, lbPoolNameHttp, oneArm)
	assert.NoError(t, err, "Load Balancer should not be found")
	assert.Empty(t, freeIPObtained, "The VIP should not be found")

	freeIPObtained, _, err = gm.GetLoadBalancer(ctx, virtualServiceNameHttps, lbPoolNameHttps, oneArm)
	assert.NoError(t, err, "Load Balancer should not be found")
	assert.Empty(t, freeIPObtained, "The VIP should not be found")

	_, err = gm.DeleteLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, oneArm, &util.AllocatedResourcesMap{})
	assert.NoError(t, err, "Repeated deletion of Load Balancer should not fail")

	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, 80, oneArm, true, "HTTP", &util.AllocatedResourcesMap{})
	assert.Error(t, err, "updating deleted HTTP Load Balancer should be an error")
	_, err = gm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"https", updatedIps, testConfig.FreeLoadBalancerIP, updatedInternalPort, 43, oneArm, true, "HTTPS", &util.AllocatedResourcesMap{})
	assert.Error(t, err, "updating deleted HTTPS Load Balancer should be an error")

	return
}

