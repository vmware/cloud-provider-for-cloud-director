package vcdsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func convertToJson(obj interface{}) (string, error) {
	objByteArr, err := json.Marshal(&obj)
	if err != nil {
		return "", fmt.Errorf("error while marshaling object: [%v]", err)
	}

	return string(objByteArr), nil
}

func TestAddToErrorSet(t *testing.T) {
	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	// get client
	vcdClient, err := getTestVCDClient(vcdConfig, nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	rdeManager := RDEManager{
		Client:                 vcdClient,
		StatusComponentName:    "capvcd",
		StatusComponentVersion: "1.0.0",
		ClusterID:              "urn:vcloud:entity:vmware:capvcdCluster:0637ac72-bf64-404b-b612-d93f019e9bf2",
	}
	ctx := context.Background()
	err1 := BackendError{
		Name:              "LoadbalancerError",
		OccurredAt:        time.Now(),
		VcdResourceId:     "3",
		AdditionalDetails: nil,
	}
	fmt.Println(fmt.Sprintf(" v [%#v]", err1))
	fmt.Println(fmt.Sprintf(" v [%s]", err1))
	rdeManager.AddToErrors(ctx, "capvcd", err1, 2)
}

func TestAddToVCDResourceSet(t *testing.T) {
	type TestCase struct {
		StatusMap      map[string]interface{}
		VCDResource    VCDResource
		ExpectedStatus map[string]interface{}
		Message        string
	}
	testCaseList := []TestCase{
		{
			Message: "add resource when absent from VCD resource set",
			StatusMap: map[string]interface{}{
				ComponentCPI: map[string]interface{}{
					"name":    "ccm",
					"version": "1.1.1",
					"vcdResourceSet": []interface{}{
						map[string]interface{}{
							"type": VcdResourceVirtualService,
							"name": "virtual-service-1",
							"id":   "12345",
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
			VCDResource: VCDResource{
				Type: VcdResourceVirtualService,
				Name: "virtual-service-2",
				ID:   "6789",
				AdditionalDetails: map[string]interface{}{
					"virtualIP": "1.2.3.4",
				},
			},
			ExpectedStatus: map[string]interface{}{
				ComponentCPI: map[string]interface{}{
					"name":    "ccm",
					"version": "1.1.1",
					"vcdResourceSet": []interface{}{
						map[string]interface{}{
							"id":   "12345",
							"name": "virtual-service-1",
							"type": VcdResourceVirtualService,
						},
						map[string]interface{}{
							"id":   "6789",
							"name": "virtual-service-2",
							"type": VcdResourceVirtualService,
							"additionalDetails": map[string]interface{}{
								"virtualIP": "1.2.3.4",
							},
						},
					},
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
				"csi": map[string]interface{}{
					"key1": "value1",
				},
			},
		},
		{
			Message: "do not add duplicate resource when it is already present in VCD resource set",
			StatusMap: map[string]interface{}{
				ComponentCPI: map[string]interface{}{
					"name":    "ccm",
					"version": "1.1.1",
					"vcdResourceSet": []interface{}{
						map[string]interface{}{
							"type": VcdResourceVirtualService,
							"name": "virtual-service-1",
							"id":   "12345",
							"additionalDetails": map[string]interface{}{
								"virtualIP": "1.2.3.4",
							},
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
			VCDResource: VCDResource{
				Type: VcdResourceVirtualService,
				Name: "virtual-service-1",
				ID:   "12345",
				AdditionalDetails: map[string]interface{}{
					"virtualIP": "1.2.3.4",
				},
			},
			ExpectedStatus: map[string]interface{}{
				ComponentCPI: map[string]interface{}{
					"name":    "ccm",
					"version": "1.1.1",
					"vcdResourceSet": []interface{}{
						map[string]interface{}{
							"id":   "12345",
							"name": "virtual-service-1",
							"type": VcdResourceVirtualService,
							"additionalDetails": map[string]interface{}{
								"virtualIP": "1.2.3.4",
							},
						},
					},
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
				"csi": map[string]interface{}{
					"key1": "value1",
				},
			},
		},
		{
			Message: "recreate CPI status if absent",
			StatusMap: map[string]interface{}{
				"csi": map[string]interface{}{
					"key1": "value1",
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
			},
			VCDResource: VCDResource{
				Type: VcdResourceVirtualService,
				Name: "virtual-service-1",
				ID:   "12345",
				AdditionalDetails: map[string]interface{}{
					"virtualIP": "1.2.3.4",
				},
			},
			ExpectedStatus: map[string]interface{}{
				ComponentCPI: map[string]interface{}{
					"name":    "ccm",
					"version": "1.1.1",
					"vcdResourceSet": []interface{}{
						map[string]interface{}{
							"id":   "12345",
							"name": "virtual-service-1",
							"type": VcdResourceVirtualService,
							"additionalDetails": map[string]interface{}{
								"virtualIP": "1.2.3.4",
							},
						},
					},
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
		updatedStatusMap, err := addToVCDResourceSet(ComponentCPI, "ccm", "1.1.1", tc.StatusMap, tc.VCDResource)
		assert.NoError(t, err, "Expected no error ", tc.Message)

		actualJson, err := convertToJson(updatedStatusMap)
		assert.NoError(t, err, tc.Message, "expected no error converting updatedStatusMap to json")
		expectedJson, err := convertToJson(tc.ExpectedStatus)
		assert.NoError(t, err, tc.Message, "expected no error when converting expected status to json")

		assert.JSONEqf(t, expectedJson, actualJson, tc.Message)
	}
}

func TestRemoveFromCPIVCDResourceSet(t *testing.T) {
	type TestCase struct {
		StatusMap       map[string]interface{}
		VCDResourceName string
		VCDResourceType string
		ExpectedStatus  map[string]interface{}
		Message         string
	}
	nameToBeRemoved := "virtual-service-1"
	testCaseList := []TestCase{
		{
			Message: "remove from resource set if name and resource type matches",
			StatusMap: map[string]interface{}{
				ComponentCPI: map[string]interface{}{
					"name":    "ccm",
					"version": "1.1.1",
					"vcdResourceSet": []VCDResource{
						{
							Type: VcdResourceVirtualService,
							Name: nameToBeRemoved,
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
			VCDResourceName: nameToBeRemoved,
			VCDResourceType: VcdResourceVirtualService,
			ExpectedStatus: map[string]interface{}{
				ComponentCPI: map[string]interface{}{
					"name":           "ccm",
					"version":        "1.1.1",
					"vcdResourceSet": []interface{}{},
				},
				"csi": map[string]interface{}{
					"key1": "value1",
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
			},
		},
		{
			Message: "no error when resource by the name and type is absent",
			StatusMap: map[string]interface{}{
				ComponentCPI: map[string]interface{}{
					"name":           "ccm",
					"version":        "1.1.1",
					"vcdResourceSet": []interface{}{},
				},
				"csi": map[string]interface{}{
					"key1": "value1",
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
			},
			VCDResourceName: nameToBeRemoved,
			VCDResourceType: VcdResourceVirtualService,
			ExpectedStatus: map[string]interface{}{
				ComponentCPI: map[string]interface{}{
					"name":           "ccm",
					"version":        "1.1.1",
					"vcdResourceSet": []interface{}{},
				},
				"csi": map[string]interface{}{
					"key1": "value1",
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
			},
		},
		{
			Message: "recreate CPI status if absent",
			StatusMap: map[string]interface{}{
				"csi": map[string]interface{}{
					"key1": "value1",
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
			},
			VCDResourceName: nameToBeRemoved,
			VCDResourceType: VcdResourceVirtualService,
			ExpectedStatus: map[string]interface{}{
				ComponentCPI: map[string]interface{}{
					"name":           "ccm",
					"version":        "1.1.1",
					"vcdResourceSet": []interface{}{},
				},
				"csi": map[string]interface{}{
					"key1": "value1",
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
			},
		},
	}
	for _, tc := range testCaseList {
		updatedStatusMap, err := removeFromVCDResourceSet(ComponentCPI, "ccm", "1.1.1",
			tc.StatusMap,
			VCDResource{
				Type: tc.VCDResourceType,
				Name: tc.VCDResourceName,
			})
		actualJson, err := convertToJson(updatedStatusMap)
		assert.NoError(t, err, tc.Message, "expected no error converting updatedStatusMap to json")
		expectedJson, err := convertToJson(tc.ExpectedStatus)
		fmt.Println(expectedJson)
		assert.NoError(t, err, tc.Message, "expected no error when converting expected status to json")

		assert.JSONEqf(t, expectedJson, actualJson, tc.Message)
	}
}
