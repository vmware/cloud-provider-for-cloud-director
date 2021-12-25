package ccm

import (
	cloudprovider "k8s.io/cloud-provider"
)

func (_ *VCDCloudProvider) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return nil, false
}
