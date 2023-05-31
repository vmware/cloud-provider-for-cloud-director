# \EdgeGatewayFirewallRulesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFirewallRules**](EdgeGatewayFirewallRulesApi.md#DeleteFirewallRules) | **Delete** /1.0.0/edgeGateways/{gatewayId}/firewall/rules | Deletes all the firewall rules for a given edge gateway.
[**GetFirewallRules**](EdgeGatewayFirewallRulesApi.md#GetFirewallRules) | **Get** /1.0.0/edgeGateways/{gatewayId}/firewall/rules | Retrieves all firewall rules for a given edge gateway.
[**UpdateFirewallRules**](EdgeGatewayFirewallRulesApi.md#UpdateFirewallRules) | **Put** /1.0.0/edgeGateways/{gatewayId}/firewall/rules | Updates firewall rules for a given edge gateway.


# **DeleteFirewallRules**
> DeleteFirewallRules(ctx, gatewayId)
Deletes all the firewall rules for a given edge gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallRules**
> EdgeFirewallRules GetFirewallRules(ctx, gatewayId)
Retrieves all firewall rules for a given edge gateway.

Retrieves all user-defined and default firewall rules for a given edge gateway. The rules are returned in the order of precedence. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**EdgeFirewallRules**](EdgeFirewallRules.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFirewallRules**
> UpdateFirewallRules(ctx, firewallRules, gatewayId)
Updates firewall rules for a given edge gateway.

Updates all the firewall rules for a given edge gateway. If a rule with the ruleId is not already present, a new rule will be created. If it already exists, the rule will be updated. Any existing rule that is not specified in the update payload will be deleted. The order of rules in payload will define the actual order in which this rules will be applied. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallRules** | [**EdgeFirewallRules**](EdgeFirewallRules.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

