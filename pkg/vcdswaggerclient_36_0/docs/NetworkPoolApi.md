# \NetworkPoolApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetworkPool**](NetworkPoolApi.md#DeleteNetworkPool) | **Delete** /1.0.0/networkPools/{networkPoolId} | Deletes a specific network pool.
[**GetNetworkPool**](NetworkPoolApi.md#GetNetworkPool) | **Get** /1.0.0/networkPools/{networkPoolId} | Retrieves a specific Network Pool.
[**SyncNetworkPool**](NetworkPoolApi.md#SyncNetworkPool) | **Post** /1.0.0/networkPools/{networkPoolId}/sync | Synchronize the VXLAN network pool.
[**UpdateNetworkPool**](NetworkPoolApi.md#UpdateNetworkPool) | **Put** /1.0.0/networkPools/{networkPoolId} | Updates a specific network pool.


# **DeleteNetworkPool**
> DeleteNetworkPool(ctx, networkPoolId)
Deletes a specific network pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkPoolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkPool**
> NetworkPool GetNetworkPool(ctx, networkPoolId)
Retrieves a specific Network Pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkPoolId** | **string**|  | 

### Return type

[**NetworkPool**](NetworkPool.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncNetworkPool**
> SyncNetworkPool(ctx, networkPoolId)
Synchronize the VXLAN network pool.

Synchronize the VXLAN network pool. If the user changes a transport zone in NSX by adding or removing clusters, synchronizing the VXLAN network pool ensures that the defined scope of the network pool is reflected in the scope of its corresponding transport zone. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkPoolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkPool**
> UpdateNetworkPool(ctx, networkPool, networkPoolId)
Updates a specific network pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkPool** | [**NetworkPool**](NetworkPool.md)|  | 
  **networkPoolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

