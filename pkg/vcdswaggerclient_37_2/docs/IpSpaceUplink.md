# IpSpaceUplink

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**Id** | **string** | The ID of the IP Space Uplink in URN format. | [optional] [default to null]
**Name** | **string** | The tenant-visible name of the IP Space Uplink. Name is unique across all IP Space Uplinks associated with a Provider Gateway.  | [default to null]
**Description** | **string** | The description of the IP Space Uplink. | [optional] [default to null]
**ExternalNetworkRef** | [***EntityReference**](EntityReference.md) | The External Network that is backed by the Provider Gateway this uplink is associated with. An External Network is used to reference the Provider Gateway since the External Network is backed by a Provider Gateway. This property is not updatable.  | [default to null]
**IpSpaceRef** | [***EntityReference**](EntityReference.md) | The IP Space associated with this uplink. This property is not updatable.  | [default to null]
**IpSpaceType** | **string** | The type of the IP Space associated with this uplink. Possible values are: PUBLIC, PRIVATE, SHARED_SERVICES. This property is read-only.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


