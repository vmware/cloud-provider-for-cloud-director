# \EdgeGatewayL2VpnTunnelApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteL2VpnTunnel**](EdgeGatewayL2VpnTunnelApi.md#DeleteL2VpnTunnel) | **Delete** /1.0.0/edgeGateways/{gatewayId}/l2vpn/tunnels/{tunnelId} | Deletes a specific L2 VPN tunnel for a given edge gateway.
[**GetL2VpnTunnel**](EdgeGatewayL2VpnTunnelApi.md#GetL2VpnTunnel) | **Get** /1.0.0/edgeGateways/{gatewayId}/l2vpn/tunnels/{tunnelId} | Retrieves a specific L2 VPN tunnel for a given edge gateway.
[**GetL2VpnTunnelStatistics**](EdgeGatewayL2VpnTunnelApi.md#GetL2VpnTunnelStatistics) | **Get** /1.0.0/edgeGateways/{gatewayId}/l2vpn/tunnels/{tunnelId}/metrics | Retrieves connection statistics for a given L2 VPN Tunnel configured on an Edge Gateway.
[**GetL2VpnTunnelStatus**](EdgeGatewayL2VpnTunnelApi.md#GetL2VpnTunnelStatus) | **Get** /1.0.0/edgeGateways/{gatewayId}/l2vpn/tunnels/{tunnelId}/status | Retrieves status of a given L2 VPN Tunnel.
[**UpdateL2VpnTunnel**](EdgeGatewayL2VpnTunnelApi.md#UpdateL2VpnTunnel) | **Put** /1.0.0/edgeGateways/{gatewayId}/l2vpn/tunnels/{tunnelId} | Updates a specific L2 VPN tunnel for a given edge gateway.


# **DeleteL2VpnTunnel**
> DeleteL2VpnTunnel(ctx, gatewayId, tunnelId)
Deletes a specific L2 VPN tunnel for a given edge gateway.

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

# **GetL2VpnTunnel**
> EdgeL2VpnTunnel GetL2VpnTunnel(ctx, gatewayId, tunnelId)
Retrieves a specific L2 VPN tunnel for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **tunnelId** | **string**|  | 

### Return type

[**EdgeL2VpnTunnel**](EdgeL2VpnTunnel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2VpnTunnelStatistics**
> EdgeL2VpnTunnelStatistics GetL2VpnTunnelStatistics(ctx, gatewayId, tunnelId)
Retrieves connection statistics for a given L2 VPN Tunnel configured on an Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **tunnelId** | **string**|  | 

### Return type

[**EdgeL2VpnTunnelStatistics**](EdgeL2VpnTunnelStatistics.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2VpnTunnelStatus**
> EdgeL2VpnTunnelStatus GetL2VpnTunnelStatus(ctx, gatewayId, tunnelId)
Retrieves status of a given L2 VPN Tunnel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **tunnelId** | **string**|  | 

### Return type

[**EdgeL2VpnTunnelStatus**](EdgeL2VpnTunnelStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateL2VpnTunnel**
> UpdateL2VpnTunnel(ctx, l2VpnTunnel, gatewayId, tunnelId)
Updates a specific L2 VPN tunnel for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **l2VpnTunnel** | [**EdgeL2VpnTunnel**](EdgeL2VpnTunnel.md)|  | 
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

