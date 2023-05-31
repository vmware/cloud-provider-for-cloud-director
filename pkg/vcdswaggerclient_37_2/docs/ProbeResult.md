# ProbeResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Localized message describing the connection result stating success or an error message with a brief summary. | [optional] [default to null]
**ResolvedIp** | **string** | The IP address the host was resolved to, if not going through a proxy. | [optional] [default to null]
**CanConnect** | **bool** | If vCD can establish a connection on the specified port. | [optional] [default to null]
**SslHandshake** | **bool** | If an SSL Handshake succeeded (secure requests only). | [optional] [default to null]
**ConnectionResult** | **string** | A code describing the result of establishing a connection. Will be one of the following: &lt;ul&gt;   &lt;li&gt;     &lt;em&gt;SUCCESS&lt;/em&gt;: The connection was successful.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;ERROR_CANNOT_RESOLVE_IP&lt;/em&gt;: The hostname could not be resolved to an IP address.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;ERROR_CANNOT_CONNECT&lt;/em&gt;: Unable to establish connection.   &lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]
**SslResult** | **string** | A code describing the result of the SSL handshake. Will be one of the following: &lt;ul&gt;   &lt;li&gt;     &lt;em&gt;SUCCESS&lt;/em&gt;: The SSL handshake was successful.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;ERROR_SSL_ERROR&lt;/em&gt;: Unable to establish SSL connection.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;ERROR_UNTRUSTED_CERTIFICATE&lt;/em&gt;: Certificate not trusted.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;ERROR_CANNOT_VERIFY_HOSTNAME&lt;/em&gt;: Hostname verification failed.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;null&lt;/em&gt;: No SSL handshake was attempted.   &lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]
**CertificateChain** | **string** | The SSL certificate chain presented by the server if a secure connection was made. | [optional] [default to null]
**AdditionalCAIssuers** | **[]string** | URLs supplied by Certificate Authorities to retrieve signing certificates, when those certificates are not included in the chain. These URLs are the locations for the &#39;caIssuers&#39; access method in the &#39;Authority Info Access&#39; extension (as described in &lt;a href&#x3D;\&quot;https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.2.1\&quot;&gt;RFC 5280 Section 4.2.2.1&lt;/a&gt;) of the certificates and gives the caller an indication where additional CA certificates may be retrieved from, to complete the chain to the trust anchor.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


