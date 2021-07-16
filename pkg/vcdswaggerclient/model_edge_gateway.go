/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// An edge gateway object 
type EdgeGateway struct {
	// Represents current status of the networking object. 
	Status *NetworkingObjectStatusType `json:"status,omitempty"`
	// The unique identifier of the edge gateway.
	Id string `json:"id,omitempty"`
	// The name of the edge gateway.
	Name string `json:"name,omitempty"`
	// The description of the edge gateway(optional).
	Description string `json:"description,omitempty"`
	// The uplink connections for the edge gateway.
	EdgeGatewayUplinks []EdgeGatewayUplink `json:"edgeGatewayUplinks,omitempty"`
	// A flag indicating whether distributed routing is enabled or not. The default is false.
	DistributedRoutingEnabled bool `json:"distributedRoutingEnabled,omitempty"`
	// The number of Org vDC networks connected to the gateway.
	OrgVdcNetworkCount int32 `json:"orgVdcNetworkCount,omitempty"`
	// The backing details of the edge gateway; only required if importing an NSX-T router.
	GatewayBacking *EdgeGatewayBacking `json:"gatewayBacking,omitempty"`
	// The organization vDC which the gateway belongs to. Property is deprecated. Please use ownerRef. 
	OrgVdc *EntityReference `json:"orgVdc,omitempty"`
	// The organization vDC or vDC Group that this edge gateway belongs to. If the ownerRef is set to a vDC Group, this gateway will be available across all the participating Organization vDCs in the vDC Group. 
	OwnerRef *EntityReference `json:"ownerRef,omitempty"`
	// The organization to which the gateway belongs.
	OrgRef *EntityReference `json:"orgRef,omitempty"`
	// The network definition in CDIR form that DNS and DHCP service on an NSX-T edge will run on. The subnet prefix length must be 27. This property applies to creating or importing an NSX-T Edge. This is not supported for VMC. If nothing is set, the default is 192.168.255.225/27.  The DHCP listener IP network is on 192.168.255.225/30. The DNS listener IP network is on 192.168.255.228/32.  This field cannot be updated. 
	ServiceNetworkDefinition string `json:"serviceNetworkDefinition,omitempty"`
	// Edge Cluster Configuration for the Edge Gateway. Can be specified if a gateway needs to be placed on a specific set of Edge Clusters. For NSX-T Edges, user should specify the ID of the NSX-T edge cluster as the value of primaryEdgeCluster's backingId. The gateway defaults to the Edge Cluster of the connected External Network's backing Tier-0 router, if nothing is specified. The value of secondaryEdgeCluster will be set to NULL for NSX-T edge gateways. For NSX-V Edges, this is read-only and the legacy API must be used for edge specific placement. 
	EdgeClusterConfig *GatewayEdgeClusterConfig `json:"edgeClusterConfig,omitempty"`
}
