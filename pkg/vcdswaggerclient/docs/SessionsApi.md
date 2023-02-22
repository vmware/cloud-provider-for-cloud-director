# \SessionsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessibleLocations**](SessionsApi.md#GetAccessibleLocations) | **Get** /1.0.0/sessions/{id:((?!provider|current).)*}/accessibleLocations | Get locations accessible to this session.
[**GetCurrentSession**](SessionsApi.md#GetCurrentSession) | **Get** /1.0.0/sessions/current | Returns the current session identified by credentials supplied using the Authorization header
[**GetCurrentSessions**](SessionsApi.md#GetCurrentSessions) | **Get** /1.0.0/sessions | List all sessions for current user
[**GetSession**](SessionsApi.md#GetSession) | **Get** /1.0.0/sessions/{id:((?!provider|current).)*} | Returns the specified session for current user
[**GetToken**](SessionsApi.md#GetToken) | **Get** /1.0.0/sessions/{id:((?!provider|current).)*}/token | Get token associated with this session.
[**Login**](SessionsApi.md#Login) | **Post** /1.0.0/sessions | Logs in a user
[**Logout**](SessionsApi.md#Logout) | **Delete** /1.0.0/sessions/{id:((?!provider|current).)*} | Logs out the current user
[**LogoutCurrentSession**](SessionsApi.md#LogoutCurrentSession) | **Delete** /1.0.0/sessions/current | Logout current session
[**ProviderLogin**](SessionsApi.md#ProviderLogin) | **Post** /1.0.0/sessions/provider | Logs in a user (Provider only)


# **GetAccessibleLocations**
> AccessibleLocations GetAccessibleLocations(ctx, page, pageSize)
Get locations accessible to this session.

Gets locations accessible to this session. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]

### Return type

[**AccessibleLocations**](AccessibleLocations.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentSession**
> Session GetCurrentSession(ctx, )
Returns the current session identified by credentials supplied using the Authorization header

Returns the specified session for the authorization token 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Session**](Session.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentSessions**
> Sessions GetCurrentSessions(ctx, page, pageSize)
List all sessions for current user

List all sessions for current user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]

### Return type

[**Sessions**](Sessions.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSession**
> Session GetSession(ctx, id)
Returns the specified session for current user

Returns the specified session for current user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Session**](Session.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetToken**
> Token GetToken(ctx, id)
Get token associated with this session.

Gets token associated with this session. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Token**](Token.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Login**
> Session Login(ctx, authorization)
Logs in a user

Logs in a user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | [**interface{}**](.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Logout**
> Logout(ctx, id)
Logs out the current user

Logs out the current user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogoutCurrentSession**
> LogoutCurrentSession(ctx, )
Logout current session

Logs out and terminates the current session identified by credentials supplied using the Authorization header 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProviderLogin**
> Session ProviderLogin(ctx, authorization)
Logs in a user (Provider only)

Logs in a user (Provider only) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | [**interface{}**](.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

