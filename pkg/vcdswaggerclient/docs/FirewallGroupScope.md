# FirewallGroupScope

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgRef** | [***EntityReference**](EntityReference.md) | The organization that this firewall group belongs to. This property is read-only and cannot be updated.  | [optional] [default to null]
**EdgeGatewayRef** | [***EntityReference**](EntityReference.md) | The edge gateway that this firewall group is scoped to. This means that this firewall group can be used when configuring firewall rules for the edge gateway. This property is now deprecated. ownerRef should be used instead  | [optional] [default to null]
**OwnerRef** | [***EntityReference**](EntityReference.md) | The vDC Group or Edge Gateway that this firewall group is scoped to. This group can be used for configuring rules for either an Edge Gateway or vDC Group. If an Edge Gateway is specified that belongs to a vDC Group, the the firewall group will be scoped to the vDC Group.  | [optional] [default to null]
**NetworkProviderScope** | **string** | The network provider scope that this object belongs to. This is a read-only property and is determined by the input context entity ID during object creation.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


