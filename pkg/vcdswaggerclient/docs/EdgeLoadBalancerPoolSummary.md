# EdgeLoadBalancerPoolSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**Id** | **string** | The unique ID of this Load Balancer Pool. On updates, the ID is required for the pool, while for create a new ID will be generated. | [optional] [default to null]
**Description** | **string** | The description of the Load Balancer Pool. | [optional] [default to null]
**Enabled** | **bool** | True if Load Balancer Pool is enabled. | [optional] [default to null]
**PassiveMonitoringEnabled** | **bool** | Whether passive monitoring for this pool is enabled or not. | [optional] [default to null]
**HealthStatus** | **string** | The current health status of the pool. Possible values are: &lt;ul&gt; &lt;li&gt; UP - The pool is operational. &lt;li&gt; RUNNING - The pool is operational, but less than 50% of the pool members are up. &lt;li&gt; DOWN - All members in the pool are down. &lt;li&gt; DISABLED - Either the pool is disabled or all of the members are disabled. &lt;li&gt; UNAVAILABLE - The pool is unavailable. Examples: pool has no members or pool is not assigned to any virtual service. &lt;li&gt; UNKNOWN - The pool state is unknown. &lt;/ul&gt;  | [optional] [default to null]
**MemberCount** | **int32** | The total number of members in the pool. | [optional] [default to null]
**EnabledMemberCount** | **int32** | The number of enabled members in the pool. | [optional] [default to null]
**UpMemberCount** | **int32** | The number of enabled members in the pool that are operational. | [optional] [default to null]
**HealthMessage** | **string** | The localized message on the health of the pool. | [optional] [default to null]
**Name** | **string** | Name for the Load Balancer Pool. Name is unique across all pools for an Edge Gateway. | [optional] [default to null]
**ActiveMonitoringEnabled** | **bool** | Whether active monitoring for this pool is enabled or not. | [optional] [default to null]
**VirtualServiceRefs** | [**[]EntityReference**](EntityReference.md) | The list of Load Balancer Virtual Services associated with this Load balancer Pool. Only first 10 Virtual Services will be returned.  | [optional] [default to null]
**MemberSslEnabled** | **bool** | Whether SSL is enabled for communicatation between the Load Balancer Virtual Services and the pool members  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


