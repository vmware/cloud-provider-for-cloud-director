# ParticipatingVdcReference

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdcRef** | [***EntityReference**](EntityReference.md) | The reference to the vDC that is part of this a vDC group. | [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | Read-only field that specifies what organization this vDC is in. | [optional] [default to null]
**SiteRef** | [***EntityReference**](EntityReference.md) | The site ID that this vDC belongs to. Required for universal vDC groups. | [optional] [default to null]
**NetworkProviderScope** | **string** | Read-only field that specifies the network provider scope of the vDC. | [optional] [default to null]
**FaultDomainTag** | **string** | Represents the fault domain of a given organization vDC. For NSX_V backed organization vDCs, this is the network provider scope. For NSX_T backed organization vDCs, this can vary (for example name of the provider vDC or compute provider scope).  | [optional] [default to null]
**RemoteOrg** | **bool** | Read-only field that specifies whether the vDC is local to this VCD site. | [optional] [default to null]
**Status** | [***VdcGroupEntityStatus**](VdcGroupEntityStatus.md) | The status that the vDC can be in. An example is if the vDC has been deleted from the system but is still part of the group. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


