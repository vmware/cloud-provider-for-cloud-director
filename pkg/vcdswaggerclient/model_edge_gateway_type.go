/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// EdgeGatewayType : The type of the edge gateway.
type EdgeGatewayType interface{}

// List of EdgeGatewayType
const (
	NSXV_BACKED_EdgeGatewayType   = "NSXV_BACKED"
	NSXT_BACKED_EdgeGatewayType   = "NSXT_BACKED"
	NSXT_IMPORTED_EdgeGatewayType = "NSXT_IMPORTED"
)
