# \UniversalRoutersApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUniversalRouter**](UniversalRoutersApi.md#CreateUniversalRouter) | **Post** /1.0.0/universalRouters | Create a new universal router
[**GetAllUniversalRoutersForVdcGroup**](UniversalRoutersApi.md#GetAllUniversalRoutersForVdcGroup) | **Get** /1.0.0/vdcGroups/{vdcGroupId}/universalRouters | Get all the universal routers defined for a vDC group


# **CreateUniversalRouter**
> CreateUniversalRouter(ctx, universalRouter)
Create a new universal router

Create a new universal router for a vDC group 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouter** | [**UniversalRouter**](UniversalRouter.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllUniversalRoutersForVdcGroup**
> UniversalRouters GetAllUniversalRoutersForVdcGroup(ctx, vdcGroupId)
Get all the universal routers defined for a vDC group

Get all the universal routers defined for a vDC group in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 

### Return type

[**UniversalRouters**](UniversalRouters.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

