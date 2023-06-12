# VdcNetworkProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryEdgeCluster** | [***EntityReference**](EntityReference.md) | The Edge Cluster where the primary appliance for an NSX-V Edge Gateway will be deployed. | [optional] [default to null]
**SecondaryEdgeCluster** | [***EntityReference**](EntityReference.md) | The Edge Cluster where the secondary appliance for an NSX-V Edge Gateway will be deployed if HA is enabled on the Edge. | [optional] [default to null]
**ServicesEdgeCluster** | [***ServicesEdgeCluster**](ServicesEdgeCluster.md) | The Edge Cluster where the DHCP server profile will be stored for NSX-T networks using NETWORK mode DHCP. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


