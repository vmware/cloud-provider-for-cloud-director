# \LoadBalancerCloudsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoadBalancerClouds**](LoadBalancerCloudsApi.md#GetLoadBalancerClouds) | **Get** /1.0.0/loadBalancer/clouds | Get all registered Load Balancer Clouds in the system.
[**RegisterLoadBalancerCloud**](LoadBalancerCloudsApi.md#RegisterLoadBalancerCloud) | **Post** /1.0.0/loadBalancer/clouds | Register a new Load Balancer Cloud.


# **GetLoadBalancerClouds**
> LoadBalancerClouds GetLoadBalancerClouds(ctx, page, pageSize, optional)
Get all registered Load Balancer Clouds in the system.

Retrieves all registered Load Balancer Clouds. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***LoadBalancerCloudsApiGetLoadBalancerCloudsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerCloudsApiGetLoadBalancerCloudsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**LoadBalancerClouds**](LoadBalancerClouds.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterLoadBalancerCloud**
> RegisterLoadBalancerCloud(ctx, loadBalancerCloud)
Register a new Load Balancer Cloud.

Register a new Load Balancer Cloud to be used with vCloud Director. If the Load Balancer Cloud is backed by NSXALB, DHCP on the NSXALB Cloud is required. vCloud Director will enable DHCP on the NSXALB Cloud if needed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerCloud** | [**LoadBalancerCloud**](LoadBalancerCloud.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

