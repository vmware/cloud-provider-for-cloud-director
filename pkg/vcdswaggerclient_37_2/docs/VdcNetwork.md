# VdcNetwork

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID for the network. This field is read-only. | [optional] [default to null]
**Name** | **string** | The name of the network. | [optional] [default to null]
**Description** | **string** | The description of the network. | [optional] [default to null]
**Subnets** | [***Subnets**](Subnets.md) | List of subnets configured for the network. | [optional] [default to null]
**BackingNetworkId** | **string** | The NSX id of the backing network. | [optional] [default to null]
**BackingNetworkType** | [***BackingNetworkType**](BackingNetworkType.md) | The object type of the backing network. | [optional] [default to null]
**ParentNetworkId** | [***EntityReference**](EntityReference.md) | The parent network if the network is a direct network, otherwise it will be null. | [optional] [default to null]
**NetworkType** | [***VdcNetworkFenceType**](VdcNetworkFenceType.md) | The type of network. Changing the network type allows converting between an isolated and routed network. Note that the \&quot;connection\&quot; field must also be set if converting from isolated to routed network.  | [optional] [default to null]
**OrgVdc** | [***EntityReference**](EntityReference.md) | The organization vDC the network belongs to. This should be unset if the network is owned by a vDC Group. For API version 35.0 and above, this field will be treated as read only. Please use ownerRef for new network creation.  | [optional] [default to null]
**OwnerRef** | [***EntityReference**](EntityReference.md) | The org vDC or vDC Group that this network belongs to. If the ownerRef is set to a vDC Group, this network will be available across all the vDCs in the vDC Group. If the vDC Group is backed by a NSX-V network provider, the org vDC network is automatically connected to the distributed router associated with the vDC Group and the \&quot;connection\&quot; field does not need to be set. For API version 35.0 and above, this field should be set for network creation.  | [optional] [default to null]
**OrgVdcIsNsxTBacked** | **bool** | For an Org vDC network, whether the vDC is backed by NSX-T. | [optional] [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | The organization to which the network belongs. | [optional] [default to null]
**Connection** | [***RouterConnection**](RouterConnection.md) | The edge gateway that the network is attached to.  | [optional] [default to null]
**IsDefaultNetwork** | **bool** | Deprecated unused field, this property will be removed in future release.  | [optional] [default to null]
**Shared** | **bool** | Whether this network is shared with other organization vDCs. | [optional] [default to null]
**EnableDualSubnetNetwork** | **bool** | Whether or not this network will support two subnets | [optional] [default to null]
**Status** | [***OrgVdcNetworkStatus**](OrgVdcNetworkStatus.md) | Description of the network&#39;s status.  | [optional] [default to null]
**LastTaskFailureMessage** | **string** | Brief failure message if the last configuration task failed. Deprecated in Api 33.0, this property will be removed in next release.  | [optional] [default to null]
**GuestVlanTaggingAllowed** | **bool** | Whether guest VLAN tagging is allowed. | [optional] [default to null]
**RetainNicResources** | **bool** | Whether network resources such as IP/MAC Addresses are to be retained. | [optional] [default to null]
**CrossVdcNetworkId** | **string** | The id of the cross vdc network if this is a stretched network, otherwise it will be null. | [optional] [default to null]
**CrossVdcNetworkLocationId** | **string** | The id of the org from which this network can be managed if this is a stretched network, otherwise it will be null. | [optional] [default to null]
**OverlayId** | **int32** | Overlay connectivity ID for this Network. This field is used on creation during POST only and will not be displayed on an object returned through GET or PUT.  | [optional] [default to null]
**TotalIpCount** | **int32** | The number of IP addresses defined by the static ip pools. If the network contains any IpV6 subnets, the total ip count will be null. | [optional] [default to null]
**UsedIpCount** | **int32** | The number of IP address used from the static ip pools. | [optional] [default to null]
**RouteAdvertised** | **bool** | Whether this network is advertised so that it can be routed out to the external networks. This applies only to routed network backed by NSX-T. Value will be unset or set to false if route advertisement is not applicable.  | [optional] [default to null]
**SecurityGroups** | [**[]EntityReference**](EntityReference.md) | The list of firewall groups of type SECURITY_GROUP/STATIC_MEMBERS that are assigned to the Org vDC Network. These groups can then be used in firewall rules to protect the Org vDC Network and allow/deny traffic.  | [optional] [default to null]
**SegmentProfileTemplateRef** | [***EntityReference**](EntityReference.md) | Reference to the Segment Profile Template that is to be used when creating/updating this network. Setting this will override any Org vDC Network Segment Profile Template defined at global level or an Org vDC level. This field is only relevant during network create/update operation and will not be returned on GETs. For specific profile types where there are no corresponding profiles defined in the template, VCD will use the default NSX-T profile. This field is only applicable for NSX-T Org vDC Networks.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


