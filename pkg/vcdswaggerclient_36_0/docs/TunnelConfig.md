# TunnelConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerfectForwardSecrecyEnabled** | **bool** | If true, perfect forward secrecy is enabled. The default value is true. | [optional] [default to null]
**DfPolicy** | [***DfPolicyType**](DfPolicyType.md) | Policy for handling degragmentation bit. The default is COPY. | [optional] [default to null]
**DhGroups** | [**[]DhGroupType**](DhGroupType.md) | The list of Diffie-Helman groups to be used is PFS is enabled. Default is GROUP14. | [default to null]
**DigestAlgorithms** | [**[]TunnelDigestAlgorithmType**](TunnelDigestAlgorithmType.md) | The list of Digest algorithms to be used for message digest. The default digest algorithm is implictly covered by default encrpyption algorithm AES_GCM_128.  | [optional] [default to null]
**EncryptionAlgorithms** | [**[]TunnelEncryptionAlgorithmType**](TunnelEncryptionAlgorithmType.md) | The list of Encryption algorithms to use in IPSec tunnel establishment. Default is AES_GCM_128. NO_ENCRYPTION_AUTH_AES_GMAC_* enables authentication on input data without encryption. If one of these options is used, digest algorithm should be empty.  | [default to null]
**SaLifeTime** | **int32** | The Security Association life time in seconds. Default is 3600 seconds. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


