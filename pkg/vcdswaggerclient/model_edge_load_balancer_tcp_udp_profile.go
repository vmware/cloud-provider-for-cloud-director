/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Specifies the TCP/UDP profile for the service ports such as TCP_PROXY or TCP_FAST_PATH. 
type EdgeLoadBalancerTcpUdpProfile struct {
	// Name of the TCP/UDP profile. Name of the system defined monitors are prefixed with System for example 'System-TCP-Proxy'.
	Name string `json:"name,omitempty"`
	// This property specifies the different ways in which TCP/UDP packets are sent to the destination server. A value of \"-\" represents an unknown type. Such profile can still be removed from the Virtual Service on update but is not valid on create. <ul> <li>TCP_PROXY - TCP connection is terminated at the Virtual Service and a new TCP connection is made to the destination server. <li>TCP_FAST_PATH - TCP packets are directly forwarded to the destination server. <li>UDP_FAST_PATH - UDP packets are directly forwarded to the destination server. </ul> 
	Type_ string `json:"type"`
	// Whether the TCP/UDP profile is system defined.
	SystemDefined bool `json:"systemDefined,omitempty"`
}
