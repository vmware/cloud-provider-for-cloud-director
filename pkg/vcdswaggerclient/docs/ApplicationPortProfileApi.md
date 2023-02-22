# \ApplicationPortProfileApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplicationPortProfile**](ApplicationPortProfileApi.md#DeleteApplicationPortProfile) | **Delete** /1.0.0/applicationPortProfiles/{applicationPortProfileId} | Deletes a specific user-defined Application Port Profile
[**GetApplicationPortProfile**](ApplicationPortProfileApi.md#GetApplicationPortProfile) | **Get** /1.0.0/applicationPortProfiles/{applicationPortProfileId} | Retrieves a specific user-defined Application Port Profile
[**UpdateApplicationPortProfile**](ApplicationPortProfileApi.md#UpdateApplicationPortProfile) | **Put** /1.0.0/applicationPortProfiles/{applicationPortProfileId} | Updates a specific user-defined Application Port Profile


# **DeleteApplicationPortProfile**
> DeleteApplicationPortProfile(ctx, applicationPortProfileId)
Deletes a specific user-defined Application Port Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationPortProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationPortProfile**
> ApplicationPortProfile GetApplicationPortProfile(ctx, applicationPortProfileId)
Retrieves a specific user-defined Application Port Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationPortProfileId** | **string**|  | 

### Return type

[**ApplicationPortProfile**](ApplicationPortProfile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationPortProfile**
> UpdateApplicationPortProfile(ctx, applicationPortProfile, applicationPortProfileId)
Updates a specific user-defined Application Port Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationPortProfile** | [**ApplicationPortProfile**](ApplicationPortProfile.md)|  | 
  **applicationPortProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

