# DefinedEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the defined entity in URN format.  | [optional] [default to null]
**EntityType** | **string** | The URN ID of the defined entity type that the entity is an instance of. This is a read-only field.  | [optional] [default to null]
**Name** | **string** | The name of the defined entity.  | [default to null]
**ExternalId** | **string** | An external entity&#39;s id that this entity may have a relation to.  | [optional] [default to null]
**Entity** | **map[string]interface{}** | A JSON value representation. The JSON will be validated against the schema of the entityType that the entity is an instance of.  | [default to null]
**State** | **string** | Every entity is created in the \&quot;PRE_CREATED\&quot; state. Once an entity is ready to be validated against its schema, it will transition in another state - RESOLVED, if the entity is valid according to the schema, or RESOLUTION_ERROR otherwise. If an entity in an \&quot;RESOLUTION_ERROR\&quot; state is updated, it will transition to the initial \&quot;PRE_CREATED\&quot; state without performing any validation. If its in the \&quot;RESOLVED\&quot; state, then it will be validated against the entity type schema and throw an exception if its invalid.  | [optional] [default to null]
**EntityState** | **string** | The \&quot;entityState\&quot; field may have the following values: \&quot;PRE_CREATED\&quot;, \&quot;RESOLVED\&quot;, \&quot;RESOLUTION_ERROR\&quot;, \&quot;IN_DELETION\&quot;. Every entity is created in the \&quot;PRE_CREATED\&quot; state. Once an entity is ready to be validated against its schema, it will transition in another state - \&quot;RESOLVED\&quot;, if the entity is valid according to the schema, or \&quot;RESOLUTION_ERROR\&quot; otherwise. If an entity in an \&quot;RESOLUTION_ERROR\&quot; state is updated, it will transition to the initial \&quot;PRE_CREATED\&quot; state without performing any validation. If its in the \&quot;RESOLVED\&quot; state, then it will be validated against the entity type schema and throw an exception if its invalid. Once an entity is ready to be deleted, if any external resources need to be cleaned up, the entity deletion process can be set up into multiple stages. The entity will transition into state IN_DELETION when the finalization process starts and will be permanently deleted once it completes successfully. If the finalization is not successful, the entity will once it completes successfully. If the finalization is not successful, the entity will remain in IN_DELETION state.  | [optional] [default to null]
**Owner** | [***EntityReference**](EntityReference.md) | The owner of the defined entity. | [optional] [default to null]
**Org** | [***EntityReference**](EntityReference.md) | The organization of the defined entity. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


