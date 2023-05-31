# EdgeBgpNeighbor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of this BGP neighbor. On updates, the id is required for the object, while for create a new id will be generated. This id is not a VCD URN.  | [optional] [default to null]
**NeighborAddress** | **string** | The IP address of the BGP neighbor. Both IPv4 and IPv6 formats are supported. | [default to null]
**RemoteASNumber** | **string** | The remote AS number of a BGP neighbor in ASPLAIN format. | [default to null]
**KeepAliveTimer** | **int32** | Specifies the time interval (in seconds) between keep alive messages sent to peer. | [optional] [default to 60]
**HoldDownTimer** | **int32** | Specifies the time interval (in seconds) before declaring a peer dead. | [optional] [default to 180]
**GracefulRestartMode** | [***GracefulRestartModeTypes**](GracefulRestartModeTypes.md) | Currently configured graceful restart configuration mode. Default is HELPER_ONLY. | [optional] [default to null]
**Bfd** | [***EdgeBgpBfdConfig**](EdgeBgpBfdConfig.md) | Specifies the BFD configuration for failure detection. Not specifying a value results in default bahavior.  | [optional] [default to null]
**AllowASIn** | **bool** | A flag indicating whether AllowAS-in is enabled or not. This specifies whether BGP neighbors can receive routes with same AS.  | [optional] [default to null]
**NeighborPassword** | **string** | Password for BGP neighbor authentication. Empty string (\&quot;\&quot;) clears existing password. Not specifying a value will be treated as \&quot;no password\&quot;.  | [optional] [default to null]
**IpAddressTypeFiltering** | **string** | Specifies IP address type based filtering in each direction. Setting the value to &#39;DISABLED&#39; will disable address family based filtering.  | [optional] [default to null]
**InRoutesFilterRef** | [***ExtObjectReference**](ExtObjectReference.md) | Specifies route filtering configuration for the BGP neighbor in IN direction. It is the reference to the prefix list, indicating which routes to filter for IN direction. Not specifying a value will be treated as \&quot;no IN route filters\&quot;.  | [optional] [default to null]
**OutRoutesFilterRef** | [***ExtObjectReference**](ExtObjectReference.md) | Specifies route filtering configuration for the BGP neighbor in OUT direction. It is the reference to the prefix list, indicating which routes to filter for OUT direction. Not specifying a value will be treated as \&quot;no OUT route filters\&quot;.  | [optional] [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


