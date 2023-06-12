# OidcRelyingParty

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the relying party. | [optional] [default to null]
**ClientId** | **string** | System generated client id of the relying party as per RFC-7591 Section 3.2.1.  | [optional] [default to null]
**ClientName** | **string** | Human readable name of the relying party.  | [default to null]
**ClientSecret** | **string** | Server generated client secret string. Must be unique for all relying parties. This field is hidden and is only returned in plaintext on a POST (during registration).  | [optional] [default to null]
**RedirectUris** | **[]string** | Supported redirect URIs for this relying party.  | [default to null]
**Scope** | **[]string** | Not configurable by the client. A fixed list of the following six scope values: &lt;ul&gt;   &lt;li&gt; openid - as required per &lt;a href&#x3D;\&quot;https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest\&quot;&gt;the OpenID Connect Core spec&lt;/a&gt;   &lt;li&gt; profile - as described in &lt;a href&#x3D;\&quot;https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims\&quot;&gt;the OpenID Connect Core spec&lt;/a&gt;   &lt;li&gt; email - as described in &lt;a href&#x3D;\&quot;https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims\&quot;&gt;the OpenID Connect Core spec&lt;/a&gt;   &lt;li&gt; phone - as described in &lt;a href&#x3D;\&quot;https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims\&quot;&gt;the OpenID Connect Core spec&lt;/a&gt;   &lt;li&gt; groups - grants access to the groups claims.   &lt;li&gt; vcd_idp - grants access to the roles, org_id, org_display_name, and org_name claims. &lt;/ul&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


