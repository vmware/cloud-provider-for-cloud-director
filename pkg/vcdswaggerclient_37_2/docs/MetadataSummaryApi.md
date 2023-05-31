# \MetadataSummaryApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetadataSummaries**](MetadataSummaryApi.md#GetMetadataSummaries) | **Get** /1.0.0/metadataSummaries/{metadataCursor} | Retrieves a map of entity ids to metadata summaries.


# **GetMetadataSummaries**
> MetadataSummaries GetMetadataSummaries(ctx, metadataCursor)
Retrieves a map of entity ids to metadata summaries.

Retrieves a map of entity ids to metadata summaries. Each entry carries only core entry data. Only entries available to the current user will be presented. If the user does not have access to the main entity, it will not be present in the map. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **metadataCursor** | **string**| A hash which identifies a list of entity ids | 

### Return type

[**MetadataSummaries**](MetadataSummaries.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

