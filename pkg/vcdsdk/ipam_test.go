/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package vcdsdk

import (
	"github.com/stretchr/testify/assert"
	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"testing"
)

func TestGetUnusedIPAddressInRange(t *testing.T) {

	type TestCase struct {
		StartIPAddress   string
		EndIPAddress     string
		UsedIPAddressMap map[string]bool
		FreeIP           string
		ErrorComment     string
	}

	testCaseList := []TestCase{
		{
			StartIPAddress: "1.2.3.4",
			EndIPAddress:   "1.2.3.40",
			UsedIPAddressMap: map[string]bool{
				"1.2.3.4": true,
				"1.2.3.5": true,
			},
			FreeIP:       "1.2.3.6",
			ErrorComment: "NormalTest: The first IP in range should be returned",
		},
		{
			StartIPAddress: "1.2.3.4",
			EndIPAddress:   "1.2.3.40",
			UsedIPAddressMap: map[string]bool{
				"1.2.3.24": true,
				"1.2.3.25": true,
			},
			FreeIP:       "1.2.3.4",
			ErrorComment: "TestEdgeCaseFirstIP: The first IP in range should be returned",
		},
		{
			StartIPAddress: "1.2.3.4",
			EndIPAddress:   "1.2.3.6",
			UsedIPAddressMap: map[string]bool{
				"1.2.3.4": true,
				"1.2.3.5": true,
			},
			FreeIP:       "1.2.3.6",
			ErrorComment: "TestEdgeCaseLastIP: The last IP in range should be returned",
		},
		{
			StartIPAddress:   "1.2.3.4",
			EndIPAddress:     "1.2.3.4",
			UsedIPAddressMap: nil,
			FreeIP:           "1.2.3.4",
			ErrorComment:     "OnlyOneIPTest: Exactly one IP is available and unused",
		},
		{
			StartIPAddress: "1.2.3.4",
			EndIPAddress:   "1.2.3.6",
			UsedIPAddressMap: map[string]bool{
				"1.2.3.4": true,
				"1.2.3.5": true,
				"1.2.3.6": true,
			},
			FreeIP:       "",
			ErrorComment: "TestNoAvailableIP: There are no IPs available",
		},
		{
			StartIPAddress: "1.2.3.4",
			EndIPAddress:   "1.2.3.4",
			UsedIPAddressMap: map[string]bool{
				"1.2.3.4": true,
			},
			FreeIP:       "",
			ErrorComment: "OnlyOneUsedIPTest: Exactly one IP is available used.",
		},
	}

	for _, testCase := range testCaseList {
		freeIP, err := getUnusedIPAddressInRange(testCase.StartIPAddress, testCase.EndIPAddress,
			testCase.UsedIPAddressMap)
		assert.NoError(t, err, "Should not get an error while finding unused IP address in range")
		assert.Equal(t, freeIP, testCase.FreeIP, testCase.ErrorComment)
	}

	return
}

func TestCheckIfIPIsAvailable(t *testing.T) {
	type TestCase struct {
		IPAddress   string
		IPRanges []*swaggerClient.IpRanges
		RetVal bool
	}

	testCaseList := []TestCase{
		{
			IPAddress: "1.2.3.4",
			IPRanges: []*swaggerClient.IpRanges{
				{
					Values: []swaggerClient.IpRange{
						{
							StartAddress: "1.2.3.1",
							EndAddress: "1.2.3.10",
						},
					},
				},
			},
			RetVal: true,
		},
		{
			IPAddress: "1.2.3.4",
			IPRanges: []*swaggerClient.IpRanges{
				{
					Values: []swaggerClient.IpRange{
						{
							StartAddress: "2.1.3.1",
							EndAddress: "2.1.3.3",
						},
					},
				},
			},
			RetVal: false,
		},
		{
			IPAddress: "1.2.3.4",
			IPRanges: []*swaggerClient.IpRanges{
				{
					Values: []swaggerClient.IpRange{
						{
							StartAddress: "1.2.3.4",
							EndAddress: "1.2.3.4",
						},
					},
				},
			},
			RetVal: true,
		},
		{
			IPAddress: "1.2.3.4",
			IPRanges: []*swaggerClient.IpRanges{
				{
					Values: []swaggerClient.IpRange{
						{
							StartAddress: "1.2.3.1",
							EndAddress: "1.2.3.4",
						},
					},
				},
			},
			RetVal: true,
		},
		{
			IPAddress: "1.2.3.4",
			IPRanges: []*swaggerClient.IpRanges{
				{
					Values: []swaggerClient.IpRange{
						{
							StartAddress: "1.2.3.4",
							EndAddress: "1.2.3.10",
						},
					},
				},
			},
			RetVal: true,
		},
		{
			IPAddress: "1.2.3.255",
			IPRanges: []*swaggerClient.IpRanges{
				{
					Values: []swaggerClient.IpRange{
						{
							StartAddress: "1.2.3.1",
							EndAddress: "1.2.4.1",
						},
					},
				},
			},
			RetVal: true,
		},
		{
			IPAddress: "1.2.3.0",
			IPRanges: []*swaggerClient.IpRanges{
				{
					Values: []swaggerClient.IpRange{
						{
							StartAddress: "1.2.3.1",
							EndAddress: "1.2.3.255",
						},
					},
				},
			},
			RetVal: false,
		},
	}

	for _, testCase := range testCaseList {
		retVal := checkIfIPIsAvailable(testCase.IPAddress, testCase.IPRanges)
		assert.Equal(t, retVal, testCase.RetVal, "Failed test for [%v]", testCase)
	}

	return
}
