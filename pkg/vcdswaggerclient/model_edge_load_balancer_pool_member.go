/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Defines a member (a destination server) in a Load Balancer Pool.
type EdgeLoadBalancerPoolMember struct {
	// The ip address of the Load Balancer Pool member.
	IpAddress string `json:"ipAddress"`
	// The port number of the Load Balancer Pool member. If unset, the port that the client used to connect will be used.
	Port int32 `json:"port,omitempty"`
	// The ratio of selecting eligible servers in the pool.
	Ratio int32 `json:"ratio,omitempty"`
	// Whether the Load Balancer Pool member is enabled or not.
	Enabled bool `json:"enabled,omitempty"`
	// The current health status of the pool member. Possible values are: <ul> <li> UP - The member is operational. <li> DOWN - The member is down. <li> DISABLED - The member is disabled <li> UNKNOWN - The state is unknown. </ul> 
	HealthStatus string `json:"healthStatus,omitempty"`
	// When the member is DOWN, the value gives the names of the health monitors that marked the member as down. If a monitor cannot be determined, the value will be UNKNOWN. 
	MarkedDownBy []string `json:"markedDownBy,omitempty"`
	// The non-localized detailed message on the health of the pool member. 
	DetailedHealthMessage string `json:"detailedHealthMessage,omitempty"`
}
