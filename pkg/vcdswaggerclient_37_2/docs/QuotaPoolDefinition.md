# QuotaPoolDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotaResourceName** | **string** | The localized name of quota resource type.  | [optional] [default to null]
**ResourceType** | **string** | The quota resource type such as memory, cpu, vm etc. Available resource types: memory, cpu, storage, urn:vcloud:legacy:vm, urn:vcloud:type:vmware.tkgcluster:1.0.0  | [default to null]
**QuotaResourceUnit** | **string** | The unit of quota defined for quota resource type. Available quota units for resource types: memory - MB cpu - MHz storage - MB urn:vcloud:legacy:vm - count  | [optional] [default to null]
**Quota** | **int64** | The quota amount for this resource.  | [default to null]
**Qualifiers** | **[]string** | The qualifiers for quota resource type, such as vm.guestOs &#x3D;&#x3D; Windows. This is optional. Qualifiers just helps in narrowing down quota resource based on values of one or more of its properties. If vm is a quota resource, from the above example, only VMs with Windows guest OS will be considered for quota eligibility. If more than one qualifier is provided, system will use AND operator to process them.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


