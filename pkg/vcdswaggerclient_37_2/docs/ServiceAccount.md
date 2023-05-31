# ServiceAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of this service account. Sorting on this field is not supported. | [optional] [default to null]
**Name** | **string** | Name of the service account. | [optional] [default to null]
**SoftwareId** | **string** | Unique identifier string to identify the client software being registered. | [optional] [default to null]
**SoftwareVersion** | **string** | Version identifier string for the client software identified by software_id. | [optional] [default to null]
**Role** | [***EntityReference**](EntityReference.md) | Entity reference of the Role assigned to this service account. | [optional] [default to null]
**Uri** | **string** | URL of a web page providing information about the client. | [optional] [default to null]
**Org** | [***EntityReference**](EntityReference.md) |  | [optional] [default to null]
**Status** | **string** | The status of the refresh token for this service account. CREATED: The client exists with no refresh token and no user codes REQUESTED: The client exists with no refresh token and has user codes GRANTED: The client exists with no refresh token and has a single user code marked as granted ACTIVE: The client exists with refresh token  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


