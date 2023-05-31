# \NsxTResourcesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGatewayQoSProfiles**](NsxTResourcesApi.md#GetGatewayQoSProfiles) | **Get** /1.0.0/nsxTResources/gatewayQoSProfiles | Get all Gateway QoS Profiles configured on an NSX-T manager.
[**GetImportableSegmentSubnet**](NsxTResourcesApi.md#GetImportableSegmentSubnet) | **Get** /1.0.0/nsxTResources/importableSegments/{nsxTManagerId}/{switchId}/subnet | Get subnet info for an importable segment
[**GetImportableSegments**](NsxTResourcesApi.md#GetImportableSegments) | **Get** /1.0.0/nsxTResources/importableSegments | Get all importable segments
[**GetImportableTier0Routers**](NsxTResourcesApi.md#GetImportableTier0Routers) | **Get** /1.0.0/nsxTResources/importableTier0Routers | Get all importable Tier-0 routers that are accessible to an organization vDC.
[**GetImportableTransportZones**](NsxTResourcesApi.md#GetImportableTransportZones) | **Get** /1.0.0/nsxTResources/importableTransportZones | Get all importable overlay transport zones that are configured on an NSX-T manager.
[**GetNsxTEdgeClusters**](NsxTResourcesApi.md#GetNsxTEdgeClusters) | **Get** /1.0.0/nsxTResources/edgeClusters | Get all edge clusters that are configured on an NSX-T manager.
[**GetSegmentIpDiscoveryProfiles**](NsxTResourcesApi.md#GetSegmentIpDiscoveryProfiles) | **Get** /1.0.0/nsxTResources/segmentIpDiscoveryProfiles | Get all segment IP Discovery Profiles configured on an NSX-T manager.
[**GetSegmentMacDiscoveryProfiles**](NsxTResourcesApi.md#GetSegmentMacDiscoveryProfiles) | **Get** /1.0.0/nsxTResources/segmentMacDiscoveryProfiles | Get all segment MAC Discovery Profiles configured on an NSX-T manager.
[**GetSegmentQoSProfiles**](NsxTResourcesApi.md#GetSegmentQoSProfiles) | **Get** /1.0.0/nsxTResources/segmentQoSProfiles | Get all segment QoS Profiles configured on an NSX-T manager.
[**GetSegmentSecurityProfiles**](NsxTResourcesApi.md#GetSegmentSecurityProfiles) | **Get** /1.0.0/nsxTResources/segmentSecurityProfiles | Get all segment Security Profiles configured on an NSX-T manager.
[**GetSegmentSpoofGuardProfiles**](NsxTResourcesApi.md#GetSegmentSpoofGuardProfiles) | **Get** /1.0.0/nsxTResources/segmentSpoofGuardProfiles | Get all segment Spoof Guard Profiles configured on an NSX-T manager.


# **GetGatewayQoSProfiles**
> GatewayQoSProfiles GetGatewayQoSProfiles(ctx, pageSize, optional)
Get all Gateway QoS Profiles configured on an NSX-T manager.

Get all Gateway QoS Profiles configured on an NSX-T manager. NSX-T manager ID (nsxTManagerRef.id), Org VDC ID (orgVdcId) or VDC Group ID (vdcGroupId) must be supplied as a filter. Results can also be filtered by a single ID (example: filter=orgVdcId==URN;id==profileId) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxTResourcesApiGetGatewayQoSProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxTResourcesApiGetGatewayQoSProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 

### Return type

[**GatewayQoSProfiles**](GatewayQoSProfiles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportableSegmentSubnet**
> NsxTImportableSegmentSubnet GetImportableSegmentSubnet(ctx, nsxTManagerId, switchId)
Get subnet info for an importable segment

Get subnet info for an importable segment. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsxTManagerId** | **string**| Id of the Nsx-T manager that this segment belongs to | 
  **switchId** | **string**| The logical switch id for the segment. Corresponds to the unique id property in Nsx-T. | 

### Return type

[**NsxTImportableSegmentSubnet**](NsxTImportableSegmentSubnet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportableSegments**
> NsxTImportableSegments GetImportableSegments(ctx, pageSize, optional)
Get all importable segments

Get all importable segments for an OrgVDC or VdcGroup 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxTResourcesApiGetImportableSegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxTResourcesApiGetImportableSegmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 

### Return type

[**NsxTImportableSegments**](NsxTImportableSegments.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
 - **Accept**: application/json;version=37.2

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
 - **Accept**: application/json;version=37.2

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
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentIpDiscoveryProfiles**
> SegmentIpDiscoveryProfiles GetSegmentIpDiscoveryProfiles(ctx, pageSize, optional)
Get all segment IP Discovery Profiles configured on an NSX-T manager.

Get all segment IP Discovery Profiles configured on an NSX-T manager. NSX-T manager ID (nsxTManagerRef.id), Org VDC ID (orgVdcId) or VDC Group ID (vdcGroupId) must be supplied as a filter. Results can also be filtered by a single profile ID (filter=nsxTManagerRef.id==nsxTManagerUrn;id==profileId). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxTResourcesApiGetSegmentIpDiscoveryProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxTResourcesApiGetSegmentIpDiscoveryProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 

### Return type

[**SegmentIpDiscoveryProfiles**](SegmentIpDiscoveryProfiles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentMacDiscoveryProfiles**
> SegmentMacDiscoveryProfiles GetSegmentMacDiscoveryProfiles(ctx, pageSize, optional)
Get all segment MAC Discovery Profiles configured on an NSX-T manager.

Get all segment MAC Discovery Profiles configured on an NSX-T manager. NSX-T manager ID (nsxTManagerRef.id), Org VDC ID (orgVdcId) or VDC Group ID (vdcGroupId) must be supplied as a filter. Results can also be filtered by a single profile ID (filter=nsxTManagerRef.id==nsxTManagerUrn;id==profileId). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxTResourcesApiGetSegmentMacDiscoveryProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxTResourcesApiGetSegmentMacDiscoveryProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 

### Return type

[**SegmentMacDiscoveryProfiles**](SegmentMacDiscoveryProfiles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentQoSProfiles**
> SegmentQoSProfiles GetSegmentQoSProfiles(ctx, pageSize, optional)
Get all segment QoS Profiles configured on an NSX-T manager.

Get all segment QoS Profiles configured on an NSX-T manager. NSX-T manager ID (nsxTManagerRef.id), Org VDC ID (orgVdcId) or VDC Group ID (vdcGroupId) must be supplied as a filter. Results can also be filtered by a single profile ID (filter=nsxTManagerRef.id==nsxTManagerUrn;id==profileId). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxTResourcesApiGetSegmentQoSProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxTResourcesApiGetSegmentQoSProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 

### Return type

[**SegmentQoSProfiles**](SegmentQoSProfiles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentSecurityProfiles**
> SegmentSecurityProfiles GetSegmentSecurityProfiles(ctx, pageSize, optional)
Get all segment Security Profiles configured on an NSX-T manager.

Get all segment Security Profiles configured on an NSX-T manager. NSX-T manager ID (nsxTManagerRef.id), Org VDC ID (orgVdcId) or VDC Group ID (vdcGroupId) must be supplied as a filter. Results can also be filtered by a single profile ID (filter=nsxTManagerRef.id==nsxTManagerUrn;id==profileId). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxTResourcesApiGetSegmentSecurityProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxTResourcesApiGetSegmentSecurityProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 

### Return type

[**SegmentSecurityProfiles**](SegmentSecurityProfiles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentSpoofGuardProfiles**
> SegmentSpoofGuardProfiles GetSegmentSpoofGuardProfiles(ctx, pageSize, optional)
Get all segment Spoof Guard Profiles configured on an NSX-T manager.

Get all segment Spoof Guard Profiles configured on an NSX-T manager. NSX-T manager ID (nsxTManagerRef.id), Org VDC ID (orgVdcId) or VDC Group ID (vdcGroupId) must be supplied as a filter. Results can also be filtered by a single profile ID (filter=nsxTManagerRef.id==nsxTManagerUrn;id==profileId). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***NsxTResourcesApiGetSegmentSpoofGuardProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NsxTResourcesApiGetSegmentSpoofGuardProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **cursor** | **optional.String**| Field used for getting next page of records. The value is supplied by the current result page. If not set, the first page is retrieved. If cursor is set, then all other pagination query parameters such as pageSize, sortDesc, sortAsc, queryFilter are ignored.  | 

### Return type

[**SegmentSpoofGuardProfiles**](SegmentSpoofGuardProfiles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

