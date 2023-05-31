# IpSpaceRanges

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRanges** | [**[]IpSpaceRange**](IpSpaceRange.md) | List of IP Ranges. | [optional] [default to null]
**DefaultFloatingIpQuota** | **int32** | This specifies the default number of IPs from the specified ranges which can be consumed by each organization using this IP Space. This is typically set for IP Spaces with type PUBLIC or SHARED_SERVICES. A Quota of -1 means there is no cap to the number of IP addresses that can be allocated. A Quota of 0 means that the IP addresses cannot be allocated. If not specified, all PUBLIC or SHARED_SERVICES IP Spaces have a default quota of 1 for Floating IP addresses and all PRIVATE IP Spaces have a default quota of -1 for Floating IP addresses.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


