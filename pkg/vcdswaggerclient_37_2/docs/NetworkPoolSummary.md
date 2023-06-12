# NetworkPoolSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**Id** | **string** | The id of the Network Pool in URN format. | [optional] [default to null]
**Name** | **string** | The name of the Network Pool. Names for Network Pools must be unique across the system. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**PoolType** | [***NetworkPoolBackingType**](NetworkPoolBackingType.md) | The network pool backing type. | [optional] [default to null]
**PromiscuousMode** | **bool** | Whether promiscuous mode is enabled on the network pool. This is only applicable for VLAN network pools. | [optional] [default to null]
**TotalBackingsCount** | **int32** | The number of backings available for use. | [optional] [default to null]
**UsedBackingsCount** | **int32** | The number of network pool backings in use. | [optional] [default to null]
**ManagingOwnerRef** | [***EntityReference**](EntityReference.md) | The Id of vCenter server or the NSX-T manager that manages backings for this network pool. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


