# IkeConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IkeVersion** | [***IkeVersionType**](IkeVersionType.md) | IKE Protocol Version to use. The default is IKE_V2. | [default to null]
**DhGroups** | [**[]DhGroupType**](DhGroupType.md) | The list of Diffie-Helman groups to be used is PFS is enabled. Default is GROUP14. | [default to null]
**DigestAlgorithms** | [**[]IkeDigestAlgorithmType**](IkeDigestAlgorithmType.md) | The list of Digest algorithms for IKE. This is used during IKE negotiation. Default is SHA2_256. | [optional] [default to null]
**EncryptionAlgorithms** | [**[]IkeEncryptionAlgorithmType**](IkeEncryptionAlgorithmType.md) | The list of Encryption algorithms for IKE. This is used during IKE negotiation. Default is AES_128. | [default to null]
**SaLifeTime** | **int32** | The Security Association life time in seconds. Default is 86400 seconds (1 day). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


