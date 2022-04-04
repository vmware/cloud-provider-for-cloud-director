package vcdclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/util"
	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"k8s.io/klog"
	"net/http"
	"strings"
)

const (
	NoRdePrefix = `NO_RDE_`

	CloudControllerManagerName    = "cloud-controller-manager"
	CloudControllerManagerVersion = "1.1.2"

	// VcdResourceVirtualService defines the type Virtual Service in the VCDResourceSet
	VcdResourceVirtualService = "virtual-service"
)

func (client *Client) GetRDEVirtualIps(ctx context.Context) ([]string, string, *swaggerClient.DefinedEntity, error) {
	if client.ClusterID == "" || strings.HasPrefix(client.ClusterID, NoRdePrefix) {
		klog.Infof("ClusterID [%s] is empty or generated", client.ClusterID)
		return nil, "", nil, nil
	}

	defEnt, _, etag, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, client.ClusterID)
	if err != nil {
		return nil, "", nil, fmt.Errorf("error when getting defined entity: [%v]", err)
	}

	virtualIpStrs, err := util.GetVirtualIPsFromRDE(&defEnt)
	if err != nil {
		return nil, "", nil, fmt.Errorf("failed to retrieve Virtual IPs from RDE [%s]: [%v]",
			client.ClusterID, err)
	}
	return virtualIpStrs, etag, &defEnt, nil
}

// This function will modify the passed in defEnt
func (client *Client) updateRDEVirtualIps(ctx context.Context, updatedIps []string, etag string,
	defEnt *swaggerClient.DefinedEntity) (*http.Response, error) {
	defEnt, err := util.ReplaceVirtualIPsInRDE(defEnt, updatedIps)
	if err != nil {
		return nil, fmt.Errorf("failed to locally edit RDE with ID [%s] with virtual IPs: [%v]", client.ClusterID, err)
	}
	// can pass invokeHooks
	_, httpResponse, err := client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, *defEnt, etag, client.ClusterID, nil)
	if err != nil {
		return httpResponse, fmt.Errorf("error when updating defined entity [%s]: [%v]", client.ClusterID, err)
	}
	return httpResponse, nil
}

func (client *Client) addVirtualIpToRDE(ctx context.Context, addIp string) error {
	if addIp == "" {
		klog.Infof("VIP is empty, hence not adding anything to RDE")
		return nil
	}
	if client.ClusterID == "" || strings.HasPrefix(client.ClusterID, NoRdePrefix) {
		klog.Infof("ClusterID [%s] is empty or generated, hence not adding VIP [%s] from RDE",
			client.ClusterID, addIp)
		return nil
	}

	numRetries := 10
	for i := 0; i < numRetries; i++ {
		currIps, etag, defEnt, err := client.GetRDEVirtualIps(ctx)
		if err != nil {
			return fmt.Errorf("error getting current vips: [%v]", err)
		}

		// check if need to update RDE
		foundAddIp := false
		for _, ip := range currIps {
			if ip == addIp {
				foundAddIp = true
				break
			}
		}
		if foundAddIp {
			return nil // no need to update RDE
		}

		updatedIps := append(currIps, addIp)
		httpResponse, err := client.updateRDEVirtualIps(ctx, updatedIps, etag, defEnt)
		if err != nil {
			if httpResponse.StatusCode == http.StatusPreconditionFailed {
				klog.Infof("Wrong ETag while adding virtual IP [%s]", addIp)
				continue
			}
			return fmt.Errorf("error when adding virtual ip [%s] to RDE: [%v]", addIp, err)
		}
		klog.Infof("Successfully updated RDE [%s] with virtual IP [%s]", client.ClusterID, addIp)
		return nil
	}

	return fmt.Errorf("unable to update rde due to incorrect etag after [%d]] tries", numRetries)
}

func (client *Client) removeVirtualIpFromRDE(ctx context.Context, removeIp string) error {
	if removeIp == "" {
		klog.Infof("VIP is empty, hence not removing anything from RDE")
		return nil
	}
	if client.ClusterID == "" || strings.HasPrefix(client.ClusterID, NoRdePrefix) {
		klog.Infof("ClusterID [%s] is empty or generated, hence not removing VIP [%s] from RDE",
			client.ClusterID, removeIp)
		return nil
	}

	numRetries := 10
	for i := 0; i < numRetries; i++ {
		currIps, etag, defEnt, err := client.GetRDEVirtualIps(ctx)
		if err != nil {
			return fmt.Errorf("error getting current vips: [%v]", err)
		}
		// currIps is guaranteed not to be nil by GetRDEVirtualIps
		if len(currIps) == 0 {
			// valid case since this could be a retry operation
			return nil
		}

		// form updated virtual ip list
		foundIdx := -1
		for idx, ip := range currIps {
			if ip == removeIp {
				foundIdx = idx
				break // for inner loop
			}
		}
		if foundIdx == -1 {
			return nil // no need to update RDE
		}
		updatedIps := append(currIps[:foundIdx], currIps[foundIdx+1:]...)

		httpResponse, err := client.updateRDEVirtualIps(ctx, updatedIps, etag, defEnt)
		if err != nil {
			if httpResponse.StatusCode == http.StatusPreconditionFailed {
				klog.Infof("Wrong ETag while removing virtual IP [%s]", removeIp)
				continue
			}
			return fmt.Errorf("error when removing virtual ip [%s] from RDE: [%v]",
				removeIp, err)
		}
		return nil
	}

	return fmt.Errorf("unable to update rde due to incorrect etag after [%d] tries", numRetries)
}

func convertCPIStatusToMap(cpiStatus util.CPIStatus) (map[string]interface{}, error) {
	cpiStatusBytes, err := json.Marshal(&cpiStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CPI status to byte array: [%v]", err)
	}
	var cpiStatusMap map[string]interface{}
	err = json.Unmarshal(cpiStatusBytes, &cpiStatusMap)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CPI status bytes to map[string]interface{}: [%v]", err)
	}
	return cpiStatusMap, nil
}

func convertMapToCPIStatus(cpiStatusMap map[string]interface{}) (*util.CPIStatus, error) {
	cpiStatusMapBytes, err := json.Marshal(&cpiStatusMap)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CPI status map to byte array: [%v]", err)
	}
	var cpiStatus util.CPIStatus
	err = json.Unmarshal(cpiStatusMapBytes, &cpiStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CPI status bytes to CPIStatus: [%v]", err)
	}
	return &cpiStatus, nil
}

func (client *Client) CreateOrUpdateCPIStatusInRDE(ctx context.Context, rdeId string) error {
	rde, resp, etag, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, rdeId)
	if resp != nil && resp.StatusCode != http.StatusOK {
		var responseMessageBytes []byte
		if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
			responseMessageBytes = gsErr.Body()
		}
		return fmt.Errorf(
			"failed to get RDE with id [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
			rdeId, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
	} else if err != nil {
		return fmt.Errorf("error while getting the RDE [%s]: [%v]", rdeId, err)
	}
	if !util.IsCAPVCDEntityType(rde.EntityType) {
		return fmt.Errorf("unable to update RDE as the RDE entity type is not [%v]", rde.EntityType)
	}
	statusEntity, ok := rde.Entity["status"]
	if !ok {
		return fmt.Errorf("failed to fetch RDE status for RDE [%s]", rdeId)
	}
	statusMap, ok := statusEntity.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to convert status of RDE [%s] to map[string]interface{}", rdeId)
	}
	cpiStatus := &util.CPIStatus{
		VCDResourceSet: nil,
		Errors:         nil,
		VirtualIPs:     nil,
	}
	cpiStatusEntity, ok := statusMap["cpi"]
	if ok {
		cpiStatusMap, ok := cpiStatusEntity.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to convert CPI status entity to map[string]interface{}")
		}
		cpiStatus, err = convertMapToCPIStatus(cpiStatusMap)
		if err != nil {
			return fmt.Errorf("failed to convert CPI status map to CPI status object: [%v]", err)
		}
	}
	cpiStatus.Name = CloudControllerManagerName
	cpiStatus.Version = CloudControllerManagerVersion
	// CPI section is missing from the RDE status. Create a new CPI status and update the RDE.
	cpiStatusMap, err := convertCPIStatusToMap(*cpiStatus)
	if err != nil {
		return fmt.Errorf("failed to convert CPI status [%v] to map[string]interface{}: [%v]", cpiStatus, err)
	}
	statusMap["cpi"] = cpiStatusMap

	_, resp, err = client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, rde, etag, rdeId, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		var responseMessageBytes []byte
		if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
			responseMessageBytes = gsErr.Body()
		}
		return fmt.Errorf(
			"failed to create CPI status for RDE [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
			rdeId, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
	} else if err != nil {
		return fmt.Errorf("error while getting the RDE [%s]: [%v]", rdeId, err)
	}
	return nil
}

func (client *Client) AddVirtualServiceToVCDResourceSet(ctx context.Context, vsName, vsID string) error {
	rde, resp, etag, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, client.ClusterID)
	if resp != nil && resp.StatusCode != http.StatusOK {
		var responseMessageBytes []byte
		if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
			responseMessageBytes = gsErr.Body()
		}
		return fmt.Errorf(
			"failed to get RDE with id [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
			client.ClusterID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
	} else if err != nil {
		return fmt.Errorf("error while updating the RDE [%s]: [%v]", client.ClusterID, err)
	}

	if !util.IsCAPVCDEntityType(rde.EntityType) {
		return fmt.Errorf("unable to update RDE as the RDE entity type is not [%v]", rde.EntityType)
	}

	statusEntity, ok := rde.Entity["status"]
	if !ok {
		return fmt.Errorf("failed to fetch RDE status for RDE [%s]", client.ClusterID)
	}
	statusMap, ok := statusEntity.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to convert status of RDE [%s] to map[string]interface{}", client.ClusterID)
	}
	cpiStatusEntity, ok := statusMap["cpi"]
	var cpiStatus *util.CPIStatus
	if !ok {
		// need to create cpi section if absent
		cpiStatus = &util.CPIStatus{
			Name:    CloudControllerManagerName,
			Version: CloudControllerManagerVersion,
			VCDResourceSet: []util.VCDResource{
				{
					Type: VcdResourceVirtualService,
					ID:   vsID,
					Name: vsName,
				},
			},
		}
	} else {
		cpiStatusMap, ok := cpiStatusEntity.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to parse CPI status entry in the RDE status")
		}
		cpiStatus, err = convertMapToCPIStatus(cpiStatusMap)
		if err != nil {
			return fmt.Errorf("failed to convert map[string]interface{} to CPIStatus object: [%v]", err)
		}
		for _, vcdResource := range cpiStatus.VCDResourceSet {
			if vcdResource.Type == VcdResourceVirtualService && vcdResource.ID == vsID {
				// VCDResource already present in the VCDResourceSet
				return nil
			}
		}
		vsResource := util.VCDResource{
			Type: VcdResourceVirtualService,
			Name: vsName,
			ID:   vsID,
		}
		// virtual service not present in the RDE
		cpiStatus.VCDResourceSet = append(cpiStatus.VCDResourceSet, vsResource)
	}
	cpiStatusMap, err := convertCPIStatusToMap(*cpiStatus)
	if err != nil {
		return fmt.Errorf("failed to update CPI status with virtual service [%s]: [%v]", vsName, err)
	}
	statusMap["cpi"] = cpiStatusMap

	_, resp, err = client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, rde, etag, client.ClusterID, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		var responseMessageBytes []byte
		if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
			responseMessageBytes = gsErr.Body()
		}
		return fmt.Errorf(
			"failed to update CPI status for RDE [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
			client.ClusterID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
	} else if err != nil {
		return fmt.Errorf("error while updating the RDE [%s]: [%v]", client.ClusterID, err)
	}
	return nil
}

func (client *Client) RemoveVirtualServiceFromVCDResourceSet(ctx context.Context, vsID string) error {
	rde, resp, etag, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, client.ClusterID)
	if resp != nil && resp.StatusCode != http.StatusOK {
		var responseMessageBytes []byte
		if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
			responseMessageBytes = gsErr.Body()
		}
		return fmt.Errorf(
			"failed to get RDE with id [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
			client.ClusterID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
	} else if err != nil {
		return fmt.Errorf("error while updating the RDE [%s]: [%v]", client.ClusterID, err)
	}

	if !util.IsCAPVCDEntityType(rde.EntityType) {
		return fmt.Errorf("unable to update RDE as the RDE entity type is not [%v]", rde.EntityType)
	}

	statusEntity, ok := rde.Entity["status"]
	if !ok {
		return fmt.Errorf("failed to fetch RDE status for RDE [%s]", client.ClusterID)
	}
	statusMap, ok := statusEntity.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to convert status of RDE [%s] to map[string]interface{}", client.ClusterID)
	}
	var cpiStatus *util.CPIStatus
	cpiStatusEntity, ok := statusMap["cpi"]
	if !ok {
		// need to create cpi section if absent
		cpiStatus = &util.CPIStatus{
			Name:           CloudControllerManagerName,
			Version:        CloudControllerManagerVersion,
			VCDResourceSet: nil,
		}
	} else {
		cpiStatusMap, ok := cpiStatusEntity.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to parse CPI status entry in the RDE status")
		}
		cpiStatus, err = convertMapToCPIStatus(cpiStatusMap)
		if err != nil {
			return fmt.Errorf("failed to convert map[string]interface{} to CPIStatus object: [%v]", err)
		}
		updatedVCDResourceSet := make([]util.VCDResource, 0)
		i := 0
		for _, vcdResource := range cpiStatus.VCDResourceSet {
			if vcdResource.Type == VcdResourceVirtualService && vcdResource.ID == vsID {
				// remove vcdResource from vcdResourceSet
				continue
			}
			updatedVCDResourceSet[i] = vcdResource
			i++
		}
		cpiStatus.VCDResourceSet = updatedVCDResourceSet
	}
	cpiStatusMap, err := convertCPIStatusToMap(*cpiStatus)
	if err != nil {
		return fmt.Errorf("failed to convert CPIStatus object to map[string]interface{} to remove virtual service [%s]: [%v]", vsID, err)
	}
	statusMap["cpi"] = cpiStatusMap

	_, resp, err = client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, rde, etag, client.ClusterID, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		var responseMessageBytes []byte
		if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
			responseMessageBytes = gsErr.Body()
		}
		return fmt.Errorf(
			"failed to update CPI status for RDE [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
			client.ClusterID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
	} else if err != nil {
		return fmt.Errorf("error while updating the RDE [%s]: [%v]", client.ClusterID, err)
	}
	return nil
}
