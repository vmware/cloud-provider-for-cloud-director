# AccessControlGrant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VcloudId URN identifier for ACL grant  | [optional] [default to null]
**Tenant** | [***EntityReference**](EntityReference.md) | The tenant this grant applies in. This is managed by the system and depends on the organization the requester operates in at the time of creation and the actual principal type.  | [optional] [default to null]
**GrantType** | **string** | Indicates whether this access control grant is based on user memberships or entitlements  | [default to null]
**ObjectId** | **string** | The object that this access control grant applies to  | [optional] [default to null]
**AccessLevelId** | **string** | The ID of the level of access which the subject will be granted.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


