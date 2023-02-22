# \LogicalVmGroupsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNamedVmGroupsToLogicalVmGroup**](LogicalVmGroupsApi.md#AddNamedVmGroupsToLogicalVmGroup) | **Post** /1.0.0/logicalVmGroups/{logicalVmGroupId}/namedVmGroups | Assign named vm groups to logical vm group.
[**CreateLogicalVmGroup**](LogicalVmGroupsApi.md#CreateLogicalVmGroup) | **Post** /1.0.0/logicalVmGroups | Creates a new logical vm group
[**DeleteLogicalVmGroup**](LogicalVmGroupsApi.md#DeleteLogicalVmGroup) | **Delete** /1.0.0/logicalVmGroups/{logicalVmGroupId} | Delete specified logical vm group.
[**GetLogicalVmGroup**](LogicalVmGroupsApi.md#GetLogicalVmGroup) | **Get** /1.0.0/logicalVmGroups/{logicalVmGroupId} | Get specified logical vm group
[**GetLogicalVmGroupNamedVmGroups**](LogicalVmGroupsApi.md#GetLogicalVmGroupNamedVmGroups) | **Get** /1.0.0/logicalVmGroups/{logicalVmGroupId}/namedVmGroups | Get all named vm groups associated with logical vm group
[**GetPvdcPoliciesForLogicalVmGroup**](LogicalVmGroupsApi.md#GetPvdcPoliciesForLogicalVmGroup) | **Get** /1.0.0/logicalVmGroups/{logicalVmGroupId}/pvdcPolicies | Get all pvdc policies associated with logical vm group
[**QueryLogicalVmGroups**](LogicalVmGroupsApi.md#QueryLogicalVmGroups) | **Get** /1.0.0/logicalVmGroups | Get list of logical vm groups.
[**UpdateLogicalVmGroup**](LogicalVmGroupsApi.md#UpdateLogicalVmGroup) | **Put** /1.0.0/logicalVmGroups/{logicalVmGroupId} | Update specified logical vm group


# **AddNamedVmGroupsToLogicalVmGroup**
> []EntityReference AddNamedVmGroupsToLogicalVmGroup(ctx, namedVmGroupRefs, logicalVmGroupId)
Assign named vm groups to logical vm group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namedVmGroupRefs** | [**[]EntityReference**](EntityReference.md)|  | 
  **logicalVmGroupId** | **string**|  | 

### Return type

[**[]EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLogicalVmGroup**
> LogicalVmGroup CreateLogicalVmGroup(ctx, newLogicalVmGroupParams)
Creates a new logical vm group

Creates a new logical vm group 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newLogicalVmGroupParams** | [**LogicalVmGroup**](LogicalVmGroup.md)|  | 

### Return type

[**LogicalVmGroup**](LogicalVmGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLogicalVmGroup**
> DeleteLogicalVmGroup(ctx, logicalVmGroupId)
Delete specified logical vm group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalVmGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalVmGroup**
> LogicalVmGroup GetLogicalVmGroup(ctx, logicalVmGroupId)
Get specified logical vm group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalVmGroupId** | **string**|  | 

### Return type

[**LogicalVmGroup**](LogicalVmGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalVmGroupNamedVmGroups**
> EntityReferences GetLogicalVmGroupNamedVmGroups(ctx, page, pageSize, logicalVmGroupId, optional)
Get all named vm groups associated with logical vm group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **logicalVmGroupId** | **string**|  | 
 **optional** | ***LogicalVmGroupsApiGetLogicalVmGroupNamedVmGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogicalVmGroupsApiGetLogicalVmGroupNamedVmGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPvdcPoliciesForLogicalVmGroup**
> EntityReferences GetPvdcPoliciesForLogicalVmGroup(ctx, page, pageSize, logicalVmGroupId, optional)
Get all pvdc policies associated with logical vm group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **logicalVmGroupId** | **string**|  | 
 **optional** | ***LogicalVmGroupsApiGetPvdcPoliciesForLogicalVmGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogicalVmGroupsApiGetPvdcPoliciesForLogicalVmGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryLogicalVmGroups**
> LogicalVmGroups QueryLogicalVmGroups(ctx, page, pageSize, optional)
Get list of logical vm groups.

Get list of logical vm groups. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***LogicalVmGroupsApiQueryLogicalVmGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogicalVmGroupsApiQueryLogicalVmGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**LogicalVmGroups**](LogicalVmGroups.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLogicalVmGroup**
> LogicalVmGroup UpdateLogicalVmGroup(ctx, updateLogicalVmGroupParams, logicalVmGroupId)
Update specified logical vm group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateLogicalVmGroupParams** | [**LogicalVmGroup**](LogicalVmGroup.md)|  | 
  **logicalVmGroupId** | **string**|  | 

### Return type

[**LogicalVmGroup**](LogicalVmGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

