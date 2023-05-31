# LocalEndpoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalId** | **string** | The local identifier for the endpoint. | [optional] [default to null]
**LocalAddress** | **string** | The IPV4 Address for the endpoint. This has to be a suballocated IP on the Edge Gateway. This is required. | [default to null]
**LocalNetworks** | **[]string** | List of local networks. These must be specified in normal Network CIDR format. Specifying no value is interpreted as 0.0.0.0/0. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


