# Token

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of this token. Sorting on this field is not supported. | [default to null]
**Name** | **string** | Name of the token | [default to null]
**Token** | **string** | User specific token that can be used to access proxies. Sorting on this field is not supported. | [optional] [default to null]
**ExpirationTimeUtc** | [**time.Time**](time.Time.md) | Time stamp representing when the token will expire (in UTC). | [optional] [default to null]
**Owner** | [***EntityReference**](EntityReference.md) | Owner of the token. Either a user or an extension | [optional] [default to null]
**Username** | **string** | Name of the user that this token is assigned to. | [optional] [default to null]
**UserId** | **string** | ID of the user that this token is assigned to. Sorting on this field is not supported.  Can not be updated.  | [optional] [default to null]
**OrgName** | **string** | Name of the organization that the assigned user belongs to. Sorting and filtering on this field is not supported. | [optional] [default to null]
**Org** | [***EntityReference**](EntityReference.md) |  | [optional] [default to null]
**Type_** | **string** | Type of the token. Can be of type PROXY or EXTENSION | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


