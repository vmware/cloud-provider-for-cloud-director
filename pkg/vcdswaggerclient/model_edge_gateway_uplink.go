/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// The uplink (i.e. external network) that the edge gateway is connected to. This also contains the external network IPs that are allocated to the edge. 
type EdgeGatewayUplink struct {
	// The identifier of the external network this edge gateway is connected to.
	UplinkId string `json:"uplinkId,omitempty"`
	// The name of the external network this edge gateway is connected to.
	UplinkName string `json:"uplinkName,omitempty"`
	// Set of subnets this edge will have access to. 
	Subnets *EdgeGatewaySubnets `json:"subnets,omitempty"`
	// Whether or not the gateway is connected to this uplink. This value defaults to true if it is not set. When filtering by this field, if the filter is false, all gateways that have 0 connected uplinks will be returned, while if it is true, all gateways with at least one connected uplink will be returned.
	Connected bool `json:"connected,omitempty"`
	// If set on create or update api calls, the specified number of IP addresses will be additionally allocated for this uplink. IPs will be allocated from multiple subnets if needed.
	QuickAddAllocatedIpCount int32 `json:"quickAddAllocatedIpCount,omitempty"`
	// If set to true, then the associated external network is exclusively used by this edge gateway.
	Dedicated bool `json:"dedicated,omitempty"`
	// Whether the associated external network is backed by a NSX-T VRF-Lite Tier-0.
	VrfLiteBacked bool `json:"vrfLiteBacked,omitempty"`
}
