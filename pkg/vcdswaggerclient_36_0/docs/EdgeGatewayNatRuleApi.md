# \EdgeGatewayNatRuleApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNatRule**](EdgeGatewayNatRuleApi.md#DeleteNatRule) | **Delete** /1.0.0/edgeGateways/{gatewayId}/nat/rules/{ruleId} | Deletes a specific NAT Rule configuration of the edge gateway based on the rule id passed in.
[**GetNatRule**](EdgeGatewayNatRuleApi.md#GetNatRule) | **Get** /1.0.0/edgeGateways/{gatewayId}/nat/rules/{ruleId} | Retrieves a specific NAT Rule configuration of the edge gateway based on the rule id passed in.
[**UpdateNatRule**](EdgeGatewayNatRuleApi.md#UpdateNatRule) | **Put** /1.0.0/edgeGateways/{gatewayId}/nat/rules/{ruleId} | Update a specific NAT Rule configuration of the edge gateway based on the rule id passed in.


# **DeleteNatRule**
> DeleteNatRule(ctx, gatewayId, ruleId)
Deletes a specific NAT Rule configuration of the edge gateway based on the rule id passed in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNatRule**
> EdgeNatRule GetNatRule(ctx, gatewayId, ruleId)
Retrieves a specific NAT Rule configuration of the edge gateway based on the rule id passed in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**EdgeNatRule**](EdgeNatRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNatRule**
> UpdateNatRule(ctx, edgeNatRule, gatewayId, ruleId)
Update a specific NAT Rule configuration of the edge gateway based on the rule id passed in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeNatRule** | [**EdgeNatRule**](EdgeNatRule.md)|  | 
  **gatewayId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

