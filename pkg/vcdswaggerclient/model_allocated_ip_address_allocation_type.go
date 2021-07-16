/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// AllocatedIpAddressAllocationType : The possible allocation types for allocated network address. VM_ALLOCATED means that the IP address is allocated by a VM, NAT_ROUTED means that the IP address is used in a NAT routed environment and VSM_ALLOCATED means that the IP address is allocated to an NSX edge gateway interface.
type AllocatedIpAddressAllocationType interface{}

// List of AllocatedIpAddressAllocationType
const (
	VM_ALLOCATED_AllocatedIpAddressAllocationType  = "VM_ALLOCATED"
	NAT_ROUTED_AllocatedIpAddressAllocationType    = "NAT_ROUTED"
	VSM_ALLOCATED_AllocatedIpAddressAllocationType = "VSM_ALLOCATED"
)
