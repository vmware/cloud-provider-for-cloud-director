# \EdgeGatewayApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEdgeGateway**](EdgeGatewayApi.md#DeleteEdgeGateway) | **Delete** /1.0.0/edgeGateways/{gatewayId} | Deletes a specific Edge Gateway
[**GetEdgeGateway**](EdgeGatewayApi.md#GetEdgeGateway) | **Get** /1.0.0/edgeGateways/{gatewayId} | Retrieves a specific Edge Gateway
[**GetUsedIpAddresses**](EdgeGatewayApi.md#GetUsedIpAddresses) | **Get** /1.0.0/edgeGateways/{gatewayId}/usedIpAddresses | Retrieve the list of IP addresses which are being used by the edge gateway.
[**UpdateEdgeGateway**](EdgeGatewayApi.md#UpdateEdgeGateway) | **Put** /1.0.0/edgeGateways/{gatewayId} | Updates a specific Edge Gateway


# **DeleteEdgeGateway**
> DeleteEdgeGateway(ctx, gatewayId, optional)
Deletes a specific Edge Gateway

Deletes a specific Edge Gateway. Only NSX-T Edge Gateways can be deleted with this endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
 **optional** | ***EdgeGatewayApiDeleteEdgeGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayApiDeleteEdgeGatewayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Value \&quot;true\&quot; means to forcefully delete the object that contains other objects even if those objects are in a state that does not allow removal. The default is \&quot;false\&quot;; therefore, objects are not removed if they are not in a state that normally allows removal. Force also implies recursive delete where other contained objects are removed. Errors may be ignored. Invalid value (not true or false) are ignored.  | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeGateway**
> EdgeGateway GetEdgeGateway(ctx, gatewayId)
Retrieves a specific Edge Gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**EdgeGateway**](EdgeGateway.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsedIpAddresses**
> GatewayUsedIpAddresses GetUsedIpAddresses(ctx, page, pageSize, gatewayId, optional)
Retrieve the list of IP addresses which are being used by the edge gateway.

Get all the IP Addresses which are being used by the Edge Gateway such as the primary IP or an IP used by a given Edge Service, such as NAT. These IP addresses are a subset of the IPs allocated from the connected external networks. If the IP is  being consumed by any of the configured services on the edge gateway then name of service will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **gatewayId** | **string**|  | 
 **optional** | ***EdgeGatewayApiGetUsedIpAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayApiGetUsedIpAddressesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**GatewayUsedIpAddresses**](GatewayUsedIpAddresses.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEdgeGateway**
> UpdateEdgeGateway(ctx, gateway, gatewayId)
Updates a specific Edge Gateway

Update a specific Edge Gateway. Only NSX-T Edge Gateways can be created with this endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gateway** | [**EdgeGateway**](EdgeGateway.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

