# GatewayQoSProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the gateway QoS profile. | [optional] [default to null]
**DisplayName** | **string** | Name of the gateway QoS profile. This corresponds to the name used in NSX-T manager&#39;s logs or GUI. | [optional] [default to null]
**Description** | **string** | The description of the gateway QoS profile. | [optional] [default to null]
**CommittedBandwidth** | **int32** | Committed bandwidth in both directions specificd in Mb/s. Bandwidth is limited to line rate when the value configured is greater than line rate. | [optional] [default to null]
**BurstSize** | **int32** | Burst size in bytes. | [optional] [default to null]
**ExcessAction** | **string** | Action on traffic exceeding bandwidth. | [optional] [default to null]
**NsxTManagerRef** | [***EntityReference**](EntityReference.md) | The NSX-T manager where this gateway QoS profile is configured. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


