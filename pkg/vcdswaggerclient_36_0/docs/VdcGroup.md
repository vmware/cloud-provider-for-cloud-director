# VdcGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID for the vDC Group (read-only). | [optional] [default to null]
**OrgId** | **string** | The organization that this group belongs to. | [default to null]
**Name** | **string** | The name of this group. The name must be unique. | [default to null]
**Description** | **string** | The description of this group. | [optional] [default to null]
**LocalEgress** | **bool** | Determines whether local egress is enabled for a universal router belonging to a universal vDC group. This value is used on create if universalNetworkingEnabled is set to true. This cannot be updated. This value is always false for local vDC groups.  | [optional] [default to null]
**ParticipatingOrgVdcs** | [**[]ParticipatingVdcReference**](ParticipatingVdcReference.md) | The list of organization vDCs that are participating in this group. | [default to null]
**UniversalNetworkingEnabled** | **bool** | True means that a vDC group router has been created. If set to true for vdc group creation, a universal router will also be created. | [optional] [default to null]
**NetworkPoolUniversalId** | **string** | The network provider&#39;s universal id that is backing the universal network pool. This field is read-only and is derived from the list of participating vDCs if a universal vDC group is created. For universal vDC groups, each participating vDC should have a universal network pool that is backed by this same id.  | [optional] [default to null]
**NetworkPoolId** | **string** | ID of network pool to use if creating a local vDC group router. Must be set if creating a local group. Ignored if creating a universal group.  | [optional] [default to null]
**Status** | [***VdcGroupEntityStatus**](VdcGroupEntityStatus.md) | The status that the group can be in. | [optional] [default to null]
**Type_** | **string** | Defines the group as LOCAL or UNIVERSAL. This cannot be changed. Local vDC Groups can have networks stretched across multiple vDCs in a single Cloud Director instance. Local vDC Groups share the same broadcast domain/transport zone and network provider scope. Universal vDC groups can have networks stretched across multiple vDCs in a single or multiple Cloud Director instance(s). Universal vDC groups are backed by a broadcast domain/transport zone that strectches across a single or multiple Cloud Director instance(s). Local vDC groups are supported for both NSX-V and NSX-T Network Provider Types. Universal vDC Groups are supported for only NSX_V Network Provider Type.  | [optional] [default to null]
**NetworkProviderType** | **string** | The values currently supported are NSX_V and NSX_T. Defines the networking provider backing the vDC Group. This is used on create. If not specified, NSX_V value will be used. NSX_V is used for existing vDC Groups and vDC Groups where Cross-VC NSX is used for the underlying technology. NSX_T is used when the networking provider type for the Organization vDCs in the group is NSX-T. NSX_T only supports groups of type LOCAL (single site).  | [optional] [default to null]
**DfwEnabled** | **bool** | Whether Distributed Firewall is enabled for this vDC Group. Only applicable for NSX_T vDC Groups. | [optional] [default to null]
**ErrorMessage** | **string** | If the group has an error status, a more detailed error message is set here. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


