# \VrealizeBrowseApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseSdkObjects**](VrealizeBrowseApi.md#BrowseSdkObjects) | **Get** /vro/servers/{vroId}/sdkObjects/{target}/{browsePaths:.+} | Browse remote vRealize Orchestrator inventory
[**GetPluginList**](VrealizeBrowseApi.md#GetPluginList) | **Get** /vro/servers/{vroId}/entityTypes | Browse remote vRealize Orchestrator inventory
[**GetPluginTypeList**](VrealizeBrowseApi.md#GetPluginTypeList) | **Get** /vro/servers/{vroId}/entityTypes/{pluginName} | Browse remote vRealize Orchestrator inventory
[**GetRemoteInventory**](VrealizeBrowseApi.md#GetRemoteInventory) | **Get** /vro/servers/{vroId}/inventory{any:.*} | Browse remote vRealize Orchestrator inventory
[**SearchSdkObjects**](VrealizeBrowseApi.md#SearchSdkObjects) | **Get** /vro/servers/{vroId}/sdkObjects/{target} | Browse remote vRealize Orchestrator inventory


# **BrowseSdkObjects**
> VroRemoteItems BrowseSdkObjects(ctx, vroId, target, browsePaths)
Browse remote vRealize Orchestrator inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vroId** | **string**| The ID of the server to browse inventory items on | 
  **target** | **string**| The &#39;plugin[:type]&#39; to serve as the starting point for the browsing | 
  **browsePaths** | **string**| Inventory search path to identify the VRO inventory node to get contents of | 

### Return type

[**VroRemoteItems**](VroRemoteItems.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPluginList**
> VroRemotePluginItems GetPluginList(ctx, vroId)
Browse remote vRealize Orchestrator inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vroId** | **string**| The ID of the server to browse inventory items on | 

### Return type

[**VroRemotePluginItems**](VroRemotePluginItems.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPluginTypeList**
> VroRemotePluginTypes GetPluginTypeList(ctx, vroId, pluginName)
Browse remote vRealize Orchestrator inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vroId** | **string**| The ID of the server to browse inventory items on | 
  **pluginName** | **string**| Inventory search path to identify the VRO inventory node to get contents of | 

### Return type

[**VroRemotePluginTypes**](VroRemotePluginTypes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRemoteInventory**
> VroRemoteInventoryItem GetRemoteInventory(ctx, vroId, any)
Browse remote vRealize Orchestrator inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vroId** | **string**| The ID of the server to browse inventory items on | 
  **any** | **string**| Inventory search path to identify the VRO inventory node to get contents of | 

### Return type

[**VroRemoteInventoryItem**](VroRemoteInventoryItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSdkObjects**
> VroRemoteItems SearchSdkObjects(ctx, vroId, target, optional)
Browse remote vRealize Orchestrator inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vroId** | **string**| The ID of the server to browse inventory items on | 
  **target** | **string**| The &#39;plugin[:type]&#39; to serve as the starting point for the browsing | 
 **optional** | ***VrealizeBrowseApiSearchSdkObjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VrealizeBrowseApiSearchSdkObjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VroRemoteItems**](VroRemoteItems.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

