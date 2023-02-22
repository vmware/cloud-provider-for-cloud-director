# \PreferencesApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPreference**](PreferencesApi.md#GetPreference) | **Get** /1.0.0/preferences/{classifier} | Get specified user preference.
[**GetPreferenceDefinition**](PreferencesApi.md#GetPreferenceDefinition) | **Get** /1.0.0/definitions/preferences/{preferenceDefinitionId} | Get specified preference definition.
[**QueryPreferenceDefinitions**](PreferencesApi.md#QueryPreferenceDefinitions) | **Get** /1.0.0/definitions/preferences | Get list of preference definitions.
[**QueryPreferences**](PreferencesApi.md#QueryPreferences) | **Get** /1.0.0/preferences | Query user preferences.
[**UpdatePreference**](PreferencesApi.md#UpdatePreference) | **Put** /1.0.0/preferences/{classifier} | Update specified user preference


# **GetPreference**
> Preference GetPreference(ctx, classifier)
Get specified user preference.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **classifier** | **string**|  | 

### Return type

[**Preference**](Preference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPreferenceDefinition**
> PreferenceDefinition GetPreferenceDefinition(ctx, preferenceDefinitionId)
Get specified preference definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **preferenceDefinitionId** | **string**|  | 

### Return type

[**PreferenceDefinition**](PreferenceDefinition.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryPreferenceDefinitions**
> PreferenceDefinitions QueryPreferenceDefinitions(ctx, page, pageSize, optional)
Get list of preference definitions.

Get list of preference definitions 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***PreferencesApiQueryPreferenceDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreferencesApiQueryPreferenceDefinitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**PreferenceDefinitions**](PreferenceDefinitions.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryPreferences**
> Preferences QueryPreferences(ctx, optional)
Query user preferences.

Query user preferences 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PreferencesApiQueryPreferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreferencesApiQueryPreferencesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter for a query.  FIQL format. | 

### Return type

[**Preferences**](Preferences.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePreference**
> Preference UpdatePreference(ctx, updatePreferenceParams, classifier)
Update specified user preference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatePreferenceParams** | [**Preference**](Preference.md)|  | 
  **classifier** | **string**|  | 

### Return type

[**Preference**](Preference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

