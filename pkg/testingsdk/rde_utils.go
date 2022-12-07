package testingsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
)

func getComponentMapInStatus(ctx context.Context, client *vcdsdk.Client, clusterId, componentName string) (map[string]interface{}, error) {
	org, err := client.VCDClient.GetOrgByName(client.ClusterOrgName)
	if err != nil {
		return nil, fmt.Errorf("unable to find org [%s] by name: [%v]", client.ClusterOrgName, err)
	}

	rde, _, _, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, clusterId, org.Org.ID)
	if err != nil {
		return nil, fmt.Errorf("unable to get defined entity [%s]: [%v]", clusterId, err)
	}

	entityStatusIf, ok := rde.Entity["status"]
	if !ok {
		return nil, fmt.Errorf("status field not found in RDE [%s]", rde.Id)
	}

	entityStatus, ok := entityStatusIf.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to convert capvcd RDE [%s] status field from [%T] to map[string]interface{}",
			rde.Id, entityStatusIf)
	}

	componentStatusIf, ok := entityStatus[componentName]
	if !ok {
		return nil, fmt.Errorf("[%s] field not found in status field of RDE [%s]", componentName, rde.Id)
	}

	componentStatus, ok := componentStatusIf.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to convert [%s] status from [%T] to map[string]interface{} for RDE [%s]",
			componentName, componentStatusIf, clusterId)
	}
	return componentStatus, nil
}

func getKubeconfigFromRDEId(client *vcdsdk.Client, clusterId string) (string, error) {
	capvcdStatusMap, err := getComponentMapInStatus(context.TODO(), client, clusterId, vcdsdk.ComponentCAPVCD)
	if err != nil {
		return "", fmt.Errorf("error retrieving [%s] field in status field of RDE [%s]: [%v]", vcdsdk.ComponentCAPVCD, clusterId, err)
	}

	privateStatusIf, ok := capvcdStatusMap["private"]
	if !ok {
		return "", fmt.Errorf("private field not found in status->capvcd of RDE [%s]", clusterId)
	}

	capvcdPrivateStatusMap, ok := privateStatusIf.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to convert RDE [%s]  status->capvcd->private from [%T] to map[string]interface{}", clusterId, privateStatusIf)
	}

	kubeConfig, ok := capvcdPrivateStatusMap["kubeConfig"]
	if !ok {
		return "", fmt.Errorf("kubeConfig field not found in status->capvcd->private of RDE [%s]", clusterId)
	}
	return kubeConfig.(string), nil
}

func getVcdResourceSetComponentMapFromRDEId(ctx context.Context, client *vcdsdk.Client, clusterId, componentName string) (interface{}, error) {
	componentStatusMap, err := getComponentMapInStatus(ctx, client, componentName, clusterId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving field [%s] in status from RDE [%s]: [%v]", componentName, clusterId, err)
	}

	// Inferred type is interface{} for both key and values. So we cannot directly type assert into []vcdsdk.VCDResource.
	vcdResourceSetArr, ok := componentStatusMap[vcdsdk.ComponentStatusFieldVCDResourceSet]
	if !ok {
		return nil, fmt.Errorf("vcdResourceSet field not found in status->[%s] of RDE [%s]", componentName, clusterId)
	}
	return vcdResourceSetArr, nil
}

// This step is required because the type inferred is []interface{}, and we cannot do type assertion here, so we must marshal and unmarshal
func convertVcdResourceSetMapToVcdResourceArr(vcdResourceSetArr interface{}) ([]vcdsdk.VCDResource, error) {
	data, err := json.Marshal(vcdResourceSetArr)
	if err != nil {
		return nil, fmt.Errorf("error marshaling VCDResourceSet [%T]: [%v]", vcdResourceSetArr, err)
	}
	var vcdResourceSet []vcdsdk.VCDResource
	err = json.Unmarshal(data, &vcdResourceSet)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling data to [%T]: [%v]", vcdResourceSet, err)
	}
	return vcdResourceSet, nil
}
