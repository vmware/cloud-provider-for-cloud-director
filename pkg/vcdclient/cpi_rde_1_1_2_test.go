package vcdclient

import (
	"github.com/stretchr/testify/assert"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/util"
	"testing"
)

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
					"vcdResourceSet": []util.VCDResource{
						{
							Type: VcdResourceVirtualService,
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
					"version": "1.1.2",
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
	}
}

func TestAddToVCDResourceSet(t *testing.T) {
	type TestCase struct {
		StatusMap      map[string]interface{}
		VCDResource    util.VCDResource
		ExpectedStatus map[string]interface{}
		Message string
	}
	testCaseList := []TestCase{
		{
			Message: "add resource when absent from VCD resource set",
			StatusMap: map[string]interface{}{
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
					"vcdResourceSet": []util.VCDResource{
						{
							Type: VcdResourceVirtualService,
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
			VCDResource: util.VCDResource{
				Type: VcdResourceVirtualService,
				Name: "virtual-service-2",
				ID: "6789",
				AdditionalDetails: map[string]interface{}{
					"virtualIP": "1.2.3.4",
				},
			},
			ExpectedStatus: map[string]interface{}{
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
					"vcdResourceSet": []interface{}{
						map[string]interface{}{
							"id":   "12345",
							"name": "virtual-service-1",
							"type": VcdResourceVirtualService,
						},
						map[string]interface{}{
							"id": "6789",
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
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
					"vcdResourceSet": []util.VCDResource{
						{
							Type: VcdResourceVirtualService,
							Name: "virtual-service-1",
							ID:   "12345",
							AdditionalDetails: map[string]interface{}{
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
			VCDResource: util.VCDResource{
				Type: VcdResourceVirtualService,
				Name: "virtual-service-1",
				ID: "12345",
				AdditionalDetails: map[string]interface{}{
					"virtualIP": "1.2.3.4",
				},
			},
			ExpectedStatus: map[string]interface{}{
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
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
			VCDResource: util.VCDResource{
				Type: VcdResourceVirtualService,
				Name: "virtual-service-1",
				ID: "12345",
				AdditionalDetails: map[string]interface{}{
					"virtualIP": "1.2.3.4",
				},
			},
			ExpectedStatus: map[string]interface{}{
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
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
		updatedStatusMap, err := AddToVCDResourceSet(tc.StatusMap, tc.VCDResource)
		assert.NoError(t, err, "Expected no error ", tc.Message)
		assert.Equal(t, updatedStatusMap, tc.ExpectedStatus, "expected updated status map to match the expected status map ", tc.Message)
	}
}

func TestRemoveFromCPIVCDResourceSet(t *testing.T) {
	type TestCase struct {
		StatusMap map[string]interface{}
		VCDResourceID string
		ExpectedStatus map[string]interface{}
		Message string
	}
	idToBeRemoved := "12345"
	testCaseList := []TestCase{
		{
			Message: "remove from resource set if ID matches",
			StatusMap: map[string]interface{}{
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
					"vcdResourceSet": []util.VCDResource{
						{
							Type: VcdResourceVirtualService,
							Name: "virtual-service-1",
							ID:   idToBeRemoved,
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
			VCDResourceID: idToBeRemoved,
			ExpectedStatus: map[string]interface{}{
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
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
			Message: "no error when ID is absent",
			StatusMap: map[string]interface{}{
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
					"vcdResourceSet": nil,
				},
				"csi": map[string]interface{}{
					"key1": "value1",
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
			},
			VCDResourceID: idToBeRemoved,
			ExpectedStatus: map[string]interface{}{
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
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
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
					"vcdResourceSet": nil,
				},
				"csi": map[string]interface{}{
					"key1": "value1",
				},
				"capvcd": map[string]interface{}{
					"key2": "value2",
				},
			},
			VCDResourceID: idToBeRemoved,
			ExpectedStatus: map[string]interface{}{
				"cpi": map[string]interface{}{
					"name": CloudControllerManagerName,
					"version": CloudControllerManagerVersion,
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
		updatedStatusMap, err := RemoveFromCPIVCDResourceSet(tc.StatusMap, tc.VCDResourceID)
		assert.NoError(t, err, "error executing test: ", tc.Message)
		assert.Equal(t, updatedStatusMap, tc.ExpectedStatus, "updated status map doesn't match expected status map: ", tc.Message)
	}
}
