# \VgpuProfilesConsumersApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryVgpuProfilesConsumers**](VgpuProfilesConsumersApi.md#QueryVgpuProfilesConsumers) | **Get** /1.0.0/vgpuProfiles/consumers | Get list of vGPU profiles consumer entities


# **QueryVgpuProfilesConsumers**
> VgpuConsumerEntities QueryVgpuProfilesConsumers(ctx, page, pageSize, optional)
Get list of vGPU profiles consumer entities

Get list of vGPU profiles consumer entities. <br> Results can be filtered by: <ul>   <li> vgpuProfileName </li>   <li> tenant.id </li>   <li> vdc.id </li>   <li> policy.id </li> <ul> <br> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VgpuProfilesConsumersApiQueryVgpuProfilesConsumersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VgpuProfilesConsumersApiQueryVgpuProfilesConsumersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VgpuConsumerEntities**](VgpuConsumerEntities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

