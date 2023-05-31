# FeatureFlag

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the feature flag in URN format. | [optional] [default to null]
**Name** | **string** | The unique name of the flag. | [optional] [default to null]
**Usage** | **string** | In what context this flag should be use. &lt;ul&gt;   &lt;li&gt; &lt;code&gt; PROD &lt;/code&gt; Indicates that this feature is production ready.   &lt;li&gt; &lt;code&gt; ALPHA &lt;/code&gt; Indicates that this feature is considered to be an alpha feature and may change in future releases. &lt;/ul&gt;  | [optional] [default to null]
**Enabled** | **bool** | True if the feature is toggled on. | [optional] [default to null]
**DisplayName** | **string** | Localized readable name of this feature. | [optional] [default to null]
**DisplayDescription** | **string** | Localized readable description of this feature. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


