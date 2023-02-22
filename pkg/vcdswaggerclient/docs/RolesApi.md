# \RolesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRightsToRole**](RolesApi.md#AddRightsToRole) | **Post** /1.0.0/roles/{id}/rights | Adds the specified rights to a role.
[**CreateRole**](RolesApi.md#CreateRole) | **Post** /1.0.0/roles | Creates a new role
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /1.0.0/roles/{id} | Delete specified role
[**GetRole**](RolesApi.md#GetRole) | **Get** /1.0.0/roles/{id} | Get specified role
[**QueryRoleRights**](RolesApi.md#QueryRoleRights) | **Get** /1.0.0/roles/{id}/rights | Gets a paged list of rights (as references) contained by a particular role
[**QueryTenantRoles**](RolesApi.md#QueryTenantRoles) | **Get** /1.0.0/roles | Get list of roles for a tenant
[**ReplaceRightsInRole**](RolesApi.md#ReplaceRightsInRole) | **Put** /1.0.0/roles/{id}/rights | Replaces the existing set of rights in role with the rights (as references) supplied.
[**UpdateRole**](RolesApi.md#UpdateRole) | **Put** /1.0.0/roles/{id} | Update specified role


# **AddRightsToRole**
> EntityReferences AddRightsToRole(ctx, rightsReferencesBody, id)
Adds the specified rights to a role.

Adds the list of rights (passed as references) to a role. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rightsReferencesBody** | [**EntityReferences**](EntityReferences.md)|  | 
  **id** | **string**|  | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRole**
> Role CreateRole(ctx, newRole)
Creates a new role

Creates a new role 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newRole** | [**Role**](Role.md)|  | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> DeleteRole(ctx, id)
Delete specified role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> Role GetRole(ctx, id)
Get specified role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryRoleRights**
> EntityReferences QueryRoleRights(ctx, page, pageSize, id, optional)
Gets a paged list of rights (as references) contained by a particular role

Get list of rights (as references) contained by a particular role 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**|  | 
 **optional** | ***RolesApiQueryRoleRightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiQueryRoleRightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryTenantRoles**
> Roles QueryTenantRoles(ctx, page, pageSize, optional)
Get list of roles for a tenant

Get list of roles for a tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***RolesApiQueryTenantRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiQueryTenantRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Roles**](Roles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRightsInRole**
> EntityReferences ReplaceRightsInRole(ctx, rightsReferencesBody, id)
Replaces the existing set of rights in role with the rights (as references) supplied.

Replaces the existing set of rights in role with the rights (as references) supplied. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rightsReferencesBody** | [**EntityReferences**](EntityReferences.md)|  | 
  **id** | **string**|  | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRole**
> Role UpdateRole(ctx, updatedRole, id)
Update specified role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedRole** | [**Role**](Role.md)|  | 
  **id** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

