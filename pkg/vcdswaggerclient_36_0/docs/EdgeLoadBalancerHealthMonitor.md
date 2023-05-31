# EdgeLoadBalancerHealthMonitor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Health monitor. Name of the system defined monitors are prefixed with System for example &#39;System-HTTP&#39;. | [optional] [default to null]
**Type_** | **string** | Type of the health monitors. A value of \&quot;-\&quot; represents an unknown type. Such monitor can still be removed from the Load Balancer Pool on update but is not valid on create. Supported values are: &lt;ul&gt; &lt;li&gt;HTTP - HTTP request/response is used to validate health. &lt;li&gt;HTTPS - Used against HTTPS encrypted web servers to validate health. &lt;li&gt;TCP - TCP connection is used to validate health. &lt;li&gt;UDP - A UDP datagram is used to validate health. &lt;li&gt;PING - An ICMP ping is used to validate health. &lt;/ul&gt;  | [default to null]
**SystemDefined** | **bool** | Whether the Health Monitor is system defined. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


