# \DefinedEntityTypeApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDefinedEntityType**](DefinedEntityTypeApi.md#CreateDefinedEntityType) | **Post** /1.0.0/entityTypes/ | Creates a defined entity type.
[**DeleteDefinedEntityType**](DefinedEntityTypeApi.md#DeleteDefinedEntityType) | **Delete** /1.0.0/entityTypes/{id} | Deletes the entity type with the unique identifier (URN)
[**GetDefinedEntityType**](DefinedEntityTypeApi.md#GetDefinedEntityType) | **Get** /1.0.0/entityTypes/{id} | Gets the entity type with the unique identifier (URN)
[**GetDefinedEntityTypes**](DefinedEntityTypeApi.md#GetDefinedEntityTypes) | **Get** /1.0.0/entityTypes/ | Gets the collection of entity types defined in the vCD instance. Allows collection refinement through traditional FIQL-based filtering
[**UpdateDefinedEntityType**](DefinedEntityTypeApi.md#UpdateDefinedEntityType) | **Put** /1.0.0/entityTypes/{id} | Updates the entity type with the unique identifier (URN)


# **CreateDefinedEntityType**
> DefinedEntityType CreateDefinedEntityType(ctx, definition)
Creates a defined entity type.

Creates a defined entity type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definition** | [**DefinedEntityType**](DefinedEntityType.md)|  | 

### Return type

[**DefinedEntityType**](DefinedEntityType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDefinedEntityType**
> DeleteDefinedEntityType(ctx, id)
Deletes the entity type with the unique identifier (URN)

Deletes the entity type with the unique identifier (URN)

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

# **GetDefinedEntityType**
> DefinedEntityType GetDefinedEntityType(ctx, id)
Gets the entity type with the unique identifier (URN)

Gets the entity type with the unique identifier (URN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DefinedEntityType**](DefinedEntityType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefinedEntityTypes**
> DefinedEntityTypes GetDefinedEntityTypes(ctx, page, pageSize, optional)
Gets the collection of entity types defined in the vCD instance. Allows collection refinement through traditional FIQL-based filtering

Gets the collection of entity types defined in the vCD instance. Allows collection refinement through traditional FIQL-based filtering

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***DefinedEntityTypeApiGetDefinedEntityTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedEntityTypeApiGetDefinedEntityTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**DefinedEntityTypes**](DefinedEntityTypes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDefinedEntityType**
> DefinedEntityType UpdateDefinedEntityType(ctx, definition, id)
Updates the entity type with the unique identifier (URN)

Updates the entity type with the unique identifier (URN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definition** | [**DefinedEntityType**](DefinedEntityType.md)|  | 
  **id** | **string**|  | 

### Return type

[**DefinedEntityType**](DefinedEntityType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

