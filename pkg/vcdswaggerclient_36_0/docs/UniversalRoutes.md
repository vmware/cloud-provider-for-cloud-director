# UniversalRoutes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepAliveTimer** | **int32** | The Keep Alive Timer is the frequency (in seconds) at which the Universal Router seeds Keep Alive messages to its egress peers. | [optional] [default to null]
**Values** | [**[]UniversalRoute**](UniversalRoute.md) | The list of Universal Routes. | [optional] [default to null]
**ProviderScopesToForceUnconfigure** | **[]string** | The list network provider scopes whose route will be unconfigure forcefully. This means that a route that is in state that normally shouldn&#39;t be unconfigurable (i.e. NSX or remote vCD site is down) will still be unconfigured. Errors may be ignored. Note that this force option also affects any route update where a network provider&#39;s egress point is modified/changed. vCD will first unconfigure any existing egress point before configuring the new egress point.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


