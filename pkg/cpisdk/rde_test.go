package cpisdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/cloud-provider-for-cloud-director/release"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/yaml"
	"net/http"
	"path/filepath"
	"testing"
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

	rm := vcdsdk.NewRDEManager(vcdClient, cloudConfig.ClusterID, release.CloudControllerManagerName, release.CpiVersion)
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

func TestUpgradeCPISectionInStatus(t *testing.T) {
	type TestCase struct {
		Name           string
		StatusMap      map[string]interface{}
		ExpectedStatus map[string]interface{}
	}
	testCaseList := []TestCase{
		{
			Name: "no modifications to status map outside CPI status",
			StatusMap: map[string]interface{}{
				"cpi": map[string]interface{}{
					"vcdResourceSet": []vcdsdk.VCDResource{
						{
							Type: vcdsdk.VcdResourceVirtualService,
							Name: "virtual-service-1",
							ID:   "12345",
						},
					},
				},
				"csi": map[string]interface{}{
					"key1": "value1",
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
			},
			ExpectedStatus: map[string]interface{}{
				"cpi": map[string]interface{}{
					"name": "cloud-controller-manager",
					"vcdResourceSet": []interface{}{
						map[string]interface{}{
							"id":   "12345",
							"name": "virtual-service-1",
							"type": "virtual-service",
						},
					},
					"version": release.CpiVersion,
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
				"csi": map[string]interface{}{
					"key1": "value1",
				},
			},
		},
	}

	for _, tc := range testCaseList {
		updatedStatus, err := UpgradeCPISectionInStatus(tc.StatusMap)
		assert.NoError(t, err, tc.Name)
		assert.Equal(t, tc.ExpectedStatus, updatedStatus, tc.Name)

		actualJson, err := convertToJson(updatedStatus)
		assert.NoError(t, err, tc.Name, "expected no error converting updatedStatusMap to json")
		expectedJson, err := convertToJson(tc.ExpectedStatus)
		assert.NoError(t, err, tc.Name, "expected no error when converting expected status to json")

		assert.JSONEqf(t, expectedJson, actualJson, tc.Name)
	}
}
