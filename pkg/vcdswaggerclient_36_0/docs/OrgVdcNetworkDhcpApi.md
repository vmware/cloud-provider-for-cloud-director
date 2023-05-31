# \OrgVdcNetworkDhcpApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDhcpBinding**](OrgVdcNetworkDhcpApi.md#CreateDhcpBinding) | **Post** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp/bindings | Creates a DHCP binding on an Org vDC Network.
[**DeleteDhcpBinding**](OrgVdcNetworkDhcpApi.md#DeleteDhcpBinding) | **Delete** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp/bindings/{bindingId} | Delete a specific DHCP binding of the Org vDC Network.
[**DeleteNetworkDhcpConfig**](OrgVdcNetworkDhcpApi.md#DeleteNetworkDhcpConfig) | **Delete** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp | Removes Dhcp configuration on a specific Org vDC network.
[**GetDhcpBinding**](OrgVdcNetworkDhcpApi.md#GetDhcpBinding) | **Get** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp/bindings/{bindingId} | Retrieve a specific DHCP binding of the Org vDC Network.
[**GetDhcpBindings**](OrgVdcNetworkDhcpApi.md#GetDhcpBindings) | **Get** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp/bindings | Retrieves all DHCP bindings for an Org vDC Network. 
[**GetNetworkDhcpConfig**](OrgVdcNetworkDhcpApi.md#GetNetworkDhcpConfig) | **Get** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp | Retrieves Dhcp configuration of a specific Org vDC network.
[**UpdateDhcpBinding**](OrgVdcNetworkDhcpApi.md#UpdateDhcpBinding) | **Put** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp/bindings/{bindingId} | Update a specific DHCP binding of the Org vDC Network.
[**UpdateNetworkDhcpConfig**](OrgVdcNetworkDhcpApi.md#UpdateNetworkDhcpConfig) | **Put** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp | Updates Dhcp configuration of a specific Org vDC network.


# **CreateDhcpBinding**
> CreateDhcpBinding(ctx, dhcpBinding, vdcNetworkId)
Creates a DHCP binding on an Org vDC Network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dhcpBinding** | [**DhcpBinding**](DhcpBinding.md)|  | 
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDhcpBinding**
> DeleteDhcpBinding(ctx, vdcNetworkId, bindingId)
Delete a specific DHCP binding of the Org vDC Network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 
  **bindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkDhcpConfig**
> DeleteNetworkDhcpConfig(ctx, vdcNetworkId)
Removes Dhcp configuration on a specific Org vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDhcpBinding**
> DhcpBinding GetDhcpBinding(ctx, vdcNetworkId, bindingId)
Retrieve a specific DHCP binding of the Org vDC Network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 
  **bindingId** | **string**|  | 

### Return type

[**DhcpBinding**](DhcpBinding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDhcpBindings**
> DhcpBindings GetDhcpBindings(ctx, pageSize, vdcNetworkId, optional)
Retrieves all DHCP bindings for an Org vDC Network. 

Retrieves all DHCP bindings for an Org vDC Network. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **vdcNetworkId** | **string**|  | 
 **optional** | ***OrgVdcNetworkDhcpApiGetDhcpBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgVdcNetworkDhcpApiGetDhcpBindingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 

### Return type

[**DhcpBindings**](DhcpBindings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkDhcpConfig**
> VdcNetworkDhcpConfig GetNetworkDhcpConfig(ctx, vdcNetworkId)
Retrieves Dhcp configuration of a specific Org vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

[**VdcNetworkDhcpConfig**](VdcNetworkDhcpConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpBinding**
> UpdateDhcpBinding(ctx, dhcpBinding, vdcNetworkId, bindingId)
Update a specific DHCP binding of the Org vDC Network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dhcpBinding** | [**DhcpBinding**](DhcpBinding.md)|  | 
  **vdcNetworkId** | **string**|  | 
  **bindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkDhcpConfig**
> UpdateNetworkDhcpConfig(ctx, dhcpConfig, vdcNetworkId)
Updates Dhcp configuration of a specific Org vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dhcpConfig** | [**VdcNetworkDhcpConfig**](VdcNetworkDhcpConfig.md)|  | 
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

