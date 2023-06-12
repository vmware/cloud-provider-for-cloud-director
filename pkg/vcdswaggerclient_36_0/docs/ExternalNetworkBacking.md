# ExternalNetworkBacking

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackingId** | **string** | Unique identifier for the network backing in NSX/vCenter. | [optional] [default to null]
**BackingType** | [***ExternalNetworkBackingType**](ExternalNetworkBackingType.md) | Type of network backing. This is deprecated in favor of string based enums going forward. | [optional] [default to null]
**BackingTypeValue** | **string** | Backing type for the network. &lt;ul&gt; &lt;li&gt;PORTGROUP&lt;/li&gt; &lt;li&gt;DV_PORTGROUP&lt;/li&gt; &lt;li&gt;NSXT_TIER0&lt;/li&gt; &lt;li&gt;NSXT_VRF_TIER0&lt;/li&gt; &lt;li&gt;IMPORTED_T_LOGICAL_SWITCH&lt;/li&gt; &lt;li&gt;UNKNOWN&lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]
**NetworkProvider** | [***EntityReference**](EntityReference.md) | The Network Provider for the backing, either a vCenter Server or NSX-T Manager. | [optional] [default to null]
**Name** | **string** | The name of this external network backing, if it exists. | [optional] [default to null]
**ParentTier0Ref** | [***ExtObjectReference**](ExtObjectReference.md) | If this external network is backed by a NSX-T VRF-Lite Tier0, then this field is set to be the parent Tier0 Router.  Otherwise, field is unset. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


