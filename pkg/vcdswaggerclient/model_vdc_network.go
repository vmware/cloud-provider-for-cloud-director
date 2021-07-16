/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// An organization vDC network. 
type VdcNetwork struct {
	// The unique ID for the network. This field is read-only.
	Id string `json:"id,omitempty"`
	// The name of the network.
	Name string `json:"name,omitempty"`
	// The description of the network.
	Description string `json:"description,omitempty"`
	// List of subnets configured for the network.
	Subnets *Subnets `json:"subnets,omitempty"`
	// The NSX id of the backing network.
	BackingNetworkId string `json:"backingNetworkId,omitempty"`
	// The object type of the backing network.
	BackingNetworkType *BackingNetworkType `json:"backingNetworkType,omitempty"`
	// The parent network if the network is a direct network, otherwise it will be null.
	ParentNetworkId *EntityReference `json:"parentNetworkId,omitempty"`
	// The type of network. Changing the network type allows converting between an isolated and routed network. Note that the \"connection\" field must also be set if converting from isolated to routed network. 
	NetworkType *VdcNetworkFenceType `json:"networkType,omitempty"`
	// The organization vDC the network belongs to. This should be unset if the network is owned by a vDC Group. For API version 35.0 and above, this field will be treated as read only. Please use ownerRef for new network creation. 
	OrgVdc *EntityReference `json:"orgVdc,omitempty"`
	// The org vDC or vDC Group that this network belongs to. If the ownerRef is set to a vDC Group, this network will be available across all the vDCs in the vDC Group. If the vDC Group is backed by a NSX-V network provider, the org vDC network is automatically connected to the distributed router associated with the vDC Group and the \"connection\" field does not need to be set. For API version 35.0 and above, this field should be set for network creation. 
	OwnerRef *EntityReference `json:"ownerRef,omitempty"`
	// The organization to which the network belongs.
	OrgRef *EntityReference `json:"orgRef,omitempty"`
	// The edge gateway that the network is attached to. 
	Connection *RouterConnection `json:"connection,omitempty"`
	// Deprecated unused field, this property will be removed in future release. 
	IsDefaultNetwork bool `json:"isDefaultNetwork,omitempty"`
	// Whether this network is shared with other organization vDCs.
	Shared bool `json:"shared,omitempty"`
	// Whether or not this network will support two subnets
	EnableDualSubnetNetwork bool `json:"enableDualSubnetNetwork,omitempty"`
	// Description of the network's status. 
	Status *OrgVdcNetworkStatus `json:"status,omitempty"`
	// Brief failure message if the last configuration task failed. Deprecated in Api 33.0, this property will be removed in next release. 
	LastTaskFailureMessage string `json:"lastTaskFailureMessage,omitempty"`
	// Whether guest VLAN tagging is allowed.
	GuestVlanTaggingAllowed bool `json:"guestVlanTaggingAllowed,omitempty"`
	// Whether network resources such as IP/MAC Addresses are to be retained.
	RetainNicResources bool `json:"retainNicResources,omitempty"`
	// The id of the cross vdc network if this is a stretched network, otherwise it will be null.
	CrossVdcNetworkId string `json:"crossVdcNetworkId,omitempty"`
	// The id of the org from which this network can be managed if this is a stretched network, otherwise it will be null.
	CrossVdcNetworkLocationId string `json:"crossVdcNetworkLocationId,omitempty"`
	// The number of IP addresses defined by the static ip pools. If the network contains any IpV6 subnets, the total ip count will be null.
	TotalIpCount int32 `json:"totalIpCount,omitempty"`
	// The number of IP address used from the static ip pools.
	UsedIpCount int32 `json:"usedIpCount,omitempty"`
	// Whether this network is advertised so that it can be routed out to the external networks. This applies only to network backed by NSX-T. Value will be unset if route advertisement is not applicable. 
	RouteAdvertised bool `json:"routeAdvertised,omitempty"`
	// The list of firewall groups of type SECURITY_GROUP/STATIC_MEMBERS that are assigned to the Org VDC Network. These groups can then be used in firewall rules to protect the Org VDC Network and allow/deny traffic. 
	SecurityGroups []EntityReference `json:"securityGroups,omitempty"`
}
