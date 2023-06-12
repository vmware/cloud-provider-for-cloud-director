# \NsxAlbResourcesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImportableClouds**](NsxAlbResourcesApi.md#GetImportableClouds) | **Get** /1.0.0/nsxAlbResources/importableClouds | Get all NSX-ALB Clouds from an NSX-ALB Controller.
[**GetImportableServiceEngineGroups**](NsxAlbResourcesApi.md#GetImportableServiceEngineGroups) | **Get** /1.0.0/nsxAlbResources/importableServiceEngineGroups | Get all importable Service Engine Groups from an NSX-ALB Cloud.


# **GetImportableClouds**
> NsxAlbClouds GetImportableClouds(ctx, page, pageSize, optional)
Get all NSX-ALB Clouds from an NSX-ALB Controller.

Get all NSX-ALB Clouds that are configured on an NSX-ALB Controller.  Clouds that are already imported are marked appropriately. The \"_context\" filter key must be set with the id of the Load Balancer Controller for which we want to get the NSX-ALB Clouds for. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxAlbResourcesApiGetImportableCloudsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxAlbResourcesApiGetImportableCloudsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**NsxAlbClouds**](NsxAlbClouds.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportableServiceEngineGroups**
> NsxAlbServiceEngineGroups GetImportableServiceEngineGroups(ctx, pageSize, optional)
Get all importable Service Engine Groups from an NSX-ALB Cloud.

Get all importable Service Engine Groups that are configured for an NSX-ALB Cloud. Service Engine Groups that are already imported are filtered out. The \"_context\" filter key must be set with the id of the Load Balancer Cloud for which we want to get the NSX-ALB Service Engine Groups for. Example: (_context==urn:vcloud:loadBalancerCloud:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxAlbResourcesApiGetImportableServiceEngineGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxAlbResourcesApiGetImportableServiceEngineGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**NsxAlbServiceEngineGroups**](NsxAlbServiceEngineGroups.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

