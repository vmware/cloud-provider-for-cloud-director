# \EdgeGatewaysApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeGateway**](EdgeGatewaysApi.md#CreateEdgeGateway) | **Post** /1.0.0/edgeGateways | Create a new edge gateway
[**GetAllEdgeGateways**](EdgeGatewaysApi.md#GetAllEdgeGateways) | **Get** /1.0.0/edgeGateways | Get all the edge gateways


# **CreateEdgeGateway**
> CreateEdgeGateway(ctx, edgeGateway)
Create a new edge gateway

Create a new edge gateway for a vDC. Only NSX-T Edge Gateways can be created with this endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeGateway** | [**EdgeGateway**](EdgeGateway.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEdgeGateways**
> EdgeGateways GetAllEdgeGateways(ctx, page, pageSize, optional)
Get all the edge gateways

Get all edge gateways. If \"ownerRef.id\" filter is not specified, then user will see all the edge gateways in the organization they have the right to view. Results can be filtered by ownerRef. Combination of ownerRef and _context. <code>(_context==includeAccessible)</code> can be used to get all the edge gateways which are available to an Org vDC including the gateways which are owned by datacenter groups but available to Org vDC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***EdgeGatewaysApiGetAllEdgeGatewaysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewaysApiGetAllEdgeGatewaysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EdgeGateways**](EdgeGateways.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

