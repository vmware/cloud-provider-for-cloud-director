# RouterConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouterRef** | [***EntityReference**](EntityReference.md) | The edge gateway that this network is attached to. | [optional] [default to null]
**ConnectionType** | [***VdcNetworkConnectionType**](VdcNetworkConnectionType.md) | How the network is connected to the edge gateway. This field is updatable to allow conversions between different types. If owner is a VDC group that is backed by a NSX-V network provider, this field does not need to be set. The organization VDC network will be automatically connected to the distributed router associated with the VDC group.  | [optional] [default to null]
**Connected** | **bool** | Whether network is marked as connected in NSX. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


