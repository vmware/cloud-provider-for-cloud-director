# \UniversalRouterDnsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDnsConfigForRouter**](UniversalRouterDnsApi.md#DeleteDnsConfigForRouter) | **Delete** /1.0.0/universalRouters/{universalRouterId}/dns | Deletes dns configuration of a universal router
[**GetDnsConfigForRouter**](UniversalRouterDnsApi.md#GetDnsConfigForRouter) | **Get** /1.0.0/universalRouters/{universalRouterId}/dns | Retrieves dns configuration of a universal router
[**UpdateDnsConfigForRouter**](UniversalRouterDnsApi.md#UpdateDnsConfigForRouter) | **Put** /1.0.0/universalRouters/{universalRouterId}/dns | Updates dns configuration of a universal Router


# **DeleteDnsConfigForRouter**
> DeleteDnsConfigForRouter(ctx, universalRouterId)
Deletes dns configuration of a universal router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDnsConfigForRouter**
> RouterDnsConfig GetDnsConfigForRouter(ctx, universalRouterId)
Retrieves dns configuration of a universal router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 

### Return type

[**RouterDnsConfig**](RouterDnsConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDnsConfigForRouter**
> UpdateDnsConfigForRouter(ctx, dnsConfig, universalRouterId)
Updates dns configuration of a universal Router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dnsConfig** | [**RouterDnsConfig**](RouterDnsConfig.md)|  | 
  **universalRouterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

