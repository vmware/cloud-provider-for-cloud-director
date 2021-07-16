/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// VdcNetworkFenceType : The network type.
type VdcNetworkFenceType string

// List of VdcNetworkFenceType
const (
	DIRECT_VdcNetworkFenceType        VdcNetworkFenceType = "DIRECT"
	DIRECT_UPLINK_VdcNetworkFenceType VdcNetworkFenceType = "DIRECT_UPLINK"
	ISOLATED_VdcNetworkFenceType      VdcNetworkFenceType = "ISOLATED"
	OPAQUE_VdcNetworkFenceType        VdcNetworkFenceType = "OPAQUE"
	NAT_ROUTED_VdcNetworkFenceType    VdcNetworkFenceType = "NAT_ROUTED"
	CROSS_VDC_VdcNetworkFenceType     VdcNetworkFenceType = "CROSS_VDC"
)
