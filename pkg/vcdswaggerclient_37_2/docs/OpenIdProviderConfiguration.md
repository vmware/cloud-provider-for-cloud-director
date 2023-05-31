# OpenIdProviderConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveKey** | [***EntityReference**](EntityReference.md) | Entity reference of the active OpenID Provider key, which will be presented in the jwks endpoint. | [default to null]
**AllowHttp** | **bool** | Determines whether relying party redirect URIs may be http | [default to null]
**WellKnownEndpoint** | **string** | OpenID Provider endpoint serving provider configuration values (as described in &lt;a href&#x3D;\&quot;https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfigurationRequest\&quot;&gt;OpenId Spec 4.1&lt;/a&gt;)  | [default to null]
**AuthorizationCodeTimeToLive** | **int64** | Authorization Code time to live (ttl) in seconds | [default to null]
**AccessTokenTimeToLive** | **int64** | Access Token time to live (ttl) in seconds | [default to null]
**IdTokenTimeToLive** | **int64** | ID Token time to live (ttl) in seconds | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


