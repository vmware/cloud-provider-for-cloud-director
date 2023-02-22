# \CrossVdcNetworksApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCrossVdcNetwork**](CrossVdcNetworksApi.md#CreateCrossVdcNetwork) | **Post** /1.0.0/crossVdcNetworks | Creates a Cross vDC network.
[**GetAllCrossVdcNetworks**](CrossVdcNetworksApi.md#GetAllCrossVdcNetworks) | **Get** /1.0.0/crossVdcNetworks | Get all Cross vDC networks in the system.
[**GetAllCrossVdcNetworksForVdcGroup**](CrossVdcNetworksApi.md#GetAllCrossVdcNetworksForVdcGroup) | **Get** /1.0.0/vdcGroups/{vdcGroupId}/crossVdcNetworks | Get all Cross vDC networks of a vDC group.


# **CreateCrossVdcNetwork**
> CreateCrossVdcNetwork(ctx, crossVdcNetwork)
Creates a Cross vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crossVdcNetwork** | [**CrossVdcNetwork**](CrossVdcNetwork.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllCrossVdcNetworks**
> CrossVdcNetworks GetAllCrossVdcNetworks(ctx, page, pageSize, optional)
Get all Cross vDC networks in the system.

Get all Cross vDC networks in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***CrossVdcNetworksApiGetAllCrossVdcNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrossVdcNetworksApiGetAllCrossVdcNetworksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**CrossVdcNetworks**](CrossVdcNetworks.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllCrossVdcNetworksForVdcGroup**
> CrossVdcNetworks GetAllCrossVdcNetworksForVdcGroup(ctx, page, pageSize, vdcGroupId, optional)
Get all Cross vDC networks of a vDC group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **vdcGroupId** | **string**|  | 
 **optional** | ***CrossVdcNetworksApiGetAllCrossVdcNetworksForVdcGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrossVdcNetworksApiGetAllCrossVdcNetworksForVdcGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**CrossVdcNetworks**](CrossVdcNetworks.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

