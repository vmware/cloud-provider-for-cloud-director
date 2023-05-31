# \LoadBalancerServiceEngineGroupAssignmentsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceEngineGroupAssignment**](LoadBalancerServiceEngineGroupAssignmentsApi.md#CreateServiceEngineGroupAssignment) | **Post** /1.0.0/loadBalancer/serviceEngineGroups/assignments | Create a new Load Balancer Service Engine Group Assignment.
[**GetServiceEngineGroupAssignments**](LoadBalancerServiceEngineGroupAssignmentsApi.md#GetServiceEngineGroupAssignments) | **Get** /1.0.0/loadBalancer/serviceEngineGroups/assignments | Get the assignments for a Load Balancer Service Engine Group.


# **CreateServiceEngineGroupAssignment**
> CreateServiceEngineGroupAssignment(ctx, assignment)
Create a new Load Balancer Service Engine Group Assignment.

Create a new Load Balancer Service Engine Group Assignment. The assignment links a Load Balancer Service Engine Group with an Edge Gateway to provide load balancing resources to the Edge Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assignment** | [**LoadBalancerServiceEngineGroupAssignment**](LoadBalancerServiceEngineGroupAssignment.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceEngineGroupAssignments**
> LoadBalancerServiceEngineGroupAssignments GetServiceEngineGroupAssignments(ctx, page, pageSize, optional)
Get the assignments for a Load Balancer Service Engine Group.

Retrieves the service engine group assignments for the Load Balancer Service Engine Group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***LoadBalancerServiceEngineGroupAssignmentsApiGetServiceEngineGroupAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerServiceEngineGroupAssignmentsApiGetServiceEngineGroupAssignmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**LoadBalancerServiceEngineGroupAssignments**](LoadBalancerServiceEngineGroupAssignments.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

