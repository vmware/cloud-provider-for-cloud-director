# \EdgeGatewayBgpNeighborApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBgpNeighbor**](EdgeGatewayBgpNeighborApi.md#DeleteBgpNeighbor) | **Delete** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/neighbors/{neighborId} | Deletes a specific BGP neighbor of edge gateway.
[**GetBgpNeighbor**](EdgeGatewayBgpNeighborApi.md#GetBgpNeighbor) | **Get** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/neighbors/{neighborId} | Retrieves a specific BGP neighbor of edge gateway.
[**GetBgpNeighborStatus**](EdgeGatewayBgpNeighborApi.md#GetBgpNeighborStatus) | **Get** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/neighbors/{neighborId}/status | Retrieves status of a specific BGP neighbor configured on an Edge Gateway.
[**UpdateBgpNeighbor**](EdgeGatewayBgpNeighborApi.md#UpdateBgpNeighbor) | **Put** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/neighbors/{neighborId} | Updates a specific BGP neighbor of edge gateway.


# **DeleteBgpNeighbor**
> DeleteBgpNeighbor(ctx, gatewayId, neighborId)
Deletes a specific BGP neighbor of edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **neighborId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBgpNeighbor**
> EdgeBgpNeighbor GetBgpNeighbor(ctx, gatewayId, neighborId)
Retrieves a specific BGP neighbor of edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **neighborId** | **string**|  | 

### Return type

[**EdgeBgpNeighbor**](EdgeBgpNeighbor.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBgpNeighborStatus**
> NetworkingObjectStatus GetBgpNeighborStatus(ctx, gatewayId, neighborId)
Retrieves status of a specific BGP neighbor configured on an Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **neighborId** | **string**|  | 

### Return type

[**NetworkingObjectStatus**](NetworkingObjectStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBgpNeighbor**
> UpdateBgpNeighbor(ctx, bgpNeighbor, gatewayId, neighborId)
Updates a specific BGP neighbor of edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bgpNeighbor** | [**EdgeBgpNeighbor**](EdgeBgpNeighbor.md)|  | 
  **gatewayId** | **string**|  | 
  **neighborId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

