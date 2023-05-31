# \VdcComputePoliciesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVdcComputePolicyToVdcs**](VdcComputePoliciesApi.md#AddVdcComputePolicyToVdcs) | **Post** /1.0.0/vdcComputePolicies/{vdcComputePolicyId}/vdcs | Assign this organization vDC compute policy to list of vDCs.
[**CreateVdcComputePolicy**](VdcComputePoliciesApi.md#CreateVdcComputePolicy) | **Post** /1.0.0/vdcComputePolicies | Creates a new organization vDC compute policy
[**DeleteVdcComputePolicy**](VdcComputePoliciesApi.md#DeleteVdcComputePolicy) | **Delete** /1.0.0/vdcComputePolicies/{vdcComputePolicyId} | Delete specified organization vDC compute policy.
[**GetVdcComputePolicy**](VdcComputePoliciesApi.md#GetVdcComputePolicy) | **Get** /1.0.0/vdcComputePolicies/{vdcComputePolicyId} | Get specified organization vDC compute policy
[**GetVdcComputePolicyVdcs**](VdcComputePoliciesApi.md#GetVdcComputePolicyVdcs) | **Get** /1.0.0/vdcComputePolicies/{vdcComputePolicyId}/vdcs | Get organization vDCs associated with this vDC compute policy
[**GetVdcComputePolicyVms**](VdcComputePoliciesApi.md#GetVdcComputePolicyVms) | **Get** /1.0.0/vdcComputePolicies/{vdcComputePolicyId}/vms | Get all VMs associated with this vDC compute policy
[**QueryVdcComputePolicies**](VdcComputePoliciesApi.md#QueryVdcComputePolicies) | **Get** /1.0.0/vdcComputePolicies | Get list of organization vDC compute policies.
[**UpdateVdcComputePolicy**](VdcComputePoliciesApi.md#UpdateVdcComputePolicy) | **Put** /1.0.0/vdcComputePolicies/{vdcComputePolicyId} | Update specified organization vDC compute policy


# **AddVdcComputePolicyToVdcs**
> []EntityReference AddVdcComputePolicyToVdcs(ctx, vdcRefs, vdcComputePolicyId)
Assign this organization vDC compute policy to list of vDCs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcRefs** | [**[]EntityReference**](EntityReference.md)|  | 
  **vdcComputePolicyId** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVdcComputePolicy**
> VdcComputePolicy CreateVdcComputePolicy(ctx, newVdcComputePolicyParams)
Creates a new organization vDC compute policy

Creates a new vDC compute policy 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newVdcComputePolicyParams** | [**VdcComputePolicy**](VdcComputePolicy.md)|  | 

### Return type

[**VdcComputePolicy**](VdcComputePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVdcComputePolicy**
> DeleteVdcComputePolicy(ctx, vdcComputePolicyId)
Delete specified organization vDC compute policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcComputePolicyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcComputePolicy**
> VdcComputePolicy GetVdcComputePolicy(ctx, vdcComputePolicyId)
Get specified organization vDC compute policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcComputePolicyId** | **string**|  | 

### Return type

[**VdcComputePolicy**](VdcComputePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcComputePolicyVdcs**
> []EntityReference GetVdcComputePolicyVdcs(ctx, vdcComputePolicyId)
Get organization vDCs associated with this vDC compute policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcComputePolicyId** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcComputePolicyVms**
> EntityReferences GetVdcComputePolicyVms(ctx, vdcComputePolicyId, page, pageSize, optional)
Get all VMs associated with this vDC compute policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcComputePolicyId** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VdcComputePoliciesApiGetVdcComputePolicyVmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VdcComputePoliciesApiGetVdcComputePolicyVmsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryVdcComputePolicies**
> VdcComputePolicies QueryVdcComputePolicies(ctx, page, pageSize, optional)
Get list of organization vDC compute policies.

Get list of organization vDC compute policy. Results can be filtered by id, pvdcId, isSizingOnly and _context. Supported contexts are: Org vDC Urn ID (_context==orgVdcUrn) - | Returns all the vDC compute policies which are available to a specific Org vDC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VdcComputePoliciesApiQueryVdcComputePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VdcComputePoliciesApiQueryVdcComputePoliciesOpts struct

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVdcComputePolicy**
> VdcComputePolicy UpdateVdcComputePolicy(ctx, updateVdcComputePolicyParams, vdcComputePolicyId)
Update specified organization vDC compute policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateVdcComputePolicyParams** | [**VdcComputePolicy**](VdcComputePolicy.md)|  | 
  **vdcComputePolicyId** | **string**|  | 

### Return type

[**VdcComputePolicy**](VdcComputePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

