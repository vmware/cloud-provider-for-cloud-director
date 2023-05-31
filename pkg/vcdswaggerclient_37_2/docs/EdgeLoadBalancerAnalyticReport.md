# EdgeLoadBalancerAnalyticReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | **string** | The gateway URN identier. | [optional] [default to null]
**ComponentId** | **string** | The virtual service or pool URN identifier. | [optional] [default to null]
**Metric** | **string** | The metric for this specific report.  Supported metrics can be determined by using the supported metrics API. | [optional] [default to null]
**Units** | **string** | The units for this specific report.  Units may be one of the following &lt;ul&gt; &lt;li&gt;PER_SECOND &lt;li&gt;METRIC_COUNT &lt;li&gt;BITS_PER_SECOND &lt;li&gt;BYTES_PER_SECOND &lt;li&gt;MILLISECONDS &lt;li&gt;PERCENT &lt;/ul&gt; Units are derived from the reported metric.  | [optional] [default to null]
**Statistics** | [***EdgeLoadBalancerMetricSummaryStats**](EdgeLoadBalancerMetricSummaryStats.md) |  | [optional] [default to null]
**Data** | [**[]EdgeLoadBalancerMetricData**](EdgeLoadBalancerMetricData.md) | The metric timeseries of data for this report. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


