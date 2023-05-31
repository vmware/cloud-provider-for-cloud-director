# \UserApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePassword**](UserApi.md#ChangePassword) | **Post** /1.0.0/users/{userUrn}/changePassword | Modify the existing password of a local user.
[**CreateUser**](UserApi.md#CreateUser) | **Post** /1.0.0/users | Create a new user.
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /1.0.0/users/{userUrn} | Delete the specified user.
[**GetUser**](UserApi.md#GetUser) | **Get** /1.0.0/users/{userUrn} | Get a specified user by id. If the id of a service account is specified instead, a simplified view of the service account is returned instead. 
[**QueryUserGroups**](UserApi.md#QueryUserGroups) | **Get** /1.0.0/users/{userUrn}/groups | Get a list of groups that the user with the given id belongs to.
[**QueryUsers**](UserApi.md#QueryUsers) | **Get** /1.0.0/users | Get a list of users.
[**TakeOwnership**](UserApi.md#TakeOwnership) | **Post** /1.0.0/users/{userUrn}/takeOwnership | Transfer ownership of this user&#39;s owned entities to the caller.
[**UpdateUser**](UserApi.md#UpdateUser) | **Put** /1.0.0/users/{userUrn} | Modify the details of a user. A non-administrator user may only modify their own password. An administrator can edit any user.


# **ChangePassword**
> ChangePassword(ctx, passwordChangeRequest, userUrn)
Modify the existing password of a local user.

Modify an existing user's own password 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **passwordChangeRequest** | [**PasswordChangeRequest**](PasswordChangeRequest.md)|  | 
  **userUrn** | **string**| userUrn | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> VcdUser CreateUser(ctx, newUser)
Create a new user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newUser** | [**VcdUser**](VcdUser.md)|  | 

### Return type

[**VcdUser**](VcdUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, userUrn)
Delete the specified user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userUrn** | **string**| userUrn | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> VcdUser GetUser(ctx, userUrn)
Get a specified user by id. If the id of a service account is specified instead, a simplified view of the service account is returned instead. 

Get a specified user by id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userUrn** | **string**| userUrn | 

### Return type

[**VcdUser**](VcdUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryUserGroups**
> EntityReferences QueryUserGroups(ctx, page, pageSize, userUrn, optional)
Get a list of groups that the user with the given id belongs to.

Get a list of references of groups that the user with the given id belongs to. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **userUrn** | **string**| userUrn | 
 **optional** | ***UserApiQueryUserGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiQueryUserGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryUsers**
> VcdUsers QueryUsers(ctx, page, pageSize, optional)
Get a list of users.

Get a list of users. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***UserApiQueryUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiQueryUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VcdUsers**](VcdUsers.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TakeOwnership**
> TakeOwnership(ctx, userUrn)
Transfer ownership of this user's owned entities to the caller.

Transfer ownership of this user's owned entities (vApps, media, etc) to the caller. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userUrn** | **string**| User URN | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> VcdUser UpdateUser(ctx, updatedUser, userUrn)
Modify the details of a user. A non-administrator user may only modify their own password. An administrator can edit any user.

Modify basic details of the specified user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedUser** | [**VcdUser**](VcdUser.md)|  | 
  **userUrn** | **string**| userUrn | 

### Return type

[**VcdUser**](VcdUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

