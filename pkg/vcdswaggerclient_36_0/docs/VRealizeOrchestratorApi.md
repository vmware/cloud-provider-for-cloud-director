# \VRealizeOrchestratorApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegisteredVRO**](VRealizeOrchestratorApi.md#GetRegisteredVRO) | **Get** /vro/servers/{id} | Retrieves service specific metadata for a vRealize Orchestrator
[**PatchService**](VRealizeOrchestratorApi.md#PatchService) | **Patch** /vro/servers/{id} | Updates service specific metadata for a vRealize Orchestrator
[**Unregister**](VRealizeOrchestratorApi.md#Unregister) | **Delete** /vro/servers/{id} | Unregisters a vRealize Orchestrator endpoint from vCloud Director
[**UpdateService**](VRealizeOrchestratorApi.md#UpdateService) | **Put** /vro/servers/{id} | Updates service specific metadata for a vRealize Orchestrator


# **GetRegisteredVRO**
> VroServiceInfo GetRegisteredVRO(ctx, id)
Retrieves service specific metadata for a vRealize Orchestrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**VroServiceInfo**](VROServiceInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchService**
> VroServiceInfo PatchService(ctx, body, id)
Updates service specific metadata for a vRealize Orchestrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VroServiceInfo**](VroServiceInfo.md)|  | 
  **id** | **string**|  | 

### Return type

[**VroServiceInfo**](VROServiceInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/merge-patch+json, application/json-patch+json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unregister**
> Unregister(ctx, id)
Unregisters a vRealize Orchestrator endpoint from vCloud Director

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

# **UpdateService**
> VroServiceInfo UpdateService(ctx, body, id)
Updates service specific metadata for a vRealize Orchestrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VroServiceInfo**](VroServiceInfo.md)|  | 
  **id** | **string**|  | 

### Return type

[**VroServiceInfo**](VROServiceInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

