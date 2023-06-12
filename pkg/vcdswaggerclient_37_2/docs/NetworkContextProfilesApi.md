# \NetworkContextProfilesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkContextProfile**](NetworkContextProfilesApi.md#CreateNetworkContextProfile) | **Post** /1.0.0/networkContextProfiles | Create a user-defined network context profile.
[**GetNetworkContextProfiles**](NetworkContextProfilesApi.md#GetNetworkContextProfiles) | **Get** /1.0.0/networkContextProfiles | Get all network context profiles.
[**SyncNetworkContextProfiles**](NetworkContextProfilesApi.md#SyncNetworkContextProfiles) | **Post** /1.0.0/networkContextProfiles/sync | Sync the network context profiles from the network provider to VCD.


# **CreateNetworkContextProfile**
> CreateNetworkContextProfile(ctx, networkContextProfile)
Create a user-defined network context profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkContextProfile** | [**NetworkContextProfile**](NetworkContextProfile.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkContextProfiles**
> NetworkContextProfiles GetNetworkContextProfiles(ctx, page, pageSize, optional)
Get all network context profiles.

Retrieves all network context profiles defined in the system. Supported contexts are: <ul> <li> Org vDC ID <code>(_context==orgVdcId)</code> - Returns all the network context profiles which are available to a specific Org vDC. <li>Network provider ID (_context==networkProviderId) - | Returns all the network context profiles which are available for a specific network provider. <li>VDC Group Id <code>(_context==vdcGroupId)</code> - | Returns all the network context profiles which are available to a specific vDC Group. </ul> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NetworkContextProfilesApiGetNetworkContextProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkContextProfilesApiGetNetworkContextProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**NetworkContextProfiles**](NetworkContextProfiles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncNetworkContextProfiles**
> SyncNetworkContextProfiles(ctx, optional)
Sync the network context profiles from the network provider to VCD.

Sync the network context profiles from the network provider to VCD. The network provider is required to be specified in the filter context. Context example: (_context==networkProviderId). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkContextProfilesApiSyncNetworkContextProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkContextProfilesApiSyncNetworkContextProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter for a query.  FIQL format. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

