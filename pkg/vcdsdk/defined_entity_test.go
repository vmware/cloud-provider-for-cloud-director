package vcdsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"net/http"
	"testing"
	"time"

	swagger "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
)

const (
	CAPVCDTypeVendor              = "vmware"
	CAPVCDTypeNss                 = "capvcdCluster"
	CAPVCDTypeVersion             = "1.1.0"
	VCDLocationHeader             = "Location"
	CAPVCDComponentRDESectionName = "capvcd"
)

var (
	CAPVCDEntityTypeID = fmt.Sprintf("urn:vcloud:type:%s:%s:%s", CAPVCDTypeVendor, CAPVCDTypeNss, CAPVCDTypeVersion)
)

func convertToJson(obj interface{}) (string, error) {
	objByteArr, err := json.Marshal(&obj)
	if err != nil {
		return "", fmt.Errorf("error while marshaling object: [%v]", err)
	}

	return string(objByteArr), nil
}

func TestCRUDOnEventSet(t *testing.T) {
	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	// get client
	vcdClient, err := getTestVCDClient(vcdConfig, nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")
	ctx := context.Background()

	// create a minimal CAPVCD RDE with almost empty spec and status
	rdeId, err := createCapvcdRDE(ctx, vcdClient, "testCluster")

	// Create some mock objects for events in capvcd section
	rdeManager := RDEManager{
		Client:                 vcdClient,
		StatusComponentName:    CAPVCDComponentRDESectionName,
		StatusComponentVersion: "1.0.0",
		ClusterID:              rdeId,
	}
	ControlPlaneInitializedEvent := BackendEvent{
		Name:              "ControlplaneInitialized",
		OccurredAt:        time.Now(),
		VcdResourceId:     "vmId",
		AdditionalDetails: nil,
	}
	InfrastractureCreatedEvent := BackendEvent{
		Name:              "InfrastructureCreated",
		OccurredAt:        time.Now(),
		VcdResourceId:     "virtualServiceId",
		AdditionalDetails: nil,
	}
	CloudInitEvent := BackendEvent{
		Name:              "CloudInitSuccessful",
		OccurredAt:        time.Now(),
		VcdResourceId:     "VmId",
		AdditionalDetails: nil,
	}

	// add few events to rde.status.capvcd.eventSet
	err = rdeManager.AddToEventSet(ctx, "capvcd", ControlPlaneInitializedEvent, 3)
	assert.NoError(t, err, "failed to add event into the eventset")
	rdeManager.AddToEventSet(ctx, "capvcd", InfrastractureCreatedEvent, 3)
	assert.NoError(t, err, "failed to add event into the eventset")
	rdeManager.AddToEventSet(ctx, "capvcd", CloudInitEvent, 3)
	assert.NoError(t, err, "failed to add event into the eventset")

	// get the rde and check if the length of errors added is same as expected
	rde, _, _, err := vcdClient.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, rdeId)
	status, _ := rde.Entity["status"].(map[string]interface{})
	capvcdComponent, _ := status[CAPVCDComponentRDESectionName].(map[string]interface{})
	eventSet, _ := capvcdComponent["eventSet"].([]interface{})
	assert.Equal(t, 3, len(eventSet), "Length of error set must match with event additions requested")

	rdeManager.AddToEventSet(ctx, "capvcd", CloudInitEvent, 3)
	assert.NoError(t, err, "failed to add event into the eventset")

	// get the rde and check if the length of events is still capped at 3- window size (even though 4 events were added)
	rde, _, _, err = vcdClient.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, rdeId)
	status, _ = rde.Entity["status"].(map[string]interface{})
	capvcdComponent, _ = status[CAPVCDComponentRDESectionName].(map[string]interface{})
	eventSet, _ = capvcdComponent["eventSet"].([]interface{})
	assert.Equal(t, 3, len(eventSet), "Length of error set must match with rollingWindowSize")

	// delete RDE
	_, _, err = vcdClient.APIClient.DefinedEntityApi.ResolveDefinedEntity(ctx, rdeId)
	_, err = vcdClient.APIClient.DefinedEntityApi.DeleteDefinedEntity(ctx,
		rdeId, nil)
	assert.NoError(t, err, "failed to delete rdeId")
}

func TestCRUDOnErrorSet(t *testing.T) {
	vcdConfig, err := getTestVCDConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	// get client
	vcdClient, err := getTestVCDClient(vcdConfig, nil)
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")
	ctx := context.Background()

	// create a minimal CAPVCD RDE with almost empty spec and status
	rdeId, err := createCapvcdRDE(ctx, vcdClient, "testCluster")

	// Create some mock objects for errors in capvcd section
	rdeManager := RDEManager{
		Client:                 vcdClient,
		StatusComponentName:    CAPVCDComponentRDESectionName,
		StatusComponentVersion: "1.0.0",
		ClusterID:              rdeId,
	}
	CloudInitError := BackendError{
		Name:              "CloudInitError",
		OccurredAt:        time.Now(),
		VcdResourceId:     "vmId",
		AdditionalDetails: nil,
	}
	LoadBalancerError := BackendError{
		Name:              "LoadBalancerError",
		OccurredAt:        time.Now(),
		VcdResourceId:     "virtualServiceId",
		AdditionalDetails: nil,
	}
	ControlPlaneError := BackendError{
		Name:              "ControlPlaneError",
		OccurredAt:        time.Now(),
		VcdResourceId:     "VmId",
		AdditionalDetails: nil,
	}
	CloudInitError1 := BackendError{
		Name:              "CloudInitError",
		OccurredAt:        time.Now(),
		VcdResourceId:     "vm1Id",
		AdditionalDetails: nil,
	}

	// add few errors to rde.status.capvcd.errorSet
	err = rdeManager.AddToErrorSet(ctx, "capvcd", CloudInitError, 8)
	assert.NoError(t, err, "failed to add error into the errorset")
	rdeManager.AddToErrorSet(ctx, "capvcd", LoadBalancerError, 8)
	assert.NoError(t, err, "failed to add error into the errorset")
	rdeManager.AddToErrorSet(ctx, "capvcd", ControlPlaneError, 8)
	assert.NoError(t, err, "failed to add error into the errorset")
	rdeManager.AddToErrorSet(ctx, "capvcd", LoadBalancerError, 8)
	assert.NoError(t, err, "failed to add error into the errorset")
	rdeManager.AddToErrorSet(ctx, "capvcd", CloudInitError1, 8)
	assert.NoError(t, err, "failed to add error into the errorset")

	// get the rde and check if the length of errors added is same as expected
	rde, _, _, err := vcdClient.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, rdeId)
	status, _ := rde.Entity["status"].(map[string]interface{})
	capvcdComponent, _ := status[CAPVCDComponentRDESectionName].(map[string]interface{})
	errorSet, _ := capvcdComponent["errorSet"].([]interface{})
	assert.Equal(t, 5, len(errorSet), "Length of error set must match with error additions requested")

	// remove few errors from rde.status.capvcd.errorSet
	err = rdeManager.RemoveErrorByNameOrIdFromErrorSet(ctx, "capvcd", "LoadBalancerError", "")
	assert.NoError(t, err, "failed to remove error from the errorset")
	err = rdeManager.RemoveErrorByNameOrIdFromErrorSet(ctx, "capvcd", "CloudInitError", "")
	assert.NoError(t, err, "failed to remove error from the errorset")
	err = rdeManager.RemoveErrorByNameOrIdFromErrorSet(ctx, "capvcd", "ControlPlaneError", "")
	assert.NoError(t, err, "failed to remove error from the errorset")

	// get the rde and check if the length of the errorSet after removing errors is same as expected
	rde, _, _, err = vcdClient.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, rdeId)
	status, _ = rde.Entity["status"].(map[string]interface{})
	capvcdComponent, _ = status[CAPVCDComponentRDESectionName].(map[string]interface{})
	errorSet, _ = capvcdComponent["errorSet"].([]interface{})
	assert.Equal(t, 0, len(errorSet), "Length of error should be reduced to 0 after removing few errors")

	// delete RDE
	_, _, err = vcdClient.APIClient.DefinedEntityApi.ResolveDefinedEntity(ctx, rdeId)
	_, err = vcdClient.APIClient.DefinedEntityApi.DeleteDefinedEntity(ctx,
		rdeId, nil)
	assert.NoError(t, err, "failed to delete rdeId")
}

func createCapvcdRDE(ctx context.Context, vcdClient *Client, clusterName string) (string, error) {
	rde := &swagger.DefinedEntity{
		EntityType: CAPVCDEntityTypeID,
		Name:       clusterName,
	}
	entityMap := map[string]interface{}{
		"spec": "",
		"status": map[string]interface{}{
			CAPVCDComponentRDESectionName: map[string]interface{}{},
		},
	}
	rde.Entity = entityMap
	resp, err := vcdClient.APIClient.DefinedEntityApi.CreateDefinedEntity(ctx, *rde, rde.EntityType, nil)
	if err != nil {
		return "", fmt.Errorf("error occurred during RDE creation for the cluster [%s]: [%v]", clusterName, err)
	}
	if resp.StatusCode != http.StatusAccepted {
		return "", fmt.Errorf("error occurred during RDE creation for the cluster [%s]", clusterName)
	}
	taskURL := resp.Header.Get(VCDLocationHeader)
	task := govcd.NewTask(&vcdClient.VCDClient.Client)
	task.Task.HREF = taskURL
	err = task.Refresh()
	if err != nil {
		return "", fmt.Errorf("error occurred during RDE creation for the cluster [%s]; error refreshing task: [%s]", clusterName, task.Task.HREF)
	}
	rdeID := task.Task.Owner.ID
	return rdeID, nil
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
