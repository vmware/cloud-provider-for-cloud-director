# EdgeBgpGracefulRestartConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [***GracefulRestartModeTypes**](GracefulRestartModeTypes.md) | Currently configured graceful restart mode. Default is HELPER_ONLY. | [optional] [default to null]
**RestartTimer** | **int32** | Maximum time taken (in seconds) for a BGP session to be established after a restart. If the session is not re-established within this timer, the receiving speaker will delete all the stale routes from that peer.  | [optional] [default to 180]
**StaleRouteTimer** | **int32** | Maximum time (in seconds) before stale routes are removed when BGP restarts. | [optional] [default to 600]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


