# \EdgeGatewayLoadBalancerApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoadBalancerConfig**](EdgeGatewayLoadBalancerApi.md#GetLoadBalancerConfig) | **Get** /1.0.0/edgeGateways/{gatewayId}/loadBalancer | Retrieves Load Balancer configuration on an Edge Gateway.
[**UpdateLoadBalancerConfig**](EdgeGatewayLoadBalancerApi.md#UpdateLoadBalancerConfig) | **Put** /1.0.0/edgeGateways/{gatewayId}/loadBalancer | Update Load Balancer configuration on an Edge Gateway.


# **GetLoadBalancerConfig**
> EdgeGatewayLoadBalancerConfig GetLoadBalancerConfig(ctx, gatewayId)
Retrieves Load Balancer configuration on an Edge Gateway.

Retrieves the current state of Load Balancer service on Edge Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**EdgeGatewayLoadBalancerConfig**](EdgeGatewayLoadBalancerConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerConfig**
> UpdateLoadBalancerConfig(ctx, loadBalancerConfig, gatewayId)
Update Load Balancer configuration on an Edge Gateway.

Updates the Load Balancer service configuration on Edge Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerConfig** | [**EdgeGatewayLoadBalancerConfig**](EdgeGatewayLoadBalancerConfig.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

