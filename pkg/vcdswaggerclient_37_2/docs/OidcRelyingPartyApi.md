# \OidcRelyingPartyApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOidcRelyingParty**](OidcRelyingPartyApi.md#DeleteOidcRelyingParty) | **Delete** /1.0.0/openIdProvider/relyingParties/{id} | Deletes a specific OIDC relying party.
[**GetOidcRelyingParty**](OidcRelyingPartyApi.md#GetOidcRelyingParty) | **Get** /1.0.0/openIdProvider/relyingParties/{id} | Retrieves a specific OIDC relying party.
[**QueryOidcRelyingParties**](OidcRelyingPartyApi.md#QueryOidcRelyingParties) | **Get** /1.0.0/openIdProvider/relyingParties | Retrieve OIDC relying parties.
[**RegenerateOidcRelyingPartySecret**](OidcRelyingPartyApi.md#RegenerateOidcRelyingPartySecret) | **Post** /1.0.0/openIdProvider/relyingParties/{id}/regenerateSecret | Regenerate the client secret of an OIDC relying party.
[**RegisterOidcRelyingParty**](OidcRelyingPartyApi.md#RegisterOidcRelyingParty) | **Post** /1.0.0/openIdProvider/relyingParties | Register an OIDC relying party.
[**UpdateOidcRelyingParty**](OidcRelyingPartyApi.md#UpdateOidcRelyingParty) | **Put** /1.0.0/openIdProvider/relyingParties/{id} | Updates a specific OIDC relying party.


# **DeleteOidcRelyingParty**
> DeleteOidcRelyingParty(ctx, id)
Deletes a specific OIDC relying party.

Delete a specific OIDC relying party. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| OIDC relying party&#39;s ID  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOidcRelyingParty**
> OidcRelyingParty GetOidcRelyingParty(ctx, id)
Retrieves a specific OIDC relying party.

Get a specific OIDC relying party. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| OIDC relying party&#39;s ID  | 

### Return type

[**OidcRelyingParty**](OidcRelyingParty.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryOidcRelyingParties**
> OidcRelyingParties QueryOidcRelyingParties(ctx, page, pageSize, optional)
Retrieve OIDC relying parties.

Get a list of all OIDC relying parties meeting the query parameters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***OidcRelyingPartyApiQueryOidcRelyingPartiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OidcRelyingPartyApiQueryOidcRelyingPartiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**OidcRelyingParties**](OidcRelyingParties.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateOidcRelyingPartySecret**
> OidcRelyingParty RegenerateOidcRelyingPartySecret(ctx, id)
Regenerate the client secret of an OIDC relying party.

Regenerates the client secret of an OIDC relying party. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| OIDC relying party&#39;s ID  | 

### Return type

[**OidcRelyingParty**](OidcRelyingParty.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterOidcRelyingParty**
> OidcRelyingParty RegisterOidcRelyingParty(ctx, relyingParty)
Register an OIDC relying party.

Registers a new OIDC relying party. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **relyingParty** | [**OidcRelyingParty**](OidcRelyingParty.md)|  | 

### Return type

[**OidcRelyingParty**](OidcRelyingParty.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOidcRelyingParty**
> OidcRelyingParty UpdateOidcRelyingParty(ctx, modifiedRelyingParty, id)
Updates a specific OIDC relying party.

Updates an OIDC relying party. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modifiedRelyingParty** | [**OidcRelyingParty**](OidcRelyingParty.md)|  | 
  **id** | **string**| OIDC relying party&#39;s ID  | 

### Return type

[**OidcRelyingParty**](OidcRelyingParty.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

