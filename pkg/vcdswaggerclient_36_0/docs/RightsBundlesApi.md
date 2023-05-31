# \RightsBundlesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRightsToRightsBundle**](RightsBundlesApi.md#AddRightsToRightsBundle) | **Post** /1.0.0/rightsBundles/{id}/rights | Adds the specified rights to a rights bundle.
[**CreateRightsBundle**](RightsBundlesApi.md#CreateRightsBundle) | **Post** /1.0.0/rightsBundles | Creates a new rights bundle
[**DeleteRightsBundle**](RightsBundlesApi.md#DeleteRightsBundle) | **Delete** /1.0.0/rightsBundles/{id} | Delete specified rights bundle
[**GetRightsBundle**](RightsBundlesApi.md#GetRightsBundle) | **Get** /1.0.0/rightsBundles/{id} | Get specified rights bundle
[**PostRightsBundlePublish**](RightsBundlesApi.md#PostRightsBundlePublish) | **Post** /1.0.0/rightsBundles/{id}/tenants/publish | Publishes the rights bundle to the specified tenants
[**PostRightsBundlePublishAll**](RightsBundlesApi.md#PostRightsBundlePublishAll) | **Post** /1.0.0/rightsBundles/{id}/tenants/publishAll | Publishes the rights bundle to all tenants
[**PostRightsBundleUnpublish**](RightsBundlesApi.md#PostRightsBundleUnpublish) | **Post** /1.0.0/rightsBundles/{id}/tenants/unpublish | Revokes publication of the rights bundle to the specified tenants
[**PostRightsBundleUnpublishAll**](RightsBundlesApi.md#PostRightsBundleUnpublishAll) | **Post** /1.0.0/rightsBundles/{id}/tenants/unpublishAll | Unpublishes the rights bundle from all tenants
[**QueryRightsBundleRights**](RightsBundlesApi.md#QueryRightsBundleRights) | **Get** /1.0.0/rightsBundles/{id}/rights | Gets a paged list of rights (as references) contained by a particular bundle
[**QueryRightsBundleTenants**](RightsBundlesApi.md#QueryRightsBundleTenants) | **Get** /1.0.0/rightsBundles/{id}/tenants | Retrieves list of tenants for whom the rights bundle is explicitly published
[**QueryRightsBundles**](RightsBundlesApi.md#QueryRightsBundles) | **Get** /1.0.0/rightsBundles | Get list of rights bundles
[**ReplaceRightsInRightsBundle**](RightsBundlesApi.md#ReplaceRightsInRightsBundle) | **Put** /1.0.0/rightsBundles/{id}/rights | Replaces the existing set of rights in bundle with the rights (as references) supplied.
[**SetRightsBundleTenants**](RightsBundlesApi.md#SetRightsBundleTenants) | **Put** /1.0.0/rightsBundles/{id}/tenants | Resets list of tenants for whom the rights bundle is explicitly published
[**UpdateRightsBundle**](RightsBundlesApi.md#UpdateRightsBundle) | **Put** /1.0.0/rightsBundles/{id} | Update specified rights bundle


# **AddRightsToRightsBundle**
> EntityReferences AddRightsToRightsBundle(ctx, rightsReferencesBody, id)
Adds the specified rights to a rights bundle.

Adds the list of rights (passed as references) to a rights bundle. 

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

# **CreateRightsBundle**
> RightsBundle CreateRightsBundle(ctx, newRightsBundle)
Creates a new rights bundle

Creates a new rights bundle 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newRightsBundle** | [**RightsBundle**](RightsBundle.md)|  | 

### Return type

[**RightsBundle**](RightsBundle.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRightsBundle**
> DeleteRightsBundle(ctx, id)
Delete specified rights bundle

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

# **GetRightsBundle**
> RightsBundle GetRightsBundle(ctx, id)
Get specified rights bundle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**RightsBundle**](RightsBundle.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRightsBundlePublish**
> EntityReferences PostRightsBundlePublish(ctx, publishTenantsBody, id)
Publishes the rights bundle to the specified tenants

Publishes the rights bundle to the specified tenants 

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
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRightsBundlePublishAll**
> EntityReferences PostRightsBundlePublishAll(ctx, id)
Publishes the rights bundle to all tenants

Publishes the rights bundle to all tenants 

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
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRightsBundleUnpublish**
> EntityReferences PostRightsBundleUnpublish(ctx, unpublishTenantsBody, id)
Revokes publication of the rights bundle to the specified tenants

Revokes publication of the rights bundle to the specified tenants 

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
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRightsBundleUnpublishAll**
> EntityReferences PostRightsBundleUnpublishAll(ctx, id)
Unpublishes the rights bundle from all tenants

Unpublishes the rights bundle from all tenants 

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
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryRightsBundleRights**
> EntityReferences QueryRightsBundleRights(ctx, page, pageSize, id, optional)
Gets a paged list of rights (as references) contained by a particular bundle

Get list of rights (as references) contained by a particular bundle 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**|  | 
 **optional** | ***RightsBundlesApiQueryRightsBundleRightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RightsBundlesApiQueryRightsBundleRightsOpts struct

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

# **QueryRightsBundleTenants**
> EntityReferences QueryRightsBundleTenants(ctx, page, pageSize, id, optional)
Retrieves list of tenants for whom the rights bundle is explicitly published

Retrieves list of tenants for whom the rights bundle is explicitly published 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**|  | 
 **optional** | ***RightsBundlesApiQueryRightsBundleTenantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RightsBundlesApiQueryRightsBundleTenantsOpts struct

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

# **QueryRightsBundles**
> RightsBundles QueryRightsBundles(ctx, page, pageSize, optional)
Get list of rights bundles

Get list of rights bundles 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***RightsBundlesApiQueryRightsBundlesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RightsBundlesApiQueryRightsBundlesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**RightsBundles**](RightsBundles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRightsInRightsBundle**
> EntityReferences ReplaceRightsInRightsBundle(ctx, rightsReferencesBody, id)
Replaces the existing set of rights in bundle with the rights (as references) supplied.

Replaces the existing set of rights in bundle with the rights (as references) supplied. 

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

# **SetRightsBundleTenants**
> EntityReferences SetRightsBundleTenants(ctx, publishTenantsBody, id)
Resets list of tenants for whom the rights bundle is explicitly published

Resets list of tenants for whom the rights bundle is explicitly published 

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
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRightsBundle**
> RightsBundle UpdateRightsBundle(ctx, updatedRightsBundle, id)
Update specified rights bundle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedRightsBundle** | [**RightsBundle**](RightsBundle.md)|  | 
  **id** | **string**|  | 

### Return type

[**RightsBundle**](RightsBundle.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

