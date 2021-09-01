/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

// VCDConfig :
type VCDConfig struct {
	Host string `yaml:"host"`
	VDC  string `yaml:"vdc"`
	Org  string `yaml:"org"`

	// It is allowed to pass the following variables using the config. However
	// that is unsafe security practice. However there can be user scenarios and
	// testing scenarios where this is sensible.
	User         string `yaml:"user" default:""`
	Secret       string `yaml:"secret" default:""`
	RefreshToken string `yaml:"refreshToken" default:""`

	VDCNetwork string `yaml:"network"`
	VIPSubnet  string `yaml:"vipSubnet"`
}

// Ports :
type Ports struct {
	HTTP  int32 `yaml:"http" default:"80"`
	HTTPS int32 `yaml:"https" default:"443"`
}

// OneArm :
type OneArm struct {
	StartIP string `yaml:"startIP"`
	EndIP   string `yaml:"endIP"`
}

// LBConfig :
type LBConfig struct {
	OneArm           *OneArm `yaml:"oneArm,omitempty"`
	Ports            Ports   `yaml:"ports"`
	CertificateAlias string  `yaml:"certAlias"`
}

// CloudConfig contains the config that will be read from the secret
type CloudConfig struct {
	VCD       VCDConfig `yaml:"vcd"`
	LB        LBConfig  `yaml:"loadbalancer"`
	ClusterID string    `yaml:"clusterid"`
}

// ParseCloudConfig : parses config and env to fill in the CloudConfig struct
func ParseCloudConfig(configReader io.Reader) (*CloudConfig, error) {
	var err error
	config := &CloudConfig{}

	decoder := yaml.NewDecoder(configReader)
	decoder.SetStrict(true)

	if err = decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("unable to decode yaml file: [%v]", err)
	}

	return config, nil
}

func SetAuthorization(config *CloudConfig) error {
	username, err := ioutil.ReadFile("/etc/kubernetes/vcloud/basic-auth/username")
	if err != nil {
		return fmt.Errorf("unable to get username: [%v]", err)
	}
	secret, err := ioutil.ReadFile("/etc/kubernetes/vcloud/basic-auth/password")
	if err != nil {
		return fmt.Errorf("unable to get password: [%v]", err)
	}
	config.VCD.User = string(username)
	config.VCD.Secret = string(secret)
	return nil
}

func ValidateCloudConfig(config *CloudConfig) error {
	// TODO: needs more validation
	if config == nil {
		return fmt.Errorf("nil config passed")
	}

	if config.VCD.Host == "" {
		return fmt.Errorf("need a valid vCloud Host")
	}
	if config.VCD.VDCNetwork == "" {
		return fmt.Errorf("need a valid ovdc network name")
	}

	if config.VCD.RefreshToken == "" {
		if config.VCD.User == "" || config.VCD.Secret == "" {
			return fmt.Errorf("credentials not passed correctly")
		}
	} else {
		// if RefreshToken is passed, let's disallow other strings
		if config.VCD.User != "" || config.VCD.Secret != "" {
			return fmt.Errorf("credentials not passed correctly")
		}
	}

	return nil
}
