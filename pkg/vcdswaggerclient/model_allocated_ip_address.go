/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Describes an allocated IP address for a network. IP addresses can be allocated to a vApp VM, an edge gateway interface or may be used in a NAT routed environment. If the address is allocated to an edge gateway, the name of the entity will be 'Edge Gateway'. 
type AllocatedIpAddress struct {
	// Unique identifier for record. This is to support UI operations.
	Id string `json:"id,omitempty"`
	// Id of the entity to which the IP address is allocated such as a VM.
	EntityId string `json:"entityId,omitempty"`
	// Name of the entity to which the IP address is allocated.
	EntityName string `json:"entityName,omitempty"`
	// Name of the vApp whose VM is using allocated IP address.
	VAppName string `json:"vAppName,omitempty"`
	// The allocated IP address.
	IpAddress string `json:"ipAddress,omitempty"`
	// Whether the entity using this IP address is currently deployed.
	Deployed bool `json:"deployed,omitempty"`
	// The possible allocation types for allocated network address like VSM allocated, VM allocated or NAT routed.
	AllocationType *AllocatedIpAddressAllocationType `json:"allocationType,omitempty"`
	// The vApp network or Org vDC network to which this entity is connected. 
	NetworkRef *EntityReference `json:"networkRef,omitempty"`
	// The organization that this entity belongs to. 
	OrgRef *EntityReference `json:"orgRef,omitempty"`
}
