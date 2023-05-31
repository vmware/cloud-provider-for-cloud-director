# \VCenterResourcesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDvSwitches**](VCenterResourcesApi.md#GetDvSwitches) | **Get** /1.0.0/virtualCenters/resources/dvSwitches | Retrieves all distributed virtual switches.
[**GetImportableDvpgs**](VCenterResourcesApi.md#GetImportableDvpgs) | **Get** /1.0.0/virtualCenters/resources/importableDvpgs | Get all DVPG network backings that are available.
[**GetImportablePortgroups**](VCenterResourcesApi.md#GetImportablePortgroups) | **Get** /1.0.0/virtualCenters/resources/importablePortgroups | Get all standard porgroups available as backings.


# **GetDvSwitches**
> DvSwitches GetDvSwitches(ctx, page, pageSize, optional)
Retrieves all distributed virtual switches.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VCenterResourcesApiGetDvSwitchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VCenterResourcesApiGetDvSwitchesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**DvSwitches**](DvSwitches.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportableDvpgs**
> DistributedPortGroups GetImportableDvpgs(ctx, page, pageSize, optional)
Get all DVPG network backings that are available.

Get all DVPG network backings that are available. The \"_context\" filter key is optional and can be set with the id of the vCenter from which to obtain the DVPG network backings. \"orgVdcId==[vdcUrn]\" can be set as a filter to show importable DVPGs for an Org vDC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VCenterResourcesApiGetImportableDvpgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VCenterResourcesApiGetImportableDvpgsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**DistributedPortGroups**](DistributedPortGroups.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportablePortgroups**
> PortGroups GetImportablePortgroups(ctx, page, pageSize, optional)
Get all standard porgroups available as backings.

Get all standard portgroups that are available as backings. The \"_context\" filter key is optional and can be set with the id of the vCenter from which to obtain the standard porgroup backings. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VCenterResourcesApiGetImportablePortgroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VCenterResourcesApiGetImportablePortgroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**PortGroups**](PortGroups.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

