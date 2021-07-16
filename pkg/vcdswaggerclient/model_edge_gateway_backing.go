/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// The backing details of the edge gateway 
type EdgeGatewayBacking struct {
	// The identifier of the edge gateways backing router.
	BackingId string `json:"backingId,omitempty"`
	// The type of the gateway. Describes if this is an NSX-T edge gateway or an NSX-V edge gateway.
	GatewayType *EdgeGatewayType `json:"gatewayType,omitempty"`
	// The backing network provider, either NSX-T or NSX-V.
	NetworkProvider *EntityReference `json:"networkProvider,omitempty"`
}
