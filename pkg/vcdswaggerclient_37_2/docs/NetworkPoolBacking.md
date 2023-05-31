# NetworkPoolBacking

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanIdRanges** | [***VlanIdRanges**](VlanIdRanges.md) | The range of backing VLAN Id&#39;s. This information is available only for VLAN backed network pools.  | [optional] [default to null]
**VdsRefs** | [**[]BackingRef**](BackingRef.md) | The information about virtual distributed switches. This information is available only for VLAN, VXLAN and Universal VXLAN backed network pools.  | [optional] [default to null]
**PortGroupRefs** | [**[]BackingRef**](BackingRef.md) | All the vSphere port groups that will be used by this network pool. This information is available only for Portgroup backed network pools.  | [optional] [default to null]
**TransportZoneRef** | [***BackingRef**](BackingRef.md) | The Id of the backing transport zone in NSX.  | [optional] [default to null]
**ProviderRef** | [***EntityReference**](EntityReference.md) | The Id of vCenter server or the NSX-T manager that owns this backing entity. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


