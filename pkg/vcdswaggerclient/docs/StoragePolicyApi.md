# \StoragePolicyApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStoragePolicies**](StoragePolicyApi.md#GetStoragePolicies) | **Get** /1.0.0/storagePolicies | Get a paged list of all Provider Storage Policies in the system
[**GetStoragePolicy**](StoragePolicyApi.md#GetStoragePolicy) | **Get** /1.0.0/storagePolicies/{storagePolicyUrn} | Get specified Storage Policy.


# **GetStoragePolicies**
> StoragePolicies GetStoragePolicies(ctx, page, pageSize, optional)
Get a paged list of all Provider Storage Policies in the system

Get a paged list of all Provider Storage Policies in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***StoragePolicyApiGetStoragePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StoragePolicyApiGetStoragePoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**StoragePolicies**](StoragePolicies.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoragePolicy**
> StoragePolicy GetStoragePolicy(ctx, storagePolicyUrn)
Get specified Storage Policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storagePolicyUrn** | **string**| storagePolicyUrn | 

### Return type

[**StoragePolicy**](StoragePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

