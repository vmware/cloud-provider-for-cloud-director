# \EdgeGatewayLoadBalancerAnalyticsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoadBalancerAnalyticReports**](EdgeGatewayLoadBalancerAnalyticsApi.md#GetLoadBalancerAnalyticReports) | **Get** /1.0.0/loadBalancer/analyticReports | Retrieves analytics for a specific load balancer.
[**GetLoadBalancerSupportedAnalyticMetrics**](EdgeGatewayLoadBalancerAnalyticsApi.md#GetLoadBalancerSupportedAnalyticMetrics) | **Get** /1.0.0/loadBalancer/analyticReports/supportedMetrics | Retrieves all the supported metrics for load balancer analytic reports.


# **GetLoadBalancerAnalyticReports**
> EdgeLoadBalancerAnalyticReports GetLoadBalancerAnalyticReports(ctx, optional)
Retrieves analytics for a specific load balancer.

Retrieves analytics for a specific load balancer.  Metrics are specified in the filter query along with time period and series resolution.  Up to 5 metric series can be specified per report.  All reports will span the same time period.  Report filters are encapsulated in a fiql filter query parameter. Sample filter:   <code>filter=(componentId==urn:vcloud:virtualservice:7d38ad7f-cd93-4501-8c40-6f61650ccda0;         metric==l4_server.avg_total_rtt;metric==l7_server.avg_application_response_time;step==500;limit==100)</code> Supported filters are:   <ul>   <li>componentId.  The URN of the virtual service or pool for which metrics will be gathered.  Only one should be specified.   This is required.   <li>metric.  One or more metrics of interest.  <code>filter=(metric==l4_server.avg_total_rtt;metric==l7_server.avg_application_response_time)</code> -   This is required.  Supported metrics can be found at the analytics/supportedMetrics endpoint.   <li>step.  The time resolution of the report, in seconds.   This is required.  Minimum supported resolution is 300 seconds (5 minutes).   <li>limit.  Optional.  The number of data points to be returned.   This is optional.  Defaults to 59 where it can't be calculated.   <li>startTime.  Start time of the series.   This is optional.  Must be in ISO 8601 format (i.e. 2020-07-24T00:00:00).  If not provided, start time is calculated from the step and end time.   <li>endTime.  End period of the series.   This is optional.  Must be in ISO 8601 format (i.e. 2020-07-24T00:00:00). Defaults to the time of latest collected data point.   </ul> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EdgeGatewayLoadBalancerAnalyticsApiGetLoadBalancerAnalyticReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayLoadBalancerAnalyticsApiGetLoadBalancerAnalyticReportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter for a query.  FIQL format. | 

### Return type

[**EdgeLoadBalancerAnalyticReports**](EdgeLoadBalancerAnalyticReports.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerSupportedAnalyticMetrics**
> EdgeLoadBalancerAnalyticMetrics GetLoadBalancerSupportedAnalyticMetrics(ctx, optional)
Retrieves all the supported metrics for load balancer analytic reports.

Retrieves all the supported metrics for load balancer analytic reports.  These metrics can be used to create runtime reports of load balancer virtual services and pools. Supported filters are: <ul>   <li>componentId.  The URN of the load balancer virtual service or pool for which we want supported metrics. Only one should be specified.   This is required. </ul> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EdgeGatewayLoadBalancerAnalyticsApiGetLoadBalancerSupportedAnalyticMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeGatewayLoadBalancerAnalyticsApiGetLoadBalancerSupportedAnalyticMetricsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter for a query.  FIQL format. | 

### Return type

[**EdgeLoadBalancerAnalyticMetrics**](EdgeLoadBalancerAnalyticMetrics.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

