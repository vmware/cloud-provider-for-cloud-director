# \QuotaPolicyAssignmentApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignQuotaPolicyToGroup**](QuotaPolicyAssignmentApi.md#AssignQuotaPolicyToGroup) | **Put** /1.0.0/groups/{groupUrn}/quotaPolicy | Assign or unassign a quota policy to a given target group. An empty payload is used to unassign an existing quota policy from a given group. 
[**AssignQuotaPolicyToOrg**](QuotaPolicyAssignmentApi.md#AssignQuotaPolicyToOrg) | **Put** /1.0.0/orgs/{orgUrn}/quotaPolicy | Assign or unassign a quota policy to a given target organization. An empty payload is used to unassign an existing quota policy from a given organization. 
[**AssignQuotaPolicyToUser**](QuotaPolicyAssignmentApi.md#AssignQuotaPolicyToUser) | **Put** /1.0.0/users/{userUrn}/quotaPolicy | Assign or unassign a quota policy to a given target user. An empty payload is used to unassign an existing quota policy from a given user. 
[**AssignQuotaPolicyToVApp**](QuotaPolicyAssignmentApi.md#AssignQuotaPolicyToVApp) | **Put** /1.0.0/vapps/{vappUrn}/quotaPolicy | Assign or unassign a quota policy to a given target vapp. An empty payload is used to unassign an existing quota policy from a given vapp. 
[**GetGroupAssignedQuotaPolicy**](QuotaPolicyAssignmentApi.md#GetGroupAssignedQuotaPolicy) | **Get** /1.0.0/groups/{groupUrn}/quotaPolicy | Get the assigned quota policy reference for the given group.
[**GetOrgAssignedQuotaPolicy**](QuotaPolicyAssignmentApi.md#GetOrgAssignedQuotaPolicy) | **Get** /1.0.0/orgs/{orgUrn}/quotaPolicy | Get the assigned quota policy reference for the given organization.
[**GetUserAssignedQuotaPolicy**](QuotaPolicyAssignmentApi.md#GetUserAssignedQuotaPolicy) | **Get** /1.0.0/users/{userUrn}/quotaPolicy | Get the assigned quota policy reference for the given user.
[**GetVAppAssignedQuotaPolicy**](QuotaPolicyAssignmentApi.md#GetVAppAssignedQuotaPolicy) | **Get** /1.0.0/vapps/{vappUrn}/quotaPolicy | Get the assigned quota policy reference for the given vapp.


# **AssignQuotaPolicyToGroup**
> AssignQuotaPolicyToGroup(ctx, quotaPolicyReference, groupUrn)
Assign or unassign a quota policy to a given target group. An empty payload is used to unassign an existing quota policy from a given group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quotaPolicyReference** | [**AssignedQuotaPolicy**](AssignedQuotaPolicy.md)|  | 
  **groupUrn** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignQuotaPolicyToOrg**
> AssignQuotaPolicyToOrg(ctx, quotaPolicyReference, orgUrn)
Assign or unassign a quota policy to a given target organization. An empty payload is used to unassign an existing quota policy from a given organization. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quotaPolicyReference** | [**AssignedQuotaPolicy**](AssignedQuotaPolicy.md)|  | 
  **orgUrn** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignQuotaPolicyToUser**
> AssignQuotaPolicyToUser(ctx, quotaPolicyReference, userUrn)
Assign or unassign a quota policy to a given target user. An empty payload is used to unassign an existing quota policy from a given user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quotaPolicyReference** | [**AssignedQuotaPolicy**](AssignedQuotaPolicy.md)|  | 
  **userUrn** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignQuotaPolicyToVApp**
> AssignQuotaPolicyToVApp(ctx, quotaPolicyReference, vappUrn)
Assign or unassign a quota policy to a given target vapp. An empty payload is used to unassign an existing quota policy from a given vapp. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quotaPolicyReference** | [**AssignedQuotaPolicy**](AssignedQuotaPolicy.md)|  | 
  **vappUrn** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupAssignedQuotaPolicy**
> AssignedQuotaPolicy GetGroupAssignedQuotaPolicy(ctx, groupUrn)
Get the assigned quota policy reference for the given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupUrn** | **string**|  | 

### Return type

[**AssignedQuotaPolicy**](AssignedQuotaPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgAssignedQuotaPolicy**
> AssignedQuotaPolicy GetOrgAssignedQuotaPolicy(ctx, orgUrn)
Get the assigned quota policy reference for the given organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgUrn** | **string**|  | 

### Return type

[**AssignedQuotaPolicy**](AssignedQuotaPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserAssignedQuotaPolicy**
> AssignedQuotaPolicy GetUserAssignedQuotaPolicy(ctx, userUrn)
Get the assigned quota policy reference for the given user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userUrn** | **string**|  | 

### Return type

[**AssignedQuotaPolicy**](AssignedQuotaPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVAppAssignedQuotaPolicy**
> AssignedQuotaPolicy GetVAppAssignedQuotaPolicy(ctx, vappUrn)
Get the assigned quota policy reference for the given vapp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vappUrn** | **string**|  | 

### Return type

[**AssignedQuotaPolicy**](AssignedQuotaPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

