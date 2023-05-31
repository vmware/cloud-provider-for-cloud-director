# VirtualMachineClass

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Virtual Machine Class.  | [optional] [default to null]
**CpuReservationRequestedMHz** | **int64** | CPU in MHz that a node reserves when this VirtualMachineClass is applied to the node of a Kubernetes cluster.  | [optional] [default to null]
**MemoryReservationRequestedMB** | **int64** | Memory in MB that a node reserves when this VirtualMachineClass is applied to the node of a Kubernetes cluster.  | [optional] [default to null]
**CpuCount** | **int32** | Number of CPUs that a node gets when this VirtualMachineClass is applied to the node of a Kubernetes cluster.  | [optional] [default to null]
**MemoryMB** | **int64** | Memory in MB that a node gets when this VirtualMachineClass is applied to the node of a Kubernetes cluster.  | [optional] [default to null]
**IsFullyReserved** | **bool** | This read-only field conveys whether CPU and memory resources are fully reserved or not when this VirtualMachineClass is applied to the node of the Kubernetes cluster.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


