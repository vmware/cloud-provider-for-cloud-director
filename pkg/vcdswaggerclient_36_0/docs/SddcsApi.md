# \SddcsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSddc**](SddcsApi.md#CreateSddc) | **Post** /1.0.0/sddcs | Create a Software-Defined Datacenter.
[**CreateSddcEndpoint**](SddcsApi.md#CreateSddcEndpoint) | **Post** /1.0.0/endpoints | Creates an SDDC endpoint.
[**DeleteSddc**](SddcsApi.md#DeleteSddc) | **Delete** /1.0.0/sddcs/{id} | Delete a specific Software-Defined Datacenter. Unless force is specified, SDDC &amp; its proxies must be disabled before they can be deleted. 
[**DeleteSddcEndpoint**](SddcsApi.md#DeleteSddcEndpoint) | **Delete** /1.0.0/endpoints/{id} | Delete a specific SDDC endpoint. Will not delete a default endpoint.
[**GetEndpointsForSddc**](SddcsApi.md#GetEndpointsForSddc) | **Get** /1.0.0/endpoints | Retrieve the endpoints for the SDDC.
[**GetProxiesForSddc**](SddcsApi.md#GetProxiesForSddc) | **Get** /1.0.0/sddcs/{id}/proxies | Retrieve the proxies for the Software-Defined Datacenter.
[**GetSddc**](SddcsApi.md#GetSddc) | **Get** /1.0.0/sddcs/{id} | Retrieve a specific Software-Defined Datacenter.
[**GetSddcEndpoint**](SddcsApi.md#GetSddcEndpoint) | **Get** /1.0.0/endpoints/{id} | Retrieves a specific SDDC endpoint.
[**GetSddcOwner**](SddcsApi.md#GetSddcOwner) | **Get** /1.0.0/sddcs/{id}/owner | Retrieve the owner of the Software-Defined Datacenter.
[**GetSddcs**](SddcsApi.md#GetSddcs) | **Get** /1.0.0/sddcs | Get the list of Software-Defined Datacenters accessible to the user.
[**QuerySddcTenants**](SddcsApi.md#QuerySddcTenants) | **Get** /1.0.0/sddcs/{id}/tenants | Retrieve the list of tenants a Software-Defined Datacenter is published to.
[**SddcPublishToTenants**](SddcsApi.md#SddcPublishToTenants) | **Post** /1.0.0/sddcs/{id}/tenants/publish | Publish a Software-Defined Datacenter to the given tenants.
[**SddcUnpublishFromTenants**](SddcsApi.md#SddcUnpublishFromTenants) | **Post** /1.0.0/sddcs/{id}/tenants/unpublish | Revoke publication of the Software-Defined Datacenter for the tenants.
[**SetSddcOwner**](SddcsApi.md#SetSddcOwner) | **Put** /1.0.0/sddcs/{id}/owner | Update the owner of the Software-Defined Datacenter.
[**SetSddcTenants**](SddcsApi.md#SetSddcTenants) | **Put** /1.0.0/sddcs/{id}/tenants | Reset the list of tenants a Software-Defined Datacenter is published to.
[**UpdateSddc**](SddcsApi.md#UpdateSddc) | **Put** /1.0.0/sddcs/{id} | Update a specific Software-Defined Datacenter.
[**UpdateSddcEndpoint**](SddcsApi.md#UpdateSddcEndpoint) | **Put** /1.0.0/endpoints/{id} | Update a specific SDDC endpoint.


# **CreateSddc**
> CreateSddc(ctx, newSddc)
Create a Software-Defined Datacenter.

Create a Software-Defined Datacenter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newSddc** | [**Sddc**](Sddc.md)| The new SDDC model. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSddcEndpoint**
> SddcEndpoint CreateSddcEndpoint(ctx, endpoint)
Creates an SDDC endpoint.

Creates an SDDC endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpoint** | [**SddcEndpoint**](SddcEndpoint.md)| The new SDDC endpoint model. | 

### Return type

[**SddcEndpoint**](SddcEndpoint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSddc**
> DeleteSddc(ctx, id, optional)
Delete a specific Software-Defined Datacenter. Unless force is specified, SDDC & its proxies must be disabled before they can be deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SDDC ID URN | 
 **optional** | ***SddcsApiDeleteSddcOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcsApiDeleteSddcOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| If true, will delete SDDC regardless of the state of the SDDC or any of its proxies.  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSddcEndpoint**
> DeleteSddcEndpoint(ctx, id)
Delete a specific SDDC endpoint. Will not delete a default endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SDDC Endpoint ID URN | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEndpointsForSddc**
> SddcEndpoints GetEndpointsForSddc(ctx, page, pageSize, optional)
Retrieve the endpoints for the SDDC.

Retrieve the endpoints for the SDDC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***SddcsApiGetEndpointsForSddcOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcsApiGetEndpointsForSddcOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**SddcEndpoints**](SddcEndpoints.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxiesForSddc**
> EntityReferences GetProxiesForSddc(ctx, page, pageSize, id, optional)
Retrieve the proxies for the Software-Defined Datacenter.

Retrieve the proxies for the Software-Defined Datacenter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**| SDDC ID URN | 
 **optional** | ***SddcsApiGetProxiesForSddcOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcsApiGetProxiesForSddcOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSddc**
> Sddc GetSddc(ctx, id)
Retrieve a specific Software-Defined Datacenter.

Retrieve a specific Software-Defined Datacenter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SDDC ID URN | 

### Return type

[**Sddc**](Sddc.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSddcEndpoint**
> SddcEndpoint GetSddcEndpoint(ctx, id)
Retrieves a specific SDDC endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SDDC Endpoint ID URN | 

### Return type

[**SddcEndpoint**](SddcEndpoint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSddcOwner**
> EntityReference GetSddcOwner(ctx, id)
Retrieve the owner of the Software-Defined Datacenter.

Retrieve the owner of the Software-Defined Datacenter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SDDC ID URN | 

### Return type

[**EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSddcs**
> Sddcs GetSddcs(ctx, page, pageSize, optional)
Get the list of Software-Defined Datacenters accessible to the user.

Get the list of Software-Defined Datacenters accessible to the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***SddcsApiGetSddcsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcsApiGetSddcsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Sddcs**](Sddcs.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuerySddcTenants**
> EntityReferences QuerySddcTenants(ctx, page, pageSize, id, optional)
Retrieve the list of tenants a Software-Defined Datacenter is published to.

Retrieve the list of tenants a Software-Defined Datacenter is published to. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**| SDDC ID URN | 
 **optional** | ***SddcsApiQuerySddcTenantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcsApiQuerySddcTenantsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SddcPublishToTenants**
> EntityReferences SddcPublishToTenants(ctx, publishTenantsBody, id)
Publish a Software-Defined Datacenter to the given tenants.

Publish a Software-Defined Datacenter to the tenants. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publishTenantsBody** | [**[]EntityReference**](EntityReference.md)| The list of tenant EntityReferences that a Software-Defined Datacenter should be published to. | 
  **id** | **string**| SDDC ID URN | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SddcUnpublishFromTenants**
> EntityReferences SddcUnpublishFromTenants(ctx, unpublishTenantsBody, id)
Revoke publication of the Software-Defined Datacenter for the tenants.

Revoke publication of the Software-Defined Datacenter for the tenants. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unpublishTenantsBody** | [**[]EntityReference**](EntityReference.md)| The list of tenant EntityReferences that a Software-Defined Datacenter should be unpublished from. | 
  **id** | **string**| SDDC ID URN | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSddcOwner**
> EntityReference SetSddcOwner(ctx, newOwner, id)
Update the owner of the Software-Defined Datacenter.

Update the owner of the Software-Defined Datacenter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newOwner** | [**EntityReference**](EntityReference.md)| The EntityReference to the owner of the SDDC. | 
  **id** | **string**| SDDC ID URN | 

### Return type

[**EntityReference**](EntityReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSddcTenants**
> EntityReferences SetSddcTenants(ctx, publishTenantsBody, id)
Reset the list of tenants a Software-Defined Datacenter is published to.

Reset the list of tenants a Software-Defined Datacenter is published to. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publishTenantsBody** | [**[]EntityReference**](EntityReference.md)| The list of tenant EntityReferences that a Software-Defined Datacenter should be published to. | 
  **id** | **string**| SDDC ID URN | 

### Return type

[**EntityReferences**](EntityReferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSddc**
> Sddc UpdateSddc(ctx, updatedSddc, id)
Update a specific Software-Defined Datacenter.

Update a specific Software-Defined Datacenter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedSddc** | [**Sddc**](Sddc.md)| The updated SDDC model. | 
  **id** | **string**| SDDC ID URN | 

### Return type

[**Sddc**](Sddc.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSddcEndpoint**
> SddcEndpoint UpdateSddcEndpoint(ctx, updatedSddcEndpoint, id)
Update a specific SDDC endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedSddcEndpoint** | [**SddcEndpoint**](SddcEndpoint.md)| The updated SDDC endpoint model. | 
  **id** | **string**| SDDC Endpoint ID URN | 

### Return type

[**SddcEndpoint**](SddcEndpoint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

