# \OrgVdcNetworksApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetwork**](OrgVdcNetworksApi.md#CreateNetwork) | **Post** /1.0.0/orgVdcNetworks | Create an organization vDC network.
[**GetAllVdcNetworks**](OrgVdcNetworksApi.md#GetAllVdcNetworks) | **Get** /1.0.0/orgVdcNetworks | Get all Org vDC networks.


# **CreateNetwork**
> CreateNetwork(ctx, vdcNetwork)
Create an organization vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetwork** | [**VdcNetwork**](VdcNetwork.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllVdcNetworks**
> VdcNetworks GetAllVdcNetworks(ctx, page, pageSize, optional)
Get all Org vDC networks.

Get all Org vDC networks. If \"ownerRef\" property is not specified in the filter, then user must have the rights to view all the vDCs within an organization in order to see all the networks in the organization. Results can be filtered by ownerRef or combination of ownerRef and _context. <code>(_context==includeAccessible)</code> can be used to get all the networks which are available to an Org vDC. In order to return only those networks which are eligible as an uplink to a vApp network, add an additional filter, <code>(vAppUplinkEligible==true)</code>. This filter must be used in conjunction with either the \"orgVdc.id\" filter, or the \"ownerRef.id\" filter with a value corresponding to an Org vDC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 16]
 **optional** | ***OrgVdcNetworksApiGetAllVdcNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgVdcNetworksApiGetAllVdcNetworksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VdcNetworks**](VdcNetworks.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

