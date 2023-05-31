# ResourcePool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Moref** | **string** | The vCenter Id of the Resource Pool. It must be in the format (resgroup-#). | [optional] [default to null]
**ClusterMoref** | **string** | The managed object reference of the Cluster of this resource pool. It must be in the format (domain-c#).  | [optional] [default to null]
**Name** | **string** | The name of the Resource Pool. | [optional] [default to null]
**VcId** | **string** | The vCenter the Resource Pool belongs to. | [optional] [default to null]
**Eligible** | **bool** | Used in resource pool tree navigation. Indicates whether the resource pool is eligible or not.  | [optional] [default to null]
**KubernetesEnabled** | **bool** | Indicates whether the resource pool&#39;s cluster is enabled for vSphere Kubernetes.  | [optional] [default to null]
**VgpuEnabled** | **bool** | Indicates whether the resource pool&#39;s cluster is enabled for vGPU.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


