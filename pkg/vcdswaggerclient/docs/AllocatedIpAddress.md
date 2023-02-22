# AllocatedIpAddress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for record. This is to support UI operations. | [optional] [default to null]
**EntityId** | **string** | Id of the entity to which the IP address is allocated such as a VM. | [optional] [default to null]
**EntityName** | **string** | Name of the entity to which the IP address is allocated. | [optional] [default to null]
**VAppName** | **string** | Name of the vApp whose VM is using allocated IP address. | [optional] [default to null]
**IpAddress** | **string** | The allocated IP address. | [optional] [default to null]
**Deployed** | **bool** | Whether the entity using this IP address is currently deployed. | [optional] [default to null]
**AllocationType** | [***AllocatedIpAddressAllocationType**](AllocatedIpAddressAllocationType.md) | The possible allocation types for allocated network address like VSM allocated, VM allocated or NAT routed. | [optional] [default to null]
**NetworkRef** | [***EntityReference**](EntityReference.md) | The vApp network or Org vDC network to which this entity is connected.  | [optional] [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | The organization that this entity belongs to.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


