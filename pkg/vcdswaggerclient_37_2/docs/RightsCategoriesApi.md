# \RightsCategoriesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRightsCategory**](RightsCategoriesApi.md#GetRightsCategory) | **Get** /1.0.0/rightsCategories/{id} | Retrieve an individual Right category.
[**QueryRightsCategories**](RightsCategoriesApi.md#QueryRightsCategories) | **Get** /1.0.0/rightsCategories | Get a list of Rights Categories visible to the logged in user


# **GetRightsCategory**
> RightsCategoryNode GetRightsCategory(ctx, id)
Retrieve an individual Right category.

Retrieves the requested Rights Category by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**RightsCategoryNode**](RightsCategoryNode.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryRightsCategories**
> RightsCategoryNodes QueryRightsCategories(ctx, page, pageSize, optional)
Get a list of Rights Categories visible to the logged in user

Get list of Rights Categories 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***RightsCategoriesApiQueryRightsCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RightsCategoriesApiQueryRightsCategoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**RightsCategoryNodes**](RightsCategoryNodes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

