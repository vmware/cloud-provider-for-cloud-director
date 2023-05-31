# IpSpaceIpAllocation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of the IP Space IP allocation in URN format. | [optional] [default to null]
**Type_** | **string** | The type of the IP allocation. Possible values are: &lt;ul&gt; &lt;li&gt; FLOATING_IP - For allocation of floating IP addresses from defined IP Space ranges. &lt;li&gt; IP_PREFIX - For allocation of IP prefix sequences from defined IP Space prefixes. &lt;/ul&gt;  | [optional] [default to null]
**Value** | **string** | An individual IP Address or an IP Prefix which is allocated. | [optional] [default to null]
**AllocationDate** | [**time.Time**](time.Time.md) | Date when the IP address or IP prefix is allocated. This property is read-only. | [optional] [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | Reference to the organization where the IP is allocated. | [optional] [default to null]
**UsageState** | **string** | Specifies current usage state of an allocated IP. Possible values are: &lt;ul&gt; &lt;li&gt; UNUSED - the allocated IP is current not being used in the system. &lt;li&gt; USED - the allocated IP is currently in use in the system. An allocated IP address or IP Prefix is considered used if it is being used in network services such as NAT rule or in Org VDC network definition. &lt;li&gt; USED_MANUAL - the allocated IP is marked for manual usage. Allocation description can be referenced to get information about the manual usage. &lt;/ul&gt;  | [optional] [default to null]
**Description** | **string** | Description about usage of an IP if the usageState is USED_MANUAL. | [optional] [default to null]
**UsedByRef** | [***EntityReference**](EntityReference.md) | Reference to the entity using the IP, such as an Edge Gateway Reference if the Floating IP is used for NAT or Org VDC network reference if IP Prefix is used for network definition. This property is read-only.  | [optional] [default to null]
**UsageCategories** | **[]string** | The list of service categories where the IP address is being used. Typically this can be one of: SNAT, DNAT, LOAD_BALANCER, IPSEC_VPN, SSL_VPN or L2_VPN. This property is read-only.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


