# \OrgVdcStoragePolicyApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgVdcStoragePolicies**](OrgVdcStoragePolicyApi.md#GetOrgVdcStoragePolicies) | **Get** /1.0.0/orgVdcStoragePolicies | Get a paged list of all organization VDC level storage policies in the system
[**GetOrgVdcStoragePolicy**](OrgVdcStoragePolicyApi.md#GetOrgVdcStoragePolicy) | **Get** /1.0.0/orgVdcStoragePolicies/{orgVdcStoragePolicyUrn} | Get specified Org VDC storage policy.
[**GetOrgVdcStoragePolicySupportedEntityTypes**](OrgVdcStoragePolicyApi.md#GetOrgVdcStoragePolicySupportedEntityTypes) | **Get** /1.0.0/orgVdcStoragePolicies/{orgVdcStoragePolicyUrn}/supportedEntityTypes | Get a paged list of the supported entity types for the specified Organization VDC storage policy.


# **GetOrgVdcStoragePolicies**
> OrgVdcStoragePolicies GetOrgVdcStoragePolicies(ctx, page, pageSize, optional)
Get a paged list of all organization VDC level storage policies in the system

Get a paged list of all organization VDC level storage policies in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***OrgVdcStoragePolicyApiGetOrgVdcStoragePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgVdcStoragePolicyApiGetOrgVdcStoragePoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**OrgVdcStoragePolicies**](OrgVdcStoragePolicies.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgVdcStoragePolicy**
> OrgVdcStoragePolicy GetOrgVdcStoragePolicy(ctx, orgVdcStoragePolicyUrn)
Get specified Org VDC storage policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgVdcStoragePolicyUrn** | **string**| orgVdcStoragePolicyUrn | 

### Return type

[**OrgVdcStoragePolicy**](OrgVdcStoragePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgVdcStoragePolicySupportedEntityTypes**
> StoragePolicySupportedEntityTypes GetOrgVdcStoragePolicySupportedEntityTypes(ctx, page, pageSize, orgVdcStoragePolicyUrn, optional)
Get a paged list of the supported entity types for the specified Organization VDC storage policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **orgVdcStoragePolicyUrn** | **string**| orgVdcStoragePolicyUrn | 
 **optional** | ***OrgVdcStoragePolicyApiGetOrgVdcStoragePolicySupportedEntityTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgVdcStoragePolicyApiGetOrgVdcStoragePolicySupportedEntityTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**StoragePolicySupportedEntityTypes**](StoragePolicySupportedEntityTypes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

