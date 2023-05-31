# \UniversalRouterHealthApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUniversalRouterHealth**](UniversalRouterHealthApi.md#GetUniversalRouterHealth) | **Get** /1.0.0/universalRouters/{universalRouterId}/health | Get Health information of a universal router


# **GetUniversalRouterHealth**
> RouterHealthReport GetUniversalRouterHealth(ctx, universalRouterId)
Get Health information of a universal router

Get Health information of a universal router and its associated entities such as egress points and routing. It includes information about reachability status of all the participating vDC's of referenced vDC group and it also includes information about all the nsx managers covered by this universal router along with associated nsx controller cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 

### Return type

[**RouterHealthReport**](RouterHealthReport.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

