/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package vcdclient

import (
	"context"
	"fmt"
	"sync"

	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"github.com/vmware/go-vcloud-director/v2/govcd"
)

var (
	clientCreatorLock sync.Mutex
	clientSingleton   *Client = nil
)

// OneArm : internal struct representing OneArm config details
type OneArm struct {
	StartIPAddress string
	EndIPAddress   string
}

// Client :
type Client struct {
	vcdAuthConfig      *VCDAuthConfig
	vcdClient          *govcd.VCDClient
	vdc                *govcd.Vdc
	apiClient          *swaggerClient.APIClient
	networkName        string
	ipamSubnet         string
	gatewayRef         *swaggerClient.EntityReference
	networkBackingType swaggerClient.BackingNetworkType
	ClusterID          string
	OneArm             *OneArm
	HTTPPort           int32
	HTTPSPort          int32
	CertificateAlias   string
	rwLock             sync.RWMutex
}

func (client *Client) RefreshBearerToken() error {
	// No need to refresh token if logged in using username and password
	if client.vcdAuthConfig.User != "" && client.vcdAuthConfig.Password != "" && client.vcdAuthConfig.RefreshToken == "" {
		return nil
	}
	client.vcdClient.Client.APIVersion = VCloudApiVersion
	accessTokenResponse, _, err := client.vcdAuthConfig.getAccessTokenFromRefreshToken(client.vcdClient.Client.IsSysAdmin)
	if err != nil {
		return fmt.Errorf("failed to refresh access token: [%v]", err)
	}
	err = client.vcdClient.SetToken(client.vcdAuthConfig.Org, "Authorization", fmt.Sprintf("Bearer %s", accessTokenResponse.AccessToken))
	if err != nil {
		return fmt.Errorf("failed to set authorization header: [%v]", err)
	}
	return nil
}

// NewVCDClientFromSecrets :
func NewVCDClientFromSecrets(host string, orgName string, vdcName string,
	networkName string, ipamSubnet string, user string, password string,
	refreshToken string, insecure bool, clusterID string, oneArm *OneArm,
	httpPort int32, httpsPort int32, certAlias string, getVdcClient bool) (*Client, error) {

	// TODO: validation of parameters

	clientCreatorLock.Lock()
	defer clientCreatorLock.Unlock()

	// Return old client if everything matches. Else create new one and cache it.
	// This is suboptimal but is not a common case.
	if clientSingleton != nil {
		if clientSingleton.vcdAuthConfig.Host == host &&
			clientSingleton.vcdAuthConfig.Org == orgName &&
			clientSingleton.vcdAuthConfig.VDC == vdcName &&
			clientSingleton.vcdAuthConfig.User == user &&
			clientSingleton.vcdAuthConfig.Password == password &&
			clientSingleton.vcdAuthConfig.RefreshToken == refreshToken &&
			clientSingleton.vcdAuthConfig.Insecure == insecure {
			return clientSingleton, nil
		}
	}

	vcdAuthConfig := NewVCDAuthConfigFromSecrets(host, user, password, refreshToken, orgName, insecure)

	vcdClient, apiClient, err := vcdAuthConfig.GetSwaggerClientFromSecrets()
	if err != nil {
		return nil, fmt.Errorf("unable to get swagger client from secrets: [%v]", err)
	}

	client := &Client{
		vcdAuthConfig:    vcdAuthConfig,
		vcdClient:        vcdClient,
		apiClient:        apiClient,
		networkName:      networkName,
		ipamSubnet:       ipamSubnet,
		gatewayRef:       nil,
		ClusterID:        clusterID,
		OneArm:           oneArm,
		HTTPPort:         httpPort,
		HTTPSPort:        httpsPort,
		CertificateAlias: certAlias,
	}

	if getVdcClient {
		org, err := vcdClient.GetOrgByName(orgName)
		if err != nil {
			return nil, fmt.Errorf("unable to get org from name [%s]: [%v]", orgName, err)
		}

		client.vdc, err = org.GetVDCByName(vdcName, true)
		if err != nil {
			return nil, fmt.Errorf("unable to get vdc [%s] from org [%s]: [%v]", vdcName, orgName, err)
		}
	}
	client.vcdClient = vcdClient
	// We will specifically cache the gateway ID that corresponds to the
	// network name since it is used frequently in the loadbalancer context.
	ctx := context.Background()
	if err = client.CacheGatewayDetails(ctx); err != nil {
		return nil, fmt.Errorf("unable to get gateway edge from network name [%s]: [%v]",
			client.networkName, err)
	}
	clientSingleton = client

	return clientSingleton, nil
}
