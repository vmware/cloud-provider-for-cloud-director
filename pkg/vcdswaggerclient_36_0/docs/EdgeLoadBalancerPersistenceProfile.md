# EdgeLoadBalancerPersistenceProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Persistence profile.  | [optional] [default to null]
**Type_** | **string** | Type of persistence strategy to use. Supported values are: &lt;ul&gt; &lt;li&gt;CLIENT_IP - The client?s IP is used as the identifier and mapped to the server. &lt;li&gt;HTTP_COOKIE - Load Balancer inserts a cookie into HTTP responses. Cookie name must be provided as value. &lt;li&gt;CUSTOM_HTTP_HEADER - Custom, static mappings of header values to specific servers are used. Header name must be provided as value. &lt;li&gt;APP_COOKIE - Load Balancer reads existing server cookies or URI embedded data such as JSessionID. Cookie name must be provided as value. &lt;li&gt;TLS - Information is embedded in the client?s SSL/TLS ticket ID. This will use default system profile \&quot;System-Persistence-TLS\&quot;. &lt;/ul&gt; Using &lt;em&gt;CUSTOM_HTTP_HEADER&lt;/em&gt;, &lt;em&gt;APP_COOKIE&lt;/em&gt;, &lt;em&gt;TLS&lt;/em&gt; persistence types may require additional licensing.  | [default to null]
**Value** | **string** | Value of attribute based on selected persistence type. This is required for HTTP_COOKIE, CUSTOM_HTTP_HEADER and APP_COOKIE persistence types. HTTP_COOKIE, APP_COOKIE must have cookie name set as the value and CUSTOM_HTTP_HEADER must have header name set as the value.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


