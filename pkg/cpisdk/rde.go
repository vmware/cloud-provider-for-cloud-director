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
	defEnt, _, etag, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, cpiRDEManager.RDEManager.ClusterID)
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
	// can pass invokeHooks
	_, httpResponse, err := client.APIClient.DefinedEntityApi.UpdateDefinedEntity(ctx, *defEnt, etag, cpiRDEManager.RDEManager.ClusterID, nil)
	if err != nil {
		return httpResponse, fmt.Errorf("error when updating defined entity [%s]: [%v]", cpiRDEManager.RDEManager.ClusterID, err)
	}
	return httpResponse, nil
}

func (cpiRDEManager *CPIRDEManager) addVirtualIpToRDE(ctx context.Context, addIp string) error {
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

func (cpiRDEManager *CPIRDEManager) removeVirtualIpFromRDE(ctx context.Context, removeIp string) error {
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

func UpgradeCPISectionInStatus(statusMap map[string]interface{}) (map[string]interface{}, error) {
	if statusMap == nil {
		return nil, fmt.Errorf("invalid value for status: [%v]", statusMap)
	}
	cpiStatus := &vcdsdk.CPIStatus{
		VCDResourceSet: nil,
		Errors:         nil,
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
// For future CPI upgrades, this method may become a factory of converters (ConvertFrom()).
func (cpiRDEManager *CPIRDEManager) UpgradeCPIStatusOfExistingRDE(ctx context.Context, rdeId string) error {
	klog.Infof("upgrading CPI section in RDE")
	client := cpiRDEManager.RDEManager.Client
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
	upgradedStatusMap, err := UpgradeCPISectionInStatus(statusMap)
	if err != nil {
		return fmt.Errorf("failed to upgrade CPI section in RDE [%s]: [%v]", rdeId, err)
	}
	rde.Entity[vcdsdk.ComponentCPI] = upgradedStatusMap

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
