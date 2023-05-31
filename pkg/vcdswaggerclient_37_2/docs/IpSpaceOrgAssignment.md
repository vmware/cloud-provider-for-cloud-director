# IpSpaceOrgAssignment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of the IP Space Org Assignment in URN format. | [optional] [default to null]
**IpSpaceRef** | [***EntityReference**](EntityReference.md) | Reference to the associated IP Space. This property is not updatable. | [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | Reference to the associated Organization. This property is not updatable. | [default to null]
**IpSpaceType** | **string** | The type of the IP Space. Possible values are: &lt;ul&gt; &lt;li&gt; PUBLIC &lt;li&gt; PRIVATE &lt;li&gt; SHARED_SERVICES &lt;/ul&gt;  | [optional] [default to null]
**DefaultQuotas** | [***IpSpaceIpQuota**](IpSpaceIpQuota.md) | Default IP Space quotas applied to the Organization. | [optional] [default to null]
**CustomQuotas** | [***IpSpaceIpQuota**](IpSpaceIpQuota.md) | Custom IP Space quotas applied to the Organization. A custom quota for floating IPs or IP Prefix overrides the default quotas applied to the associated Organization. Setting the customQuotas property to NULL will result in resetting all the overridden quotas to their respective defaults. Setting the specific overridden quota value within customQuotas will result in resetting the overridden quota value to its default value. A Quota of -1 means there is no cap to the number of IP addresses or IP Prefixes that can be allocated. A Quota of 0 means that the IP addresses or IP Prefixes cannot be allocated.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


