# \EdgeGatewayFirewallRuleApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFirewallRule**](EdgeGatewayFirewallRuleApi.md#DeleteFirewallRule) | **Delete** /1.0.0/edgeGateways/{gatewayId}/firewall/rules/{ruleId} | Deletes a specific firewall rule for a given edge gateway.
[**GetFirewallRule**](EdgeGatewayFirewallRuleApi.md#GetFirewallRule) | **Get** /1.0.0/edgeGateways/{gatewayId}/firewall/rules/{ruleId} | Retrieves a specific firewall rule for a given edge gateway.
[**UpdateFirewallRule**](EdgeGatewayFirewallRuleApi.md#UpdateFirewallRule) | **Put** /1.0.0/edgeGateways/{gatewayId}/firewall/rules/{ruleId} | Updates a specific firewall rule for a given edge gateway.


# **DeleteFirewallRule**
> DeleteFirewallRule(ctx, gatewayId, ruleId)
Deletes a specific firewall rule for a given edge gateway.

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
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallRule**
> EdgeFirewallRule GetFirewallRule(ctx, gatewayId, ruleId)
Retrieves a specific firewall rule for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**EdgeFirewallRule**](EdgeFirewallRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFirewallRule**
> UpdateFirewallRule(ctx, firewallRule, gatewayId, ruleId)
Updates a specific firewall rule for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallRule** | [**EdgeFirewallRule**](EdgeFirewallRule.md)|  | 
  **gatewayId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

