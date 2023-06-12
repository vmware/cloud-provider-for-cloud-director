# \DfwPoliciesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDfwPolicies**](DfwPoliciesApi.md#GetDfwPolicies) | **Get** /1.0.0/vdcGroups/{vdcGroupId}/dfwPolicies | Retrieves DFW security policies configuration.
[**UpdateDfwPolicies**](DfwPoliciesApi.md#UpdateDfwPolicies) | **Put** /1.0.0/vdcGroups/{vdcGroupId}/dfwPolicies | Update DFW security policies configuration.


# **GetDfwPolicies**
> DfwPolicies GetDfwPolicies(ctx, vdcGroupId)
Retrieves DFW security policies configuration.

Retrieves the current state of DFW along with all the DFW security policies for a given networking and security domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 

### Return type

[**DfwPolicies**](DfwPolicies.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDfwPolicies**
> UpdateDfwPolicies(ctx, dfwPolicies, vdcGroupId)
Update DFW security policies configuration.

Updates the DFW security policies for a given networking and security domain. A default security policy will be created when DFW is enabled. Removing a security policy will result in removal of the policy and all of its associated firewall rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dfwPolicies** | [**DfwPolicies**](DfwPolicies.md)|  | 
  **vdcGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

