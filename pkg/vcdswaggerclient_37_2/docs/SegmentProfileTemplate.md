# SegmentProfileTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The URN of the Segment Profile Template Object. | [optional] [default to null]
**Name** | **string** | The name of the Segment Profile Template. | [default to null]
**Description** | **string** | The description of the Segment Profile Template. | [optional] [default to null]
**SourceNsxTManagerRef** | [***EntityReference**](EntityReference.md) | The NSX-T manager providing the source segment profiles. | [default to null]
**QosProfile** | [***ExtObjectReference**](ExtObjectReference.md) | The Quality of Service (QoS) profile. | [optional] [default to null]
**MacDiscoveryProfile** | [***ExtObjectReference**](ExtObjectReference.md) | The MAC Discovery profile. Defines how the segment discovers MAC addresses. | [optional] [default to null]
**IpDiscoveryProfile** | [***ExtObjectReference**](ExtObjectReference.md) | The IP Discovery profile. Defines how the segment discovers IP addresses. | [optional] [default to null]
**SegmentSecurityProfile** | [***ExtObjectReference**](ExtObjectReference.md) | The Segment Security profile. Enables stateless L2 and L3 security on the segment. | [optional] [default to null]
**SpoofGuardProfile** | [***ExtObjectReference**](ExtObjectReference.md) | The Spoof Guard profile. | [optional] [default to null]
**LastModified** | [**time.Time**](time.Time.md) | The last time the segment profile template was modified. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


