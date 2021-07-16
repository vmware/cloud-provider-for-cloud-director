/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Specifies the Load Balancer pool configuration.
type EdgeLoadBalancerPool struct {
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
	Name string `json:"name"`
	// The destination server port used by the traffic sent to the member.
	DefaultPort int32 `json:"defaultPort,omitempty"`
	// Maximum time (in minutes) to gracefully disable a member. Virtual service waits for the specified time before terminating the existing connections to the members that are disabled. <code>Special values: 0 represents 'Immediate', -1 represents 'Infinite'.</code> 
	GracefulTimeoutPeriod int32 `json:"gracefulTimeoutPeriod,omitempty"`
	// The algorithm for choosing a member within the pool's list of available members for each new connection. Default value is \"LEAST_CONNECTIONS\". Supported algorithms are: <ul> <li>LEAST_CONNECTIONS <li>ROUND_ROBIN <li>CONSISTENT_HASH <li>FASTEST_RESPONSE <li>LEAST_LOAD <li>FEWEST_SERVERS <li>RANDOM <li>FEWEST_TASKS <li>CORE_AFFINITY </ul> <em>CONSISTENT_HASH</em> uses Source IP Address hash. Using <em>FASTEST_RESPONSE</em>, <em>LEAST_LOAD</em>, <em>FEWEST_SERVERS</em>, <em>RANDOM</em>, <em>FEWEST_TASKS</em>, <em>CORE_AFFINITY</em> algorithms may require additional licensing. 
	Algorithm string `json:"algorithm,omitempty"`
	// Member server's health can be monitored by using one or more health monitors. Active monitors generate synthetic traffic and mark a server up or down based on the response. 
	HealthMonitors []EdgeLoadBalancerHealthMonitor `json:"healthMonitors,omitempty"`
	// Selected persistence profile for the Load Balancer Pool.
	PersistenceProfile *EdgeLoadBalancerPersistenceProfile `json:"persistenceProfile,omitempty"`
	// The list of destination servers which are used by the Load Balancer Pool to direct load balanced traffic. 
	Members []EdgeLoadBalancerPoolMember `json:"members,omitempty"`
	// The list of Load Balancer Virtual Services associated with this Load balancer Pool.
	VirtualServiceRefs []EntityReference `json:"virtualServiceRefs,omitempty"`
	// The Edge Gateway associated with this Load Balancer Pool.
	GatewayRef *EntityReference `json:"gatewayRef"`
	// The root certificates to use when validating certificates presented by the pool members.
	CaCertificateRefs []EntityReference `json:"caCertificateRefs,omitempty"`
	// Whether to check the common name of the certificate presented by the pool member. This cannot be enabled if no caCertificateRefs are specified. 
	CommonNameCheckEnabled bool `json:"commonNameCheckEnabled,omitempty"`
	// A list of domain names which will be used to verify the common names or subject alternative names presented by the pool member certificates. It is performed only when common name check (commonNameCheckEnabled) is enabled. If common name check is enabled, but domain names are not specified then the incoming host header will be used to check the certificate. 
	DomainNames []string `json:"domainNames,omitempty"`
}
