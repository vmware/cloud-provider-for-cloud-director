# VCenterServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcId** | **string** | URN of the associated vCenter server. | [optional] [default to null]
**Name** | **string** | The name of the vCenter server. | [default to null]
**Description** | **string** | Optional description. | [optional] [default to null]
**Username** | **string** | User name to connect to the server. | [default to null]
**Password** | **string** | Cleartext password to connect to the server. | [optional] [default to null]
**Url** | **string** | URL of the server. | [default to null]
**IsEnabled** | **bool** | True if the vCenter server is enabled for use with vCloud Director. | [optional] [default to null]
**VsphereWebClientServerUrl** | **string** | The URL of vCenter web client server. | [optional] [default to null]
**HasProxy** | **bool** | When present, indicates that a proxy exists within vCloud Director that proxies this vCenter server for access by authorized end-users. Setting this field to true when registering a vCenter server will result in a proxy being created for the vCenter server, and another for the corresponding SSO endpoint (if different from the vCenter server&#39;s endpoint). This field is immutable after the vCenter Server is registered, and will be updated by the system when/if the proxy is removed. | [optional] [default to null]
**RootFolder** | **string** | vCenter root folder in which the vCloud Director system folder will be created. This parameter only takes the folder name and not directory structure. | [optional] [default to null]
**VcNoneNetwork** | **string** | Network in Vcenter to be used as &#39;NONE&#39; network by vCD. | [optional] [default to null]
**TenantVisibleName** | **string** | Public label of this vCenter server visible to all tenants. | [optional] [default to null]
**IsConnected** | **bool** | True if the vCenter server is connected. | [optional] [default to null]
**Mode** | **string** | The vcenter mode. One of &lt;ul&gt; &lt;li&gt;NONE - undetermined&lt;/li&gt; &lt;li&gt;IAAS - provider scoped&lt;/li&gt; &lt;li&gt;SDDC - tenant scoped&lt;/li&gt; &lt;li&gt;MIXED&lt;/li&gt; &lt;/ul&gt; IAAS indicates this vCenter server is scoped to the provider. SDDC indicates that this vCenter server is scoped to tenants, while MIXED indicates mixed mode, where both uses are allowed in this vCenter server. | [optional] [default to null]
**ListenerState** | **string** | The vcenter listener state. One of &lt;ul&gt; &lt;li&gt;INITIAL&lt;/li&gt; &lt;li&gt;INVALID_SETTINGS&lt;/li&gt; &lt;li&gt;UNSUPPORTED&lt;/li&gt; &lt;li&gt;DISCONNECTED&lt;/li&gt; &lt;li&gt;CONNECTING&lt;/li&gt; &lt;li&gt;CONNECTED_SYNCING&lt;/li&gt; &lt;li&gt;CONNECTED&lt;/li&gt; &lt;li&gt;STOP_REQ&lt;/li&gt; &lt;li&gt;STOP_AND_PURGE_REQ&lt;/li&gt; &lt;li&gt;STOP_ACK&lt;/li&gt; &lt;/ul&gt; | [optional] [default to null]
**ClusterHealthStatus** | **string** | The overall health status of clusters in this vCenter server. One of &lt;ul&gt; &lt;li&gt;GRAY&lt;/li&gt; &lt;li&gt;RED&lt;/li&gt; &lt;li&gt;YELLOW&lt;/li&gt; &lt;li&gt;GREEN&lt;/li&gt; &lt;/ul&gt; | [optional] [default to null]
**VcVersion** | **string** | The version of the VIM server. | [optional] [default to null]
**BuildNumber** | **string** | The build number of the VIM server. | [optional] [default to null]
**Uuid** | **string** | The instance UUID property of the vCenter server. | [optional] [default to null]
**NsxVManager** | [***NsxVManager**](NsxVManager.md) | the NSX-V attached to this Virtual Center server, when present. | [optional] [default to null]
**ProxyConfigurationUrn** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


