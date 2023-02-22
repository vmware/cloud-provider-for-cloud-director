# AvailableIpPoolSubnet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | **string** | The gateway for the subnet on which IPs are available. | [optional] [default to null]
**PrefixLength** | **int32** | The netmask prefix length of the subnet. | [optional] [default to null]
**Enabled** | **bool** | Indicates whether the network subnet is currently enabled. | [optional] [default to null]
**IpRanges** | [***IpRanges**](IpRanges.md) | Range of IPs which are available for use. | [optional] [default to null]
**TotalIpCount** | **int32** | The total number of available IP addresses. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


