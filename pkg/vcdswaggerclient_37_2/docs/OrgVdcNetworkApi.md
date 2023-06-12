# \OrgVdcNetworkApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetwork**](OrgVdcNetworkApi.md#DeleteNetwork) | **Delete** /1.0.0/orgVdcNetworks/{vdcNetworkId} | Deletes a specific Org vDC network.
[**GetAdditionalProperties**](OrgVdcNetworkApi.md#GetAdditionalProperties) | **Get** /1.0.0/orgVdcNetworks/{vdcNetworkId}/additionalProperties | Returns the additional properties of a specific Org vDC network. This is a System Administrator only API.
[**GetAllocatedIpAddresses**](OrgVdcNetworkApi.md#GetAllocatedIpAddresses) | **Get** /1.0.0/orgVdcNetworks/{vdcNetworkId}/allocatedIpAddresses | Retrieve the list of IP addresses allocated to the network.
[**GetNetworkSegmentProfiles**](OrgVdcNetworkApi.md#GetNetworkSegmentProfiles) | **Get** /1.0.0/orgVdcNetworks/{vdcNetworkId}/segmentProfiles | Retrieves the segment profiles configuration for an Org vDC Network.
[**GetOrgVdcNetwork**](OrgVdcNetworkApi.md#GetOrgVdcNetwork) | **Get** /1.0.0/orgVdcNetworks/{vdcNetworkId} | Retrieves a specific Org vDC network.
[**GetSecondaryAllocatedIpAddresses**](OrgVdcNetworkApi.md#GetSecondaryAllocatedIpAddresses) | **Get** /1.0.0/orgVdcNetworks/{vdcNetworkId}/secondaryAllocatedIpAddresses | Retrieve the list of secondary IP addresses allocated to the network, if the network is a dual stack network.
[**ResetNetwork**](OrgVdcNetworkApi.md#ResetNetwork) | **Post** /1.0.0/orgVdcNetworks/{vdcNetworkId}/reset | Reset a specific isolated Org vDC network.
[**SyncOrgVdcNetwork**](OrgVdcNetworkApi.md#SyncOrgVdcNetwork) | **Post** /1.0.0/orgVdcNetworks/{vdcNetworkId}/sync | Sync/repair a specific Org vDC network.
[**SyncSyslogSettingsOfNetwork**](OrgVdcNetworkApi.md#SyncSyslogSettingsOfNetwork) | **Post** /1.0.0/orgVdcNetworks/{vdcNetworkId}/syncSyslog | Synchronize syslog server settings for a Org vDC network.
[**UpdateNetwork**](OrgVdcNetworkApi.md#UpdateNetwork) | **Put** /1.0.0/orgVdcNetworks/{vdcNetworkId} | Updates a specific Org vDC network.
[**UpdateNetworkSegmentProfiles**](OrgVdcNetworkApi.md#UpdateNetworkSegmentProfiles) | **Put** /1.0.0/orgVdcNetworks/{vdcNetworkId}/segmentProfiles | Updates the segment profiles configuration for an Org vDC Network.


# **DeleteNetwork**
> DeleteNetwork(ctx, vdcNetworkId, optional)
Deletes a specific Org vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 
 **optional** | ***OrgVdcNetworkApiDeleteNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgVdcNetworkApiDeleteNetworkOpts struct

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

# **GetAdditionalProperties**
> OrgVdcNetworkAdditionalProperties GetAdditionalProperties(ctx, vdcNetworkId)
Returns the additional properties of a specific Org vDC network. This is a System Administrator only API.

Returns the set of additional properties for the given Org vDC Network 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

[**OrgVdcNetworkAdditionalProperties**](OrgVdcNetworkAdditionalProperties.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllocatedIpAddresses**
> AllocatedIpAddresses GetAllocatedIpAddresses(ctx, page, pageSize, vdcNetworkId, optional)
Retrieve the list of IP addresses allocated to the network.

Get all the allocated IPs for a given Org vDC network. This returns all the IP addresses of network which are allocated to a vApp VM, an edge gateway interface and the addresses being used in a NAT routed environment. Results can be filtered by IP address. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **vdcNetworkId** | **string**|  | 
 **optional** | ***OrgVdcNetworkApiGetAllocatedIpAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgVdcNetworkApiGetAllocatedIpAddressesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**AllocatedIpAddresses**](AllocatedIpAddresses.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkSegmentProfiles**
> VdcNetworkSegmentProfileConfig GetNetworkSegmentProfiles(ctx, vdcNetworkId)
Retrieves the segment profiles configuration for an Org vDC Network.

Retrieves the segment profiles configured for an Org vDC Network. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

[**VdcNetworkSegmentProfileConfig**](VdcNetworkSegmentProfileConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgVdcNetwork**
> VdcNetwork GetOrgVdcNetwork(ctx, vdcNetworkId)
Retrieves a specific Org vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

[**VdcNetwork**](VdcNetwork.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecondaryAllocatedIpAddresses**
> AllocatedIpAddresses GetSecondaryAllocatedIpAddresses(ctx, page, pageSize, vdcNetworkId, optional)
Retrieve the list of secondary IP addresses allocated to the network, if the network is a dual stack network.

Get all the secondary allocated IPs for a given Org vDC network. This returns all the IP addresses of network which are allocated to a vApp VM, an edge gateway interface, from the IPv6 subnet of the Org vDC network. Results can be filtered by IP address. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **vdcNetworkId** | **string**|  | 
 **optional** | ***OrgVdcNetworkApiGetSecondaryAllocatedIpAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgVdcNetworkApiGetSecondaryAllocatedIpAddressesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**AllocatedIpAddresses**](AllocatedIpAddresses.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetNetwork**
> ResetNetwork(ctx, vdcNetworkId)
Reset a specific isolated Org vDC network.

Reset a specific isolated Org vDC network.  Reset involves redeploying the internal edge gateway of the isolated Org vDC Network if present. An error is returned if the network is not isolated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncOrgVdcNetwork**
> SyncOrgVdcNetwork(ctx, vdcNetworkId)
Sync/repair a specific Org vDC network.

Sync/repair the vDC Group Cross vDC network. An example usage is to realize a network in the participating vDC which was unreachable when the network was created. This operation is only allowed for VIRTUAL_WIRE backed cross vDC networks. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncSyslogSettingsOfNetwork**
> SyncSyslogSettingsOfNetwork(ctx, vdcNetworkId)
Synchronize syslog server settings for a Org vDC network.

When the IP addresses of the primary or secondary syslog server are updated, this api can synchronize syslog server settings of an Isolated Org vDC Network against the vCD Syslog Settings. An error is returned if network is a direct network. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetwork**
> UpdateNetwork(ctx, vdcNetwork, vdcNetworkId)
Updates a specific Org vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetwork** | [**VdcNetwork**](VdcNetwork.md)|  | 
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkSegmentProfiles**
> UpdateNetworkSegmentProfiles(ctx, networkSegmentProfiles, vdcNetworkId)
Updates the segment profiles configuration for an Org vDC Network.

Updates the segment profiles configuration for an Org vDC Network. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkSegmentProfiles** | [**VdcNetworkSegmentProfileConfig**](VdcNetworkSegmentProfileConfig.md)|  | 
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

