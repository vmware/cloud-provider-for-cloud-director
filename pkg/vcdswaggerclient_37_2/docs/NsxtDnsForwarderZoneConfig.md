# NsxtDnsForwarderZoneConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the DNS forwarder zone. If value is unset, a new zone is created. If value is set, an update is done on the zone.  | [optional] [default to null]
**DisplayName** | **string** | User friendly name for the zone. | [optional] [default to null]
**DnsDomainNames** | **[]string** | List of domain names on which conditional forwarding is based. This field is required if the DNS Zone is being used for a conditional forwarder. This field will also be used for conditional reverse lookup. This field should not be set if the zone is used as default forwarder zone.  | [optional] [default to null]
**UpstreamServers** | **[]string** | DNS servers to which the DNS request needs to be forwarded. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


