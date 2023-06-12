# ProbeResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Localized message describing the connection result stating success or an error message with a brief summary. | [optional] [default to null]
**ResolvedIp** | **string** | The IP address the host was resolved to. | [optional] [default to null]
**CanConnect** | **bool** | If vCD can establish a connection on the specified port. | [optional] [default to null]
**SslHandshake** | **bool** | If an SSL Handshake succeeded (secure requests only). | [optional] [default to null]
**CertificateChain** | **string** | The SSL certificate chain presented by the server if a secure connection was made. | [optional] [default to null]
**AdditionalCAIssuers** | **[]string** | URLs supplied by Certificate Authorities to retrieve signing certificates, when those certificates are not included in the chain. These URLs are the locations for the &#39;caIssuers&#39; access method in the &#39;Authority Info Access&#39; extension (as described in &lt;a href&#x3D;\&quot;https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.2.1\&quot;&gt;RFC 5280 Section 4.2.2.1&lt;/a&gt;) of the certificates and gives the caller an indication where additional CA certificates may be retrieved from, to complete the chain to the trust anchor.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


