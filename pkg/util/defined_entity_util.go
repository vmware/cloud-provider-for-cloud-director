package util

import (
	"fmt"
	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"strings"
)

const (
	CAPVCDEntityTypeVendor = "vmware"
	CAPVCDEntityTypeNss    = "capvcdCluster"

	NativeClusterEntityTypeVendor = "cse"
	NativeClusterEntityTypeNss    = "nativeCluster"
)

type VCDResource struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CPIStatus struct {
	Name           string        `json:"name,omitempty"`
	Version        string        `json:"version:omitempty"`
	VCDResourceSet []VCDResource `json:"vcdResourceSet,omitempty"`
	Errors         []string      `json:"errors,omitempty"`
	VirtualIPs     []string      `json:"virtualIPs,omitempty"`
}

func isCAPVCDEntityType(entityTypeID string) bool {
	entityTypeIDSplit := strings.Split(entityTypeID, ":")
	// format is urn:vcloud:type:<vendor>:<nss>:<version>
	if len(entityTypeIDSplit) != 6 {
		return false
	}
	return entityTypeIDSplit[3] == CAPVCDEntityTypeVendor && entityTypeIDSplit[4] == CAPVCDEntityTypeNss
}

func isNativeClusterEntityType(entityTypeID string) bool {
	entityTypeIDSplit := strings.Split(entityTypeID, ":")
	// format is urn:vcloud:type:<vendor>:<nss>:<version>
	if len(entityTypeIDSplit) != 6 {
		return false
	}
	return entityTypeIDSplit[3] == NativeClusterEntityTypeVendor && entityTypeIDSplit[4] == NativeClusterEntityTypeNss
}

func GetVirtualIPsFromRDE(rde *swaggerClient.DefinedEntity) ([]string, error) {
	statusEntry, ok := rde.Entity["status"]
	if !ok {
		return nil, fmt.Errorf("could not find 'status' entry in defined entity")
	}
	statusMap, ok := statusEntry.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unable to convert [%T] to map", statusEntry)
	}

	var virtualIpInterfaces interface{}
	if isCAPVCDEntityType(rde.EntityType) {
		// TODO: upgrade existing RDEs to 1.1.0 version
		cpiStatusInterface, ok := statusMap["cpi"]
		if !ok {
			return nil, fmt.Errorf("CAPVCD entity [%s] is missing CPI status", rde.Id)
		}
		cpiStatusMap, ok := cpiStatusInterface.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to convert value for CPI in RDE status to map[string]interface{}")
		}
		virtualIpInterfaces = cpiStatusMap["virtualIPs"]
	}
	if isNativeClusterEntityType(rde.EntityType) {
		virtualIpInterfaces = statusMap["virtual_IPs"]
	} else {
		return nil, fmt.Errorf("entity type %s not supported by CPI", rde.EntityType)
	}

	if virtualIpInterfaces == nil {
		return make([]string, 0), nil
	}

	virtualIpInterfacesSlice, ok := virtualIpInterfaces.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unable to convert [%T] to slice of interface", virtualIpInterfaces)
	}
	virtualIpStrs := make([]string, len(virtualIpInterfacesSlice))
	for ind, ipInterface := range virtualIpInterfacesSlice {
		currIp, ok := ipInterface.(string)
		if !ok {
			return nil, fmt.Errorf("unable to convert [%T] to string", ipInterface)
		}
		virtualIpStrs[ind] = currIp
	}
	return virtualIpStrs, nil
}

// ReplaceVirtualIPsInRDE replaces the virtual IPs array in the inputted rde. It does not make an API call to update
// the RDE.
func ReplaceVirtualIPsInRDE(rde *swaggerClient.DefinedEntity, updatedIps []string) (*swaggerClient.DefinedEntity, error) {
	statusEntry, ok := rde.Entity["status"]
	if !ok {
		return nil, fmt.Errorf("could not find 'status' entry in defined entity")
	}
	statusMap, ok := statusEntry.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unable to convert [%T] to map", statusEntry)
	}
	if isCAPVCDEntityType(rde.EntityType) {
		// TODO: upgrade existing entities to capvcdCluster 1.1.0
		cpiStatusInterface, ok := statusMap["cpi"]
		var cpiStatusMap map[string]interface{}
		if !ok {
			cpiStatusMap = make(map[string]interface{})
			cpiStatusMap["virtualIPs"] = updatedIps
		} else {
			cpiStatusMap, ok = cpiStatusInterface.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("failed to parse CPI status")
			}
			cpiStatusMap["virtualIPs"] = updatedIps
		}
		statusMap["cpi"] = cpiStatusMap
	} else if isNativeClusterEntityType(rde.EntityType) {
		statusMap["virtual_IPs"] = updatedIps
	}
	return rde, nil
}
