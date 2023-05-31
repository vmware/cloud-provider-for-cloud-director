# SegmentIpDiscoveryProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the segment profile. | [optional] [default to null]
**DisplayName** | **string** | Name of the segment profile. This corresponds to the name used in NSX-T manager&#39;s logs or GUI. | [optional] [default to null]
**Description** | **string** | The description of the segment profile. | [optional] [default to null]
**NsxTManagerRef** | [***EntityReference**](EntityReference.md) | The NSX-T manager where this segment profile is configured. | [optional] [default to null]
**IsArpSnoopingEnabled** | **bool** | Whether ARP snooping is enabled. | [optional] [default to null]
**ArpBindingLimit** | **int32** | Indicates the number of arp snooped IP addresses to be remembered per LogicalPort. | [optional] [default to null]
**IsDhcpSnoopingV4Enabled** | **bool** | Whether DHCP snooping for IPv4 is enabled. | [optional] [default to null]
**IsVmToolsV4Enabled** | **bool** | Whether fetching IPv4 address using vm-tools is enabled. This option is only supported on ESX where vm-tools is installed. | [optional] [default to null]
**IsDhcpSnoopingV6Enabled** | **bool** |  | [optional] [default to null]
**IsVmToolsV6Enabled** | **bool** | Whether fetching IPv6 address using vm-tools is enabled. This will learn the IPv6 addresses which are configured on interfaces of a VM with the help of the VMTools software.  | [optional] [default to null]
**IsNdSnoopingEnabled** | **bool** | Whether ND snooping is enabled.  If true, this method will snoop the NS (Neighbor Solicitation) and NA (Neighbor Advertisement) messages in the ND (Neighbor Discovery Protocol) family of messages which are transmitted by a VM. From the NS messages, we will learn about the source which sent this NS message. From the NA message, we will learn the resolved address in the message which the VM is a recipient of. Addresses snooped by this method are subject to TOFU.  | [optional] [default to null]
**NdSnoopingLimit** | **int32** | Maximum number of ND (Neighbor Discovery Protocol) snooped IPv6 addresses. | [optional] [default to null]
**ArpNdBindingTimeout** | **int32** | ARP and ND cache timeout (in minutes). | [optional] [default to null]
**IsDuplicateIpDetectionEnabled** | **bool** | Whether duplicate IP detection is enabled. Duplicate IP detection is used to determine if there is any IP conflict with any other port on the same logical switch. If a conflict is detected, then the IP is marked as a duplicate on the port where the IP was discovered last.  | [optional] [default to null]
**IsTofuEnabled** | **bool** | Whether \&quot;Trust on First Use(TOFU)\&quot; paradigm is enabled. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


