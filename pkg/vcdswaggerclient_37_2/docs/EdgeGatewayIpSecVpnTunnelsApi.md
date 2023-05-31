# \EdgeGatewayIpSecVpnTunnelsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpSecVpnTunnel**](EdgeGatewayIpSecVpnTunnelsApi.md#CreateIpSecVpnTunnel) | **Post** /1.0.0/edgeGateways/{gatewayId}/ipsec/tunnels | Creates an IPSec tunnel on the Edge Gateway.
[**GetIpSecVpnTunnels**](EdgeGatewayIpSecVpnTunnelsApi.md#GetIpSecVpnTunnels) | **Get** /1.0.0/edgeGateways/{gatewayId}/ipsec/tunnels | Retrieves all IPSec tunnels for a given edge gateway.


# **CreateIpSecVpnTunnel**
> CreateIpSecVpnTunnel(ctx, ipsecVpnTunnel, gatewayId)
Creates an IPSec tunnel on the Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnTunnel** | [**EdgeIpSecVpnTunnel**](EdgeIpSecVpnTunnel.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSecVpnTunnels**
> EdgeIpSecVpnTunnels GetIpSecVpnTunnels(ctx, pageSize, gatewayId, optional)
Retrieves all IPSec tunnels for a given edge gateway.

Retrieves all IPSec VPN tunnels that are configured for an edge gateway. Pagination is supported, use response header to get the next page. Results can be sorted by only a single parameter. Sorting by combination of parameters (sortAsc=foo&sortDesc=bar) is not allowed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **gatewayId** | **string**|  | 
 **optional** | ***EdgeGatewayIpSecVpnTunnelsApiGetIpSecVpnTunnelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayIpSecVpnTunnelsApiGetIpSecVpnTunnelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EdgeIpSecVpnTunnels**](EdgeIpSecVpnTunnels.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

