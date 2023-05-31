# \ConfigurationsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig**](ConfigurationsApi.md#GetConfig) | **Get** /1.0.0/site/configurations/{urn} | Get the current setting for the specified configuration property.
[**GetConfigs**](ConfigurationsApi.md#GetConfigs) | **Get** /1.0.0/site/configurations | Returns all configuration properties. At present this will always return an EMPTY list.
[**SetConfig**](ConfigurationsApi.md#SetConfig) | **Put** /1.0.0/site/configurations/{urn} | Sets a configuration property to the provided value.


# **GetConfig**
> Config GetConfig(ctx, urn)
Get the current setting for the specified configuration property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urn** | **string**| URN containing the name of the configuration property such as (urn:vcloud:configuration:&lt;configuration name&gt;) | 

### Return type

[**Config**](Config.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigs**
> Configs GetConfigs(ctx, page, pageSize, optional)
Returns all configuration properties. At present this will always return an EMPTY list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***ConfigurationsApiGetConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigurationsApiGetConfigsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Configs**](Configs.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetConfig**
> Config SetConfig(ctx, newConfigValue, urn)
Sets a configuration property to the provided value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newConfigValue** | [**Config**](Config.md)| The configuration with the desired value. | 
  **urn** | **string**| URN containing the name of the configuration property such as (urn:vcloud:configuration:&lt;configuration name&gt;) | 

### Return type

[**Config**](Config.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

