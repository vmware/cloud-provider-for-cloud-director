# \EmailSettingsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestEmailSettings**](EmailSettingsApi.md#TestEmailSettings) | **Post** /1.0.0/smtp/test | Tests Email Settings and Connection


# **TestEmailSettings**
> TestEmailSettings(ctx, testEmailRequest)
Tests Email Settings and Connection

Tests that Email SMTP Settings are valid 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **testEmailRequest** | [**TestEmailRequest**](TestEmailRequest.md)| Email Settings and destination email address for Email Test call | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

