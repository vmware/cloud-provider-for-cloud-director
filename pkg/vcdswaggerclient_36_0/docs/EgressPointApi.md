# \EgressPointApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEgressPoint**](EgressPointApi.md#DeleteEgressPoint) | **Delete** /1.0.0/universalRouters/{universalRouterId}/routing/egressPoints/{egressPointId} | Deletes a specific Universal Egress Point.
[**GetEgressPoint**](EgressPointApi.md#GetEgressPoint) | **Get** /1.0.0/universalRouters/{universalRouterId}/routing/egressPoints/{egressPointId} | Retrieves a specific Universal Egress Point.
[**SyncEgressPoint**](EgressPointApi.md#SyncEgressPoint) | **Post** /1.0.0/universalRouters/{universalRouterId}/routing/egressPoints/{egressPointId}/sync | Sync/repair the egress point.


# **DeleteEgressPoint**
> DeleteEgressPoint(ctx, universalRouterId, egressPointId, optional)
Deletes a specific Universal Egress Point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 
  **egressPointId** | **string**|  | 
 **optional** | ***EgressPointApiDeleteEgressPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EgressPointApiDeleteEgressPointOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| Value \&quot;true\&quot; means to forcefully delete the object that contains other objects even if those objects are in a state that does not allow removal. The default is \&quot;false\&quot;; therefore, objects are not removed if they are not in a state that normally allows removal. Force also implies recursive delete where other contained objects are removed. Errors may be ignored. Invalid value (not true or false) are ignored.  | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEgressPoint**
> EgressPoint GetEgressPoint(ctx, universalRouterId, egressPointId)
Retrieves a specific Universal Egress Point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 
  **egressPointId** | **string**|  | 

### Return type

[**EgressPoint**](EgressPoint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncEgressPoint**
> SyncEgressPoint(ctx, universalRouterId, egressPointId)
Sync/repair the egress point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **universalRouterId** | **string**|  | 
  **egressPointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

