# \QuotaPoliciesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuotaPolicy**](QuotaPoliciesApi.md#CreateQuotaPolicy) | **Post** /1.0.0/quotaPolicies | Creates a new quota policy
[**DeleteQuotaPolicy**](QuotaPoliciesApi.md#DeleteQuotaPolicy) | **Delete** /1.0.0/quotaPolicies/{quotaPolicyId} | Delete the specified quota policy.
[**GetQuotaPolicy**](QuotaPoliciesApi.md#GetQuotaPolicy) | **Get** /1.0.0/quotaPolicies/{quotaPolicyId} | Get the specified quota policy
[**QueryQuotaPolicies**](QuotaPoliciesApi.md#QueryQuotaPolicies) | **Get** /1.0.0/quotaPolicies | Get list of quota policies.
[**UpdateQuotaPolicy**](QuotaPoliciesApi.md#UpdateQuotaPolicy) | **Put** /1.0.0/quotaPolicies/{quotaPolicyId} | Update the specified quota policy


# **CreateQuotaPolicy**
> QuotaPolicy CreateQuotaPolicy(ctx, quotaPolicyParams)
Creates a new quota policy

Creates a new quota policy 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quotaPolicyParams** | [**QuotaPolicy**](QuotaPolicy.md)|  | 

### Return type

[**QuotaPolicy**](QuotaPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuotaPolicy**
> DeleteQuotaPolicy(ctx, quotaPolicyId)
Delete the specified quota policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quotaPolicyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotaPolicy**
> QuotaPolicy GetQuotaPolicy(ctx, quotaPolicyId)
Get the specified quota policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quotaPolicyId** | **string**|  | 

### Return type

[**QuotaPolicy**](QuotaPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryQuotaPolicies**
> QuotaPolicies QueryQuotaPolicies(ctx, page, pageSize, optional)
Get list of quota policies.

Get list of quota policies. Results can be filtered by id Returns all the quota policies which are available in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***QuotaPoliciesApiQueryQuotaPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuotaPoliciesApiQueryQuotaPoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**QuotaPolicies**](QuotaPolicies.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuotaPolicy**
> QuotaPolicy UpdateQuotaPolicy(ctx, updateQuotaPolicyParams, quotaPolicyId)
Update the specified quota policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateQuotaPolicyParams** | [**QuotaPolicy**](QuotaPolicy.md)|  | 
  **quotaPolicyId** | **string**|  | 

### Return type

[**QuotaPolicy**](QuotaPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

