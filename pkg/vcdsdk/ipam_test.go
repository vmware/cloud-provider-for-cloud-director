/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package vcdsdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUnusedIPAddressInAllowedRange(t *testing.T) {

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
		freeIP, err := getUnusedIPAddressInAllowedRange(testCase.StartIPAddress, testCase.EndIPAddress,
			testCase.UsedIPAddressMap, nil)
		assert.NoError(t, err, "Should not get an error while finding unused IP address in range")
		assert.Equal(t, freeIP, testCase.FreeIP, testCase.ErrorComment)
	}

	return
}

func TestGetUnusedIPAddressInRange(t *testing.T) {

	type TestCase struct {
		IPRangeList      []IPRange
		UsedIPAddressMap map[string]bool
		FreeIP           string
		ErrorComment     string
	}

	testCaseList := []TestCase{
		{
			IPRangeList: []IPRange{
				{
					StartIP: "1.2.3.3",
					EndIP:   "1.2.3.5",
				},
			},
			UsedIPAddressMap: map[string]bool{
				"1.2.3.4": true,
				"1.2.3.5": true,
			},
			FreeIP:       "1.2.3.3",
			ErrorComment: "NormalTest: The first IP in range should be returned",
		},
		{
			IPRangeList: []IPRange{
				{
					StartIP: "1.2.3.4",
					EndIP:   "1.2.3.5",
				},
			},
			UsedIPAddressMap: map[string]bool{
				"1.2.3.4": true,
				"1.2.3.5": true,
			},
			FreeIP:       "",
			ErrorComment: "NoIPTest",
		},
		{
			IPRangeList: []IPRange{
				{
					StartIP: "1.2.3.4",
					EndIP:   "1.2.3.5",
				},
				{
					StartIP: "1.2.3.9",
					EndIP:   "1.2.3.11",
				},
			},
			UsedIPAddressMap: map[string]bool{
				"1.2.3.4": true,
				"1.2.3.5": true,
				"1.2.3.9": true,
			},
			FreeIP:       "1.2.3.10",
			ErrorComment: "NoIPTest",
		},
	}

	for _, testCase := range testCaseList {
		freeIP, err := getUnusedIPAddressInRange(testCase.UsedIPAddressMap, testCase.IPRangeList)
		assert.NoError(t, err, "Should not get an error while finding unused IP address in range")
		assert.Equal(t, freeIP, testCase.FreeIP, testCase.ErrorComment)
	}

	return
}

func TestCheckIfIPIsAvailable(t *testing.T) {
	type TestCase struct {
		IPAddress string
		IPRanges  []IPRange
		RetVal    bool
	}

	testCaseList := []TestCase{
		{
			IPAddress: "1.2.3.4",
			IPRanges: []IPRange{
				{
					StartIP: "1.2.3.1",
					EndIP:   "1.2.3.10",
				},
			},
			RetVal: true,
		},
		{
			IPAddress: "1.2.3.4",
			IPRanges: []IPRange{
				{
					StartIP: "2.1.3.1",
					EndIP:   "2.1.3.3",
				},
			},
			RetVal: false,
		},
		{
			IPAddress: "1.2.3.4",
			IPRanges: []IPRange{
				{
					StartIP: "1.2.3.4",
					EndIP:   "1.2.3.4",
				},
			},
			RetVal: true,
		},
		{
			IPAddress: "1.2.3.4",
			IPRanges: []IPRange{
				{
					StartIP: "1.2.3.1",
					EndIP:   "1.2.3.4",
				},
			},
			RetVal: true,
		},
		{
			IPAddress: "1.2.3.4",
			IPRanges: []IPRange{
				{
					StartIP: "1.2.3.4",
					EndIP:   "1.2.3.10",
				},
			},
			RetVal: true,
		},
		{
			IPAddress: "1.2.3.255",
			IPRanges: []IPRange{
				{
					StartIP: "1.2.3.1",
					EndIP:   "1.2.4.1",
				},
			},
			RetVal: true,
		},
		{
			IPAddress: "1.2.3.0",
			IPRanges: []IPRange{
				{
					StartIP: "1.2.3.1",
					EndIP:   "1.2.3.255",
				},
			},
			RetVal: false,
		},
		{
			IPAddress: "1.2.4.0",
			IPRanges: []IPRange{
				{
					StartIP: "1.2.3.1",
					EndIP:   "1.2.3.255",
				},
			},
			RetVal: false,
		},
		{
			IPAddress: "1.2.3.255",
			IPRanges: []IPRange{
				{
					StartIP: "1.2.3.1",
					EndIP:   "1.2.3.255",
				},
			},
			RetVal: true,
		},
	}

	for _, testCase := range testCaseList {
		retVal := checkIfIPInRanges(testCase.IPAddress, &testCase.IPRanges)
		assert.Equal(t, retVal, testCase.RetVal, "Failed test for [%v]", testCase)
	}

	return
}
