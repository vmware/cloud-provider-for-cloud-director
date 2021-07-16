/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// VdcNetworkConnectionType : Describes how a network is connected to a gateway.
type VdcNetworkConnectionType string

// List of VdcNetworkConnectionType
const (
	DISTRIBUTED_VdcNetworkConnectionType  VdcNetworkConnectionType = "DISTRIBUTED"
	INTERNAL_VdcNetworkConnectionType     VdcNetworkConnectionType = "INTERNAL"
	SUBINTERFACE_VdcNetworkConnectionType VdcNetworkConnectionType = "SUBINTERFACE"
)
