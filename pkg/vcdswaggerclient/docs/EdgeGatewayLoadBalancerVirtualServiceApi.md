# \EdgeGatewayLoadBalancerVirtualServiceApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVirtualService**](EdgeGatewayLoadBalancerVirtualServiceApi.md#DeleteVirtualService) | **Delete** /1.0.0/loadBalancer/virtualServices/{virtualServiceId} | Delete the specified Virtual Service.
[**GetVirtualService**](EdgeGatewayLoadBalancerVirtualServiceApi.md#GetVirtualService) | **Get** /1.0.0/loadBalancer/virtualServices/{virtualServiceId} | Get Virtual Service.
[**UpdateVirtualService**](EdgeGatewayLoadBalancerVirtualServiceApi.md#UpdateVirtualService) | **Put** /1.0.0/loadBalancer/virtualServices/{virtualServiceId} | Update specified Virtual Service.


# **DeleteVirtualService**
> DeleteVirtualService(ctx, virtualServiceId)
Delete the specified Virtual Service.

Delete a Virtual Service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualServiceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVirtualService**
> EdgeLoadBalancerVirtualService GetVirtualService(ctx, virtualServiceId)
Get Virtual Service.

Retrieves a specific Virtual Service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualServiceId** | **string**|  | 

### Return type

[**EdgeLoadBalancerVirtualService**](EdgeLoadBalancerVirtualService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVirtualService**
> UpdateVirtualService(ctx, virtualServiceConfig, virtualServiceId)
Update specified Virtual Service.

Update a Virtual Service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualServiceConfig** | [**EdgeLoadBalancerVirtualService**](EdgeLoadBalancerVirtualService.md)|  | 
  **virtualServiceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

