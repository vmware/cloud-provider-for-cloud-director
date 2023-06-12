# \TrustedCertificatesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCertificate**](TrustedCertificatesApi.md#DeleteCertificate) | **Delete** /1.0.0/ssl/trustedCertificates/{trustedCertificate} | Revoke trusting specified certificate
[**GetCertificate**](TrustedCertificatesApi.md#GetCertificate) | **Get** /1.0.0/ssl/trustedCertificates/{trustedCertificate} | Get specified certificate
[**QueryTrustedCertificates**](TrustedCertificatesApi.md#QueryTrustedCertificates) | **Get** /1.0.0/ssl/trustedCertificates | Get currently trusted certificates
[**TrustCertificate**](TrustedCertificatesApi.md#TrustCertificate) | **Post** /1.0.0/ssl/trustedCertificates | Add to list of currently trusted certificates
[**UpdateCertificate**](TrustedCertificatesApi.md#UpdateCertificate) | **Put** /1.0.0/ssl/trustedCertificates/{trustedCertificate} | Updates an existing trusted certificate


# **DeleteCertificate**
> DeleteCertificate(ctx, trustedCertificate)
Revoke trusting specified certificate

Revoke trusting specified certificate 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trustedCertificate** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertificate**
> TrustedCertificate GetCertificate(ctx, trustedCertificate)
Get specified certificate

Get the PEM-encoded certificate with the requested URN 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trustedCertificate** | **string**|  | 

### Return type

[**TrustedCertificate**](TrustedCertificate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryTrustedCertificates**
> Certificates QueryTrustedCertificates(ctx, page, pageSize, optional)
Get currently trusted certificates

Get currently trusted certificates 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***TrustedCertificatesApiQueryTrustedCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrustedCertificatesApiQueryTrustedCertificatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Certificates**](Certificates.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrustCertificate**
> TrustedCertificate TrustCertificate(ctx, newCertificate)
Add to list of currently trusted certificates

Add to list of currently trusted certificates 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newCertificate** | [**TrustedCertificate**](TrustedCertificate.md)|  | 

### Return type

[**TrustedCertificate**](TrustedCertificate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCertificate**
> TrustedCertificate UpdateCertificate(ctx, modifiedCertificate, trustedCertificate)
Updates an existing trusted certificate

Updates an existing trusted certificate 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **modifiedCertificate** | [**TrustedCertificate**](TrustedCertificate.md)|  | 
  **trustedCertificate** | **string**|  | 

### Return type

[**TrustedCertificate**](TrustedCertificate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

