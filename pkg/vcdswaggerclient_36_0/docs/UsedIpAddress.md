# UsedIpAddress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Id of the entity using the IP address, such as a VM. | [optional] [default to null]
**EntityName** | **string** | Name of the entity using the IP address. | [optional] [default to null]
**VAppName** | **string** | Name of the vApp whose VM is using this IP address. | [optional] [default to null]
**IpAddress** | **string** | The IP address in use. | [optional] [default to null]
**Deployed** | **bool** | Whether the entity using this IP address is currently deployed. | [optional] [default to null]
**AllocationType** | [***UsedIpAddressAllocationType**](UsedIpAddressAllocationType.md) | The possible allocation types for network address like VSM allocated, VM allocated or NAT routed. | [optional] [default to null]
**NetworkRef** | [***EntityReference**](EntityReference.md) | The vApp network or Org vDC network to which this entity is connected.  | [optional] [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | The organization that this entity belongs to.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


