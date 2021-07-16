/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Specifies the application profile for the virtual service such as whether it's HTTP, HTTPS, or TCP/UDP. 
type EdgeLoadBalancerApplicationProfile struct {
	// Name of the application profile. Name of the system defined monitors are prefixed with System for example 'System-HTTP'.
	Name string `json:"name,omitempty"`
	// The profile type of application that this Virtual Service is configured with. A value of \"-\" represents an unknown type. Such profile can still be removed from the Virtual Service on update but is not valid on create. Values are below. <ul> <li>HTTP - Virtual Service supports HTTP protocol. <li>HTTPS - Virtual Service supports HTTPS protocol. <li>L4 - Virtual Service supports Layer 4 (Transport) using UDP/TCP protocol. <li>L4_TLS - Virtual Service supports Layer 4 (Transport) using UDP/TCP protocol with TLS. </ul> L4_TLS is allowed only with additional licensing. 
	Type_ string `json:"type"`
	// Whether the Application Profile is system defined.
	SystemDefined bool `json:"systemDefined,omitempty"`
}
