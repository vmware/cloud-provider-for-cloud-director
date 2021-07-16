/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// List of configured NAT Rules.
type EdgeNatRules struct {
	// Represents current status of the networking object. 
	Status *NetworkingObjectStatusType `json:"status,omitempty"`
	// The list of NAT Rules.
	Values []EdgeNatRule `json:"values,omitempty"`
}
