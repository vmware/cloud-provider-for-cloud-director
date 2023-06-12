# \LoadBalancerControllersApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoadBalancerControllers**](LoadBalancerControllersApi.md#GetLoadBalancerControllers) | **Get** /1.0.0/loadBalancer/controllers | Get all registered Load Balancer Controllers in the system.
[**RegisterLoadBalancerController**](LoadBalancerControllersApi.md#RegisterLoadBalancerController) | **Post** /1.0.0/loadBalancer/controllers | Register a new Load Balancer Controller


# **GetLoadBalancerControllers**
> LoadBalancerControllers GetLoadBalancerControllers(ctx, page, pageSize, optional)
Get all registered Load Balancer Controllers in the system.

Retrieves all registered Load Balancer Controllers. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***LoadBalancerControllersApiGetLoadBalancerControllersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerControllersApiGetLoadBalancerControllersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**LoadBalancerControllers**](LoadBalancerControllers.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterLoadBalancerController**
> RegisterLoadBalancerController(ctx, loadBalancerController)
Register a new Load Balancer Controller

Register a new Load Balancer Controller to be used with vCloud Director. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerController** | [**LoadBalancerController**](LoadBalancerController.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

