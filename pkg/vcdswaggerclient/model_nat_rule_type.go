/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// NatRuleType : The Nat Rule Type for a given NAT rule.
type NatRuleType interface{}

// List of NatRuleType
const (
	DNAT_NatRuleType    = "DNAT"
	SNAT_NatRuleType    = "SNAT"
	NO_DNAT_NatRuleType = "NO_DNAT"
	NO_SNAT_NatRuleType = "NO_SNAT"
)
