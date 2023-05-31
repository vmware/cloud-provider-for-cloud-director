# \LDAPApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchLdapGroups**](LDAPApi.md#SearchLdapGroups) | **Get** /1.0.0/ldap/search/group | Search LDAP Groups
[**SearchLdapUsers**](LDAPApi.md#SearchLdapUsers) | **Get** /1.0.0/ldap/search/user | Search LDAP Users
[**SyncLdap**](LDAPApi.md#SyncLdap) | **Post** /1.0.0/ldap/sync | Synchronize LDAP users/settings
[**TestLdap**](LDAPApi.md#TestLdap) | **Post** /1.0.0/ldap/test | Tests LDAP Connection and Settings


# **SearchLdapGroups**
> []UserGroup SearchLdapGroups(ctx, optional)
Search LDAP Groups

Searches LDAP for given group(s) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LDAPApiSearchLdapGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LDAPApiSearchLdapGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| String to search for via LDAP | 

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLdapUsers**
> []User SearchLdapUsers(ctx, optional)
Search LDAP Users

Searches LDAP for given user(s) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LDAPApiSearchLdapUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LDAPApiSearchLdapUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| String to search for via LDAP | 

### Return type

[**[]User**](User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncLdap**
> SyncLdap(ctx, )
Synchronize LDAP users/settings

Begins the LDAP sync task 

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

# **TestLdap**
> LdapTestResult TestLdap(ctx, ldapSettings, optional)
Tests LDAP Connection and Settings

Tests that custom LDAP settings are valid, and that the system can use them to search for a user or group 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapSettings** | [**LdapSettings**](LdapSettings.md)| LDAP Settings to use for testing connection | 
 **optional** | ***LDAPApiTestLdapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LDAPApiTestLdapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **optional.String**| Username to use when testing LDAP search | 

### Return type

[**LdapTestResult**](LdapTestResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

