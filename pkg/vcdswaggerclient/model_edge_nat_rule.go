/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// The configuration for a given NAT Rule.
type EdgeNatRule struct {
	// The unique id of the NAT Rule. This must be supplied when updating a given NAT Rule. On creation, an unique id is generated for the NAT Rule. 
	Id string `json:"id,omitempty"`
	// User friendly name for the NAT Rule. Name must be provided.
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	// A flag indicating whether the individual nat rule is enabled or not. The default is true.
	Enabled bool `json:"enabled,omitempty"`
	// Represents the type of NAT Rule. SNAT translates an internal IP to an external IP and is used for outbound traffic. DNAT translates the external IP to an internal IP and is used for inbound traffic. 
	RuleType *NatRuleType `json:"ruleType"`
	// Represents the application ports on which the NAT Rule will be applied. An application port profile id in the form of URN format must be provided. If not provided then the port will be considered as \"ANY\". For a DNAT Rule, the source port on the application port profile represents the port from which the traffic is originating from. For a DNAT rule, the destination port on the application port profile represents the internal port on the workloads where the traffic is terminating. For a SNAT rule, the source port on the application port profile represents the internal port on the workloads where the traffic is originating from. For a SNAT rule, the destination port application port profile represents the port where the traffic is terminating. 
	ApplicationPortProfile *EntityReference `json:"applicationPortProfile,omitempty"`
	// The external addresses for the NAT Rule. This must be supplied as a single IP or Network CIDR. For a DNAT rule, this is the external facing IP Address for incoming traffic. For an SNAT rule, this is the external facing IP Address for outgoing traffic. These ips are typically allocated/suballocated IP Addresses on the Edge Gateway. 
	ExternalAddresses string `json:"externalAddresses"`
	// The internal addresses for the NAT Rule. This must be supplied as a single IP or Network CIDR. For a DNAT rule, this is the internal IP Address for incoming traffic. For an SNAT rule, this is the internal IP Address for outgoing traffic. These ips are typically the Private IPs that are allocated to workloads. 
	InternalAddresses string `json:"internalAddresses"`
	// Port number or port range for incoming network traffic. If Any Traffic is selected for the Service, the default internal port is \"ANY\". Note that this field has been deprecated. Please use dnatExternalPort to set port forwarding for DNAT rules. This typically should not be set for SNAT rules as the rule would not be able to support IP Translation with multiple ports. 
	InternalPort string `json:"internalPort,omitempty"`
	// This represents the external port number or port range when doing DNAT port forwarding from external to internal. The default dnatExternalPort is \"ANY\" meaning traffic on any port for the given IPs selected will be translated. 
	DnatExternalPort string `json:"dnatExternalPort,omitempty"`
	// A flag indicating whether logging for the individual nat rule is enabled or not. The default is false.
	Logging bool `json:"logging,omitempty"`
	// A flag indicating whether this NAT rule is managed by the system. This is not user editable
	SystemRule bool `json:"systemRule,omitempty"`
	// The destination addresses to match in the SNAT Rule. This must be supplied as a single IP or Network CIDR. Providing no value for this field results in match with ANY destination network. These IPs are typically the Private IPs that are allocated to destination workloads. 
	SnatDestinationAddresses string `json:"snatDestinationAddresses,omitempty"`
	// Determines how the firewall matches the address during NATing if firewall stage is not skipped.  Below are valid values. <ul>   <li> <code> MATCH_INTERNAL_ADDRESS </code> indicates the firewall will be applied to internal address of a NAT rule. For SNAT, the internal address is        the original source address before NAT is done. For DNAT, the internal address is the translated destination address after NAT is done.   <li> <code> MATCH_EXTERNAL_ADDRESS </code> indicates the firewall will be applied to external address of a NAT rule. For SNAT, the external address is        the translated source address after NAT is done. For DNAT, the external address is the original destination address before NAT is done.   <li> <code> BYPASS </code> firewall stage will be skipped. </ul> The default is MATCH_INTERNAL_ADDRESS. 
	FirewallMatch string `json:"firewallMatch,omitempty"`
	// If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a higher precedence for this rule.
	Priority int32 `json:"priority,omitempty"`
	Version *ObjectVersion `json:"version,omitempty"`
}
