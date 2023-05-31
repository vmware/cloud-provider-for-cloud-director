# EdgeLoadBalancerPool

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
**Name** | **string** | Name for the Load Balancer Pool. Name is unique across all pools for an Edge Gateway. | [default to null]
**DefaultPort** | **int32** | The destination server port used by the traffic sent to the member. | [optional] [default to 80]
**GracefulTimeoutPeriod** | **int32** | Maximum time (in minutes) to gracefully disable a member. Virtual service waits for the specified time before terminating the existing connections to the members that are disabled. &lt;code&gt;Special values: 0 represents &#39;Immediate&#39;, -1 represents &#39;Infinite&#39;.&lt;/code&gt;  | [optional] [default to 1]
**Algorithm** | **string** | The algorithm for choosing a member within the pool&#39;s list of available members for each new connection. Default value is \&quot;LEAST_CONNECTIONS\&quot;. Supported algorithms are: &lt;ul&gt; &lt;li&gt;LEAST_CONNECTIONS &lt;li&gt;ROUND_ROBIN &lt;li&gt;CONSISTENT_HASH &lt;li&gt;FASTEST_RESPONSE &lt;li&gt;LEAST_LOAD &lt;li&gt;FEWEST_SERVERS &lt;li&gt;RANDOM &lt;li&gt;FEWEST_TASKS &lt;li&gt;CORE_AFFINITY &lt;/ul&gt; &lt;em&gt;CONSISTENT_HASH&lt;/em&gt; uses Source IP Address hash. Using &lt;em&gt;FASTEST_RESPONSE&lt;/em&gt;, &lt;em&gt;LEAST_LOAD&lt;/em&gt;, &lt;em&gt;FEWEST_SERVERS&lt;/em&gt;, &lt;em&gt;RANDOM&lt;/em&gt;, &lt;em&gt;FEWEST_TASKS&lt;/em&gt;, &lt;em&gt;CORE_AFFINITY&lt;/em&gt; algorithms may require additional licensing.  | [optional] [default to null]
**HealthMonitors** | [**[]EdgeLoadBalancerHealthMonitor**](EdgeLoadBalancerHealthMonitor.md) | Member server&#39;s health can be monitored by using one or more health monitors. Active monitors generate synthetic traffic and mark a server up or down based on the response.  | [optional] [default to null]
**PersistenceProfile** | [***EdgeLoadBalancerPersistenceProfile**](EdgeLoadBalancerPersistenceProfile.md) | Selected persistence profile for the Load Balancer Pool. | [optional] [default to null]
**Members** | [**[]EdgeLoadBalancerPoolMember**](EdgeLoadBalancerPoolMember.md) | The list of destination servers which are used by the Load Balancer Pool to direct load balanced traffic.  | [optional] [default to null]
**VirtualServiceRefs** | [**[]EntityReference**](EntityReference.md) | The list of Load Balancer Virtual Services associated with this Load balancer Pool. | [optional] [default to null]
**GatewayRef** | [***EntityReference**](EntityReference.md) | The Edge Gateway associated with this Load Balancer Pool. | [default to null]
**CaCertificateRefs** | [**[]EntityReference**](EntityReference.md) | The root certificates to use when validating certificates presented by the pool members. | [optional] [default to null]
**CommonNameCheckEnabled** | **bool** | Whether to check the common name of the certificate presented by the pool member. This cannot be enabled if no caCertificateRefs are specified.  | [optional] [default to null]
**DomainNames** | **[]string** | A list of domain names which will be used to verify the common names or subject alternative names presented by the pool member certificates. It is performed only when common name check (commonNameCheckEnabled) is enabled. If common name check is enabled, but domain names are not specified then the incoming host header will be used to check the certificate.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


