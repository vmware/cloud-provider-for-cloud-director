# \UniversalRoutingApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUniversalEgressRouting**](UniversalRoutingApi.md#GetUniversalEgressRouting) | **Get** /1.0.0/universalRouters/{universalRouterId}/routing | Retrieves Universal Egress Points and routing configuration for a Universal Router.
[**GetUniversalRoutes**](UniversalRoutingApi.md#GetUniversalRoutes) | **Get** /1.0.0/universalRouters/{universalRouterId}/routing/routes | Retrieves routing configuration for a Universal Router.
[**SyncUniversalRoutes**](UniversalRoutingApi.md#SyncUniversalRoutes) | **Post** /1.0.0/universalRouters/{universalRouterId}/routing/routes/sync | Sync/repair the routing configuration for a Universal Router.
[**UpdateUniversalEgressRouting**](UniversalRoutingApi.md#UpdateUniversalEgressRouting) | **Put** /1.0.0/universalRouters/{universalRouterId}/routing | Updates the routing configuration using the specified egress points in the universal routes. Any egress point that does not exist will be created before updating routing. Any egress point that currently exists and is not in use by any of the specified routes will be deleted. If the new egress points for routing fail to create, routing will not be updated. 
[**UpdateUniversalRoutes**](UniversalRoutingApi.md#UpdateUniversalRoutes) | **Put** /1.0.0/universalRouters/{universalRouterId}/routing/routes | Updates routing configuration for a Universal Router.


# **GetUniversalEgressRouting**
> UniversalEgressRoutes GetUniversalEgressRouting(ctx, universalRouterId)
Retrieves Universal Egress Points and routing configuration for a Universal Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 

### Return type

[**UniversalEgressRoutes**](UniversalEgressRoutes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniversalRoutes**
> UniversalRoutes GetUniversalRoutes(ctx, universalRouterId)
Retrieves routing configuration for a Universal Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 

### Return type

[**UniversalRoutes**](UniversalRoutes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncUniversalRoutes**
> SyncUniversalRoutes(ctx, universalRouterId)
Sync/repair the routing configuration for a Universal Router.

Sync/repair the universal routes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUniversalEgressRouting**
> UpdateUniversalEgressRouting(ctx, universalEgressRoutes, universalRouterId)
Updates the routing configuration using the specified egress points in the universal routes. Any egress point that does not exist will be created before updating routing. Any egress point that currently exists and is not in use by any of the specified routes will be deleted. If the new egress points for routing fail to create, routing will not be updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalEgressRoutes** | [**UniversalEgressRoutes**](UniversalEgressRoutes.md)|  | 
  **universalRouterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUniversalRoutes**
> UpdateUniversalRoutes(ctx, routes, universalRouterId)
Updates routing configuration for a Universal Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routes** | [**UniversalRoutes**](UniversalRoutes.md)|  | 
  **universalRouterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

