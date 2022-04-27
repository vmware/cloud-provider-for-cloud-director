package vcdsdk

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
	CloudControllerManagerName    = "cloud-controller-manager"
	CloudControllerManagerVersion = "1.2.0"

	MaxRDEUpdateRetries = 10

	// VCD resource types in VCDResourceSet
	VcdResourceVirtualService = "virtual-service"
	VcdResourceLoadBalancerPool = "lb-pool"
	VcdResourceDNATRule = "dnat-rule"
	VcdResourceAppPortProfile = "app-port-profile"
)

type NonCAPVCDEntityError struct {
	EntityTypeID string
}

func (e NonCAPVCDEntityError) Error() string {
	return fmt.Sprintf("Non CAPVCD entity type found: [%s]", e.EntityTypeID)
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

func UpgradeCPISectionInStatus(statusMap map[string]interface{}) (map[string]interface{}, error) {
	if statusMap == nil {
		return nil, fmt.Errorf("invalid value for status: [%v]", statusMap)
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
			return nil, fmt.Errorf("failed to convert CPI status in RDE from interface{} to map[string]interface{} when upgrading CPI status section")
		}
		var err error
		cpiStatus, err = convertMapToCPIStatus(cpiStatusMap)
		if err != nil {
			return nil, fmt.Errorf("failed to convert CPI status map in RDE to CPI status object: [%v]", err)
		}
	}
	cpiStatus.Name = CloudControllerManagerName
	cpiStatus.Version = CloudControllerManagerVersion
	// CPI section is missing from the RDE status. Create a new CPI status and update the RDE.
	cpiStatusMap, err := convertCPIStatusToMap(*cpiStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CPI status [%v] in RDE to map[string]interface{} when upgrading CPI status section: [%v]", cpiStatus, err)
	}
	statusMap["cpi"] = cpiStatusMap
	return statusMap, nil
}

// UpgradeCPIStatusOfExistingRDE Creates an empty cpi section with just the name and version populated if CPI status
// section is missing from CAPVCD RDE. This method is intended to be called when CPI is started up.
// EnsureLoadBalancer function, which is called for every instance of load balancer type service during start up,
// will take care of lazily updating rest of the information related to CPI.
// For future CPI upgrades, this method may become a factory of converters (ConvertFrom()).
func (client *Client) UpgradeCPIStatusOfExistingRDE(ctx context.Context, rdeId string) error {
	klog.Infof("upgrading CPI section in RDE")
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
		nonCapvcdEntityError := NonCAPVCDEntityError{
			EntityTypeID: rde.EntityType,
		}
		return nonCapvcdEntityError
	}

	statusEntity, ok := rde.Entity["status"]
	if !ok {
		return fmt.Errorf("failed to fetch status section in RDE [%s] when upgrading CPI section", rdeId)
	}
	statusMap, ok := statusEntity.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to convert status section of RDE [%s] to map[string]interface{} during RDE upgrade", rdeId)
	}
	upgradedStatusMap, err := UpgradeCPISectionInStatus(statusMap)
	if err != nil {
		return fmt.Errorf("failed to upgrade CPI section in RDE [%s]: [%v]", rdeId, err)
	}
	rde.Entity["status"] = upgradedStatusMap

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

func extractCPIStatus(statusMap map[string]interface{}) (*util.CPIStatus, error) {
	var cpiStatus *util.CPIStatus
	cpiStatusIf, ok := statusMap["cpi"]
	if !ok {
		// CPI section in the status is absent.
		// Recover by creating a new section
		cpiStatus = &util.CPIStatus{
			Name: CloudControllerManagerName,
			Version: CloudControllerManagerVersion,
			VCDResourceSet: nil,
		}
	} else {
		cpiStatusMap, ok := cpiStatusIf.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to parse CPI status into map[string]interface{}")
		}
		var err error
		cpiStatus, err = convertMapToCPIStatus(cpiStatusMap)
		if err != nil {
			return nil, fmt.Errorf("failed to convert map[string]interface{} to CPIStatus object: [%v]", err)
		}
	}
	if cpiStatus.VCDResourceSet == nil {
		cpiStatus.VCDResourceSet = make([]util.VCDResource, 0)
	}
	return cpiStatus, nil
}

func AddToVCDResourceSet(statusMap map[string]interface{}, vcdResource util.VCDResource) (map[string]interface{}, error) {
	cpiStatus, err := extractCPIStatus(statusMap)
	if err != nil {
		return nil, fmt.Errorf("failed to add VCDResource [%s] of type [%s] to CPI status: [%v]", vcdResource.ID, vcdResource.Type, err)
	}
	resourceFound := false
	for _, r := range cpiStatus.VCDResourceSet {
		if r.Type == vcdResource.Type && r.ID == vcdResource.ID {
			resourceFound = true
			r.Name = vcdResource.Name
			r.AdditionalDetails = vcdResource.AdditionalDetails
		}
	}
	if !resourceFound {
		cpiStatus.VCDResourceSet = append(cpiStatus.VCDResourceSet, vcdResource)
	}
	updatedCPIMap, err := convertCPIStatusToMap(*cpiStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CPI status object to map[string]interface{} to add [%v] to VCDResourceSet: [%v]", vcdResource, err)
	}
	statusMap["cpi"] = updatedCPIMap
	return statusMap, nil
}

func (client *Client) AddToVCDResourceSet(ctx context.Context, resourceType string, resourceName string, resourceId string, additionalDetails map[string]interface{}, rdeID string) error {
	if rdeID == "" || strings.HasPrefix(rdeID, NoRdePrefix) {
		klog.Infof("ClusterID [%s] is empty or generated, hence not removing VCDResource [%s:%s] from RDE",
			rdeID, resourceType, resourceId)
		return nil
	}
	for i := MaxRDEUpdateRetries ; i > 1 ; i -- {
		rde, resp, etag, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, rdeID)
		if resp != nil && resp.StatusCode != http.StatusOK {
			var responseMessageBytes []byte
			if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
				responseMessageBytes = gsErr.Body()
			}
			return fmt.Errorf(
				"failed to get RDE [%s] when adding virtual service to VCD resource set ; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
				rdeID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
		} else if err != nil {
			return fmt.Errorf("error while updating the RDE [%s]: [%v]", rdeID, err)
		}
		if !util.IsCAPVCDEntityType(rde.EntityType) {
			nonCapvcdEntityError := NonCAPVCDEntityError{
				EntityTypeID: rde.EntityType,
			}
			return nonCapvcdEntityError
		}
		statusIf, ok := rde.Entity["status"]
		if !ok {
			return fmt.Errorf("failed to update RDE [%s] with virtual status information", rdeID)
		}
		statusMap, ok := statusIf.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to convert RDE status [%s] to map[string]interface{}", rdeID)
		}
		vcdResource := util.VCDResource{
			Type:              resourceType,
			ID:                resourceId,
			Name:              resourceName,
			AdditionalDetails: additionalDetails,
		}
		updatedStatusMap, err := AddToVCDResourceSet(statusMap, vcdResource)
		if err != nil {
			return fmt.Errorf("error occurred when updating VCDResource set of CPI status in RDE [%s]: [%v]", rdeID, err)
		}
		rde.Entity["status"] = updatedStatusMap
		_, resp, err = client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, rde, etag, rdeID, nil)
		if resp != nil {
			if resp.StatusCode == http.StatusPreconditionFailed {
				klog.Errorf("wrong etag while adding [%v] to VCDResourceSet in RDE [%s]. Retry attempts remaining: [%d]", vcdResource, rdeID, i - 1)
				continue
			} else if resp.StatusCode != http.StatusOK {
				var responseMessageBytes []byte
				if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
					responseMessageBytes = gsErr.Body()
				}
				return fmt.Errorf(
					"failed to add resource [%s] having ID [%s] to VCDResourseSet of CPI in RDE [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
					vcdResource.Name, vcdResource.ID, rdeID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
			}
		} else if err != nil {
			return fmt.Errorf("error while updating the RDE [%s]: [%v]", rdeID, err)
		} else {
			return fmt.Errorf("invalid response obtained when updating VCDResoruceSet of CPI in RDE [%s]", rdeID)
		}
	}
	return nil
}

func RemoveFromCPIVCDResourceSet(statusMap map[string]interface{}, vcdResource util.VCDResource) (map[string]interface{}, error) {
	cpiStatus, err := extractCPIStatus(statusMap)
	if err != nil {
		return nil, fmt.Errorf("failed to remove VCDResource [%s] to CPI status: [%v]", vcdResource.ID, err)
	}
	updatedVcdResourceSet := make([]util.VCDResource, 0)
	for _, r := range cpiStatus.VCDResourceSet {
		if r.Type == vcdResource.Type && r.ID == vcdResource.ID {
			continue
		}
		updatedVcdResourceSet = append(updatedVcdResourceSet, r)
	}
	cpiStatus.VCDResourceSet = updatedVcdResourceSet
	updatedCPIMap, err := convertCPIStatusToMap(*cpiStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CPI status object to map[string]interface{} to remove [%s] from VCDResourceSet: [%v]", vcdResource.ID, err)
	}
	statusMap["cpi"] = updatedCPIMap
	return statusMap, nil
}

func (client *Client) RemoveFromVCDResourceSet(ctx context.Context, resourceType, resourceId string, rdeID string) error {
	if rdeID == "" || strings.HasPrefix(rdeID, NoRdePrefix) {
		klog.Infof("ClusterID [%s] is empty or generated, hence not removing VCDResource [%s:%s] from RDE",
			rdeID, resourceType, resourceId)
		return nil
	}
	for i := MaxRDEUpdateRetries ; i > 1 ; i -- {
		rde, resp, etag, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, rdeID)
		if resp != nil && resp.StatusCode != http.StatusOK {
			var responseMessageBytes []byte
			if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
				responseMessageBytes = gsErr.Body()
			}
			return fmt.Errorf(
				"failed to get RDE with id [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
				rdeID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
		} else if err != nil {
			return fmt.Errorf("error while updating the RDE [%s]: [%v]", rdeID, err)
		}

		if !util.IsCAPVCDEntityType(rde.EntityType) {
			nonCapvcdEntityError := NonCAPVCDEntityError{
				EntityTypeID: rde.EntityType,
			}
			return nonCapvcdEntityError
		}

		statusEntity, ok := rde.Entity["status"]
		if !ok {
			return fmt.Errorf("failed to parse status in RDE [%s] to remove virtual service [%s] from CPI status", rdeID, resourceId)
		}
		statusMap, ok := statusEntity.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to parse status in RDE [%s] into map[string]interface{} to remove virtual service [%s] from CPI status", rdeID, resourceId)
		}
		updatedStatus, err := RemoveFromCPIVCDResourceSet(statusMap, util.VCDResource{
			Type: resourceType,
			ID: resourceId,
		})
		if err != nil {
			return fmt.Errorf("failed to remove resource [%s] from VCDResourceSet in CPI status section of RDE [%s]: [%v]",resourceId, rdeID, err)
		}
		rde.Entity["status"] = updatedStatus

		_, resp, err = client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, rde, etag, rdeID, nil)
		if resp != nil {
			if resp.StatusCode == http.StatusPreconditionFailed {
				klog.Errorf("wrong etag while removing [%v] from VCDResourceSet in RDE [%s]. Retry attempts remaining: [%d]", resourceId, rdeID, i - 1)
				continue
			} else if resp.StatusCode != http.StatusOK {
				var responseMessageBytes []byte
				if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
					responseMessageBytes = gsErr.Body()
				}
				return fmt.Errorf(
					"failed to update CPI status for RDE [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
					rdeID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
			}
		} else if err != nil {
			return fmt.Errorf("error while removing virtual service [%s] from the RDE [%s]: [%v]", resourceId, rdeID, err)
		} else {
			return fmt.Errorf("invalid response obtained when updating VCDResoruceSet of CPI in RDE [%s]", rdeID)
		}
	}
	return nil
}
