# SegmentMacDiscoveryProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the segment profile. | [optional] [default to null]
**DisplayName** | **string** | Name of the segment profile. This corresponds to the name used in NSX-T manager&#39;s logs or GUI. | [optional] [default to null]
**Description** | **string** | The description of the segment profile. | [optional] [default to null]
**NsxTManagerRef** | [***EntityReference**](EntityReference.md) | The NSX-T manager where this segment profile is configured. | [optional] [default to null]
**IsMacChangeEnabled** | **bool** | Whether source MAC address change is enabled. | [optional] [default to null]
**IsMacLearningEnabled** | **bool** | Whether source MAC address learning is enabled. | [optional] [default to null]
**MacLearningAgingTime** | **int32** | Aging time in seconds for learned MAC address. Indicates how long learned MAC address remain. | [optional] [default to null]
**MacLimit** | **int32** | The maximum number of MAC addresses that can be learned on this port. | [optional] [default to null]
**MacPolicy** | **string** | The policy after MAC Limit is exceeded. It can be either &#39;ALLOW&#39; or &#39;DROP&#39;. | [optional] [default to null]
**IsUnknownUnicastFloodingEnabled** | **bool** | Whether unknown unicast flooding rule is enabled. This allows flooding for unlearned MAC for ingress traffic. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


