# EdgeGatewayUplink

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UplinkId** | **string** | The identifier of the external network this edge gateway is connected to. | [optional] [default to null]
**UplinkName** | **string** | The name of the external network this edge gateway is connected to. | [optional] [default to null]
**Subnets** | [***EdgeGatewaySubnets**](EdgeGatewaySubnets.md) | Set of subnets this edge will have access to.  | [optional] [default to null]
**Connected** | **bool** | Whether or not the gateway is connected to this uplink. This value defaults to true if it is not set. When filtering by this field, if the filter is false, all gateways that have 0 connected uplinks will be returned, while if it is true, all gateways with at least one connected uplink will be returned. | [optional] [default to null]
**QuickAddAllocatedIpCount** | **int32** | If set on create or update api calls, the specified number of IP addresses will be additionally allocated for this uplink. IPs will be allocated from multiple subnets if needed. | [optional] [default to null]
**Dedicated** | **bool** | If set to true, then the associated external network is exclusively used by this edge gateway. | [optional] [default to null]
**VrfLiteBacked** | **bool** | Whether the associated external network is backed by a NSX-T VRF-Lite Tier-0. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


