# \TypeAccessControlsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEntityTypeAccessControlGrant**](TypeAccessControlsApi.md#CreateEntityTypeAccessControlGrant) | **Post** /1.0.0/entityTypes/{objectId}/accessControls | Creates an access-control grant for an entity type
[**GetEntityTypeAccessControlGrant**](TypeAccessControlsApi.md#GetEntityTypeAccessControlGrant) | **Get** /1.0.0/entityTypes/{objectId}/accessControls/{accessControlId} | Get the specified access-control grant.
[**QueryEntityTypeAccessControlGrants**](TypeAccessControlsApi.md#QueryEntityTypeAccessControlGrants) | **Get** /1.0.0/entityTypes/{objectId}/accessControls | Get the access-control list for the specified vCD entity type.
[**RemoveEntityTypeAccessControlGrant**](TypeAccessControlsApi.md#RemoveEntityTypeAccessControlGrant) | **Delete** /1.0.0/entityTypes/{objectId}/accessControls/{accessControlId} | Removes the specified access-control grant from the vCD entity type access-control list.
[**UpdateEntityTypeAccessControlGrant**](TypeAccessControlsApi.md#UpdateEntityTypeAccessControlGrant) | **Put** /1.0.0/entityTypes/{objectId}/accessControls/{accessControlId} | Updates the specified access-control grant.


# **CreateEntityTypeAccessControlGrant**
> AccessControlGrant CreateEntityTypeAccessControlGrant(ctx, objectId, accessControlGrant)
Creates an access-control grant for an entity type

Creates an access-control grant, giving the user the level of access for the vCD entity type. 

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntityTypeAccessControlGrant**
> AccessControlGrant GetEntityTypeAccessControlGrant(ctx, objectId, accessControlId)
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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryEntityTypeAccessControlGrants**
> AccessControlGrants QueryEntityTypeAccessControlGrants(ctx, objectId, page, pageSize)
Get the access-control list for the specified vCD entity type.

Get the access-control list for the specified vCD entity type. 

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveEntityTypeAccessControlGrant**
> RemoveEntityTypeAccessControlGrant(ctx, objectId, accessControlId)
Removes the specified access-control grant from the vCD entity type access-control list.

Removes the specified access-control grant from the vCD entity type access-control list. 

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
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEntityTypeAccessControlGrant**
> AccessControlGrant UpdateEntityTypeAccessControlGrant(ctx, objectId, accessControlId, accessControlGrant)
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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

