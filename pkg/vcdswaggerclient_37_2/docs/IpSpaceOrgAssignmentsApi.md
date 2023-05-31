# \IpSpaceOrgAssignmentsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpSpaceOrgAssignment**](IpSpaceOrgAssignmentsApi.md#CreateIpSpaceOrgAssignment) | **Post** /1.0.0/ipSpaces/orgAssignments | Create a new IP Space Org Assignment.
[**DeleteIpSpaceOrgAssignment**](IpSpaceOrgAssignmentsApi.md#DeleteIpSpaceOrgAssignment) | **Delete** /1.0.0/ipSpaces/orgAssignments/{assignmentId} | Delete an IP Space Org Assignment.
[**GetIpSpaceOrgAssignment**](IpSpaceOrgAssignmentsApi.md#GetIpSpaceOrgAssignment) | **Get** /1.0.0/ipSpaces/orgAssignments/{assignmentId} | Get a an IP Space Org Assignment.
[**GetIpSpaceOrgAssignments**](IpSpaceOrgAssignmentsApi.md#GetIpSpaceOrgAssignments) | **Get** /1.0.0/ipSpaces/orgAssignments | Get the Org assignments for an IP Space.
[**UpdateIpSpaceOrgAssignment**](IpSpaceOrgAssignmentsApi.md#UpdateIpSpaceOrgAssignment) | **Put** /1.0.0/ipSpaces/orgAssignments/{assignmentId} | Update an IP Space Org Assignment.


# **CreateIpSpaceOrgAssignment**
> CreateIpSpaceOrgAssignment(ctx, assignment)
Create a new IP Space Org Assignment.

Create a new IP Space Org Assignment. The assignment links an Organization to an IP Space by providing the Organization access to the specified IP Space. It also defines the various IP Space quotas applied to an Organization. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assignment** | [**IpSpaceOrgAssignment**](IpSpaceOrgAssignment.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpSpaceOrgAssignment**
> DeleteIpSpaceOrgAssignment(ctx, assignmentId)
Delete an IP Space Org Assignment.

Deletes a specific IP Space Org Assignment. The organization will no longer be able to use the IP Space. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assignmentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSpaceOrgAssignment**
> IpSpaceOrgAssignment GetIpSpaceOrgAssignment(ctx, assignmentId)
Get a an IP Space Org Assignment.

Retrieves a specific IP Space Org Assignment. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assignmentId** | **string**|  | 

### Return type

[**IpSpaceOrgAssignment**](IpSpaceOrgAssignment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSpaceOrgAssignments**
> IpSpaceOrgAssignments GetIpSpaceOrgAssignments(ctx, page, pageSize, optional)
Get the Org assignments for an IP Space.

Retrieves the Org assignments for an IP Space. Either \"ipSpaceRef\" or \"orgRef\" filter is required. \"ipSpaceRef\" filter can also be combined with onlyIncludeCustomQuotas==true filter to filter out only those Org Assignments where custom quotas are applied. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***IpSpaceOrgAssignmentsApiGetIpSpaceOrgAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IpSpaceOrgAssignmentsApiGetIpSpaceOrgAssignmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**IpSpaceOrgAssignments**](IpSpaceOrgAssignments.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpSpaceOrgAssignment**
> UpdateIpSpaceOrgAssignment(ctx, assignment, assignmentId)
Update an IP Space Org Assignment.

Updates a specific IP Space Org Assignment. Only custom quotas applied to Organization can be modified. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assignment** | [**IpSpaceOrgAssignment**](IpSpaceOrgAssignment.md)|  | 
  **assignmentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

