# EdgeGatewayUplink

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UplinkId** | **string** | The identifier of the external network/provider gateway this edge gateway is connected to. | [optional] [default to null]
**UplinkName** | **string** | The name of the external network/provider gateway this edge gateway is connected to. | [optional] [default to null]
**Subnets** | [***EdgeGatewaySubnets**](EdgeGatewaySubnets.md) | Set of subnets this edge will have access to.  | [optional] [default to null]
**Connected** | **bool** | Whether or not the gateway is connected to this uplink. This value defaults to true if it is not set. When filtering by this field, if the filter is false, all gateways that have 0 connected uplinks will be returned, while if it is true, all gateways with at least one connected uplink will be returned. | [optional] [default to null]
**QuickAddAllocatedIpCount** | **int32** | If set on create or update api calls, the specified number of IP addresses will be additionally allocated for this uplink. IPs will be allocated from multiple subnets if needed. | [optional] [default to null]
**Dedicated** | **bool** | This property can be set to make the provider gateway be exclusively used by this edge gateway. This property is read-only for Provider gateways using IP Space. If the associated provider gateway is using IP Spaces, user can update the provider gateway itself to dedicate it to the Organization. Provider gateways using IP Spaces cannot be dedicated to an Edge Gateway.  | [optional] [default to null]
**UsingIpSpace** | **bool** | This fields is read-only and is set to true if the provider gateway backing the uplink (external network) is using IP Space; else, false.  | [optional] [default to null]
**VrfLiteBacked** | **bool** | Whether the associated external network is backed by a NSX-T VRF-Lite Tier-0. | [optional] [default to null]
**BackingType** | **string** | Type of backing for the network this interface is connected to. Can be NSXT_TIER0, NSXT_VRF_TIER0 or IMPORTED_T_LOGICAL_SWITCH | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


