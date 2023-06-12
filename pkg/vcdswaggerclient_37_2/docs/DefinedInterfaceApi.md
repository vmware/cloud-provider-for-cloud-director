# \DefinedInterfaceApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInterface**](DefinedInterfaceApi.md#CreateInterface) | **Post** /1.0.0/interfaces | Creates a defined interface.
[**DeleteInterface**](DefinedInterfaceApi.md#DeleteInterface) | **Delete** /1.0.0/interfaces/{id} | Deletes the interface with the unique identifier (URN)
[**GetInterface**](DefinedInterfaceApi.md#GetInterface) | **Get** /1.0.0/interfaces/{id} | Gets the interface with the unique identifier (URN)
[**QueryInterfaces**](DefinedInterfaceApi.md#QueryInterfaces) | **Get** /1.0.0/interfaces | Gets the collection of interfaces defined in the vCD instance. Allows collection refinement through traditional FIQL-based filtering
[**UpdateInterface**](DefinedInterfaceApi.md#UpdateInterface) | **Put** /1.0.0/interfaces/{id} | Updates the interface with the unique identifier (URN)


# **CreateInterface**
> DefinedInterface CreateInterface(ctx, definedInterface)
Creates a defined interface.

Creates a defined interface. The version must follow semantic versioning rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definedInterface** | [**DefinedInterface**](DefinedInterface.md)|  | 

### Return type

[**DefinedInterface**](DefinedInterface.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInterface**
> DeleteInterface(ctx, id)
Deletes the interface with the unique identifier (URN)

Deletes the interface with the unique identifier (URN) 

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

# **GetInterface**
> DefinedInterface GetInterface(ctx, id)
Gets the interface with the unique identifier (URN)

Gets the interface with the unique identifier (URN) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DefinedInterface**](DefinedInterface.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryInterfaces**
> DefinedInterfaces QueryInterfaces(ctx, page, pageSize, optional)
Gets the collection of interfaces defined in the vCD instance. Allows collection refinement through traditional FIQL-based filtering

Gets the collection of interfaces defined in the vCD instance. Allows collection refinement through traditional FIQL-based filtering

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***DefinedInterfaceApiQueryInterfacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedInterfaceApiQueryInterfacesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**DefinedInterfaces**](DefinedInterfaces.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInterface**
> DefinedInterface UpdateInterface(ctx, definedInterface, id)
Updates the interface with the unique identifier (URN)

Updates the interface with the unique identifier (URN) The version must follow semantic versioning rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definedInterface** | [**DefinedInterface**](DefinedInterface.md)|  | 
  **id** | **string**|  | 

### Return type

[**DefinedInterface**](DefinedInterface.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

