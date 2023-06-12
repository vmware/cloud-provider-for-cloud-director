# NetworkContextProfileAttribute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | This describes the type of attribute value. &lt;ul&gt;   &lt;li&gt; APP_ID - Values represents layer 7 App Ids. For example: ACTIVDIR   &lt;li&gt; DOMAIN_NAME - Values represents Domain names (FQDN). For example: *.live.com &lt;/ul&gt;  | [default to null]
**Values** | **[]string** | Values for attribute. | [default to null]
**SubAttributes** | [**[]NetworkContextProfileSubAttribute**](NetworkContextProfileSubAttribute.md) | List of sub attributes for an attribute. These are specified with the attributes such as SSL or CIFS, which can have different cipher suites or TLS versions as values.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


