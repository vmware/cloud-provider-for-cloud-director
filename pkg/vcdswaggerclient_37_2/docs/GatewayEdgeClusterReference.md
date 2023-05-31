# GatewayEdgeClusterReference

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeClusterRef** | [***EntityReference**](EntityReference.md) | The reference to VCD Edge Cluster. | [optional] [default to null]
**BackingId** | **string** | The Id of the edge cluster in NSX-T manager. The user should specify the id of NSX-T edge cluster during edge gateway create/update. VCD will automatically create a corresponding VCD Edge cluster object referencing the specified NSX-T edge cluster. For NSX-V Edges, this is set to NULL.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


