# EdgePrefixList

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of this prefix list. On updates, the id is required for the list, while for create a new id will be generated. This id is not a VCD URN.  | [optional] [default to null]
**Name** | **string** | Name for the prefix list. | [default to null]
**Description** | **string** | Description for this prefix list. | [optional] [default to null]
**Prefixes** | [**[]EdgePrefixListEntry**](EdgePrefixListEntry.md) | List of network prefixes. | [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


