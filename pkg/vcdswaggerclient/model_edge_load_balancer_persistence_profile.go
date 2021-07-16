/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Specifies the Persistence profile for a Load Balancer Pool. Persistence profile will ensure that the same user sticks to the same server for a desired duration of time. If the persistence profile is unmanaged by Cloud Director, updates that leave the values unchanged will continue to use the same unmanaged profile. Any changes made to the persistence profile will cause Cloud Director to switch the pool to a profile managed by Cloud Director. 
type EdgeLoadBalancerPersistenceProfile struct {
	// Name of the Persistence profile. 
	Name string `json:"name,omitempty"`
	// Type of persistence strategy to use. Supported values are: <ul> <li>CLIENT_IP - The client?s IP is used as the identifier and mapped to the server. <li>HTTP_COOKIE - Load Balancer inserts a cookie into HTTP responses. Cookie name must be provided as value. <li>CUSTOM_HTTP_HEADER - Custom, static mappings of header values to specific servers are used. Header name must be provided as value. <li>APP_COOKIE - Load Balancer reads existing server cookies or URI embedded data such as JSessionID. Cookie name must be provided as value. <li>TLS - Information is embedded in the client?s SSL/TLS ticket ID. This will use default system profile \"System-Persistence-TLS\". </ul> Using <em>CUSTOM_HTTP_HEADER</em>, <em>APP_COOKIE</em>, <em>TLS</em> persistence types may require additional licensing. 
	Type_ string `json:"type"`
	// Value of attribute based on selected persistence type. This is required for HTTP_COOKIE, CUSTOM_HTTP_HEADER and APP_COOKIE persistence types. HTTP_COOKIE, APP_COOKIE must have cookie name set as the value and CUSTOM_HTTP_HEADER must have header name set as the value. 
	Value string `json:"value,omitempty"`
}
