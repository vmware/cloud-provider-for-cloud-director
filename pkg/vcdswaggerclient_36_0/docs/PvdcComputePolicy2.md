# PvdcComputePolicy2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | UUID for pVDC compute policy. This is immutable.  | [optional] [default to null]
**Name** | **string** | Display name.  | [default to null]
**PolicyType** | **string** | The discriminator type is used to differentiate among various sub policy types.  | [default to null]
**Description** | **string** |  | [optional] [default to null]
**PvdcId** | **string** | URN for Provider vDC. If not provided, then this policy is considered Global.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


