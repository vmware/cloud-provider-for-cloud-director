# \VroWorkflowPresentationApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWorkflowPresenationInstance**](VroWorkflowPresentationApi.md#DeleteWorkflowPresenationInstance) | **Delete** /workflows/{workflowId}/presentation/instances/{presentationExecutionId} | Delete workflow presentation execution
[**GetAllWorkflowPresentationInstances**](VroWorkflowPresentationApi.md#GetAllWorkflowPresentationInstances) | **Get** /workflows/{workflowId}/presentation/instances | Get all presentations
[**GetWorkflowPresentation**](VroWorkflowPresentationApi.md#GetWorkflowPresentation) | **Get** /workflows/{workflowId}/presentation | Get presentation
[**GetWorkflowPresentationInstance**](VroWorkflowPresentationApi.md#GetWorkflowPresentationInstance) | **Get** /workflows/{workflowId}/presentation/instances/{presentationExecutionId} | Load Execution
[**StartWorkflowPresentation**](VroWorkflowPresentationApi.md#StartWorkflowPresentation) | **Post** /workflows/{workflowId}/presentation/instances | Start presentation
[**UpdateWorkflowPresentationInstance**](VroWorkflowPresentationApi.md#UpdateWorkflowPresentationInstance) | **Put** /workflows/{workflowId}/presentation/instances/{presentationExecutionId} | Update presentation


# **DeleteWorkflowPresenationInstance**
> DeleteWorkflowPresenationInstance(ctx, workflowId, presentationExecutionId)
Delete workflow presentation execution

Cancels the execution of a workflow presentation instance. This API call cancels only the workflow presentation execution. To cancel the workflow execution, use APIs under <b>/cloudapi/workflows/{workflowId}/instances</b>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 
  **presentationExecutionId** | **string**| presentationExecutionId | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllWorkflowPresentationInstances**
> PresentationExecutionsList GetAllWorkflowPresentationInstances(ctx, workflowId)
Get all presentations

Retrieves a list of the presentation instances for a workflow that you specify. To retrieve the list of workflow presentations, make an HTTP GET request at the workflow presentations list URL. The returned list contains all of the currently running workflow presentation instances, and all completed instances based on the data from the workflow executions. If the user has admin rights, all presentation instances for all users are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 

### Return type

[**PresentationExecutionsList**](PresentationExecutionsList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflowPresentation**
> Presentation GetWorkflowPresentation(ctx, workflowId)
Get presentation

Retrieves the definition of a workflow presentation. To retrieve the workflow presentation definition localized, add Accept-Language header, with the appropriate locale. In advance, localization resource should be present for the workflow, otherwise it defaults to the standard workflow presentation definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 

### Return type

[**Presentation**](Presentation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflowPresentationInstance**
> PresentationExecution GetWorkflowPresentationInstance(ctx, workflowId, presentationExecutionId)
Load Execution

Retrieves a specific workflow presentation instance. Presentation instances are removed after the workflow starts. If the presentation instance under requested <b>executionId</b> does not exists, a new presentation instance is created by using the parameters from the workflow execution with the same ID. To retrieve the workflow presentation localized, add Accept-Language header, with the appropriate locale. In advance, localization resource should be present for the workflow, otherwise it defaults to the standard workflow presentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 
  **presentationExecutionId** | **string**| presentationExecutionId | 

### Return type

[**PresentationExecution**](PresentationExecution.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartWorkflowPresentation**
> PresentationExecution StartWorkflowPresentation(ctx, workflowId, executionContext)
Start presentation

 Creates a new instance of the presentation of a workflow, by using the passed parameters. To create a new instance of a workflow presentation, make an HTTP GET request at the URL that contains the instances of the workflow presentation. Presentation's fields are populated with input parameter values and are validated. If there are any validation errors, they are collected and attached to each field. The presentation is marked as invalid. In order the returned workflow presentation to be localized, add Accept-Language header, with the appropriate locale. In advance, localization resource should be present for the workflow, otherwise it defaults to the standard workflow presentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 
  **executionContext** | [**ExecutionContext**](ExecutionContext.md)| executionContext | 

### Return type

[**PresentationExecution**](PresentationExecution.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWorkflowPresentationInstance**
> PresentationExecution UpdateWorkflowPresentationInstance(ctx, workflowId, presentationExecutionId, executionContext)
Update presentation

Update a specific workflow presentation instance. Presentation fields are populated with input parameter values and are validated. If there are any validation errors, they are collected and attached to each field. The presentation is marked as invalid. If the parameter's 'updated' flag is set to true, the dependent field values are recalculated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 
  **presentationExecutionId** | **string**| presentationExecutionId | 
  **executionContext** | [**ExecutionContext**](ExecutionContext.md)| executionContext | 

### Return type

[**PresentationExecution**](PresentationExecution.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

