# \LoadBalancerControllerApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoadBalancerController**](LoadBalancerControllerApi.md#GetLoadBalancerController) | **Get** /1.0.0/loadBalancer/controllers/{loadBalancerControllerId} | Get Load Balancer Controller
[**UnregisterLoadBalancerController**](LoadBalancerControllerApi.md#UnregisterLoadBalancerController) | **Delete** /1.0.0/loadBalancer/controllers/{loadBalancerControllerId} | Unregister the specified Load Balancer Controller.
[**UpdateLoadBalancerController**](LoadBalancerControllerApi.md#UpdateLoadBalancerController) | **Put** /1.0.0/loadBalancer/controllers/{loadBalancerControllerId} | Update specified Load Balancer Controller


# **GetLoadBalancerController**
> LoadBalancerController GetLoadBalancerController(ctx, loadBalancerControllerId)
Get Load Balancer Controller

Retrieves a specific Load Balancer Controller. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerControllerId** | **string**|  | 

### Return type

[**LoadBalancerController**](LoadBalancerController.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnregisterLoadBalancerController**
> UnregisterLoadBalancerController(ctx, loadBalancerControllerId)
Unregister the specified Load Balancer Controller.

Unregister an Load Balancer Controller. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerControllerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerController**
> UpdateLoadBalancerController(ctx, loadBalancerController, loadBalancerControllerId)
Update specified Load Balancer Controller

Update an Load Balancer Controller. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerController** | [**LoadBalancerController**](LoadBalancerController.md)|  | 
  **loadBalancerControllerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

