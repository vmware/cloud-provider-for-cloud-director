# SddcProxy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**Id** | **string** |  | [optional] [default to null]
**SddcId** | **string** | URN of the parent SDDC. This is not editable once the proxy has been created. | [default to null]
**Enabled** | **bool** | True if the proxy is enabled. Proxy can only be enabled/disabled by privileged users. A disabled proxy cannot be activated and thus, cannot be used. When a proxy is disabled, all active sessions are terminated.  | [optional] [default to null]
**TenantVisible** | **bool** | Whether this proxy has been published to tenants. | [optional] [default to null]
**TargetHost** | **string** | IP address or FQDN of the host being proxied. Lower case formatting will be applied to the value of the property. This is not editable once the proxy has been created.  | [default to null]
**Active** | **bool** | True if the proxy is currently active for the user session associated with the request made to get the proxy. An inactive proxy cannot be used.  | [optional] [default to null]
**Token** | **string** | The generated read-only token that should be used as the password when using this proxy. To generate a new token, activate the proxy. The token is tied to the user session that activated the proxy. If the proxy is inactive, this value will be null.  | [optional] [default to null]
**DefaultProxy** | **bool** | True if this is the default proxy for the parent SDDC. A proxy being the default for the SDDC means that this proxy&#39;s UI will be launched when the SDDC tile is clicked in the H5 Tenant UI of VCD. If no default proxy is set, clicking the SDDC tile will be a no-op.  | [optional] [default to null]
**ParentProxyId** | **string** | The URN of the parent proxy. If a proxy has a parent, the proxy is activated along with its parent and shares the token with its parent. Each proxy may only have one parent. A parent proxy cannot have a parent of its own.  | [optional] [default to null]
**UiUrl** | **string** | The URL of the proxied component&#39;s UI endpoint. This is the URL that the browser tab  will be pointed to when the proxy is launched via the H5 UI of VCD.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


