# FirewallRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of this firewall rule. If a rule with the ruleId is not already present, a new rule will be created. If it already exists, the rule will be updated.  | [optional] [default to null]
**Name** | **string** | Name for the rule. | [default to null]
**Description** | **string** |  | [optional] [default to null]
**SourceFirewallGroups** | [**[]EntityReference**](EntityReference.md) | List of source groups for firewall rule. It specifies the sources of network traffic for the firewall rule. Null value or an empty list will be treated as \&quot;ANY\&quot; which means traffic from any source. For Distributed Firewall rules, an entry with an id of urn:vcloud:firewallGroup:internal can be used to specify all internal vDC Group network traffic.  | [optional] [default to null]
**DestinationFirewallGroups** | [**[]EntityReference**](EntityReference.md) | List of source groups for firewall rule. It specifies the destinations of network traffic for the firewall rule. Null value or an empty list will be treated as \&quot;ANY\&quot; which means traffic to any destination. For Distributed Firewall rules, an entry with an id of urn:vcloud:firewallGroup:internal can be used to specify all internal vDC Group network traffic.  | [optional] [default to null]
**ApplicationPortProfiles** | [**[]EntityReference**](EntityReference.md) | The list of application ports where this firewall rule is applicable. Null value or an empty list will be treated as \&quot;ANY\&quot; which means rule applies to all ports.  | [optional] [default to null]
**IpProtocol** | [***FirewallRuleIpProtocol**](FirewallRuleIpProtocol.md) | Type of IP packet that should be matched while enforcing the rule. Default value is IPV4_IPV6.  | [optional] [default to null]
**Action** | [***FirewallRuleAction**](FirewallRuleAction.md) | The action to be applied to all the traffic that meets the firewall rule criteria. It determines if the rule permits or blocks traffic. This property is now deprecated and replaced with actionValue. Property is required if actionValue is not set.  | [optional] [default to null]
**ActionValue** | **string** | The action to be applied to all the traffic that meets the firewall rule criteria. It determines if the rule permits or blocks traffic. Property is required if action is not set. Below are valid values. &lt;ul&gt;   &lt;li&gt; &lt;code&gt; ALLOW &lt;/code&gt; permits traffic to go through the firewall.   &lt;li&gt; &lt;code&gt; DROP &lt;/code&gt; blocks the traffic at the firewall. No response is sent back to the source.   &lt;li&gt; &lt;code&gt; REJECT &lt;/code&gt; blocks the traffic at the firewall. A response is sent back to the source. &lt;/ul&gt;  | [optional] [default to null]
**Direction** | [***FirewallRuleDirection**](FirewallRuleDirection.md) | Specifies the direction of the network traffic. Default value is IN_OUT.  | [optional] [default to null]
**Logging** | **bool** | Whether packet logging is enabled for firewall rule. | [optional] [default to null]
**NetworkContextProfiles** | [**[]EntityReference**](EntityReference.md) | The list of layer 7 network context profiles where this firewall rule is applicable. Null value or an empty list will be treated as \&quot;ANY\&quot; which means rule applies to all applications and domains.  | [optional] [default to null]
**Enabled** | **bool** | Whether the firewall rule is enabled. | [optional] [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


