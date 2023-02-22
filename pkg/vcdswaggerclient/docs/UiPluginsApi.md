# \UiPluginsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUiPlugin**](UiPluginsApi.md#AddUiPlugin) | **Post** /extensions/ui | Adds plugin metadata for a new UI Extension
[**GetUiPlugins**](UiPluginsApi.md#GetUiPlugins) | **Get** /extensions/ui | Get a list of all UI Extensions


# **AddUiPlugin**
> UiPluginMetadataResponse AddUiPlugin(ctx, body)
Adds plugin metadata for a new UI Extension

Creates a new UI extension and sets the provided plugin metadata for it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UiPluginMetadata**](UiPluginMetadata.md)|  | 

### Return type

[**UiPluginMetadataResponse**](UiPluginMetadataResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUiPlugins**
> []UiPluginMetadataResponse GetUiPlugins(ctx, )
Get a list of all UI Extensions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UiPluginMetadataResponse**](UiPluginMetadataResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

