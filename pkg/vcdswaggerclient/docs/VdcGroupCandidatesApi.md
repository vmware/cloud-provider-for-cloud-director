# \VdcGroupCandidatesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkingVdcGroupCandidates**](VdcGroupCandidatesApi.md#GetNetworkingVdcGroupCandidates) | **Get** /1.0.0/vdcGroups/networkingCandidateVdcs | Get a list of candidate vDCs that can be added to a single vDC Group in the networking context.


# **GetNetworkingVdcGroupCandidates**
> NetworkingCandidateVdcs GetNetworkingVdcGroupCandidates(ctx, page, pageSize, optional)
Get a list of candidate vDCs that can be added to a single vDC Group in the networking context.

Get a list of candidate vDCs. Results can be filtered by the type of vDC group and contexts of vDC ID or network pool ID. Below are the supported contexts: <ul> <li>Local vDC Group<code>(_context==LOCAL;_context==orgVdcId)</code> - Returns all vDCs that have the same network pool assigned to the org vDC. <li>Universal vDC Group backed by NSX-V network provider <code>(_context==UNIVERSAL;_context==orgVdcId;_context==networkPoolUniversalId)</code> - If the orgVdc is specified, returns all vDCs that have the same assigned universal network pool as that of the org vDC. A universal network pool is backed by a broadcast domain that can stretch to multiple VMware Cloud Director sites. If the org vDC is not specified or not found and the universal network pool is specified, return all vDCs whose network pool are backed by that universal network pool. This case supports multisite/associated organization queries where user wants to retrieve remote vDCs on different sites/associated organizations with same backing universal network pool. <li>Existing vDC Group<code>(_context==vdcGroupId;_context==networkPoolUniversalId)</code> - If vDC Group is specified, return all vDCs that have the?same assigned network pool as that of the vDC Group and also all vDCs that are currently participating in the vDC Group. The assigned network pool can also be universal if the vDC Group type is UNIVERSAL. If vDC Group is not specified or not found and the universal network pool is specified, return all vDCs whose network pool are backed by that universal network pool. This case supports multisite/associated organization queries where user wants to retrieve remote vDCs on different sites/associated organizations with same backing universal network pool. </ul> Note that multisite calls to get the candidate vDCs for a local vDC Group will only return the vDCs of the local site's associated organizations. Remote site's vDCs are not returned since the org vDC or the vDC group is not found there. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VdcGroupCandidatesApiGetNetworkingVdcGroupCandidatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VdcGroupCandidatesApiGetNetworkingVdcGroupCandidatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**NetworkingCandidateVdcs**](NetworkingCandidateVdcs.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

