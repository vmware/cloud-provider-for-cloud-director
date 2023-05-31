# IpSpaceSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**Id** | **string** | The identifier of the IP Space in URN format. | [optional] [default to null]
**Name** | **string** | The name of the IP Space. Name is unique across all IP Spaces of a given type and organization. | [default to null]
**Description** | **string** | The description of the IP Space. | [optional] [default to null]
**Type_** | **string** | The type of the IP Space. Possible values are: &lt;ul&gt; &lt;li&gt; PUBLIC - These can be consumed by multiple organizations. These are created by System Administrators only, for managing public IPs. The IP addresses and IP Prefixes from this IP space are allocated to specific organizations for consumption. &lt;li&gt; PRIVATE - These can be consumed by only a single organization. All the IPs within this IP Space are allocated to the particular organization. &lt;li&gt; SHARED_SERVICES - These are for internal use only. The IP addresses and IP Prefixes from this IP space can be consumed by multiple organizations but those IP addresses and IP Prefixes will not be not visible to the individual users within the organization. These are created by System Administrators only, typically for a service or for management networks. &lt;/ul&gt; Only &lt;em&gt;SHARED_SERVICES&lt;/em&gt; type can be changed to &lt;em&gt;PUBLIC&lt;/em&gt; type. No oher type changes are allowed.  | [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | The organization this IP Space belongs to. This property is only applicable for IP Spaces with type &lt;em&gt;PRIVATE&lt;/em&gt;. This property is required for IP Spaces with type &lt;em&gt;PRIVATE&lt;/em&gt;.  | [optional] [default to null]
**Utilization** | [***IpSpaceUtilizationSummary**](IpSpaceUtilizationSummary.md) | Utilization summary for this IP space. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


