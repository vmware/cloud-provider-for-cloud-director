# \EdgeGatewayRouteAdvertisementApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRouteAdvertisement**](EdgeGatewayRouteAdvertisementApi.md#GetRouteAdvertisement) | **Get** /1.0.0/edgeGateways/{gatewayId}/routing/advertisement | Retrieve the list of subnets that will be advertised so that the Edge Gateway can route out to the connected external network.
[**UpdateRouteAdvertisement**](EdgeGatewayRouteAdvertisementApi.md#UpdateRouteAdvertisement) | **Put** /1.0.0/edgeGateways/{gatewayId}/routing/advertisement | Updates the list of subnets that will be advertised so that the Edge Gateway can route out to the connected external network.


# **GetRouteAdvertisement**
> RouteAdvertisement GetRouteAdvertisement(ctx, gatewayId)
Retrieve the list of subnets that will be advertised so that the Edge Gateway can route out to the connected external network.

Retrieve the list of subnets that will be advertised so that the Edge Gateway can route out to the connected external network. Org vDC networks that are in any of these subnets can then be routed out to the external networks. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**RouteAdvertisement**](RouteAdvertisement.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouteAdvertisement**
> UpdateRouteAdvertisement(ctx, routeAdvertisement, gatewayId)
Updates the list of subnets that will be advertised so that the Edge Gateway can route out to the connected external network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeAdvertisement** | [**RouteAdvertisement**](RouteAdvertisement.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

