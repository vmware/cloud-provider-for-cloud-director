# \EdgeGatewayL2VpnTunnelsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateL2VpnTunnel**](EdgeGatewayL2VpnTunnelsApi.md#CreateL2VpnTunnel) | **Post** /1.0.0/edgeGateways/{gatewayId}/l2vpn/tunnels | Creates an L2 VPN tunnel on the Edge Gateway.
[**GetL2VpnTunnels**](EdgeGatewayL2VpnTunnelsApi.md#GetL2VpnTunnels) | **Get** /1.0.0/edgeGateways/{gatewayId}/l2vpn/tunnels | Retrieves all L2 VPN tunnels for a given edge gateway.


# **CreateL2VpnTunnel**
> CreateL2VpnTunnel(ctx, l2VpnTunnel, gatewayId)
Creates an L2 VPN tunnel on the Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **l2VpnTunnel** | [**EdgeL2VpnTunnel**](EdgeL2VpnTunnel.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2VpnTunnels**
> EdgeL2VpnTunnels GetL2VpnTunnels(ctx, gatewayId, optional)
Retrieves all L2 VPN tunnels for a given edge gateway.

Retrieves all L2 VPN tunnels that are configured for an edge gateway. Results can be sorted by only a single parameter. Sorting by combination of parameters (sortAsc=foo&sortDesc=bar) is not allowed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
 **optional** | ***EdgeGatewayL2VpnTunnelsApiGetL2VpnTunnelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayL2VpnTunnelsApiGetL2VpnTunnelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EdgeL2VpnTunnels**](EdgeL2VpnTunnels.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

