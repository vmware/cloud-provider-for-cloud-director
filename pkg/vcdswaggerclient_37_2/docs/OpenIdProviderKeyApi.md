# \OpenIdProviderKeyApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOpenIdProviderKey**](OpenIdProviderKeyApi.md#CreateOpenIdProviderKey) | **Post** /1.0.0/openIdProvider/keys | Adds a new OpenID Provider signing key to list of available keys.
[**DeleteOpenIdProviderKey**](OpenIdProviderKeyApi.md#DeleteOpenIdProviderKey) | **Delete** /1.0.0/openIdProvider/keys/{openIdProviderKeyId} | Delete the specified OpenID Provider key.
[**GetOpenIdProviderKey**](OpenIdProviderKeyApi.md#GetOpenIdProviderKey) | **Get** /1.0.0/openIdProvider/keys/{openIdProviderKeyId} | Get the specified OpenID Provider key.
[**QueryOpenIdProviderKeys**](OpenIdProviderKeyApi.md#QueryOpenIdProviderKeys) | **Get** /1.0.0/openIdProvider/keys | Query the list of all configured OpenID Provider keys.
[**UpdateOpenIdProviderKey**](OpenIdProviderKeyApi.md#UpdateOpenIdProviderKey) | **Put** /1.0.0/openIdProvider/keys/{openIdProviderKeyId} | Updates the specified OpenID Provider key entry.


# **CreateOpenIdProviderKey**
> OpenIdProviderKey CreateOpenIdProviderKey(ctx, newOpenIdProviderKey)
Adds a new OpenID Provider signing key to list of available keys.

Adds the provided private and public key pairs to the list of configured signing keys. This key is NOT automatically made the active signing key. The existing active key will continue to be used. <P>The provided keys, in addition to being of one of acceptable types, will be validated to confirm that they are a cryptographic pair and that they conform to the minimum key size in the SSL settings for the product. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newOpenIdProviderKey** | [**OpenIdProviderKey**](OpenIdProviderKey.md)|  | 

### Return type

[**OpenIdProviderKey**](openIdProviderKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOpenIdProviderKey**
> DeleteOpenIdProviderKey(ctx, openIdProviderKeyId)
Delete the specified OpenID Provider key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **openIdProviderKeyId** | **string**| OpenID Provider key URN | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpenIdProviderKey**
> OpenIdProviderKey GetOpenIdProviderKey(ctx, openIdProviderKeyId)
Get the specified OpenID Provider key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **openIdProviderKeyId** | **string**| OpenID Provider key URN | 

### Return type

[**OpenIdProviderKey**](openIdProviderKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryOpenIdProviderKeys**
> OpenIdProviderKeys QueryOpenIdProviderKeys(ctx, page, pageSize, optional)
Query the list of all configured OpenID Provider keys.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***OpenIdProviderKeyApiQueryOpenIdProviderKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenIdProviderKeyApiQueryOpenIdProviderKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**OpenIdProviderKeys**](openIdProviderKeys.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOpenIdProviderKey**
> OpenIdProviderKey UpdateOpenIdProviderKey(ctx, updatedOpenIdProviderKey, openIdProviderKeyId)
Updates the specified OpenID Provider key entry.

The description of the specified key entry can be updated. Attempt to modify any other field will result in a bad request error.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedOpenIdProviderKey** | [**OpenIdProviderKey**](OpenIdProviderKey.md)| The updated OpenID Provider key. | 
  **openIdProviderKeyId** | **string**| OpenID Provider key URN | 

### Return type

[**OpenIdProviderKey**](openIdProviderKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

