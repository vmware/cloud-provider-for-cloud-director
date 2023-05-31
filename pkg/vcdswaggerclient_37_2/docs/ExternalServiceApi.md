# \ExternalServiceApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalService**](ExternalServiceApi.md#CreateExternalService) | **Post** /extensions/api | Create external service.
[**DeleteExternalService**](ExternalServiceApi.md#DeleteExternalService) | **Delete** /extensions/api/{id} | Delete an external service.
[**GetExternalService**](ExternalServiceApi.md#GetExternalService) | **Get** /extensions/api/{id} | Retrieve an external service.
[**GetExternalServices**](ExternalServiceApi.md#GetExternalServices) | **Get** /extensions/api | Query external services.
[**UpdateExternalService**](ExternalServiceApi.md#UpdateExternalService) | **Put** /extensions/api/{id} | Update an external service.


# **CreateExternalService**
> ExternalService CreateExternalService(ctx, service)
Create external service.

Create an external service. Once created, the combination of vendor, name and version cannot be modified and must be unique. Each extension will have its own MQTT topics. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | [**ExternalService**](ExternalService.md)|  | 

### Return type

[**ExternalService**](ExternalService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExternalService**
> DeleteExternalService(ctx, id)
Delete an external service.

Delete an external service. The extension must be disabled or the deletion will fail. 

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

# **GetExternalService**
> ExternalService GetExternalService(ctx, id)
Retrieve an external service.

Extensions, created from other APIs will not be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ExternalService**](ExternalService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalServices**
> ExternalServices GetExternalServices(ctx, page, pageSize, optional)
Query external services.

Query all external services. Extensions, created from other APIs will not be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***ExternalServiceApiGetExternalServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalServiceApiGetExternalServicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**ExternalServices**](ExternalServices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExternalService**
> ExternalService UpdateExternalService(ctx, service, id)
Update an external service.

Update an external service. Vendor, name and version cannot be updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | [**ExternalService**](ExternalService.md)|  | 
  **id** | **string**|  | 

### Return type

[**ExternalService**](ExternalService.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

