# \UiPluginResourceApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUiPluginResource**](UiPluginResourceApi.md#DeleteUiPluginResource) | **Delete** /extensions/ui/{id}/plugin | Delete the plugin for this extension
[**UploadUiPluginResource**](UiPluginResourceApi.md#UploadUiPluginResource) | **Post** /extensions/ui/{id}/plugin | Upload the plugin for this extension


# **DeleteUiPluginResource**
> DeleteUiPluginResource(ctx, id)
Delete the plugin for this extension

Deletes the actual plugin for this extension 

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

# **UploadUiPluginResource**
> UploadUiPluginResource(ctx, pluginUploadSpec, id)
Upload the plugin for this extension

Initiates an upload for the plugin for this extension using the Transfer service A unique transfer service URL is returned where the plugin may be uploaded. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginUploadSpec** | [**UploadSpec**](UploadSpec.md)|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

