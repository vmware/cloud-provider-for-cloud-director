# \ServicesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /serviceLibrary | Creates a new service
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /serviceLibrary/{id} | Delete specified service
[**DeleteServiceItem**](ServicesApi.md#DeleteServiceItem) | **Delete** /serviceItem/{id} | Deletes specified service item
[**GetService**](ServicesApi.md#GetService) | **Get** /serviceLibrary/{id} | Get specified service
[**GetServiceItem**](ServicesApi.md#GetServiceItem) | **Get** /serviceItem/{id} | Get the specified item
[**GetServiceItemTenants**](ServicesApi.md#GetServiceItemTenants) | **Get** /serviceItem/{id}/tenants | Retrieves list of tenants for whom the service item is explicitly published
[**GetWorkflowServiceItems**](ServicesApi.md#GetWorkflowServiceItems) | **Get** /serviceLibrary/{id}/workflows | This endpoint will not produce results. It is a placeholder to enforce code generation of VroWorkflowServiceItem
[**ImportVroWorkflows**](ServicesApi.md#ImportVroWorkflows) | **Post** /serviceLibrary/{id}/workflows | Add VRO remote workflows to this service
[**PostServiceItemPublish**](ServicesApi.md#PostServiceItemPublish) | **Post** /serviceItem/{id}/tenants/publish | Publishes the service item to the specified tenants
[**PostServiceItemPublishAll**](ServicesApi.md#PostServiceItemPublishAll) | **Post** /serviceItem/{id}/tenants/publishAll | Publishes the service item to all tenants
[**PostServiceItemUnpublish**](ServicesApi.md#PostServiceItemUnpublish) | **Post** /serviceItem/{id}/tenants/unpublish | Revokes publication of the service item to the specified tenants
[**PostServiceItemUnpublishAll**](ServicesApi.md#PostServiceItemUnpublishAll) | **Post** /serviceItem/{id}/tenants/unpublishAll | Unpublishes the service item from all tenants
[**QueryServiceItems**](ServicesApi.md#QueryServiceItems) | **Get** /serviceItem | Get all items across all services
[**QueryServices**](ServicesApi.md#QueryServices) | **Get** /serviceLibrary | Get list of services
[**SetServiceItemTenants**](ServicesApi.md#SetServiceItemTenants) | **Put** /serviceItem/{id}/tenants | Resets list of tenants for whom the service item is explicitly published
[**UpdateService**](ServicesApi.md#UpdateService) | **Put** /serviceLibrary/{id} | Update specified service
[**UpdateServiceItem**](ServicesApi.md#UpdateServiceItem) | **Put** /serviceItem/{id} | Update specified service item


# **CreateService**
> Service CreateService(ctx, newService)
Creates a new service

Creates a new service 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newService** | [**Service**](Service.md)|  | 

### Return type

[**Service**](Service.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteService**
> DeleteService(ctx, id, optional)
Delete specified service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ServicesApiDeleteServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiDeleteServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **optional.Bool**| If true, the Service and all its service items will be deleted. If false, and there are service items in the service, delete will fail. | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceItem**
> DeleteServiceItem(ctx, id)
Deletes specified service item

Deletes specified service item 

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

# **GetService**
> Service GetService(ctx, id)
Get specified service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Service**](Service.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceItem**
> ServiceItem GetServiceItem(ctx, id)
Get the specified item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ServiceItem**](ServiceItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceItemTenants**
> []EntityReference GetServiceItemTenants(ctx, id)
Retrieves list of tenants for whom the service item is explicitly published

Retrieves list of item for whom the service item is explicitly published 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflowServiceItems**
> []VroWorkflowServiceItem GetWorkflowServiceItems(ctx, id)
This endpoint will not produce results. It is a placeholder to enforce code generation of VroWorkflowServiceItem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**[]VroWorkflowServiceItem**](VroWorkflowServiceItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportVroWorkflows**
> []ServiceItem ImportVroWorkflows(ctx, remoteWorkflows, id)
Add VRO remote workflows to this service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **remoteWorkflows** | [**[]VroRemoteWorkflowItem**](VroRemoteWorkflowItem.md)|  | 
  **id** | **string**|  | 

### Return type

[**[]ServiceItem**](ServiceItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServiceItemPublish**
> []EntityReference PostServiceItemPublish(ctx, publishTenantsBody, id)
Publishes the service item to the specified tenants

Publishes the service item to the specified tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publishTenantsBody** | [**[]EntityReference**](EntityReference.md)|  | 
  **id** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServiceItemPublishAll**
> []EntityReference PostServiceItemPublishAll(ctx, id)
Publishes the service item to all tenants

Publishes the service item to all tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServiceItemUnpublish**
> []EntityReference PostServiceItemUnpublish(ctx, unpublishTenantsBody, id)
Revokes publication of the service item to the specified tenants

Revokes publication of the service item to the specified tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unpublishTenantsBody** | [**[]EntityReference**](EntityReference.md)|  | 
  **id** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServiceItemUnpublishAll**
> []EntityReference PostServiceItemUnpublishAll(ctx, id)
Unpublishes the service item from all tenants

Unpublishes the service item from all tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryServiceItems**
> ServiceItems QueryServiceItems(ctx, page, pageSize, optional)
Get all items across all services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***ServicesApiQueryServiceItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiQueryServiceItemsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 
 **getExternalData** | **optional.Bool**| Flag indicating whether data stored outside of vCloud Director should be included in results | [default to true]

### Return type

[**ServiceItems**](ServiceItems.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryServices**
> Services QueryServices(ctx, page, pageSize, optional)
Get list of services

Get list of services 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***ServicesApiQueryServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicesApiQueryServicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Services**](Services.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetServiceItemTenants**
> []EntityReference SetServiceItemTenants(ctx, publishTenantsBody, id)
Resets list of tenants for whom the service item is explicitly published

Resets list of tenants for whom the service item is explicitly published 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publishTenantsBody** | [**[]EntityReference**](EntityReference.md)|  | 
  **id** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateService**
> Service UpdateService(ctx, updatedService, id)
Update specified service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedService** | [**Service**](Service.md)|  | 
  **id** | **string**|  | 

### Return type

[**Service**](Service.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceItem**
> ServiceItem UpdateServiceItem(ctx, updatedService, id)
Update specified service item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedService** | [**ServiceItem**](ServiceItem.md)|  | 
  **id** | **string**|  | 

### Return type

[**ServiceItem**](ServiceItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

