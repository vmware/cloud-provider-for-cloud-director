# \BrandingApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandingTheme**](BrandingApi.md#CreateBrandingTheme) | **Post** /branding/themes | Create a new custom theme
[**DeleteBrandingTheme**](BrandingApi.md#DeleteBrandingTheme) | **Delete** /branding/themes/{name} | Delete an existing custom theme
[**DeleteBrandingThemeContents**](BrandingApi.md#DeleteBrandingThemeContents) | **Delete** /branding/themes/{name}/contents | Delete an existing custom theme&#39;s contents
[**DeleteSystemIcon**](BrandingApi.md#DeleteSystemIcon) | **Delete** /branding/icon | Delete system level icon
[**DeleteSystemLogo**](BrandingApi.md#DeleteSystemLogo) | **Delete** /branding/logo | Delete system level logo
[**DeleteTenantIcon**](BrandingApi.md#DeleteTenantIcon) | **Delete** /branding/tenant/{org}/icon | Delete system level icon
[**DeleteTenantLogo**](BrandingApi.md#DeleteTenantLogo) | **Delete** /branding/tenant/{org}/logo | Delete the org-specific logo
[**GetBrandingTheme**](BrandingApi.md#GetBrandingTheme) | **Get** /branding/themes/{name} | Retrieve a specified theme identified by name
[**GetBrandingThemeCss**](BrandingApi.md#GetBrandingThemeCss) | **Get** /branding/themes/{name}/css | Retrieve the custom CSS for this theme, if any
[**GetBrandingThemes**](BrandingApi.md#GetBrandingThemes) | **Get** /branding/themes | Get a list of themes
[**GetSystemBranding**](BrandingApi.md#GetSystemBranding) | **Get** /branding | Gets the system level branding
[**GetSystemIcon**](BrandingApi.md#GetSystemIcon) | **Get** /branding/icon | Gets the system level browser icon
[**GetSystemLogo**](BrandingApi.md#GetSystemLogo) | **Get** /branding/logo | Gets the system level logo
[**GetTenantBranding**](BrandingApi.md#GetTenantBranding) | **Get** /branding/tenant/{org} | Gets org-specific branding
[**GetTenantIcon**](BrandingApi.md#GetTenantIcon) | **Get** /branding/tenant/{org}/icon | Gets the system level browser icon
[**GetTenantLogo**](BrandingApi.md#GetTenantLogo) | **Get** /branding/tenant/{org}/logo | Gets the org-specific logo
[**PutSystemBranding**](BrandingApi.md#PutSystemBranding) | **Put** /branding | Sets default branding
[**PutSystemIcon**](BrandingApi.md#PutSystemIcon) | **Put** /branding/icon | Sets the system level icon
[**PutSystemLogo**](BrandingApi.md#PutSystemLogo) | **Put** /branding/logo | Sets the system level logo
[**PutTenantBranding**](BrandingApi.md#PutTenantBranding) | **Put** /branding/tenant/{org} | Sets org-specific branding
[**PutTenantIcon**](BrandingApi.md#PutTenantIcon) | **Put** /branding/tenant/{org}/icon | Sets the system level icon
[**PutTenantLogo**](BrandingApi.md#PutTenantLogo) | **Put** /branding/tenant/{org}/logo | Sets the org-specific logo
[**RemoveTenantBranding**](BrandingApi.md#RemoveTenantBranding) | **Delete** /branding/tenant/{org} | Remove org-specific branding
[**UpdateBrandingTheme**](BrandingApi.md#UpdateBrandingTheme) | **Put** /branding/themes/{name} | Update an existing custom theme
[**UploadBrandingThemeContents**](BrandingApi.md#UploadBrandingThemeContents) | **Post** /branding/themes/{name}/contents | Upload the contents for this theme


# **CreateBrandingTheme**
> UiTheme CreateBrandingTheme(ctx, newTheme)
Create a new custom theme

Create a new custom theme, uniquely identified by name.  This can be used to set the theme in \"/branding\". 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newTheme** | [**UiTheme**](UiTheme.md)|  | 

### Return type

[**UiTheme**](UiTheme.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBrandingTheme**
> DeleteBrandingTheme(ctx, name)
Delete an existing custom theme

Delete a custom theme 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBrandingThemeContents**
> DeleteBrandingThemeContents(ctx, name)
Delete an existing custom theme's contents

Delete a custom theme's contents 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSystemIcon**
> DeleteSystemIcon(ctx, )
Delete system level icon

Delete the system level icon, forcing the get method to return the vCloud Director default icon. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSystemLogo**
> DeleteSystemLogo(ctx, )
Delete system level logo

Delete the system level logo, forcing the get method to return the vCloud Director default logo. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTenantIcon**
> DeleteTenantIcon(ctx, org)
Delete system level icon

Delete the system level icon, forcing the get method to return the vCloud Director default icon. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization for whom branding is being set | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTenantLogo**
> DeleteTenantLogo(ctx, org)
Delete the org-specific logo

Delete the org-specific logo, forcing the get method to return the system default logo. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization for whom branding is being set | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrandingTheme**
> UiTheme GetBrandingTheme(ctx, name)
Retrieve a specified theme identified by name

Update an existing theme, uniquely identified by name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**UiTheme**](UiTheme.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrandingThemeCss**
> string GetBrandingThemeCss(ctx, name)
Retrieve the custom CSS for this theme, if any

Retrieve the CSS for this theme if it has been set 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/css;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrandingThemes**
> []UiTheme GetBrandingThemes(ctx, )
Get a list of themes

Get a list of supported themes, uniquely identified by their names.  This can be used to set the theme in \"/branding\". 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UiTheme**](UiTheme.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemBranding**
> UiBranding GetSystemBranding(ctx, )
Gets the system level branding

Get the system level branding information including the portal name, portal color, selected theme and custom URLs. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UiBranding**](UiBranding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemIcon**
> string GetSystemIcon(ctx, )
Gets the system level browser icon

Get the system level icon as raw image data suitable for use in an image tag's src attribute.  If a custom icon is not set then the vCloud Director default icon is sent. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png;version=36.0, image/x-icon;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemLogo**
> string GetSystemLogo(ctx, )
Gets the system level logo

Get the system level logo as raw image data suitable for use in an image tag's src attribute.  If a custom logo is not set then the vCloud Director default logo is sent. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png;version=36.0, image/jpeg;version=36.0, image/svg+xml;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTenantBranding**
> UiBranding GetTenantBranding(ctx, org)
Gets org-specific branding

Get org-specific branding information including the portal name, portal color, selected theme and custom URLs. If no org branding has been specified, retrieve the default system branding. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization for whom branding is being set | 

### Return type

[**UiBranding**](UiBranding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTenantIcon**
> string GetTenantIcon(ctx, org)
Gets the system level browser icon

Get the system level icon as raw image data suitable for use in an image tag's src attribute.  If a custom icon is not set then the vCloud Director default icon is sent. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization for whom branding is being set | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png;version=36.0, image/x-icon;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTenantLogo**
> string GetTenantLogo(ctx, org)
Gets the org-specific logo

Get the org-specific logo as raw image data suitable for use in an image tag's src attribute. If an org-specific logo is not set, get the default system logo.  If a custom logo is not set then the vCloud Director default logo is sent. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization for whom branding is being set | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png;version=36.0, image/jpeg;version=36.0, image/svg+xml;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSystemBranding**
> UiBranding PutSystemBranding(ctx, body)
Sets default branding

Sets the branding information including the portal name, portal color, selected theme and custom URLs for a specific org or system default. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UiBranding**](UiBranding.md)|  | 

### Return type

[**UiBranding**](UiBranding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSystemIcon**
> PutSystemIcon(ctx, contentType, body)
Sets the system level icon

Set the system icon data. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | [**interface{}**](.md)|  | 
  **body** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: image/png, image/x-icon
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSystemLogo**
> PutSystemLogo(ctx, contentType, body)
Sets the system level logo

Set the system logo data. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | [**interface{}**](.md)|  | 
  **body** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: image/png, image/jpeg, image/svg+xml
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTenantBranding**
> UiBranding PutTenantBranding(ctx, body, org)
Sets org-specific branding

Sets the branding information including the portal name, portal color, selected theme and custom URLs for a specific org. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UiBranding**](UiBranding.md)|  | 
  **org** | **string**| Organization for whom branding is being set | 

### Return type

[**UiBranding**](UiBranding.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTenantIcon**
> PutTenantIcon(ctx, contentType, body, org)
Sets the system level icon

Set the system icon data. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | [**interface{}**](.md)|  | 
  **body** | **string**|  | 
  **org** | **string**| Organization for whom branding is being set | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: image/png, image/x-icon
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTenantLogo**
> PutTenantLogo(ctx, contentType, body, org)
Sets the org-specific logo

Set the org-specific data. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | [**interface{}**](.md)|  | 
  **body** | **string**|  | 
  **org** | **string**| Organization for whom branding is being set | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: image/png, image/jpeg, image/svg+xml
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTenantBranding**
> RemoveTenantBranding(ctx, org)
Remove org-specific branding

Removes org-specific branding information if specified 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization for whom branding is being set | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBrandingTheme**
> UiTheme UpdateBrandingTheme(ctx, updatedTheme, name)
Update an existing custom theme

Update an existing custom theme, uniquely identified by name.  This can be used to set the theme in \"/branding\". 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatedTheme** | [**UiTheme**](UiTheme.md)|  | 
  **name** | **string**|  | 

### Return type

[**UiTheme**](UiTheme.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBrandingThemeContents**
> UploadBrandingThemeContents(ctx, pluginUploadSpec, name)
Upload the contents for this theme

Update an existing custom theme's contents, uniquely identified by name. Currently, this is limited to a single CSS file. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginUploadSpec** | [**UploadSpec**](UploadSpec.md)|  | 
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

