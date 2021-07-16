/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// OrgVdcNetworkStatus : Represents status of a Org Vdc network. This value will be PENDING if the network has been recorded by VCD but has not been fully configured, CONFIGURING if the network is in transition, REALIZED if the existing state of the network has been fully realized, or REALIZED_FAILED if there was an error creating the network.
type OrgVdcNetworkStatus interface{}

// List of OrgVdcNetworkStatus
const (
	PENDING_OrgVdcNetworkStatus            = "PENDING"
	CONFIGURING_OrgVdcNetworkStatus        = "CONFIGURING"
	REALIZED_OrgVdcNetworkStatus           = "REALIZED"
	REALIZATION_FAILED_OrgVdcNetworkStatus = "REALIZATION_FAILED"
)
