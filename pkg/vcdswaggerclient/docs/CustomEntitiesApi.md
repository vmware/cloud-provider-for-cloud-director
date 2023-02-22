# \CustomEntitiesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomEntityType**](CustomEntitiesApi.md#CreateCustomEntityType) | **Post** /customEntityTypes | Creates a new custom entity type
[**CreateCustomEntityTypeAction**](CustomEntitiesApi.md#CreateCustomEntityTypeAction) | **Post** /customEntityTypes/{customEntityTypeId}/actions | Creates a new custom entity type action.
[**DeleteCustomEntityType**](CustomEntitiesApi.md#DeleteCustomEntityType) | **Delete** /customEntityTypes/{customEntityTypeId} | Delete specified custom entity type.
[**DeleteCustomEntityTypeAction**](CustomEntitiesApi.md#DeleteCustomEntityTypeAction) | **Delete** /customEntityTypes/{customEntityTypeId}/actions/{workflowId} | Delete specified custom entity type action
[**GetCustomEntities**](CustomEntitiesApi.md#GetCustomEntities) | **Get** /customEntities | Get list of all custom entities
[**GetCustomEntity**](CustomEntitiesApi.md#GetCustomEntity) | **Get** /customEntities/{customEntityId} | Get specified custom entity
[**GetCustomEntityAsSdkObject**](CustomEntitiesApi.md#GetCustomEntityAsSdkObject) | **Get** /customEntities/{customEntityId}/sdkObject | Get specified custom entity represented as on Sdk-Object
[**GetCustomEntityType**](CustomEntitiesApi.md#GetCustomEntityType) | **Get** /customEntityTypes/{customEntityTypeId} | Get specified custom entity type
[**GetCustomEntityTypeActions**](CustomEntitiesApi.md#GetCustomEntityTypeActions) | **Get** /customEntityTypes/{customEntityTypeId}/actions | Get all actions associated with this custom entity type
[**GetCustomEntityTypeTenants**](CustomEntitiesApi.md#GetCustomEntityTypeTenants) | **Get** /customEntityTypes/{customEntityTypeId}/tenants | Retrieves list of tenants for whom the custom entity type is explicitly published
[**PostCustomEntityTypePublish**](CustomEntitiesApi.md#PostCustomEntityTypePublish) | **Post** /customEntityTypes/{customEntityTypeId}/tenants/publish | Publishes the custom entity type to the specified tenants
[**PostCustomEntityTypePublishAll**](CustomEntitiesApi.md#PostCustomEntityTypePublishAll) | **Post** /customEntityTypes/{customEntityTypeId}/tenants/publishAll | Publishes the custom entity type to all tenants
[**PostCustomEntityTypeUnpublish**](CustomEntitiesApi.md#PostCustomEntityTypeUnpublish) | **Post** /customEntityTypes/{customEntityTypeId}/tenants/unpublish | Revokes publication of the custom entity type to the specified tenants
[**PostCustomEntityTypeUnpublishAll**](CustomEntitiesApi.md#PostCustomEntityTypeUnpublishAll) | **Post** /customEntityTypes/{customEntityTypeId}/tenants/unpublishAll | Unpublishes the custom entity type from all tenants
[**QueryCustomEntityTypes**](CustomEntitiesApi.md#QueryCustomEntityTypes) | **Get** /customEntityTypes | Get list of all custom entity types
[**SetCustomEntityTypeTenants**](CustomEntitiesApi.md#SetCustomEntityTypeTenants) | **Put** /customEntityTypes/{customEntityTypeId}/tenants | Resets list of tenants for whom the custom entity type is explicitly published
[**UpdateCustomEntityType**](CustomEntitiesApi.md#UpdateCustomEntityType) | **Put** /customEntityTypes/{customEntityTypeId} | Update specified custom entity type


# **CreateCustomEntityType**
> CustomEntityType CreateCustomEntityType(ctx, newCustomEntityType)
Creates a new custom entity type

Creates a new custom entity type 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newCustomEntityType** | [**CustomEntityType**](CustomEntityType.md)|  | 

### Return type

[**CustomEntityType**](CustomEntityType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomEntityTypeAction**
> CustomEntityTypeAction CreateCustomEntityTypeAction(ctx, newCustomEntityTypeAction, customEntityTypeId)
Creates a new custom entity type action.

Creates a new custom entity type action 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newCustomEntityTypeAction** | [**CreateCustomEntityTypeAction**](CreateCustomEntityTypeAction.md)|  | 
  **customEntityTypeId** | **string**|  | 

### Return type

[**CustomEntityTypeAction**](CustomEntityTypeAction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomEntityType**
> DeleteCustomEntityType(ctx, recursive, customEntityTypeId)
Delete specified custom entity type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recursive** | **bool**| if true, will recursively delete both custom entity type, all its instances and associated actions | 
  **customEntityTypeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomEntityTypeAction**
> DeleteCustomEntityTypeAction(ctx, customEntityTypeId, workflowId)
Delete specified custom entity type action

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEntityTypeId** | **string**|  | 
  **workflowId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomEntities**
> CustomEntities GetCustomEntities(ctx, optional)
Get list of all custom entities

Get list of custom entities 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomEntitiesApiGetCustomEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomEntitiesApiGetCustomEntitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**CustomEntities**](CustomEntities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomEntity**
> CustomEntity GetCustomEntity(ctx, customEntityId)
Get specified custom entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEntityId** | **string**|  | 

### Return type

[**CustomEntity**](CustomEntity.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomEntityAsSdkObject**
> SdkObject GetCustomEntityAsSdkObject(ctx, customEntityId)
Get specified custom entity represented as on Sdk-Object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEntityId** | **string**|  | 

### Return type

[**SdkObject**](SdkObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomEntityType**
> CustomEntityType GetCustomEntityType(ctx, customEntityTypeId)
Get specified custom entity type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEntityTypeId** | **string**|  | 

### Return type

[**CustomEntityType**](CustomEntityType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomEntityTypeActions**
> CustomEntityTypeActions GetCustomEntityTypeActions(ctx, customEntityTypeId)
Get all actions associated with this custom entity type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEntityTypeId** | **string**|  | 

### Return type

[**CustomEntityTypeActions**](CustomEntityTypeActions.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomEntityTypeTenants**
> []EntityReference GetCustomEntityTypeTenants(ctx, customEntityTypeId)
Retrieves list of tenants for whom the custom entity type is explicitly published

Retrieves list of item for whom the custom entity type is explicitly published 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEntityTypeId** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCustomEntityTypePublish**
> []EntityReference PostCustomEntityTypePublish(ctx, publishTenantsBody, customEntityTypeId)
Publishes the custom entity type to the specified tenants

Publishes the custom entity type to the specified tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publishTenantsBody** | [**[]EntityReference**](EntityReference.md)|  | 
  **customEntityTypeId** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCustomEntityTypePublishAll**
> []EntityReference PostCustomEntityTypePublishAll(ctx, customEntityTypeId)
Publishes the custom entity type to all tenants

Publishes the custom entity type to all tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEntityTypeId** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCustomEntityTypeUnpublish**
> []EntityReference PostCustomEntityTypeUnpublish(ctx, unpublishTenantsBody, customEntityTypeId)
Revokes publication of the custom entity type to the specified tenants

Revokes publication of the custom entity type to the specified tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unpublishTenantsBody** | [**[]EntityReference**](EntityReference.md)|  | 
  **customEntityTypeId** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCustomEntityTypeUnpublishAll**
> []EntityReference PostCustomEntityTypeUnpublishAll(ctx, customEntityTypeId)
Unpublishes the custom entity type from all tenants

Unpublishes the custom entity type from all tenants 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEntityTypeId** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryCustomEntityTypes**
> CustomEntityTypes QueryCustomEntityTypes(ctx, page, pageSize, optional)
Get list of all custom entity types

Get list of custom entity types. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***CustomEntitiesApiQueryCustomEntityTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomEntitiesApiQueryCustomEntityTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**CustomEntityTypes**](CustomEntityTypes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetCustomEntityTypeTenants**
> []EntityReference SetCustomEntityTypeTenants(ctx, publishTenantsBody, customEntityTypeId)
Resets list of tenants for whom the custom entity type is explicitly published

Resets list of tenants for whom the custom entity type is explicitly published 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publishTenantsBody** | [**[]EntityReference**](EntityReference.md)|  | 
  **customEntityTypeId** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomEntityType**
> CustomEntityType UpdateCustomEntityType(ctx, updatedCustomEntityType, customEntityTypeId)
Update specified custom entity type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedCustomEntityType** | [**CustomEntityType**](CustomEntityType.md)|  | 
  **customEntityTypeId** | **string**|  | 

### Return type

[**CustomEntityType**](CustomEntityType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

