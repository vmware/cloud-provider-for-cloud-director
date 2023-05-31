# \CapabilitiesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgVdcStoragePolicyCapabilities**](CapabilitiesApi.md#GetOrgVdcStoragePolicyCapabilities) | **Get** /1.0.0/orgVdcStoragePolicies/{id}/capabilities | Retrieves capabilities of a specific organization VDC storage policy.
[**GetPvdcStoragePolicyCapabilities**](CapabilitiesApi.md#GetPvdcStoragePolicyCapabilities) | **Get** /1.0.0/pvdcStoragePolicies/{id}/capabilities | Retrieves capabilities of a specific provider VDC storage policy.
[**GetQuotaPolicyCapabilities**](CapabilitiesApi.md#GetQuotaPolicyCapabilities) | **Get** /1.0.0/quotaPolicy/capabilities | Retrieves capabilities for quotaPolicy feature.
[**GetVdcCapabilities**](CapabilitiesApi.md#GetVdcCapabilities) | **Get** /1.0.0/vdcs/{orgVdcId}/capabilities | Retrieves capabilities for the given Organization vDC.
[**GetVdcGroupCapabilities**](CapabilitiesApi.md#GetVdcGroupCapabilities) | **Get** /1.0.0/vdcGroups/{vdcGroupId}/capabilities | Retrieves the supported capabilities of the specified vDC Group.


# **GetOrgVdcStoragePolicyCapabilities**
> Capabilities GetOrgVdcStoragePolicyCapabilities(ctx, id)
Retrieves capabilities of a specific organization VDC storage policy.

Retrieves the current capabilities configured on a specific organization VDC storage policy. These cannot be edited. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Capabilities**](Capabilities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPvdcStoragePolicyCapabilities**
> Capabilities GetPvdcStoragePolicyCapabilities(ctx, id)
Retrieves capabilities of a specific provider VDC storage policy.

Retrieves the current capabilities configured on a specific provider VDC storage policy. These cannot be edited. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Capabilities**](Capabilities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotaPolicyCapabilities**
> Capabilities GetQuotaPolicyCapabilities(ctx, )
Retrieves capabilities for quotaPolicy feature.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Capabilities**](Capabilities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcCapabilities**
> Capabilities GetVdcCapabilities(ctx, orgVdcId)
Retrieves capabilities for the given Organization vDC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgVdcId** | **string**|  | 

### Return type

[**Capabilities**](Capabilities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcGroupCapabilities**
> Capabilities GetVdcGroupCapabilities(ctx, vdcGroupId)
Retrieves the supported capabilities of the specified vDC Group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 

### Return type

[**Capabilities**](Capabilities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

