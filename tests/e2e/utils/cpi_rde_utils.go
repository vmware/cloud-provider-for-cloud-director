package utils

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/testingsdk"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
)

func IsExternalIpInVCDResourceSet(ctx context.Context, tc *testingsdk.TestClient, externalIp string) (bool, error) {
	virtualIps, err := getVirtualIpsFromRDE(ctx, tc.VcdClient, tc.ClusterId)
	if err != nil {
		return false, fmt.Errorf("error retrieving virtual ips from VCDResourceSet in [%s] status: [%v]", vcdsdk.ComponentCPI, err)
	}

	if virtualIps.Contains(externalIp) {
		return true, nil
	}
	return false, fmt.Errorf("external ip [%s] does not exist in VCDResourceSet from [%s]", externalIp, vcdsdk.ComponentCPI)
}

func getVirtualIpsFromRDE(ctx context.Context, client *vcdsdk.Client, clusterId string) (*testingsdk.Set, error) {
	vcdResources, err := testingsdk.GetVCDResourceSet(ctx, client, clusterId, vcdsdk.ComponentCPI)
	if err != nil {
		return nil, fmt.Errorf("error getting VCDResourceSet in [%s] status: [%v]", vcdsdk.ComponentCPI, err)
	}

	listOfIps := testingsdk.NewSet()
	for _, vcdResource := range vcdResources {
		if vcdResource.Type == vcdsdk.VcdResourceVirtualService && vcdResource.AdditionalDetails != nil {
			virtualIp := vcdResource.AdditionalDetails["virtualIP"].(string)
			listOfIps.Add(virtualIp)
		}
	}
	return listOfIps, nil
}
