# RouteAdvertisement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | True means that the subnets will be advertised. The default is true. | [optional] [default to null]
**Subnets** | **[]string** | List of subnets that will be advertised so that the Edge Gateway can route out to the connected external network. Each value is in CIDR format. Note that the CIDR value will automatically be converted to its network definition based on the prefix length.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


