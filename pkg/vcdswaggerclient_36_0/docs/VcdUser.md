# VcdUser

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Immutable user name of the user. | [default to null]
**FullName** | **string** | Full name (display name) of the user. | [optional] [default to null]
**Description** | **string** | Description of the user. | [optional] [default to null]
**Id** | **string** | Unique id for the user. | [optional] [default to null]
**RoleEntityRefs** | [**[]EntityReference**](EntityReference.md) | The role(s) of the user. If isGroupRole is true, this user inherits their role(s) from group membership(s) | [default to null]
**OrgEntityRef** | [***EntityReference**](EntityReference.md) | Organization that the user belongs to. | [optional] [default to null]
**Password** | **string** | Password for the user. Must be null for external users. | [optional] [default to null]
**DeployedVmQuota** | **int32** | The deployed VM quota for this user. Defaults to 0 which means unlimited. This property is deprecated in 35.0. Use user quotas API. | [optional] [default to null]
**StoredVmQuota** | **int32** | The stored VM quota for this user. Defaults to 0 which means unlimited. This property is deprecated in 35.0. Use user quotas API. | [optional] [default to null]
**Email** | **string** | A user&#39;s email address. Based on org email preferences, notifications can be sent to the user via email. | [optional] [default to null]
**NameInSource** | **string** | Name of the user in its source. | [optional] [default to null]
**Enabled** | **bool** | Enabled state of the user. Defaults to true. | [optional] [default to null]
**IsGroupRole** | **bool** | Determines if this user&#39;s role is inherited from a group. Defaults to false. | [optional] [default to null]
**ProviderType** | **string** | Provider type of the user. It is immutable and must be one of: LOCAL, LDAP, SAML, OAUTH. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


