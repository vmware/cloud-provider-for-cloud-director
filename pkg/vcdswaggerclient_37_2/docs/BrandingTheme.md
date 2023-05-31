# BrandingTheme

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the branding theme (read-only). | [optional] [default to null]
**Name** | **string** | Display name for the Branding Theme  | [default to null]
**ThemeType** | **string** | Type of the Branding Theme. Supported theme types are: &lt;ul&gt;   &lt;li&gt;     &lt;em&gt;BUILT_IN&lt;/em&gt; - Theme pre-bundled with Cloud Director   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;CUSTOM&lt;/em&gt; - Custom Branding Theme   &lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]
**IsLegacy** | **bool** | If true, the branding theme is represented in old format that can not be edited.  | [optional] [default to null]
**IsDefault** | **bool** | If true, the branding theme is used by any organization with no explicit theme assignment.  | [optional] [default to null]
**Active** | **bool** | If true, the branding theme is used for branding customization of the organization. Only one branding theme can be active for organization.  | [optional] [default to null]
**ThemeBase** | **string** | This indicates whether the theme css is based of the light or dark css base. Missing value indicates the base is unknown. Enum options - CLARITY_LIGHT_THEME, CLARITY_DARK_THEME.  | [optional] [default to null]
**LogoThumbprint** | **string** | base64 encoded Logo thumbprint generated from the provided logo resource. | [optional] [default to null]
**Version** | **string** | Version of the theme  | [optional] [default to null]
**PreviewStyles** | **string** | Subset of the provided custom resource styles. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


