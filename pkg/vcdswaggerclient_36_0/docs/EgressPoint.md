# EgressPoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID for the Universal Egress Point (read-only). | [optional] [default to null]
**VdcRef** | [***EntityReference**](EntityReference.md) | The Org vDC the Universal Egress Point belongs to. | [optional] [default to null]
**GatewayRef** | [***EntityReference**](EntityReference.md) | The Edge Gateway the Egress Point is referring to. | [optional] [default to null]
**NetworkProviderScope** | **string** | Read-only field that specifies the network provider scope of the Universal Egress Point (inherited from the Org vDC). | [optional] [default to null]
**IsUsedForRouting** | **bool** | Specifies whether the egress point is being used for Universal Routing. This is a read-only field. | [optional] [default to null]
**Status** | [***VdcGroupEntityStatus**](VdcGroupEntityStatus.md) | The status of the Universal Egress Point. | [optional] [default to null]
**ErrorMessage** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


