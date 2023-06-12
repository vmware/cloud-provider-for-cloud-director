# \EdgeGatewayDhcpApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDhcpForwarder**](EdgeGatewayDhcpApi.md#GetDhcpForwarder) | **Get** /1.0.0/edgeGateways/{gatewayId}/dhcpForwarder | Retrieves DHCP Forwarder configuration on an Edge Gateway.
[**UpdateDhcpForwarder**](EdgeGatewayDhcpApi.md#UpdateDhcpForwarder) | **Put** /1.0.0/edgeGateways/{gatewayId}/dhcpForwarder | Update DHCP Forwarder configuration on an Edge Gateway.


# **GetDhcpForwarder**
> EdgeDhcpForwarder GetDhcpForwarder(ctx, gatewayId)
Retrieves DHCP Forwarder configuration on an Edge Gateway.

Retrieves the DHCP Forwarder configuration on an Edge Gateway. A routed Org vDC network connected to this edge can choose to configure its DHCP configuration in RELAY mode which will use this DHCP forwarder. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**EdgeDhcpForwarder**](EdgeDhcpForwarder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpForwarder**
> UpdateDhcpForwarder(ctx, dhcpForwarder, gatewayId)
Update DHCP Forwarder configuration on an Edge Gateway.

Updates the DHCP Forwarder configuration on an Edge Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dhcpForwarder** | [**EdgeDhcpForwarder**](EdgeDhcpForwarder.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

