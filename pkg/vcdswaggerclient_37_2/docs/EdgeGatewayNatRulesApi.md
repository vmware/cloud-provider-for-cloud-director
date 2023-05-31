# \EdgeGatewayNatRulesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNatRule**](EdgeGatewayNatRulesApi.md#CreateNatRule) | **Post** /1.0.0/edgeGateways/{gatewayId}/nat/rules | Creates a NAT Rule on the Edge Gateway.
[**GetNatRules**](EdgeGatewayNatRulesApi.md#GetNatRules) | **Get** /1.0.0/edgeGateways/{gatewayId}/nat/rules | Retrieves all NAT Rules on the edge gateway. 


# **CreateNatRule**
> CreateNatRule(ctx, edgeNatRule, gatewayId)
Creates a NAT Rule on the Edge Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeNatRule** | [**EdgeNatRule**](EdgeNatRule.md)|  | 
  **gatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNatRules**
> EdgeNatRules GetNatRules(ctx, pageSize, gatewayId, optional)
Retrieves all NAT Rules on the edge gateway. 

Retrieves all NAT Rules on the edge gateway.  Pagination is supported to get the next page in the header response. Results can be sorted by only a single parameter. Sorting by combination of parameters (sortAsc=foo&sortDesc=bar) is not allowed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **gatewayId** | **string**|  | 
 **optional** | ***EdgeGatewayNatRulesApiGetNatRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayNatRulesApiGetNatRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EdgeNatRules**](EdgeNatRules.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

