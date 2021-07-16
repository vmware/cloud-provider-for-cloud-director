/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Describes the connection between a routed network and an edge gateway. This must be set if converting from an isolated to a routed network. 
type RouterConnection struct {
	// The edge gateway that this network is attached to.
	RouterRef *EntityReference `json:"routerRef,omitempty"`
	// How the network is connected to the edge gateway. This field is updatable to allow conversions between different types. If owner is a VDC group that is backed by a NSX-V network provider, this field does not need to be set. The organization VDC network will be automatically connected to the distributed router associated with the VDC group. 
	ConnectionType *VdcNetworkConnectionType `json:"connectionType,omitempty"`
	// Whether network is marked as connected in NSX.
	Connected bool `json:"connected,omitempty"`
}
