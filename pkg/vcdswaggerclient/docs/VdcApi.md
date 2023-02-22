# \VdcApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComputePolicies**](VdcApi.md#GetComputePolicies) | **Get** /1.0.0/vdcs/{orgVdcId}/computePolicies | Retrieves all compute policies of a vDC
[**GetVdc**](VdcApi.md#GetVdc) | **Get** /1.0.0/vdcs/{orgVdcId} | Retrieves a specific vDC via URN.
[**GetVdcMaxComputePolicy**](VdcApi.md#GetVdcMaxComputePolicy) | **Get** /1.0.0/vdcs/{orgVdcId}/maxComputePolicy | Retrieves Max Compute Policy of the vDC.
[**QueryVdcs**](VdcApi.md#QueryVdcs) | **Get** /1.0.0/vdcs | Retrieves a list of vDCs
[**UpdateVdcMaxComputePolicy**](VdcApi.md#UpdateVdcMaxComputePolicy) | **Put** /1.0.0/vdcs/{orgVdcId}/maxComputePolicy | Updates Max Compute Policy of the vDC.


# **GetComputePolicies**
> VdcComputePolicies GetComputePolicies(ctx, page, pageSize, orgVdcId, optional)
Retrieves all compute policies of a vDC

Retrieves all compute policies of a vDC 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **orgVdcId** | **string**|  | 
 **optional** | ***VdcApiGetComputePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VdcApiGetComputePoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VdcComputePolicies**](VdcComputePolicies.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdc**
> OrgVdc GetVdc(ctx, orgVdcId)
Retrieves a specific vDC via URN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgVdcId** | **string**|  | 

### Return type

[**OrgVdc**](OrgVdc.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcMaxComputePolicy**
> VdcComputePolicy GetVdcMaxComputePolicy(ctx, orgVdcId)
Retrieves Max Compute Policy of the vDC.

Retrieves Max Compute Policy of the vDC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgVdcId** | **string**|  | 

### Return type

[**VdcComputePolicy**](VdcComputePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryVdcs**
> OrgVdcs QueryVdcs(ctx, page, pageSize, optional)
Retrieves a list of vDCs

Retrieves a list of Org-scoped (if applicable) vDCs. Results can be filtered by id, name, allocationType, and computePolicyType.  Supported filters for computePolicyType are:   (computePolicyType==VdcKubernetesPolicy) - |     Returns a list of all vDCs that have a VdcKubernetesPolicy compute policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VdcApiQueryVdcsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VdcApiQueryVdcsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**OrgVdcs**](OrgVdcs.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVdcMaxComputePolicy**
> VdcComputePolicy UpdateVdcMaxComputePolicy(ctx, newVdcComputePolicyParams, orgVdcId)
Updates Max Compute Policy of the vDC.

Updates Max Compute Policy of the vDC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newVdcComputePolicyParams** | [**VdcComputePolicy**](VdcComputePolicy.md)|  | 
  **orgVdcId** | **string**|  | 

### Return type

[**VdcComputePolicy**](VdcComputePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

