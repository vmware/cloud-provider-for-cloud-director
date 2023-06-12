# \VdcGroupsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVdcGroup**](VdcGroupsApi.md#CreateVdcGroup) | **Post** /1.0.0/vdcGroups | Creates a vDC Group. A universal router will also be created if universalNetworkingEnabled is set to true.
[**GetVdcGroups**](VdcGroupsApi.md#GetVdcGroups) | **Get** /1.0.0/vdcGroups | Get a list of vDC Groups.


# **CreateVdcGroup**
> CreateVdcGroup(ctx, vdcGroup)
Creates a vDC Group. A universal router will also be created if universalNetworkingEnabled is set to true.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroup** | [**VdcGroup**](VdcGroup.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcGroups**
> VdcGroups GetVdcGroups(ctx, page, pageSize, optional)
Get a list of vDC Groups.

Get a list of vDC Groups. To find all vDC Groups that contains a specific Organization vDC, user can use the filter \"participatingOrgVdcs.vdcRef.id\" key. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VdcGroupsApiGetVdcGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VdcGroupsApiGetVdcGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VdcGroups**](VdcGroups.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

