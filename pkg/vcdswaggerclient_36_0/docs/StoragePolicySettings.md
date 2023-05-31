# StoragePolicySettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsAllValidEntityTypes** | **bool** | If true, designates that this storage policy can be used with all valid entity types | [optional] [default to null]
**DiskIopsMax** | **int64** | Maximum IOPS for any disk associated with this storage policy. | [optional] [default to 0]
**DiskIopsPerGbMax** | **int64** | Maximum IOPS that can be assigned to any disk associated with this storage policy based on the size of the disk (in GB). This is also the default IOPS value used for any disk associated with this policy. If set to zero, Default Disk IOPS is used as the default IOPS to be assigned to any disk associated with this storage policy.  | [optional] [default to 0]
**DiskIopsDefault** | **int64** | Default IOPS value to use for any disk associated with the storage policy. This default is only used when Disk IOPS Per GB Max is set to zero.  | [optional] [default to 0]
**StoragePolicyIopsLimit** | **int64** | The sum of IOPS across all disks associated with this policy will be limited to this value.  | [optional] [default to 0]
**IsIopsLimitingEnabled** | **bool** | Whether IOPS limiting is enabled. | [optional] [default to null]
**IgnoreIopsPlacement** | **bool** | Whether VCD IOPS placement should be ignored for disks using this policy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


