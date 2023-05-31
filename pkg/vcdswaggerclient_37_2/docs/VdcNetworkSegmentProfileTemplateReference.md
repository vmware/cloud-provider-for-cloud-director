# VdcNetworkSegmentProfileTemplateReference

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateRef** | [***EntityReference**](EntityReference.md) | Reference to the Segment Profile Template that is used when creating/updating this network. This reference will be returned if the original template still exists and all of the segment profiles on the network match exactly with the profiles configured on the template.  | [optional] [default to null]
**Source** | **string** | Where the referenced template is sourced from. &lt;ul&gt; &lt;li&gt;The possible values are GLOBAL, ORG_VDC and NETWORK&lt;/li&gt; &lt;li&gt;GLOBAL - The template comes from the global default for the VCD site.&lt;/li&gt; &lt;li&gt;ORG_VDC - The template comes from the default for the Org vDC that owns this network.&lt;/li&gt; &lt;li&gt;NETWORK - The template was directly specified on the network during creation or update.&lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


