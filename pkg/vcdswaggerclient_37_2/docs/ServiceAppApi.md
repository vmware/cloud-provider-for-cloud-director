# \ServiceAppApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceApp**](ServiceAppApi.md#DeleteServiceApp) | **Delete** /1.0.0/serviceApps/{serviceAppId} | Deletes a specific VMware service application.
[**GetServiceApp**](ServiceAppApi.md#GetServiceApp) | **Get** /1.0.0/serviceApps/{serviceAppId} | Retrieves a specific VMware service application
[**UpdateServiceApp**](ServiceAppApi.md#UpdateServiceApp) | **Put** /1.0.0/serviceApps/{serviceAppId} | Updates a specific VMware service application.


# **DeleteServiceApp**
> DeleteServiceApp(ctx, serviceAppId)
Deletes a specific VMware service application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceAppId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceApp**
> ServiceApp GetServiceApp(ctx, serviceAppId)
Retrieves a specific VMware service application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceAppId** | **string**|  | 

### Return type

[**ServiceApp**](ServiceApp.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceApp**
> ServiceApp UpdateServiceApp(ctx, serviceApp, serviceAppId)
Updates a specific VMware service application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceApp** | [**ServiceApp**](ServiceApp.md)|  | 
  **serviceAppId** | **string**|  | 

### Return type

[**ServiceApp**](ServiceApp.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

