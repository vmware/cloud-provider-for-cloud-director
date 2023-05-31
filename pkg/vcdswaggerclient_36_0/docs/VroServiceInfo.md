# VroServiceInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is a read-only field in the client. Values set on this field will be ignored by the server.  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**VroEndpoint** | **string** |  | [optional] [default to null]
**Username** | **string** |  | [optional] [default to null]
**Password** | **string** | The password is hidden using a Password Mask represented by a series of 6 asterisks (\&quot;******\&quot;) in the response. The Password Mask is not an acceptable password during VRO Server registration. During an update, the Password Mask set for the &#39;password&#39; field is interpreted as no change to the field and is ignored.  | [optional] [default to null]
**Version** | **string** |  | [optional] [default to null]
**TrustAnchor** | **string** | SSL Certificate chain for the VRO endpoint (deprecated) | [optional] [default to null]
**VcId** | [***EntityReference**](EntityReference.md) | Reference to the associated vCenter server. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


