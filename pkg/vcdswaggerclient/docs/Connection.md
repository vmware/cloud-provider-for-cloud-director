# Connection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The host (or IP address) to connect to. | [default to null]
**Port** | **int32** | The port to use when connecting. | [default to null]
**Secure** | **bool** | If the connection should use https. | [optional] [default to null]
**Timeout** | **int32** | Maximum time (in seconds) any step in the test should wait for a response. | [optional] [default to 10]
**HostnameVerificationAlgorithm** | **string** | Endpoint/Hostname verification algorithm to be used during SSL/TLS/DTLS handshake. Their values are as follows: &lt;ul&gt;   &lt;li&gt;     &lt;em&gt;HTTPS&lt;/em&gt;: use https hostname verification algorithm as described in     &lt;a href&#x3D;\&quot;https://datatracker.ietf.org/doc/html/rfc2818\&quot;&gt;RFC 2818&lt;/a&gt;   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;LDAPS&lt;/em&gt;: use ldaps hostname verification algorithm as described in     &lt;a href&#x3D;\&quot;https://datatracker.ietf.org/doc/html/rfc2830\&quot;&gt;RFC 2830&lt;/a&gt;   &lt;/li&gt; &lt;/ul&gt; When this field is not set, the default value &lt;em&gt;null&lt;/em&gt; indicates no hostname verification will be performed.  | [optional] [default to null]
**AdditionalCAIssuers** | **[]string** | A list of URLs being authorized by the user to retrieve additional CA certificates from,  if necessary, to complete the certificate chain to its trust anchor. &lt;p&gt; Upon retrieving the certificate chain presented during the handshake, if signing CA certificates were not included, but a location is specified for the &#39;caIssuers&#39; access method of the &#39;Authority Info Access&#39; extension  (as described in &lt;a href&#x3D;\&quot;https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.2.1\&quot;&gt;RFC 5280 Section 4.2.2.1&lt;/a&gt;) of those certificates and that location is one of these supplied URLs, then additional certificates will be retrieved from those URLs in accordance with the protocol laid out in the RFC. &lt;p&gt; Any failure to retrieve this certificate will NOT fail the test connection request, nor will the error associated with this failure be returned. &lt;p&gt; In the unlikely event that the CA Issuers URL specifies &#x60;https&#x60; instead of &#x60;http&#x60;, the original certificate, that included that URL, will be temporarily used to trust the server during ssl handshake  | [optional] [default to null]
**ProxyConnection** | [***ProxyConnection**](ProxyConnection.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


