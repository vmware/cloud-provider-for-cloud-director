# \LoadBalancerServiceEngineGroupAssignmentApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceEngineGroupAssignment**](LoadBalancerServiceEngineGroupAssignmentApi.md#DeleteServiceEngineGroupAssignment) | **Delete** /1.0.0/loadBalancer/serviceEngineGroups/assignments/{assignmentId} | Delete the specified Load Balancer Service Engine Group Assignment.
[**GetServiceEngineGroupAssignment**](LoadBalancerServiceEngineGroupAssignmentApi.md#GetServiceEngineGroupAssignment) | **Get** /1.0.0/loadBalancer/serviceEngineGroups/assignments/{assignmentId} | Get a Load Balancer Service Engine Group Assignment.
[**UpdateServiceEngineGroupAssignment**](LoadBalancerServiceEngineGroupAssignmentApi.md#UpdateServiceEngineGroupAssignment) | **Put** /1.0.0/loadBalancer/serviceEngineGroups/assignments/{assignmentId} | Update a Load Balancer Service Engine Group Assignment.


# **DeleteServiceEngineGroupAssignment**
> DeleteServiceEngineGroupAssignment(ctx, assignmentId)
Delete the specified Load Balancer Service Engine Group Assignment.

Delete a Load Balancer Service Engine Group Assignment. The Edge Gateway will no longer be able to use the Load Balancer Service Engine Group for load balancing resources. 

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

# **GetServiceEngineGroupAssignment**
> LoadBalancerServiceEngineGroupAssignment GetServiceEngineGroupAssignment(ctx, assignmentId)
Get a Load Balancer Service Engine Group Assignment.

Retrieves a specific Load Balancer Service Engine Group Assignment. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assignmentId** | **string**|  | 

### Return type

[**LoadBalancerServiceEngineGroupAssignment**](LoadBalancerServiceEngineGroupAssignment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceEngineGroupAssignment**
> UpdateServiceEngineGroupAssignment(ctx, assignment, assignmentId)
Update a Load Balancer Service Engine Group Assignment.

Update a Load Balancer Service Engine Group Assignment. Updates are not allowed if the associated Load Balancer Service Engine Group has reservation type 'DEDICATED'. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assignment** | [**LoadBalancerServiceEngineGroupAssignment**](LoadBalancerServiceEngineGroupAssignment.md)|  | 
  **assignmentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

