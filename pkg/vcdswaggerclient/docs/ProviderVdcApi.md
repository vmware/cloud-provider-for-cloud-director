# \ProviderVdcApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllProviderVDCs**](ProviderVdcApi.md#GetAllProviderVDCs) | **Get** /1.0.0/providerVdcs | Get all provider VDCs.
[**GetChildResourcePools**](ProviderVdcApi.md#GetChildResourcePools) | **Get** /1.0.0/providerVdcs/{pvdcUrn}/infra/resourcePools/browse/{moref} | Browse valid root resource pools hierarchy to back a Provider VDC.
[**GetRootResourcePools**](ProviderVdcApi.md#GetRootResourcePools) | **Get** /1.0.0/providerVdcs/{pvdcUrn}/infra/resourcePools/browse/ | Browse valid root resource pools hierarchy to back a Provider VDC.


# **GetAllProviderVDCs**
> ProviderVdcs GetAllProviderVDCs(ctx, page, pageSize, optional)
Get all provider VDCs.

Retrieve a list of all provider VDCs. Results can be filtered by context (_context). Supported contexts are: DVS (_context==dvs-NN;vimServer.id==urn:vcloud:vimserver:uuid) - | Returns all the provider VDCs which are related to the DVS. The VimServer is required for this filtering. External Network (_context==urn:vcloud:network:uuid) - | Returns all the provider VDCs which are related to the external network. Network Pool (_context==urn:vcloud:networkpool:uuid) - | Returns all the provider VDCs which are accessible to the network pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***ProviderVdcApiGetAllProviderVDCsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProviderVdcApiGetAllProviderVDCsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**ProviderVdcs**](ProviderVdcs.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChildResourcePools**
> ResourcePools GetChildResourcePools(ctx, pvdcUrn, moref, page, pageSize)
Browse valid root resource pools hierarchy to back a Provider VDC.

Get list of child resource pools of the specified parent. If a resource pool is ineligible, but is in the response, this means it has children, which are eligible. A resource pool will be ineligible, unless the cluster has an ESXi host on it. The list will be sorted by name, case insensitive. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcUrn** | **string**|  | 
  **moref** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]

### Return type

[**ResourcePools**](ResourcePools.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRootResourcePools**
> ResourcePools GetRootResourcePools(ctx, pvdcUrn, page, pageSize)
Browse valid root resource pools hierarchy to back a Provider VDC.

Get a list of all root resource pools. If a resource pool is ineligible, but is in the response, this means it has children, which are eligible. A resource pool will be ineligible, unless the cluster has an ESXi host on it. The list will be sorted by name, case insensitive. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcUrn** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]

### Return type

[**ResourcePools**](ResourcePools.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

