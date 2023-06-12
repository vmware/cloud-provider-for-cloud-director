# RouterHealthReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NsxManagersHealth** | [**[]RouterComponentHealth**](RouterComponentHealth.md) | Status of all the nsx managers spanned by the universal router.  | [optional] [default to null]
**ControlVmHealth** | [**[]RouterComponentHealth**](RouterComponentHealth.md) | Status of all the universal router control VMs within each network provider scope. This also includes status of egress points and routes configured for the universal router.  | [optional] [default to null]
**ControlPlaneHealth** | [***RouterComponentHealth**](RouterComponentHealth.md) | Status of the nsx controller cluster associated with the nsx managers.  | [optional] [default to null]
**VdcReachabilityStatus** | [**[]RouterComponentHealth**](RouterComponentHealth.md) | Reachability status for local and remote participating vDC&#39;s of referenced vDC group. States whether a vDC is reachable from this local site.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


