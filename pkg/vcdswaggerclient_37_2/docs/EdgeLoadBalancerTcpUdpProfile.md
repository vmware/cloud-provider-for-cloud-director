# EdgeLoadBalancerTcpUdpProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the TCP/UDP profile. Name of the system defined monitors are prefixed with System for example &#39;System-TCP-Proxy&#39;. | [optional] [default to null]
**Type_** | **string** | This property specifies the different ways in which TCP/UDP packets are sent to the destination server. A value of \&quot;-\&quot; represents an unknown type. Such profile can still be removed from the Virtual Service on update but is not valid on create. &lt;ul&gt; &lt;li&gt;TCP_PROXY - TCP connection is terminated at the Virtual Service and a new TCP connection is made to the destination server. &lt;li&gt;TCP_FAST_PATH - TCP packets are directly forwarded to the destination server. &lt;li&gt;UDP_FAST_PATH - UDP packets are directly forwarded to the destination server. &lt;/ul&gt;  | [default to null]
**SystemDefined** | **bool** | Whether the TCP/UDP profile is system defined. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


