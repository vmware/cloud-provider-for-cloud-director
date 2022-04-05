package vcdclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/util"
	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"net/http"
)

const (
	CloudControllerManagerName    = "cloud-controller-manager"
	CloudControllerManagerVersion = "1.1.2"

	// VcdResourceVirtualService defines the type Virtual Service in the VCDResourceSet
	VcdResourceVirtualService = "virtual-service"
)

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

// UpgradeCPIStatusOfExistingRDE updates the CPI section of the existing RDEs to the latest model (CPI 1.1.2).
// As this section is empty in all the older existing RDEs, will always create a new section.
// But the future versions of CPI need to compare CPIStatus.Version property in the existing RDEs to determine whether an RDE upgrade is needed.
func (client *Client) UpgradeCPIStatusOfExistingRDE(ctx context.Context, rdeId string) error {
	rde, resp, etag, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, rdeId)
	if resp != nil && resp.StatusCode != http.StatusOK {
		var responseMessageBytes []byte
		if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
			responseMessageBytes = gsErr.Body()
		}
		return fmt.Errorf(
			"failed to get RDE with id [%s] when upgrading CPI status; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
			rdeId, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
	} else if err != nil {
		return fmt.Errorf("error while getting the RDE [%s]: [%v]", rdeId, err)
	}

	if !util.IsCAPVCDEntityType(rde.EntityType) {
		return fmt.Errorf("unable to upgrade CPI status in RDE [%s] as the RDE entity type is not [%v]",rdeId, rde.EntityType)
	}
	statusEntity, ok := rde.Entity["status"]
	if !ok {
		return fmt.Errorf("failed to fetch status section in RDE [%s] when upgrading CPI section", rdeId)
	}
	statusMap, ok := statusEntity.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to convert status section of RDE [%s] to map[string]interface{} during RDE upgrade", rdeId)
	}

	var virtualIPs []string
	virtualIPsIf, ok := statusMap["virtualIPs"]
	if !ok || virtualIPsIf == nil {
		virtualIPs = nil
	} else {
		virtualIPs = virtualIPsIf.([]string)
	}
	cpiStatus := &util.CPIStatus{
		VCDResourceSet: nil,
		Errors:         nil,
		VirtualIPs:     virtualIPs,
	}
	cpiStatusEntity, ok := statusMap["cpi"]
	if ok {
		cpiStatusMap, ok := cpiStatusEntity.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to convert CPI status in RDE [%s] from interface{} to map[string]interface{} when upgrading CPI status section", rdeId)
		}
		cpiStatus, err = convertMapToCPIStatus(cpiStatusMap)
		if err != nil {
			return fmt.Errorf("failed to convert CPI status map in RDE [%s] to CPI status object: [%v]", rdeId, err)
		}
	}
	cpiStatus.Name = CloudControllerManagerName
	cpiStatus.Version = CloudControllerManagerVersion
	// CPI section is missing from the RDE status. Create a new CPI status and update the RDE.
	cpiStatusMap, err := convertCPIStatusToMap(*cpiStatus)
	if err != nil {
		return fmt.Errorf("failed to convert CPI status [%v] in RDE [%s] to map[string]interface{} when upgrading CPI status section: [%v]", cpiStatus, rdeId, err)
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
			"failed to get RDE [%s] when adding virtual service to VCD resource set ; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
			client.ClusterID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
	} else if err != nil {
		return fmt.Errorf("error while updating the RDE [%s]: [%v]", client.ClusterID, err)
	}

	if !util.IsCAPVCDEntityType(rde.EntityType) {
		return fmt.Errorf("unable to add virtual service to RDE [%s] as the RDE entity type is not of type [%v]", rde.EntityType)
	}

	statusEntity, ok := rde.Entity["status"]
	if !ok {
		return fmt.Errorf("failed to update RDE [%s] with virtual status information", client.ClusterID)
	}
	statusMap, ok := statusEntity.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to convert RDE status [%s] to map[string]interface{}", client.ClusterID)
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
			return fmt.Errorf("failed to parse CPI status entry in the RDE [%s] to update virtual service infromation", client.ClusterID)
		}
		cpiStatus, err = convertMapToCPIStatus(cpiStatusMap)
		if err != nil {
			return fmt.Errorf("failed to convert map[string]interface{} to CPIStatus object in RDE [%s] to update virtual service information: [%v]", client.ClusterID, err)
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
		return fmt.Errorf("failed to add virtual service [%s] information in CPI status of RDE [%s]: [%v]", vsName, client.ClusterID, err)
	}
	statusMap["cpi"] = cpiStatusMap

	_, resp, err = client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, rde, etag, client.ClusterID, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		var responseMessageBytes []byte
		if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
			responseMessageBytes = gsErr.Body()
		}
		return fmt.Errorf(
			"failed to add virtual service [%s] information in CPI status for RDE [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
			vsName, client.ClusterID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
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
		return fmt.Errorf("unable to remove virtual service [%s] from CPI status in RDE [%s] as the RDE entity type is not [%v]", vsID, client.ClusterID, rde.EntityType)
	}

	statusEntity, ok := rde.Entity["status"]
	if !ok {
		return fmt.Errorf("failed to parse status in RDE [%s] to remove virtual service [%s] from CPI status", client.ClusterID, vsID)
	}
	statusMap, ok := statusEntity.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to parse status in RDE [%s] into map[string]interface{} to remove virtual service [%s] from CPI status", client.ClusterID, vsID)
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
			return fmt.Errorf("failed to parse CPI status entry in status of RDE [%s] to remove virtual service [%s]", client.ClusterID, vsID)
		}
		cpiStatus, err = convertMapToCPIStatus(cpiStatusMap)
		if err != nil {
			return fmt.Errorf("failed to convert map[string]interface{} to CPIStatus object while removing virtual service [%s] from RDE [%s]: [%v]", vsID, client.ClusterID, err)
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
		return fmt.Errorf("failed to convert CPIStatus object to map[string]interface{} to remove virtual service [%s] from RDE [%s]: [%v]", vsID, client.ClusterID, err)
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
		return fmt.Errorf("error while removing virtual service [%s] from the RDE [%s]: [%v]", vsID, client.ClusterID, err)
	}
	return nil
}
