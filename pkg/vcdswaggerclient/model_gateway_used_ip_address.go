/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Describes an IP address currently be used by the edge gateway. The IP belongs to the IPs allocated to the edge. 
type GatewayUsedIpAddress struct {
	// The external network that this IP Address belongs to.
	NetworkRef *EntityReference `json:"networkRef,omitempty"`
	// The IP address in used.
	IpAddress string `json:"ipAddress,omitempty"`
	// The catagory that an IP can be used for.
	Category *GatewayUsedIpAddressCategory `json:"category,omitempty"`
}
