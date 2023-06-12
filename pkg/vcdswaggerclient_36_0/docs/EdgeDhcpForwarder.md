# EdgeDhcpForwarder

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A flag indicating whether DHCP Forwarder is enabled or not. This should be enabled in order to use the DHCP relay on routed Org vDC networks.  | [optional] [default to null]
**DhcpServers** | **[]string** | DHCP server IP addresses. Both IPv4 and IPv6 addresses are supported. The DHCP servers can be in any subnet, outside the SDDC, or in the physical network.  | [optional] [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


