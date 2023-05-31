# EdgeGatewayLoadBalancerConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A flag indicating whether Load Balancer Service is enabled or not. | [default to null]
**ServiceNetworkDefinition** | **string** | The network definition in Gateway CIDR format which will be used by Load Balancer service on edge. All the load balancer service engines associated with the Service Engine Group will be attached to this network. The subnet prefix length must be 25. If nothing is set, the default is &lt;code&gt;192.168.255.1/25&lt;/code&gt;. Default cidr can be configured. This field cannot be updated.  | [optional] [default to null]
**LoadBalancerCloudRef** | [***EntityReference**](EntityReference.md) | Reference to the Load Balancer Cloud. This cloud is used by Edge Gateway&#39;s Load Balancer Virtual Services and Pools. | [optional] [default to null]
**LicenseType** | **string** | The license type of the backing Load Balancer Cloud. &lt;ul&gt; &lt;li&gt;BASIC - Basic edition of the NSX Advanced Load Balancer. &lt;li&gt;ENTERPRISE - Full featured edition of the NSX Advanced Load Balancer. &lt;/ul&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


