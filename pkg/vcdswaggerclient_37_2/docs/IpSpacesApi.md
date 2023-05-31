# \IpSpacesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Allocate**](IpSpacesApi.md#Allocate) | **Post** /1.0.0/ipSpaces/{ipSpaceId}/allocate | Allocate floating IP addresses or IP Prefix blocks from an IP Space.
[**CreateIpSpace**](IpSpacesApi.md#CreateIpSpace) | **Post** /1.0.0/ipSpaces | Create a new IP Space.
[**DeleteIpSpace**](IpSpacesApi.md#DeleteIpSpace) | **Delete** /1.0.0/ipSpaces/{ipSpaceId} | Delete an IP Space.
[**DeleteIpSpaceIpAllocation**](IpSpacesApi.md#DeleteIpSpaceIpAllocation) | **Delete** /1.0.0/ipSpaces/{ipSpaceId}/allocations/{ipAllocationId} | Delete an IP Space IP Allocation.
[**GetFloatingIpSuggestions**](IpSpacesApi.md#GetFloatingIpSuggestions) | **Get** /1.0.0/ipSpaces/floatingIpSuggestions | Suggests IP addresses to use in Edge Gateway services.
[**GetIpPrefixSequences**](IpSpacesApi.md#GetIpPrefixSequences) | **Get** /1.0.0/ipSpaces/ipPrefixSequences | Retrieves the generated sequences for an IP Prefix.
[**GetIpPrefixSuggestions**](IpSpacesApi.md#GetIpPrefixSuggestions) | **Get** /1.0.0/ipSpaces/ipPrefixSuggestions | Suggests IP Prefixes to use for network definitions.
[**GetIpSpace**](IpSpacesApi.md#GetIpSpace) | **Get** /1.0.0/ipSpaces/{ipSpaceId} | Get an IP Space.
[**GetIpSpaceIpAllocation**](IpSpacesApi.md#GetIpSpaceIpAllocation) | **Get** /1.0.0/ipSpaces/{ipSpaceId}/allocations/{ipAllocationId} | Get an IP Space IP allocation.
[**GetIpSpaceIpAllocations**](IpSpacesApi.md#GetIpSpaceIpAllocations) | **Get** /1.0.0/ipSpaces/{ipSpaceId}/allocations | Retrieves all the allocated IP addresses or IP Prefixes of an IP Space.
[**GetIpSpaceSummaries**](IpSpacesApi.md#GetIpSpaceSummaries) | **Get** /1.0.0/ipSpaces/summaries | Retrieves all the IP Space summaries.
[**UpdateIpSpace**](IpSpacesApi.md#UpdateIpSpace) | **Put** /1.0.0/ipSpaces/{ipSpaceId} | Update an IP Space.
[**UpdateIpSpaceIpAllocation**](IpSpacesApi.md#UpdateIpSpaceIpAllocation) | **Put** /1.0.0/ipSpaces/{ipSpaceId}/allocations/{ipAllocationId} | Update an IP Space IP Allocation.


# **Allocate**
> Allocate(ctx, allocationConfig, ipSpaceId)
Allocate floating IP addresses or IP Prefix blocks from an IP Space.

Allocate floating IP addresses or IP Prefix blocks from an IP Space. This results in reserving the IP address or IP Prefix block for the specified organization. The organization can then use the IP address for network services such as NAT or use the IP Prefix as the network CIDR definition during Org VDC network creation. An IP Space IP allocation request can either request a specific IP address or IP prefix, or a request can allocate a given number of any free IP Addresses or IP Prefixes within an IP Space. These two types of requests cannot be combined to request both a specific IP Address/Prefix or any number of IP Addresses/Prefixes simultaneously. Please either request a specific value or request a variable number of IP Addresses/Prefixes with different POST requests. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **allocationConfig** | [**IpSpaceIpAllocationRequest**](IpSpaceIpAllocationRequest.md)|  | 
  **ipSpaceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIpSpace**
> CreateIpSpace(ctx, ipSpaceConfig)
Create a new IP Space.

Create a new IP Space in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceConfig** | [**IpSpace**](IpSpace.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpSpace**
> DeleteIpSpace(ctx, ipSpaceId)
Delete an IP Space.

Deletes the specified IP Space. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpSpaceIpAllocation**
> DeleteIpSpaceIpAllocation(ctx, ipSpaceId, ipAllocationId)
Delete an IP Space IP Allocation.

Deletes the specified IP Space IP Allocation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceId** | **string**|  | 
  **ipAllocationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFloatingIpSuggestions**
> IpSpaceFloatingIpSuggestions GetFloatingIpSuggestions(ctx, page, pageSize, optional)
Suggests IP addresses to use in Edge Gateway services.

Suggests IP addresses to use for Edge Gateway Services. \"gatewayId\" filter is required. Based on the specified Edge Gateway, VCD will query all the applicable IP Spaces and suggest some IP addresses which can be utilized to configure the network services on the Edge Gateway. IP Space IP addresses which are are allocated but not currently used for any network services are considered. Results can also be filtered by IPV4 or IPV6 IP address types. Filter examples:<code>(filter=gatewayId==URN)</code>, <code>(filter=gatewayId==URN;ipType==IPV6)</code> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***IpSpacesApiGetFloatingIpSuggestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IpSpacesApiGetFloatingIpSuggestionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**IpSpaceFloatingIpSuggestions**](IpSpaceFloatingIpSuggestions.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpPrefixSequences**
> IpPrefixSequenceList GetIpPrefixSequences(ctx, optional)
Retrieves the generated sequences for an IP Prefix.

Get the sequences of network addresses which can be generated from an IP Prefix. \"startingPrefixIpAddress\", \"prefixLength\" and \"prefixCount\" filters are required. For Example: An IP Prefix with startingPrefixIpAddress 192.168.0.0 and prefixLength 30 and prefixCount 3 will result in generation of 3 sequences each with 4 IP addresses as: 192.168.0.0/30, 192.168.0.4/30 and 192.168.0.8/30 Filter example:<code>(filter=startingPrefixIpAddress==192.168.0.0;prefixLength==30;prefixCount==3)</code> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IpSpacesApiGetIpPrefixSequencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IpSpacesApiGetIpPrefixSequencesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter for a query.  FIQL format. | 

### Return type

[**IpPrefixSequenceList**](IpPrefixSequenceList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpPrefixSuggestions**
> IpSpaceIpPrefixSuggestions GetIpPrefixSuggestions(ctx, page, pageSize, optional)
Suggests IP Prefixes to use for network definitions.

Suggests IP Prefixes to use for network definitions. \"orgVdcId\"/\"vdcGroupId\" and \"networkType\" filters are required. Based on the specified Org VDC ID/vDC Group ID and network type, VCD will query all the applicable IP Spaces and suggest some IP prefixes which can be utilized for creation of network definition. IP Space IP prefixes which are are allocated but not currently used for any network definitions are considered. Allowed values for networkType filter are  ISOLATED and ROUTED. If the networkType is ROUTED, \"gatewayId\" filter must be specified. For ROUTED networks, VCD will query all the IP Spaces associated with the Edge Gateway. For ISOLATED networks, VCD will query all the available private IP Spaces. Results can also be filtered by IPV4 or IPV6 IP address types. Filter examples:<code>(filter=orgVdcId==URN;networkType==ROUTED;gatewayId==URN)</code>, <code>(filter=vdcGroupId==URN;networkType==ISOLATED)</code>, <code>(filter=orgVdcId==URN;networkType==ROUTED;gatewayId==URN;ipType==IPV4)</code> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***IpSpacesApiGetIpPrefixSuggestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IpSpacesApiGetIpPrefixSuggestionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**IpSpaceIpPrefixSuggestions**](IpSpaceIpPrefixSuggestions.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSpace**
> IpSpace GetIpSpace(ctx, ipSpaceId)
Get an IP Space.

Retrieves the specified IP Space. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceId** | **string**|  | 

### Return type

[**IpSpace**](IpSpace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSpaceIpAllocation**
> IpSpaceIpAllocation GetIpSpaceIpAllocation(ctx, ipSpaceId, ipAllocationId)
Get an IP Space IP allocation.

Retrieves the specified IP Space IP Allocation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceId** | **string**|  | 
  **ipAllocationId** | **string**|  | 

### Return type

[**IpSpaceIpAllocation**](IpSpaceIpAllocation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSpaceIpAllocations**
> IpSpaceIpAllocations GetIpSpaceIpAllocations(ctx, page, pageSize, ipSpaceId, optional)
Retrieves all the allocated IP addresses or IP Prefixes of an IP Space.

Retrieves all the allocated IP addresses or IP Prefixes of an IP Space. The allocation type is required to be specified in the filter. example: (type==FLOATING_IP). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **ipSpaceId** | **string**|  | 
 **optional** | ***IpSpacesApiGetIpSpaceIpAllocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IpSpacesApiGetIpSpaceIpAllocationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**IpSpaceIpAllocations**](IpSpaceIpAllocations.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpSpaceSummaries**
> IpSpaceSummaries GetIpSpaceSummaries(ctx, page, pageSize, optional)
Retrieves all the IP Space summaries.

Get all the IP space summaries in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***IpSpacesApiGetIpSpaceSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IpSpacesApiGetIpSpaceSummariesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**IpSpaceSummaries**](IpSpaceSummaries.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpSpace**
> UpdateIpSpace(ctx, ipSpaceConfig, ipSpaceId)
Update an IP Space.

Updates the specified IP Space. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceConfig** | [**IpSpace**](IpSpace.md)|  | 
  **ipSpaceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpSpaceIpAllocation**
> UpdateIpSpaceIpAllocation(ctx, ipSpaceIpAllocation, ipSpaceId, ipAllocationId)
Update an IP Space IP Allocation.

Updates the specified IP Space IP Allocation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSpaceIpAllocation** | [**IpSpaceIpAllocation**](IpSpaceIpAllocation.md)|  | 
  **ipSpaceId** | **string**|  | 
  **ipAllocationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

