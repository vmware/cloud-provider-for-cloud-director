# \LoadBalancerCloudApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoadBalancerCloud**](LoadBalancerCloudApi.md#GetLoadBalancerCloud) | **Get** /1.0.0/loadBalancer/clouds/{loadBalancerCloudId} | Get Load Balancer Cloud.
[**UnregisterLoadBalancerCloud**](LoadBalancerCloudApi.md#UnregisterLoadBalancerCloud) | **Delete** /1.0.0/loadBalancer/clouds/{loadBalancerCloudId} | Unregister the specified Load Balancer Cloud.
[**UpdateLoadBalancerCloud**](LoadBalancerCloudApi.md#UpdateLoadBalancerCloud) | **Put** /1.0.0/loadBalancer/clouds/{loadBalancerCloudId} | Update specified Load Balancer Cloud.


# **GetLoadBalancerCloud**
> LoadBalancerCloud GetLoadBalancerCloud(ctx, loadBalancerCloudId)
Get Load Balancer Cloud.

Retrieves a specific Load Balancer Cloud. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerCloudId** | **string**|  | 

### Return type

[**LoadBalancerCloud**](LoadBalancerCloud.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnregisterLoadBalancerCloud**
> UnregisterLoadBalancerCloud(ctx, loadBalancerCloudId)
Unregister the specified Load Balancer Cloud.

Unregister an Load Balancer Cloud. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerCloudId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerCloud**
> UpdateLoadBalancerCloud(ctx, loadBalancerCloud, loadBalancerCloudId)
Update specified Load Balancer Cloud.

Update an Load Balancer Cloud. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerCloud** | [**LoadBalancerCloud**](LoadBalancerCloud.md)|  | 
  **loadBalancerCloudId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

