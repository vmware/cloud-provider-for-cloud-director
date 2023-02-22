# \NsxTResourcesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImportableTier0Routers**](NsxTResourcesApi.md#GetImportableTier0Routers) | **Get** /1.0.0/nsxTResources/importableTier0Routers | Get all importable Tier-0 routers that are accessible to an organization vDC.
[**GetImportableTransportZones**](NsxTResourcesApi.md#GetImportableTransportZones) | **Get** /1.0.0/nsxTResources/importableTransportZones | Get all importable overlay transport zones that are configured on an NSX-T manager.
[**GetNsxTEdgeClusters**](NsxTResourcesApi.md#GetNsxTEdgeClusters) | **Get** /1.0.0/nsxTResources/edgeClusters | Get all edge clusters that are configured on an NSX-T manager.


# **GetImportableTier0Routers**
> Tier0Routers GetImportableTier0Routers(ctx, pageSize, optional)
Get all importable Tier-0 routers that are accessible to an organization vDC.

Get all Tier-0 routers that are accessible to an organization vDC. Routers that are already associated with an External Network are filtered out. The \"_context\" filter key must be set with the id of the NSX-T manager for which we want to get the Tier-0 routers for. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxTResourcesApiGetImportableTier0RoutersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxTResourcesApiGetImportableTier0RoutersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Tier0Routers**](Tier0Routers.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportableTransportZones**
> NsxTTransportZones GetImportableTransportZones(ctx, pageSize, optional)
Get all importable overlay transport zones that are configured on an NSX-T manager.

Get all importable overlay transport zones that are configured on an NSX-T manager. Transport zones that are already associated with a network pool are filtered out. The \"_context\" filter key must be set with the id of the NSX-T manager which we want to get the transport zones for. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxTResourcesApiGetImportableTransportZonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxTResourcesApiGetImportableTransportZonesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**NsxTTransportZones**](NsxTTransportZones.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNsxTEdgeClusters**
> NsxTEdgeClusters GetNsxTEdgeClusters(ctx, pageSize, optional)
Get all edge clusters that are configured on an NSX-T manager.

Returns all the configured NSX-T edge clusters for an Org vDC or a VDC Group or a Provider VDC. Supported filters are: <ul> <li>orgVdcId - | The filter orgVdcId must be set equal to the id of the NSX-T backed Org vDC for which we want to get the edge clusters. Example: (orgVdcId==urn:vcloud:vdc:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx) <li>vdcGroupId - | The filter vdcGroupId must be set equal to the id of the NSX-T VDC Group for which we want to get the edge clusters. Example: (vdcGroupId==urn:vcloud:vdcGroup:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx) <li>pvdcId - | The filter pvdcId must be set equal to the id of the NSX-T backed Provider VDC for which we want to get the edge clusters. pvdcId filter is supported from version 35.2 Example: (pvdcId==urn:vcloud:providervdc:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx) </ul> Use of \"_context\" filter has been deprecated. Please use supported filters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxTResourcesApiGetNsxTEdgeClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxTResourcesApiGetNsxTEdgeClustersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**NsxTEdgeClusters**](NsxTEdgeClusters.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

