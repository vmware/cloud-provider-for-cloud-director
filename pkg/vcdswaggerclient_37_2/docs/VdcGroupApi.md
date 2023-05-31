# \VdcGroupApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVdcGroup**](VdcGroupApi.md#DeleteVdcGroup) | **Delete** /1.0.0/vdcGroups/{vdcGroupId} | Deletes a vDC Group
[**GetVdcGroup**](VdcGroupApi.md#GetVdcGroup) | **Get** /1.0.0/vdcGroups/{vdcGroupId} | Retrieves a specific vDC Group.
[**SyncVdcGroup**](VdcGroupApi.md#SyncVdcGroup) | **Post** /1.0.0/vdcGroups/{vdcGroupId}/sync | Sync/repair the vDC group. An example usage is to detect if a vDC still exists/is valid. If an Organization vDC referenced by the VDC group is deleted or if it is not participating in universal networking, it&#39;s status will be updated to OBJECT_NOT_FOUND and the vdc group will be marked as NOT_REALIZED. This will also initiate a sync of associated router, if any. The router entities like egress points and universal routes will also be marked as NOT_REALIZED if they reference the removed Organization vDC. 
[**UpdateVdcGroup**](VdcGroupApi.md#UpdateVdcGroup) | **Put** /1.0.0/vdcGroups/{vdcGroupId} | Updates a specific vDC Group.  Example is to add/remove a participarting vDC.


# **DeleteVdcGroup**
> DeleteVdcGroup(ctx, vdcGroupId, optional)
Deletes a vDC Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 
 **optional** | ***VdcGroupApiDeleteVdcGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VdcGroupApiDeleteVdcGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Value \&quot;true\&quot; means to forcefully delete the object that contains other objects even if those objects are in a state that does not allow removal. The default is \&quot;false\&quot;; therefore, objects are not removed if they are not in a state that normally allows removal. Force also implies recursive delete where other contained objects are removed. Errors may be ignored. Invalid value (not true or false) are ignored.  | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcGroup**
> VdcGroup GetVdcGroup(ctx, vdcGroupId)
Retrieves a specific vDC Group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 

### Return type

[**VdcGroup**](VdcGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncVdcGroup**
> SyncVdcGroup(ctx, vdcGroupId)
Sync/repair the vDC group. An example usage is to detect if a vDC still exists/is valid. If an Organization vDC referenced by the VDC group is deleted or if it is not participating in universal networking, it's status will be updated to OBJECT_NOT_FOUND and the vdc group will be marked as NOT_REALIZED. This will also initiate a sync of associated router, if any. The router entities like egress points and universal routes will also be marked as NOT_REALIZED if they reference the removed Organization vDC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVdcGroup**
> UpdateVdcGroup(ctx, vdcGroup, vdcGroupId)
Updates a specific vDC Group.  Example is to add/remove a participarting vDC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcGroup** | [**VdcGroup**](VdcGroup.md)|  | 
  **vdcGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

