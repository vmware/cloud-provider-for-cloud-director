# \FirewallGroupApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFirewallGroup**](FirewallGroupApi.md#DeleteFirewallGroup) | **Delete** /1.0.0/firewallGroups/{firewallGroupId} | Deletes a Firewall Group.
[**GetFirewallGroup**](FirewallGroupApi.md#GetFirewallGroup) | **Get** /1.0.0/firewallGroups/{firewallGroupId} | Retrieves a specific firewall group.
[**GetFirewallGroupAssociatedVMs**](FirewallGroupApi.md#GetFirewallGroupAssociatedVMs) | **Get** /1.0.0/firewallGroups/{firewallGroupId}/associatedVMs | Retrieves associated VMs for a specific firewall group.
[**UpdateFirewallGroup**](FirewallGroupApi.md#UpdateFirewallGroup) | **Put** /1.0.0/firewallGroups/{firewallGroupId} | Updates the Firewall Group.


# **DeleteFirewallGroup**
> DeleteFirewallGroup(ctx, firewallGroupId)
Deletes a Firewall Group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallGroup**
> FirewallGroupDetails GetFirewallGroup(ctx, firewallGroupId)
Retrieves a specific firewall group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallGroupId** | **string**|  | 

### Return type

[**FirewallGroupDetails**](FirewallGroupDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallGroupAssociatedVMs**
> FirewallGroupAssociatedVMs GetFirewallGroupAssociatedVMs(ctx, page, pageSize, firewallGroupId, optional)
Retrieves associated VMs for a specific firewall group.

Get all associated VMs for a specific firewall group. Associated VM members can only be obtained for firewall groups with typeValue <em>STATIC_MEMBERS</em> or <em>VM_CRITERIA</em>. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **firewallGroupId** | **string**|  | 
 **optional** | ***FirewallGroupApiGetFirewallGroupAssociatedVMsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallGroupApiGetFirewallGroupAssociatedVMsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**FirewallGroupAssociatedVMs**](FirewallGroupAssociatedVMs.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFirewallGroup**
> UpdateFirewallGroup(ctx, firewallGroup, firewallGroupId)
Updates the Firewall Group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallGroup** | [**FirewallGroupDetails**](FirewallGroupDetails.md)|  | 
  **firewallGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

