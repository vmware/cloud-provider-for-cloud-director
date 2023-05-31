# \IpSpaceUplinksApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpSpaceUplink**](IpSpaceUplinksApi.md#CreateIpSpaceUplink) | **Post** /1.0.0/ipSpaceUplinks | Create a new IP Space Uplink.
[**DeleteIpSpaceUplink**](IpSpaceUplinksApi.md#DeleteIpSpaceUplink) | **Delete** /1.0.0/ipSpaceUplinks/{ipSpaceUplinkId} | Delete an IP Space Uplink.
[**GetIpSpaceUplink**](IpSpaceUplinksApi.md#GetIpSpaceUplink) | **Get** /1.0.0/ipSpaceUplinks/{ipSpaceUplinkId} | Get an IP Space Uplink.
[**GetIpSpaceUplinks**](IpSpaceUplinksApi.md#GetIpSpaceUplinks) | **Get** /1.0.0/ipSpaceUplinks | Retrieves all the IP Space Uplinks.
[**UpdateIpSpaceUplink**](IpSpaceUplinksApi.md#UpdateIpSpaceUplink) | **Put** /1.0.0/ipSpaceUplinks/{ipSpaceUplinkId} | Update an IP Space Uplink.


# **CreateIpSpaceUplink**
> CreateIpSpaceUplink(ctx, ipSpaceUplinkConfig)
Create a new IP Space Uplink.

Create a new IP Space Uplink in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceUplinkConfig** | [**IpSpaceUplink**](IpSpaceUplink.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpSpaceUplink**
> DeleteIpSpaceUplink(ctx, ipSpaceUplinkId)
Delete an IP Space Uplink.

Deletes the specified IP Space Uplink. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceUplinkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSpaceUplink**
> IpSpaceUplink GetIpSpaceUplink(ctx, ipSpaceUplinkId)
Get an IP Space Uplink.

Retrieves the specified IP Space Uplink. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceUplinkId** | **string**|  | 

### Return type

[**IpSpaceUplink**](IpSpaceUplink.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSpaceUplinks**
> IpSpaceUplinks GetIpSpaceUplinks(ctx, page, pageSize, optional)
Retrieves all the IP Space Uplinks.

Get all the IP Space Uplinks for a specified Provider Gateway. Note that the filter parameter \"externalNetworkRef.id\" is required. An External Network is used to reference the Provider Gateway since the External Network is backed by a Provider Gateway. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***IpSpaceUplinksApiGetIpSpaceUplinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IpSpaceUplinksApiGetIpSpaceUplinksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**IpSpaceUplinks**](IpSpaceUplinks.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpSpaceUplink**
> UpdateIpSpaceUplink(ctx, ipSpaceUplinkConfig, ipSpaceUplinkId)
Update an IP Space Uplink.

Updates the specified IP Space Uplink. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceUplinkConfig** | [**IpSpaceUplink**](IpSpaceUplink.md)|  | 
  **ipSpaceUplinkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

