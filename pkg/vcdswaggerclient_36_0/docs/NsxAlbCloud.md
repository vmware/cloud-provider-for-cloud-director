# NsxAlbCloud

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the cloud. | [optional] [default to null]
**DisplayName** | **string** | Name of the Cloud used in NSX-ALB Controller&#39;s logs or GUI. | [optional] [default to null]
**AlreadyImported** | **bool** | True if the Cloud is already imported. Cloud cannot be imported again. | [optional] [default to null]
**NetworkPoolRef** | [***EntityReference**](EntityReference.md) | The Network Pool associated with this Cloud. If unset, this Cloud cannot be imported. | [optional] [default to null]
**TransportZoneName** | **string** | Name of the transport zone in NSX-T associated with the NSX-ALB Cloud. If unset, the tranport zone associated with the Load Balancer Cloud is not found in NSX-T.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


