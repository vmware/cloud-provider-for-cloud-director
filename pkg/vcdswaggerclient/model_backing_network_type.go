/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// BackingNetworkType : The backing network type.
type BackingNetworkType interface{}

// List of BackingNetworkType
const (
	DV_PORTGROUP_BackingNetworkType              = "DV_PORTGROUP"
	VIRTUAL_WIRE_BackingNetworkType              = "VIRTUAL_WIRE"
	IMPORTED_T_LOGICAL_SWITCH_BackingNetworkType = "IMPORTED_T_LOGICAL_SWITCH"
	OPAQUE_NETWORK_BackingNetworkType            = "OPAQUE_NETWORK"
	NSXT_FIXED_SEGMENT_BackingNetworkType        = "NSXT_FIXED_SEGMENT"
	NSXT_FLEXIBLE_SEGMENT_BackingNetworkType     = "NSXT_FLEXIBLE_SEGMENT"
)
