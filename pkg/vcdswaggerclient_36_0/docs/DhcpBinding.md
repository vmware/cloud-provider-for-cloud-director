# DhcpBinding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the DHCP binding. | [optional] [default to null]
**Name** | **string** | Display name for the DHCP binding. | [default to null]
**Description** | **string** | Description for the DHCP binding. | [optional] [default to null]
**MacAddress** | **string** | MAC address of the host. | [default to null]
**IpAddress** | **string** | IP Address assigned to host. This address must belong to the subnet of Org vDC network. For IPv4, this is required. For IPv6, when not specified, no ip address will be assigned to client host.  | [optional] [default to null]
**LeaseTime** | **int64** | The amount of time in seconds of how long a DHCP IP will be leased out for. The minimum is 60s while the maximum is 4,294,967,295s, which is roughly 49,710 days. Default is 24 hours.  | [optional] [default to 86400]
**DnsServers** | **[]string** | DNS nameservers to be set to client host. | [optional] [default to null]
**BindingType** | **string** | The type of DHCP binding. &lt;ul&gt;   &lt;li&gt;IPV4 - an IPv4 DHCP binding.   &lt;li&gt;IPV6 - an IPv6 DHCP binding. &lt;/ul&gt;  | [optional] [default to null]
**DhcpV4BindingConfig** | [***DhcpV4BindingConfig**](DhcpV4BindingConfig.md) | Additional configuration for IPv4 DHCP binding. This is ignored for IPV6 binding.  | [optional] [default to null]
**DhcpV6BindingConfig** | [***DhcpV6BindingConfig**](DhcpV6BindingConfig.md) | Additional configuration for IPv6 DHCP binding. This is ignored for IPV4 binding.  | [optional] [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


