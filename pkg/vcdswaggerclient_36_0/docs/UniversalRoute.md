# UniversalRoute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultEgress** | [***EgressPointReference**](EgressPointReference.md) | For a given route, routing by default will go through the default egress point. Every valid/configured route will have a default egress point. | [optional] [default to null]
**StandbyEgress** | [***EgressPointReference**](EgressPointReference.md) | A standby egress point can be configured for failover of routes in case there is an issue routign to the default egress point. | [optional] [default to null]
**NetworkProviderScope** | **string** | The network provider scope for the given route. There can only be one route per fault domain for a Universal Route. | [default to null]
**Status** | [***UniversalRoutingStatus**](UniversalRoutingStatus.md) | The status of the route (whether it&#39;s realized, failed, etc.). | [optional] [default to null]
**ErrorMessage** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


