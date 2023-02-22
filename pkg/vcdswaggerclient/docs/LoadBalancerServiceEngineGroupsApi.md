# \LoadBalancerServiceEngineGroupsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceEngineGroup**](LoadBalancerServiceEngineGroupsApi.md#CreateServiceEngineGroup) | **Post** /1.0.0/loadBalancer/serviceEngineGroups | Create a new Load Balancer Service Engine Group.
[**GetServiceEngineGroups**](LoadBalancerServiceEngineGroupsApi.md#GetServiceEngineGroups) | **Get** /1.0.0/loadBalancer/serviceEngineGroups | Get all Load Balancer Service Engine Groups in the system.


# **CreateServiceEngineGroup**
> CreateServiceEngineGroup(ctx, loadBalancerServiceEngineGroup)
Create a new Load Balancer Service Engine Group.

Create a new Load Balancer Service Engine Group to be used with VMware Cloud Director. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerServiceEngineGroup** | [**LoadBalancerServiceEngineGroup**](LoadBalancerServiceEngineGroup.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceEngineGroups**
> LoadBalancerServiceEngineGroups GetServiceEngineGroups(ctx, page, pageSize, optional)
Get all Load Balancer Service Engine Groups in the system.

Retrieves all Load Balancer Service Engine Groups. Supported contexts are: Gateway ID <code>(_context==gatewayId)</code> - | Returns all Load Balancer Service Engine Groups that are accessible to the gateway. Assignable Gateway ID <code>(_context=gatewayId;_context==assignable)</code> - | Returns all Load Balancer Service Engine Groups that are assignable to the gateway. This filters out any Load Balancer Service Engine groups that are already assigned to the gateway or assigned to another gateway if the reservation type is 'DEDICATED'. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***LoadBalancerServiceEngineGroupsApiGetServiceEngineGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerServiceEngineGroupsApiGetServiceEngineGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**LoadBalancerServiceEngineGroups**](LoadBalancerServiceEngineGroups.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

