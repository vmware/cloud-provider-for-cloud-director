# BrandingThemeResourceCustomLink

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the custom link. The Id is provided by the client in the CustomLinks file.  | [optional] [default to null]
**Name** | **string** | The string used to display in the UI portal this custom link.  | [default to null]
**LinkType** | **string** | Type of the custom link. &lt;ul&gt;   &lt;li&gt;     &lt;em&gt;override&lt;/em&gt;- Override existing/pre-built custom item  (well known links that can be overridden, Default values are About, Help, VMRC)   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;link&lt;/em&gt; - Represent an custom link   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;separator&lt;/em&gt; - Used as a visual separator between the links/sections   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;section&lt;/em&gt; - Used as a group of links. Use parentId to group children links.   &lt;/li&gt; &lt;/ul&gt;  | [default to null]
**Scope** | **string** | Defines the place where the link should appear in the UI portal. &lt;ul&gt;   &lt;li&gt;     &lt;em&gt;LOGIN&lt;/em&gt; - Custom link to be shown on the login screen   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;LOGOUT&lt;/em&gt;- Custom link to be shown on the logout screen   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;PORTAL&lt;/em&gt;- Custom link to be shown within the UI portal after authorization   &lt;/li&gt; &lt;/ul&gt;  | [default to null]
**Url** | **string** | The URL to be used in the UI portal to navigate user. To be used when linkType is either &lt;em&gt;OVERRIDE&lt;/em&gt; or &lt;em&gt;LINK&lt;/em&gt;, otherwise it will be ignored.  | [optional] [default to null]
**Order** | **float32** | Used to order the links during display in the UI portal. Links with the same order should be ordered by name  | [optional] [default to null]
**ParentId** | **string** | Id of parent Custom link of type GROUP. Custom links that share a parentId will be grouped together. The Id is provided by the client in the CustomLinks file.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


