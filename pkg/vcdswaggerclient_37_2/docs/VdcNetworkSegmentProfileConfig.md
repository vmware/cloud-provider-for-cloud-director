# VdcNetworkSegmentProfileConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentProfileTemplate** | [***VdcNetworkSegmentProfileTemplateReference**](VdcNetworkSegmentProfileTemplateReference.md) | Reference to the Segment Profile Template that is used when creating/updating this network. This reference will be returned if the original template still exists and all of the segment profiles on the network match exactly with the profiles configured on the template.  | [optional] [default to null]
**QosProfile** | [***ExtObjectReference**](ExtObjectReference.md) | The Quality of Service (QoS) profile. | [optional] [default to null]
**MacDiscoveryProfile** | [***ExtObjectReference**](ExtObjectReference.md) | The MAC Discovery profile. Defines how the segment discovers MAC addresses. | [optional] [default to null]
**IpDiscoveryProfile** | [***ExtObjectReference**](ExtObjectReference.md) | The IP Discovery profile. Defines how the segment discovers IP addresses. | [optional] [default to null]
**SegmentSecurityProfile** | [***ExtObjectReference**](ExtObjectReference.md) | The Segment Security profile. Enables stateless L2 and L3 security on the segment. | [optional] [default to null]
**SpoofGuardProfile** | [***ExtObjectReference**](ExtObjectReference.md) | The Spoof Guard profile. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


