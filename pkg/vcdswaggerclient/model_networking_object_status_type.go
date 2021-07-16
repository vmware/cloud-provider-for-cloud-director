/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// NetworkingObjectStatusType : Represents status type of a networking entity. <ul>   <li> PENDING - Desired entity configuration has been received by system and is pending realization.   <li> CONFIGURING - The system is in process of realizing the entity.   <li> REALIZED - The entity is successfully realized in the system.   <li> REALIZATION_FAILED - There are some issues and the system is not able to realize the entity.   <li> UNKNOWN - Current state of entity is unknown. </ul>
type NetworkingObjectStatusType interface{}

// List of NetworkingObjectStatusType
const (
	PENDING_NetworkingObjectStatusType            = "PENDING"
	CONFIGURING_NetworkingObjectStatusType        = "CONFIGURING"
	REALIZED_NetworkingObjectStatusType           = "REALIZED"
	REALIZATION_FAILED_NetworkingObjectStatusType = "REALIZATION_FAILED"
	UNKNOWN_NetworkingObjectStatusType            = "UNKNOWN"
)
