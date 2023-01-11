/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package vcdsdk

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	gitRoot string = ""
)

type VcdConfig struct {
	Host             string `yaml:"host"`
	TenantOrg        string `yaml:"tenantOrg"`
	TenantVdc        string `yaml:"tenantVdc"`
	OvdcNetwork      string `yaml:"ovdcNetwork"`
	User             string `yaml:"user"`
	UserOrg          string `yaml:"userOrg"`
	Password         string `yaml:"password"`
	RefreshToken     string `yaml:"refreshToken"`
	VIPSubnet        string `yaml:"vipSubnet"`
	CertificateAlias string `yaml:"certificateAlias"`
	ClusterID          string `yaml:"clusterID"`
	FreeLoadBalancerIP string `yaml:"freeLoadBalancerIP"`
}

type authorizationDetails struct {
	Username               string `yaml:"username"`
	Password               string `yaml:"password"`
	RefreshToken           string `yaml:"refreshToken"`
	UserOrg                string `yaml:"userOrg"`
	SystemUser             string `yaml:"systemUser"`
	SystemUserPassword     string `yaml:"systemUserPassword"`
	SystemUserRefreshToken string `yaml:"systemUserRefreshToken"`
}

func init() {
	gitRoot = os.Getenv("GITROOT")
	if gitRoot == "" {
		// It is okay to panic here as this will be caught during dev
		panic("GITROOT should be set")
	}
}

func getStrValStrict(val interface{}, defaultVal string) string {
	if strVal, ok := val.(string); ok {
		return strVal
	}

	return defaultVal
}

func getBoolValStrict(val interface{}, defaultVal bool) bool {
	if boolVal, ok := val.(bool); ok {
		return boolVal
	}

	return defaultVal
}

func getTestVCDConfig() (*VcdConfig, error) {
	testVcdInfoFilePath := filepath.Join(gitRoot, "testdata/vcd_info.yaml")
	vcdInfoContent, err := ioutil.ReadFile(testVcdInfoFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading the vcd_info.yaml file contents: [%v]", err)
	}

	var vcdInfo VcdConfig
	err = yaml.Unmarshal(vcdInfoContent, &vcdInfo)
	if err != nil {
		return nil, fmt.Errorf("error parsing vcd_info.yaml file content: [%v]", err)
	}
	return &vcdInfo, nil
}

// getTestVCDClient returns a client created using the file testdata/vcd_info.yaml
func getTestVCDClient(vcdConfig *VcdConfig, inputMap map[string]interface{}) (*Client, error) {
	vcdConfigCopy := *vcdConfig
	insecure := true
	getVdcClient := false
	if inputMap != nil {
		for key, val := range inputMap {
			switch key {
			case "host":
				vcdConfigCopy.Host = getStrValStrict(val, vcdConfigCopy.Host)
			case "org":
				vcdConfigCopy.TenantOrg = getStrValStrict(val, vcdConfigCopy.TenantOrg)
			case "user":
				vcdConfigCopy.User = getStrValStrict(val, vcdConfigCopy.User)
			case "secret":
				vcdConfigCopy.Password = getStrValStrict(val, vcdConfigCopy.Password)
			case "insecure":
				insecure = getBoolValStrict(val, true)
			case "getVdcClient":
				getVdcClient = getBoolValStrict(val, false)
			case "refreshToken":
				vcdConfigCopy.RefreshToken = getStrValStrict(val, vcdConfigCopy.RefreshToken)
			case "userOrg":
				vcdConfigCopy.UserOrg = getStrValStrict(val, vcdConfigCopy.UserOrg)
			}
		}
	}

	return NewVCDClientFromSecrets(
		strings.TrimRight(vcdConfigCopy.Host, "/"),
		vcdConfigCopy.TenantOrg,
		vcdConfigCopy.TenantVdc,
		vcdConfigCopy.UserOrg,
		vcdConfigCopy.User,
		vcdConfigCopy.Password,
		vcdConfigCopy.RefreshToken,
		insecure,
		getVdcClient,
	)
}
