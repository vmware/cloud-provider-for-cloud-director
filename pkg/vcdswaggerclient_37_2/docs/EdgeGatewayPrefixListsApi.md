# \EdgeGatewayPrefixListsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrefixList**](EdgeGatewayPrefixListsApi.md#CreatePrefixList) | **Post** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/prefixLists | Creates a new Prefix list on the edge gateway.
[**GetPrefixLists**](EdgeGatewayPrefixListsApi.md#GetPrefixLists) | **Get** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/prefixLists | Retrieves all Prefix lists for a given edge gateway.


# **CreatePrefixList**
> CreatePrefixList(ctx, prefixList, gatewayId)
Creates a new Prefix list on the edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **prefixList** | [**EdgePrefixList**](EdgePrefixList.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrefixLists**
> EdgePrefixLists GetPrefixLists(ctx, gatewayId, optional)
Retrieves all Prefix lists for a given edge gateway.

Retrieves all Prefix lists for a given edge gateway. Results can be sorted by only a single parameter. Sorting by combination of parameters (sortAsc=foo&sortDesc=bar) is not allowed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
 **optional** | ***EdgeGatewayPrefixListsApiGetPrefixListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayPrefixListsApiGetPrefixListsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EdgePrefixLists**](EdgePrefixLists.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

