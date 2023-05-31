# \GlobalRolesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRightsToGlobalRole**](GlobalRolesApi.md#AddRightsToGlobalRole) | **Post** /1.0.0/globalRoles/{id}/rights | Adds the specified rights to a global role.
[**CreateGlobalRole**](GlobalRolesApi.md#CreateGlobalRole) | **Post** /1.0.0/globalRoles | Creates a new global role
[**DeleteGlobalRole**](GlobalRolesApi.md#DeleteGlobalRole) | **Delete** /1.0.0/globalRoles/{id} | Delete specified global role
[**GetGlobalRole**](GlobalRolesApi.md#GetGlobalRole) | **Get** /1.0.0/globalRoles/{id} | Get specified global role
[**PostGlobalRolePublish**](GlobalRolesApi.md#PostGlobalRolePublish) | **Post** /1.0.0/globalRoles/{id}/tenants/publish | Publishes the global role to the specified tenants
[**PostGlobalRolePublishAll**](GlobalRolesApi.md#PostGlobalRolePublishAll) | **Post** /1.0.0/globalRoles/{id}/tenants/publishAll | Publishes the global role to all tenants
[**PostGlobalRoleUnpublish**](GlobalRolesApi.md#PostGlobalRoleUnpublish) | **Post** /1.0.0/globalRoles/{id}/tenants/unpublish | Revokes publication of the global role to the specified tenants
[**PostGlobalRoleUnpublishAll**](GlobalRolesApi.md#PostGlobalRoleUnpublishAll) | **Post** /1.0.0/globalRoles/{id}/tenants/unpublishAll | Unpublishes the global role from all tenants
[**QueryGlobalRoleRights**](GlobalRolesApi.md#QueryGlobalRoleRights) | **Get** /1.0.0/globalRoles/{id}/rights | Gets a paged list of rights (as references) contained by a particular global role
[**QueryGlobalRoleTenants**](GlobalRolesApi.md#QueryGlobalRoleTenants) | **Get** /1.0.0/globalRoles/{id}/tenants | Retrieves list of tenants for whom the global role is explicitly published
[**QueryGlobalRoles**](GlobalRolesApi.md#QueryGlobalRoles) | **Get** /1.0.0/globalRoles | Get list of global roles
[**ReplaceRightsInGlobalRole**](GlobalRolesApi.md#ReplaceRightsInGlobalRole) | **Put** /1.0.0/globalRoles/{id}/rights | Replaces the existing set of rights in global role with the rights (as references) supplied.
[**SetGlobalRoleTenants**](GlobalRolesApi.md#SetGlobalRoleTenants) | **Put** /1.0.0/globalRoles/{id}/tenants | Resets list of tenants for whom the global role is explicitly published
[**UpdateGlobalRole**](GlobalRolesApi.md#UpdateGlobalRole) | **Put** /1.0.0/globalRoles/{id} | Update specified global role


# **AddRightsToGlobalRole**
> EntityReferences AddRightsToGlobalRole(ctx, rightsReferencesBody, id)
Adds the specified rights to a global role.

Adds the list of rights (passed as references) to a global role. 

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGlobalRole**
> GlobalRole CreateGlobalRole(ctx, newGlobalRole)
Creates a new global role

Creates a new global role 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newGlobalRole** | [**GlobalRole**](GlobalRole.md)|  | 

### Return type

[**GlobalRole**](GlobalRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGlobalRole**
> DeleteGlobalRole(ctx, id)
Delete specified global role

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
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalRole**
> GlobalRole GetGlobalRole(ctx, id)
Get specified global role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**GlobalRole**](GlobalRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGlobalRolePublish**
> EntityReferences PostGlobalRolePublish(ctx, publishTenantsBody, id)
Publishes the global role to the specified tenants

Publishes the global role to the specified tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publishTenantsBody** | [**EntityReferences**](EntityReferences.md)|  | 
  **id** | **string**|  | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGlobalRolePublishAll**
> EntityReferences PostGlobalRolePublishAll(ctx, id)
Publishes the global role to all tenants

Publishes the global role to all tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGlobalRoleUnpublish**
> EntityReferences PostGlobalRoleUnpublish(ctx, unpublishTenantsBody, id)
Revokes publication of the global role to the specified tenants

Revokes publication of the global role to the specified tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unpublishTenantsBody** | [**EntityReferences**](EntityReferences.md)|  | 
  **id** | **string**|  | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGlobalRoleUnpublishAll**
> EntityReferences PostGlobalRoleUnpublishAll(ctx, id)
Unpublishes the global role from all tenants

Unpublishes the global role from all tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryGlobalRoleRights**
> EntityReferences QueryGlobalRoleRights(ctx, page, pageSize, id, optional)
Gets a paged list of rights (as references) contained by a particular global role

Get list of rights (as references) contained by a particular global role 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**|  | 
 **optional** | ***GlobalRolesApiQueryGlobalRoleRightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalRolesApiQueryGlobalRoleRightsOpts struct

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryGlobalRoleTenants**
> EntityReferences QueryGlobalRoleTenants(ctx, page, pageSize, id, optional)
Retrieves list of tenants for whom the global role is explicitly published

Retrieves list of tenants for whom the global role is explicitly published 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**|  | 
 **optional** | ***GlobalRolesApiQueryGlobalRoleTenantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalRolesApiQueryGlobalRoleTenantsOpts struct

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryGlobalRoles**
> GlobalRoles QueryGlobalRoles(ctx, page, pageSize, optional)
Get list of global roles

Get list of global roles 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***GlobalRolesApiQueryGlobalRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalRolesApiQueryGlobalRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**GlobalRoles**](GlobalRoles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRightsInGlobalRole**
> EntityReferences ReplaceRightsInGlobalRole(ctx, rightsReferencesBody, id)
Replaces the existing set of rights in global role with the rights (as references) supplied.

Replaces the existing set of rights in global role with the rights (as references) supplied. 

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetGlobalRoleTenants**
> EntityReferences SetGlobalRoleTenants(ctx, publishTenantsBody, id)
Resets list of tenants for whom the global role is explicitly published

Resets list of tenants for whom the global role is explicitly published 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publishTenantsBody** | [**EntityReferences**](EntityReferences.md)|  | 
  **id** | **string**|  | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGlobalRole**
> GlobalRole UpdateGlobalRole(ctx, updatedGlobalRole, id)
Update specified global role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedGlobalRole** | [**GlobalRole**](GlobalRole.md)|  | 
  **id** | **string**|  | 

### Return type

[**GlobalRole**](GlobalRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

