# EdgeLoadBalancerVirtualServiceSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**Id** | **string** | The identifier of the Virtual Service in URN format | [optional] [default to null]
**Name** | **string** | The name of the Virtual Service. Name is unique across all Virtual Services for an Edge Gateway. | [default to null]
**Description** | **string** | The description of the Virtual Service. | [optional] [default to null]
**Enabled** | **bool** | A flag indicating whether Virtual Service is enabled or not. | [default to null]
**VirtualIpAddress** | **string** | The virtual IP Address (VIP) of the Virtual Service. This IP can be an allocated IP to the Gateway from the External Network or it can be an arbitrary internal IP address used for internal load balancing. It it&#39;s an internal IP Address, this IP cannot be part of any existing subnet attached to the Edge Gateway or any vDC Group network if the Edge Gateway is scoped accordingly.  | [default to null]
**LoadBalancerPoolRef** | [***EntityReference**](EntityReference.md) | The Load Balancer Pool associated with this Virtual Service. | [default to null]
**GatewayRef** | [***EntityReference**](EntityReference.md) | The Edge Gateway associated with this Virtual Service. | [default to null]
**ServiceEngineGroupRef** | [***EntityReference**](EntityReference.md) | The Load Balancer Service Engine Group that is assigned to the Edge Gateway. This Virtual Service will be deployed to this Service Engine Group.  | [default to null]
**CertificateRef** | [***EntityReference**](EntityReference.md) | The certificate used for SSL termination for the Virtual Service. This is required if the service port type is \&quot;HTTPS\&quot; or \&quot;L4_TLS\&quot;. | [optional] [default to null]
**ServicePorts** | [**[]EdgeLoadBalancerServicePort**](EdgeLoadBalancerServicePort.md) | A list of service ports supported by this Virtual Service.  Multiple service ports are allowed only with additional licensing.  | [default to null]
**HealthStatus** | **string** | The current health status of the virtual service. Possible values are: &lt;ul&gt; &lt;li&gt; UP - The virtual service is healthy. &lt;li&gt; DOWN - The virtual service is down, inactive, or has failed. &lt;li&gt; DISABLED - The virtual service is disabled. &lt;li&gt; UNAVAILABLE - The virtual service is unavailable. &lt;li&gt; PENDING - The virtual service is being creating or resources are being allocated. &lt;li&gt; UNKNOWN - The virtual service state is unknown. &lt;/ul&gt;  | [optional] [default to null]
**HealthMessage** | **string** | The localized message on the health of the virtual service. | [optional] [default to null]
**DetailedHealthMessage** | **string** | The non-localized detailed message on the health of the virtual service. | [optional] [default to null]
**ApplicationProfileType** | **string** | The profile type of application that this Virtual Service is configured with. A value of \&quot;-\&quot; represents an unknown type. &lt;ul&gt; &lt;li&gt;HTTP - Virtual Service supports HTTP protocol. &lt;li&gt;HTTPS - Virtual Service supports HTTPS protocol. &lt;li&gt;L4 - Virtual Service supports Layer 4 (Transport) using UDP/TCP protocol. &lt;li&gt;L4_TLS - Virtual Service supports Layer 4 (Transport) using UDP/TCP protocol with TLS. &lt;/ul&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


