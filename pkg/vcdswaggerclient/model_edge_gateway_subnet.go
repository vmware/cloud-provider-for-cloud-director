/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Subnet for edge gateway that contains the IPs from the external network that are allocated to the edge. 
type EdgeGatewaySubnet struct {
	// The gateway for the subnet.
	Gateway string `json:"gateway"`
	// The prefix length of the subnet.
	PrefixLength int32 `json:"prefixLength"`
	// The DNS suffix that VMs attached to this network will use.
	DnsSuffix string `json:"dnsSuffix,omitempty"`
	// The first DNS server that VMs attached to this network will use.
	DnsServer1 string `json:"dnsServer1,omitempty"`
	// The second DNS server that VMs attached to this network will use.
	DnsServer2 string `json:"dnsServer2,omitempty"`
	// Range of IPs within the subnet that can be used in this network. A VM attached to this network is assigned one of these IPs.
	IpRanges *IpRanges `json:"ipRanges,omitempty"`
	// Indicates whether the external network subnet is currently enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The number of IP addresses defined by the static IP ranges.
	TotalIpCount int32 `json:"totalIpCount,omitempty"`
	// The number of IP address used from the static IP ranges.
	UsedIpCount int32 `json:"usedIpCount,omitempty"`
	// The primary IP address allocated for this subnet. If not specified, this IP is auto-allocated.  This IP belongs to the external network and can be used for system-configured NAT rules such as DNS forwarder configuration. 
	PrimaryIp string `json:"primaryIp,omitempty"`
	// Used for create and update api calls. If set to true, IP Ranges are automatically generated based on totalIpCount.
	AutoAllocateIpRanges bool `json:"autoAllocateIpRanges,omitempty"`
}
