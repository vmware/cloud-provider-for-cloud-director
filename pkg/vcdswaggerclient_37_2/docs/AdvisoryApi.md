# \AdvisoryApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdvisoryDefinition**](AdvisoryApi.md#CreateAdvisoryDefinition) | **Post** /1.0.0/definitions/advisories | Create a new advisory definition.
[**DeleteAdvisoryDefinition**](AdvisoryApi.md#DeleteAdvisoryDefinition) | **Delete** /1.0.0/definitions/advisories/{advisoryId} | Delete the advisory with the associated specified id.
[**GetAdvisory**](AdvisoryApi.md#GetAdvisory) | **Get** /1.0.0/advisories/{advisoryId} | Get the advisory with the specified id.
[**GetAdvisoryDefinition**](AdvisoryApi.md#GetAdvisoryDefinition) | **Get** /1.0.0/definitions/advisories/{advisoryId} | Get the advisory definition with the specified id.
[**QueryAdvisories**](AdvisoryApi.md#QueryAdvisories) | **Get** /1.0.0/advisories | Get a list of all advisories accessible to the user.
[**QueryAdvisoryDefinitions**](AdvisoryApi.md#QueryAdvisoryDefinitions) | **Get** /1.0.0/definitions/advisories | Get a list of all advisory definitions.
[**UpdateAdvisory**](AdvisoryApi.md#UpdateAdvisory) | **Put** /1.0.0/advisories/{advisoryId} | Updates an advisory.


# **CreateAdvisoryDefinition**
> AdvisoryDefinition CreateAdvisoryDefinition(ctx, newAdvisoryDefinition)
Create a new advisory definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newAdvisoryDefinition** | [**AdvisoryDefinition**](AdvisoryDefinition.md)|  | 

### Return type

[**AdvisoryDefinition**](AdvisoryDefinition.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAdvisoryDefinition**
> DeleteAdvisoryDefinition(ctx, advisoryId)
Delete the advisory with the associated specified id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advisoryId** | **string**| advisory URN | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdvisory**
> Advisory GetAdvisory(ctx, advisoryId)
Get the advisory with the specified id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advisoryId** | **string**| advisory URN | 

### Return type

[**Advisory**](Advisory.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdvisoryDefinition**
> AdvisoryDefinition GetAdvisoryDefinition(ctx, advisoryId)
Get the advisory definition with the specified id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advisoryId** | **string**| advisory URN | 

### Return type

[**AdvisoryDefinition**](AdvisoryDefinition.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryAdvisories**
> Advisories QueryAdvisories(ctx, page, pageSize, optional)
Get a list of all advisories accessible to the user.

Get a list of all advisories accessible to the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***AdvisoryApiQueryAdvisoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdvisoryApiQueryAdvisoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Advisories**](Advisories.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryAdvisoryDefinitions**
> AdvisoryDefinitions QueryAdvisoryDefinitions(ctx, page, pageSize, optional)
Get a list of all advisory definitions.

Get a list of all advisory definitions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***AdvisoryApiQueryAdvisoryDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdvisoryApiQueryAdvisoryDefinitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**AdvisoryDefinitions**](AdvisoryDefinitions.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAdvisory**
> Advisory UpdateAdvisory(ctx, modifiedAdvisory, advisoryId)
Updates an advisory.

Updates an advisory for the active session using a specified id. Advisories that have a MANDATORY priority may not be updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modifiedAdvisory** | [**Advisory**](Advisory.md)|  | 
  **advisoryId** | **string**| advisory URN | 

### Return type

[**Advisory**](Advisory.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

