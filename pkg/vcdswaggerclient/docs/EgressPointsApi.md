# \EgressPointsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEgressPoint**](EgressPointsApi.md#CreateEgressPoint) | **Post** /1.0.0/universalRouters/{universalRouterId}/routing/egressPoints | Creates a new Universal Egress Point from a specified Edge Gateway and Org vDC.
[**GetEgressPoints**](EgressPointsApi.md#GetEgressPoints) | **Get** /1.0.0/universalRouters/{universalRouterId}/routing/egressPoints | Get a list of Universal Egress Points for a Universal Router.


# **CreateEgressPoint**
> CreateEgressPoint(ctx, egressPoint, universalRouterId)
Creates a new Universal Egress Point from a specified Edge Gateway and Org vDC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **egressPoint** | [**EgressPoint**](EgressPoint.md)|  | 
  **universalRouterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEgressPoints**
> EgressPoints GetEgressPoints(ctx, universalRouterId)
Get a list of Universal Egress Points for a Universal Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 

### Return type

[**EgressPoints**](EgressPoints.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

