# \EdgeGatewayDnsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEdgeGatewayDns**](EdgeGatewayDnsApi.md#DeleteEdgeGatewayDns) | **Delete** /1.0.0/edgeGateways/{gatewayId}/dns | Deletes DNS configuration of the edge gateway.
[**GetEdgeGatewayDns**](EdgeGatewayDnsApi.md#GetEdgeGatewayDns) | **Get** /1.0.0/edgeGateways/{gatewayId}/dns | Retrieves DNS configuration of the edge gateway.
[**UpdateEdgeGatewayDns**](EdgeGatewayDnsApi.md#UpdateEdgeGatewayDns) | **Put** /1.0.0/edgeGateways/{gatewayId}/dns | Updates DNS configuration of the edge gateway.


# **DeleteEdgeGatewayDns**
> DeleteEdgeGatewayDns(ctx, gatewayId)
Deletes DNS configuration of the edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeGatewayDns**
> EdgeDnsConfig GetEdgeGatewayDns(ctx, gatewayId)
Retrieves DNS configuration of the edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**EdgeDnsConfig**](EdgeDnsConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEdgeGatewayDns**
> UpdateEdgeGatewayDns(ctx, dnsConfig, gatewayId)
Updates DNS configuration of the edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dnsConfig** | [**EdgeDnsConfig**](EdgeDnsConfig.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

