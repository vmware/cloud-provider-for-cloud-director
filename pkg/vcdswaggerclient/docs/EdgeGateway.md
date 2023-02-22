# EdgeGateway

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***NetworkingObjectStatusType**](NetworkingObjectStatusType.md) | Represents current status of the networking object.  | [optional] [default to null]
**Id** | **string** | The unique identifier of the edge gateway. | [optional] [default to null]
**Name** | **string** | The name of the edge gateway. | [optional] [default to null]
**Description** | **string** | The description of the edge gateway(optional). | [optional] [default to null]
**EdgeGatewayUplinks** | [**[]EdgeGatewayUplink**](EdgeGatewayUplink.md) | The uplink connections for the edge gateway. | [optional] [default to null]
**DistributedRoutingEnabled** | **bool** | A flag indicating whether distributed routing is enabled or not. The default is false. | [optional] [default to null]
**OrgVdcNetworkCount** | **int32** | The number of Org vDC networks connected to the gateway. | [optional] [default to null]
**GatewayBacking** | [***EdgeGatewayBacking**](EdgeGatewayBacking.md) | The backing details of the edge gateway; only required if importing an NSX-T router. | [optional] [default to null]
**OrgVdc** | [***EntityReference**](EntityReference.md) | The organization vDC which the gateway belongs to. Property is deprecated. Please use ownerRef.  | [optional] [default to null]
**OwnerRef** | [***EntityReference**](EntityReference.md) | The organization vDC or vDC Group that this edge gateway belongs to. If the ownerRef is set to a vDC Group, this gateway will be available across all the participating Organization vDCs in the vDC Group.  | [optional] [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | The organization to which the gateway belongs. | [optional] [default to null]
**ServiceNetworkDefinition** | **string** | The network definition in CDIR form that DNS and DHCP service on an NSX-T edge will run on. The subnet prefix length must be 27. This property applies to creating or importing an NSX-T Edge. This is not supported for VMC. If nothing is set, the default is 192.168.255.225/27.  The DHCP listener IP network is on 192.168.255.225/30. The DNS listener IP network is on 192.168.255.228/32.  This field cannot be updated.  | [optional] [default to null]
**DistributedRouterUplinkNetworkDefinition** | **string** | The uplink network is the network that is used to connect the distributed router to the gateway. This is in CIDR form. This is not set if distributed routing is disabled. This field cannot be updated. This applies to NSX-V edges only.  | [optional] [default to null]
**EdgeClusterConfig** | [***GatewayEdgeClusterConfig**](GatewayEdgeClusterConfig.md) | Edge Cluster Configuration for the Edge Gateway. Can be specified if a gateway needs to be placed on a specific set of Edge Clusters. For NSX-T Edges, user should specify the ID of the NSX-T edge cluster as the value of primaryEdgeCluster&#39;s backingId. The gateway defaults to the Edge Cluster of the connected External Network&#39;s backing Tier-0 router, if nothing is specified. The value of secondaryEdgeCluster will be set to NULL for NSX-T edge gateways. For NSX-V Edges, this is read-only and the legacy API must be used for edge specific placement.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


