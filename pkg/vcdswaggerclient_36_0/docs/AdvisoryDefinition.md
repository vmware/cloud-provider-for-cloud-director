# AdvisoryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the advisory (read-only). | [optional] [default to null]
**TargetId** | **string** | The id reference to the target entity this advisory is for. | [optional] [default to null]
**Message** | **string** | A localized message for this advisory. | [optional] [default to null]
**Priority** | **string** | Priority for an advisory that indicates the level of urgency. These priorities are listed in ascending sort order. &lt;ul&gt;   &lt;li&gt;     &lt;em&gt;MANDATORY&lt;/em&gt;: A mandatory message which is always displayed;     these advisories cannot be snoozed or dismissed (see documentation     on displayStart and displayEnd)   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;CRITICAL&lt;/em&gt;: A high priority, potentially actionable message which can be     snoozed or dismissed   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;IMPORTANT&lt;/em&gt;: A potentially actionable message which can be snoozed or dismissed   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;NOTICE&lt;/em&gt;: An informational message which can be dismissed (but not snoozed)   &lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]
**DisplayStart** | [**time.Time**](time.Time.md) | The ISO-8601 timestamp from which this advisory is applicable. Defaults to the server&#39;s current time if unspecified. If permissible, users may update this value to a time in the future to snooze this advisory.  | [optional] [default to null]
**DisplayEnd** | [**time.Time**](time.Time.md) | The ISO-8601 timestamp representing when this advisory is no longer applicable. If permissible, users may update this value to a time in the past to dismiss this advisory. The displayEnd timestamp must be &gt;&#x3D; displayStart.  | [optional] [default to null]
**Source** | **string** | Represents where the advisory is being generated from. This is a read-only field. Can be of type USER or INTERNAL.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


