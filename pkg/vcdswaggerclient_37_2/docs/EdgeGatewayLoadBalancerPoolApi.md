# \EdgeGatewayLoadBalancerPoolApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLoadBalancerPool**](EdgeGatewayLoadBalancerPoolApi.md#DeleteLoadBalancerPool) | **Delete** /1.0.0/loadBalancer/pools/{poolId} | Deletes a specific Load Balancer Pool.
[**GetLoadBalancerPool**](EdgeGatewayLoadBalancerPoolApi.md#GetLoadBalancerPool) | **Get** /1.0.0/loadBalancer/pools/{poolId} | Retrieves a specific Load Balancer Pool.
[**UpdateLoadBalancerPool**](EdgeGatewayLoadBalancerPoolApi.md#UpdateLoadBalancerPool) | **Put** /1.0.0/loadBalancer/pools/{poolId} | Updates a specific Load Balancer Pool.


# **DeleteLoadBalancerPool**
> DeleteLoadBalancerPool(ctx, poolId)
Deletes a specific Load Balancer Pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerPool**
> EdgeLoadBalancerPool GetLoadBalancerPool(ctx, poolId)
Retrieves a specific Load Balancer Pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

[**EdgeLoadBalancerPool**](EdgeLoadBalancerPool.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerPool**
> UpdateLoadBalancerPool(ctx, loadBalancerPool, poolId)
Updates a specific Load Balancer Pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerPool** | [**EdgeLoadBalancerPool**](EdgeLoadBalancerPool.md)|  | 
  **poolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

