# \DefinedEntityApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDefinedEntity**](DefinedEntityApi.md#CreateDefinedEntity) | **Post** /1.0.0/entityTypes/{id} | Creates a defined entity based on the entity type (URN).
[**CreateMetadataEntry**](DefinedEntityApi.md#CreateMetadataEntry) | **Post** /1.0.0/entities/{id}/metadata | Creates a new entry.
[**DeleteDefinedEntity**](DefinedEntityApi.md#DeleteDefinedEntity) | **Delete** /1.0.0/entities/{id} | Deletes the defined entity with the unique identifier (URN)
[**DeleteMetadataEntry**](DefinedEntityApi.md#DeleteMetadataEntry) | **Delete** /1.0.0/entities/{id}/metadata/{metadataId} | Delete a single metadata entry.
[**GetDefinedEntitiesByEntityType**](DefinedEntityApi.md#GetDefinedEntitiesByEntityType) | **Get** /1.0.0/entities/types/{vendor}/{nss}/{version} | Gets the collection of defined entities for the vCD-defined type with the specified vendor, nss and version.
[**GetDefinedEntitiesByEntityTypeNoVersionSpecified**](DefinedEntityApi.md#GetDefinedEntitiesByEntityTypeNoVersionSpecified) | **Get** /1.0.0/entities/types/{vendor}/{nss} | Gets the collection of defined entities for the vCD-defined type with the specified vendor and nss.
[**GetDefinedEntitiesByInterface**](DefinedEntityApi.md#GetDefinedEntitiesByInterface) | **Get** /1.0.0/entities/interfaces/{vendor}/{nss}/{version} | Gets the collection of defined entities for the vCD-defined interface with the specified vendor, nss and version
[**GetDefinedEntity**](DefinedEntityApi.md#GetDefinedEntity) | **Get** /1.0.0/entities/{id} | Gets the defined entity with the unique identifier (URN)
[**GetMetadata**](DefinedEntityApi.md#GetMetadata) | **Get** /1.0.0/entities/{id}/metadata | Retrieves all the metadata for the entity.
[**GetMetadataEntry**](DefinedEntityApi.md#GetMetadataEntry) | **Get** /1.0.0/entities/{id}/metadata/{metadataId} | Get a single metadata entry.
[**GetMetadataFileContent**](DefinedEntityApi.md#GetMetadataFileContent) | **Get** /1.0.0/entities/{id}/metadata/{metadataId}/content | Download the binary content of a file entry
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
 **resolveEntity** | [**optional.Interface of interface{}**](.md)| The default value is &#39;false&#39;.  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMetadataEntry**
> MetadataEntry CreateMetadataEntry(ctx, entry, id)
Creates a new entry.

Creates a new entry. This operation is allowed only if the user has at least a read access level to the main entity. Additionally file entries require the user to have the &#39;Metadata File Entry: Create/Modify&#39; right. 

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDefinedEntity**
> DeleteDefinedEntity(ctx, id, optional)
Deletes the defined entity with the unique identifier (URN)

Deletes the defined entity with the unique identifier (URN). A multi-stage entity deletion process can achieved using the PreDelete and PostDelete RDE lifecycle hooks. When deleting a defined entity the PreDelete hook is executed first and if invocation fails, deletion is aborted and entity remains unchanged. If PreDelete hook execution succeeds, the entity is moved into IN_DELETION state and PostDelete hook execution is started. If the PostDelete hook succeeds, the entity is deleted. Otherwise, it remains in IN_DELETION state. An entity can always be deleted by setting the invokeHooks parameter to 'false'. 

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
 - **Accept**: *_/_*;version=37.2

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
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefinedEntitiesByEntityType**
> DefinedEntities GetDefinedEntitiesByEntityType(ctx, vendor, nss, version, page, pageSize, optional)
Gets the collection of defined entities for the vCD-defined type with the specified vendor, nss and version.

Gets the collection of defined entities for the vCD-defined type with the specified vendor, nss and version. The version can act as a wildcard. If only '1' is specified as the version, all entity types with a major version of '1' will be matched (e.g. 1.0.0, 1.1.2). If '1.0' is specified, all entity types with a major version of '1' and a minor version of '0' will be included (e.g. 1.0.0, 1.0.1). If the full semver is specified, then no search will be performed. Depending on the requested items per page, and the number of returned entities, one or more metadata summary cursor links will be returned in the headers. In order to retrieve the summaries of all the entities, clients need to fetch each separate cursor and merge the results. 

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
 **metadata** | **optional.String**| Metadata filter for a query.  FIQL format. A single statement is of the form namespace|key operator value. The namespace is optional, has to be separated by &#39;|&#39; from the key and therefore this character is not supported as part of the namespace or key. The value is not optional, however one can omit a value search if &#39;*&#39; is specified. Note that API clients will need to encode these characters accordingly. The framework will try to infer the type of the value in the following order&amp;#58; * if it starts and ends with single unescaped quotes it is a string and the quotes are removed from the beginning and end of the string * else if it parses to a long it is a long * else if it is either &#39;true&#39; or &#39;false&#39;(case insensitive) it is a boolean * else an error id returned  Examples&amp;#58; namespace|com:vmware:key1&#x3D;&#x3D;&amp;#39;42&amp;#39; here the value 42 will be searched as a string entry com:vmware:key2&#x3D;&#x3D;&amp;#39;&amp;#39;string&amp;#39;&amp;#39; here the value &amp;#39;string&amp;#39; (with the quotes) com.key3&#x3D;gt&#x3D;42 here a search for a number entry will be performed, with a value greater than 42 ns|com.key4&#x3D;&#x3D;* here a search for an entry with the namespace &amp;#39;ns&amp;#39; and key &amp;#39;key4&amp;#39; and any value will be performed  | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**DefinedEntities**](DefinedEntities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefinedEntitiesByEntityTypeNoVersionSpecified**
> DefinedEntities GetDefinedEntitiesByEntityTypeNoVersionSpecified(ctx, vendor, nss, page, pageSize, optional)
Gets the collection of defined entities for the vCD-defined type with the specified vendor and nss.

Gets the collection of defined entities for the vCD-defined type with the specified vendor and nss without restrictions on the version. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendor** | **string**|  | 
  **nss** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***DefinedEntityApiGetDefinedEntitiesByEntityTypeNoVersionSpecifiedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedEntityApiGetDefinedEntitiesByEntityTypeNoVersionSpecifiedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **metadata** | **optional.String**| Metadata filter for a query.  FIQL format. A single statement is of the form namespace|key operator value. The namespace is optional, has to be separated by &#39;|&#39; from the key and therefore this character is not supported as part of the namespace or key. The value is not optional, however one can omit a value search if &#39;*&#39; is specified. Note that API clients will need to encode these characters accordingly. The framework will try to infer the type of the value in the following order&amp;#58; * if it starts and ends with single unescaped quotes it is a string and the quotes are removed from the beginning and end of the string * else if it parses to a long it is a long * else if it is either &#39;true&#39; or &#39;false&#39;(case insensitive) it is a boolean * else an error id returned  Examples&amp;#58; namespace|com:vmware:key1&#x3D;&#x3D;&amp;#39;42&amp;#39; here the value 42 will be searched as a string entry com:vmware:key2&#x3D;&#x3D;&amp;#39;&amp;#39;string&amp;#39;&amp;#39; here the value &amp;#39;string&amp;#39; (with the quotes) com.key3&#x3D;gt&#x3D;42 here a search for a number entry will be performed, with a value greater than 42 ns|com.key4&#x3D;&#x3D;* here a search for an entry with the namespace &amp;#39;ns&amp;#39; and key &amp;#39;key4&amp;#39; and any value will be performed  | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**DefinedEntities**](DefinedEntities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefinedEntitiesByInterface**
> DefinedEntities GetDefinedEntitiesByInterface(ctx, vendor, nss, version, page, pageSize, optional)
Gets the collection of defined entities for the vCD-defined interface with the specified vendor, nss and version

Gets the collection of defined entities for the vCD-defined interface with the specified vendor, nss and version. The version can act as a wildcard. If only '1' is specified as the version, all entity types with a major version of '1' will be matched (e.g. 1.0.0, 1.1.2). If '1.0' is specified, all entity types with a major version of '1' and a minor version of '0' will be included (e.g. 1.0.0, 1.0.1). If the full semver is specified, then no search will be performed. Depending on the requested items per page, and the number of returned entities, one or more metadata summary cursor links will be returned in the headers. In order to retrieve the summaries of all the entities, clients need to fetch each separate cursor and merge the results. 

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
 **metadata** | **optional.String**| Metadata filter for a query.  FIQL format. A single statement is of the form namespace|key operator value. The namespace is optional, has to be separated by &#39;|&#39; from the key and therefore this character is not supported as part of the namespace or key. The value is not optional, however one can omit a value search if &#39;*&#39; is specified. Note that API clients will need to encode these characters accordingly. The framework will try to infer the type of the value in the following order&amp;#58; * if it starts and ends with single unescaped quotes it is a string and the quotes are removed from the beginning and end of the string * else if it parses to a long it is a long * else if it is either &#39;true&#39; or &#39;false&#39;(case insensitive) it is a boolean * else an error id returned  Examples&amp;#58; namespace|com:vmware:key1&#x3D;&#x3D;&amp;#39;42&amp;#39; here the value 42 will be searched as a string entry com:vmware:key2&#x3D;&#x3D;&amp;#39;&amp;#39;string&amp;#39;&amp;#39; here the value &amp;#39;string&amp;#39; (with the quotes) com.key3&#x3D;gt&#x3D;42 here a search for a number entry will be performed, with a value greater than 42 ns|com.key4&#x3D;&#x3D;* here a search for an entry with the namespace &amp;#39;ns&amp;#39; and key &amp;#39;key4&amp;#39; and any value will be performed  | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**DefinedEntities**](DefinedEntities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefinedEntity**
> DefinedEntity GetDefinedEntity(ctx, id, optional)
Gets the defined entity with the unique identifier (URN)

Gets the defined entity with the unique identifier (URN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***DefinedEntityApiGetDefinedEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedEntityApiGetDefinedEntityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityVersion** | [**optional.Interface of interface{}**](.md)| Requests that the entity is returned in the specified version of its type. The classification of the requested type must be the same as the current entity type, only the version may differ. The returned entity contents will be converted to the requested type version according to the &#39;required&#39;, &#39;additionalProperties&#39;, and &#39;default&#39; properties of the type versions schema. If the entity is RESOLVED, then the converted entity will be re-validated against the requested type version and an error may be returned if the validation fails. The conversion only affects the returned contents. The entity itself is not modified. To modify the version of the entity permanently, one must update it with a spec of a newer version or perform an upgrade/mass upgrade request.  | 

### Return type

[**DefinedEntity**](DefinedEntity.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetadata**
> MetadataEntries GetMetadata(ctx, page, pageSize, id, optional)
Retrieves all the metadata for the entity.

Retrieves all the metadata for the entity. User can view the entries if user can view the entity. 

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
 - **Accept**: application/json;version=37.2

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetadataFileContent**
> string GetMetadataFileContent(ctx, id, metadataId)
Download the binary content of a file entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| the URN of the entity the entry is attached to. | 
  **metadataId** | **string**| a metadata vcloud id urn | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

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
 - **Accept**: application/json;version=37.2

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
 - **Accept**: application/json;version=37.2

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

