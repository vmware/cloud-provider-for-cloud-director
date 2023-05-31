# \EdgeClusterApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEdgeCluster**](EdgeClusterApi.md#DeleteEdgeCluster) | **Delete** /1.0.0/edgeClusters/{edgeClusterId} | Deletes a specific Edge Cluster
[**GetEdgeCluster**](EdgeClusterApi.md#GetEdgeCluster) | **Get** /1.0.0/edgeClusters/{edgeClusterId} | Retrieves a specific Edge Cluster
[**UpdateEdgeCluster**](EdgeClusterApi.md#UpdateEdgeCluster) | **Put** /1.0.0/edgeClusters/{edgeClusterId} | Updates a specific Edge Cluster


# **DeleteEdgeCluster**
> DeleteEdgeCluster(ctx, edgeClusterId)
Deletes a specific Edge Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeClusterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeCluster**
> EdgeCluster GetEdgeCluster(ctx, edgeClusterId)
Retrieves a specific Edge Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeClusterId** | **string**|  | 

### Return type

[**EdgeCluster**](EdgeCluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEdgeCluster**
> UpdateEdgeCluster(ctx, edgeCluster, edgeClusterId)
Updates a specific Edge Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeCluster** | [**EdgeCluster**](EdgeCluster.md)|  | 
  **edgeClusterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

