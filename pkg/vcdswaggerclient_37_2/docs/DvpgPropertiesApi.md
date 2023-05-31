# \DvpgPropertiesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDvpgProperties**](DvpgPropertiesApi.md#GetDvpgProperties) | **Get** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dvpgProperties | Returns the DVPG properties, such as promiscuous mode and forged transmit, of a specific Org vDC network. This is a SysAdmin only API.
[**UpdateDvpgProperties**](DvpgPropertiesApi.md#UpdateDvpgProperties) | **Put** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dvpgProperties | Toggles the DVPG properties of a specific Org vDC network. Allows for update of promiscuous mode and forged transmit. This is a SysAdmin only API.


# **GetDvpgProperties**
> DvpgProperties GetDvpgProperties(ctx, vdcNetworkId)
Returns the DVPG properties, such as promiscuous mode and forged transmit, of a specific Org vDC network. This is a SysAdmin only API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

[**DvpgProperties**](DvpgProperties.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDvpgProperties**
> UpdateDvpgProperties(ctx, dvpgProperties, vdcNetworkId)
Toggles the DVPG properties of a specific Org vDC network. Allows for update of promiscuous mode and forged transmit. This is a SysAdmin only API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dvpgProperties** | [**DvpgProperties**](DvpgProperties.md)|  | 
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

