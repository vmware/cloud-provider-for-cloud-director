# \VroWorkflowExecutionApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelWorkflowExecution**](VroWorkflowExecutionApi.md#CancelWorkflowExecution) | **Delete** /workflows/{workflowId}/executions/{executionId}/state | Cancel workflow run
[**DeleteWorkflowExecution**](VroWorkflowExecutionApi.md#DeleteWorkflowExecution) | **Delete** /workflows/{workflowId}/executions/{executionId} | Delete workflow run
[**GetAllWorkflowExecutions**](VroWorkflowExecutionApi.md#GetAllWorkflowExecutions) | **Get** /workflows/{workflowId}/executions | Get all executions
[**GetParameterTypes**](VroWorkflowExecutionApi.md#GetParameterTypes) | **Get** /workflows/parameterTypes | This endpoint will not produce results. It is a placeholder to enforce code generation of ParameterTypes.
[**GetSupportedDecorators**](VroWorkflowExecutionApi.md#GetSupportedDecorators) | **Get** /workflows/decorators | This endpoint will not produce results. It is a placeholder to enforce code generation of SupportedDecorators.
[**GetSupportedPresentationElements**](VroWorkflowExecutionApi.md#GetSupportedPresentationElements) | **Get** /workflows/presentationElements | This endpoint will not produce results. It is a placeholder to enforce code generation of SupportedPresentationElements.
[**GetSupportedconstraints**](VroWorkflowExecutionApi.md#GetSupportedconstraints) | **Get** /workflows/constraints | This endpoint will not produce results. It is a placeholder to enforce code generation of SupportedConstraints.
[**GetWorkflowExecution**](VroWorkflowExecutionApi.md#GetWorkflowExecution) | **Get** /workflows/{workflowId}/executions/{executionId} | Get workflow execution
[**GetWorkflowExecutionState**](VroWorkflowExecutionApi.md#GetWorkflowExecutionState) | **Get** /workflows/{workflowId}/executions/{executionId}/state | Get workflow execution state
[**StartWorkflowExecution**](VroWorkflowExecutionApi.md#StartWorkflowExecution) | **Post** /workflows/{workflowId}/executions | Start workflow execution


# **CancelWorkflowExecution**
> CancelWorkflowExecution(ctx, workflowId, executionId)
Cancel workflow run

Cancels a workflow run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 
  **executionId** | **string**| executionId | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWorkflowExecution**
> DeleteWorkflowExecution(ctx, workflowId, executionId)
Delete workflow run

Deletes a specific workflow run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 
  **executionId** | **string**| executionId | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllWorkflowExecutions**
> PresentationExecutionsList GetAllWorkflowExecutions(ctx, workflowId)
Get all executions

Retrieves all workflow runs for a given workflow

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
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParameterTypes**
> ParameterTypes GetParameterTypes(ctx, )
This endpoint will not produce results. It is a placeholder to enforce code generation of ParameterTypes.

This endpoint will not produce results. It is a placeholder to enforce code generation of ParameterTypes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ParameterTypes**](ParameterTypes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedDecorators**
> SupportedDecorators GetSupportedDecorators(ctx, )
This endpoint will not produce results. It is a placeholder to enforce code generation of SupportedDecorators.

This endpoint will not produce results. It is a placeholder to enforce code generation of SupportedDecorators.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SupportedDecorators**](SupportedDecorators.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedPresentationElements**
> SupportedPresentationElements GetSupportedPresentationElements(ctx, )
This endpoint will not produce results. It is a placeholder to enforce code generation of SupportedPresentationElements.

This endpoint will not produce results. It is a placeholder to enforce code generation of SupportedPresentationElements.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SupportedPresentationElements**](SupportedPresentationElements.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedconstraints**
> SupportedConstraints GetSupportedconstraints(ctx, )
This endpoint will not produce results. It is a placeholder to enforce code generation of SupportedConstraints.

This endpoint will not produce results. It is a placeholder to enforce code generation of SupportedConstraints.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SupportedConstraints**](SupportedConstraints.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflowExecution**
> WsWorkflowExecution GetWorkflowExecution(ctx, workflowId, executionId)
Get workflow execution

Retrieves a workflow run for a given workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 
  **executionId** | **string**| executionId | 

### Return type

[**WsWorkflowExecution**](WsWorkflowExecution.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflowExecutionState**
> ExecutionState GetWorkflowExecutionState(ctx, workflowId, executionId)
Get workflow execution state

Retrieves the current state for a requested workflow run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 
  **executionId** | **string**| executionId | 

### Return type

[**ExecutionState**](ExecutionState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartWorkflowExecution**
> WsWorkflowExecution StartWorkflowExecution(ctx, workflowId, executionContext)
Start workflow execution

Instantiates a workflow run, by using the passed parameters. The workflow run is asynchronous, so the call returns a pointer to a task that can be used to track the workflow run. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| workflowId | 
  **executionContext** | [**ExecutionContext**](ExecutionContext.md)| executionContext | 

### Return type

[**WsWorkflowExecution**](WsWorkflowExecution.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

