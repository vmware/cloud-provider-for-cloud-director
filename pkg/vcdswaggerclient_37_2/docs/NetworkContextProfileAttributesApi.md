# \NetworkContextProfileAttributesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkContextProfileAttributes**](NetworkContextProfileAttributesApi.md#GetNetworkContextProfileAttributes) | **Get** /1.0.0/networkContextProfiles/attributes | List all supported network context profile attributes and sub-attributes for the given NSX-T manager.


# **GetNetworkContextProfileAttributes**
> NetworkContextProfileAttributes GetNetworkContextProfileAttributes(ctx, optional)
List all supported network context profile attributes and sub-attributes for the given NSX-T manager.

Retrieves all available network context profile attributes and sub-attributes for the given NSX-T manager, based on the filter parameter given in FIQL format (e.g. filter=\"_context==urn:vcloud:nsxtmanager:<some_id>\"). Optionally filter by attribute type by adding a FIQL name parameter to the above filter context (e.g. filter=\"_context==<urn>;name=APP_ID\") 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkContextProfileAttributesApiGetNetworkContextProfileAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkContextProfileAttributesApiGetNetworkContextProfileAttributesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**NetworkContextProfileAttributes**](NetworkContextProfileAttributes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

