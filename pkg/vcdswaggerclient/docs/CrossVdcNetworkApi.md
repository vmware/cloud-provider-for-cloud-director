# \CrossVdcNetworkApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrossVdcNetwork**](CrossVdcNetworkApi.md#DeleteCrossVdcNetwork) | **Delete** /1.0.0/crossVdcNetworks/{crossVdcNetworkId} | Deletes a specific Cross vDC network
[**GetCrossVdcNetwork**](CrossVdcNetworkApi.md#GetCrossVdcNetwork) | **Get** /1.0.0/crossVdcNetworks/{crossVdcNetworkId} | Retrieves a specific Cross vDC network.
[**SyncCrossVdcNetwork**](CrossVdcNetworkApi.md#SyncCrossVdcNetwork) | **Post** /1.0.0/crossVdcNetworks/{crossVdcNetworkId}/sync | Sync/repair a specific Cross vDC network.
[**UpdateCrossVdcNetwork**](CrossVdcNetworkApi.md#UpdateCrossVdcNetwork) | **Put** /1.0.0/crossVdcNetworks/{crossVdcNetworkId} | Updates a specific Cross vDC network.


# **DeleteCrossVdcNetwork**
> DeleteCrossVdcNetwork(ctx, crossVdcNetworkId, optional)
Deletes a specific Cross vDC network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crossVdcNetworkId** | **string**|  | 
 **optional** | ***CrossVdcNetworkApiDeleteCrossVdcNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrossVdcNetworkApiDeleteCrossVdcNetworkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Value \&quot;true\&quot; means to forcefully delete the object that contains other objects even if those objects are in a state that does not allow removal. The default is \&quot;false\&quot;; therefore, objects are not removed if they are not in a state that normally allows removal. Force also implies recursive delete where other contained objects are removed. Errors may be ignored. Invalid value (not true or false) are ignored.  | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCrossVdcNetwork**
> CrossVdcNetwork GetCrossVdcNetwork(ctx, crossVdcNetworkId)
Retrieves a specific Cross vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crossVdcNetworkId** | **string**|  | 

### Return type

[**CrossVdcNetwork**](CrossVdcNetwork.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncCrossVdcNetwork**
> SyncCrossVdcNetwork(ctx, crossVdcNetworkId)
Sync/repair a specific Cross vDC network.

Sync/repair the specific cross vdc network 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crossVdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCrossVdcNetwork**
> UpdateCrossVdcNetwork(ctx, crossVdcNetwork, crossVdcNetworkId)
Updates a specific Cross vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crossVdcNetwork** | [**CrossVdcNetwork**](CrossVdcNetwork.md)|  | 
  **crossVdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

