# EdgeDnsConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | True means that the forwarder is enabled. False means it&#39;s disabled. | [optional] [default to null]
**ListenerIp** | **string** | The IP on which the DNS forwarder listens. If the Edge Gateway has a dedicated external network, this can be changed. | [optional] [default to null]
**DefaultForwarderZone** | [***NsxtDnsForwarderZoneConfig**](NsxtDnsForwarderZoneConfig.md) | The default forwarder zone to use if there&#39;s no matching domain in the conditional forwarder zone. | [optional] [default to null]
**ConditionalForwarderZones** | [**[]NsxtDnsForwarderZoneConfig**](NsxtDnsForwarderZoneConfig.md) | The list of forwarder zones with its matching DNS domains. | [optional] [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [optional] [default to null]
**SnatRuleEnabled** | **bool** | Whether there is an SNAT rule exists for the DNS forwarder or not. In NAT routed environments, an SNAT rule is required for the Edge DNS forwarder to send traffic to an upstream server. In fully routed environments, this is not needed if the listener IP is on an advertised subnet. If the Edge Gateway has a dedicated external network and the listener IP has been changed, there will not be an SNAT rule for the DNS forwarder. In all other cases the SNAT rule will exist.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


