# \DefinedEntityApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDefinedEntity**](DefinedEntityApi.md#CreateDefinedEntity) | **Post** /1.0.0/entityTypes/{id} | Creates a defined entity based on the entity type (URN).
[**CreateMetadataEntry**](DefinedEntityApi.md#CreateMetadataEntry) | **Post** /1.0.0/entities/{id}/metadata | Creates a new entry.
[**DeleteDefinedEntity**](DefinedEntityApi.md#DeleteDefinedEntity) | **Delete** /1.0.0/entities/{id} | Deletes the defined entity with the unique identifier (URN)
[**DeleteMetadataEntry**](DefinedEntityApi.md#DeleteMetadataEntry) | **Delete** /1.0.0/entities/{id}/metadata/{metadataId} | Delete a single metadata entry.
[**GetDefinedEntitiesByEntityType**](DefinedEntityApi.md#GetDefinedEntitiesByEntityType) | **Get** /1.0.0/entities/types/{vendor}/{nss}/{version} | Gets the collection of defined entities for the vCD-defined type with the specified vendor, nss and version.
[**GetDefinedEntitiesByInterface**](DefinedEntityApi.md#GetDefinedEntitiesByInterface) | **Get** /1.0.0/entities/interfaces/{vendor}/{nss}/{version} | Gets the collection of defined entities for the vCD-defined interface with the specified vendor, nss and version
[**GetDefinedEntity**](DefinedEntityApi.md#GetDefinedEntity) | **Get** /1.0.0/entities/{id} | Gets the defined entity with the unique identifier (URN)
[**GetMetadata**](DefinedEntityApi.md#GetMetadata) | **Get** /1.0.0/entities/{id}/metadata | Retrieves all the metadata for the entity.
[**GetMetadataEntry**](DefinedEntityApi.md#GetMetadataEntry) | **Get** /1.0.0/entities/{id}/metadata/{metadataId} | Get a single metadata entry.
[**ResolveDefinedEntity**](DefinedEntityApi.md#ResolveDefinedEntity) | **Post** /1.0.0/entities/{id}/resolve | Validates the defined entity against the entity type schema.
[**UpdateDefinedEntity**](DefinedEntityApi.md#UpdateDefinedEntity) | **Put** /1.0.0/entities/{id} | Updates the defined entity with the unique identifier (URN)
[**UpdateMetadataEntry**](DefinedEntityApi.md#UpdateMetadataEntry) | **Put** /1.0.0/entities/{id}/metadata/{metadataId} | Update the value of a single key-value metadata entry.


# **CreateDefinedEntity**
> CreateDefinedEntity(ctx, entity, id, optional)
Creates a defined entity based on the entity type (URN).

Creates a defined entity based on the entity type (URN).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entity** | [**DefinedEntity**](DefinedEntity.md)|  | 
  **id** | **string**|  | 
 **optional** | ***DefinedEntityApiCreateDefinedEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedEntityApiCreateDefinedEntityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **invokeHooks** | [**optional.Interface of interface{}**](.md)| Only users with Admin FullControl access to the Entity Type can pass this parameter. The default value is &#39;true&#39;.  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMetadataEntry**
> MetadataEntry CreateMetadataEntry(ctx, entry, id)
Creates a new entry.

Creates a new entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entry** | [**MetadataEntry**](MetadataEntry.md)|  | 
  **id** | **string**| the URN of the entity the entry is attached to. | 

### Return type

[**MetadataEntry**](MetadataEntry.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDefinedEntity**
> DeleteDefinedEntity(ctx, id, optional)
Deletes the defined entity with the unique identifier (URN)

Deletes the defined entity with the unique identifier (URN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***DefinedEntityApiDeleteDefinedEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedEntityApiDeleteDefinedEntityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invokeHooks** | [**optional.Interface of interface{}**](.md)| Only users with Admin FullControl access to the Entity Type can pass this parameter. The default value is &#39;true&#39;.  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMetadataEntry**
> DeleteMetadataEntry(ctx, id, metadataId)
Delete a single metadata entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| the URN of the entity the entry is attached to. | 
  **metadataId** | **string**| a metadata vcloud id urn | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefinedEntitiesByEntityType**
> DefinedEntities GetDefinedEntitiesByEntityType(ctx, vendor, nss, version, page, pageSize, optional)
Gets the collection of defined entities for the vCD-defined type with the specified vendor, nss and version.

Gets the collection of defined entities for the vCD-defined type with the specified vendor, nss and version. The version can act as a wildcard. If only '1' is specified as the version, all entity types with a major version of '1' will be matched (e.g. 1.0.0, 1.1.2). If '1.0' is specified, all entity types with a major version of '1' and a minor version of '0' will be included (e.g. 1.0.0, 1.0.1). If the full semver is specified, then no search will be performed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendor** | **string**|  | 
  **nss** | **string**|  | 
  **version** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***DefinedEntityApiGetDefinedEntitiesByEntityTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedEntityApiGetDefinedEntitiesByEntityTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**DefinedEntities**](DefinedEntities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefinedEntitiesByInterface**
> DefinedEntities GetDefinedEntitiesByInterface(ctx, vendor, nss, version, page, pageSize, optional)
Gets the collection of defined entities for the vCD-defined interface with the specified vendor, nss and version

Gets the collection of defined entities for the vCD-defined interface with the specified vendor, nss and version. The version can act as a wildcard. If only '1' is specified as the version, all entity types with a major version of '1' will be matched (e.g. 1.0.0, 1.1.2). If '1.0' is specified, all entity types with a major version of '1' and a minor version of '0' will be included (e.g. 1.0.0, 1.0.1). If the full semver is specified, then no search will be performed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendor** | **string**|  | 
  **nss** | **string**|  | 
  **version** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***DefinedEntityApiGetDefinedEntitiesByInterfaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedEntityApiGetDefinedEntitiesByInterfaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**DefinedEntities**](DefinedEntities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefinedEntity**
> DefinedEntity GetDefinedEntity(ctx, id)
Gets the defined entity with the unique identifier (URN)

Gets the defined entity with the unique identifier (URN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DefinedEntity**](DefinedEntity.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetadata**
> MetadataEntries GetMetadata(ctx, page, pageSize, id, optional)
Retrieves all the metadata for the entity.

Retrieves all the metadata for the VM. User can view the tags if user can view the VM. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**| the URN of the entity the entry is attached to. | 
 **optional** | ***DefinedEntityApiGetMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedEntityApiGetMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**MetadataEntries**](MetadataEntries.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetadataEntry**
> MetadataEntry GetMetadataEntry(ctx, id, metadataId)
Get a single metadata entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| the URN of the entity the entry is attached to. | 
  **metadataId** | **string**| a metadata vcloud id urn | 

### Return type

[**MetadataEntry**](MetadataEntry.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResolveDefinedEntity**
> EntityState ResolveDefinedEntity(ctx, id)
Validates the defined entity against the entity type schema.

Validates the defined entity against the entity type schema. If the validation is successful, the entity will transition to a \"RESOLVED\" state. Otherwise, it will transition to an \"ERROR\" state. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**EntityState**](EntityState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDefinedEntity**
> DefinedEntity UpdateDefinedEntity(ctx, entity, id, optional)
Updates the defined entity with the unique identifier (URN)

Update the defined entity with the unique identifier (URN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entity** | [**DefinedEntity**](DefinedEntity.md)|  | 
  **id** | **string**|  | 
 **optional** | ***DefinedEntityApiUpdateDefinedEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedEntityApiUpdateDefinedEntityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **invokeHooks** | [**optional.Interface of interface{}**](.md)| Only users with Admin FullControl access to the Entity Type can pass this parameter. The default value is &#39;true&#39;.  | 

### Return type

[**DefinedEntity**](DefinedEntity.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMetadataEntry**
> MetadataEntry UpdateMetadataEntry(ctx, entry, id, metadataId)
Update the value of a single key-value metadata entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entry** | [**MetadataEntry**](MetadataEntry.md)|  | 
  **id** | **string**| the URN of the entity the entry is attached to. | 
  **metadataId** | **string**| a metadata vcloud id urn | 

### Return type

[**MetadataEntry**](MetadataEntry.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

