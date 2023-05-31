# EdgeLoadBalancerApplicationProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the application profile. Name of the system defined monitors are prefixed with System for example &#39;System-HTTP&#39;. | [optional] [default to null]
**Type_** | **string** | The profile type of application that this Virtual Service is configured with. A value of \&quot;-\&quot; represents an unknown type. Such profile can still be removed from the Virtual Service on update but is not valid on create. Values are below. &lt;ul&gt; &lt;li&gt;HTTP - Virtual Service supports HTTP protocol. &lt;li&gt;HTTPS - Virtual Service supports HTTPS protocol. &lt;li&gt;L4 - Virtual Service supports Layer 4 (Transport) using UDP/TCP protocol. &lt;li&gt;L4_TLS - Virtual Service supports Layer 4 (Transport) using UDP/TCP protocol with TLS. &lt;/ul&gt; L4_TLS is allowed only with additional licensing.  | [default to null]
**SystemDefined** | **bool** | Whether the Application Profile is system defined. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


