# \VdcGroupsSettingsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVdcGroupSettings**](VdcGroupsSettingsApi.md#GetVdcGroupSettings) | **Get** /1.0.0/vdcGroups/settings | Retrieves the global vDC groups settings. These settings apply to all vDC Groups in the system and can only be retrieved by the provider.
[**UpdateVdcGroupSettings**](VdcGroupsSettingsApi.md#UpdateVdcGroupSettings) | **Put** /1.0.0/vdcGroups/settings | Updates the global vDC groups settings. These settings apply to all vDC Groups in the system and can only be updated by the provider.


# **GetVdcGroupSettings**
> VdcGroupSettings GetVdcGroupSettings(ctx, )
Retrieves the global vDC groups settings. These settings apply to all vDC Groups in the system and can only be retrieved by the provider.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**VdcGroupSettings**](VdcGroupSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVdcGroupSettings**
> VdcGroupSettings UpdateVdcGroupSettings(ctx, vdcGroupSettings)
Updates the global vDC groups settings. These settings apply to all vDC Groups in the system and can only be updated by the provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupSettings** | [**VdcGroupSettings**](VdcGroupSettings.md)|  | 

### Return type

[**VdcGroupSettings**](VdcGroupSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

