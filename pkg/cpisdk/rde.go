package cpisdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"github.com/vmware/cloud-provider-for-cloud-director/release"
	"k8s.io/klog"
	"net/http"
	"strings"
	"time"
)

const (
	// Errors
	CreateLoadbalancerError  = "CreateLoadbalancerError"
	UpdateLoadbalancerError  = "UpdateLoadbalancerError"
	DeleteLoadbalancerError  = "DeleteLoadbalancerError"
	GetLoadbalancerError     = "GetLoadbalancerError"
	CPIStatusUpgradeRdeError = "CPIStatusUpgradeRdeError"
	RemoveVIPFromRdeError    = "RemoveVirtualIPFromRdeError"
	AddVIPToRdeError         = "AddVirtualIPToRdeError"

	// Events
	ClientAuthenticated  = "ClientAuthenticated"
	CreatedLoadbalancer  = "CreatedLoadbalancer"
	UpdatedLoadbalancer  = "UpdatedLoadbalancer"
	DeletedLoadbalancer  = "DeletedLoadbalancer"
	CPIStatusRDEUpgraded = "CPIStatusRDEUpgraded"
)

type CPIRDEManager struct {
	// Client will be refreshed separately
	RDEManager *vcdsdk.RDEManager
}

func NewCPIRDEManager(rdeManager *vcdsdk.RDEManager) *CPIRDEManager {
	return &CPIRDEManager{
		RDEManager: rdeManager,
	}
}

func (cpiRDEManager *CPIRDEManager) GetRDEVirtualIps(ctx context.Context) ([]string, string, *swaggerClient.DefinedEntity, error) {
	if cpiRDEManager.RDEManager.ClusterID == "" || strings.HasPrefix(cpiRDEManager.RDEManager.ClusterID, vcdsdk.NoRdePrefix) {
		klog.Infof("ClusterID [%s] is empty or generated", cpiRDEManager.RDEManager.ClusterID)
		return nil, "", nil, nil
	}

	client := cpiRDEManager.RDEManager.Client
	clusterOrg, err := client.VCDClient.GetOrgByName(client.ClusterOrgName)
	if err != nil {
		return nil, "", nil, fmt.Errorf("unable to get org for org [%s]: [%v]", client.ClusterOrgName, err)
	}
	if clusterOrg == nil || clusterOrg.Org == nil {
		return nil, "", nil, fmt.Errorf("obtained nil org for name [%s]", client.ClusterOrgName)
	}
	defEnt, _, etag, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, cpiRDEManager.RDEManager.ClusterID, clusterOrg.Org.ID)
	if err != nil {
		return nil, "", nil, fmt.Errorf("error when getting defined entity: [%v]", err)
	}

	virtualIpStrs, err := vcdsdk.GetVirtualIPsFromRDE(&defEnt)
	if err != nil {
		capvcdEntityFoundErr, ok := err.(vcdsdk.CapvcdRdeFoundError)
		if ok {
			return nil, "", nil, capvcdEntityFoundErr
		}
		return nil, "", nil, fmt.Errorf("failed to retrieve Virtual IPs from RDE [%s]: [%v]",
			cpiRDEManager.RDEManager.ClusterID, err)
	}
	return virtualIpStrs, etag, &defEnt, nil
}

// This function will modify the passed in defEnt
func (cpiRDEManager *CPIRDEManager) updateRDEVirtualIps(ctx context.Context, updatedIps []string, etag string,
	defEnt *swaggerClient.DefinedEntity) (*http.Response, error) {
	defEnt, err := vcdsdk.ReplaceVirtualIPsInRDE(defEnt, updatedIps)
	if err != nil {
		if capvcdEntityFoundErr, ok := err.(vcdsdk.CapvcdRdeFoundError); ok {
			return nil, capvcdEntityFoundErr
		}
		return nil, fmt.Errorf("failed to locally edit RDE with ID [%s] with virtual IPs: [%v]", cpiRDEManager.RDEManager.ClusterID, err)
	}
	client := cpiRDEManager.RDEManager.Client
	clusterOrg, err := client.VCDClient.GetOrgByName(client.ClusterOrgName)
	if err != nil {
		return nil, fmt.Errorf("unable to get org for org [%s]: [%v]", client.ClusterOrgName, err)
	}
	if clusterOrg == nil || clusterOrg.Org == nil {
		return nil, fmt.Errorf("obtained nil org for name [%s]", client.ClusterOrgName)
	}
	// can pass invokeHooks
	_, httpResponse, err := client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, *defEnt, etag, cpiRDEManager.RDEManager.ClusterID, clusterOrg.Org.ID, nil)
	if err != nil {
		return httpResponse, fmt.Errorf("error when updating defined entity [%s]: [%v]", cpiRDEManager.RDEManager.ClusterID, err)
	}
	return httpResponse, nil
}

func (cpiRDEManager *CPIRDEManager) AddVirtualIpToRDE(ctx context.Context, addIp string) error {
	if addIp == "" {
		klog.Infof("VIP is empty, hence not adding anything to RDE")
		return nil
	}
	if cpiRDEManager.RDEManager.ClusterID == "" || strings.HasPrefix(cpiRDEManager.RDEManager.ClusterID, vcdsdk.NoRdePrefix) {
		klog.Infof("ClusterID [%s] is empty or generated, hence not adding VIP [%s] from RDE",
			cpiRDEManager.RDEManager.ClusterID, addIp)
		return nil
	}

	numRetries := 10
	for i := 0; i < numRetries; i++ {
		currIps, etag, defEnt, err := cpiRDEManager.GetRDEVirtualIps(ctx)
		if err != nil {
			if _, ok := err.(vcdsdk.CapvcdRdeFoundError); ok {
				klog.Infof("CAPVCD entity type found. Skipping adding RDE VIPs to status")
				return nil
			}
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
		httpResponse, err := cpiRDEManager.updateRDEVirtualIps(ctx, updatedIps, etag, defEnt)
		if err != nil {
			if capvcdEntityFoundErr, ok := err.(vcdsdk.CapvcdRdeFoundError); ok {
				return capvcdEntityFoundErr
			}
			if httpResponse.StatusCode == http.StatusPreconditionFailed {
				klog.Infof("Wrong ETag while adding virtual IP [%s]", addIp)
				continue
			}
			return fmt.Errorf("error when adding virtual ip [%s] to RDE: [%v]", addIp, err)
		}
		klog.Infof("Successfully updated RDE [%s] with virtual IP [%s]", cpiRDEManager.RDEManager.ClusterID, addIp)
		return nil
	}

	return fmt.Errorf("unable to update rde due to incorrect etag after [%d]] tries", numRetries)
}

func (cpiRDEManager *CPIRDEManager) RemoveVirtualIpFromRDE(ctx context.Context, removeIp string) error {
	if removeIp == "" {
		klog.Infof("VIP is empty, hence not removing anything from RDE")
		return nil
	}
	if cpiRDEManager.RDEManager.ClusterID == "" || strings.HasPrefix(cpiRDEManager.RDEManager.ClusterID, vcdsdk.NoRdePrefix) {
		klog.Infof("ClusterID [%s] is empty or generated, hence not removing VIP [%s] from RDE",
			cpiRDEManager.RDEManager.ClusterID, removeIp)
		return nil
	}

	numRetries := 10
	for i := 0; i < numRetries; i++ {
		currIps, etag, defEnt, err := cpiRDEManager.GetRDEVirtualIps(ctx)
		if err != nil {
			if _, ok := err.(vcdsdk.CapvcdRdeFoundError); ok {
				klog.Infof("CAPVCD entity found. Skip removing VIPs from RDE in the status")
				return nil
			}
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

		httpResponse, err := cpiRDEManager.updateRDEVirtualIps(ctx, updatedIps, etag, defEnt)
		if err != nil {
			if capvcdEntityFoundErr, ok := err.(vcdsdk.CapvcdRdeFoundError); ok {
				return capvcdEntityFoundErr
			}
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

func convertCPIStatusToMap(cpiStatus vcdsdk.CPIStatus) (map[string]interface{}, error) {
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

func convertMapToCPIStatus(cpiStatusMap map[string]interface{}) (*vcdsdk.CPIStatus, error) {
	cpiStatusMapBytes, err := json.Marshal(&cpiStatusMap)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CPI status map to byte array: [%v]", err)
	}
	var cpiStatus vcdsdk.CPIStatus
	err = json.Unmarshal(cpiStatusMapBytes, &cpiStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CPI status bytes to CPIStatus: [%v]", err)
	}
	return &cpiStatus, nil
}

// UpgradeCPISectionInStatus provides a helper function to update the latest CPI status into local entity
// Provide an automatic conversion of the content in srcCapvcdEntity.entity.status.cpi content to the latest RDE version format (vcdsdk.CPIStatus)
// Add the placeholder for any special conversion logic inside vcdsdk.CPIStatus (for developers)
// Convert vcdsdk.CPIStatus to cpiStatusMap as the content in srcCapvcdEntity.entity.status.cpi; return cpiStatusMap
func UpgradeCPISectionInStatus(statusMap map[string]interface{}) (map[string]interface{}, error) {
	if statusMap == nil {
		return nil, fmt.Errorf("invalid value for status: [%v]", statusMap)
	}
	cpiStatus := &vcdsdk.CPIStatus{
		VCDResourceSet: nil,
		Errors:         nil,
		Events:         nil,
		VirtualIPs:     nil,
	}
	cpiStatusEntity, ok := statusMap[vcdsdk.ComponentCPI]
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
	// ******************  placeHolder: add any special conversion logic for CPIStatus  ******************
	// Say CPI 1.1 has properties {X, Y} and CPI 1.2 introduces new property Z; {X, Y, Z}
	// CPIStatus should update with {X, Y, Z=default}. Developers should update property Z at this place.
	//PUT RDE.status.CPI should update with {X, Y, Z=updatedValue}
	cpiStatus.Name = release.CloudControllerManagerName
	cpiStatus.Version = release.CpiVersion
	// CPI section is missing from the RDE status. Create a new CPI status and update the RDE.
	cpiStatusMap, err := convertCPIStatusToMap(*cpiStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to convert CPI status [%v] in RDE to map[string]interface{} when upgrading CPI status section: [%v]", cpiStatus, err)
	}
	statusMap[vcdsdk.ComponentCPI] = cpiStatusMap
	return statusMap, nil
}

// UpgradeCPIStatusOfExistingRDE Creates an empty cpi section with just the name and version populated if CPI status
// section is missing from CAPVCD RDE. This method is intended to be called when CPI is started up.
// EnsureLoadBalancer function, which is called for every instance of load balancer type service during start up,
// will take care of lazily updating rest of the information related to CPI.
// Provide an automatic conversion of the content in srcEntity.entity.status.cpi content to the latest RDE version format (CPIStatus)
// Add the placeholder for any special conversion logic inside CSIStatus inside UpgradeCPISectionInStatus() (for developers)
// Call an API call (PUT) to update CAPVCD entity and persist data into VCD
// For future CPI upgrades, this method may become a factory of converters (ConvertFrom()).
func (cpiRDEManager *CPIRDEManager) UpgradeCPIStatusOfExistingRDE(ctx context.Context, rdeId string) error {
	klog.Infof("upgrading CPI section in RDE")
	client := cpiRDEManager.RDEManager.Client
	clusterOrg, err := client.VCDClient.GetOrgByName(client.ClusterOrgName)
	if err != nil {
		return fmt.Errorf("unable to get org for org [%s]: [%v]", client.ClusterOrgName, err)
	}
	if clusterOrg == nil || clusterOrg.Org == nil {
		return fmt.Errorf("obtained nil org for name [%s]", client.ClusterOrgName)
	}
	for retries := 0; retries < vcdsdk.MaxRDEUpdateRetries; retries++ {
		rde, resp, etag, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, rdeId, clusterOrg.Org.ID)
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

		if !vcdsdk.IsCAPVCDEntityType(rde.EntityType) {
			nonCapvcdEntityError := vcdsdk.NonCAPVCDEntityError{
				EntityTypeID: rde.EntityType,
			}
			return nonCapvcdEntityError
		}

		statusEntity, ok := rde.Entity["status"]
		if !ok {
			return fmt.Errorf("failed to fetch status section in RDE [%s] when upgrading CPI section",
				rdeId)
		}
		statusMap, ok := statusEntity.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to convert status section of RDE [%s] to map[string]interface{} during RDE upgrade",
				rdeId)
		}
		_, err = UpgradeCPISectionInStatus(statusMap)
		if err != nil {
			return fmt.Errorf("failed to upgrade CPI section in RDE [%s]: [%v]", rdeId, err)
		}

		_, resp, err = client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, rde, etag, rdeId, clusterOrg.Org.ID, nil)
		if resp != nil && resp.StatusCode != http.StatusOK {
			var responseMessageBytes []byte
			if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
				responseMessageBytes = gsErr.Body()
			}
			klog.Warningf(
				"failed to create CPI status for RDE [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]. Remaining retry attempts: [%d]",
				rdeId, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err, vcdsdk.MaxRDEUpdateRetries-retries+1)
			continue
		} else if err != nil {
			klog.Warningf("error while getting the RDE [%s]: [%v]. Remaining retry attempts: [%d]", rdeId, err, vcdsdk.MaxRDEUpdateRetries-retries+1)
			continue
		}
		klog.Infof("successfully upgraded CPI status section of the RDE [%s]", rdeId)
		return nil
	}
	return fmt.Errorf("failed to upgrade CPI status section of the RDE [%s] after [%d] retries", rdeId, vcdsdk.MaxRDEUpdateRetries)
}

// AddVIPToVCDResourceSet adds virtual IP to the RDE and removes the VIP from older "status.virtualIPs" section if present
func (cpiRDEManager *CPIRDEManager) AddVIPToVCDResourceSet(ctx context.Context, vsName string, vsID string, externalIP string) error {
	if cpiRDEManager.RDEManager == nil {
		return fmt.Errorf("RDEManager not initialized in [%T]", cpiRDEManager)
	}
	// - remove the VIP from status["virtualIP"] which matches "externalIP" value
	// - add VIP to vcdresourceset[]
	// - update rde
	if cpiRDEManager.RDEManager.ClusterID == "" || strings.HasPrefix(cpiRDEManager.RDEManager.ClusterID, vcdsdk.NoRdePrefix) {
		// Indicates that the RDE ID is either empty or it was auto-generated.
		klog.Infof("ClusterID [%s] is empty or generated, hence not adding VCDResource [%s] of type [%s] from RDE",
			vsName, vcdsdk.VcdResourceVirtualService, vsID)
		return nil
	}
	client := cpiRDEManager.RDEManager.Client
	clusterOrg, err := client.VCDClient.GetOrgByName(client.ClusterOrgName)
	if err != nil {
		return fmt.Errorf("unable to get org for org [%s]: [%v]", client.ClusterOrgName, err)
	}
	if clusterOrg == nil || clusterOrg.Org == nil {
		return fmt.Errorf("obtained nil org for name [%s]", client.ClusterOrgName)
	}
	for i := vcdsdk.MaxRDEUpdateRetries; i > 1; i-- {
		rde, resp, etag, err := cpiRDEManager.RDEManager.Client.APIClient.DefinedEntityApi.GetDefinedEntity(
			ctx, cpiRDEManager.RDEManager.ClusterID, clusterOrg.Org.ID)
		if resp != nil && resp.StatusCode != http.StatusOK {
			var responseMessageBytes []byte
			if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
				responseMessageBytes = gsErr.Body()
			}
			return fmt.Errorf(
				"failed to get RDE [%s] when adding resourse to VCD resource set ; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
				cpiRDEManager.RDEManager.ClusterID, http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
		} else if err != nil {
			return fmt.Errorf("error while updating the RDE [%s]: [%v]", cpiRDEManager.RDEManager.ClusterID, err)
		}
		// Only entity of type capvcdCluster should be updated.
		if !vcdsdk.IsCAPVCDEntityType(rde.EntityType) {
			klog.V(3).Infof("entity type of RDE [%s] is [%s]. skipping adding resource [%s] of type [%s] to status of component [%s]",
				rde.Id, rde.EntityType, vsName, vcdsdk.VcdResourceVirtualService, vcdsdk.ComponentCPI)
			return nil
		}
		statusIf, ok := rde.Entity["status"]
		if !ok {
			return fmt.Errorf("failed to update RDE [%s] with VCDResource set information", cpiRDEManager.RDEManager.ClusterID)
		}
		statusMap, ok := statusIf.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to convert RDE status [%s] to map[string]interface{}", cpiRDEManager.RDEManager.ClusterID)
		}
		vcdResource := vcdsdk.VCDResource{
			Type: vcdsdk.VcdResourceVirtualService,
			ID:   vsID,
			Name: vsName,
			AdditionalDetails: map[string]interface{}{
				"virtualIP": externalIP,
			},
		}
		updatedStatusMap, rdeUpdateRequired, err := vcdsdk.AddVCDResourceToStatusMap(vcdsdk.ComponentCPI, cpiRDEManager.RDEManager.StatusComponentName,
			cpiRDEManager.RDEManager.StatusComponentVersion, statusMap, vcdResource)
		if err != nil {
			return fmt.Errorf("error occurred when updating VCDResource set of %s status in RDE [%s]: [%v]",
				cpiRDEManager.RDEManager.ClusterID, vcdsdk.ComponentCPI, err)
		}
		if !rdeUpdateRequired {
			klog.V(3).Info("VCD resource set for the RDE [%s(%s)] already has the resource [%v] in the status of the component [%s]",
				rde.Name, rde.Id, vcdResource, vcdsdk.ComponentCPI)
			return nil
		}

		// remove the VIP from old "status.virtualIPs" key if present
		if virtualIPsIf, ok := updatedStatusMap["virtualIPs"]; ok {
			if virtualIPsArr, ok := virtualIPsIf.([]interface{}); ok {
				updatedVIPs := make([]string, 0)
				for _, vipIf := range virtualIPsArr {
					vipStr, ok := vipIf.(string)
					if ok && vipStr != externalIP {
						updatedVIPs = append(updatedVIPs, vipStr)
					}
					if !ok {
						klog.V(3).Infof("failed to convert VIP [%v] in RDE [%s] to string",
							vipIf, cpiRDEManager.RDEManager.ClusterID)
					}
				}
				if len(updatedVIPs) == 0 {
					delete(updatedStatusMap, "virtualIPs")
				} else {
					updatedStatusMap["virtualIPs"] = updatedVIPs
				}
			}
		}

		rde.Entity["status"] = updatedStatusMap

		_, resp, err = cpiRDEManager.RDEManager.Client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, rde, etag,
			cpiRDEManager.RDEManager.ClusterID, clusterOrg.Org.ID, nil)
		if resp != nil {
			if resp.StatusCode == http.StatusPreconditionFailed {
				klog.Errorf("wrong etag while adding [%v] to VCDResourceSet in RDE [%s]. Retry attempts remaining: [%d]",
					vcdResource, cpiRDEManager.RDEManager.ClusterID, i-1)
				continue
			} else if resp.StatusCode != http.StatusOK {
				var responseMessageBytes []byte
				if gsErr, ok := err.(swaggerClient.GenericSwaggerError); ok {
					responseMessageBytes = gsErr.Body()
				}
				return fmt.Errorf(
					"failed to add resource [%s] having ID [%s] to VCDResourseSet of %s in RDE [%s]; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
					vcdResource.Name, vcdResource.ID, vcdsdk.ComponentCPI, cpiRDEManager.RDEManager.ClusterID,
					http.StatusOK, resp.StatusCode, string(responseMessageBytes), err)
			}
			// resp.StatusCode is http.StatusOK
			klog.Infof("successfully added resource [%s] having ID [%s] to VCDResourceSet of [%s] in RDE [%s]",
				vcdResource.Name, vcdResource.ID, vcdsdk.ComponentCPI, cpiRDEManager.RDEManager.ClusterID)
			return nil
		} else if err != nil {
			return fmt.Errorf("error while updating the RDE [%s]: [%v]", cpiRDEManager.RDEManager.ClusterID, err)
		} else {
			return fmt.Errorf("invalid response obtained when updating VCDResoruceSet of %s in RDE [%s]",
				vcdsdk.ComponentCPI, cpiRDEManager.RDEManager.ClusterID)
		}
	}

	return fmt.Errorf("failed to update RDE [%s] with VCDResourceSet [%s] of type [%s] in [%s] status section",
		cpiRDEManager.RDEManager.ClusterID, vsName, vcdsdk.VcdResourceVirtualService, vcdsdk.ComponentCPI)
}
func (cpiRdeManager *CPIRDEManager) AddToErrorSet(ctx context.Context, errorName, vcdResourceId, detailedErrorMessage string) error {
	backendErr := vcdsdk.BackendError{
		Name:          errorName,
		OccurredAt:    time.Now(),
		VcdResourceId: vcdResourceId,
	}
	if detailedErrorMessage != "" {
		backendErr.AdditionalDetails = map[string]interface{}{"Detailed Error": detailedErrorMessage}
	}
	return cpiRdeManager.RDEManager.AddToErrorSet(ctx, vcdsdk.ComponentCPI, backendErr, vcdsdk.DefaultRollingWindowSize)
}

func (cpiRdeManager *CPIRDEManager) AddToErrorSetWithNameAndId(ctx context.Context, errorName, vcdResourceId, vcdResourceName, detailedErrorMessage string) error {
	backendErr := vcdsdk.BackendError{
		Name:            errorName,
		OccurredAt:      time.Now(),
		VcdResourceId:   vcdResourceId,
		VcdResourceName: vcdResourceName,
	}
	if detailedErrorMessage != "" {
		backendErr.AdditionalDetails = map[string]interface{}{"Detailed Error": detailedErrorMessage}
	}
	return cpiRdeManager.RDEManager.AddToErrorSet(ctx, vcdsdk.ComponentCPI, backendErr, vcdsdk.DefaultRollingWindowSize)
}

func (cpiRdeManager *CPIRDEManager) AddToEventSet(ctx context.Context, eventName, vcdResourceId, detailedEventMsg string) error {
	backendEvent := vcdsdk.BackendEvent{
		Name:          eventName,
		OccurredAt:    time.Now(),
		VcdResourceId: vcdResourceId,
	}
	if detailedEventMsg != "" {
		backendEvent.AdditionalDetails = map[string]interface{}{"Detailed Event": detailedEventMsg}
	}
	return cpiRdeManager.RDEManager.AddToEventSet(ctx, vcdsdk.ComponentCPI, backendEvent, vcdsdk.DefaultRollingWindowSize)
}

func (cpiRdeManager *CPIRDEManager) AddToEventSetWithNameAndId(ctx context.Context, eventName, vcdResourceId, vcdResourceName, detailedEventMsg string) error {
	backendEvent := vcdsdk.BackendEvent{
		Name:            eventName,
		OccurredAt:      time.Now(),
		VcdResourceId:   vcdResourceId,
		VcdResourceName: vcdResourceName,
	}
	if detailedEventMsg != "" {
		backendEvent.AdditionalDetails = map[string]interface{}{"Detailed Event": detailedEventMsg}
	}
	return cpiRdeManager.RDEManager.AddToEventSet(ctx, vcdsdk.ComponentCPI, backendEvent, vcdsdk.DefaultRollingWindowSize)
}
