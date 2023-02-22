# \SddcProxyApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSddcProxy**](SddcProxyApi.md#CreateSddcProxy) | **Post** /1.0.0/sddcProxies | Creates an SDDC proxy.
[**DeleteSddcProxy**](SddcProxyApi.md#DeleteSddcProxy) | **Delete** /1.0.0/sddcProxies/{id} | Delete a specific SDDC proxy. Will not delete an enabled proxy unless force is specified.
[**GetSddcProxies**](SddcProxyApi.md#GetSddcProxies) | **Get** /1.0.0/sddcProxies | Gets a paged list of SDDC proxies for a tenant.
[**GetSddcProxiesPacFile**](SddcProxyApi.md#GetSddcProxiesPacFile) | **Get** /1.0.0/sddcProxiesPac | Gets the .pac file for the user&#39;s accessible proxies.
[**GetSddcProxy**](SddcProxyApi.md#GetSddcProxy) | **Get** /1.0.0/sddcProxies/{id} | Retrieves a specific SDDC proxy.
[**GetSddcProxyCRL**](SddcProxyApi.md#GetSddcProxyCRL) | **Get** /1.0.0/sddcProxies/{id}/crl | Retrieve a SDDC proxy certificate revocation list in PEM format.
[**GetSddcProxyCertificate**](SddcProxyApi.md#GetSddcProxyCertificate) | **Get** /1.0.0/sddcProxies/{id}/cert | Retrieve a SDDC proxy SSL certificate chain in PEM format.
[**GetSddcProxyCertificateThumbprint**](SddcProxyApi.md#GetSddcProxyCertificateThumbprint) | **Get** /1.0.0/sddcProxies/{id}/thumbprint | Retrieve a SDDC Proxy SSL certificate thumbprint. The thumbprint is the SHA-1 hash of the DER encoding of the certificate.
[**UpdateSddcProxy**](SddcProxyApi.md#UpdateSddcProxy) | **Put** /1.0.0/sddcProxies/{id} | Update a specific SDDC proxy.
[**UpdateSddcProxyCRL**](SddcProxyApi.md#UpdateSddcProxyCRL) | **Put** /1.0.0/sddcProxies/{id}/crl | Update a SDDC proxy certificate revocation list in PEM format.
[**UpdateSddcProxyCertificate**](SddcProxyApi.md#UpdateSddcProxyCertificate) | **Put** /1.0.0/sddcProxies/{id}/cert | Update a SDDC proxy certificate chain in PEM format.


# **CreateSddcProxy**
> CreateSddcProxy(ctx, proxy)
Creates an SDDC proxy.

Creates an SDDC proxy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxy** | [**SddcProxy**](SddcProxy.md)| The new SDDC proxy model. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSddcProxy**
> DeleteSddcProxy(ctx, id, optional)
Delete a specific SDDC proxy. Will not delete an enabled proxy unless force is specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SDDC Proxy ID URN | 
 **optional** | ***SddcProxyApiDeleteSddcProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcProxyApiDeleteSddcProxyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| If true, will delete proxy regardless of proxy state. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSddcProxies**
> SddcProxies GetSddcProxies(ctx, page, pageSize, optional)
Gets a paged list of SDDC proxies for a tenant.

Gets a paged list of SDDC proxies for a tenant. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***SddcProxyApiGetSddcProxiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SddcProxyApiGetSddcProxiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**SddcProxies**](SddcProxies.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSddcProxiesPacFile**
> string GetSddcProxiesPacFile(ctx, )
Gets the .pac file for the user's accessible proxies.

Gets the .pac file for the user's accessible proxies. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-ns-proxy-autoconfig;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSddcProxy**
> SddcProxy GetSddcProxy(ctx, id)
Retrieves a specific SDDC proxy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SDDC Proxy ID URN | 

### Return type

[**SddcProxy**](SddcProxy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSddcProxyCRL**
> string GetSddcProxyCRL(ctx, id)
Retrieve a SDDC proxy certificate revocation list in PEM format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-pkcs7-crl;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSddcProxyCertificate**
> string GetSddcProxyCertificate(ctx, id)
Retrieve a SDDC proxy SSL certificate chain in PEM format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SDDC Proxy ID URN | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-x509-ca-cert;version=36.0, application/x-pem-file;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSddcProxyCertificateThumbprint**
> string GetSddcProxyCertificateThumbprint(ctx, id)
Retrieve a SDDC Proxy SSL certificate thumbprint. The thumbprint is the SHA-1 hash of the DER encoding of the certificate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SDDC Proxy ID URN | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSddcProxy**
> SddcProxy UpdateSddcProxy(ctx, updatedSddcProxy, id)
Update a specific SDDC proxy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedSddcProxy** | [**SddcProxy**](SddcProxy.md)| The updated SDDC proxy model. | 
  **id** | **string**| SDDC Proxy ID URN | 

### Return type

[**SddcProxy**](SddcProxy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSddcProxyCRL**
> string UpdateSddcProxyCRL(ctx, proxyCRL, id)
Update a SDDC proxy certificate revocation list in PEM format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyCRL** | **string**|  | 
  **id** | **string**|  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/x-pkcs7-crl
 - **Accept**: application/x-pkcs7-crl;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSddcProxyCertificate**
> string UpdateSddcProxyCertificate(ctx, proxyTrustAnchor, id)
Update a SDDC proxy certificate chain in PEM format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyTrustAnchor** | **string**| The updated SDDC proxy certificate chain in PEM format. | 
  **id** | **string**| SDDC Proxy ID URN | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/x-x509-ca-cert, application/x-pem-file
 - **Accept**: application/x-x509-ca-cert;version=36.0, application/x-pem-file;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

