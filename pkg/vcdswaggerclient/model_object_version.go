/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// This property describes the current version of the entity. To prevent clients from overwriting each other's changes, update operations must include the version which can be obtained by issuing a GET operation. If the version number on an update call is missing, the operation will be rejected. This is only needed on update calls. 
type ObjectVersion struct {
	// version number for the given entity.
	Version int32 `json:"version"`
}
