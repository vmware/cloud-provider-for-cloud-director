# QoSDscpConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrustMode** | **string** | It specifies DSCP trust mode. Values are below. &lt;ul&gt; &lt;li&gt;TRUSTED - With Trusted mode the inner header DSCP value is applied to the outer IP header for IP/IPv6 traffic. For non IP/IPv6 traffic, the outer IP header takes the default value. &lt;li&gt;UNTRUSTED - Untrusted mode is supported on overlay-based and VLAN-based logical port. &lt;/ul&gt;  | [optional] [default to null]
**Priority** | **int32** | The range of backing VLAN Id&#39;s. This information is available | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


