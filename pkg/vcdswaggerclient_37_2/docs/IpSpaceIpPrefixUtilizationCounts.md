# IpSpaceIpPrefixUtilizationCounts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **string** | The number of IP addresses or IP Prefixes defined by the IP Space. If user does not own this IP Space, this is the quota that the user&#39;s organization is granted. A \&quot;-1\&quot; value means that the user&#39;s organization has no cap on the quota; for this case, allocatedPercentage is unset.  | [optional] [default to null]
**AllocatedCount** | **string** | The number of allocated IP addresses or IP Prefixes. | [optional] [default to null]
**UsedCount** | **string** | The number of used IP addresses or IP Prefixes. An allocated IP address or IP Prefix is considered used if it is being used in network services such as NAT rule or in Org VDC network definition.  | [optional] [default to null]
**UnusedCount** | **string** | The number of unused IP addresses or IP Prefixes. An IP address or an IP Prefix is considered unused if it is allocated but not being used by any network service or any Org vDC network definition.  | [optional] [default to null]
**AllocatedPercentage** | **float32** | Specifies the percentage of allocated IP addresses or IP Prefixes out of all defined IP addresses or IP Prefixes. | [optional] [default to null]
**UsedPercentage** | **float32** | Specifies the percentage of used IP addresses or IP Prefixes out of total allocated IP addresses or IP Prefixes. | [optional] [default to null]
**PrefixLengthUtilizations** | [**[]IpSpacePrefixLengthUtilization**](IpSpacePrefixLengthUtilization.md) | Utilization summary grouped by IP Prefix&#39;s prefix length. This information will only be returned for an individual IP Prefix.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


