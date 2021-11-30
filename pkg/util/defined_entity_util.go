package util

import (
	"fmt"
	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
)

const (
	EntityTypePrefix = "urn:vcloud:type"
	CAPVCDEntityTypeVendor = "vmware"
	CAPVCDEntityTypeNss = "capvcd"
	CAPVCDEntityTypeVersion = "1.0.0"

	NativeClusterEntityTypeVendor = "cse"
	NativeClusterEntityTypeNss = "nativeCluster"
	NativeClusterEntityTypeVersion = "2.0.0"
)

var (
	CAPVCDEntityTypeID = fmt.Sprintf("%s:%s:%s:%s", EntityTypePrefix, CAPVCDEntityTypeVendor, CAPVCDEntityTypeNss, CAPVCDEntityTypeVersion)
	NativeEntityTypeID = fmt.Sprintf("%s:%s:%s:%s", EntityTypePrefix, NativeClusterEntityTypeVendor, NativeClusterEntityTypeNss, NativeClusterEntityTypeVersion)
)

func GetVirtualIPFromRDE(rde  *swaggerClient.DefinedEntity) ([]string, error) {
	statusEntry, ok := rde.Entity["status"]
	if !ok {
		return nil, fmt.Errorf("could not find 'status' entry in defined entity")
	}
	statusMap, ok := statusEntry.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unable to convert [%T] to map", statusEntry)
	}

	var virtualIpInterfaces interface{}
	if rde.EntityType == CAPVCDEntityTypeID {
		virtualIpInterfaces := statusMap["virtualIPs"]
		if virtualIpInterfaces == nil {
			return make([]string, 0), nil
		}
	} else if rde.EntityType == NativeEntityTypeID {
		virtualIpInterfaces := statusMap["virtual_IPs"]
		if virtualIpInterfaces == nil {
			return make([]string, 0), nil
		}
	} else {
		return nil, fmt.Errorf("entity type %s not supported by CPI", rde.EntityType)
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

func UpdateVirtualIPsInRDE(rde *swaggerClient.DefinedEntity, updatedIps []string) (*swaggerClient.DefinedEntity, error) {
	statusEntry, ok := rde.Entity["status"]
	if !ok {
		return nil, fmt.Errorf("could not find 'status' entry in defined entity")
	}
	statusMap, ok := statusEntry.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unable to convert [%T] to map", statusEntry)
	}
	if rde.EntityType == CAPVCDEntityTypeID {
		statusMap["virtualIPs"] = updatedIps
	} else if rde. EntityType == NativeEntityTypeID {
		statusMap["virtual_IPs"] = updatedIps
	}
	return rde, nil
}
