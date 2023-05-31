# EdgeStaticRoute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of this static route. On updates, the ID is required for the object, while for create a new ID will be generated. This ID is not a VCD URN.  | [optional] [default to null]
**Name** | **string** | Name for the static route. | [default to null]
**Description** | **string** | Description for this static route. | [optional] [default to null]
**NetworkCidr** | **string** | The network prefix in CIDR format. Both IPv4 and IPv6 formats are supported.  | [default to null]
**NextHops** | [**[]EdgeStaticRouteNextHop**](EdgeStaticRouteNextHop.md) | The list of next hops to use within the static route. List must contain at least one valid next hop.  | [default to null]
**SystemOwned** | **bool** | A flag indicating whether this static route is managed by the system. Property is read-only. | [optional] [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


