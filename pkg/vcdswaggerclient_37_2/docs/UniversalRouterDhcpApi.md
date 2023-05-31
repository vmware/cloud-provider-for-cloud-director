# \UniversalRouterDhcpApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDhcpConfigForUniversalRouter**](UniversalRouterDhcpApi.md#DeleteDhcpConfigForUniversalRouter) | **Delete** /1.0.0/universalRouters/{universalRouterId}/dhcp | Deletes Dhcp configuration of a specific Universal Router
[**GetDhcpConfigForUniversalRouter**](UniversalRouterDhcpApi.md#GetDhcpConfigForUniversalRouter) | **Get** /1.0.0/universalRouters/{universalRouterId}/dhcp | Retrieves Dhcp configuration of a specific Universal Router
[**UpdateDhcpConfigForUniversalRouter**](UniversalRouterDhcpApi.md#UpdateDhcpConfigForUniversalRouter) | **Put** /1.0.0/universalRouters/{universalRouterId}/dhcp | Updates Dhcp configuration for a specific Universal Router


# **DeleteDhcpConfigForUniversalRouter**
> DeleteDhcpConfigForUniversalRouter(ctx, universalRouterId)
Deletes Dhcp configuration of a specific Universal Router

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
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDhcpConfigForUniversalRouter**
> UniversalRouterDhcpConfig GetDhcpConfigForUniversalRouter(ctx, universalRouterId)
Retrieves Dhcp configuration of a specific Universal Router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 

### Return type

[**UniversalRouterDhcpConfig**](UniversalRouterDhcpConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpConfigForUniversalRouter**
> UpdateDhcpConfigForUniversalRouter(ctx, dhcpConfig, universalRouterId)
Updates Dhcp configuration for a specific Universal Router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dhcpConfig** | [**UniversalRouterDhcpConfig**](UniversalRouterDhcpConfig.md)|  | 
  **universalRouterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

