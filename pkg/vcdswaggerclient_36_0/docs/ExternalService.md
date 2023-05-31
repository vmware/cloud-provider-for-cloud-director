# ExternalService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the external service. | [default to null]
**Id** | **string** | The unique id of the external service. | [optional] [default to null]
**Version** | **string** | The extension&#39;s version. The version should follow semantic versioning rules. The combination of vendor-namespace-version must be unique.  | [default to null]
**Vendor** | **string** | The vendor name. The combination of vendor-namespace-version must be unique.  | [default to null]
**Priority** | **int32** | Extension service priority. An integer between 0-100. A value of 50 denotes a neutral priority.  | [optional] [default to 50]
**Enabled** | **bool** | Whether the extension is enabled or not. | [default to null]
**AuthorizationEnabled** | **bool** | Whether authorization is enabled for the service. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**MqttTopics** | **map[string]string** | The map of MQTT topics this extension will communicate through.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


