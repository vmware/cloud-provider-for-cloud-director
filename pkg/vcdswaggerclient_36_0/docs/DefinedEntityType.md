# DefinedEntityType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the defined entity type in URN format.  | [optional] [default to null]
**Name** | **string** | The name of the defined entity type.  | [optional] [default to null]
**Description** | **string** | Description of the defined entity type.  | [optional] [default to null]
**Nss** | **string** | A unique namespace specific string. The combination of nss and version must be unique.  | [optional] [default to null]
**Version** | **string** | The version of the defined entity type. The combination of nss and version must be unique. The version string must follow semantic versioning rules.  | [optional] [default to null]
**InheritedVersion** | **string** | To be used when creating a new version of a defined entity type. Specifies the version of the type that will be the template for the authorization configuration of a the new version. The Type ACLs and the access requirements of the Type Behaviors of the new version will be copied from those of the inherited version. If the value of this property is &#39;0&#39;, then the new type version will not inherit another version and will have the default authorization settings, just like the first version of a new type.  | [optional] [default to null]
**ExternalId** | **string** | An external entity&#39;s id that this definition may apply to.  | [optional] [default to null]
**Schema** | [**map[string]interface{}**](interface{}.md) | The JSON-Schema valid definition of the defined entity type. If no JSON Schema version is specified, version 4 will be assumed.  | [optional] [default to null]
**Vendor** | **string** | The vendor name.  | [optional] [default to null]
**Interfaces** | **[]string** | List of interface ids that this defined entity type is referenced by.  | [optional] [default to null]
**Hooks** | [**map[string]interface{}**](interface{}.md) | A mapping defining which behaviors should be invoked upon specific lifecycle events, like PostCreate, PostUpdate, PreDelete. For example: \&quot;hooks\&quot;: { \&quot;PostCreate\&quot;: \&quot;urn:vcloud:behavior-interface:postCreateHook:vendorA:containerCluster:1.0.0\&quot; }  | [optional] [default to null]
**Readonly** | **bool** | True if the entity type cannot be modified. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


