# \StorageContainerApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatastoreClusterDatastores**](StorageContainerApi.md#GetDatastoreClusterDatastores) | **Get** /1.0.0/storageContainers/{storageContainerUrn}/datastores | Retrieves all Datastores associated with the specified Datastore Cluster.
[**GetStorageContainer**](StorageContainerApi.md#GetStorageContainer) | **Get** /1.0.0/storageContainers/{storageContainerUrn} | Get specified Datastore or Datastore Cluster.
[**GetStorageContainers**](StorageContainerApi.md#GetStorageContainers) | **Get** /1.0.0/storageContainers | Get a paged list of all standalone Datastores and Datastore Clusters in the system


# **GetDatastoreClusterDatastores**
> StorageContainers GetDatastoreClusterDatastores(ctx, page, pageSize, storageContainerUrn, optional)
Retrieves all Datastores associated with the specified Datastore Cluster.

Retrieves all Datastores associated with the specified Datastore Cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **storageContainerUrn** | **string**|  | 
 **optional** | ***StorageContainerApiGetDatastoreClusterDatastoresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageContainerApiGetDatastoreClusterDatastoresOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**StorageContainers**](StorageContainers.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStorageContainer**
> StorageContainer GetStorageContainer(ctx, storageContainerUrn)
Get specified Datastore or Datastore Cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storageContainerUrn** | **string**| storageContainerUrn | 

### Return type

[**StorageContainer**](StorageContainer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStorageContainers**
> StorageContainers GetStorageContainers(ctx, page, pageSize, optional)
Get a paged list of all standalone Datastores and Datastore Clusters in the system

Get a paged list of all standalone Datastores and Datastore Clusters in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***StorageContainerApiGetStorageContainersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageContainerApiGetStorageContainersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**StorageContainers**](StorageContainers.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

