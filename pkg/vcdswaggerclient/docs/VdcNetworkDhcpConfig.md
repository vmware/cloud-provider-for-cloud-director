# VdcNetworkDhcpConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the DHCP service is currently enabled on network. | [optional] [default to null]
**LeaseTime** | **int64** | The amount of time in seconds of how long a DHCP IP will be leased out for. The minimum is 60s while the maximum is 4,294,967,295s, which is roughly 49,710 days.  | [optional] [default to 86400]
**DhcpPools** | [**[]VdcNetworkDhcpPool**](VdcNetworkDhcpPool.md) | Range of DHCP IP addresses | [optional] [default to null]
**Mode** | **string** | This value describes how the DHCP service is configured for this network. Once a DHCP service has been created, the mode attribute cannot be changed. The mode field will default to &#39;EDGE&#39; if it is not provided. This field only applies to networks backed by an NSX-T network provider. &lt;ul&gt; &lt;li&gt;The supported values are EDGE and NETWORK.&lt;/li&gt; &lt;li&gt;If EDGE is specified, the DHCP service of the edge is used to obtain DHCP IPs.&lt;/li&gt; &lt;li&gt;If NETWORK is specified, a DHCP server is created for use by this network.&lt;/li&gt; &lt;/ul&gt; In order to use DHCP for IPV6, NETWORK mode must be used. Routed networks which are using NETWORK DHCP services can be disconnected from the edge gateway and still retain their DHCP configuration, however network using EDGE DHCP cannot be disconnected from the gateway until DHCP has been disabled.  | [optional] [default to null]
**IpAddress** | **string** | The IP address of the DHCP service. This is required upon create if using NETWORK mode. This field only applies to networks backed by an NSX-T network provider.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


