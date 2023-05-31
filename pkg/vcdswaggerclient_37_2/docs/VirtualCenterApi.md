# \VirtualCenterApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVirtualCenter**](VirtualCenterApi.md#AttachVirtualCenter) | **Post** /1.0.0/virtualCenters | Attach a Virtual Center server
[**DeleteVirtualCenter**](VirtualCenterApi.md#DeleteVirtualCenter) | **Delete** /1.0.0/virtualCenters/{vcUrn} | Detach the specified Virtual Center server
[**GetAllResourcePools**](VirtualCenterApi.md#GetAllResourcePools) | **Get** /1.0.0/virtualCenters/{vcUrn}/resourcePools/browseAll | Get resource pools of a VC
[**GetChildResourcePools**](VirtualCenterApi.md#GetChildResourcePools) | **Get** /1.0.0/virtualCenters/{vcUrn}/resourcePools/browse/{moref} | Get resource pools of a VC
[**GetNsxVManagerSettings**](VirtualCenterApi.md#GetNsxVManagerSettings) | **Get** /1.0.0/virtualCenters/{vcUrn}/nsxVSettings | Get Virtual Center server network settings
[**GetResourcePoolKubernetesConfig**](VirtualCenterApi.md#GetResourcePoolKubernetesConfig) | **Get** /1.0.0/virtualCenters/{vcUrn}/resourcePools/{moref}/kubernetesConfig | Get kubernetes configuration for a resource pool.
[**GetRootResourcePools**](VirtualCenterApi.md#GetRootResourcePools) | **Get** /1.0.0/virtualCenters/{vcUrn}/resourcePools/browse | Get resource pools of a VC
[**GetStorageProfiles**](VirtualCenterApi.md#GetStorageProfiles) | **Get** /1.0.0/virtualCenters/{vcUrn}/storageProfiles | Get storage profiles of a VC
[**GetSupportedHardwareVersions**](VirtualCenterApi.md#GetSupportedHardwareVersions) | **Get** /1.0.0/virtualCenters/{vcUrn}/resourcePools/{moref}/hwv | Get supported hardware versions of a resource pool
[**GetUnmanagedVirtualMachines**](VirtualCenterApi.md#GetUnmanagedVirtualMachines) | **Get** /1.0.0/virtualCenters/{vcUrn}/unmanagedVirtualMachines | Get a list of unmanaged virtual machines from vCenter Server
[**GetVcStoragePolicyCapabilities**](VirtualCenterApi.md#GetVcStoragePolicyCapabilities) | **Get** /1.0.0/virtualCenters/{vcUrn}/storageProfiles/{moref}/capabilities | Retrieves capabilities of a specific Virtual Center storage policy.
[**GetVirtualCenter**](VirtualCenterApi.md#GetVirtualCenter) | **Get** /1.0.0/virtualCenters/{vcUrn} | Get Virtual Center server
[**GetVirtualCenterMetrics**](VirtualCenterApi.md#GetVirtualCenterMetrics) | **Get** /1.0.0/virtualCenters/{vcUrn}/metrics | Get Virtual Center server metrics
[**QueryVirtualCenters**](VirtualCenterApi.md#QueryVirtualCenters) | **Get** /1.0.0/virtualCenters | Gets a paged list of Virtual Center servers.
[**QueryVirtualMachineClasses**](VirtualCenterApi.md#QueryVirtualMachineClasses) | **Get** /1.0.0/virtualCenters/{vcUrn}/resourcePools/{moref}/virtualMachineClasses | Get a list of Virtual Machine Classes associated with this resource pool.
[**RetrieveVsphereVmca**](VirtualCenterApi.md#RetrieveVsphereVmca) | **Get** /1.0.0/virtualCenters/{vcUrn}/certificateAuthority/vmca | Retrieve the VMCA certificate
[**UpdateNsxVManagerSettings**](VirtualCenterApi.md#UpdateNsxVManagerSettings) | **Put** /1.0.0/virtualCenters/{vcUrn}/nsxVSettings | Update specified Virtual Center server network settings
[**UpdateVirtualCenter**](VirtualCenterApi.md#UpdateVirtualCenter) | **Put** /1.0.0/virtualCenters/{vcUrn} | Update specified Virtual Center server


# **AttachVirtualCenter**
> AttachVirtualCenter(ctx, vimserver)
Attach a Virtual Center server

Attach a Virtual Center server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vimserver** | [**VCenterServer**](VCenterServer.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVirtualCenter**
> DeleteVirtualCenter(ctx, vcUrn)
Detach the specified Virtual Center server

Unregister a vCenter server. This operation is asynchronous and returns a task that you can monitor to track the progress of the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllResourcePools**
> ResourcePools GetAllResourcePools(ctx, vcUrn, page, pageSize, optional)
Get resource pools of a VC

Get a list of all resource pools in the specified vCenter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VirtualCenterApiGetAllResourcePoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualCenterApiGetAllResourcePoolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**ResourcePools**](ResourcePools.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChildResourcePools**
> ResourcePools GetChildResourcePools(ctx, vcUrn, moref, page, pageSize)
Get resource pools of a VC

Get list of child resource pools of the specified parent that are eligible for consumption. If a resource pool is ineligible but is in the response, this means it has children which are eligible for consumption. A resource pool will be ineligible, unless the cluster has an ESXi host on it. The list will be sorted by name, case insensitive. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **moref** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]

### Return type

[**ResourcePools**](ResourcePools.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNsxVManagerSettings**
> NsxVManagerSettings GetNsxVManagerSettings(ctx, vcUrn)
Get Virtual Center server network settings

Retrieve the networking configuration of a registered vCenter server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 

### Return type

[**NsxVManagerSettings**](NsxVManagerSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourcePoolKubernetesConfig**
> ResourcePoolKubernetesConfig GetResourcePoolKubernetesConfig(ctx, vcUrn, moref)
Get kubernetes configuration for a resource pool.

Get kubernetes configuration for a resource pool which is backed by a kubernetes enabled cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **moref** | **string**|  | 

### Return type

[**ResourcePoolKubernetesConfig**](ResourcePoolKubernetesConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRootResourcePools**
> ResourcePools GetRootResourcePools(ctx, vcUrn, page, pageSize)
Get resource pools of a VC

Get a list of all root resource pools that are eligible for consumption. If a resource pool is ineligible but is in the response, this means it has children which are eligible for consumption. A resource pool will be ineligible, unless the cluster has an ESXi host on it. The list will be sorted by name, case insensitive. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]

### Return type

[**ResourcePools**](ResourcePools.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStorageProfiles**
> StorageProfiles GetStorageProfiles(ctx, vcUrn, page, pageSize, optional)
Get storage profiles of a VC

Get a list of all storage profiles for a VC. Supported contexts are: Resource Pool Moref (_context==moref) - | Returns all the storage profiles which are related to a specific Resoure Pool. Example: /cloudapi/{ver}/virtualCenters/{urn}/storageProfiles?filter=_context==resgroup-N 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VirtualCenterApiGetStorageProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualCenterApiGetStorageProfilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**StorageProfiles**](StorageProfiles.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedHardwareVersions**
> HardwareVersions GetSupportedHardwareVersions(ctx, vcUrn, moref)
Get supported hardware versions of a resource pool

Get a set of all supported hardware versions by the ESXi hosts in the resource pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **moref** | **string**|  | 

### Return type

[**HardwareVersions**](HardwareVersions.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnmanagedVirtualMachines**
> UnmanagedVirtualMachines GetUnmanagedVirtualMachines(ctx, vcUrn, page, pageSize, optional)
Get a list of unmanaged virtual machines from vCenter Server

Get a list of unmanaged virtual machines from vCenter Server 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VirtualCenterApiGetUnmanagedVirtualMachinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualCenterApiGetUnmanagedVirtualMachinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**UnmanagedVirtualMachines**](UnmanagedVirtualMachines.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVcStoragePolicyCapabilities**
> Capabilities GetVcStoragePolicyCapabilities(ctx, vcUrn, moref)
Retrieves capabilities of a specific Virtual Center storage policy.

Retrieves the current capabilities configured on a specific Virtual Center storage policy. These cannot be edited. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **moref** | **string**|  | 

### Return type

[**Capabilities**](Capabilities.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVirtualCenter**
> VCenterServer GetVirtualCenter(ctx, vcUrn)
Get Virtual Center server

Retrieve the representation of a vCenter server registered and managed by vCD. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 

### Return type

[**VCenterServer**](VCenterServer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVirtualCenterMetrics**
> VCenterServerMetrics GetVirtualCenterMetrics(ctx, vcUrn)
Get Virtual Center server metrics

Retrieve the metrics of a registered vCenter server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 

### Return type

[**VCenterServerMetrics**](VCenterServerMetrics.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryVirtualCenters**
> VCenterServers QueryVirtualCenters(ctx, page, pageSize, optional)
Gets a paged list of Virtual Center servers.

Retrieves a paged list of all Virtual Center servers in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VirtualCenterApiQueryVirtualCentersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualCenterApiQueryVirtualCentersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VCenterServers**](VCenterServers.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryVirtualMachineClasses**
> VirtualMachineClasses QueryVirtualMachineClasses(ctx, vcUrn, moref, page, pageSize, optional)
Get a list of Virtual Machine Classes associated with this resource pool.

Get a list of Virtual Machine Classes associated with this resource pool. This API throws 400 BadRequestException if called against a resource pool which is not Kubernetes enabled. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **moref** | **string**|  | 
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***VirtualCenterApiQueryVirtualMachineClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualCenterApiQueryVirtualMachineClassesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VirtualMachineClasses**](VirtualMachineClasses.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveVsphereVmca**
> VCenterVmca RetrieveVsphereVmca(ctx, vcUrn)
Retrieve the VMCA certificate

Within VSphere's Certificate management, the VMCA is a designated CA certificate that signs vsphere infrastructure endpoint certificates.<BR/>This API retrieves that certificate 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 

### Return type

[**VCenterVmca**](VCenterVMCA.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNsxVManagerSettings**
> UpdateNsxVManagerSettings(ctx, vcUrn, updateVCenterServerParams)
Update specified Virtual Center server network settings

Update the network settings of a registered vCenter server. This operation is asynchronous and returns a task that you can monitor to track the progress of the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **updateVCenterServerParams** | [**NsxVManagerSettings**](NsxVManagerSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVirtualCenter**
> UpdateVirtualCenter(ctx, vcUrn, updateVCenterServerParams)
Update specified Virtual Center server

Update the representation of a registered vCenter server. This operation is asynchronous and returns a task that you can monitor to track the progress of the request. Starting with API version 36.0, a null nsxVManager will attempt to remove the NSX-V Manager from the vCenter, and a non-null nsxVManager will attempt to add the NSX-V Manager to the vCenter if there is none registered or update the NSX-V Manager if there is one already registered to the vCenter. If you don't want to update the NSX-V Manager, provide the same NSX-V Manager settings as the existing one. For API versions before 36.0, no changes or updates to the nsxVManager will be made, regardless of whether an nsxVManger is provided or if it's different than the existing one. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vcUrn** | **string**|  | 
  **updateVCenterServerParams** | [**VCenterServer**](VCenterServer.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

