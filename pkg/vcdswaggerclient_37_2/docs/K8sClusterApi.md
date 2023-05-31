# \K8sClusterApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateK8sCluster**](K8sClusterApi.md#CreateK8sCluster) | **Post** /1.0.0/tkgClusters | Creates a new Kubernetes cluster. This operation is asynchronous and returns a task that you can monitor to track the progress of the request. 
[**DeleteK8sCluster**](K8sClusterApi.md#DeleteK8sCluster) | **Delete** /1.0.0/tkgClusters/{urn} | 
[**GenerateKubeconfig**](K8sClusterApi.md#GenerateKubeconfig) | **Post** /1.0.0/tkgClusters/{urn}/kubeconfig | Generate kubeconfig file for corresponding cluster
[**GetK8sCluster**](K8sClusterApi.md#GetK8sCluster) | **Get** /1.0.0/tkgClusters/{urn} | Get specified Kubernetes Cluster
[**QueryK8sClusters**](K8sClusterApi.md#QueryK8sClusters) | **Get** /1.0.0/tkgClusters | Retrieves all K8s clusters
[**UpdateK8sCluster**](K8sClusterApi.md#UpdateK8sCluster) | **Put** /1.0.0/tkgClusters/{urn} | Update the desired state of the Kubernetes cluster. This operation is asynchronous and returns a task that you can monitor to track the progress of the request. 


# **CreateK8sCluster**
> CreateK8sCluster(ctx, k8sCluster)
Creates a new Kubernetes cluster. This operation is asynchronous and returns a task that you can monitor to track the progress of the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sCluster** | [**Cluster**](Cluster.md)| &lt;strong&gt;Example Value&lt;/strong&gt; for CSE Native cluster &lt;pre&gt; {   \&quot;metadata\&quot;: {     \&quot;name\&quot;: \&quot;my-CSE-Native-cluster\&quot;,     \&quot;site\&quot;: \&quot;\&quot;,     \&quot;orgName\&quot;: \&quot;my-organization\&quot;,     \&quot;virtualDataCenterName\&quot;: \&quot;my-org-vdc\&quot;   },   \&quot;apiVersion\&quot;: \&quot;cse.vmware.com/v2.0\&quot;,   \&quot;kind\&quot;: \&quot;native\&quot;,   \&quot;spec\&quot;: {     \&quot;expose\&quot;: false,     \&quot;settings\&quot;: {       \&quot;sshKey\&quot;: null,       \&quot;network\&quot;: null,       \&quot;ovdcNetwork\&quot;: \&quot;network-in-my-org-vdc\&quot;,       \&quot;rollbackOnFailure\&quot;: true     },     \&quot;topology\&quot;: {       \&quot;workers\&quot;: {         \&quot;count\&quot;: 1       },       \&quot;controlPlane\&quot;: {         \&quot;count\&quot;: 1       }     },     \&quot;distribution\&quot;: {       \&quot;templateName\&quot;: \&quot;ubuntu-16.04_k8-1.18_weave-2.6.5\&quot;,       \&quot;templateRevision\&quot;: 2     }   } } &lt;/pre&gt; &lt;strong&gt;Example Value&lt;/strong&gt; for TKGS cluster &lt;pre&gt; {   \&quot;kind\&quot;: \&quot;TanzuKubernetesCluster\&quot;,   \&quot;metadata\&quot;: {     \&quot;name\&quot;: \&quot;my-TKGS-cluster\&quot;,     \&quot;placementPolicy\&quot;: \&quot;my-placement-policy\&quot;,     \&quot;virtualDataCenterName\&quot;: \&quot;my-org-vdc\&quot;   },   \&quot;spec\&quot;: {     \&quot;distribution\&quot;: {       \&quot;version\&quot;: \&quot;v1.20.2\&quot;     },     \&quot;topology\&quot;: {       \&quot;controlPlane\&quot;: {         \&quot;class\&quot;: \&quot;best-effort-xsmall\&quot;,         \&quot;count\&quot;: 1,         \&quot;storageClass\&quot;: \&quot;my-storage-class\&quot;       },       \&quot;workers\&quot;: {         \&quot;class\&quot;: \&quot;best-effort-xsmall\&quot;,         \&quot;count\&quot;: 1,         \&quot;storageClass\&quot;: \&quot;my-storage-class\&quot;       }     }   } } &lt;/pre&gt;  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sCluster**
> DeleteK8sCluster(ctx, urn)


Deletes the Kubernetes cluster with the unique identifier (URN). This operation is asynchronous and returns a task that you can monitor to track the progress of the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urn** | **string**| A URN corresponding to a Kubernetes cluster previously created by a POST to OpenAPI tkgClusters endpoint.  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateKubeconfig**
> GenerateKubeconfig(ctx, urn)
Generate kubeconfig file for corresponding cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urn** | **string**| A URN corresponding to a Kubernetes cluster previously created by a POST to OpenAPI tkgClusters endpoint.  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetK8sCluster**
> Cluster GetK8sCluster(ctx, urn)
Get specified Kubernetes Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urn** | **string**| A URN corresponding to a Kubernetes cluster previously created by a POST to OpenAPI tkgClusters endpoint.  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryK8sClusters**
> Clusters QueryK8sClusters(ctx, page, pageSize, optional)
Retrieves all K8s clusters

Retrieves all K8s clusters 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***K8sClusterApiQueryK8sClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sClusterApiQueryK8sClustersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Clusters**](Clusters.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sCluster**
> UpdateK8sCluster(ctx, k8sCluster, urn)
Update the desired state of the Kubernetes cluster. This operation is asynchronous and returns a task that you can monitor to track the progress of the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sCluster** | [**Cluster**](Cluster.md)|  | 
  **urn** | **string**| A URN corresponding to a Kubernetes cluster previously created by a POST to OpenAPI tkgClusters endpoint.  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

