# LoadBalancerServiceEngineGroupAssignment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of the Load Balancer Service Engine Groups in URN format. | [optional] [default to null]
**MaxVirtualServices** | **int32** | The maximum number of virtual services the Edge Gateway is allowed to use. This is required if the Load Balancer Service Engine Group has reservation type &#39;SHARED&#39;. This must be unset if the Load Balancer Service Engine Group has reservation type &#39;DEDICATED&#39;.  | [optional] [default to null]
**MinVirtualServices** | **int32** | The number of guaranteed virtual services available to the Edge Gateway. This is required if the Load Balancer Service Engine Group has reservation type &#39;SHARED&#39;. This must be unset if the Load Balancer Service Engine Group has reservation type &#39;DEDICATED&#39;.  | [optional] [default to null]
**NumDeployedVirtualServices** | **int32** | The current number of deployed virutal services.  | [optional] [default to null]
**ServiceEngineGroupRef** | [***EntityReference**](EntityReference.md) | The associated Load Balancer Service Engine Group. | [default to null]
**GatewayRef** | [***EntityReference**](EntityReference.md) | The associated Edge Gateway. | [default to null]
**GatewayOwnerRef** | [***EntityReference**](EntityReference.md) | The owner of the associated Edge Gateway. This can be a vDC or vDC Group. | [optional] [default to null]
**GatewayOrgRef** | [***EntityReference**](EntityReference.md) | The organization of the associated Edge Gateway. | [optional] [default to null]
**ServiceEngineGroupSupportedFeatureSet** | **string** | The feature set supported by the Load Balancer Service Engine Group. | [optional] [default to null]
**GatewaySupportedFeatureSet** | **string** | The feature set supported by the Edge Gateway Load Balancer. This will be NULL if Load Balancer is not active on the associated Edge Gateway.  | [optional] [default to null]
**ServiceEngineGroupHaMode** | **string** | The service engine group&#39;s High Availability Mode. &lt;ul&gt; &lt;li&gt;ELASTIC_N_PLUS_M_BUFFER - Service Engines will scale out to N active nodes with M nodes as buffer. &lt;li&gt;ELASTIC_ACTIVE_ACTIVE - Active-Active with scale out. &lt;li&gt;LEGACY_ACTIVE_STANDBY - Traditional single Active-Standby configuration &lt;/ul&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


