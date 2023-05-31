# \VdcNetworkProfileApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVdcNetworkProfile**](VdcNetworkProfileApi.md#DeleteVdcNetworkProfile) | **Delete** /1.0.0/vdcs/{orgVdcId}/networkProfile | Deletes/Reset a vDC Network Profile.
[**GetVdcNetworkProfile**](VdcNetworkProfileApi.md#GetVdcNetworkProfile) | **Get** /1.0.0/vdcs/{orgVdcId}/networkProfile | Retrieves the vDC Network Profile.
[**UpdateVdcNetworkProfile**](VdcNetworkProfileApi.md#UpdateVdcNetworkProfile) | **Put** /1.0.0/vdcs/{orgVdcId}/networkProfile | Updates the vDC Network Profile.


# **DeleteVdcNetworkProfile**
> DeleteVdcNetworkProfile(ctx, orgVdcId)
Deletes/Reset a vDC Network Profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgVdcId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcNetworkProfile**
> VdcNetworkProfile GetVdcNetworkProfile(ctx, orgVdcId)
Retrieves the vDC Network Profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgVdcId** | **string**|  | 

### Return type

[**VdcNetworkProfile**](VdcNetworkProfile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVdcNetworkProfile**
> UpdateVdcNetworkProfile(ctx, vdcNetworkProfile, orgVdcId)
Updates the vDC Network Profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkProfile** | [**VdcNetworkProfile**](VdcNetworkProfile.md)|  | 
  **orgVdcId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

