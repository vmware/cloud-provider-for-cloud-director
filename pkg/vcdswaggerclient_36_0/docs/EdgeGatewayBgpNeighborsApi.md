# \EdgeGatewayBgpNeighborsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBgpNeighbor**](EdgeGatewayBgpNeighborsApi.md#CreateBgpNeighbor) | **Post** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/neighbors | Creates a new BGP neighbor for the edge gateway.
[**GetBgpNeighbors**](EdgeGatewayBgpNeighborsApi.md#GetBgpNeighbors) | **Get** /1.0.0/edgeGateways/{gatewayId}/routing/bgp/neighbors | Retrieves all BGP neighbors configured for the edge gateway.


# **CreateBgpNeighbor**
> CreateBgpNeighbor(ctx, bgpNeighbor, gatewayId)
Creates a new BGP neighbor for the edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bgpNeighbor** | [**EdgeBgpNeighbor**](EdgeBgpNeighbor.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBgpNeighbors**
> EdgeBgpNeighbors GetBgpNeighbors(ctx, gatewayId)
Retrieves all BGP neighbors configured for the edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**EdgeBgpNeighbors**](EdgeBgpNeighbors.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

