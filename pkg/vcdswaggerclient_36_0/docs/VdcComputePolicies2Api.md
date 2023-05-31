# \VdcComputePolicies2Api

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVdcComputePolicyToVdcs**](VdcComputePolicies2Api.md#AddVdcComputePolicyToVdcs) | **Post** /2.0.0/vdcComputePolicies/{vdcComputePolicyId}/vdcs | Assign this organization VDC compute policy to list of VDCs.
[**CreateVdcComputePolicy**](VdcComputePolicies2Api.md#CreateVdcComputePolicy) | **Post** /2.0.0/vdcComputePolicies | Creates a new organization VDC compute policy
[**DeleteVdcComputePolicy**](VdcComputePolicies2Api.md#DeleteVdcComputePolicy) | **Delete** /2.0.0/vdcComputePolicies/{vdcComputePolicyId} | Delete specified organization VDC compute policy.
[**GetVdcComputePolicy**](VdcComputePolicies2Api.md#GetVdcComputePolicy) | **Get** /2.0.0/vdcComputePolicies/{vdcComputePolicyId} | Get specified organization VDC compute policy
[**GetVdcComputePolicyMetrics**](VdcComputePolicies2Api.md#GetVdcComputePolicyMetrics) | **Get** /2.0.0/vdcComputePolicies/{vdcComputePolicyId}/metrics | Get specified organization VDC compute policy metrics.
[**GetVdcComputePolicyVdcs**](VdcComputePolicies2Api.md#GetVdcComputePolicyVdcs) | **Get** /2.0.0/vdcComputePolicies/{vdcComputePolicyId}/vdcs | Get orgatization VDCs this VDC compute policy has been assigned/published to
[**QueryVdcComputePolicies**](VdcComputePolicies2Api.md#QueryVdcComputePolicies) | **Get** /2.0.0/vdcComputePolicies | Get list of organization VDC compute policies.
[**QueryVirtualMachineClasses**](VdcComputePolicies2Api.md#QueryVirtualMachineClasses) | **Get** /2.0.0/vdcComputePolicies/{vdcComputePolicyId}/virtualMachineClasses | Get a list of Virtual Machine Classes associated with this policy.
[**UpdateVdcComputePolicy**](VdcComputePolicies2Api.md#UpdateVdcComputePolicy) | **Put** /2.0.0/vdcComputePolicies/{vdcComputePolicyId} | Update specified organization VDC compute policy


# **AddVdcComputePolicyToVdcs**
> []EntityReference AddVdcComputePolicyToVdcs(ctx, vdcRefs, vdcComputePolicyId)
Assign this organization VDC compute policy to list of VDCs.

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
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVdcComputePolicy**
> VdcComputePolicy2 CreateVdcComputePolicy(ctx, newVdcComputePolicyParams)
Creates a new organization VDC compute policy

Creates a new VDC compute policy. if PolicyType is VdcKubernetesPolicy, then the response is a 202 with task URL in location header. For VdcVmPolicy type, the response is 201 with created policy in response body. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newVdcComputePolicyParams** | [**VdcComputePolicy2**](VdcComputePolicy2.md)|  | 

### Return type

[**VdcComputePolicy2**](VdcComputePolicy2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVdcComputePolicy**
> DeleteVdcComputePolicy(ctx, vdcComputePolicyId)
Delete specified organization VDC compute policy.

Deletes vDC compute policy. if PolicyType is VdcKubernetesPolicy, then the response is a 202 with task URL in location header, else 204 is returned. 

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
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcComputePolicy**
> VdcComputePolicy2 GetVdcComputePolicy(ctx, vdcComputePolicyId)
Get specified organization VDC compute policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcComputePolicyId** | **string**|  | 

### Return type

[**VdcComputePolicy2**](VdcComputePolicy2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcComputePolicyMetrics**
> VdcComputePolicyMetrics GetVdcComputePolicyMetrics(ctx, vdcComputePolicyId)
Get specified organization VDC compute policy metrics.

Returns metrics for the given compute policy. Returns 400 if the policyType does not support metrics. For example, if type of the policy is VdcVmPolicy, then this API returns a 400 BadRequest response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcComputePolicyId** | **string**|  | 

### Return type

[**VdcComputePolicyMetrics**](VdcComputePolicyMetrics.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcComputePolicyVdcs**
> []EntityReference GetVdcComputePolicyVdcs(ctx, vdcComputePolicyId)
Get orgatization VDCs this VDC compute policy has been assigned/published to

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
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryVdcComputePolicies**
> VdcComputePolicies2 QueryVdcComputePolicies(ctx, page, pageSize, optional)
Get list of organization VDC compute policies.

Get list of organization VDC compute policy. <br> Results can be filtered by: <ul>   <li> id   <li> name   <li> pvdcId   <li> isSizingOnly   <li> policyType   <li> vdc.id   <li> pvdcComputePolicy.id   <li> publishableToVdc   <li> pvdc   <li> isAutoGenerated   <li> _context  <ul> <br> Supported contexts are: Org VDC Urn ID (_context==orgVdcUrn) - | <br> Returns all the VDC compute policies which are available to a specific Org VDC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VdcComputePolicies2ApiQueryVdcComputePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VdcComputePolicies2ApiQueryVdcComputePoliciesOpts struct

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

# **QueryVirtualMachineClasses**
> VirtualMachineClasses QueryVirtualMachineClasses(ctx, page, pageSize, vdcComputePolicyId, optional)
Get a list of Virtual Machine Classes associated with this policy.

Get a list of Virtual Machine Classes associated with this policy. Returns 400 if policy type is anything but VdcKubernetesPolicy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **vdcComputePolicyId** | **string**|  | 
 **optional** | ***VdcComputePolicies2ApiQueryVirtualMachineClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VdcComputePolicies2ApiQueryVirtualMachineClassesOpts struct

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
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVdcComputePolicy**
> VdcComputePolicy2 UpdateVdcComputePolicy(ctx, updateVdcComputePolicyParams, vdcComputePolicyId)
Update specified organization VDC compute policy

Updates vDC compute policy. if PolicyType is VdcKubernetesPolicy, then the response is a 202 with task URL in location header. For VdcVmPolicy type, the response is 200 with updated policy in response body. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateVdcComputePolicyParams** | [**VdcComputePolicy2**](VdcComputePolicy2.md)|  | 
  **vdcComputePolicyId** | **string**|  | 

### Return type

[**VdcComputePolicy2**](VdcComputePolicy2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

