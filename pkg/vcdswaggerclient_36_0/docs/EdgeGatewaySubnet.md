# EdgeGatewaySubnet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | **string** | The gateway for the subnet. | [default to null]
**PrefixLength** | **int32** | The prefix length of the subnet. | [default to null]
**DnsSuffix** | **string** | The DNS suffix that VMs attached to this network will use. | [optional] [default to null]
**DnsServer1** | **string** | The first DNS server that VMs attached to this network will use. | [optional] [default to null]
**DnsServer2** | **string** | The second DNS server that VMs attached to this network will use. | [optional] [default to null]
**IpRanges** | [***IpRanges**](IpRanges.md) | Range of IPs within the subnet that can be used in this network. A VM attached to this network is assigned one of these IPs. | [optional] [default to null]
**Enabled** | **bool** | Indicates whether the external network subnet is currently enabled. | [optional] [default to null]
**TotalIpCount** | **int32** | The number of IP addresses defined by the static IP ranges. | [optional] [default to null]
**UsedIpCount** | **int32** | The number of IP address used from the static IP ranges. | [optional] [default to null]
**PrimaryIp** | **string** | The primary IP address allocated for this subnet. If not specified, this IP is auto-allocated.  This IP belongs to the external network and can be used for system-configured NAT rules such as DNS forwarder configuration.  | [optional] [default to null]
**AutoAllocateIpRanges** | **bool** | Used for create and update api calls. If set to true, IP Ranges are automatically generated based on totalIpCount. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


