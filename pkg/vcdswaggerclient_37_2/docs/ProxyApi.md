# \ProxyApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProxy**](ProxyApi.md#CreateProxy) | **Post** /1.0.0/proxies | Creates a proxy.
[**DeleteProxy**](ProxyApi.md#DeleteProxy) | **Delete** /1.0.0/proxies/{id} | Delete a specific proxy. Will not delete an enabled proxy unless force is specified.
[**GetProxies**](ProxyApi.md#GetProxies) | **Get** /1.0.0/proxies | Gets a paged list of proxies for a tenant.
[**GetProxiesPacFileForTenant**](ProxyApi.md#GetProxiesPacFileForTenant) | **Get** /1.0.0/sddcProxiesPac/{id} | Gets the tenant-specific .pac file listing proxies accessible to the tenant.
[**GetProxy**](ProxyApi.md#GetProxy) | **Get** /1.0.0/proxies/{id} | Retrieves a specific proxy.
[**GetProxyCRL**](ProxyApi.md#GetProxyCRL) | **Get** /1.0.0/proxies/{id}/crl | Retrieve a proxy certificate revocation list in PEM format.
[**GetProxyCertificate**](ProxyApi.md#GetProxyCertificate) | **Get** /1.0.0/proxies/{id}/cert | Retrieve a proxy SSL certificate chain in PEM format.
[**GetProxyCertificateThumbprint**](ProxyApi.md#GetProxyCertificateThumbprint) | **Get** /1.0.0/proxies/{id}/thumbprint | Retrieve a Proxy SSL certificate thumbprint and algorithm used for calculation. Only SHA-256 is supported.
[**UpdateProxy**](ProxyApi.md#UpdateProxy) | **Put** /1.0.0/proxies/{id} | Update a specific proxy.
[**UpdateProxyCRL**](ProxyApi.md#UpdateProxyCRL) | **Put** /1.0.0/proxies/{id}/crl | Update a proxy certificate revocation list in PEM format.
[**UpdateProxyCertificate**](ProxyApi.md#UpdateProxyCertificate) | **Put** /1.0.0/proxies/{id}/cert | Update a proxy certificate chain in PEM format.


# **CreateProxy**
> CreateProxy(ctx, proxy)
Creates a proxy.

Creates a proxy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxy** | [**Proxy**](Proxy.md)| The new proxy model. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProxy**
> DeleteProxy(ctx, id, optional)
Delete a specific proxy. Will not delete an enabled proxy unless force is specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Proxy ID URN | 
 **optional** | ***ProxyApiDeleteProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiDeleteProxyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| If true, will delete proxy regardless of proxy state. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxies**
> Proxies GetProxies(ctx, page, pageSize, optional)
Gets a paged list of proxies for a tenant.

Gets a paged list of proxies for a tenant. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***ProxyApiGetProxiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProxyApiGetProxiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Proxies**](Proxies.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxiesPacFileForTenant**
> string GetProxiesPacFileForTenant(ctx, id)
Gets the tenant-specific .pac file listing proxies accessible to the tenant.

Gets the tenant-specific .pac file listing proxies accessible to the tenant. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier for the pac file configured for your organization | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-ns-proxy-autoconfig;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxy**
> Proxy GetProxy(ctx, id)
Retrieves a specific proxy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Proxy ID URN | 

### Return type

[**Proxy**](Proxy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxyCRL**
> string GetProxyCRL(ctx, id)
Retrieve a proxy certificate revocation list in PEM format.

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
 - **Accept**: application/x-pkcs7-crl;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxyCertificate**
> string GetProxyCertificate(ctx, id)
Retrieve a proxy SSL certificate chain in PEM format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Proxy ID URN | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-x509-ca-cert;version=37.2, application/x-pem-file;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProxyCertificateThumbprint**
> ThumbprintAndAlgorithm GetProxyCertificateThumbprint(ctx, id)
Retrieve a Proxy SSL certificate thumbprint and algorithm used for calculation. Only SHA-256 is supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Proxy ID URN | 

### Return type

[**ThumbprintAndAlgorithm**](ThumbprintAndAlgorithm.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProxy**
> Proxy UpdateProxy(ctx, updatedProxy, id)
Update a specific proxy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedProxy** | [**Proxy**](Proxy.md)| The updated proxy model. | 
  **id** | **string**| Proxy ID URN | 

### Return type

[**Proxy**](Proxy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProxyCRL**
> string UpdateProxyCRL(ctx, proxyCRL, id)
Update a proxy certificate revocation list in PEM format.

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
 - **Accept**: application/x-pkcs7-crl;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProxyCertificate**
> string UpdateProxyCertificate(ctx, proxyTrustAnchor, id)
Update a proxy certificate chain in PEM format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyTrustAnchor** | **string**| The updated proxy certificate chain in PEM format. Any extraneous whitespace or other information is removed.  | 
  **id** | **string**| Proxy ID URN | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/x-x509-ca-cert, application/x-pem-file
 - **Accept**: application/x-x509-ca-cert;version=37.2, application/x-pem-file;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

