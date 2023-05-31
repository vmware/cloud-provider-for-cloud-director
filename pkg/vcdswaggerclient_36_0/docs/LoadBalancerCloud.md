# LoadBalancerCloud

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**Id** | **string** | The identifier of the registered Load Balancer Cloud in URN format | [optional] [default to null]
**Name** | **string** | The name of the Load Balancer Cloud.  | [default to null]
**Description** | **string** | Description for the registered Load Balancer Cloud. | [optional] [default to null]
**LoadBalancerCloudBacking** | [***LoadBalancerCloudBacking**](LoadBalancerCloudBacking.md) | The backing that uniquely identifies a Load Balancer Cloud configured within a Load Balancer Controller.  At the present, VCD only supports NSX-T Clouds configured within an NSX-ALB Controller deployment. This is not updatable once it&#39;s created.  | [default to null]
**NetworkPoolRef** | [***EntityReference**](EntityReference.md) | The Network Pool associated with this Cloud | [optional] [default to null]
**HealthStatus** | **string** | The current health status of the Load Balancer Cloud. Possible values are: &lt;ul&gt; &lt;li&gt; UP - The cloud is healthy and ready to enable Load Balancer for an Edge Gateway. &lt;li&gt; DOWN - The cloud is in a failure state. Enabling Load balancer on an Edge Gateway may not be possible. &lt;li&gt; RUNNING - The cloud is currently processing. An example is if it&#39;s enabling a Load Balancer for an Edge Gateway. &lt;li&gt; UNAVAILABLE - The cloud is unavailable. &lt;li&gt; UNKNOWN - The cloud state is unknown. &lt;/ul&gt;  | [optional] [default to null]
**DetailedHealthMessage** | **string** | The non-localized detailed message on the health of the Cloud. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


