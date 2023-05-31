# \VgpuProfilesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVgpuProfile**](VgpuProfilesApi.md#DeleteVgpuProfile) | **Delete** /1.0.0/vgpuProfiles/{vgpuProfileId:((?!consumers|commonVgpuProfile|vgpuProfileConsumer).)*} | Delete specified vGPU profile.
[**GetVgpuProfile**](VgpuProfilesApi.md#GetVgpuProfile) | **Get** /1.0.0/vgpuProfiles/{vgpuProfileId:((?!consumers|commonVgpuProfile|vgpuProfileConsumer).)*} | Get specified vGPU profile
[**QueryVgpuProfiles**](VgpuProfilesApi.md#QueryVgpuProfiles) | **Get** /1.0.0/vgpuProfiles | Get list of vGPU profiles
[**UpdateVgpuProfile**](VgpuProfilesApi.md#UpdateVgpuProfile) | **Put** /1.0.0/vgpuProfiles/{vgpuProfileId:((?!consumers|commonVgpuProfile|vgpuProfileConsumer).)*} | Update specified vGPU profile


# **DeleteVgpuProfile**
> DeleteVgpuProfile(ctx, vgpuProfileId)
Delete specified vGPU profile.

Deletes specified vGPU profile. A vGPU profile is eligible for deletion only when it is no longer available in any vCD managed Provider vDC cluster and is not in use by any active vGPU policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vgpuProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVgpuProfile**
> VgpuProfile GetVgpuProfile(ctx, vgpuProfileId)
Get specified vGPU profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vgpuProfileId** | **string**|  | 

### Return type

[**VgpuProfile**](VgpuProfile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryVgpuProfiles**
> VgpuProfiles QueryVgpuProfiles(ctx, page, pageSize, optional)
Get list of vGPU profiles

Get list of vGPU profiles available in the system. <br> Results can be filtered by: <ul>   <li> name   <li> tenantFacingName   <li> pvdcId - |     The URN of provider vDC.     Filters all the vGPU profiles that are accessible to the supplied provider vDC.     Example: (pvdcId==urn:vcloud:providervdc:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)   </li> <ul> <br> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VgpuProfilesApiQueryVgpuProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VgpuProfilesApiQueryVgpuProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VgpuProfiles**](VgpuProfiles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVgpuProfile**
> VgpuProfile UpdateVgpuProfile(ctx, updatedVgpuProfile, vgpuProfileId)
Update specified vGPU profile

vGPU profile's name cannot be updated but other information such as tenantFacingName and instructions can be updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedVgpuProfile** | [**VgpuProfile**](VgpuProfile.md)|  | 
  **vgpuProfileId** | **string**|  | 

### Return type

[**VgpuProfile**](VgpuProfile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

