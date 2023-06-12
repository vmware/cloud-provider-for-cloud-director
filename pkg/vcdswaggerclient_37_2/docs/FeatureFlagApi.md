# \FeatureFlagApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeatureFlag**](FeatureFlagApi.md#GetFeatureFlag) | **Get** /1.0.0/featureFlags/{urn} | Retrieves a specific feature flag.
[**UpdateFeatureFlag**](FeatureFlagApi.md#UpdateFeatureFlag) | **Put** /1.0.0/featureFlags/{urn} | Updates a specific feature flag to either enable or disable it.


# **GetFeatureFlag**
> FeatureFlag GetFeatureFlag(ctx, urn)
Retrieves a specific feature flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urn** | **string**| URN containing the name of the feature flag such as (urn:vcloud:featureflag:&lt;feature flag name&gt;) | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeatureFlag**
> FeatureFlag UpdateFeatureFlag(ctx, featureFlag, urn)
Updates a specific feature flag to either enable or disable it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureFlag** | [**FeatureFlag**](FeatureFlag.md)|  | 
  **urn** | **string**| URN containing the name of the feature flag such as (urn:vcloud:featureflag:&lt;feature flag name&gt;) | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

