# \UiPluginApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUiPlugin**](UiPluginApi.md#DeleteUiPlugin) | **Delete** /extensions/ui/{id} | Delete system level logo
[**GetExtensionPointSummary**](UiPluginApi.md#GetExtensionPointSummary) | **Get** /extensions/ui/extensionPoints | Retrieves a map of extension Points and an ordered list of items registered with that extension point
[**GetUiPlugin**](UiPluginApi.md#GetUiPlugin) | **Get** /extensions/ui/{id} | Retrieves extension specific plugin metadata
[**PutExtensionPointSummary**](UiPluginApi.md#PutExtensionPointSummary) | **Put** /extensions/ui/extensionPoints | Customizes the order and enables/disables extension Points
[**PutUiPlugin**](UiPluginApi.md#PutUiPlugin) | **Put** /extensions/ui/{id} | Updates extension specific plugin&#39;s metadata


# **DeleteUiPlugin**
> DeleteUiPlugin(ctx, id)
Delete system level logo

Delete the system level logo, forcing the get method to return the vCloud Director default logo. 

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

# **GetExtensionPointSummary**
> ExtensionPointSummary GetExtensionPointSummary(ctx, )
Retrieves a map of extension Points and an ordered list of items registered with that extension point

Retrieves a map of extension Points and an ordered list of items registered with that extension point

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ExtensionPointSummary**](ExtensionPointSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUiPlugin**
> UiPluginMetadataResponse GetUiPlugin(ctx, id)
Retrieves extension specific plugin metadata

Retrieve the plugin metadata for this extension 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**UiPluginMetadataResponse**](UiPluginMetadataResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutExtensionPointSummary**
> ExtensionPointSummary PutExtensionPointSummary(ctx, extensionPointSummaryBody)
Customizes the order and enables/disables extension Points

Customizes the order and enables/disables extension Points 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionPointSummaryBody** | [**ExtensionPointSummary**](ExtensionPointSummary.md)|  | 

### Return type

[**ExtensionPointSummary**](ExtensionPointSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUiPlugin**
> UiPluginMetadataResponse PutUiPlugin(ctx, pluginMetadataBody, id)
Updates extension specific plugin's metadata

Update the plugin metadata for this extension clobbering existing information and returns the updated plugin metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginMetadataBody** | [**UiPluginMetadata**](UiPluginMetadata.md)|  | 
  **id** | **string**|  | 

### Return type

[**UiPluginMetadataResponse**](UiPluginMetadataResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

