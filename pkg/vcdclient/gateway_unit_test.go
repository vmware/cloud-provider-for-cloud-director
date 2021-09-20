/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package vcdclient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
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
		freeIP := getUnusedIPAddressInRange(testCase.StartIPAddress, testCase.EndIPAddress,
			testCase.UsedIPAddressMap)
		assert.Equal(t, freeIP, testCase.FreeIP, testCase.ErrorComment)
	}

	return
}

func TestGetCursor(t *testing.T) {

	type TestCase struct {
		Resp         *http.Response
		Cursor       string
		ErrorComment string
	}

	testCaseList := []TestCase{
		{
			Resp: &http.Response{
				Header: http.Header{
					"x-vmware-vcloud-ceip-id":    []string{"ca6aef23-f33e-48b4-b8ac-17664023fa78"},
					"x-vmware-vcloud-request-id": []string{"23330a4d-88e0-489e-81c6-1ea623e845a8"},
					"Link": []string{
						`<https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn%3Avcloud%3Agateway%3A35835bef-664f-4afe-a45b-07b25ea0b6e4/nat/rules/24ebb8ba-01de-4d5c-a1f4-be787154174f>;rel="down edit remove";title="24ebb8ba-01de-4d5c-a1f4-be787154174f";type="application/json";model="EdgeNatRule"`,
						`<https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn:vcloud:gateway:35835bef-664f-4afe-a45b-07b25ea0b6e4/nat/rules?cursor=eyJORVRXT1JLSU5HX0NVUlNPUl9PRkZTRVQiOiIwIiwicGFnZVNpemUiOjEsIk5FVFdPUktJTkdfQ1VSU09SIjoiMDAwMTEifQ%3D%3D>;rel="nextPage";type="application/json"`,
						`<https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn%3Avcloud%3Agateway%3A35835bef-664f-4afe-a45b-07b25ea0b6e4>;rel="up";type="application/json";model="EdgeGateway"`,
					},
					"Content-Type": []string{
						fmt.Sprintf("application/json;version=%s", ApiVersion),
					},
					"Cache-Control": []string{
						"no-store, must-revalidate",
					},
					"Content-Location": []string{
						"https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn:vcloud:gateway:35835bef-664f-4afe-a45b-07b25ea0b6e4/nat/rules",
					},
				},
			},
			Cursor:       "eyJORVRXT1JLSU5HX0NVUlNPUl9PRkZTRVQiOiIwIiwicGFnZVNpemUiOjEsIk5FVFdPUktJTkdfQ1VSU09SIjoiMDAwMTEifQ==",
			ErrorComment: "Unable to find cursor",
		},
		{
			Resp: &http.Response{
				Header: http.Header{
					"Link": []string{
						`<https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn%3Avcloud%3Agateway%3A35835bef-664f-4afe-a45b-07b25ea0b6e4/nat/rules>;rel="add";type="application/json";model="EdgeNatRule"`,
						`<https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn%3Avcloud%3Agateway%3A35835bef-664f-4afe-a45b-07b25ea0b6e4>;rel="up";type="application/json";model="EdgeGateway"`,
						`<https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn%3Avcloud%3Agateway%3A35835bef-664f-4afe-a45b-07b25ea0b6e4/nat/rules/525ef9cf-088a-4a30-aae2-454aabeead11>;rel="remove edit down";title="525ef9cf-088a-4a30-aae2-454aabeead11";type="application/json";model="EdgeNatRule"`,
						`<https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn:vcloud:gateway:35835bef-664f-4afe-a45b-07b25ea0b6e4/nat/rules?cursor=eyJORVRXT1JLSU5HX0NVUlNPUl9PRkZTRVQiOiIwIiwicGFnZVNpemUiOjEsIk5FVFdPUktJTkdfQ1VSU09SIjoiMDAwMTIifQ%3D%3D>;rel="nextPage";type="application/json"`,
					},
				},
			},
			Cursor:       "eyJORVRXT1JLSU5HX0NVUlNPUl9PRkZTRVQiOiIwIiwicGFnZVNpemUiOjEsIk5FVFdPUktJTkdfQ1VSU09SIjoiMDAwMTIifQ==",
			ErrorComment: "Unable to find cursor",
		},
		{
			Resp: &http.Response{
				Header: http.Header{
					"Link": []string{
						`<https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn%3Avcloud%3Agateway%3A35835bef-664f-4afe-a45b-07b25ea0b6e4/nat/rules>;rel="add";type="application/json";model="EdgeNatRule"`,
						`<https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn%3Avcloud%3Agateway%3A35835bef-664f-4afe-a45b-07b25ea0b6e4>;rel="up";type="application/json";model="EdgeGateway"`,
						`<https://bos1-vcloud-static-176-199.eng.vmware.com/cloudapi/1.0.0/edgeGateways/urn%3Avcloud%3Agateway%3A35835bef-664f-4afe-a45b-07b25ea0b6e4/nat/rules/d602a8af-a9d8-4aa3-85ba-c8e3bd0d54b7>;rel="down edit remove";title="d602a8af-a9d8-4aa3-85ba-c8e3bd0d54b7";type="application/json";model="EdgeNatRule"`,
					},
				},
			},
			Cursor:       "",
			ErrorComment: "There should be no cursor",
		},
	}

	for _, testCase := range testCaseList {
		cursor, err := getCursor(testCase.Resp)
		assert.NoError(t, err, "Error is not expected")
		assert.Equal(t, cursor, testCase.Cursor, testCase.ErrorComment)
	}

	return
}
