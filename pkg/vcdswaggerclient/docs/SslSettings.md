# SslSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledSslProtocols** | **[]string** | SSL protocols | [optional] [default to null]
**EnabledSslCiphers** | **[]string** | SSL ciphers | [optional] [default to null]
**KeySize** | **int32** | Size of keys generated | [optional] [default to null]
**CertificateValidityDays** | **int32** | Number of days generated certificates are valid for | [optional] [default to null]
**CertificateSignatureAlgorithm** | **string** | Algorithm used to sign generated certificates | [optional] [default to null]
**FipsMode** | **string** | The desired FIPS mode of this server group. &lt;ul&gt;   &lt;li&gt;     &lt;em&gt;ON&lt;/em&gt;: Indicates FIPS mode is desired for this server group.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;OFF&lt;/em&gt;: Indicates FIPS mode is not desired for this server group.   &lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]
**Status** | **string** | The current SSL settings status for this server group. &lt;ul&gt;   &lt;li&gt;     &lt;em&gt;CURRENT&lt;/em&gt;: Indicates that all of the SSL settings for this server group are up to date.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;UPDATING&lt;/em&gt;: Indicates that there is at least one cell in the server group which has not yet     applied the SSL settings.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;AWAITING_RESTART&lt;/em&gt;: Indicates that the SSL settings for this server group have been applied on each cell,     and that each cell needs to be rebooted for the settings to take place. When each cell has been rebooted this field     will change to CURRENT.   &lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


