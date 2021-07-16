/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// A Virtual Service port configuration. 
type EdgeLoadBalancerServicePort struct {
	// Specifies the TCP/UDP profile for the service ports such as TCP_PROXY or TCP_FAST_PATH. If unset, a profile using the system-defined type of TCP_PROXY is used. This is valid for all Application Profiles. For L4 application profile type, user can set the profile to a different type such as TCP_FAST_PATH. For any other application profile type, the TCP/UDP profile type must be TCP_PROXY. 
	TcpUdpProfile *EdgeLoadBalancerTcpUdpProfile `json:"tcpUdpProfile,omitempty"`
	// The starting port number in the port range of the Virtual Service. If a single port is desired, only this field needs to be specified. An example is that HTTP uses port 80. 
	PortStart int32 `json:"portStart"`
	// The ending port number in the port range of the Virtual Service.  This is not required if a single port is desired.
	PortEnd int32 `json:"portEnd,omitempty"`
	// A flag indicating whether SSL termination is enabled at the port or not. This can only be enabled if HTTPS and L4_TLS Application Profile types are used. At least one service port must have ssl enabled when using HTTPS or L4_TLS Application Profile types. For HTTPS, if SSL is disabled for a specific port, traffic from that port will be redirected to the first SSL port specified. This allows for automatic HTTP to HTTPS redirects. Disabling SSL for HTTPS and L4_TLS is allowed only with additional licensing. 
	SslEnabled bool `json:"sslEnabled,omitempty"`
}
