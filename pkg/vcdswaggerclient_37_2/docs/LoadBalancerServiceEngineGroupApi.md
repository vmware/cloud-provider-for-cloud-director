# \LoadBalancerServiceEngineGroupApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceEngineGroup**](LoadBalancerServiceEngineGroupApi.md#DeleteServiceEngineGroup) | **Delete** /1.0.0/loadBalancer/serviceEngineGroups/{loadBalancerServiceEngineGroupId} | Delete the specified Load Balancer Service Engine Group.
[**GetServiceEngineGroup**](LoadBalancerServiceEngineGroupApi.md#GetServiceEngineGroup) | **Get** /1.0.0/loadBalancer/serviceEngineGroups/{loadBalancerServiceEngineGroupId} | Get Load Balancer Service Engine Group.
[**SyncServiceEngineGroup**](LoadBalancerServiceEngineGroupApi.md#SyncServiceEngineGroup) | **Post** /1.0.0/loadBalancer/serviceEngineGroups/{loadBalancerServiceEngineGroupId}/sync | Sync Load Balancer Service Engine Group.
[**UpdateServiceEngineGroup**](LoadBalancerServiceEngineGroupApi.md#UpdateServiceEngineGroup) | **Put** /1.0.0/loadBalancer/serviceEngineGroups/{loadBalancerServiceEngineGroupId} | Update specified Load Balancer Service Engine Group.


# **DeleteServiceEngineGroup**
> DeleteServiceEngineGroup(ctx, loadBalancerServiceEngineGroupId)
Delete the specified Load Balancer Service Engine Group.

Delete a Load Balancer Service Engine Group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerServiceEngineGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceEngineGroup**
> LoadBalancerServiceEngineGroup GetServiceEngineGroup(ctx, loadBalancerServiceEngineGroupId)
Get Load Balancer Service Engine Group.

Retrieves a specific Load Balancer Service Engine Group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerServiceEngineGroupId** | **string**|  | 

### Return type

[**LoadBalancerServiceEngineGroup**](LoadBalancerServiceEngineGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncServiceEngineGroup**
> SyncServiceEngineGroup(ctx, loadBalancerServiceEngineGroupId)
Sync Load Balancer Service Engine Group.

Syncs a specified Load Balancer Service Engine Group. Requests the HA mode and the maximum number of supported Virtual Services for this Service Engine Group from the Load Balancer, and updates vCD's local record of these properties. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerServiceEngineGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceEngineGroup**
> UpdateServiceEngineGroup(ctx, loadBalancerServiceEngineGroup, loadBalancerServiceEngineGroupId)
Update specified Load Balancer Service Engine Group.

Update a Load Balancer Service Engine Group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerServiceEngineGroup** | [**LoadBalancerServiceEngineGroup**](LoadBalancerServiceEngineGroup.md)|  | 
  **loadBalancerServiceEngineGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

