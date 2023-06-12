# \UniversalRouterApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUniversalRouter**](UniversalRouterApi.md#DeleteUniversalRouter) | **Delete** /1.0.0/universalRouters/{universalRouterId} | Deletes a specific Universal Router
[**GetUniversalRouter**](UniversalRouterApi.md#GetUniversalRouter) | **Get** /1.0.0/universalRouters/{universalRouterId} | Retrieves a specific Universal Router
[**SyncUniversalRouter**](UniversalRouterApi.md#SyncUniversalRouter) | **Post** /1.0.0/universalRouters/{universalRouterId}/sync | Sync/repair the Universal Router
[**UpdateUniversalRouter**](UniversalRouterApi.md#UpdateUniversalRouter) | **Put** /1.0.0/universalRouters/{universalRouterId} | Updates a specific Universal Router


# **DeleteUniversalRouter**
> DeleteUniversalRouter(ctx, universalRouterId, optional)
Deletes a specific Universal Router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 
 **optional** | ***UniversalRouterApiDeleteUniversalRouterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniversalRouterApiDeleteUniversalRouterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Value \&quot;true\&quot; means to forcefully delete the object that contains other objects even if those objects are in a state that does not allow removal. The default is \&quot;false\&quot;; therefore, objects are not removed if they are not in a state that normally allows removal. Force also implies recursive delete where other contained objects are removed. Errors may be ignored. Invalid value (not true or false) are ignored.  | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniversalRouter**
> UniversalRouter GetUniversalRouter(ctx, universalRouterId)
Retrieves a specific Universal Router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 

### Return type

[**UniversalRouter**](UniversalRouter.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncUniversalRouter**
> SyncUniversalRouter(ctx, universalRouterId)
Sync/repair the Universal Router

Sync/repair the Universal Router 

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

# **UpdateUniversalRouter**
> UpdateUniversalRouter(ctx, router, universalRouterId)
Updates a specific Universal Router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **router** | [**UniversalRouter**](UniversalRouter.md)|  | 
  **universalRouterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

