# RemoteEndpoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | This Remote ID is needed to uniquely identify the peer site. If this tunnel is using PSK authentication, the Remote ID is the public IP Address of the remote device terminating the VPN Tunnel. When NAT is configured on the Remote ID, enter the private IP Address of the Remote Site. If the remote ID is not set, VCD will set the remote id to the remote address. If this tunnel is using certificate authentication, enter the distinguished name of the certificate used to secure the remote endpoint (for example, C&#x3D;US,ST&#x3D;Massachusetts,O&#x3D;VMware,OU&#x3D;VCD,CN&#x3D;Edge1). The remote id must be provided in this case.  | [optional] [default to null]
**RemoteAddress** | **string** | IPV4 Address of the remote endpoint on the remote site. This is the Public IPv4 Address of the remote device terminating the VPN connection. | [default to null]
**RemoteNetworks** | **[]string** | List of remote networks. These must be specified in normal Network CIDR format. Specifying no value is interpreted as 0.0.0.0/0 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


