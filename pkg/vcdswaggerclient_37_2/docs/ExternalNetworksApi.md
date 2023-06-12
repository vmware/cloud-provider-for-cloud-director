# \ExternalNetworksApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllExternalNetworks**](ExternalNetworksApi.md#GetAllExternalNetworks) | **Get** /1.0.0/externalNetworks | Get all external networks.


# **GetAllExternalNetworks**
> ExternalNetworks GetAllExternalNetworks(ctx, page, pageSize, optional)
Get all external networks.

Get all external networks. Results can be filtered by id, name, backing type (networkBackings.values.backingType) and context (_context). Supported contexts are: <ul> <li>Provider vDC ID (_context==providerVdcId) - | Returns all the external networks which are available to a specific Provider vDC. <li>Org vDC ID (_context==orgVdcId) - | Returns all the external networks which are available to a specific Org vDC. <li>vCenter ID And Resource Pool Moref (_context==vCenterId;_context==rpMoref) - | Returns all the external networks accessible to a given vCenter resource pool. <li>Org vDC ID And Edge Deployment Mode (_context==orgVdcId;_context==anEdgeDeploymentMode) - | Returns all the external networks to which an edge gateway can connect. Edge Deployment Mode can be 'standaloneEdgeDeployment' or 'haEdgeDeployment'. Deployment mode specifies whether to use both primary edge cluster and secondary edge cluster or just primary edge cluster to determine external network accessibility. Edge clusters are determined via vDC Network Profile for input Org vDC. <li>Org vDC ID And Dedicatable External Networks (_context==orgVdcId;_context==dedicatable) - | Dedicatable only shows external networks that have no connected Edge Gateways. </ul> <em>orgAssociated==true</em> filter can be specified to get the list of external networks which are already associated with the Organization. An External network is already associated with the Organization, if either the Organization is the owner of the external network or if the Organization has an Edge Gateway which is connected to the external network. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***ExternalNetworksApiGetAllExternalNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalNetworksApiGetAllExternalNetworksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**ExternalNetworks**](ExternalNetworks.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

