package cpisdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient_37_2"
	"github.com/vmware/cloud-provider-for-cloud-director/release"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/yaml"
	"net/http"
	"path/filepath"
	"strings"
	"testing"
)

const (
	CAPVCDEntityTypeVersion = "1.1.0"
	TestEntityName          = "test"
)

func foundStringInSlice(find string, slice []string) bool {
	for _, currElement := range slice {
		if currElement == find {
			return true
		}
	}
	return false
}

func convertToJson(obj interface{}) (string, error) {
	objByteArr, err := json.Marshal(&obj)
	if err != nil {
		return "", fmt.Errorf("error while marshaling object: [%v]", err)
	}

	return string(objByteArr), nil
}

func TestUpdateRDEUsingEtag(t *testing.T) {
	// TODO: This test will currently fail unless the code below is uncommented. Refer to VCDA-3600

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	vcdClient, err := getTestVCDClient(cloudConfig, map[string]interface{}{
		"user":    authDetails.Username,
		"secret":  authDetails.Password,
		"userOrg": authDetails.UserOrg,
	})

	ctx := context.Background()

	rm := vcdsdk.NewRDEManager(vcdClient, release.CloudControllerManagerName, release.Version, cloudConfig.ClusterID)
	cpiRdeManager := NewCPIRDEManager(rm)
	// get rde Vips
	rdeVips1, etag1, defEnt1, err := cpiRdeManager.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips on first attempt")
	rdeVips2, etag2, defEnt2, err := cpiRdeManager.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips on second attempt")
	assert.Equal(t, etag1, etag2, "etags from consecutive RDE GET calls should match")
	origRdeVips := make([]string, len(rdeVips1))
	copy(origRdeVips, rdeVips1)

	// Test successfully updating using first etag
	addIp1 := "1.2.3.4"
	addIp2 := "2.3.4.5"
	updatedRdeVips1 := append(rdeVips1, addIp1)
	httpResponse1, err := cpiRdeManager.updateRDEVirtualIps(ctx, updatedRdeVips1, etag1, defEnt1)
	assert.NoError(t, err, "RDE should be updated")
	assert.Equal(t, http.StatusOK, httpResponse1.StatusCode, "RDE update status code should be 200 (OK)")
	rdeVips3, _, _, err := cpiRdeManager.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips successfully")
	assert.True(t, foundStringInSlice(addIp1, rdeVips3), "ip [%s] should be found in rde vips", addIp1)

	// Test adding addIp2 with outdated etag
	updatedRdeVips2 := append(rdeVips2, addIp2)
	httpResponse2, err := cpiRdeManager.updateRDEVirtualIps(ctx, updatedRdeVips2, etag2, defEnt2)
	assert.Error(t, err, "Should have an error updating RDE with outdated etag")
	assert.Equal(t, http.StatusPreconditionFailed, httpResponse2.StatusCode, "RDE update status code should be 412 (Precondition failed)")
	rdeVips3, etag3, defEnt3, err := cpiRdeManager.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips successfully")
	assert.False(t, foundStringInSlice(addIp2, rdeVips3), "ip [%s] should not be found in rde vips", addIp2)

	// Try adding addIp2 with correct etag
	updatedRdeVips3 := append(rdeVips3, addIp2)
	httpResponse3, err := cpiRdeManager.updateRDEVirtualIps(ctx, updatedRdeVips3, etag3, defEnt3)
	assert.NoError(t, err, "RDE should be updated")
	assert.Equal(t, http.StatusOK, httpResponse3.StatusCode, "RDE update status code should be 200 (OK)")
	rdeVips4, etag4, defEnt4, err := cpiRdeManager.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips successfully")
	assert.True(t, foundStringInSlice(addIp2, rdeVips4), "ip [%s] should be found in rde vips", addIp2)

	// reset RDE vips to original state
	httpResponse4, err := cpiRdeManager.updateRDEVirtualIps(ctx, rdeVips1, etag4, defEnt4)
	assert.NoError(t, err, "RDE should be updated")
	assert.Equal(t, http.StatusOK, httpResponse4.StatusCode, "RDE update status code should be 200 (OK)")
	// no check to ensure ip's removed because they may have been previously present in the RDE vips
	rdeVips5, _, _, err := cpiRdeManager.GetRDEVirtualIps(ctx)
	assert.NoError(t, err, "Should retrieve RDE vips to check added ips are removed")
	assert.False(t, foundStringInSlice(addIp1, rdeVips5), "ip [%s] should not be found in rde vips", addIp1)
	assert.False(t, foundStringInSlice(addIp2, rdeVips5), "ip [%s] should not be found in rde vips", addIp2)
}

func cleanUpEntitiesWithName(t *testing.T, vcdClient *vcdsdk.Client, entityName string, failMessage string) {
	nameFilter := &swaggerClient.DefinedEntityApiGetDefinedEntitiesByEntityTypeOpts{
		Filter: optional.NewString(fmt.Sprintf("name==%s", entityName)),
	}
	definedEntities, resp, err := vcdClient.APIClient.DefinedEntityApi.GetDefinedEntitiesByEntityType(context.TODO(),
		vcdsdk.CAPVCDEntityTypeVendor, vcdsdk.CAPVCDEntityTypeNss, CAPVCDEntityTypeVersion, "", 1, 25, nameFilter)
	assert.NoError(t, err, "expected no error executing list entities by entity type with name filter", "message", failMessage)
	assert.NotNil(t, resp, "list entity by entity type response should not be nil", "message", failMessage)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "invalid response for list entity by entity type", "message", failMessage)
	for _, e := range definedEntities.Values {
		// resolve before delete - may fail with resolution errors.
		entityState, _, _ := vcdClient.APIClient.DefinedEntityApi.ResolveDefinedEntity(context.TODO(), e.Id, "")
		if entityState.State != "RESOLVED" {
			fmt.Println("entity resolution failed for RDE ", entityState.Id, " with message: ", entityState.Message)
		}
		resp, err := vcdClient.APIClient.DefinedEntityApi.DeleteDefinedEntity(context.TODO(), e.Id, "", nil)
		assert.NoError(t, err, "expected no error cleaning up the defined entity", "message", failMessage)
		assert.NotNil(t, resp, "did not expect a nil response while cleaning up defined entity by name", "message", failMessage)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode, "got unexpected response status code",
			"message", failMessage)
	}
}

func TestAddVIPToVCDResourceSet(t *testing.T) {
	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := getTestVCDClient(cloudConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	rdeManager := vcdsdk.RDEManager{
		Client:                 vcdClient,
		StatusComponentName:    release.CloudControllerManagerName,
		StatusComponentVersion: release.Version,
		ClusterID:              "", // will be filled up during tests
	}

	assert.True(t, rdeManager.IsCapvcdEntityTypeRegistered(CAPVCDEntityTypeVersion),
		fmt.Sprintf("entity type capvcdCluster:[%s] is not registered", CAPVCDEntityTypeVersion))

	// clean up before the test -
	cleanUpEntitiesWithName(t, vcdClient, TestEntityName, "initial cleanup")

	testVip := "1.2.3.4"
	var vipsIf interface{}
	vipsIf = []string{
		testVip,
	}
	entityTypeID := strings.Join(
		[]string{
			vcdsdk.CAPVCDEntityTypePrefix,
			CAPVCDEntityTypeVersion,
		},
		":")
	type TestCase struct {
		EntityCreated             map[string]interface{}
		ExpectedEntityAfterUpdate map[string]interface{}
		Message                   string
	}
	vsName := "dummyVs"
	vsId := "dummyVsId"
	testCaseArr := []TestCase{
		TestCase{
			EntityCreated: map[string]interface{}{
				"status": map[string]interface{}{
					"key1": "value1",
					"cpi": map[string]interface{}{
						"name":    "oldName",
						"version": "oldVersion",
					},
					"virtualIPs": vipsIf,
				},
			},
			ExpectedEntityAfterUpdate: map[string]interface{}{
				"status": map[string]interface{}{
					"key1": "value1",
					"cpi": map[string]interface{}{
						"name":    release.CloudControllerManagerName,
						"version": release.Version,
						"vcdResourceSet": []vcdsdk.VCDResource{
							vcdsdk.VCDResource{
								Type: vcdsdk.VcdResourceVirtualService,
								Name: vsName,
								ID:   vsId,
								AdditionalDetails: map[string]interface{}{
									"virtualIP": testVip,
								},
							},
						},
					},
				},
			},
			Message: "add to vcd resource set and remove virtualIPs section if empty",
		},
		TestCase{
			EntityCreated: map[string]interface{}{
				"status": map[string]interface{}{
					"key1": "value1",
					"cpi": map[string]interface{}{
						"name":    "oldName",
						"version": "oldVersion",
					},
				},
			},
			ExpectedEntityAfterUpdate: map[string]interface{}{
				"status": map[string]interface{}{
					"key1": "value1",
					"cpi": map[string]interface{}{
						"name":    release.CloudControllerManagerName,
						"version": release.Version,
						"vcdResourceSet": []vcdsdk.VCDResource{
							vcdsdk.VCDResource{
								Type: vcdsdk.VcdResourceVirtualService,
								Name: vsName,
								ID:   vsId,
								AdditionalDetails: map[string]interface{}{
									"virtualIP": testVip,
								},
							},
						},
					},
				},
			},
			Message: "Add VIP to VCD ResourceSet only if entity.status.virtualIPs is absent",
		},
		TestCase{
			EntityCreated: map[string]interface{}{
				"status": map[string]interface{}{
					"key1": "value1",
					"cpi": map[string]interface{}{
						"name":    "oldName",
						"version": "oldVersion",
					},
					"virtualIPs": []string{
						testVip,
						"2.3.4.5",
					},
				},
			},
			ExpectedEntityAfterUpdate: map[string]interface{}{
				"status": map[string]interface{}{
					"key1": "value1",
					"cpi": map[string]interface{}{
						"name":    release.CloudControllerManagerName,
						"version": release.Version,
						"vcdResourceSet": []vcdsdk.VCDResource{
							vcdsdk.VCDResource{
								Type: vcdsdk.VcdResourceVirtualService,
								Name: vsName,
								ID:   vsId,
								AdditionalDetails: map[string]interface{}{
									"virtualIP": testVip,
								},
							},
						},
					},
					"virtualIPs": []string{
						"2.3.4.5",
					},
				},
			},
			Message: "retain entity.status.virtualIPs section if empty after removal",
		},
		TestCase{
			EntityCreated: map[string]interface{}{
				"status": map[string]interface{}{
					"key1": "value1",
					"cpi": map[string]interface{}{
						"name":    "oldName",
						"version": "oldVersion",
					},
					"virtualIPs": []string{
						"3.4.5.6",
						"2.3.4.5",
					},
				},
			},
			ExpectedEntityAfterUpdate: map[string]interface{}{
				"status": map[string]interface{}{
					"key1": "value1",
					"cpi": map[string]interface{}{
						"name":    release.CloudControllerManagerName,
						"version": release.Version,
						"vcdResourceSet": []vcdsdk.VCDResource{
							vcdsdk.VCDResource{
								Type: vcdsdk.VcdResourceVirtualService,
								Name: vsName,
								ID:   vsId,
								AdditionalDetails: map[string]interface{}{
									"virtualIP": testVip,
								},
							},
						},
					},
					"virtualIPs": []string{
						"3.4.5.6",
						"2.3.4.5",
					},
				},
			},
			Message: "don't remove from entity.status.virtualIPs if relevant VIP is not present",
		},
	}
	for _, tc := range testCaseArr {
		// create entity
		entityToBeCreated := swaggerClient.DefinedEntity{
			Name:       TestEntityName,
			EntityType: entityTypeID,
			Entity:     tc.EntityCreated,
		}
		resp, err := vcdClient.APIClient.DefinedEntityApi.CreateDefinedEntity(context.TODO(), entityToBeCreated, entityTypeID, "", nil)
		assert.NoError(t, err, "expected no error creating defined entity", "test", tc.Message)
		assert.NotNil(t, resp, "did not expect a nil response while creating defined entity", "test", tc.Message)
		assert.Equal(t, http.StatusAccepted, resp.StatusCode, "status code is not as expected", "test", tc.Message)
		taskURL := resp.Header.Get("Location")
		task := govcd.NewTask(&vcdClient.VCDClient.Client)
		task.Task.HREF = taskURL
		err = task.Refresh()
		assert.NoError(t, err, "no error refreshing entity create task", "test", tc.Message)
		rdeID := task.Task.Owner.ID

		rdeManager.ClusterID = rdeID
		capvcdRdeManager := NewCPIRDEManager(&rdeManager)

		// execute test
		err = capvcdRdeManager.AddVIPToVCDResourceSet(context.TODO(), vsName, vsId, testVip)
		assert.NoError(t, err, "expected no error adding VIP to VCD resource set", "test", tc.Message)

		// verify information in the RDE
		definedEntity, resp, _, err := vcdClient.APIClient.DefinedEntityApi.GetDefinedEntity(context.TODO(), rdeID,
			"", nil)
		assert.NoError(t, err, "expected no error fetching defined entity after updating RDE with virtual service", "test", tc.Message)
		assert.NotNil(t, definedEntity, "expected defined entity to be not nil", "test", tc.Message)
		assert.NotNil(t, resp, "invalid response - got nil response", "test", tc.Message)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "unexpected response code", "test", tc.Message)

		actualJson, err := convertToJson(definedEntity.Entity)
		assert.NoError(t, err, tc.Message, "expected no error converting updatedStatusMap to json", "test", tc.Message)
		expectedJson, err := convertToJson(tc.ExpectedEntityAfterUpdate)
		assert.NoError(t, err, tc.Message, "expected no error when converting expected status to json", "test", tc.Message)
		assert.JSONEqf(t, expectedJson, actualJson, tc.Message, "test", tc.Message)

		// clean up entity
		cleanUpEntitiesWithName(t, vcdClient, TestEntityName, tc.Message)
	}
}
