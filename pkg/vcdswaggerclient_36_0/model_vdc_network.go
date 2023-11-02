/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

/*
 * VMware Cloud Director OpenAPI
 *
 * VMware Cloud Director OpenAPI is a new API that is defined using the OpenAPI standards.<br/> This ReSTful API borrows some elements of the legacy VMware Cloud Director API and establishes new patterns for use as described below. <h4>Authentication</h4> Authentication and Authorization schemes are the same as those for the legacy APIs. You can authenticate using the JWT token via the <code>Authorization</code> header or specifying a session using <code>x-vcloud-authorization</code> (The latter form is deprecated). <h4>Operation Patterns</h4> This API follows the following general guidelines to establish a consistent CRUD pattern: <table> <tr>   <th>Operation</th><th>Description</th><th>Response Code</th><th>Response Content</th> </tr><tr>   <td>GET /items<td>Returns a paginated list of items<td>200<td>Response will include Navigational links to the items in the list. </tr><tr>   <td>POST /items<td>Returns newly created item<td>201<td>Content-Location header links to the newly created item </tr><tr>   <td>GET /items/urn<td>Returns an individual item<td>200<td>A single item using same data type as that included in list above </tr><tr>   <td>PUT /items/urn<td>Updates an individual item<td>200<td>Updated view of the item is returned </tr><tr>   <td>DELETE /items/urn<td>Deletes the item<td>204<td>No content is returned. </tr> </table> <h5>Asynchronous operations</h5> Asynchronous operations are determined by the server. In those cases, instead of responding as described above, the server responds with an HTTP Response code 202 and an empty body. The tracking task (which is the same task as all legacy API operations use) is linked via the URI provided in the <code>Location</code> header.<br/> All API calls can choose to service a request asynchronously or synchronously as determined by the server upon interpreting the request. Operations that choose to exhibit this dual behavior will have both options documented by specifying both response code(s) below. The caller must be prepared to handle responses to such API calls by inspecting the HTTP Response code. <h5>Error Conditions</h5> <b>All</b> operations report errors using the following error reporting rules: <ul>   <li>400: Bad Request - In event of bad request due to incorrect data or other user error</li>   <li>401: Bad Request - If user is unauthenticated or their session has expired</li>   <li>403: Forbidden - If the user is not authorized or the entity does not exist</li> </ul> <h4>OpenAPI Design Concepts and Principles</h4> <ul>   <li>IDs are full Uniform Resource Names (URNs).</li>   <li>OpenAPI's <code>Content-Type</code> is always <code>application/json</code></li>   <li>REST links are in the Link header.</li>   <ul>     <li>Multiple relationships for any link are represented by multiple values in a space-separated list.</li>     <li>Links have a custom VMware Cloud Director-specific &quot;model&quot; attribute that hints at the applicable data         type for the links.</li>     <li>title + rel + model attributes evaluates to a unique link.</li>     <li>Links follow Hypermedia as the Engine of Application State (HATEOAS) principles. Links are present if         certain operations are present and permitted for the user&quot;s current role and the state of the         referred entities.</li>   </ul>   <li>APIs follow a flat structure relying on cross-referencing other entities instead of the navigational style       used by the legacy VMware Cloud Director APIs.</li>   <li>Most endpoints that return a list support filtering and sorting similar to the query service in the legacy       VMware Cloud Director APIs.</li>   <li>Accept header must be included to specify the API version for the request similar to calls to existing legacy       VMware Cloud Director APIs.</li>   <li>Each feature has a version in the path element present in its URL.<br/>       <b>Note</b> API URL's without a version in their paths must be considered experimental.</li> </ul>
 *
 * API version: 36.0
 * Contact: https://code.vmware.com/support
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// An organization vDC network.
type VdcNetwork struct {
	// The unique ID for the network. This field is read-only.
	Id string `json:"id"`
	// The name of the network.
	Name string `json:"name"`
	// The description of the network.
	Description string `json:"description"`
	// List of subnets configured for the network.
	Subnets *Subnets `json:"subnets"`
	// The NSX id of the backing network.
	BackingNetworkId string `json:"backingNetworkId"`
	// The object type of the backing network.
	BackingNetworkType *BackingNetworkType `json:"backingNetworkType"`
	// The parent network if the network is a direct network, otherwise it will be null.
	ParentNetworkId *EntityReference `json:"parentNetworkId"`
	// The type of network. Changing the network type allows converting between an isolated and routed network. Note that the \"connection\" field must also be set if converting from isolated to routed network.
	NetworkType *VdcNetworkFenceType `json:"networkType"`
	// The organization vDC the network belongs to. This should be unset if the network is owned by a vDC Group. For API version 35.0 and above, this field will be treated as read only. Please use ownerRef for new network creation.
	OrgVdc *EntityReference `json:"orgVdc"`
	// The org vDC or vDC Group that this network belongs to. If the ownerRef is set to a vDC Group, this network will be available across all the vDCs in the vDC Group. If the vDC Group is backed by a NSX-V network provider, the org vDC network is automatically connected to the distributed router associated with the vDC Group and the \"connection\" field does not need to be set. For API version 35.0 and above, this field should be set for network creation.
	OwnerRef *EntityReference `json:"ownerRef"`
	// For an Org vDC network, whether the vDC is backed by NSX-T.
	OrgVdcIsNsxTBacked bool `json:"orgVdcIsNsxTBacked"`
	// The organization to which the network belongs.
	OrgRef *EntityReference `json:"orgRef"`
	// The edge gateway that the network is attached to.
	Connection *RouterConnection `json:"connection"`
	// Deprecated unused field, this property will be removed in future release.
	IsDefaultNetwork bool `json:"isDefaultNetwork"`
	// Whether this network is shared with other organization vDCs.
	Shared bool `json:"shared"`
	// Whether or not this network will support two subnets
	EnableDualSubnetNetwork bool `json:"enableDualSubnetNetwork"`
	// Description of the network's status.
	Status *OrgVdcNetworkStatus `json:"status"`
	// Brief failure message if the last configuration task failed. Deprecated in Api 33.0, this property will be removed in next release.
	LastTaskFailureMessage string `json:"lastTaskFailureMessage"`
	// Whether guest VLAN tagging is allowed.
	GuestVlanTaggingAllowed bool `json:"guestVlanTaggingAllowed"`
	// Whether network resources such as IP/MAC Addresses are to be retained.
	RetainNicResources bool `json:"retainNicResources"`
	// The id of the cross vdc network if this is a stretched network, otherwise it will be null.
	CrossVdcNetworkId string `json:"crossVdcNetworkId"`
	// The id of the org from which this network can be managed if this is a stretched network, otherwise it will be null.
	CrossVdcNetworkLocationId string `json:"crossVdcNetworkLocationId"`
	// Overlay connectivity ID for this Network. This field is used on creation during POST only and will not be displayed on an object returned through GET or PUT.
	OverlayId int32 `json:"overlayId"`
	// The number of IP addresses defined by the static ip pools. If the network contains any IpV6 subnets, the total ip count will be null.
	TotalIpCount int32 `json:"totalIpCount"`
	// The number of IP address used from the static ip pools.
	UsedIpCount int32 `json:"usedIpCount"`
	// Whether this network is advertised so that it can be routed out to the external networks. This applies only to network backed by NSX-T. Value will be unset if route advertisement is not applicable.
	RouteAdvertised bool `json:"routeAdvertised"`
	// The list of firewall groups of type SECURITY_GROUP/STATIC_MEMBERS that are assigned to the Org VDC Network. These groups can then be used in firewall rules to protect the Org VDC Network and allow/deny traffic.
	SecurityGroups []EntityReference `json:"securityGroups"`
}
