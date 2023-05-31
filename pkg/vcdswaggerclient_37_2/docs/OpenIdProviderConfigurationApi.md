# \OpenIdProviderConfigurationApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOpenIdProviderConfiguration**](OpenIdProviderConfigurationApi.md#GetOpenIdProviderConfiguration) | **Get** /1.0.0/openIdProvider | Get the OpenID Provider configuration.
[**UpdateOpenIdProviderConfiguration**](OpenIdProviderConfigurationApi.md#UpdateOpenIdProviderConfiguration) | **Put** /1.0.0/openIdProvider | Updates the OpenID Provider configuration. 


# **GetOpenIdProviderConfiguration**
> OpenIdProviderConfiguration GetOpenIdProviderConfiguration(ctx, )
Get the OpenID Provider configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OpenIdProviderConfiguration**](OpenIdProviderConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOpenIdProviderConfiguration**
> OpenIdProviderConfiguration UpdateOpenIdProviderConfiguration(ctx, updatedOpenIdProviderConfiguration)
Updates the OpenID Provider configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedOpenIdProviderConfiguration** | [**OpenIdProviderConfiguration**](OpenIdProviderConfiguration.md)| The updated OpenID Provider configuration. | 

### Return type

[**OpenIdProviderConfiguration**](OpenIdProviderConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

