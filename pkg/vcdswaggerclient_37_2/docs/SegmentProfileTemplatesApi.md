# \SegmentProfileTemplatesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegmentProfileTemplate**](SegmentProfileTemplatesApi.md#CreateSegmentProfileTemplate) | **Post** /1.0.0/segmentProfileTemplates | Creates a new Segment Profile Template.
[**DeleteSegmentProfileTemplate**](SegmentProfileTemplatesApi.md#DeleteSegmentProfileTemplate) | **Delete** /1.0.0/segmentProfileTemplates/{segmentProfileTemplateId} | Deletes the specified Segment Profile Template.
[**GetGlobalDefaultSegmentProfileTemplates**](SegmentProfileTemplatesApi.md#GetGlobalDefaultSegmentProfileTemplates) | **Get** /1.0.0/segmentProfileTemplates/default | Retrieve the global defaults for segment profile templates
[**GetSegmentProfileTemplate**](SegmentProfileTemplatesApi.md#GetSegmentProfileTemplate) | **Get** /1.0.0/segmentProfileTemplates/{segmentProfileTemplateId} | Retrieves a Segment Profile Template.
[**GetSegmentProfileTemplates**](SegmentProfileTemplatesApi.md#GetSegmentProfileTemplates) | **Get** /1.0.0/segmentProfileTemplates | Retrieves all Segment Profile Templates.
[**SyncSegmentProfileTemplate**](SegmentProfileTemplatesApi.md#SyncSegmentProfileTemplate) | **Post** /1.0.0/segmentProfileTemplates/{segmentProfileTemplateId}/sync | Sync the Segment Profile Template.
[**UpdateGlobalDefaultSegmentProfileTemplates**](SegmentProfileTemplatesApi.md#UpdateGlobalDefaultSegmentProfileTemplates) | **Put** /1.0.0/segmentProfileTemplates/default | Updates the global defaults for segment profile templates
[**UpdateSegmentProfileTemplate**](SegmentProfileTemplatesApi.md#UpdateSegmentProfileTemplate) | **Put** /1.0.0/segmentProfileTemplates/{segmentProfileTemplateId} | Updates a Segment Profile Template.


# **CreateSegmentProfileTemplate**
> CreateSegmentProfileTemplate(ctx, segmentProfileTemplate)
Creates a new Segment Profile Template.

Creates a new Segment Profile Template. If needed, the segment profiles referenced in the template will be synced from the source NSX-T Manager to all known NSX-T Managers in Cloud Director. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentProfileTemplate** | [**SegmentProfileTemplate**](SegmentProfileTemplate.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSegmentProfileTemplate**
> DeleteSegmentProfileTemplate(ctx, segmentProfileTemplateId)
Deletes the specified Segment Profile Template.

Deletes the Segment Profile Template with the given ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentProfileTemplateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalDefaultSegmentProfileTemplates**
> SegmentProfileTemplateDefaultDefinition GetGlobalDefaultSegmentProfileTemplates(ctx, )
Retrieve the global defaults for segment profile templates

Retrieve the global default segment profile templates. These segment profile templates apply to all NSX-T backed networks created by Cloud Director unless overridden explicitly during create/update or by an Org vDC defined default. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SegmentProfileTemplateDefaultDefinition**](SegmentProfileTemplateDefaultDefinition.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentProfileTemplate**
> SegmentProfileTemplate GetSegmentProfileTemplate(ctx, segmentProfileTemplateId)
Retrieves a Segment Profile Template.

Retrieves a singular Segment Profile Template with the given ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentProfileTemplateId** | **string**|  | 

### Return type

[**SegmentProfileTemplate**](SegmentProfileTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentProfileTemplates**
> SegmentProfileTemplates GetSegmentProfileTemplates(ctx, page, pageSize, optional)
Retrieves all Segment Profile Templates.

Retrieves all the Segment Profile Templates available to the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***SegmentProfileTemplatesApiGetSegmentProfileTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SegmentProfileTemplatesApiGetSegmentProfileTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**SegmentProfileTemplates**](SegmentProfileTemplates.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncSegmentProfileTemplate**
> SyncSegmentProfileTemplate(ctx, segmentProfileTemplateId)
Sync the Segment Profile Template.

Sync the Segment Profile Template. An example usage is to detect if a segment profile referenced by this template still exists/is valid. The segment profiles referenced in the template will be synced from the source NSX-T Manager to all known NSX-T Managers in Cloud Director, if needed. If previously synced, this will overwrite the profiles on the target NSX-T managers with the source profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentProfileTemplateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGlobalDefaultSegmentProfileTemplates**
> UpdateGlobalDefaultSegmentProfileTemplates(ctx, segmentProfileTemplateDefaultDefinition)
Updates the global defaults for segment profile templates

Updates the global default segment profile templates. These segment profile templates apply to all NSX-T backed networks created by Cloud Director unless overridden explicitly during create/update or by an Org vDC defined default. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentProfileTemplateDefaultDefinition** | [**SegmentProfileTemplateDefaultDefinition**](SegmentProfileTemplateDefaultDefinition.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSegmentProfileTemplate**
> UpdateSegmentProfileTemplate(ctx, segmentProfileTemplate, segmentProfileTemplateId)
Updates a Segment Profile Template.

Updates the Segment Profile Template with the given ID. If needed, the segment profiles referenced in the template will be synced from the source NSX-T Manager to all known NSX-T Managers in Cloud Director. If the source NSX-T Manager is updated, all updates to profiles will be ignored within the same request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentProfileTemplate** | [**SegmentProfileTemplate**](SegmentProfileTemplate.md)|  | 
  **segmentProfileTemplateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

