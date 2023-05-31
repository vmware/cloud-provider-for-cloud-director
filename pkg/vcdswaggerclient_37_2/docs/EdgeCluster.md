# EdgeCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Edge Cluster in URN format. | [optional] [default to null]
**Name** | **string** | The name of the Edge Cluster. Names for Edge Clusters must be unique across the system. | [default to null]
**Description** | **string** |  | [optional] [default to null]
**ResourcePool** | [***ResourcePool**](ResourcePool.md) | The Resource Pool in vCenter where the Edge VM will be deployed. | [optional] [default to null]
**StorageProfileName** | **string** | Name of the Storage Profile. This will define the set of datastores where the edge vm will be deployed. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


