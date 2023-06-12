# \EdgeGatewayIpSecVpnTunnelConnectionPropertiesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIpSecVpnTunnelConnectionProperties**](EdgeGatewayIpSecVpnTunnelConnectionPropertiesApi.md#GetIpSecVpnTunnelConnectionProperties) | **Get** /1.0.0/edgeGateways/{gatewayId}/ipsec/tunnels/{tunnelId}/connectionProperties | Retrieves connection properties for a given IPSec VPN Tunnel configured on an Edge Gateway.
[**UpdateIpSecVpnTunnelConnectionProperties**](EdgeGatewayIpSecVpnTunnelConnectionPropertiesApi.md#UpdateIpSecVpnTunnelConnectionProperties) | **Put** /1.0.0/edgeGateways/{gatewayId}/ipsec/tunnels/{tunnelId}/connectionProperties | Updates the connection properties for a given IPSec VPN Tunnel configured on an Edge Gateway.


# **GetIpSecVpnTunnelConnectionProperties**
> EdgeIpSecVpnTunnelConnectionProperties GetIpSecVpnTunnelConnectionProperties(ctx, gatewayId, tunnelId)
Retrieves connection properties for a given IPSec VPN Tunnel configured on an Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **tunnelId** | **string**|  | 

### Return type

[**EdgeIpSecVpnTunnelConnectionProperties**](EdgeIpSecVpnTunnelConnectionProperties.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpSecVpnTunnelConnectionProperties**
> UpdateIpSecVpnTunnelConnectionProperties(ctx, ipSecVpnTunnelConnectionProperties, gatewayId, tunnelId)
Updates the connection properties for a given IPSec VPN Tunnel configured on an Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSecVpnTunnelConnectionProperties** | [**EdgeIpSecVpnTunnelConnectionProperties**](EdgeIpSecVpnTunnelConnectionProperties.md)|  | 
  **gatewayId** | **string**|  | 
  **tunnelId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

