/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Represents the Edge Cluster Configuration for a given Edge Gateway. 
type GatewayEdgeClusterConfig struct {
	// This represents the Primary Edge Cluster used for the gateway. For NSX-T Edges, this means the Edge Cluster the Tier 1 SR resides on. The user should specify the ID of the NSX-T edge cluster as the value of primaryEdgeCluster's backingId. For NSX-V Edges, this means the primary appliance for the gateway. 
	PrimaryEdgeCluster *GatewayEdgeClusterReference `json:"primaryEdgeCluster"`
	// This represents the Secondary Edge Cluster used for the gateway. It is only applicable for NSX-V Edges when High Availability is enabled. If HA is enabled and no secondary edge cluster is specified, both appliances will be deployed on the primary edge cluster. If there is a specific secondary edge cluster, the standby/secondary appliance will be deployed on the secondary edge cluster. For NSX-T Edges, the value of secondaryEdgeCluster should be set to NULL. 
	SecondaryEdgeCluster *GatewayEdgeClusterReference `json:"secondaryEdgeCluster,omitempty"`
}
