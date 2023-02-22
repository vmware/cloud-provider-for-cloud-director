# \ProxyConfigurationApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProxyConfiguration**](ProxyConfigurationApi.md#CreateProxyConfiguration) | **Post** /1.0.0/proxyConfigurations | Creates a proxy configuration.
[**DeleteProxyConfiguration**](ProxyConfigurationApi.md#DeleteProxyConfiguration) | **Delete** /1.0.0/proxyConfigurations/{id} | Delete a specific proxy configuration.
[**GetProxyConfiguration**](ProxyConfigurationApi.md#GetProxyConfiguration) | **Get** /1.0.0/proxyConfigurations/{id} | Retrieves a specific proxy configuration.
[**QueryProxyConfigurations**](ProxyConfigurationApi.md#QueryProxyConfigurations) | **Get** /1.0.0/proxyConfigurations | Gets a paged list of proxy configurations.
[**UpdateProxyConfiguration**](ProxyConfigurationApi.md#UpdateProxyConfiguration) | **Put** /1.0.0/proxyConfigurations/{id} | Update a specific proxy configuration.


# **CreateProxyConfiguration**
> CreateProxyConfiguration(ctx, proxyConfiguration)
Creates a proxy configuration.

Creates a proxy configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyConfiguration** | [**ProxyConfiguration**](ProxyConfiguration.md)| The new proxy configuration API model. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProxyConfiguration**
> DeleteProxyConfiguration(ctx, id)
Delete a specific proxy configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Proxy Configuration ID URN | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxyConfiguration**
> ProxyConfiguration GetProxyConfiguration(ctx, id)
Retrieves a specific proxy configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Proxy Configuration ID URN | 

### Return type

[**ProxyConfiguration**](ProxyConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryProxyConfigurations**
> ProxyConfigurations QueryProxyConfigurations(ctx, page, pageSize, optional)
Gets a paged list of proxy configurations.

Gets a paged list of proxy configurations. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***ProxyConfigurationApiQueryProxyConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyConfigurationApiQueryProxyConfigurationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**ProxyConfigurations**](ProxyConfigurations.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProxyConfiguration**
> ProxyConfiguration UpdateProxyConfiguration(ctx, updatedProxyConfiguration, id)
Update a specific proxy configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedProxyConfiguration** | [**ProxyConfiguration**](ProxyConfiguration.md)| The updated proxy configuration API model. | 
  **id** | **string**| Proxy Configuration ID URN | 

### Return type

[**ProxyConfiguration**](ProxyConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

