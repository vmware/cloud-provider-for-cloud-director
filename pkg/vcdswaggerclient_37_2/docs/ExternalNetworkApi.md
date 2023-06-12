# \ExternalNetworkApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalNetwork**](ExternalNetworkApi.md#CreateExternalNetwork) | **Post** /1.0.0/externalNetworks | Create an external network
[**DeleteExternalNetwork**](ExternalNetworkApi.md#DeleteExternalNetwork) | **Delete** /1.0.0/externalNetworks/{externalNetworkId} | Deletes a specific external network.
[**GetAvailableIpAddresses**](ExternalNetworkApi.md#GetAvailableIpAddresses) | **Get** /1.0.0/externalNetworks/{externalNetworkId}/availableIpAddresses | Retrieve the list of IP addresses available for use on the network.
[**GetEdgeGatewaysForProviderGatewayTopology**](ExternalNetworkApi.md#GetEdgeGatewaysForProviderGatewayTopology) | **Get** /1.0.0/externalNetworks/{externalNetworkId}/topology/edgeGateways | Retrieve information about all the Edge Gateways which are associated with the Provider Gateway. This endpoint is only supported for external networks which are backed by NSX-T Tier-0 router. 
[**GetExternalNetwork**](ExternalNetworkApi.md#GetExternalNetwork) | **Get** /1.0.0/externalNetworks/{externalNetworkId} | Retrieves a specific external network.
[**GetIpSpaceUplinksForProviderGatewayTopology**](ExternalNetworkApi.md#GetIpSpaceUplinksForProviderGatewayTopology) | **Get** /1.0.0/externalNetworks/{externalNetworkId}/topology/ipSpaceUplinks | Retrieve information about all the IP Space uplinks which are associated with the Provider Gateway. This endpoint is only supported for external networks which are backed by NSX-T Tier-0 router. 
[**GetUsedIpAddresses**](ExternalNetworkApi.md#GetUsedIpAddresses) | **Get** /1.0.0/externalNetworks/{externalNetworkId}/usedIpAddresses | Retrieve the list of IP addresses which are being used from the network.
[**UpdateExternalNetwork**](ExternalNetworkApi.md#UpdateExternalNetwork) | **Put** /1.0.0/externalNetworks/{externalNetworkId} | Updates a specific external network.


# **CreateExternalNetwork**
> CreateExternalNetwork(ctx, externalNetwork)
Create an external network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalNetwork** | [**ExternalNetwork**](ExternalNetwork.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExternalNetwork**
> DeleteExternalNetwork(ctx, externalNetworkId, optional)
Deletes a specific external network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalNetworkId** | **string**|  | 
 **optional** | ***ExternalNetworkApiDeleteExternalNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalNetworkApiDeleteExternalNetworkOpts struct

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

# **GetAvailableIpAddresses**
> AvailableIpPoolSubnets GetAvailableIpAddresses(ctx, externalNetworkId)
Retrieve the list of IP addresses available for use on the network.

Get all the available IPs for a given external network. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalNetworkId** | **string**|  | 

### Return type

[**AvailableIpPoolSubnets**](AvailableIpPoolSubnets.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeGatewaysForProviderGatewayTopology**
> TopologyEdgeGateways GetEdgeGatewaysForProviderGatewayTopology(ctx, page, pageSize, externalNetworkId, optional)
Retrieve information about all the Edge Gateways which are associated with the Provider Gateway. This endpoint is only supported for external networks which are backed by NSX-T Tier-0 router. 

Get all the Edge Gateways which are associated with this Provider Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **externalNetworkId** | **string**|  | 
 **optional** | ***ExternalNetworkApiGetEdgeGatewaysForProviderGatewayTopologyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalNetworkApiGetEdgeGatewaysForProviderGatewayTopologyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**TopologyEdgeGateways**](TopologyEdgeGateways.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalNetwork**
> ExternalNetwork GetExternalNetwork(ctx, externalNetworkId)
Retrieves a specific external network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalNetworkId** | **string**|  | 

### Return type

[**ExternalNetwork**](ExternalNetwork.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSpaceUplinksForProviderGatewayTopology**
> TopologyIpSpaceUplinks GetIpSpaceUplinksForProviderGatewayTopology(ctx, page, pageSize, externalNetworkId, optional)
Retrieve information about all the IP Space uplinks which are associated with the Provider Gateway. This endpoint is only supported for external networks which are backed by NSX-T Tier-0 router. 

Get all the IP Space uplinks which are associated with this Provider Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **externalNetworkId** | **string**|  | 
 **optional** | ***ExternalNetworkApiGetIpSpaceUplinksForProviderGatewayTopologyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalNetworkApiGetIpSpaceUplinksForProviderGatewayTopologyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**TopologyIpSpaceUplinks**](TopologyIpSpaceUplinks.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsedIpAddresses**
> UsedIpAddresses GetUsedIpAddresses(ctx, page, pageSize, externalNetworkId, optional)
Retrieve the list of IP addresses which are being used from the network.

Get all the used IPs for a given external network. This returns all the IP addresses of network which are being used by a vApp VM or by an edge gateway connected to this external network. Results can be filtered by IP address. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **externalNetworkId** | **string**|  | 
 **optional** | ***ExternalNetworkApiGetUsedIpAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalNetworkApiGetUsedIpAddressesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**UsedIpAddresses**](UsedIpAddresses.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExternalNetwork**
> UpdateExternalNetwork(ctx, externalNetwork, externalNetworkId)
Updates a specific external network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalNetwork** | [**ExternalNetwork**](ExternalNetwork.md)|  | 
  **externalNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

