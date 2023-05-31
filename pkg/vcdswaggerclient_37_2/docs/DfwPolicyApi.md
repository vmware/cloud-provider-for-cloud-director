# \DfwPolicyApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDfwPolicy**](DfwPolicyApi.md#DeleteDfwPolicy) | **Delete** /1.0.0/vdcGroups/{vdcGroupId}/dfwPolicies/{policyId} | Deletes a specific DFW security policy. Removing a security policy will result in removal of the policy and all of its associated firewall rules. 
[**DeleteDfwRule**](DfwPolicyApi.md#DeleteDfwRule) | **Delete** /1.0.0/vdcGroups/{vdcGroupId}/dfwPolicies/{policyId}/rules/{ruleId} | Deletes a specific firewall rule for a given DFW security policy.
[**GetDfwPolicy**](DfwPolicyApi.md#GetDfwPolicy) | **Get** /1.0.0/vdcGroups/{vdcGroupId}/dfwPolicies/{policyId} | Retrieves a specific DFW security policy.
[**GetDfwRule**](DfwPolicyApi.md#GetDfwRule) | **Get** /1.0.0/vdcGroups/{vdcGroupId}/dfwPolicies/{policyId}/rules/{ruleId} | Retrieves a specific firewall rule for a given DFW security policy.
[**GetDfwRules**](DfwPolicyApi.md#GetDfwRules) | **Get** /1.0.0/vdcGroups/{vdcGroupId}/dfwPolicies/{policyId}/rules | Retrieves all firewall rules for a given DFW security policy.
[**UpdateDfwPolicy**](DfwPolicyApi.md#UpdateDfwPolicy) | **Put** /1.0.0/vdcGroups/{vdcGroupId}/dfwPolicies/{policyId} | Updates a specific DFW security policy.
[**UpdateDfwRule**](DfwPolicyApi.md#UpdateDfwRule) | **Put** /1.0.0/vdcGroups/{vdcGroupId}/dfwPolicies/{policyId}/rules/{ruleId} | Updates a specific firewall rule for a given DFW security policy.
[**UpdateDfwRules**](DfwPolicyApi.md#UpdateDfwRules) | **Put** /1.0.0/vdcGroups/{vdcGroupId}/dfwPolicies/{policyId}/rules | Updates firewall rules for a given DFW security policy.


# **DeleteDfwPolicy**
> DeleteDfwPolicy(ctx, vdcGroupId, policyId)
Deletes a specific DFW security policy. Removing a security policy will result in removal of the policy and all of its associated firewall rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDfwRule**
> DeleteDfwRule(ctx, vdcGroupId, policyId, ruleId)
Deletes a specific firewall rule for a given DFW security policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 
  **policyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDfwPolicy**
> DfwPolicy GetDfwPolicy(ctx, vdcGroupId, policyId)
Retrieves a specific DFW security policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

[**DfwPolicy**](DfwPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDfwRule**
> DfwRule GetDfwRule(ctx, vdcGroupId, policyId, ruleId)
Retrieves a specific firewall rule for a given DFW security policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 
  **policyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**DfwRule**](DfwRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDfwRules**
> DfwRules GetDfwRules(ctx, vdcGroupId, policyId)
Retrieves all firewall rules for a given DFW security policy.

Retrieves all firewall rules for a given DFW security policy. The rules are returned in the order of precedence. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

[**DfwRules**](DfwRules.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDfwPolicy**
> UpdateDfwPolicy(ctx, dfwPolicy, vdcGroupId, policyId)
Updates a specific DFW security policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dfwPolicy** | [**DfwPolicy**](DfwPolicy.md)|  | 
  **vdcGroupId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDfwRule**
> UpdateDfwRule(ctx, dfwRule, vdcGroupId, policyId, ruleId)
Updates a specific firewall rule for a given DFW security policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dfwRule** | [**DfwRule**](DfwRule.md)|  | 
  **vdcGroupId** | **string**|  | 
  **policyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDfwRules**
> UpdateDfwRules(ctx, dfwRules, vdcGroupId, policyId)
Updates firewall rules for a given DFW security policy.

Updates all the firewall rules for a given DFW security policy. If a rule with the ruleId is not already present, a new rule will be created. If it already exists, the rule will be updated. Any existing rule that is not specified in the update payload will be deleted. The order of rules in payload will define the actual order in which this rules will be applied. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dfwRules** | [**DfwRules**](DfwRules.md)|  | 
  **vdcGroupId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

