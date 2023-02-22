# \ApplicationPortProfilesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationPortProfile**](ApplicationPortProfilesApi.md#CreateApplicationPortProfile) | **Post** /1.0.0/applicationPortProfiles | Create a user-defined application port profile.
[**GetApplicationPortProfiles**](ApplicationPortProfilesApi.md#GetApplicationPortProfiles) | **Get** /1.0.0/applicationPortProfiles | Get all Application Port Profiles.
[**SyncApplicationPortProfiles**](ApplicationPortProfilesApi.md#SyncApplicationPortProfiles) | **Post** /1.0.0/applicationPortProfiles/sync | Sync the application port profiles from the network provider to VCD.


# **CreateApplicationPortProfile**
> CreateApplicationPortProfile(ctx, applicationPortProfile)
Create a user-defined application port profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationPortProfile** | [**ApplicationPortProfile**](ApplicationPortProfile.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationPortProfiles**
> ApplicationPortProfiles GetApplicationPortProfiles(ctx, page, pageSize, optional)
Get all Application Port Profiles.

Retrieves all Application Port Profiles. This includes user-defined profiles and default system profile. Supported contexts are: Org vDC ID (_context==orgVdcId) - | Returns all the application port profiles which are available to a specific Org vDC. Network provider ID (_context==networkProviderId) - | Returns all the application port profiles which are available under a specific network provider. VDC Group Id <code>(_context==vdcGroupId)</code> - | Returns all the application port profiles which are available to a specific vDC Group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***ApplicationPortProfilesApiGetApplicationPortProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationPortProfilesApiGetApplicationPortProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**ApplicationPortProfiles**](ApplicationPortProfiles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncApplicationPortProfiles**
> SyncApplicationPortProfiles(ctx, optional)
Sync the application port profiles from the network provider to VCD.

Sync the application port profiles from the network provider to VCD. The network provider is required to be specified in the filter context. Context example: (_context==networkProviderId). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationPortProfilesApiSyncApplicationPortProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationPortProfilesApiSyncApplicationPortProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter for a query.  FIQL format. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

