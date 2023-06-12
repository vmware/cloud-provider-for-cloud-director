# OrgVdcStoragePolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique VCD Id for the policy. | [optional] [default to null]
**Name** | **string** | Unique name for the policy. | [default to null]
**IsEnabled** | **bool** | Enabled state of the policy, defaults to true. | [optional] [default to null]
**IsDefaultStoragePolicy** | **bool** | Storage policy is marked as default, defaults to false. | [optional] [default to null]
**StorageLimitMb** | **int64** | Storage limit for the policy. | [optional] [default to null]
**OrgVdcRef** | [***EntityReference**](EntityReference.md) | The Org VDC that this policy belongs to. | [optional] [default to null]
**PvdcStoragePolicyRef** | [***EntityReference**](EntityReference.md) | The parent PVDC storage policy for this Organization VDC storage policy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


