# ApplicationPortProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | The organization that this object belongs to. This property is required during creation and cannot be updated.  | [optional] [default to null]
**ContextEntityId** | **string** | The context under which this object is created. The context can the id of the organization vDC, vDC Group, or network provider. This field is required on object creation and is unset on object reads. This same context ID can be used in the context filter field when querying for the list of objects.  | [optional] [default to null]
**NetworkProviderScope** | **string** | The network provider scope that this object belongs to. This is a read-only property and is determined by the input context entity ID during object creation.  | [optional] [default to null]
**Id** | **string** | The id of the Application Port Profile in URN format. | [optional] [default to null]
**Name** | **string** | The name of the Application Port Profile. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Scope** | [***ApplicationPortProfileScopeType**](ApplicationPortProfileScopeType.md) | The scope of the application port profile definition. | [optional] [default to null]
**ApplicationPorts** | [**[]ApplicationPortList**](ApplicationPortList.md) | A list of protocol and ports supported by this application port profile. | [optional] [default to null]
**UsableForNAT** | **bool** | True means that the port profile can be used for NAT configuration.  A port profile can be used for NAT if there is only 1 application port with at most 1 destination port.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


