/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// This represents the reference to an Edge Cluster used for the gateway. 
type GatewayEdgeClusterReference struct {
	// The reference to VCD Edge Cluster.
	EdgeClusterRef *EntityReference `json:"edgeClusterRef,omitempty"`
	// The Id of the edge cluster in NSX-T manager. The user should specify the id of NSX-T edge cluster during edge gateway create/update. VCD will automatically create a corresponding VCD Edge cluster object referencing the specified NSX-T edge cluster. For NSX-V Edges, this is set to NULL. 
	BackingId string `json:"backingId,omitempty"`
}
