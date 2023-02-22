# \AccessControlsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEntityAccessControlGrant**](AccessControlsApi.md#CreateEntityAccessControlGrant) | **Post** /1.0.0/entities/{objectId}/accessControls | Creates an access-control grant
[**GetEntityAccessControlGrant**](AccessControlsApi.md#GetEntityAccessControlGrant) | **Get** /1.0.0/entities/{objectId}/accessControls/{accessControlId} | Get the specified access-control grant.
[**QueryEntityAccessControlGrants**](AccessControlsApi.md#QueryEntityAccessControlGrants) | **Get** /1.0.0/entities/{objectId}/accessControls | Get the access-control list for the specified vCD entity.
[**RemoveEntityAccessControlGrant**](AccessControlsApi.md#RemoveEntityAccessControlGrant) | **Delete** /1.0.0/entities/{objectId}/accessControls/{accessControlId} | Removes the specified access-control grant from the vCD entities access-control list.
[**UpdateEntityAccessControlGrant**](AccessControlsApi.md#UpdateEntityAccessControlGrant) | **Put** /1.0.0/entities/{objectId}/accessControls/{accessControlId} | Updates the specified access-control grant.


# **CreateEntityAccessControlGrant**
> AccessControlGrant CreateEntityAccessControlGrant(ctx, objectId, accessControlGrant)
Creates an access-control grant

Creates an access-control grant, giving the user the level of access for the vCD entity. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 
  **accessControlGrant** | [**AccessControlGrant**](AccessControlGrant.md)|  | 

### Return type

[**AccessControlGrant**](AccessControlGrant.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntityAccessControlGrant**
> AccessControlGrant GetEntityAccessControlGrant(ctx, objectId, accessControlId)
Get the specified access-control grant.

Get the specified access-control grant. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 
  **accessControlId** | **string**|  | 

### Return type

[**AccessControlGrant**](AccessControlGrant.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryEntityAccessControlGrants**
> AccessControlGrants QueryEntityAccessControlGrants(ctx, objectId, page, pageSize)
Get the access-control list for the specified vCD entity.

Get the access-control list for the specified vCD entity. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]

### Return type

[**AccessControlGrants**](AccessControlGrants.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveEntityAccessControlGrant**
> RemoveEntityAccessControlGrant(ctx, objectId, accessControlId)
Removes the specified access-control grant from the vCD entities access-control list.

Removes the specified access-control grant from the vCD entities access-control list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 
  **accessControlId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEntityAccessControlGrant**
> AccessControlGrant UpdateEntityAccessControlGrant(ctx, objectId, accessControlId, accessControlGrant)
Updates the specified access-control grant.

Updates the specified access-control grant. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 
  **accessControlId** | **string**|  | 
  **accessControlGrant** | [**AccessControlGrant**](AccessControlGrant.md)|  | 

### Return type

[**AccessControlGrant**](AccessControlGrant.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

