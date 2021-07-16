/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// GatewayUsedIpAddressCategory : The categories that an IP can be used for.
type GatewayUsedIpAddressCategory interface{}

// List of GatewayUsedIpAddressCategory
const (
	SNAT_GatewayUsedIpAddressCategory          = "SNAT"
	DNAT_GatewayUsedIpAddressCategory          = "DNAT"
	LOAD_BALANCER_GatewayUsedIpAddressCategory = "Load_Balancer"
	I_PSEC_VPN_GatewayUsedIpAddressCategory    = "IPsec_VPN"
	SSL_VPN_GatewayUsedIpAddressCategory       = "SSL_VPN"
	L2_VPN_GatewayUsedIpAddressCategory        = "L2_VPN"
	PRIMARY_IP_GatewayUsedIpAddressCategory    = "PRIMARY_IP"
)
