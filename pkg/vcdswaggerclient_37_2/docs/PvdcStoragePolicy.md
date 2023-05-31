# PvdcStoragePolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique VCD Id for the policy. | [optional] [default to null]
**Name** | **string** | Unique name for the policy. | [default to null]
**StoragePolicyMoref** | **string** | Unique Id in Virtual Center of the policy. | [default to null]
**IsEnabled** | **bool** | Enabled state of the policy, defaults to true. | [optional] [default to null]
**ProviderVdcRef** | [***EntityReference**](EntityReference.md) | The PVDC that this policy belongs to. | [optional] [default to null]
**VcRef** | [***EntityReference**](EntityReference.md) | The VC that this policy belongs to. | [optional] [default to null]
**TotalCapacityMb** | **int64** | Total capacity in MB for this storage policy | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


