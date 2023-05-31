# \BrandingThemeApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertOldThemes**](BrandingThemeApi.md#ConvertOldThemes) | **Post** /1.0.0/brandingThemes/convert | Initiates converting for the old branding api themes
[**CreateBrandingTheme**](BrandingThemeApi.md#CreateBrandingTheme) | **Post** /1.0.0/brandingThemes | Creates a new branding theme
[**CreateBrandingThemeAssignment**](BrandingThemeApi.md#CreateBrandingThemeAssignment) | **Post** /1.0.0/brandingThemes/{brandingThemeId}/assignments | Set active theme for an organization
[**CreateBrandingThemeAssignments**](BrandingThemeApi.md#CreateBrandingThemeAssignments) | **Put** /1.0.0/brandingThemes/{brandingThemeId}/assignments | Set active theme for a list of organizations
[**DeleteBrandingTheme**](BrandingThemeApi.md#DeleteBrandingTheme) | **Delete** /1.0.0/brandingThemes/{brandingThemeId} | Deletes existing branding theme.
[**DeleteBrandingThemeAssignment**](BrandingThemeApi.md#DeleteBrandingThemeAssignment) | **Delete** /1.0.0/brandingThemes/{brandingThemeId}/assignments/{organizationId} | Remove active branding theme of tenant.
[**GetActiveBackgroundThemeResource**](BrandingThemeApi.md#GetActiveBackgroundThemeResource) | **Get** /1.0.0/brandingThemes/{brandingThemeId}/background | Gets the background resource for the active theme for the current organization.
[**GetActiveCSSThemeResource**](BrandingThemeApi.md#GetActiveCSSThemeResource) | **Get** /1.0.0/brandingThemes/{brandingThemeId}/css | Gets the CSS resource for the active theme for the current organization.
[**GetActiveCustomLinksThemeResource**](BrandingThemeApi.md#GetActiveCustomLinksThemeResource) | **Get** /1.0.0/brandingThemes/{brandingThemeId}/customLinks | Gets the custom links for the active theme for the current organization.
[**GetActiveFavIconThemeResource**](BrandingThemeApi.md#GetActiveFavIconThemeResource) | **Get** /1.0.0/brandingThemes/{brandingThemeId}/favIcon | Gets the favIcon resource for the active theme for the current organization.
[**GetActiveLogoThemeResource**](BrandingThemeApi.md#GetActiveLogoThemeResource) | **Get** /1.0.0/brandingThemes/{brandingThemeId}/logo | Gets the logo resource for the active theme for the current organization.
[**GetActivePortalNameThemeResource**](BrandingThemeApi.md#GetActivePortalNameThemeResource) | **Get** /1.0.0/brandingThemes/{brandingThemeId}/portalName | Gets the portal name for the active theme for the current organization.
[**GetBrandingTheme**](BrandingThemeApi.md#GetBrandingTheme) | **Get** /1.0.0/brandingThemes/{brandingThemeId} | Gets an existing branding theme
[**GetBrandingThemeAssignments**](BrandingThemeApi.md#GetBrandingThemeAssignments) | **Get** /1.0.0/brandingThemes/{brandingThemeId}/assignments | Gets organizations using this branding theme as active
[**GetBrandingThemes**](BrandingThemeApi.md#GetBrandingThemes) | **Get** /1.0.0/brandingThemes | Gets the list of all available branding themes
[**GetBrandingThemesAssignments**](BrandingThemeApi.md#GetBrandingThemesAssignments) | **Get** /1.0.0/brandingThemes/assignments | Get organizations and their active themes
[**GetThemeZip**](BrandingThemeApi.md#GetThemeZip) | **Get** /1.0.0/brandingThemes/{brandingThemeId}/resources | Gets theme resources as a zip archive.
[**UpdateBrandingTheme**](BrandingThemeApi.md#UpdateBrandingTheme) | **Put** /1.0.0/brandingThemes/{brandingThemeId} | Updates existing branding theme
[**UploadBrandingResource**](BrandingThemeApi.md#UploadBrandingResource) | **Put** /1.0.0/brandingThemes/{brandingThemeId}/resources | Sets resources for a theme


# **ConvertOldThemes**
> ConvertOldThemes(ctx, optional)
Initiates converting for the old branding api themes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BrandingThemeApiConvertOldThemesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandingThemeApiConvertOldThemesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assign** | **optional.Bool**| Convert and assign themes to tenants, also sets converted provider theme as default if needed | 
 **dryrun** | **optional.Bool**| Perform only a dry run, adding results in the job status | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBrandingTheme**
> BrandingTheme CreateBrandingTheme(ctx, brandingThemeObject)
Creates a new branding theme

Creates a new branding theme 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeObject** | [**BrandingTheme**](BrandingTheme.md)|  | 

### Return type

[**BrandingTheme**](BrandingTheme.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBrandingThemeAssignment**
> BrandingThemeAssignment CreateBrandingThemeAssignment(ctx, organizationReference, brandingThemeId, optional)
Set active theme for an organization

Set active theme for an organization 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationReference** | [**EntityReference**](EntityReference.md)|  | 
  **brandingThemeId** | **string**| theme URN | 
 **optional** | ***BrandingThemeApiCreateBrandingThemeAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandingThemeApiCreateBrandingThemeAssignmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| Indicates whether current organization assignments can be overwritten | [default to false]

### Return type

[**BrandingThemeAssignment**](BrandingThemeAssignment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBrandingThemeAssignments**
> CreateBrandingThemeAssignments(ctx, organizationReferences, brandingThemeId, optional)
Set active theme for a list of organizations

Set active theme for a list of organizations 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationReferences** | [**[]EntityReference**](EntityReference.md)|  | 
  **brandingThemeId** | **string**| theme URN | 
 **optional** | ***BrandingThemeApiCreateBrandingThemeAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandingThemeApiCreateBrandingThemeAssignmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| Indicates whether current organization assignments can be overwritten | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBrandingTheme**
> DeleteBrandingTheme(ctx, brandingThemeId, optional)
Deletes existing branding theme.

Deletes existing branding theme. If a theme is assigned to any organizations and forceDelete flag is specified it will delete the scoping as well. If a theme is assigned but forceDelete is not specified - error will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeId** | **string**| branding theme URN or &#39;active&#39; for the current organization active theme | 
 **optional** | ***BrandingThemeApiDeleteBrandingThemeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandingThemeApiDeleteBrandingThemeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Value \&quot;true\&quot; means to forcefully delete the object that contains other objects even if those objects are in a state that does not allow removal. The default is \&quot;false\&quot;; therefore, objects are not removed if they are not in a state that normally allows removal. Force also implies recursive delete where other contained objects are removed. Errors may be ignored. Invalid value (not true or false) are ignored.  | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBrandingThemeAssignment**
> DeleteBrandingThemeAssignment(ctx, brandingThemeId, organizationId)
Remove active branding theme of tenant.

Remove active branding theme of tenant. Tenant will inherit the default branding theme. The branding theme instance is not deleted and can be set as active again at later point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeId** | **string**| theme URN | 
  **organizationId** | **string**| organization URN | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveBackgroundThemeResource**
> string GetActiveBackgroundThemeResource(ctx, brandingThemeId)
Gets the background resource for the active theme for the current organization.

Gets the background resource for the active theme for the current organization. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeId** | **string**| branding theme URN or &#39;active&#39; for the current active theme | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png;version=37.2, image/jpeg;version=37.2, image/svg+xml;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveCSSThemeResource**
> string GetActiveCSSThemeResource(ctx, brandingThemeId)
Gets the CSS resource for the active theme for the current organization.

Gets the CSS resource for the active theme for the current organization. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeId** | **string**| branding theme URN or &#39;active&#39; for the current active theme | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/css;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveCustomLinksThemeResource**
> BrandingThemeResourceCustomLinks GetActiveCustomLinksThemeResource(ctx, page, pageSize, brandingThemeId, optional)
Gets the custom links for the active theme for the current organization.

Gets the custom links for the active theme for the current organization. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **brandingThemeId** | **string**| branding theme URN or &#39;active&#39; for the current active theme | 
 **optional** | ***BrandingThemeApiGetActiveCustomLinksThemeResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandingThemeApiGetActiveCustomLinksThemeResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**BrandingThemeResourceCustomLinks**](BrandingThemeResourceCustomLinks.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveFavIconThemeResource**
> string GetActiveFavIconThemeResource(ctx, brandingThemeId)
Gets the favIcon resource for the active theme for the current organization.

Gets the favIcon resource for the active theme for the current organization. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeId** | **string**| branding theme URN or &#39;active&#39; for the current active theme | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png;version=37.2, image/x-icon;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveLogoThemeResource**
> string GetActiveLogoThemeResource(ctx, brandingThemeId)
Gets the logo resource for the active theme for the current organization.

Gets the logo resource for the active theme for the current organization. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeId** | **string**| branding theme URN or &#39;active&#39; for the current active theme | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png;version=37.2, image/jpeg;version=37.2, image/svg+xml;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivePortalNameThemeResource**
> string GetActivePortalNameThemeResource(ctx, brandingThemeId)
Gets the portal name for the active theme for the current organization.

Gets the portal name for the active theme for the current organization. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeId** | **string**| branding theme URN or &#39;active&#39; for the current active theme | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrandingTheme**
> BrandingTheme GetBrandingTheme(ctx, brandingThemeId)
Gets an existing branding theme

Gets an existing branding theme 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeId** | **string**| branding theme URN or &#39;active&#39; for the current organization active theme | 

### Return type

[**BrandingTheme**](BrandingTheme.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrandingThemeAssignments**
> BrandingThemesAssignments GetBrandingThemeAssignments(ctx, page, pageSize, brandingThemeId, optional)
Gets organizations using this branding theme as active

Gets organizations using this branding theme as active 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **brandingThemeId** | **string**| theme URN | 
 **optional** | ***BrandingThemeApiGetBrandingThemeAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandingThemeApiGetBrandingThemeAssignmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**BrandingThemesAssignments**](BrandingThemesAssignments.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrandingThemes**
> BrandingThemes GetBrandingThemes(ctx, page, pageSize, optional)
Gets the list of all available branding themes

Gets the list of all available branding themes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***BrandingThemeApiGetBrandingThemesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandingThemeApiGetBrandingThemesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**BrandingThemes**](BrandingThemes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrandingThemesAssignments**
> BrandingThemesAssignments GetBrandingThemesAssignments(ctx, page, pageSize, optional)
Get organizations and their active themes

Get organizations and their active themes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***BrandingThemeApiGetBrandingThemesAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandingThemeApiGetBrandingThemesAssignmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**BrandingThemesAssignments**](BrandingThemesAssignments.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetThemeZip**
> string GetThemeZip(ctx, brandingThemeId)
Gets theme resources as a zip archive.

Gets theme resources as a zip archive. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeId** | **string**| theme URN | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBrandingTheme**
> BrandingTheme UpdateBrandingTheme(ctx, brandingThemeObject, brandingThemeId)
Updates existing branding theme

Updates existing branding theme 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeObject** | [**BrandingTheme**](BrandingTheme.md)|  | 
  **brandingThemeId** | **string**| branding theme URN or &#39;active&#39; for the current organization active theme | 

### Return type

[**BrandingTheme**](BrandingTheme.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBrandingResource**
> UploadBrandingResource(ctx, brandingThemeResourceUploadSpec, brandingThemeId)
Sets resources for a theme

Initiate an upload for the resources of this branding theme using the Transfer Service. A unique transfer service URL is returned where the plugin can be uploaded. The resources should be bundled in an archive, which must also contain a manfiest.json describing the various resources bundled in this archive. Example manifest.json {   Name: \"cusome_name.zip\",   Description: \"Autogenerated branding theme resources\",   Vendor: \"Vmware\",   Version: 0.0.1   Resources: {     FavIcon: <name of the favIcon file>,     Logo: <name of the logo file>,     Background: <name of the background file>,     Localizations: <name of the localization file>,     CustomLinks: <name of the custom links file>,     Styles: <name of the CSS file representing the overwritten styles>,     PortalName: <id_of_localization_or_just_text_to_be_displayed>   } } 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandingThemeResourceUploadSpec** | [**BrandingUploadSpec**](BrandingUploadSpec.md)|  | 
  **brandingThemeId** | **string**| theme URN | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

