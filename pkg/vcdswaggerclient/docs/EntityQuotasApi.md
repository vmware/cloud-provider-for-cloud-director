# \EntityQuotasApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignQuotasToGroup**](EntityQuotasApi.md#AssignQuotasToGroup) | **Put** /1.0.0/groups/{groupUrn}/quotas | Assign quotas to given target group.
[**AssignQuotasToOrg**](EntityQuotasApi.md#AssignQuotasToOrg) | **Put** /1.0.0/orgs/{orgUrn}/quotas | Assign quotas to given target organization.
[**AssignQuotasToUser**](EntityQuotasApi.md#AssignQuotasToUser) | **Put** /1.0.0/users/{userUrn}/quotas | Assign quotas to given target user.
[**GetQuotasForGroup**](EntityQuotasApi.md#GetQuotasForGroup) | **Get** /1.0.0/groups/{groupUrn}/quotas | Get the effective quotas applicable for the given group.
[**GetQuotasForOrg**](EntityQuotasApi.md#GetQuotasForOrg) | **Get** /1.0.0/orgs/{orgUrn}/quotas | Get the effective quotas applicable for the given organization.
[**GetQuotasForUser**](EntityQuotasApi.md#GetQuotasForUser) | **Get** /1.0.0/users/{userUrn}/quotas | Get the effective quotas applicable for the given user.
[**GetQuotasForVApp**](EntityQuotasApi.md#GetQuotasForVApp) | **Get** /1.0.0/vapps/{vappUrn}/quotas | Get the effective quotas applicable for the given vapp.


# **AssignQuotasToGroup**
> AssignQuotasToGroup(ctx, quotas, groupUrn)
Assign quotas to given target group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quotas** | [**Quotas**](Quotas.md)|  | 
  **groupUrn** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignQuotasToOrg**
> AssignQuotasToOrg(ctx, quotas, orgUrn)
Assign quotas to given target organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quotas** | [**Quotas**](Quotas.md)|  | 
  **orgUrn** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignQuotasToUser**
> AssignQuotasToUser(ctx, quotas, userUrn)
Assign quotas to given target user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quotas** | [**Quotas**](Quotas.md)|  | 
  **userUrn** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotasForGroup**
> Quotas GetQuotasForGroup(ctx, groupUrn)
Get the effective quotas applicable for the given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupUrn** | **string**|  | 

### Return type

[**Quotas**](Quotas.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotasForOrg**
> Quotas GetQuotasForOrg(ctx, orgUrn)
Get the effective quotas applicable for the given organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgUrn** | **string**|  | 

### Return type

[**Quotas**](Quotas.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotasForUser**
> Quotas GetQuotasForUser(ctx, userUrn)
Get the effective quotas applicable for the given user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userUrn** | **string**|  | 

### Return type

[**Quotas**](Quotas.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotasForVApp**
> Quotas GetQuotasForVApp(ctx, vappUrn)
Get the effective quotas applicable for the given vapp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vappUrn** | **string**|  | 

### Return type

[**Quotas**](Quotas.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

