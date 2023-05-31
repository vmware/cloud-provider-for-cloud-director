# ProviderVdc

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the pvdc. | [optional] [default to null]
**Name** | **string** | The name of the pvdc. | [optional] [default to null]
**Description** | **string** | The description of the pvdc. | [optional] [default to null]
**IsEnabled** | **bool** | Whether the pvdc is enabled or not. | [optional] [default to null]
**MaxSupportedHwVersion** | **string** | The maximum hardware version this pvdc supports. | [optional] [default to null]
**NsxTManager** | [***EntityReference**](EntityReference.md) | The NSX-T manager of the pvdc, if any. | [optional] [default to null]
**VimServer** | [***EntityReference**](EntityReference.md) | The vCenter server the pvdc belongs to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


