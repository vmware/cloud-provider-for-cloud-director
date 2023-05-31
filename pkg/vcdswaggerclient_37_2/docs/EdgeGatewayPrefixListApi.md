# \EdgeGatewayPrefixListApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePrefixList**](EdgeGatewayPrefixListApi.md#DeletePrefixList) | **Delete** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/prefixLists/{listId} | Deletes a specific Prefix list for a given edge gateway.
[**GetPrefixList**](EdgeGatewayPrefixListApi.md#GetPrefixList) | **Get** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/prefixLists/{listId} | Retrieves a specific Prefix list for a given edge gateway.
[**UpdatePrefixList**](EdgeGatewayPrefixListApi.md#UpdatePrefixList) | **Put** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/prefixLists/{listId} | Updates a specific Prefix list for a given edge gateway.


# **DeletePrefixList**
> DeletePrefixList(ctx, gatewayId, listId)
Deletes a specific Prefix list for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **listId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrefixList**
> EdgePrefixList GetPrefixList(ctx, gatewayId, listId)
Retrieves a specific Prefix list for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **listId** | **string**|  | 

### Return type

[**EdgePrefixList**](EdgePrefixList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrefixList**
> UpdatePrefixList(ctx, prefixList, gatewayId, listId)
Updates a specific Prefix list for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **prefixList** | [**EdgePrefixList**](EdgePrefixList.md)|  | 
  **gatewayId** | **string**|  | 
  **listId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

