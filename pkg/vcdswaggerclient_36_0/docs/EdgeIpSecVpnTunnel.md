# EdgeIpSecVpnTunnel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of this IPSec VPN tunnel. On updates, the id is required for the tunnel, while for create a new id will be generated.  | [optional] [default to null]
**Name** | **string** | Name for the tunnel. | [default to null]
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** | Described whether the tunnel is enabled or not. The default is true. | [optional] [default to null]
**LocalEndpoint** | [***LocalEndpoint**](LocalEndpoint.md) | The Local Endpoint which corresponds to the Edge Gateway the tunnel is being configured on. Local Endpoint requires an IP. That IP must be suballocated to the edge gateway.  | [default to null]
**RemoteEndpoint** | [***RemoteEndpoint**](RemoteEndpoint.md) | The Remote Endpoints correspoinds to the device on the remote site terminating the VPN tunnel.  | [default to null]
**AuthenticationMode** | **string** | The authentication mode this IPSec tunnel will use to authenticate with the peer endpoint. The default is a pre-shared key (PSK). &lt;ul&gt;   &lt;li&gt;PSK - A known key is shared between each site before the tunnel is established.&lt;/li&gt;   &lt;li&gt;CERTIFICATE ? Incoming connections are required to present an identifying digital certificate, which VCD verifies has been signed by a trusted certificate authority.&lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]
**PreSharedKey** | **string** | This is the Pre-shared key used for authentication. | [optional] [default to null]
**CertificateRef** | [***EntityReference**](EntityReference.md) | The server certificate which will be used to secure the tunnel&#39;s local endpoint. | [optional] [default to null]
**CaCertificateRef** | [***EntityReference**](EntityReference.md) | The CA authority used to verify the remote endpoint?s certificate. | [optional] [default to null]
**ConnectorInitiationMode** | [***ConnectorInitiationMode**](ConnectorInitiationMode.md) | This is the mode used by the local endpoint to establish an IKE Connection with the remote site. The default is INITIATOR. | [optional] [default to null]
**SecurityType** | **string** | This is the security type used for the IPSec Tunnel. If nothing is specified, this will be set to &#39;DEFAULT&#39; in which the default settings in NSX will be used. For custom settings, one should use the connectionProperties endpoint to specify custom settings. The security type will then appropriately reflect itself as &#39;CUSTOM&#39;.  | [optional] [default to null]
**Logging** | **bool** | Whether logging for the tunnel is enabled or not. The default is false. | [optional] [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


