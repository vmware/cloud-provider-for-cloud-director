# \SecurityTagsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecurityTaggedEntities**](SecurityTagsApi.md#GetSecurityTaggedEntities) | **Get** /1.0.0/securityTags/entities | Retrieves the list of entities that have at least one tag assigned to it.
[**GetTagValues**](SecurityTagsApi.md#GetTagValues) | **Get** /1.0.0/securityTags/values | Retrieves the list of security tags that are in the organization.
[**GetVmTags**](SecurityTagsApi.md#GetVmTags) | **Get** /1.0.0/securityTags/vm/{id} | Retrieves the list of tags for a specific VM.
[**UpdateSecurityTag**](SecurityTagsApi.md#UpdateSecurityTag) | **Put** /1.0.0/securityTags/tag | Updates a specific tag
[**UpdateVmTags**](SecurityTagsApi.md#UpdateVmTags) | **Put** /1.0.0/securityTags/vm/{id} | Update the list of tags for a specific VM.


# **GetSecurityTaggedEntities**
> SecurityTaggedEntities GetSecurityTaggedEntities(ctx, page, pageSize, optional)
Retrieves the list of entities that have at least one tag assigned to it.

Retrieves the list of entities that have at least one tag assigned to it. Besides entityType, additional supported filters are: <ul> <li>tag - The tag to search by </ul> Example: <code>filter=(tag==Web;entityType==vm)</code> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***SecurityTagsApiGetSecurityTaggedEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityTagsApiGetSecurityTaggedEntitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**SecurityTaggedEntities**](SecurityTaggedEntities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagValues**
> SecurityTagValues GetTagValues(ctx, page, pageSize, optional)
Retrieves the list of security tags that are in the organization.

Retrieves the list of security tags that are in the organization and can be reused to tag an entity. The list of tags include tags assigned to entities within the organization. This API is meant for organization user only (i.e. not system provider). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***SecurityTagsApiGetTagValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecurityTagsApiGetTagValuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**SecurityTagValues**](SecurityTagValues.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVmTags**
> EntitySecurityTags GetVmTags(ctx, id)
Retrieves the list of tags for a specific VM.

Retrieves the list of tags for a specific VM. If user has view right to the VM, user can view its tags. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| the URN of the VM to manage tags for. | 

### Return type

[**EntitySecurityTags**](EntitySecurityTags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecurityTag**
> UpdateSecurityTag(ctx, securityTag)
Updates a specific tag

Only the list of tagged entities can be updated. The name cannot be updated. Any other existing entities not in the list will be untagged. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **securityTag** | [**SecurityTag**](SecurityTag.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVmTags**
> EntitySecurityTags UpdateVmTags(ctx, entityTags, id)
Update the list of tags for a specific VM.

Update the list of tags for a specific VM. An empty list of tags means to delete all dags for the VM. If user has edit permission on the VM, user can edit its tags. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityTags** | [**EntitySecurityTags**](EntitySecurityTags.md)| the list of tags to update. | 
  **id** | **string**| the URN of the VM to manage tags for. | 

### Return type

[**EntitySecurityTags**](EntitySecurityTags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

