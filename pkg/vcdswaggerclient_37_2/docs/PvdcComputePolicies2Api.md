# \PvdcComputePolicies2Api

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePvdcComputePolicy**](PvdcComputePolicies2Api.md#CreatePvdcComputePolicy) | **Post** /2.0.0/pvdcComputePolicies | Creates a new provider vDC compute policy
[**DeletePvdcComputePolicy**](PvdcComputePolicies2Api.md#DeletePvdcComputePolicy) | **Delete** /2.0.0/pvdcComputePolicies/{pvdcComputePolicyId} | Delete specified provider vDC compute policy.
[**GetPvdcComputePolicy**](PvdcComputePolicies2Api.md#GetPvdcComputePolicy) | **Get** /2.0.0/pvdcComputePolicies/{pvdcComputePolicyId} | Get specified provider vDC compute policy
[**QueryPvdcComputePolicies**](PvdcComputePolicies2Api.md#QueryPvdcComputePolicies) | **Get** /2.0.0/pvdcComputePolicies | Get list of provider vDC compute policies.
[**QueryVirtualMachineClasses**](PvdcComputePolicies2Api.md#QueryVirtualMachineClasses) | **Get** /2.0.0/pvdcComputePolicies/{pvdcComputePolicyId}/virtualMachineClasses | Get a list of Virtual Machine Classes associated with this policy.
[**UpdatePvdcComputePolicy**](PvdcComputePolicies2Api.md#UpdatePvdcComputePolicy) | **Put** /2.0.0/pvdcComputePolicies/{pvdcComputePolicyId} | Update specified provider vDC compute policy


# **CreatePvdcComputePolicy**
> PvdcComputePolicy2 CreatePvdcComputePolicy(ctx, pvdcComputePolicy)
Creates a new provider vDC compute policy

Creates a new pVDC compute policy 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcComputePolicy** | [**PvdcComputePolicy2**](PvdcComputePolicy2.md)|  | 

### Return type

[**PvdcComputePolicy2**](PvdcComputePolicy2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePvdcComputePolicy**
> DeletePvdcComputePolicy(ctx, pvdcComputePolicyId)
Delete specified provider vDC compute policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcComputePolicyId** | **string**| ID of provider vDC Compute Policy | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPvdcComputePolicy**
> PvdcComputePolicy2 GetPvdcComputePolicy(ctx, pvdcComputePolicyId)
Get specified provider vDC compute policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcComputePolicyId** | **string**| ID of provider vDC Compute Policy | 

### Return type

[**PvdcComputePolicy2**](PvdcComputePolicy2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryPvdcComputePolicies**
> PvdcComputePolicies2 QueryPvdcComputePolicies(ctx, page, pageSize, optional)
Get list of provider vDC compute policies.

Get list of provider vDC compute policies. Only filtering by provider vDC compute policy name is supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***PvdcComputePolicies2ApiQueryPvdcComputePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PvdcComputePolicies2ApiQueryPvdcComputePoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**PvdcComputePolicies2**](PvdcComputePolicies2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryVirtualMachineClasses**
> VirtualMachineClasses QueryVirtualMachineClasses(ctx, page, pageSize, pvdcComputePolicyId, optional)
Get a list of Virtual Machine Classes associated with this policy.

Get a list of Virtual Machine Classes associated with this policy. Returns 400 if policy type is anything but PvdcKubernetesPolicy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **pvdcComputePolicyId** | **string**|  | 
 **optional** | ***PvdcComputePolicies2ApiQueryVirtualMachineClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PvdcComputePolicies2ApiQueryVirtualMachineClassesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VirtualMachineClasses**](VirtualMachineClasses.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePvdcComputePolicy**
> PvdcComputePolicy2 UpdatePvdcComputePolicy(ctx, pvdcComputePolicyId, pvdcComputePolicy)
Update specified provider vDC compute policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcComputePolicyId** | **string**| ID of provider vDC Compute Policy | 
  **pvdcComputePolicy** | [**PvdcComputePolicy2**](PvdcComputePolicy2.md)|  | 

### Return type

[**PvdcComputePolicy2**](PvdcComputePolicy2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

