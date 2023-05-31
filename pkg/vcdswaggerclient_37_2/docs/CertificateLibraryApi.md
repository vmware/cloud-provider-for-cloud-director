# \CertificateLibraryApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCertificateLibraryItem**](CertificateLibraryApi.md#AddCertificateLibraryItem) | **Post** /1.0.0/ssl/certificateLibrary | Add an item to the certificate library
[**AddConsumerRefToCertLibraryItem**](CertificateLibraryApi.md#AddConsumerRefToCertLibraryItem) | **Post** /1.0.0/ssl/certificateLibrary/{certLibraryItemId}/consumers | Adds the specified consumer reference to a library item.
[**DeleteCertificateLibraryItem**](CertificateLibraryApi.md#DeleteCertificateLibraryItem) | **Delete** /1.0.0/ssl/certificateLibrary/{id} | Remove certificate library item.
[**GetCertificateLibraryItem**](CertificateLibraryApi.md#GetCertificateLibraryItem) | **Get** /1.0.0/ssl/certificateLibrary/{id} | Get the specified certificate library item
[**QueryCertLibraryItemConsumerRefs**](CertificateLibraryApi.md#QueryCertLibraryItemConsumerRefs) | **Get** /1.0.0/ssl/certificateLibrary/{certLibraryItemId}/consumers | Gets a paged list of consumers (as references) of a particular certificate library item
[**QueryCertificateLibrary**](CertificateLibraryApi.md#QueryCertificateLibrary) | **Get** /1.0.0/ssl/certificateLibrary | Get the certificate library items
[**ReplaceCertLibraryItemConsumerRefs**](CertificateLibraryApi.md#ReplaceCertLibraryItemConsumerRefs) | **Put** /1.0.0/ssl/certificateLibrary/{certLibraryItemId}/consumers | Replaces the existing consumer refs with the used by references supplied.
[**UpdateCertificateLibraryItem**](CertificateLibraryApi.md#UpdateCertificateLibraryItem) | **Put** /1.0.0/ssl/certificateLibrary/{id} | Update the specified certificate library item.


# **AddCertificateLibraryItem**
> CertificateLibraryItem AddCertificateLibraryItem(ctx, newCertificateLibraryItem)
Add an item to the certificate library

Add an item to the certificate library 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newCertificateLibraryItem** | [**CertificateLibraryItem**](CertificateLibraryItem.md)|  | 

### Return type

[**CertificateLibraryItem**](CertificateLibraryItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddConsumerRefToCertLibraryItem**
> EntityReference AddConsumerRefToCertLibraryItem(ctx, consumerReference, certLibraryItemId)
Adds the specified consumer reference to a library item.

Adds the specified consumer reference to a library item. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerReference** | [**EntityReference**](EntityReference.md)|  | 
  **certLibraryItemId** | **string**|  | 

### Return type

[**EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCertificateLibraryItem**
> DeleteCertificateLibraryItem(ctx, id)
Remove certificate library item.

Delete the specified certificate library item. Only items that are not in use can be deleted. Note: This API also supports a former (erroneously spelt) alternate path /cetificateLibrary/{id} as a Deprecated API (deprecated-in and removed after API version 36.0) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertificateLibraryItem**
> CertificateLibraryItem GetCertificateLibraryItem(ctx, id)
Get the specified certificate library item

Retrieves the specified certificate library item. Note: This API also supports a former (erroneously spelt) alternate path /cetificateLibrary/{id} as a Deprecated API (deprecated-in and removed after API version 36.0) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**CertificateLibraryItem**](CertificateLibraryItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryCertLibraryItemConsumerRefs**
> EntityReferences QueryCertLibraryItemConsumerRefs(ctx, page, pageSize, certLibraryItemId, optional)
Gets a paged list of consumers (as references) of a particular certificate library item

Get list of consumers (as references) of a particular certificate library item 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **certLibraryItemId** | **string**|  | 
 **optional** | ***CertificateLibraryApiQueryCertLibraryItemConsumerRefsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificateLibraryApiQueryCertLibraryItemConsumerRefsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryCertificateLibrary**
> CertificateLibraryItems QueryCertificateLibrary(ctx, page, pageSize, optional)
Get the certificate library items

Get a list of the certificate library items 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***CertificateLibraryApiQueryCertificateLibraryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificateLibraryApiQueryCertificateLibraryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**CertificateLibraryItems**](CertificateLibraryItems.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceCertLibraryItemConsumerRefs**
> EntityReferences ReplaceCertLibraryItemConsumerRefs(ctx, consumerRefs, certLibraryItemId)
Replaces the existing consumer refs with the used by references supplied.

Replaces the existing consumer refs with the consumer references supplied. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerRefs** | [**EntityReferences**](EntityReferences.md)|  | 
  **certLibraryItemId** | **string**|  | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCertificateLibraryItem**
> CertificateLibraryItem UpdateCertificateLibraryItem(ctx, modifiedCertificatLibraryItem, id)
Update the specified certificate library item.

Updates the specified certificate library item. Only the alias and description fields may be edited Note: This API also supports a former (erroneously spelt) alternate path /cetificateLibrary/{id} as a Deprecated API (deprecated-in and removed after API version 36.0) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modifiedCertificatLibraryItem** | [**CertificateLibraryItem**](CertificateLibraryItem.md)|  | 
  **id** | **string**|  | 

### Return type

[**CertificateLibraryItem**](CertificateLibraryItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

