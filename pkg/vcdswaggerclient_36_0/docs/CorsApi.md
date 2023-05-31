# \CorsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryAllowedOrigins**](CorsApi.md#QueryAllowedOrigins) | **Get** /1.0.0/site/settings/cors | Queries CORS allowed origins.
[**SetAllowedOrigins**](CorsApi.md#SetAllowedOrigins) | **Put** /1.0.0/site/settings/cors | Sets allowed origins to the given set of origins.


# **QueryAllowedOrigins**
> AllowedOrigins QueryAllowedOrigins(ctx, page, pageSize, optional)
Queries CORS allowed origins.

Queries the set of allowed origins. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***CorsApiQueryAllowedOriginsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorsApiQueryAllowedOriginsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**AllowedOrigins**](AllowedOrigins.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAllowedOrigins**
> AllowedOrigins SetAllowedOrigins(ctx, allowedOrigins)
Sets allowed origins to the given set of origins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **allowedOrigins** | [**AllowedOrigins**](AllowedOrigins.md)|  | 

### Return type

[**AllowedOrigins**](AllowedOrigins.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

