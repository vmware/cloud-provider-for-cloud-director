# SegmentSecurityProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the segment profile. | [optional] [default to null]
**DisplayName** | **string** | Name of the segment profile. This corresponds to the name used in NSX-T manager&#39;s logs or GUI. | [optional] [default to null]
**Description** | **string** | The description of the segment profile. | [optional] [default to null]
**NsxTManagerRef** | [***EntityReference**](EntityReference.md) | The NSX-T manager where this segment profile is configured. | [optional] [default to null]
**IsBpduFilterEnabled** | **bool** | Whether BPDU filter is enabled. | [optional] [default to null]
**BpduFilterAllowList** | **[]string** | Pre-defined list of allowed MAC addresses to be excluded from BPDU filtering. | [optional] [default to null]
**IsDhcpClientBlockV4Enabled** | **bool** | Whether DHCP Client block IPv4 is enabled. This filters DHCP Client IPv4 traffic. | [optional] [default to null]
**IsDhcpClientBlockV6Enabled** | **bool** | Whether DHCP Client block IPv6 is enabled. This filters DHCP Client IPv6 traffic. | [optional] [default to null]
**IsDhcpServerBlockV4Enabled** | **bool** | Whether DHCP Server block IPv4 is enabled. This filters DHCP Server IPv4 traffic. | [optional] [default to null]
**IsDhcpServerBlockV6Enabled** | **bool** | Whether DHCP Server block IPv6 is enabled. This filters DHCP Server IPv6 traffic. | [optional] [default to null]
**IsNonIpTrafficBlockEnabled** | **bool** | Whether non IP traffic block is enabled. If true, it blocks all traffic except IP/(G)ARP/BPDU. | [optional] [default to null]
**IsRaGuardEnabled** | **bool** | Whether Router Advertisement Guard is enabled. This filters DHCP Server IPv6 traffic. | [optional] [default to null]
**IsRateLimitingEnabled** | **bool** | Whether Rate Limiting is enabled. | [optional] [default to null]
**RateLimits** | [***TrafficRateLimitConfig**](TrafficRateLimitConfig.md) | Allows configuration of rate limits for broadcast and multicast traffic. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


