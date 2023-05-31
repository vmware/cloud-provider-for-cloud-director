# EdgeIpSecVpnTunnelStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSubnet** | **string** | Local subnet to which a tunnel belongs. | [optional] [default to null]
**PeerSubnet** | **string** | Peer subnet to which a tunnel belongs. | [optional] [default to null]
**PacketsIn** | **int64** | Number of packets received. | [optional] [default to null]
**PacketsOut** | **int64** | Number of packets sent. | [optional] [default to null]
**BytesIn** | **int64** | Number of bytes received. | [optional] [default to null]
**BytesOut** | **int64** | Number of bytes sent. | [optional] [default to null]
**PacketsSentError** | **int64** | Number of packets sent with errors. | [optional] [default to null]
**PacketsReceivedError** | **int64** | Number of packets reveived with errors. | [optional] [default to null]
**PacketsInDropped** | **int64** | Number of packets dropped while receiving. | [optional] [default to null]
**PacketsOutDropped** | **int64** | Number of packets dropped while sending. | [optional] [default to null]
**EncryptionErrors** | **int64** | Number of encryption errors. | [optional] [default to null]
**DecryptionErrors** | **int64** | Number of decryption errors. | [optional] [default to null]
**OverflowErrors** | **int64** | Number of errors due to overflow. | [optional] [default to null]
**ReplayErrors** | **int64** | Number of replay errors. | [optional] [default to null]
**IntegrityErrors** | **int64** | Number of integrity check errors. | [optional] [default to null]
**SaMismatchInErrors** | **int64** | Number of SA mismatch errors while receiving. | [optional] [default to null]
**SaMismatchOutErrors** | **int64** | Number of SA mismatch errors while sending. | [optional] [default to null]
**NoMatchingPolicyErrors** | **int64** | Number of packets dropped because of no matching policy is available. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


