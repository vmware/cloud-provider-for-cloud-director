/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Specifies a summary of a Load Balancer pool.
type EdgeLoadBalancerPoolSummary struct {
	// Represents current status of the networking object. 
	Status *NetworkingObjectStatusType `json:"status,omitempty"`
	// The unique ID of this Load Balancer Pool. On updates, the ID is required for the pool, while for create a new ID will be generated.
	Id string `json:"id,omitempty"`
	// The description of the Load Balancer Pool.
	Description string `json:"description,omitempty"`
	// True if Load Balancer Pool is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Whether passive monitoring for this pool is enabled or not.
	PassiveMonitoringEnabled bool `json:"passiveMonitoringEnabled,omitempty"`
	// The current health status of the pool. Possible values are: <ul> <li> UP - The pool is operational. <li> RUNNING - The pool is operational, but less than 50% of the pool members are up. <li> DOWN - All members in the pool are down. <li> DISABLED - Either the pool is disabled or all of the members are disabled. <li> UNAVAILABLE - The pool is unavailable. Examples: pool has no members or pool is not assigned to any virtual service. <li> UNKNOWN - The pool state is unknown. </ul> 
	HealthStatus string `json:"healthStatus,omitempty"`
	// The total number of members in the pool.
	MemberCount int32 `json:"memberCount,omitempty"`
	// The number of enabled members in the pool.
	EnabledMemberCount int32 `json:"enabledMemberCount,omitempty"`
	// The number of enabled members in the pool that are operational.
	UpMemberCount int32 `json:"upMemberCount,omitempty"`
	// The localized message on the health of the pool.
	HealthMessage string `json:"healthMessage,omitempty"`
	// Name for the Load Balancer Pool. Name is unique across all pools for an Edge Gateway.
	Name string `json:"name,omitempty"`
	// Whether active monitoring for this pool is enabled or not.
	ActiveMonitoringEnabled bool `json:"activeMonitoringEnabled,omitempty"`
	// The list of Load Balancer Virtual Services associated with this Load balancer Pool. Only first 10 Virtual Services will be returned. 
	VirtualServiceRefs []EntityReference `json:"virtualServiceRefs,omitempty"`
	// Whether SSL is enabled for communicatation between the Load Balancer Virtual Services and the pool members 
	MemberSslEnabled bool `json:"memberSslEnabled,omitempty"`
}
