# \TestConnectionApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Test**](TestConnectionApi.md#Test) | **Post** /1.0.0/testConnection | Test a connection


# **Test**
> TestResult Test(ctx, connection)
Test a connection

Tests a connection, including SSL handshake and hostname verification. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connection** | [**Connection**](Connection.md)|  | 

### Return type

[**TestResult**](TestResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=36.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

