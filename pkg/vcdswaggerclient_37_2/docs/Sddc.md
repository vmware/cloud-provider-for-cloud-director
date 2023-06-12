# Sddc

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**Id** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**VcId** | **string** | URN of the associated vCenter. This is not editable once the SDDC has been created. | [default to null]
**Version** | **string** | Version of the associated vCenter. This is not editable. | [optional] [default to null]
**OverallStatus** | [***OverallStatus**](OverallStatus.md) | Overall status of the associated vCenter. This is not editable. | [optional] [default to null]
**Stats** | [***SddcStats**](SddcStats.md) | Associated read-only statistics. | [optional] [default to null]
**DefaultProxy** | [***SddcProxy**](SddcProxy.md) | Default proxy for the SDDC. This field is read-only. To set a new default, edit the proxy you wish to make the new default by modifying its defaultProxy flag. Deprecated in Api 34.0.  | [optional] [default to null]
**DefaultEndpoint** | [***SddcEndpoint**](SddcEndpoint.md) | Default endpoint for the SDDC. There is an endpoint available for an SDDC even if there are no proxies configured for the SDDC. This indicates that an endpoint is available that doesn&#39;t require proxying, either because it is publicly accessible or because it assumes an established VPM connection. The field is read-only.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


