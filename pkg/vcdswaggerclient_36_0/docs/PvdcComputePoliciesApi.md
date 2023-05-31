# \PvdcComputePoliciesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePvdcComputePolicy**](PvdcComputePoliciesApi.md#CreatePvdcComputePolicy) | **Post** /1.0.0/pvdcComputePolicies | Creates a new provider vDC compute policy
[**DeletePvdcComputePolicy**](PvdcComputePoliciesApi.md#DeletePvdcComputePolicy) | **Delete** /1.0.0/pvdcComputePolicies/{pvdcComputePolicyId} | Delete specified provider vDC compute policy.
[**GetPvdcComputePolicy**](PvdcComputePoliciesApi.md#GetPvdcComputePolicy) | **Get** /1.0.0/pvdcComputePolicies/{pvdcComputePolicyId} | Get specified provider vDC compute policy
[**GetPvdcComputePolicyVms**](PvdcComputePoliciesApi.md#GetPvdcComputePolicyVms) | **Get** /1.0.0/pvdcComputePolicies/{pvdcComputePolicyId}/vms | Get all VMs associated with this pVDC compute policy
[**QueryPvdcComputePolicies**](PvdcComputePoliciesApi.md#QueryPvdcComputePolicies) | **Get** /1.0.0/pvdcComputePolicies | Get list of provider vDC compute policies.
[**UpdatePvdcComputePolicy**](PvdcComputePoliciesApi.md#UpdatePvdcComputePolicy) | **Put** /1.0.0/pvdcComputePolicies/{pvdcComputePolicyId} | Update specified provider vDC compute policy


# **CreatePvdcComputePolicy**
> PvdcComputePolicy CreatePvdcComputePolicy(ctx, pvdcComputePolicy)
Creates a new provider vDC compute policy

Creates a new pVDC compute policy 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcComputePolicy** | [**PvdcComputePolicy**](PvdcComputePolicy.md)|  | 

### Return type

[**PvdcComputePolicy**](PvdcComputePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePvdcComputePolicy**
> DeletePvdcComputePolicy(ctx, pvdcComputePolicyId)
Delete specified provider vDC compute policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcComputePolicyId** | **string**| ID of provider VDC Compute Policy | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPvdcComputePolicy**
> PvdcComputePolicy GetPvdcComputePolicy(ctx, pvdcComputePolicyId)
Get specified provider vDC compute policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcComputePolicyId** | **string**| ID of provider VDC Compute Policy | 

### Return type

[**PvdcComputePolicy**](PvdcComputePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPvdcComputePolicyVms**
> EntityReferences GetPvdcComputePolicyVms(ctx, pvdcComputePolicyId, page, pageSize, optional)
Get all VMs associated with this pVDC compute policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcComputePolicyId** | **string**| ID of provider VDC Compute Policy | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***PvdcComputePoliciesApiGetPvdcComputePolicyVmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PvdcComputePoliciesApiGetPvdcComputePolicyVmsOpts struct

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
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryPvdcComputePolicies**
> PvdcComputePolicies QueryPvdcComputePolicies(ctx, page, pageSize, optional)
Get list of provider vDC compute policies.

Get list of provider vDC compute policies. Only filtering by pvdc compute policy name is supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***PvdcComputePoliciesApiQueryPvdcComputePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PvdcComputePoliciesApiQueryPvdcComputePoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**PvdcComputePolicies**](PvdcComputePolicies.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePvdcComputePolicy**
> PvdcComputePolicy UpdatePvdcComputePolicy(ctx, pvdcComputePolicyId, pvdcComputePolicy)
Update specified provider vDC compute policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcComputePolicyId** | **string**| ID of provider VDC Compute Policy | 
  **pvdcComputePolicy** | [**PvdcComputePolicy**](PvdcComputePolicy.md)|  | 

### Return type

[**PvdcComputePolicy**](PvdcComputePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

