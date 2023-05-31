# VdcGroupSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NsxTVdcGroupFaultDomainTagType** | **string** | &lt;ul&gt; Defines the fault domain tag that will be shown for participating/candidate Organization vDCs during creation or update of a vDC Group whose networkProviderType is NSX_T. The options a provider can choose from are COMPUTE_PROVIDER_SCOPE and NETWORK_PROVIDER_SCOPE. This should be selected based on the provider&#39;s compute or networking infrastructure. The default value is COMPUTE_PROVIDER_SCOPE.  &lt;li&gt; NETWORK_PROVIDER_SCOPE: This represents the tenant-facing name of the backing network provider for the organization vDC (NSX-V/NSX-T manager). This should be used when the fault domains in the provider infrastructure are separated by network provider. &lt;/li&gt; &lt;li&gt; COMPUTE_PROVIDER_SCOPE: This represents the tenant-facing name of the backing compute provider for the organization vDC (Provider VDC). This should be used when the fault domains in the provider infrastructure are separated by compute provider. &lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


