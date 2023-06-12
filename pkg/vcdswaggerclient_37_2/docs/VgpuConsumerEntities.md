# VgpuConsumerEntities

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultTotal** | **int32** | How many results there are in total (i.e., considering all pages). | [optional] [default to null]
**PageCount** | **int32** | How many pages there are in total. | [optional] [default to null]
**Page** | **int32** | The page that was fetched, 1-indexed. | [optional] [default to null]
**PageSize** | **int32** | Result count for page that was fetched. | [optional] [default to null]
**Associations** | [**[]Association**](Association.md) | Association info for each result. | [optional] [default to null]
**TotalConsumptionCount** | **int32** | Total number of vGPUs consumed across all the pages for the applied filters. | [optional] [default to null]
**Values** | [**[]VgpuConsumerEntity**](VgpuConsumerEntity.md) | The current page of VgpuConsumerEntity objects. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


