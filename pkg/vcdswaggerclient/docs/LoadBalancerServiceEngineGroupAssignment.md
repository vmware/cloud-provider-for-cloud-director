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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


