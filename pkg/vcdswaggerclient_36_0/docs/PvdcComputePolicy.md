# PvdcComputePolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | UUID for pVDC compute policy. This is immutable.  | [optional] [default to null]
**Name** | **string** | Display name.  | [default to null]
**Description** | **string** |  | [optional] [default to null]
**PvdcId** | **string** | URN for Provider VDC. If not provided, then this policy is considered Global.  | [optional] [default to null]
**NamedVmGroupReferences** | [**[]EntityReference**](EntityReference.md) | This field cannot be updated and is a read-only field in the client after creation. Deprecated in Api 33.0, this property will be removed in future release.  | [optional] [default to null]
**NamedVmGroups** | [**[][]EntityReference**](array.md) | List of list of vmGroups grouped together in a meaningful manner. A group of vmGroups would consist of one functionally equal vmGroup picked from each cluster of the pvdc.  | [optional] [default to null]
**LogicalVmGroupReferences** | [**[]EntityReference**](EntityReference.md) | This field cannot be updated and is a read-only field after creation.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


