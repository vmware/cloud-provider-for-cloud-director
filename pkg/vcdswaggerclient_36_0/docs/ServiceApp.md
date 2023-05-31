# ServiceApp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID for the service application. This field is read-only. | [optional] [default to null]
**Name** | **string** | The name of the service application | [default to null]
**ClientId** | **string** | The client ID | [default to null]
**ClientSecret** | **string** | The client secret | [default to null]
**OrgId** | **string** | The external organization id of the service application | [optional] [default to null]
**ServiceScope** | [***ServiceAppScope**](ServiceAppScope.md) | The scope of the service application | [default to null]
**AccessTokenTTL** | **int32** | The read-only time-to-live of the access token in seconds. | [optional] [default to null]
**AccessTokenExpiration** | [**time.Time**](time.Time.md) | The read-only expiration date of the access token. | [optional] [default to null]
**Enabled** | **bool** | Whether the application is currently enabled. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


