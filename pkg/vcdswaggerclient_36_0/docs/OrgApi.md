# \OrgApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrg**](OrgApi.md#CreateOrg) | **Post** /1.0.0/orgs | Create a new organization
[**DeleteOrg**](OrgApi.md#DeleteOrg) | **Delete** /1.0.0/orgs/{orgUrn} | Delete the specified org.
[**GetOrg**](OrgApi.md#GetOrg) | **Get** /1.0.0/orgs/{orgUrn} | Get specified organization.
[**QueryOrgs**](OrgApi.md#QueryOrgs) | **Get** /1.0.0/orgs | Get list of all orgs accessible to the user.
[**UpdateOrg**](OrgApi.md#UpdateOrg) | **Put** /1.0.0/orgs/{orgUrn} | Updates an organization.


# **CreateOrg**
> Org CreateOrg(ctx, newOrg)
Create a new organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newOrg** | [**Org**](Org.md)|  | 

### Return type

[**Org**](Org.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrg**
> DeleteOrg(ctx, orgUrn, optional)
Delete the specified org.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgUrn** | **string**| orgUrn | 
 **optional** | ***OrgApiDeleteOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgApiDeleteOrgOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Value \&quot;true\&quot; means to forcefully delete the object that contains other objects even if those objects are in a state that does not allow removal. The default is \&quot;false\&quot;; therefore, objects are not removed if they are not in a state that normally allows removal. Force also implies recursive delete where other contained objects are removed. Errors may be ignored. Invalid value (not true or false) are ignored.  | [default to false]
 **recursive** | **optional.Bool**| Value \&quot;true\&quot; means to remove an object that contains other objects as long as those objects are in a state that normally allows removal. The default is \&quot;false\&quot;, which means that an exception will be thrown if any of those objects are in a state that does not allow removal. Invalid value (not true or false) are ignored.  | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrg**
> Org GetOrg(ctx, orgUrn)
Get specified organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgUrn** | **string**| orgUrn | 

### Return type

[**Org**](Org.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryOrgs**
> Orgs QueryOrgs(ctx, page, pageSize, optional)
Get list of all orgs accessible to the user.

Get list of all orgs accessible to the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***OrgApiQueryOrgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgApiQueryOrgsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Orgs**](Orgs.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrg**
> Org UpdateOrg(ctx, modifiedOrg, orgUrn)
Updates an organization.

Updates an organization. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modifiedOrg** | [**Org**](Org.md)|  | 
  **orgUrn** | **string**| orgUrn | 

### Return type

[**Org**](Org.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

