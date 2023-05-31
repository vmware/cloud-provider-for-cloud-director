# NetworkingCandidateVdc

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the candidate vDC. | [optional] [default to null]
**Name** | **string** | The name of the candidate vDC. | [optional] [default to null]
**OrgRef** | [***EntityReference**](EntityReference.md) | The reference to the organization this vDC is in. | [optional] [default to null]
**SiteRef** | [***EntityReference**](EntityReference.md) | The site ID that this vDC belongs to. | [optional] [default to null]
**NetworkProviderScope** | **string** | The network provider scope of the vDC. | [optional] [default to null]
**FaultDomainTag** | **string** | Represents the fault domain of a given organization vDC. For NSX_V backed organization vDCs, this is the network provider scope. For NSX_T backed organization vDCs, this can vary (for example name of the provider vDC or compute provider scope).  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


