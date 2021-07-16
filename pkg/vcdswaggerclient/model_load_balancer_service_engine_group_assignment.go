/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// An assignment of a Load Balancer Service Engine Group to an Edge Gateway
type LoadBalancerServiceEngineGroupAssignment struct {
	// The identifier of the Load Balancer Service Engine Groups in URN format.
	Id string `json:"id,omitempty"`
	// The maximum number of virtual services the Edge Gateway is allowed to use. This is required if the Load Balancer Service Engine Group has reservation type 'SHARED'. This must be unset if the Load Balancer Service Engine Group has reservation type 'DEDICATED'. 
	MaxVirtualServices int32 `json:"maxVirtualServices,omitempty"`
	// The number of guaranteed virtual services available to the Edge Gateway. This is required if the Load Balancer Service Engine Group has reservation type 'SHARED'. This must be unset if the Load Balancer Service Engine Group has reservation type 'DEDICATED'. 
	MinVirtualServices int32 `json:"minVirtualServices,omitempty"`
	// The current number of deployed virutal services. 
	NumDeployedVirtualServices int32 `json:"numDeployedVirtualServices,omitempty"`
	// The associated Load Balancer Service Engine Group.
	ServiceEngineGroupRef *EntityReference `json:"serviceEngineGroupRef"`
	// The associated Edge Gateway.
	GatewayRef *EntityReference `json:"gatewayRef"`
	// The owner of the associated Edge Gateway. This can be a vDC or vDC Group.
	GatewayOwnerRef *EntityReference `json:"gatewayOwnerRef,omitempty"`
	// The organization of the associated Edge Gateway.
	GatewayOrgRef *EntityReference `json:"gatewayOrgRef,omitempty"`
}
