# EdgePrefixListEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | The network prefix in CIDR format. If the value is not specified, it will be treated as \&quot;ANY\&quot; which means match all networks. Both IPv4 and IPv6 formats are supported.  | [optional] [default to null]
**Action** | **string** | Action for the prefix list. This specifies whether the packet from specified network is advertised or not for routing purposes.  | [optional] [default to null]
**GreaterThanEqualTo** | **int32** | The value which the prefix length must be greater than or equal to. Must be less than or equal to &#39;lessThanEqualTo&#39; | [optional] [default to null]
**LessThanEqualTo** | **int32** | The value which the prefix length must be less than or equal to. Must be greater than or equal to &#39;greaterThanEqualTo&#39; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


