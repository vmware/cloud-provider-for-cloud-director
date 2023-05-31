# \EdgeGatewayIpSecVpnTunnelApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIpSecVpnTunnel**](EdgeGatewayIpSecVpnTunnelApi.md#DeleteIpSecVpnTunnel) | **Delete** /1.0.0/edgeGateways/{gatewayId}/ipsec/tunnels/{tunnelId} | Deletes a specific IPSec tunnel for a given edge gateway.
[**GetIpSecVpnTunnel**](EdgeGatewayIpSecVpnTunnelApi.md#GetIpSecVpnTunnel) | **Get** /1.0.0/edgeGateways/{gatewayId}/ipsec/tunnels/{tunnelId} | Retrieves a specific IPSec tunnel for a given edge gateway.
[**GetIpSecVpnTunnelDefaultConnectionProperties**](EdgeGatewayIpSecVpnTunnelApi.md#GetIpSecVpnTunnelDefaultConnectionProperties) | **Get** /1.0.0/edgeGateways/{gatewayId}/ipsec/tunnels/defaultConnectionProperties | Retrieves the default connection properties that are used for a given IPSec Tunnel in NSX-T when default is set or no security type is specified.
[**GetIpSecVpnTunnelStatistics**](EdgeGatewayIpSecVpnTunnelApi.md#GetIpSecVpnTunnelStatistics) | **Get** /1.0.0/edgeGateways/{gatewayId}/ipsec/tunnels/{tunnelId}/statistics | Retrieves connection statistics for a given IPSec VPN Tunnel configured on an Edge Gateway.
[**GetIpSecVpnTunnelStatus**](EdgeGatewayIpSecVpnTunnelApi.md#GetIpSecVpnTunnelStatus) | **Get** /1.0.0/edgeGateways/{gatewayId}/ipsec/tunnels/{tunnelId}/status | Retrieves status of a given IPSec VPN Tunnel configured on an Edge Gateway.
[**UpdateIpSecVpnTunnel**](EdgeGatewayIpSecVpnTunnelApi.md#UpdateIpSecVpnTunnel) | **Put** /1.0.0/edgeGateways/{gatewayId}/ipsec/tunnels/{tunnelId} | Updates a specific IPSec tunnel for a given edge gateway.


# **DeleteIpSecVpnTunnel**
> DeleteIpSecVpnTunnel(ctx, gatewayId, tunnelId)
Deletes a specific IPSec tunnel for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **tunnelId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSecVpnTunnel**
> EdgeIpSecVpnTunnel GetIpSecVpnTunnel(ctx, gatewayId, tunnelId)
Retrieves a specific IPSec tunnel for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **tunnelId** | **string**|  | 

### Return type

[**EdgeIpSecVpnTunnel**](EdgeIpSecVpnTunnel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSecVpnTunnelDefaultConnectionProperties**
> EdgeIpSecVpnTunnelConnectionProperties GetIpSecVpnTunnelDefaultConnectionProperties(ctx, gatewayId)
Retrieves the default connection properties that are used for a given IPSec Tunnel in NSX-T when default is set or no security type is specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**EdgeIpSecVpnTunnelConnectionProperties**](EdgeIpSecVpnTunnelConnectionProperties.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSecVpnTunnelStatistics**
> EdgeIpSecVpnTunnelStatistics GetIpSecVpnTunnelStatistics(ctx, gatewayId, tunnelId)
Retrieves connection statistics for a given IPSec VPN Tunnel configured on an Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **tunnelId** | **string**|  | 

### Return type

[**EdgeIpSecVpnTunnelStatistics**](EdgeIpSecVpnTunnelStatistics.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSecVpnTunnelStatus**
> EdgeIpSecVpnTunnelStatus GetIpSecVpnTunnelStatus(ctx, gatewayId, tunnelId)
Retrieves status of a given IPSec VPN Tunnel configured on an Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **tunnelId** | **string**|  | 

### Return type

[**EdgeIpSecVpnTunnelStatus**](EdgeIpSecVpnTunnelStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpSecVpnTunnel**
> UpdateIpSecVpnTunnel(ctx, ipsecVpnTunnel, gatewayId, tunnelId)
Updates a specific IPSec tunnel for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnTunnel** | [**EdgeIpSecVpnTunnel**](EdgeIpSecVpnTunnel.md)|  | 
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

