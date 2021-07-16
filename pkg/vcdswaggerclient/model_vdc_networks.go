/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// List of Org vDC networks 
type VdcNetworks struct {
	// How many results there are in total (i.e., considering all pages).
	ResultTotal int32 `json:"resultTotal,omitempty"`
	// How many pages there are in total.
	PageCount int32 `json:"pageCount,omitempty"`
	// The page that was fetched, 1-indexed.
	Page int32 `json:"page,omitempty"`
	// Result count for page that was fetched.
	PageSize int32 `json:"pageSize,omitempty"`
	// Association info for each result.
	Associations []Association `json:"associations,omitempty"`
	Values []VdcNetwork `json:"values,omitempty"`
}
