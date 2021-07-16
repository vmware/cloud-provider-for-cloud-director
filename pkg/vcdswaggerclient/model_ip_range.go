/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// An IpRange 
type IpRange struct {
	// Starting IP address in the range.
	StartAddress string `json:"startAddress,omitempty"`
	// Ending IP address in the range.
	EndAddress string `json:"endAddress,omitempty"`
}
