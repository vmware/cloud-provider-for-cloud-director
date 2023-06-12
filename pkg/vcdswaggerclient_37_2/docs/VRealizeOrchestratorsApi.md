# \VRealizeOrchestratorsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverVroVcenter**](VRealizeOrchestratorsApi.md#DiscoverVroVcenter) | **Post** /vro/servers/discovery | Discover vCenter to work with the provided vRealize Orchestrator service
[**GetRegisteredVROs**](VRealizeOrchestratorsApi.md#GetRegisteredVROs) | **Get** /vro/servers | Get a list of registered vRealize Orchestrator services
[**Register**](VRealizeOrchestratorsApi.md#Register) | **Post** /vro/servers | Register a vRealize Orchestrator endpoint with vCloud Director


# **DiscoverVroVcenter**
> DiscoverVroVcenter(ctx, body)
Discover vCenter to work with the provided vRealize Orchestrator service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VroServiceInfo**](VroServiceInfo.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegisteredVROs**
> []VroServiceInfo GetRegisteredVROs(ctx, )
Get a list of registered vRealize Orchestrator services

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]VroServiceInfo**](VROServiceInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Register**
> VroServiceInfo Register(ctx, body)
Register a vRealize Orchestrator endpoint with vCloud Director

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VroServiceInfo**](VroServiceInfo.md)|  | 

### Return type

[**VroServiceInfo**](VROServiceInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

