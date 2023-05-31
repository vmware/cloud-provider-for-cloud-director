# \EdgeGatewayQoSApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGatewayQoSConfig**](EdgeGatewayQoSApi.md#GetGatewayQoSConfig) | **Get** /1.0.0/edgeGateways/{gatewayId}/qos | Retrieves the Edge Gateway Rate Limiting (QoS) configuration.
[**UpdateGatewayQoSConfig**](EdgeGatewayQoSApi.md#UpdateGatewayQoSConfig) | **Put** /1.0.0/edgeGateways/{gatewayId}/qos | Update Rate Limiting (QoS) configuration on an Edge Gateway.


# **GetGatewayQoSConfig**
> EdgeQoSConfig GetGatewayQoSConfig(ctx, gatewayId)
Retrieves the Edge Gateway Rate Limiting (QoS) configuration.

Retrieves the Rate Limiting (QoS) configuration on an Edge Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**EdgeQoSConfig**](EdgeQoSConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGatewayQoSConfig**
> UpdateGatewayQoSConfig(ctx, edgeQoSConfig, gatewayId)
Update Rate Limiting (QoS) configuration on an Edge Gateway.

Updates the Rate Limiting (QoS) configuration on an Edge Gateway. Allows configuration of rate limits for traffic passing through this gateway. This defines QoS profiles which contains configuration which can be applied in ingress and egress directions on Edge Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeQoSConfig** | [**EdgeQoSConfig**](EdgeQoSConfig.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

