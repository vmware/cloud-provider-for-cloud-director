# \VrealizeWorkflowsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRemoteWorkflows**](VrealizeWorkflowsApi.md#GetRemoteWorkflows) | **Get** /vro/servers/{vroServerId}/workflows | Browse remote vRealize Orchestrator workflows


# **GetRemoteWorkflows**
> VroRemoteWorkflowItems GetRemoteWorkflows(ctx, page, pageSize, vroServerId, optional)
Browse remote vRealize Orchestrator workflows

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **vroServerId** | **string**| The ID of the server to browse workflows on | 
 **optional** | ***VrealizeWorkflowsApiGetRemoteWorkflowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VrealizeWorkflowsApiGetRemoteWorkflowsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VroRemoteWorkflowItems**](VroRemoteWorkflowItems.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

