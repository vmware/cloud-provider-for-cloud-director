# \EdgeGatewayStaticRoutesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStaticRoute**](EdgeGatewayStaticRoutesApi.md#CreateStaticRoute) | **Post** /1.0.0/edgeGateways/{gatewayId}/routing/staticRoutes | Creates a new static route configured on an Edge Gateway.
[**DeleteStaticRoute**](EdgeGatewayStaticRoutesApi.md#DeleteStaticRoute) | **Delete** /1.0.0/edgeGateways/{gatewayId}/routing/staticRoutes/{routeId} | Deletes a specific static route of edge gateway.
[**GetStaticRoute**](EdgeGatewayStaticRoutesApi.md#GetStaticRoute) | **Get** /1.0.0/edgeGateways/{gatewayId}/routing/staticRoutes/{routeId} | Retrieves a specific static route configured on an Edge Gateway.
[**GetStaticRoutes**](EdgeGatewayStaticRoutesApi.md#GetStaticRoutes) | **Get** /1.0.0/edgeGateways/{gatewayId}/routing/staticRoutes | Retrieves all static routes configured for the edge gateway.
[**UpdateStaticRoute**](EdgeGatewayStaticRoutesApi.md#UpdateStaticRoute) | **Put** /1.0.0/edgeGateways/{gatewayId}/routing/staticRoutes/{routeId} | Updates a specific static route configured on an Edge Gateway.


# **CreateStaticRoute**
> CreateStaticRoute(ctx, staticRoute, gatewayId)
Creates a new static route configured on an Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **staticRoute** | [**EdgeStaticRoute**](EdgeStaticRoute.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStaticRoute**
> DeleteStaticRoute(ctx, gatewayId, routeId)
Deletes a specific static route of edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **routeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStaticRoute**
> EdgeStaticRoute GetStaticRoute(ctx, gatewayId, routeId)
Retrieves a specific static route configured on an Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **routeId** | **string**|  | 

### Return type

[**EdgeStaticRoute**](EdgeStaticRoute.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStaticRoutes**
> EdgeStaticRoutes GetStaticRoutes(ctx, pageSize, gatewayId, optional)
Retrieves all static routes configured for the edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **gatewayId** | **string**|  | 
 **optional** | ***EdgeGatewayStaticRoutesApiGetStaticRoutesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayStaticRoutesApiGetStaticRoutesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EdgeStaticRoutes**](EdgeStaticRoutes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStaticRoute**
> UpdateStaticRoute(ctx, staticRoute, gatewayId, routeId)
Updates a specific static route configured on an Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **staticRoute** | [**EdgeStaticRoute**](EdgeStaticRoute.md)|  | 
  **gatewayId** | **string**|  | 
  **routeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

