# \DeviceAuthorizationApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DenyDeviceAuthorization**](DeviceAuthorizationApi.md#DenyDeviceAuthorization) | **Post** /1.0.0/deviceLookup/deny | Denies a device authorization request
[**FindDeviceAuthorizationRequest**](DeviceAuthorizationApi.md#FindDeviceAuthorizationRequest) | **Post** /1.0.0/deviceLookup | Looks up a device authorization request for approval
[**GrantDeviceAuthorization**](DeviceAuthorizationApi.md#GrantDeviceAuthorization) | **Post** /1.0.0/deviceLookup/grant | Grants a device authorization request


# **DenyDeviceAuthorization**
> DenyDeviceAuthorization(ctx, userCode)
Denies a device authorization request

A device's request for access on behalf of a service account, as identified by the specified user code, is denied 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userCode** | [**DeviceInfo**](DeviceInfo.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindDeviceAuthorizationRequest**
> FindDeviceAuthorizationRequest(ctx, userCode)
Looks up a device authorization request for approval

Looks up a service account identified by the specified user code for processing its authorization request 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userCode** | [**DeviceInfo**](DeviceInfo.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrantDeviceAuthorization**
> GrantDeviceAuthorization(ctx, userCode)
Grants a device authorization request

Grants access to service account identified by the specified user code. Subsequent polling by the device will result in access token to be transmitted as per device code flow specification 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userCode** | [**DeviceInfo**](DeviceInfo.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

