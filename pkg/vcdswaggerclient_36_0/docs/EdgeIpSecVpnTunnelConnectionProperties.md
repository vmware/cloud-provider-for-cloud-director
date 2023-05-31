# EdgeIpSecVpnTunnelConnectionProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityType** | **string** | This is the security type used for the IPSec Tunnel. If nothing is specified, this will be set to &#39;DEFAULT&#39; in which the default settings in NSX will be used. If &#39;CUSTOM&#39; is specified, then ike, tunnel, and dpd configurations can be set.  | [optional] [default to null]
**IkeConfiguration** | [***IkeConfig**](IkeConfig.md) | The IKE Configuration to be used for the tunnel. If nothing is explictly set, the system defaults will be used.  | [optional] [default to null]
**TunnelConfiguration** | [***TunnelConfig**](TunnelConfig.md) | The Tunnel Configuration, which contains parameters such as encryption algorithm to be used. If nothing is explicitly set, the system defaults will be used.  | [optional] [default to null]
**DpdConfiguration** | [***DpdConfig**](DpdConfig.md) | The Dead Peer Detection configuration. If nothing is explictly set, the system defaults will be used.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


