# EntityState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the defined entity in URN format.  | [optional] [default to null]
**Entity** | **map[string]interface{}** | A JSON entity. This entity will be validated against the provided entityType.  | [optional] [default to null]
**State** | **string** | The current state of the entity  | [optional] [default to null]
**EntityState** | **string** | The \&quot;entityState\&quot; field may have the following values: \&quot;RESOLVED\&quot;, \&quot;RESOLUTION_ERROR\&quot;. It represents the current state of the entity.  | [optional] [default to null]
**Message** | **string** | The error message(s), if the entity could not be resolved.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


