# \NetworkPoolsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkPool**](NetworkPoolsApi.md#CreateNetworkPool) | **Post** /1.0.0/networkPools | Create a new network pool.
[**GetNetworkPoolsSummary**](NetworkPoolsApi.md#GetNetworkPoolsSummary) | **Get** /1.0.0/networkPools/networkPoolSummaries | Get summary of all the Network Pools in the system.


# **CreateNetworkPool**
> CreateNetworkPool(ctx, networkPool)
Create a new network pool.

Create a network pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkPool** | [**NetworkPool**](NetworkPool.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkPoolsSummary**
> NetworkPoolsSummary GetNetworkPoolsSummary(ctx, page, pageSize, optional)
Get summary of all the Network Pools in the system.

Retrieves summary of all Network Pools in the system. Results can be filtered by context `(_context)`. Supported contexts are: <ul> <li>Provider vDC ID <code>(_context==providerVdcId)</code> - Returns all the network pools which are available to a specific Provider vDC. <li>Org vDC ID <code>(_context==orgVdcId)</code> - Returns all the network pools which are available to a specific Org vDC. <li>Virtual Center ID <code>(managingOwnerRef.id==vcId)</code> + Resource Pool Moref <code>(_context==moref)</code> - Returns all the network pools which are related to a specific Resoure Pool. </ul> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NetworkPoolsApiGetNetworkPoolsSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkPoolsApiGetNetworkPoolsSummaryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**NetworkPoolsSummary**](NetworkPoolsSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

