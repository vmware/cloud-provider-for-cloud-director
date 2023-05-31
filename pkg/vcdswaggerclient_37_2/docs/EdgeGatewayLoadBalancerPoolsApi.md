# \EdgeGatewayLoadBalancerPoolsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoadBalancerPool**](EdgeGatewayLoadBalancerPoolsApi.md#CreateLoadBalancerPool) | **Post** /1.0.0/loadBalancer/pools | Creates a Load Balancer Pool.
[**GetPoolSummariesForGateway**](EdgeGatewayLoadBalancerPoolsApi.md#GetPoolSummariesForGateway) | **Get** /1.0.0/edgeGateways/{gatewayId}/loadBalancer/poolSummaries | Retrieves all Load Balancer Pool Summaries for a given Edge Gateway.


# **CreateLoadBalancerPool**
> CreateLoadBalancerPool(ctx, loadBalancerPool)
Creates a Load Balancer Pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerPool** | [**EdgeLoadBalancerPool**](EdgeLoadBalancerPool.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPoolSummariesForGateway**
> EdgeLoadBalancerPoolSummaries GetPoolSummariesForGateway(ctx, page, pageSize, gatewayId, optional)
Retrieves all Load Balancer Pool Summaries for a given Edge Gateway.

Retrieves summaries for all of the Load Balancer Pools that are configured for an Edge Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **gatewayId** | **string**|  | 
 **optional** | ***EdgeGatewayLoadBalancerPoolsApiGetPoolSummariesForGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayLoadBalancerPoolsApiGetPoolSummariesForGatewayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EdgeLoadBalancerPoolSummaries**](EdgeLoadBalancerPoolSummaries.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

