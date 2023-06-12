# \NsxVResourcesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImportableTransportZones**](NsxVResourcesApi.md#GetImportableTransportZones) | **Get** /1.0.0/nsxVResources/importableTransportZones | Get all importable vxlan transport zones that are configured on a vCenter backed by a NSX-V manager.


# **GetImportableTransportZones**
> NsxVTransportZones GetImportableTransportZones(ctx, optional)
Get all importable vxlan transport zones that are configured on a vCenter backed by a NSX-V manager.

Get all importable transport zones that are configured on a vCenter backed by a NSX-V manager. Transport zones that are already associated with a network pool are filtered out. Only \"_context\" filter is supported and \"_context\" filter key must be set with the id of the vCenter which we want to get the transport zones for. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NsxVResourcesApiGetImportableTransportZonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxVResourcesApiGetImportableTransportZonesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter for a query.  FIQL format. | 

### Return type

[**NsxVTransportZones**](NsxVTransportZones.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

