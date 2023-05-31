# ServicesEdgeCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeClusterRef** | [***EntityReference**](EntityReference.md) | The reference to VCD Edge Cluster, if it exists. | [optional] [default to null]
**BackingId** | **string** | The Id of the Edge Cluster in the NSX-T manager. If the user provides the id of NSX-T edge cluster during update, VCD will automatically create a corresponding VCD Edge cluster object referencing the specified NSX-T edge cluster. For NSX-V Edges, this is set to NULL.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


