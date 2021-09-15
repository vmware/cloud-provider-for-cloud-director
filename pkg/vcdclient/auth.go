/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package vcdclient

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"k8s.io/klog"
)

// VCDAuthConfig : contains config related to vcd auth
type VCDAuthConfig struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	RefreshToken	string `json:"refreshToken"`
	Org          string `json:"org"`
	Host         string `json:"host"`
	CloudAPIHref string `json:"cloudapihref"`
	VDC          string `json:"vdc"`
	Insecure     bool   `json:"insecure"`
	Token        string `json:"token"`
}

func (config *VCDAuthConfig) GetBearerToken() (*govcd.VCDClient, *http.Response, error) {
	href := fmt.Sprintf("%s/api", config.Host)
	u, err := url.ParseRequestURI(href)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to parse url [%s]: %s", href, err)
	}

	vcdClient := govcd.NewVCDClient(*u, config.Insecure)
	vcdClient.Client.APIVersion = "35.0"
	klog.Infof("Using VCD OpenAPI version [%s]", vcdClient.Client.APIVersion)

	var resp *http.Response
	if config.RefreshToken != "" {
		// Authenticate using refresh token
		accessTokenResponse, resp, err := config.getAccessTokenFromRefreshToken()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get access token from refresh token [%s]: [%v]", config.RefreshToken, err)
		}
		err = vcdClient.SetToken(config.Org, "Authorization", fmt.Sprintf("Bearer %s", accessTokenResponse.AccessToken))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to set authorization header: [%v]", err)
		}
		return vcdClient, resp, nil
	} else {
		resp, err := vcdClient.GetAuthResponse(config.User, config.Password, config.Org)
		if err != nil {
			return nil, resp, fmt.Errorf("unable to authenticate [%s/%s] for url [%s]: [%+v] : [%v]",
				config.Org, config.User, href, resp, err)
		}
	}
	return vcdClient, resp, nil
}

func (config *VCDAuthConfig) GetSwaggerClientFromSecrets() (*govcd.VCDClient, *swaggerClient.APIClient, error) {

	vcdClient, _, err := config.GetBearerToken()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to get bearer token from serets: [%v]", err)
	}
	var authHeader string
	if config.RefreshToken == "" {
		authHeader = fmt.Sprintf("Bearer %s", vcdClient.Client.VCDToken)
	} else {
		authHeader = vcdClient.Client.VCDToken
	}


	swaggerConfig := swaggerClient.NewConfiguration()
	swaggerConfig.BasePath = fmt.Sprintf("%s/cloudapi", config.Host)
	swaggerConfig.AddDefaultHeader("Authorization", authHeader)
	swaggerConfig.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	return vcdClient, swaggerClient.NewAPIClient(swaggerConfig), nil
}

func (config *VCDAuthConfig) GetPlainClientFromSecrets() (*govcd.VCDClient, error) {

	href := fmt.Sprintf("%s/api", config.Host)
	u, err := url.ParseRequestURI(href)
	if err != nil {
		return nil, fmt.Errorf("unable to parse url: [%s]: [%v]", href, err)
	}

	vcdClient := govcd.NewVCDClient(*u, config.Insecure)
	vcdClient.Client.APIVersion = "35.0"
	klog.Infof("Using VCD XML API version [%s]", vcdClient.Client.APIVersion)
	if err = vcdClient.Authenticate(config.User, config.Password, config.Org); err != nil {
		return nil, fmt.Errorf("cannot authenticate with vcd: [%v]", err)
	}

	return vcdClient, nil
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType  string `json:"token_type"`
	ExpiresIn  int32 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func (config *VCDAuthConfig) getAccessTokenFromRefreshToken() (*tokenResponse, *http.Response, error) {
	accessTokenUrl := ""
	if strings.ToLower(config.Org) == "system" {
		accessTokenUrl = fmt.Sprintf("%s/oauth/provider/token", config.Host)
	} else {
		accessTokenUrl = fmt.Sprintf("%s/oauth/tenant/%s/token", config.Host, config.Org)
	}

	payload := url.Values{}
	payload.Set("grant_type", "refresh_token")
	payload.Set("refresh_token", config.RefreshToken)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	r, err := http.NewRequest("POST", accessTokenUrl, strings.NewReader(payload.Encode())) // URL-encoded payload
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get access token from refresh token: [%v]", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(r)
	if err != nil {
		return nil, res, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}
	var accessTokenResponse tokenResponse
	if err = json.Unmarshal(body, &accessTokenResponse); err != nil {
		return nil, res, err
	}
	return &accessTokenResponse, res, nil
}

func NewVCDAuthConfigFromSecrets(host string, user string, secret string, refreshToken string, org string, insecure bool) *VCDAuthConfig {
	return &VCDAuthConfig{
		Host:     host,
		User:     user,
		Password: secret,
		RefreshToken: refreshToken,
		Org:      org,
		Insecure: insecure,
	}
}
