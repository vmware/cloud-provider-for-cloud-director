# DefinedEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the defined entity in URN format.  | [optional] [default to null]
**EntityType** | **string** | The URN ID of the defined entity type that the entity is an instance of. This is a read-only field.  | [optional] [default to null]
**Name** | **string** | The name of the defined entity.  | [default to null]
**ExternalId** | **string** | An external entity&#39;s id that this entity may have a relation to.  | [optional] [default to null]
**Entity** | [**map[string]interface{}**](interface{}.md) | A JSON value representation. The JSON will be validated against the schema of the entityType that the entity is an instance of.  | [default to null]
**State** | **string** | Every entity is created in the \&quot;PRE_CREATED\&quot; state. Once an entity is ready to be validated against its schema, it will transition in another state - RESOLVED, if the entity is valid according to the schema, or RESOLUTION_ERROR otherwise. If an entity in an \&quot;RESOLUTION_ERROR\&quot; state is updated, it will transition to the inital \&quot;PRE_CREATED\&quot; state without performing any validation. If its in the \&quot;RESOLVED\&quot; state, then it will be validated against the entity type schema and throw an exception if its invalid.  | [optional] [default to null]
**Owner** | [***EntityReference**](EntityReference.md) | The owner of the defined entity. | [optional] [default to null]
**Org** | [***EntityReference**](EntityReference.md) | The organization of the defined entity. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


