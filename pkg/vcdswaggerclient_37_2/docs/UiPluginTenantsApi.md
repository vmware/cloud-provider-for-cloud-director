# \UiPluginTenantsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUiPluginTenants**](UiPluginTenantsApi.md#GetUiPluginTenants) | **Get** /extensions/ui/{id}/tenants | Retrieves list of tenants for whom the plugin is explicitly published
[**PostUiPluginPublish**](UiPluginTenantsApi.md#PostUiPluginPublish) | **Post** /extensions/ui/{id}/tenants/publish | Publishes the UI plugin to the specified tenants
[**PostUiPluginPublishAll**](UiPluginTenantsApi.md#PostUiPluginPublishAll) | **Post** /extensions/ui/{id}/tenants/publishAll | Publishes the UI plugin to all tenants
[**PostUiPluginUnpublish**](UiPluginTenantsApi.md#PostUiPluginUnpublish) | **Post** /extensions/ui/{id}/tenants/unpublish | Revokes publication of the UI plugin to the specified tenants
[**PostUiPluginUnpublishAll**](UiPluginTenantsApi.md#PostUiPluginUnpublishAll) | **Post** /extensions/ui/{id}/tenants/unpublishAll | Unpublishes the UI plugin from all tenants


# **GetUiPluginTenants**
> EntityReferences GetUiPluginTenants(ctx, id, page, pageSize, optional)
Retrieves list of tenants for whom the plugin is explicitly published

Retrieves list of tenants for whom the plugin is explicitly published 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***UiPluginTenantsApiGetUiPluginTenantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UiPluginTenantsApiGetUiPluginTenantsOpts struct

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

# **PostUiPluginPublish**
> EntityReferences PostUiPluginPublish(ctx, publishTenantsBody, id)
Publishes the UI plugin to the specified tenants

Publishes the UI plugin to the specified tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publishTenantsBody** | [**[]EntityReference**](EntityReference.md)|  | 
  **id** | **string**|  | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUiPluginPublishAll**
> EntityReferences PostUiPluginPublishAll(ctx, id)
Publishes the UI plugin to all tenants

Publishes the UI plugin to all tenants 

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

# **PostUiPluginUnpublish**
> EntityReferences PostUiPluginUnpublish(ctx, unpublishTenantsBody, id)
Revokes publication of the UI plugin to the specified tenants

Revokes publication of the UI plugin to the specified tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unpublishTenantsBody** | [**[]EntityReference**](EntityReference.md)|  | 
  **id** | **string**|  | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUiPluginUnpublishAll**
> EntityReferences PostUiPluginUnpublishAll(ctx, id)
Unpublishes the UI plugin from all tenants

Unpublishes the UI plugin from all tenants 

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
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

