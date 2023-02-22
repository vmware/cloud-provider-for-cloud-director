# Session

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of session | [optional] [default to null]
**User** | [***EntityReference**](EntityReference.md) | User of this session | [optional] [default to null]
**Org** | [***EntityReference**](EntityReference.md) | Organization user is logged into for this session | [optional] [default to null]
**Location** | **string** | The accessible location this session is valid for | [optional] [default to null]
**Roles** | **[]string** | User&#39;s roles for this session | [optional] [default to null]
**RoleRefs** | [**[]EntityReference**](EntityReference.md) | References to user&#39;s roles | [optional] [default to null]
**SessionIdleTimeoutMinutes** | **int32** | The session idle timeout in minutes | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


