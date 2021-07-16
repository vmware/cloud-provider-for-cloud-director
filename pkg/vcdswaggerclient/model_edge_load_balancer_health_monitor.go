/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Specifies the health monitoring configuration for a Load Balancer Pool. Health monitors are used to monitor health of the Load Balancer Pool's member servers. 
type EdgeLoadBalancerHealthMonitor struct {
	// Name of the Health monitor. Name of the system defined monitors are prefixed with System for example 'System-HTTP'.
	Name string `json:"name,omitempty"`
	// Type of the health monitors. A value of \"-\" represents an unknown type. Such monitor can still be removed from the Load Balancer Pool on update but is not valid on create. Supported values are: <ul> <li>HTTP - HTTP request/response is used to validate health. <li>HTTPS - Used against HTTPS encrypted web servers to validate health. <li>TCP - TCP connection is used to validate health. <li>UDP - A UDP datagram is used to validate health. <li>PING - An ICMP ping is used to validate health. </ul> 
	Type_ string `json:"type"`
	// Whether the Health Monitor is system defined.
	SystemDefined bool `json:"systemDefined,omitempty"`
}
