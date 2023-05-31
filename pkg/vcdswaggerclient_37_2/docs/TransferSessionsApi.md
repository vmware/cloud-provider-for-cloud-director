# \TransferSessionsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransferSession**](TransferSessionsApi.md#GetTransferSession) | **Get** /1.0.0/transferSessions/{transferSessionUrn} | Get specified transfer session
[**QueryTransferSessionItems**](TransferSessionsApi.md#QueryTransferSessionItems) | **Get** /1.0.0/transferSessions/{transferSessionUrn}/transferItems | Get list of transfer items associated with this transfer session.
[**QueryTransferSessions**](TransferSessionsApi.md#QueryTransferSessions) | **Get** /1.0.0/transferSessions | Get list of active transfer sessions


# **GetTransferSession**
> TransferSession GetTransferSession(ctx, transferSessionUrn)
Get specified transfer session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transferSessionUrn** | **string**|  | 

### Return type

[**TransferSession**](TransferSession.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryTransferSessionItems**
> TransferSessionItems QueryTransferSessionItems(ctx, transferSessionUrn, page, pageSize, optional)
Get list of transfer items associated with this transfer session.

Get list of transfer session items associated with this transfer session. <br> Results can be filtered by: <ul>   <li> name </li> <ul> <br> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transferSessionUrn** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***TransferSessionsApiQueryTransferSessionItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransferSessionsApiQueryTransferSessionItemsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**TransferSessionItems**](TransferSessionItems.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryTransferSessions**
> TransferSessions QueryTransferSessions(ctx, page, pageSize, optional)
Get list of active transfer sessions

Get list of active transfer sessions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***TransferSessionsApiQueryTransferSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransferSessionsApiQueryTransferSessionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**TransferSessions**](TransferSessions.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

