/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// A network subnet defined by its gateway, 
type Subnet struct {
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
}
