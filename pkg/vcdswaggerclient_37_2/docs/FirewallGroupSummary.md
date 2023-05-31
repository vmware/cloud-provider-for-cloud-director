# FirewallGroupSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | The organization that this firewall group belongs to. This property is read-only and cannot be updated.  | [optional] [default to null]
**EdgeGatewayRef** | [***EntityReference**](EntityReference.md) | The edge gateway that this firewall group is scoped to. This means that this firewall group can be used when configuring firewall rules for the edge gateway. This property is now deprecated. ownerRef should be used instead  | [optional] [default to null]
**OwnerRef** | [***EntityReference**](EntityReference.md) | The vDC Group or Edge Gateway that this firewall group is scoped to. This group can be used for configuring rules for either an Edge Gateway or vDC Group. If an Edge Gateway is specified that belongs to a vDC Group, the the firewall group will be scoped to the vDC Group.  | [optional] [default to null]
**NetworkProviderScope** | **string** | The network provider scope that this object belongs to. This is a read-only property and is determined by the input context entity ID during object creation.  | [optional] [default to null]
**Id** | **string** | The id of the firewall group. | [optional] [default to null]
**Name** | **string** | The name of the firewall group. | [default to null]
**Description** | **string** | The description of the firewall group | [optional] [default to null]
**Type_** | [***FirewallGroupType**](FirewallGroupType.md) | Defines the type of Firewall Group which determines what can be members of this group such as IP Addresses, Org vDC networks, or VMs based on dynamic criteria. This property is now deprecated and replaced with typeValue.  | [optional] [default to null]
**TypeValue** | **string** | Defines the type of Firewall Group which determines what can be members of this group such as IP Addresses, Org vDC networks, or VMs based on dynamic criteria.  Below are valid values. &lt;ul&gt;   &lt;li&gt; &lt;code&gt; IP_SET &lt;/code&gt; should be used when using particular IP Addresses of VMs, Networks, etc.   &lt;li&gt; &lt;code&gt; STATIC_MEMBERS &lt;/code&gt; should be used when specifying exact members such as a particular Org vDC Network.   &lt;li&gt; &lt;code&gt; VM_CRITERIA &lt;/code&gt; should be used when specifying some dynamic criteria that matches a VM member such as VM name or Operating System name.        This type is valid only if the firewall group is scoped to a vDC Group. &lt;/ul&gt; The default is IP_SET.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


