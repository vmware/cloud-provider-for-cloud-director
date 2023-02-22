# SddcEndpoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | This is a required property. | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**Endpoint** | **string** | A network endpoint that the SDDC exposes for communication.Deprecated in Api 35.0. | [optional] [default to null]
**ProxyId** | **string** | An optional identifier to a Proxy that can be used to establish a connection to the endpoint. Deprecated in Api 35.0 and replaced by proxy.  | [optional] [default to null]
**TargetUrl** | **string** | The URL target of the SDDC endpoint. This is the URL that the browser tab  will be pointed to when the endpoint is launched via the H5 UI of VCD. This is a required property.  | [optional] [default to null]
**Proxy** | [***EntityReference**](EntityReference.md) | A reference to an optional Proxy that can be used to establish a connection to the endpoint | [optional] [default to null]
**Sddc** | [***EntityReference**](EntityReference.md) | The EntityReference of the parent SDDC entity. This is a required property to create the endpoint and once set cannot be edited. | [optional] [default to null]
**IsDefault** | **bool** | True if this is the default endpoint for the parent SDDC. An endpoint being the default for the SDDC means that this endpoint&#39;s target will be launched when the SDDC tile is clicked in the H5 Tenant UI of VCD.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


