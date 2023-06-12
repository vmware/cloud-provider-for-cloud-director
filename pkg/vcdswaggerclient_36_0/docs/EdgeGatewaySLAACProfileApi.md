# \EdgeGatewaySLAACProfileApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSLAACProfile**](EdgeGatewaySLAACProfileApi.md#GetSLAACProfile) | **Get** /1.0.0/edgeGateways/{gatewayId}/slaacProfile | Retrieves the SLAAC profile on the edge gateway. 
[**UpdateSLAACProfile**](EdgeGatewaySLAACProfileApi.md#UpdateSLAACProfile) | **Put** /1.0.0/edgeGateways/{gatewayId}/slaacProfile | Creates a SLAAC profile or updates the existing one if it already exists.


# **GetSLAACProfile**
> SlaacProfile GetSLAACProfile(ctx, gatewayId)
Retrieves the SLAAC profile on the edge gateway. 

Retrieves the SLAAC profile on the edge gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**|  | 

### Return type

[**SlaacProfile**](SLAACProfile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSLAACProfile**
> UpdateSLAACProfile(ctx, sLAACProfile, gatewayId)
Creates a SLAAC profile or updates the existing one if it already exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sLAACProfile** | [**SlaacProfile**](SlaacProfile.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

