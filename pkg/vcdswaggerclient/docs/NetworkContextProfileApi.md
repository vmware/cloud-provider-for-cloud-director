# \NetworkContextProfileApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetworkContextProfile**](NetworkContextProfileApi.md#DeleteNetworkContextProfile) | **Delete** /1.0.0/networkContextProfiles/{profileId} | Deletes a specific network context profile, removing the associated firewall rule and permitting the traffic this profile restricts.
[**GetNetworkContextProfile**](NetworkContextProfileApi.md#GetNetworkContextProfile) | **Get** /1.0.0/networkContextProfiles/{profileId} | Get a specific network context profile.
[**UpdateNetworkContextProfile**](NetworkContextProfileApi.md#UpdateNetworkContextProfile) | **Put** /1.0.0/networkContextProfiles/{profileId} | Updates a specific user-defined network context profile, changing the associated firewall and modifying the traffic this profile restricts.


# **DeleteNetworkContextProfile**
> DeleteNetworkContextProfile(ctx, profileId)
Deletes a specific network context profile, removing the associated firewall rule and permitting the traffic this profile restricts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkContextProfile**
> NetworkContextProfile GetNetworkContextProfile(ctx, profileId)
Get a specific network context profile.

Retrieves a single network context profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

[**NetworkContextProfile**](NetworkContextProfile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkContextProfile**
> UpdateNetworkContextProfile(ctx, networkContextProfile, profileId)
Updates a specific user-defined network context profile, changing the associated firewall and modifying the traffic this profile restricts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkContextProfile** | [**NetworkContextProfile**](NetworkContextProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

