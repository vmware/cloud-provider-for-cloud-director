# IpSpacePrefixes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpPrefixSequence** | [**[]IpSpacePrefix**](IpSpacePrefix.md) | A sequence of IP prefixes with same prefix length. All the IP Prefix sequences with the same prefix length are treated as one logical unit for allocation purpose.  | [optional] [default to null]
**DefaultQuotaForPrefixLength** | **int32** | This specifies the number of prefixes from the specified sequence which can be consumed by each organization using this IP Space. All the IP Prefix sequences with the same prefix length are treated as one logical unit for allocation purpose. This is typically set for IP Spaces with type PUBLIC or SHARED_SERVICES. A Quota of -1 means there is no cap to the number of IP Prefixes that can be allocated. A Quota of 0 means that the IP Prefixes cannot be allocated. If not specified, all PUBLIC or SHARED_SERVICES IP Spaces have a default quota of 0 for IP Prefixes and all PRIVATE IP Spaces have a default quota of -1 for IP Prefixes.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


