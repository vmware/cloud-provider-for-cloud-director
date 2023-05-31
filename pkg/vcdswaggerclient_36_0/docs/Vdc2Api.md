# \Vdc2Api

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComputePolicies**](Vdc2Api.md#GetComputePolicies) | **Get** /2.0.0/vdcs/{orgVdcId}/computePolicies | Retrieves all compute policies of a vDC
[**GetVdcMaxComputePolicy**](Vdc2Api.md#GetVdcMaxComputePolicy) | **Get** /2.0.0/vdcs/{orgVdcId}/maxComputePolicy | Retrieves Max Compute Policy of the vDC.
[**UpdateVdcMaxComputePolicy**](Vdc2Api.md#UpdateVdcMaxComputePolicy) | **Put** /2.0.0/vdcs/{orgVdcId}/maxComputePolicy | Updates Max Compute Policy of the vDC.


# **GetComputePolicies**
> VdcComputePolicies2 GetComputePolicies(ctx, page, pageSize, orgVdcId, optional)
Retrieves all compute policies of a vDC

Retrieves all compute policies of a vDC 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **orgVdcId** | **string**|  | 
 **optional** | ***Vdc2ApiGetComputePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Vdc2ApiGetComputePoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VdcComputePolicies2**](VdcComputePolicies2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcMaxComputePolicy**
> VdcComputePolicy2 GetVdcMaxComputePolicy(ctx, orgVdcId)
Retrieves Max Compute Policy of the vDC.

Retrieves Max Compute Policy of the vDC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgVdcId** | **string**|  | 

### Return type

[**VdcComputePolicy2**](VdcComputePolicy2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVdcMaxComputePolicy**
> VdcComputePolicy2 UpdateVdcMaxComputePolicy(ctx, newVdcComputePolicy2Params, orgVdcId)
Updates Max Compute Policy of the vDC.

Updates Max Compute Policy of the vDC. Returns 400 if policy type is not VdcVmPolicy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newVdcComputePolicy2Params** | [**VdcComputePolicy2**](VdcComputePolicy2.md)|  | 
  **orgVdcId** | **string**|  | 

### Return type

[**VdcComputePolicy2**](VdcComputePolicy2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

