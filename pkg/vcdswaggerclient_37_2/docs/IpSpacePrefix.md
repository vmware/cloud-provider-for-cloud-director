# IpSpacePrefix

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the IP Space Prefix. For new IP Prefix, this is unset. Existing IP Prefix will have an ID which is used for update identification. | [optional] [default to null]
**StartingPrefixIpAddress** | **string** | Starting IP address for the IP prefix. Note that if the IP is a host IP and not the network definition IP for the specific prefix length, VCD will automatically modify starting IP to the network definition&#39;s IP for the specified host IP. An example is that for prefix length 30, the starting IP of 192.169.0.2 will be automatically modified to 192.169.0.0.  192.169.0.6 will be modified to 192.169.0.4. 192.169.0.0/30 and 192.169.0.4/30 are network definition CIDRs for host IPs 192.169.0.2 and 192.169.0.6, respectively.  | [default to null]
**PrefixLength** | **int32** | The prefix length. | [default to null]
**TotalPrefixCount** | **int32** | The number of prefix blocks defined by this IP prefix. | [default to null]
**AllocatedPrefixCount** | **int32** | The number of allocated IP prefix blocks. | [optional] [default to null]
**AllocatedPrefixPercentage** | **float32** | Specifies the percentage of allocated IP prefix blocks out of total specified IP prefix blocks. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


