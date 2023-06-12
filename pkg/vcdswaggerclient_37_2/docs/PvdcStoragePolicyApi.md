# \PvdcStoragePolicyApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSupportedStorageEntityTypes**](PvdcStoragePolicyApi.md#GetAllSupportedStorageEntityTypes) | **Get** /1.0.0/storagePolicySupportedEntityTypes | Get a paged list of all supported entity types configured for storage policies in the system
[**GetPvdcStoragePolicies**](PvdcStoragePolicyApi.md#GetPvdcStoragePolicies) | **Get** /1.0.0/pvdcStoragePolicies | Get a paged list of all Provider VDC level storage policies in the system
[**GetPvdcStoragePolicy**](PvdcStoragePolicyApi.md#GetPvdcStoragePolicy) | **Get** /1.0.0/pvdcStoragePolicies/{pvdcStoragePolicyUrn} | Get specified Provider VDC storage policy.
[**GetPvdcStoragePolicyInheritableSettings**](PvdcStoragePolicyApi.md#GetPvdcStoragePolicyInheritableSettings) | **Get** /1.0.0/pvdcStoragePolicies/{id}/inheritableSettings | Retrieves the settings that child Org VDC storage policies of this provider VDC storage policy should inherit. 
[**GetPvdcStoragePolicySupportedEntityTypes**](PvdcStoragePolicyApi.md#GetPvdcStoragePolicySupportedEntityTypes) | **Get** /1.0.0/pvdcStoragePolicies/{pvdcStoragePolicyUrn}/supportedEntityTypes | Get a paged list of the supported entity types for the specified Provider VDC storage policy.
[**UpdateAllStoragePolicySupportedEntityTypes**](PvdcStoragePolicyApi.md#UpdateAllStoragePolicySupportedEntityTypes) | **Put** /1.0.0/storagePolicySupportedEntityTypes | Updates the supported entity types for the specified provider VDC storage policy. 
[**UpdatePvdcStoragePolicyInheritableSettings**](PvdcStoragePolicyApi.md#UpdatePvdcStoragePolicyInheritableSettings) | **Put** /1.0.0/pvdcStoragePolicies/{id}/inheritableSettings | Updates the settings that child Org VDC storage policies of this provider VDC storage policy should inherit. 
[**UpdatePvdcStoragePolicySupportedEntityTypes**](PvdcStoragePolicyApi.md#UpdatePvdcStoragePolicySupportedEntityTypes) | **Put** /1.0.0/pvdcStoragePolicies/{pvdcStoragePolicyUrn}/supportedEntityTypes | Updates the supported entity types for the specified provider VDC storage policy. 


# **GetAllSupportedStorageEntityTypes**
> StoragePolicySupportedEntityTypes GetAllSupportedStorageEntityTypes(ctx, page, pageSize, optional)
Get a paged list of all supported entity types configured for storage policies in the system

Get a paged list of all supported entity types configured for storage policies in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***PvdcStoragePolicyApiGetAllSupportedStorageEntityTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PvdcStoragePolicyApiGetAllSupportedStorageEntityTypesOpts struct

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPvdcStoragePolicies**
> PvdcStoragePolicies GetPvdcStoragePolicies(ctx, page, pageSize, optional)
Get a paged list of all Provider VDC level storage policies in the system

Get a paged list of all Provider VDC level storage policies in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***PvdcStoragePolicyApiGetPvdcStoragePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PvdcStoragePolicyApiGetPvdcStoragePoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**PvdcStoragePolicies**](PvdcStoragePolicies.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPvdcStoragePolicy**
> PvdcStoragePolicy GetPvdcStoragePolicy(ctx, pvdcStoragePolicyUrn)
Get specified Provider VDC storage policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pvdcStoragePolicyUrn** | **string**| pvdcStoragePolicyUrn | 

### Return type

[**PvdcStoragePolicy**](PvdcStoragePolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPvdcStoragePolicyInheritableSettings**
> StoragePolicySettings GetPvdcStoragePolicyInheritableSettings(ctx, id)
Retrieves the settings that child Org VDC storage policies of this provider VDC storage policy should inherit. 

Retrieves the settings that child Org VDC storage policies of this provider VDC storage policy should inherit. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**StoragePolicySettings**](StoragePolicySettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPvdcStoragePolicySupportedEntityTypes**
> StoragePolicySupportedEntityTypes GetPvdcStoragePolicySupportedEntityTypes(ctx, page, pageSize, pvdcStoragePolicyUrn, optional)
Get a paged list of the supported entity types for the specified Provider VDC storage policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **pvdcStoragePolicyUrn** | **string**| pvdcStoragePolicyUrn | 
 **optional** | ***PvdcStoragePolicyApiGetPvdcStoragePolicySupportedEntityTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PvdcStoragePolicyApiGetPvdcStoragePolicySupportedEntityTypesOpts struct

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAllStoragePolicySupportedEntityTypes**
> StoragePolicySupportedEntityTypes UpdateAllStoragePolicySupportedEntityTypes(ctx, updatedSupportedEntityTypes)
Updates the supported entity types for the specified provider VDC storage policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedSupportedEntityTypes** | [**[]StoragePolicySupportedEntityType**](StoragePolicySupportedEntityType.md)| The updated list of supported entity types. | 

### Return type

[**StoragePolicySupportedEntityTypes**](StoragePolicySupportedEntityTypes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePvdcStoragePolicyInheritableSettings**
> StoragePolicySettings UpdatePvdcStoragePolicyInheritableSettings(ctx, updatedSettings, id)
Updates the settings that child Org VDC storage policies of this provider VDC storage policy should inherit. 

Updates the settings that child Org VDC storage policies of this provider VDC storage policy should inherit. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedSettings** | [**StoragePolicySettings**](StoragePolicySettings.md)| The updated inheritable settings. | 
  **id** | **string**|  | 

### Return type

[**StoragePolicySettings**](StoragePolicySettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePvdcStoragePolicySupportedEntityTypes**
> StoragePolicySupportedEntityTypes UpdatePvdcStoragePolicySupportedEntityTypes(ctx, updatedSupportedEntityTypes, pvdcStoragePolicyUrn)
Updates the supported entity types for the specified provider VDC storage policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedSupportedEntityTypes** | [**[]StoragePolicySupportedEntityType**](StoragePolicySupportedEntityType.md)| The updated list of supported entity types. | 
  **pvdcStoragePolicyUrn** | **string**| pvdcStoragePolicyUrn | 

### Return type

[**StoragePolicySupportedEntityTypes**](StoragePolicySupportedEntityTypes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

