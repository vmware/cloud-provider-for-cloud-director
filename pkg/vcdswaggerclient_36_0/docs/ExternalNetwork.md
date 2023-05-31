# ExternalNetwork

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**Id** | **string** | The unique ID for the network. This field is read-only. | [optional] [default to null]
**Name** | **string** | The name of the network. | [optional] [default to null]
**Description** | **string** | The description of the network. | [optional] [default to null]
**Subnets** | [***Subnets**](Subnets.md) | List of subnets configured for the network. | [optional] [default to null]
**NetworkBackings** | [***ExternalNetworkBackings**](ExternalNetworkBackings.md) | Backings for this external network. Describes if this external network is backed by port groups, vCenter standard switch or an NSX-T Tier-0 router.  | [optional] [default to null]
**TotalIpCount** | **int32** | The number of IP addresses defined by the static ip pools. If the network contains any IpV6 subnets, the total ip count will be null. | [optional] [default to null]
**UsedIpCount** | **int32** | The number of IP address used from the static ip pools. | [optional] [default to null]
**DedicatedOrg** | [***EntityReference**](EntityReference.md) | The Organization that this external network is dedicated to. This is null if this is not a dedicated external network. | [optional] [default to null]
**DedicatedEdgeGateway** | [***EntityReference**](EntityReference.md) | The Edge Gateway that this external network is dedicated to. This is null if this is not a dedicated external network. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


