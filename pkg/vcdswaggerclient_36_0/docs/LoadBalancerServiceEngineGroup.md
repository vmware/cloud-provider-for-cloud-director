# LoadBalancerServiceEngineGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**Id** | **string** | The identifier of the Load Balancer Service Engine Groups in URN format | [optional] [default to null]
**Name** | **string** | The name of the Load Balancer Service Engine Group.  | [default to null]
**Description** | **string** | Description of the Load Balancer Service Engine Group. | [optional] [default to null]
**ServiceEngineGroupBacking** | [***LoadBalancerServiceEngineGroupBacking**](LoadBalancerServiceEngineGroupBacking.md) | The backing that uniquely identifies a Load Balancer Service Engine Group configured within a Load Balancer Cloud.  | [default to null]
**HaMode** | **string** | The service engine group&#39;s High Availability Mode. &lt;ul&gt; &lt;li&gt;ELASTIC_N_PLUS_M_BUFFER - Service Engines will scale out to N active nodes with M nodes as buffer. &lt;li&gt;ELASTIC_ACTIVE_ACTIVE - Active-Active with scale out. &lt;li&gt;LEGACY_ACTIVE_STANDBY - Traditional single Active-Standby configuration &lt;/ul&gt;  | [optional] [default to null]
**ReservationType** | **string** | The reservation model for virutal services on the Load Balancer Service Engine Group. &lt;ul&gt; &lt;li&gt;DEDICATED - Dedicated to a single Edge Gateway and can only be assigned to a single Edge Gateway. &lt;li&gt;SHARED - Shared between multiple Edge Gateways. Can be assigned to multiple Edge Gateways. &lt;/ul&gt;  | [optional] [default to null]
**MaxVirtualServices** | **int32** | The maximum number of virtual services supported on the Load Balancer Service Engine Group.  | [optional] [default to null]
**NumDeployedVirtualServices** | **int32** | The number of virtual services currently deployed on the Load Balancer Service Engine Group.  | [optional] [default to null]
**ReservedVirtualServices** | **int32** | The number of virtual services already reserved on the Load Balancer Service Engine Group. This value is the sum of the guaranteed virtual services given to Edge Gateways assigned to the Load Balancer Service Engine Group.  | [optional] [default to null]
**OverAllocated** | **bool** | Indicates whether the maximum number of virtual services supported on the Load Balancer Service Engine Group has been surpassed by the current number of reserved virtual services.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


