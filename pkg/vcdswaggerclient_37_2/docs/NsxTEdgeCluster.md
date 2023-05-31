# NsxTEdgeCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the edge cluster on the NSX-T manager. | [optional] [default to null]
**Name** | **string** | Name of edge cluster. | [optional] [default to null]
**Description** | **string** | Description of the edge cluster. | [optional] [default to null]
**NodeCount** | **int32** | Number of edge transport nodes in the edge cluster. If this information is not available, nodeCount will be set to -1. This field has been deprecated.  | [optional] [default to null]
**NodeType** | [***NsxTEdgeClusterNodeType**](NsxTEdgeClusterNodeType.md) | Type of transport nodes in the edge cluster. All the nodes in the edge cluster are of same type.  | [optional] [default to null]
**DeploymentType** | [***NsxTEdgeClusterDeploymentType**](NsxTEdgeClusterDeploymentType.md) | Deployment type for transport nodes in the edge cluster. The nodes in the edge cluster may have different deployment types. This field has been deprecated.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


