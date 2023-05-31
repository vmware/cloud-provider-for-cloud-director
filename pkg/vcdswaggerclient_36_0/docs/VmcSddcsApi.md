# \VmcSddcsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVmcSddcs**](VmcSddcsApi.md#GetVmcSddcs) | **Get** /1.0.0/serviceApps/{serviceAppId}/sddcs | Retrieves a list of SDDCs for a specific VMware service application that are available
[**RegisterVmcSDDC**](VmcSddcsApi.md#RegisterVmcSDDC) | **Post** /1.0.0/serviceApps/{serviceAppId}/sddcs | Register a VMware service application SDDC&#39;s components


# **GetVmcSddcs**
> VmcSddcs GetVmcSddcs(ctx, serviceAppId)
Retrieves a list of SDDCs for a specific VMware service application that are available

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceAppId** | **string**|  | 

### Return type

[**VmcSddcs**](VmcSddcs.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterVmcSDDC**
> RegisterVmcSDDC(ctx, serviceAppSddc, serviceAppId)
Register a VMware service application SDDC's components

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceAppSddc** | [**VmcSddc**](VmcSddc.md)|  | 
  **serviceAppId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

