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
	ClusterID        string `yaml:"clusterID"`
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
	insecure := true
	getVdcClient := false
	if inputMap != nil {
		for key, val := range inputMap {
			switch key {
			case "host":
				vcdConfig.Host = getStrValStrict(val, vcdConfig.Host)
			case "org":
				vcdConfig.TenantOrg = getStrValStrict(val, vcdConfig.TenantOrg)
			case "user":
				vcdConfig.User = getStrValStrict(val, vcdConfig.User)
			case "secret":
				vcdConfig.Password = getStrValStrict(val, vcdConfig.Password)
			case "insecure":
				insecure = getBoolValStrict(val, true)
			case "getVdcClient":
				getVdcClient = getBoolValStrict(val, false)
			case "refreshToken":
				vcdConfig.RefreshToken = getStrValStrict(val, vcdConfig.RefreshToken)
			case "userOrg":
				vcdConfig.UserOrg = getStrValStrict(val, vcdConfig.UserOrg)
			}
		}
	}

	return NewVCDClientFromSecrets(
		vcdConfig.Host,
		vcdConfig.TenantOrg,
		vcdConfig.TenantVdc,
		vcdConfig.UserOrg,
		vcdConfig.User,
		vcdConfig.Password,
		vcdConfig.RefreshToken,
		insecure,
		getVdcClient,
	)
}
