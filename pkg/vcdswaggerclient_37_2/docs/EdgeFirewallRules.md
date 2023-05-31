# EdgeFirewallRules

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**SystemRules** | [**[]EdgeFirewallRule**](EdgeFirewallRule.md) | The ordered list of system defined edge firewall rules. System rules are applied before user defined rules in the order in which they are returned.  | [optional] [default to null]
**UserDefinedRules** | [**[]EdgeFirewallRule**](EdgeFirewallRule.md) | The ordered list of user defined edge firewall rules. Users are allowed to add/modify/delete rules only to this list.  | [optional] [default to null]
**DefaultRules** | [**[]EdgeFirewallRule**](EdgeFirewallRule.md) | The ordered list of default edge firewall rules. Default rules are applied after the user defined rules in the order in which they are returned.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


