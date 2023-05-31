# StorageContainer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique VCD Id for the Storage Container. | [optional] [default to null]
**Name** | **string** | Name for the Storage Container. | [optional] [default to null]
**DatastoreType** | **string** | Datastore Type of the Storage Container. | [optional] [default to null]
**IsDatastoreCluster** | **bool** | True if this Storage Container is a Datastore Cluster. | [optional] [default to null]
**IsEnabled** | **bool** | True if this Storage Container is enabled. Unset if this storage container is part of a Datastore Cluster | [optional] [default to null]
**IsDeleted** | **bool** | Deleted state of the Storage Container. | [optional] [default to null]
**Moref** | **string** | Unique Id in Virtual Center of the Storage Container. | [optional] [default to null]
**VcRef** | [***EntityReference**](EntityReference.md) | The VC that this Storage Container belongs to. | [optional] [default to null]
**TotalStorageMb** | **int64** | Total storage in MB for this Storage Container. | [optional] [default to null]
**UsedStorageMb** | **int64** | Total used storage in MB for this Storage Container. | [optional] [default to null]
**ProvisionedStorageMb** | **int64** | Total provisioned storage in MB for this Storage Container. | [optional] [default to null]
**IopsCapacity** | **int64** | Total IOPS capacity for this Storage Container. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


