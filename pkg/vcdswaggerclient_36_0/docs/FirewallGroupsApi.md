# \FirewallGroupsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirewallGroup**](FirewallGroupsApi.md#CreateFirewallGroup) | **Post** /1.0.0/firewallGroups | Create a firewall group
[**GetFirewallGroups**](FirewallGroupsApi.md#GetFirewallGroups) | **Get** /1.0.0/firewallGroups/summaries | Retrieves the Firewall Groups.


# **CreateFirewallGroup**
> CreateFirewallGroup(ctx, firewallGroup)
Create a firewall group

Create a firewall group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallGroup** | [**FirewallGroupDetails**](FirewallGroupDetails.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallGroups**
> FirewallGroups GetFirewallGroups(ctx, page, pageSize, optional)
Retrieves the Firewall Groups.

Get all firewall groups. Results can be filtered by name and context (_context). Supported contexts are:   <ul>   <li>Org Vdc Network ID <code>(_context==networkId)</code> -   Returns all the firewall groups which the specified network is a member of.   <li>Edge Gateway ID <code>(_context==edgeGatewayId)</code> -   Returns all the firewall groups which are available to the specific edge gateway.   <li>Network Provider ID <code>(_context==networkProviderId)</code> -   Returns all the firewall groups which are available under a specific network provider. This context requires system admin privilege.   </ul> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***FirewallGroupsApiGetFirewallGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallGroupsApiGetFirewallGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**FirewallGroups**](FirewallGroups.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

