# EdgeBgpConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A flag indicating whether BGP configuration is enabled or not. | [optional] [default to null]
**Ecmp** | **bool** | A flag indicating whether ECMP is enabled or not. | [optional] [default to null]
**LocalASNumber** | **string** | BGP AS number to advertise to BGP peers. BGP AS number can be specified in either ASPLAIN or ASDOT formats, like ASPLAIN format :- &#39;65546&#39;, ASDOT format :- &#39;1.10&#39;. Read only if using a VRF-Lite backed external network.  | [optional] [default to null]
**GracefulRestart** | [***EdgeBgpGracefulRestartConfig**](EdgeBgpGracefulRestartConfig.md) | BGP Graceful Restart configuration. Not specifying a value results in default bahavior. Read only if using a VRF-Lite backed external network.  | [optional] [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


