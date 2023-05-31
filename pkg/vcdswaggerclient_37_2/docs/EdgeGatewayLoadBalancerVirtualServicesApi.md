# \EdgeGatewayLoadBalancerVirtualServicesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVirtualService**](EdgeGatewayLoadBalancerVirtualServicesApi.md#CreateVirtualService) | **Post** /1.0.0/loadBalancer/virtualServices | Create a new Virtual Service for a specific Edge Gateway.
[**GetVirtualServiceSummariesForGateway**](EdgeGatewayLoadBalancerVirtualServicesApi.md#GetVirtualServiceSummariesForGateway) | **Get** /1.0.0/edgeGateways/{gatewayId}/loadBalancer/virtualServiceSummaries | Get all Virtual Service Summaries for an Edge Gateway.


# **CreateVirtualService**
> CreateVirtualService(ctx, virtualServiceConfig)
Create a new Virtual Service for a specific Edge Gateway.

Create a new Virtual Service for a specific Edge Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualServiceConfig** | [**EdgeLoadBalancerVirtualService**](EdgeLoadBalancerVirtualService.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVirtualServiceSummariesForGateway**
> EdgeLoadBalancerVirtualServiceSummaries GetVirtualServiceSummariesForGateway(ctx, page, pageSize, gatewayId, optional)
Get all Virtual Service Summaries for an Edge Gateway.

Retrieves all Virtual Service Summaries for an Edge Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **gatewayId** | **string**|  | 
 **optional** | ***EdgeGatewayLoadBalancerVirtualServicesApiGetVirtualServiceSummariesForGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayLoadBalancerVirtualServicesApiGetVirtualServiceSummariesForGatewayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EdgeLoadBalancerVirtualServiceSummaries**](EdgeLoadBalancerVirtualServiceSummaries.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

