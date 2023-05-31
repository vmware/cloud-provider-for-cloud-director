# SecurityTaggedEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the entity in URN format.  | [optional] [default to null]
**Name** | **string** | The name of the entity. | [optional] [default to null]
**ParentRef** | [***EntityReference**](EntityReference.md) | The parent of the entity such as vApp if the entity is a VM. If not applicable, field is not set. | [optional] [default to null]
**OwnerRef** | [***EntityReference**](EntityReference.md) | The owner of the specified entity such as vDC or vDC Group. If not applicable, field is not set. | [optional] [default to null]
**EntityType** | **string** | The type of entity. Currently only \&quot;vm\&quot; is supported. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


