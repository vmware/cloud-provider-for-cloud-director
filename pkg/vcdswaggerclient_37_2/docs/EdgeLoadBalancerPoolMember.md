# EdgeLoadBalancerPoolMember

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | The ip address of the Load Balancer Pool member. | [default to null]
**Port** | **int32** | The port number of the Load Balancer Pool member. If unset, the port that the client used to connect will be used. | [optional] [default to null]
**Ratio** | **int32** | The ratio of selecting eligible servers in the pool. | [optional] [default to 1]
**Enabled** | **bool** | Whether the Load Balancer Pool member is enabled or not. | [optional] [default to null]
**HealthStatus** | **string** | The current health status of the pool member. Possible values are: &lt;ul&gt; &lt;li&gt; UP - The member is operational. &lt;li&gt; DOWN - The member is down. &lt;li&gt; DISABLED - The member is disabled &lt;li&gt; UNKNOWN - The state is unknown. &lt;/ul&gt;  | [optional] [default to null]
**MarkedDownBy** | **[]string** | When the member is DOWN, the value gives the names of the health monitors that marked the member as down. If a monitor cannot be determined, the value will be UNKNOWN.  | [optional] [default to null]
**DetailedHealthMessage** | **string** | The non-localized detailed message on the health of the pool member.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


