# \ServiceAccountApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceAccount**](ServiceAccountApi.md#DeleteServiceAccount) | **Delete** /1.0.0/serviceAccounts/{id} | Deletes a service account
[**GetServiceAccount**](ServiceAccountApi.md#GetServiceAccount) | **Get** /1.0.0/serviceAccounts/{id} | Retrieves a specific service account
[**QueryServiceAccounts**](ServiceAccountApi.md#QueryServiceAccounts) | **Get** /1.0.0/serviceAccounts | Retrieve service accounts
[**RevokeServiceAccount**](ServiceAccountApi.md#RevokeServiceAccount) | **Post** /1.0.0/serviceAccounts/{id}/revoke | Revokes the token associated with the given service account.
[**TakeOwnership**](ServiceAccountApi.md#TakeOwnership) | **Post** /1.0.0/serviceAccounts/{id}/takeOwnership | Transfer ownership of this service account&#39;s owned entities to the caller.
[**UpdateServiceAccount**](ServiceAccountApi.md#UpdateServiceAccount) | **Put** /1.0.0/serviceAccounts/{id} | Updates a service account


# **DeleteServiceAccount**
> DeleteServiceAccount(ctx, id)
Deletes a service account

Deletes a service account 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Service account ID URN | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceAccount**
> ServiceAccount GetServiceAccount(ctx, id)
Retrieves a specific service account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Service account ID URN | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryServiceAccounts**
> ServiceAccounts QueryServiceAccounts(ctx, page, pageSize, optional)
Retrieve service accounts

Get a list of all service accounts. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***ServiceAccountApiQueryServiceAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAccountApiQueryServiceAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**ServiceAccounts**](ServiceAccounts.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeServiceAccount**
> RevokeServiceAccount(ctx, id)
Revokes the token associated with the given service account.

Revokes the token associated with given service account URN, invalidates any existing sessions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Service account ID URN | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TakeOwnership**
> TakeOwnership(ctx, id)
Transfer ownership of this service account's owned entities to the caller.

Transfer ownership of this service account's owned entities (vApps, media, etc) to the caller. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Service account ID URN | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceAccount**
> ServiceAccount UpdateServiceAccount(ctx, modifiedServiceAccount, id)
Updates a service account

Updates a service account 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modifiedServiceAccount** | [**ServiceAccount**](ServiceAccount.md)|  | 
  **id** | **string**| Service account ID URN | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

