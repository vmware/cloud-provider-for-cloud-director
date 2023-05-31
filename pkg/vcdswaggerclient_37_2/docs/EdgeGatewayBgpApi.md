# \EdgeGatewayBgpApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBgpConfig**](EdgeGatewayBgpApi.md#GetBgpConfig) | **Get** /1.0.0/edgeGateways/{gatewayId}/routing/bgp | Retrieves the BGP configuration for a given Edge Gateway.
[**UpdateBgpConfig**](EdgeGatewayBgpApi.md#UpdateBgpConfig) | **Put** /1.0.0/edgeGateways/{gatewayId}/routing/bgp | Updates the BGP configuration on the Edge Gateway.


# **GetBgpConfig**
> EdgeBgpConfig GetBgpConfig(ctx, gatewayId)
Retrieves the BGP configuration for a given Edge Gateway.

Retrieves the general BGP configuration for an edge gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**EdgeBgpConfig**](EdgeBgpConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBgpConfig**
> UpdateBgpConfig(ctx, bgpConfig, gatewayId)
Updates the BGP configuration on the Edge Gateway.

Updates the general BGP configuration on an edge gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bgpConfig** | [**EdgeBgpConfig**](EdgeBgpConfig.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

