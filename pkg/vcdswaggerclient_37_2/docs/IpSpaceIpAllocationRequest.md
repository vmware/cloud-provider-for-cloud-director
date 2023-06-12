# IpSpaceIpAllocationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | The type of the IP allocation requested. Possible values are: &lt;ul&gt; &lt;li&gt; FLOATING_IP - For allocation of floating IP addresses from defined IP Space ranges. &lt;li&gt; IP_PREFIX - For allocation of IP prefix sequences from defined IP Space prefixes. &lt;/ul&gt;  | [default to null]
**Quantity** | **int32** | The number of IP addresses or IP Prefix blocks to allocate. Specifying quantity will allocate the given number of any free IP addresses or IP Prefixes within the IP Space. To use a specific IP address or IP Prefix, please use the value field to request a specific value.  | [optional] [default to null]
**PrefixLength** | **int32** | The prefix length of an IP Prefix to allocate. This is required if type is &lt;em&gt;IP_PREFIX&lt;/em&gt;. This field is only required if the request is for a specific quantity of IP Prefixes and not needed if request value is specified.  | [optional] [default to null]
**Value** | **string** | The specific IP address or IP Prefix to allocate. If an IP address or IP Prefix is specified, the quantity value should not be set.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


