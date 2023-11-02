/*
 * VMware Cloud Director OpenAPI
 *
 * VMware Cloud Director OpenAPI is a new API that is defined using the OpenAPI standards.<br/> This ReSTful API borrows some elements of the legacy VMware Cloud Director API and establishes new patterns for use as described below. <h4>Authentication</h4> Authentication and Authorization schemes are the same as those for the legacy APIs. You can authenticate using the JWT token via the <code>Authorization</code> header or specifying a session using <code>x-vcloud-authorization</code> (The latter form is deprecated). <h4>Operation Patterns</h4> This API follows the following general guidelines to establish a consistent CRUD pattern: <table> <tr>   <th>Operation</th><th>Description</th><th>Response Code</th><th>Response Content</th> </tr><tr>   <td>GET /items<td>Returns a paginated list of items<td>200<td>Response will include Navigational links to the items in the list. </tr><tr>   <td>POST /items<td>Returns newly created item<td>201<td>Content-Location header links to the newly created item </tr><tr>   <td>GET /items/urn<td>Returns an individual item<td>200<td>A single item using same data type as that included in list above </tr><tr>   <td>PUT /items/urn<td>Updates an individual item<td>200<td>Updated view of the item is returned </tr><tr>   <td>DELETE /items/urn<td>Deletes the item<td>204<td>No content is returned. </tr> </table> <h5>Asynchronous operations</h5> Asynchronous operations are determined by the server. In those cases, instead of responding as described above, the server responds with an HTTP Response code 202 and an empty body. The tracking task (which is the same task as all legacy API operations use) is linked via the URI provided in the <code>Location</code> header.<br/> All API calls can choose to service a request asynchronously or synchronously as determined by the server upon interpreting the request. Operations that choose to exhibit this dual behavior will have both options documented by specifying both response code(s) below. The caller must be prepared to handle responses to such API calls by inspecting the HTTP Response code. <h5>Error Conditions</h5> <b>All</b> operations report errors using the following error reporting rules: <ul>   <li>400: Bad Request - In event of bad request due to incorrect data or other user error</li>   <li>401: Bad Request - If user is unauthenticated or their session has expired</li>   <li>403: Forbidden - If the user is not authorized or the entity does not exist</li> </ul> <h4>OpenAPI Design Concepts and Principles</h4> <ul>   <li>IDs are full Uniform Resource Names (URNs).</li>   <li>OpenAPI's <code>Content-Type</code> is always <code>application/json</code></li>   <li>REST links are in the Link header.</li>   <ul>     <li>Multiple relationships for any link are represented by multiple values in a space-separated list.</li>     <li>Links have a custom VMware Cloud Director-specific &quot;model&quot; attribute that hints at the applicable data         type for the links.</li>     <li>title + rel + model attributes evaluates to a unique link.</li>     <li>Links follow Hypermedia as the Engine of Application State (HATEOAS) principles. Links are present if         certain operations are present and permitted for the user&quot;s current role and the state of the         referred entities.</li>   </ul>   <li>APIs follow a flat structure relying on cross-referencing other entities instead of the navigational style       used by the legacy VMware Cloud Director APIs.</li>   <li>Most endpoints that return a list support filtering and sorting similar to the query service in the legacy       VMware Cloud Director APIs.</li>   <li>Accept header must be included to specify the API version for the request similar to calls to existing legacy       VMware Cloud Director APIs.</li>   <li>Each feature has a version in the path element present in its URL.<br/>       <b>Note</b> API URL's without a version in their paths must be considered experimental.</li> </ul> 
 *
 * API version: 37.2
 * Contact: https://code.vmware.com/support
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Describes the DHCP Binding configuration of an Org vDC Network. 
type DhcpBinding struct {
	// The unique id of the DHCP binding.
	Id string `json:"id"`
	// Display name for the DHCP binding.
	Name string `json:"name"`
	// Description for the DHCP binding.
	Description string `json:"description"`
	// MAC address of the host.
	MacAddress string `json:"macAddress"`
	// IP Address assigned to host. This address must belong to the subnet of Org vDC network. For IPv4, this is required. For IPv6, when not specified, Stateless Address Autoconfiguration (SLAAC) is used to auto-assign an IPv6 address to the DHCPv6 clients. 
	IpAddress string `json:"ipAddress"`
	// The amount of time in seconds of how long a DHCP IP will be leased out for. The minimum is 60s while the maximum is 4,294,967,295s, which is roughly 49,710 days. Default is 24 hours. 
	LeaseTime int64 `json:"leaseTime"`
	// DNS nameservers to be set to client host.
	DnsServers []string `json:"dnsServers"`
	// The type of DHCP binding. <ul>   <li>IPV4 - an IPv4 DHCP binding.   <li>IPV6 - an IPv6 DHCP binding. </ul> 
	BindingType string `json:"bindingType"`
	// Additional configuration for IPv4 DHCP binding. This is ignored for IPV6 binding. 
	DhcpV4BindingConfig *DhcpV4BindingConfig `json:"dhcpV4BindingConfig"`
	// Additional configuration for IPv6 DHCP binding. This is ignored for IPV4 binding. 
	DhcpV6BindingConfig *DhcpV6BindingConfig `json:"dhcpV6BindingConfig"`
	Version *ObjectVersion                   `json:"version"`
}
