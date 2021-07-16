/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// List of Subnets of a network 
type Subnets struct {
	Values []Subnet `json:"values,omitempty"`
}
