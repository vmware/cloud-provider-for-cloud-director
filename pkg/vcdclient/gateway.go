/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package vcdclient

import (
	"context"
	"fmt"
	"k8s.io/klog"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/antihax/optional"
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/peterhellberg/link"

	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"github.com/vmware/go-vcloud-director/v2/govcd"
)

// CacheGatewayDetails : get gateway reference and cache some details in client object
func (client *Client) CacheGatewayDetails(ctx context.Context) error {

	if client.networkName == "" {
		return fmt.Errorf("network name should not be empty")
	}

	ovdcNetworksAPI := client.apiClient.OrgVdcNetworksApi
	pageNum := int32(1)
	ovdcNetworkID := ""
	for {
		ovdcNetworks, resp, err := ovdcNetworksAPI.GetAllVdcNetworks(ctx, pageNum, 32, nil)
		if err != nil {
			// TODO: log resp in debug mode only
			return fmt.Errorf("unable to get all ovdc networks: [%+v]: [%v]", resp, err)
		}

		if len(ovdcNetworks.Values) == 0 {
			break
		}

		for _, ovdcNetwork := range ovdcNetworks.Values {
			if ovdcNetwork.Name == client.networkName {
				ovdcNetworkID = ovdcNetwork.Id
				break
			}
		}

		if ovdcNetworkID != "" {
			break
		}
		pageNum++
	}
	if ovdcNetworkID == "" {
		return fmt.Errorf("unable to obtain ID for ovdc network name [%s]",
			client.networkName)
	}

	ovdcNetworkAPI := client.apiClient.OrgVdcNetworkApi
	ovdcNetwork, resp, err := ovdcNetworkAPI.GetOrgVdcNetwork(ctx, ovdcNetworkID)
	if err != nil {
		return fmt.Errorf("unable to get network for id [%s]: [%+v]: [%v]", ovdcNetworkID, resp, err)
	}

	// Cache gateway reference
	client.gatewayRef = &swaggerClient.EntityReference{
		Name: ovdcNetwork.Connection.RouterRef.Name,
		Id:   ovdcNetwork.Connection.RouterRef.Id,
	}

	// Cache backing type
	client.networkBackingType = *ovdcNetwork.BackingNetworkType

	klog.Infof("Obtained Gateway [%s] for Network Name [%s] of type [%v]\n",
		client.gatewayRef.Name, client.networkName, client.networkBackingType)

	return nil
}

func getUnusedIPAddressInRange(startIPAddress string, endIPAddress string,
	usedIPAddresses map[string]bool) string {

	// This is not the best approach and can be optimized further by skipping ranges.
	freeIP := ""
	startIP, _, _ := net.ParseCIDR(fmt.Sprintf("%s/32", startIPAddress))
	endIP, _, _ := net.ParseCIDR(fmt.Sprintf("%s/32", endIPAddress))
	for !startIP.Equal(endIP) && usedIPAddresses[startIP.String()] {
		startIP = cidr.Inc(startIP)
	}
	// either the last IP is free or an intermediate IP is not yet used
	if !startIP.Equal(endIP) || startIP.Equal(endIP) && !usedIPAddresses[startIP.String()] {
		freeIP = startIP.String()
	}

	if freeIP != "" {
		klog.Infof("Obtained unused IP [%s] in range [%s-%s]\n", freeIP, startIPAddress, endIPAddress)
	}

	return freeIP
}

func (client *Client) getUnusedInternalIPAddress(ctx context.Context) (string, error) {

	if client.gatewayRef == nil {
		return "", fmt.Errorf("gateway reference should not be nil")
	}

	usedIPAddress := make(map[string]bool)
	pageNum := int32(1)
	for {
		lbVSSummaries, resp, err := client.apiClient.EdgeGatewayLoadBalancerVirtualServicesApi.GetVirtualServiceSummariesForGateway(
			ctx, pageNum, 25, client.gatewayRef.Id, nil)
		if err != nil {
			return "", fmt.Errorf("unable to get virtual service summaries for gateway [%s]: resp: [%v]: [%v]",
				client.gatewayRef.Name, resp, err)
		}
		if len(lbVSSummaries.Values) == 0 {
			break
		}
		for _, lbVSSummary := range lbVSSummaries.Values {
			usedIPAddress[lbVSSummary.VirtualIpAddress] = true
		}

		pageNum++
	}

	freeIP := getUnusedIPAddressInRange(client.OneArm.StartIPAddress,
		client.OneArm.EndIPAddress, usedIPAddress)
	if freeIP == "" {
		return "", fmt.Errorf("unable to find unused IP address in range [%s-%s]",
			client.OneArm.StartIPAddress, client.OneArm.EndIPAddress)
	}

	return freeIP, nil
}

// There are races here since there is no 'acquisition' of an IP. However, since k8s retries, it will
// be correct.
func (client *Client) getUnusedExternalIPAddress(ctx context.Context, ipamSubnet string) (string, error) {
	if client.gatewayRef == nil {
		return "", fmt.Errorf("gateway reference should not be nil")
	}

	// First, get list of ip ranges for the ipamSubnet subnet mask
	edgeGW, resp, err := client.apiClient.EdgeGatewayApi.GetEdgeGateway(ctx, client.gatewayRef.Id)
	if err != nil {
		return "", fmt.Errorf("unable to retrieve edge gateway details for [%s]: resp [%+v]: [%v]",
			client.gatewayRef.Name, resp, err)
	}

	ipRangesList := make([]*swaggerClient.IpRanges, 0)
	for _, edgeGWUplink := range edgeGW.EdgeGatewayUplinks {
		for _, subnet := range edgeGWUplink.Subnets.Values {
			subnetMask := fmt.Sprintf("%s/%d", subnet.Gateway, subnet.PrefixLength)
			// if there is no specified subnet, look at all ranges
			if ipamSubnet == "" {
				ipRangesList = append(ipRangesList, subnet.IpRanges)
			} else if subnetMask == ipamSubnet {
				ipRangesList = append(ipRangesList, subnet.IpRanges)
				break
			}
		}
	}
	if len(ipRangesList) == 0 {
		return "", fmt.Errorf(
			"unable to get appropriate ipRange corresponding to IPAM subnet mask [%s]",
			ipamSubnet)
	}

	// Next, get the list of used IP addresses for this gateway
	usedIPs := make(map[string]bool)
	pageNum := int32(1)
	for {
		gwUsedIPAddresses, resp, err := client.apiClient.EdgeGatewayApi.GetUsedIpAddresses(ctx, pageNum, 25,
			client.gatewayRef.Id, nil)
		if err != nil {
			return "", fmt.Errorf("unable to get used IP addresses of gateway [%s]: [%+v]: [%v]",
				client.gatewayRef.Name, resp, err)
		}
		if len(gwUsedIPAddresses.Values) == 0 {
			break
		}

		for _, gwUsedIPAddress := range gwUsedIPAddresses.Values {
			usedIPs[gwUsedIPAddress.IpAddress] = true
		}

		pageNum++
	}

	// Now get a free IP that is not used.
	freeIP := ""
	for _, ipRanges := range ipRangesList {
		for _, ipRange := range ipRanges.Values {
			freeIP = getUnusedIPAddressInRange(ipRange.StartAddress,
				ipRange.EndAddress, usedIPs)
			if freeIP != "" {
				break
			}
		}
	}
	if freeIP == "" {
		return "", fmt.Errorf("unable to obtain free IP from gateway [%s]; all are used",
			client.gatewayRef.Name)
	}
	klog.Infof("Using unused IP [%s] on gateway [%v]\n", freeIP, client.gatewayRef.Name)

	return freeIP, nil
}

// TODO: There could be a race here as we don't book a slot. Retry repeatedly to get a LB Segment.
func (client *Client) getLoadBalancerSEG(ctx context.Context) (*swaggerClient.EntityReference, error) {
	if client.gatewayRef == nil {
		return nil, fmt.Errorf("gateway reference should not be nil")
	}

	pageNum := int32(1)
	var chosenSEGAssignment *swaggerClient.LoadBalancerServiceEngineGroupAssignment = nil
	for {
		segAssignments, resp, err := client.apiClient.LoadBalancerServiceEngineGroupAssignmentsApi.GetServiceEngineGroupAssignments(
			ctx, pageNum, 25,
			&swaggerClient.LoadBalancerServiceEngineGroupAssignmentsApiGetServiceEngineGroupAssignmentsOpts{
				Filter: optional.NewString(fmt.Sprintf("gatewayRef.id==%s", client.gatewayRef.Id)),
			},
		)
		if err != nil {
			return nil, fmt.Errorf("unable to get service engine group for gateway [%s]: resp: [%v]: [%v]",
				client.gatewayRef.Name, resp, err)
		}
		if len(segAssignments.Values) == 0 {
			return nil, fmt.Errorf("obtained no service engine group assignment for gateway [%s]: [%v]", client.gatewayRef.Name, err)
		}

		for _, segAssignment := range segAssignments.Values {
			if segAssignment.NumDeployedVirtualServices < segAssignment.MaxVirtualServices {
				chosenSEGAssignment = &segAssignment
				break
			}
		}
		if chosenSEGAssignment != nil {
			break
		}

		pageNum++
	}

	if chosenSEGAssignment == nil {
		return nil, fmt.Errorf("unable to find service engine group with free instances")
	}

	klog.Infof("Using service engine group [%v] on gateway [%v]\n", chosenSEGAssignment.ServiceEngineGroupRef, client.gatewayRef.Name)

	return chosenSEGAssignment.ServiceEngineGroupRef, nil
}

func getCursor(resp *http.Response) (string, error) {
	cursorURI := ""
	for _, linklet := range resp.Header["Link"] {
		for _, l := range link.Parse(linklet) {
			if l.Rel == "nextPage" {
				cursorURI = l.URI
				break
			}
		}
		if cursorURI != "" {
			break
		}
	}
	if cursorURI == "" {
		return "", nil
	}

	u, err := url.Parse(cursorURI)
	if err != nil {
		return "", fmt.Errorf("unable to parse cursor URI [%s]: [%v]", cursorURI, err)
	}

	cursorStr := ""
	keyMap, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", fmt.Errorf("unable to parse raw query [%s]: [%v]", u.RawQuery, err)
	}

	if cursorStrList, ok := keyMap["cursor"]; ok {
		cursorStr = cursorStrList[0]
	}

	return cursorStr, nil
}

type NatRuleRef struct {
	Name         string
	ID           string
	ExternalIP   string
	InternalIP   string
	ExternalPort int
	InternalPort int
}

// getNATRuleRef: returns nil if the rule is not found;
func (client *Client) getNATRuleRef(ctx context.Context, natRuleName string) (*NatRuleRef, error) {

	if client.gatewayRef == nil {
		return nil, fmt.Errorf("gateway reference should not be nil")
	}

	var natRuleRef *NatRuleRef = nil
	cursor := optional.EmptyString()
	for {
		natRules, resp, err := client.apiClient.EdgeGatewayNatRulesApi.GetNatRules(
			ctx, 128, client.gatewayRef.Id,
			&swaggerClient.EdgeGatewayNatRulesApiGetNatRulesOpts{
				Cursor: cursor,
			})
		if err != nil {
			return nil, fmt.Errorf("unable to get nat rules: resp: [%+v]: [%v]", resp, err)
		}
		if len(natRules.Values) == 0 {
			break
		}

		for _, rule := range natRules.Values {
			if rule.Name == natRuleName {
				externalPort := 0
				if rule.DnatExternalPort != "" {
					externalPort, err = strconv.Atoi(rule.DnatExternalPort)
					if err != nil {
						return nil, fmt.Errorf("Unable to convert external port [%s] to int: [%v]",
							rule.DnatExternalPort, err)
					}
				}

				internalPort := 0
				if rule.InternalPort != "" {
					internalPort, err = strconv.Atoi(rule.InternalPort)
					if err != nil {
						return nil, fmt.Errorf("Unable to convert internal port [%s] to int: [%v]",
							rule.InternalPort, err)
					}
				}

				natRuleRef = &NatRuleRef{
					ID:           rule.Id,
					Name:         rule.Name,
					ExternalIP:   rule.ExternalAddresses,
					InternalIP:   rule.InternalAddresses,
					ExternalPort: externalPort,
					InternalPort: internalPort,
				}
				break
			}
		}
		if natRuleRef != nil {
			break
		}

		cursorStr, err := getCursor(resp)
		if err != nil {
			return nil, fmt.Errorf("error while parsing response [%+v]: [%v]", resp, err)
		}
		if cursorStr == "" {
			break
		}
		cursor = optional.NewString(cursorStr)
	}

	if natRuleRef == nil {
		return nil, nil // this is not an error
	}

	return natRuleRef, nil
}

func getDNATRuleName(virtualServiceName string) string {
	return fmt.Sprintf("dnat-%s", virtualServiceName)
}

func (client *Client) createDNATRule(ctx context.Context, dnatRuleName string,
	externalIP string, internalIP string, port int32) error {

	if client.gatewayRef == nil {
		return fmt.Errorf("gateway reference should not be nil")
	}

	dnatRuleRef, err := client.getNATRuleRef(ctx, dnatRuleName)
	if err != nil {
		return fmt.Errorf("unexpected error while looking for nat rule [%s] in gateway [%s]: [%v]",
			dnatRuleName, client.gatewayRef.Name, err)
	}
	if dnatRuleRef != nil {
		klog.Infof("DNAT Rule [%s] already exists", dnatRuleName)
		return nil
	}

	ruleType := swaggerClient.NatRuleType(swaggerClient.DNAT_NatRuleType)
	edgeNatRule := swaggerClient.EdgeNatRule{
		Name:              dnatRuleName,
		Enabled:           true,
		RuleType:          &ruleType,
		ExternalAddresses: externalIP,
		InternalAddresses: internalIP,
		DnatExternalPort:  fmt.Sprintf("%d", port),
	}
	resp, err := client.apiClient.EdgeGatewayNatRulesApi.CreateNatRule(ctx, edgeNatRule, client.gatewayRef.Id)
	if resp != nil && resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unable to create dnat rule [%s]: [%s]=>[%s]; expected http response [%v], obtained [%v]: [%v]",
			dnatRuleName, externalIP, internalIP, http.StatusAccepted, resp.StatusCode, err)
	} else if err != nil {
		return fmt.Errorf("unable to create dnat rule [%s]: [%s:%d]=>[%s:%d]: [%v]", dnatRuleName,
			externalIP, port, internalIP, port, err)
	}

	taskURL := resp.Header.Get("Location")
	task := govcd.NewTask(&client.vcdClient.Client)
	task.Task.HREF = taskURL
	if err = task.WaitTaskCompletion(); err != nil {
		return fmt.Errorf("unable to create dnat rule [%s]: [%s]=>[%s]; creation task [%s] did not complete: [%v]",
			dnatRuleName, externalIP, internalIP, taskURL, err)
	}

	klog.Infof("Created DNAT rule [%s]: [%s:%d] => [%s] on gateway [%s]\n", dnatRuleName,
		externalIP, port, internalIP, client.gatewayRef.Name)

	return nil
}

func (client *Client) deleteDNATRule(ctx context.Context, dnatRuleName string,
	failIfAbsent bool) error {

	if client.gatewayRef == nil {
		return fmt.Errorf("gateway reference should not be nil")
	}

	dnatRuleRef, err := client.getNATRuleRef(ctx, dnatRuleName)
	if err != nil {
		return fmt.Errorf("unexpected error while finding dnat rule [%s]: [%v]", dnatRuleName, err)
	}
	if dnatRuleRef == nil {
		if failIfAbsent {
			return fmt.Errorf("dnat rule [%s] does not exist", dnatRuleName)
		}

		return nil
	}

	resp, err := client.apiClient.EdgeGatewayNatRuleApi.DeleteNatRule(ctx,
		client.gatewayRef.Id, dnatRuleRef.ID)
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unable to delete dnat rule [%s]: expected http response [%v], obtained [%v]",
			dnatRuleName, http.StatusAccepted, resp.StatusCode)
	}

	taskURL := resp.Header.Get("Location")
	task := govcd.NewTask(&client.vcdClient.Client)
	task.Task.HREF = taskURL
	if err = task.WaitTaskCompletion(); err != nil {
		return fmt.Errorf("unable to delete dnat rule [%s]: deletion task [%s] did not complete: [%v]",
			dnatRuleName, taskURL, err)
	}
	klog.Infof("Deleted DNAT rule [%s] on gateway [%s]\n", dnatRuleName, client.gatewayRef.Name)

	return nil
}

func (client *Client) getLoadBalancerPoolSummary(ctx context.Context,
	lbPoolName string) (*swaggerClient.EdgeLoadBalancerPoolSummary, error) {
	if client.gatewayRef == nil {
		return nil, fmt.Errorf("gateway reference should not be nil")
	}

	// This should return exactly one result, so no need to accumulate results
	lbPoolSummaries, resp, err := client.apiClient.EdgeGatewayLoadBalancerPoolsApi.GetPoolSummariesForGateway(
		ctx, 1, 25, client.gatewayRef.Id,
		&swaggerClient.EdgeGatewayLoadBalancerPoolsApiGetPoolSummariesForGatewayOpts{
			Filter: optional.NewString(fmt.Sprintf("name==%s", lbPoolName)),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to get reference for LB pool [%s]: [%+v]: [%v]",
			lbPoolName, resp, err)
	}
	if len(lbPoolSummaries.Values) != 1 {
		return nil, nil // this is not an error
	}

	return &lbPoolSummaries.Values[0], nil
}

func (client *Client) getLoadBalancerPool(ctx context.Context,
	lbPoolName string) (*swaggerClient.EntityReference, error) {

	lbPoolSummary, err := client.getLoadBalancerPoolSummary(ctx, lbPoolName)
	if err != nil {
		return nil, fmt.Errorf("error when getting LB Pool: [%v]", err)
	}
	if lbPoolSummary == nil {
		return nil, nil // this is not an error
	}

	return &swaggerClient.EntityReference{
		Name: lbPoolSummary.Name,
		Id:   lbPoolSummary.Id,
	}, nil
}

func (client *Client) formLoadBalancerPool(lbPoolName string, ips []string,
	internalPort int32) (swaggerClient.EdgeLoadBalancerPool, []swaggerClient.EdgeLoadBalancerPoolMember) {
	lbPoolMembers := make([]swaggerClient.EdgeLoadBalancerPoolMember, len(ips))
	for i, ip := range ips {
		lbPoolMembers[i].IpAddress = ip
		lbPoolMembers[i].Port = internalPort
		lbPoolMembers[i].Ratio = 1
		lbPoolMembers[i].Enabled = true
	}

	lbPool := swaggerClient.EdgeLoadBalancerPool{
		Enabled:               true,
		Name:                  lbPoolName,
		DefaultPort:           internalPort,
		GracefulTimeoutPeriod: 0, // when service outage occurs, immediately mark as bad
		Members:               lbPoolMembers,
		GatewayRef:            client.gatewayRef,
	}
	return lbPool, lbPoolMembers
}

func (client *Client) createLoadBalancerPool(ctx context.Context, lbPoolName string,
	ips []string, internalPort int32) (*swaggerClient.EntityReference, error) {

	if client.gatewayRef == nil {
		return nil, fmt.Errorf("gateway reference should not be nil")
	}

	lbPoolRef, err := client.getLoadBalancerPool(ctx, lbPoolName)
	if err != nil {
		return nil, fmt.Errorf("unexpected error when querying for pool [%s]: [%v]",
			lbPoolName, err)
	}
	if lbPoolRef != nil {
		klog.Infof("LoadBalancer Pool [%s] already exists", lbPoolName)
		return lbPoolRef, nil
	}

	lbPool, lbPoolMembers := client.formLoadBalancerPool(lbPoolName, ips, internalPort)
	resp, err := client.apiClient.EdgeGatewayLoadBalancerPoolsApi.CreateLoadBalancerPool(ctx, lbPool)
	if err != nil {
		return nil, fmt.Errorf("unable to create loadbalancer pool with name [%s], members [%+v]: resp [%+v]: [%v]",
			lbPoolName, lbPoolMembers, resp, err)
	}
	if resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("unable to create loadbalancer pool; expected http response [%v], obtained [%v]",
			http.StatusAccepted, resp.StatusCode)
	}

	taskURL := resp.Header.Get("Location")
	task := govcd.NewTask(&client.vcdClient.Client)
	task.Task.HREF = taskURL
	if err = task.WaitTaskCompletion(); err != nil {
		return nil, fmt.Errorf("unable to create loadbalancer pool; creation task [%s] did not complete: [%v]",
			taskURL, err)
	}

	// Get the pool to return it
	lbPoolRef, err = client.getLoadBalancerPool(ctx, lbPoolName)
	if err != nil {
		return nil, fmt.Errorf("unexpected error when querying for pool [%s]: [%v]",
			lbPoolName, err)
	}
	if lbPoolRef == nil {
		return nil, fmt.Errorf("unable to query for loadbalancer pool [%s] that was freshly created: [%v]",
			lbPoolName, err)
	}
	klog.Infof("Created lb pool [%v] on gateway [%v]\n", lbPoolRef, client.gatewayRef.Name)

	return lbPoolRef, nil
}

func (client *Client) deleteLoadBalancerPool(ctx context.Context, lbPoolName string,
	failIfAbsent bool) error {

	if client.gatewayRef == nil {
		return fmt.Errorf("gateway reference should not be nil")
	}

	lbPoolRef, err := client.getLoadBalancerPool(ctx, lbPoolName)
	if err != nil {
		return fmt.Errorf("unexpected error in retrieving loadbalancer pool [%s]: [%v]",
			lbPoolName, err)
	}
	if lbPoolRef == nil {
		if failIfAbsent {
			return fmt.Errorf("LoadBalancer pool [%s] does not exist", lbPoolName)
		}

		return nil
	}

	resp, err := client.apiClient.EdgeGatewayLoadBalancerPoolApi.DeleteLoadBalancerPool(ctx, lbPoolRef.Id)
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unable to delete lb pool; expected http response [%v], obtained [%v]",
			http.StatusAccepted, resp.StatusCode)
	}

	taskURL := resp.Header.Get("Location")
	task := govcd.NewTask(&client.vcdClient.Client)
	task.Task.HREF = taskURL
	if err = task.WaitTaskCompletion(); err != nil {
		return fmt.Errorf("unable to delete lb pool; deletion task [%s] did not complete: [%v]",
			taskURL, err)
	}
	klog.Infof("Deleted loadbalancer pool [%s]\n", lbPoolName)

	return nil
}

func (client *Client) updateLoadBalancerPool(ctx context.Context, lbPoolName string, ips []string,
	internalPort int32) (*swaggerClient.EntityReference, error) {
	lbPoolRef, err := client.getLoadBalancerPool(ctx, lbPoolName)
	if err != nil {
		return nil, fmt.Errorf("unexpected error when querying for pool [%s]: [%v]", lbPoolName, err)
	}
	if lbPoolRef == nil {
		return nil, fmt.Errorf("no lb pool found with name [%s]: [%v]", lbPoolName, err)
	}

	lbPool, lbPoolMembers := client.formLoadBalancerPool(lbPoolName, ips, internalPort)
	resp, err := client.apiClient.EdgeGatewayLoadBalancerPoolApi.UpdateLoadBalancerPool(ctx, lbPool, lbPoolRef.Id)
	if err != nil {
		return nil, fmt.Errorf("unable to update loadbalancer pool with name [%s], members [%+v]: resp [%+v]: [%v]",
			lbPoolName, lbPoolMembers, resp, err)
	}
	if resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("unable to update loadbalancer pool; expected http response [%v], obtained [%v]",
			http.StatusAccepted, resp.StatusCode)
	}

	taskURL := resp.Header.Get("Location")
	task := govcd.NewTask(&client.vcdClient.Client)
	task.Task.HREF = taskURL
	if err = task.WaitTaskCompletion(); err != nil {
		return nil, fmt.Errorf("unable to update loadbalancer pool; update task [%s] did not complete: [%v]",
			taskURL, err)
	}

	// Get the pool to return it
	lbPoolRef, err = client.getLoadBalancerPool(ctx, lbPoolName)
	if err != nil {
		return nil, fmt.Errorf("unexpected error when querying for pool [%s]: [%v]",
			lbPoolName, err)
	}
	if lbPoolRef == nil {
		return nil, fmt.Errorf("unable to query for loadbalancer pool [%s] that was updated: [%v]",
			lbPoolName, err)
	}
	klog.Infof("Updated lb pool [%v] on gateway [%v]\n", lbPoolRef, client.gatewayRef.Name)

	return lbPoolRef, nil
}

func (client *Client) getVirtualService(ctx context.Context,
	virtualServiceName string) (*swaggerClient.EdgeLoadBalancerVirtualServiceSummary, error) {

	if client.gatewayRef == nil {
		return nil, fmt.Errorf("gateway reference should not be nil")
	}

	// This should return exactly one result, so no need to accumulate results
	lbVSSummaries, resp, err := client.apiClient.EdgeGatewayLoadBalancerVirtualServicesApi.GetVirtualServiceSummariesForGateway(
		ctx, 1, 25, client.gatewayRef.Id,
		&swaggerClient.EdgeGatewayLoadBalancerVirtualServicesApiGetVirtualServiceSummariesForGatewayOpts{
			Filter: optional.NewString(fmt.Sprintf("name==%s", virtualServiceName)),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to get reference for LB VS [%s]: resp: [%v]: [%v]",
			virtualServiceName, resp, err)
	}
	if len(lbVSSummaries.Values) != 1 {
		return nil, nil // this is not an error
	}

	return &lbVSSummaries.Values[0], nil
}

func (client *Client) createVirtualService(ctx context.Context, virtualServiceName string,
	lbPoolRef *swaggerClient.EntityReference, segRef *swaggerClient.EntityReference,
	freeIP string, rdeVIP string, vsType string, externalPort int32,
	certificateAlias string) (*swaggerClient.EntityReference, error) {

	if client.gatewayRef == nil {
		return nil, fmt.Errorf("gateway reference should not be nil")
	}

	vsSummary, err := client.getVirtualService(ctx, virtualServiceName)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while getting summary for LB VS [%s]: [%v]",
			virtualServiceName, err)
	}
	if vsSummary != nil {
		klog.Infof("LoadBalancer Virtual Service [%s] already exists", virtualServiceName)
		return &swaggerClient.EntityReference{
			Name: vsSummary.Name,
			Id:   vsSummary.Id,
		}, nil
	}

	var virtualServiceConfig *swaggerClient.EdgeLoadBalancerVirtualService = nil
	switch vsType {
	case "HTTP":
		virtualServiceConfig = &swaggerClient.EdgeLoadBalancerVirtualService{
			Name:                  virtualServiceName,
			Enabled:               true,
			VirtualIpAddress:      freeIP,
			LoadBalancerPoolRef:   lbPoolRef,
			GatewayRef:            client.gatewayRef,
			ServiceEngineGroupRef: segRef,
			ServicePorts: []swaggerClient.EdgeLoadBalancerServicePort{
				{
					TcpUdpProfile: &swaggerClient.EdgeLoadBalancerTcpUdpProfile{
						Type_: "TCP_PROXY",
					},
					PortStart:  externalPort,
					SslEnabled: false,
				},
			},
			ApplicationProfile: &swaggerClient.EdgeLoadBalancerApplicationProfile{
				Name:          "System-HTTP",
				Type_:         vsType,
				SystemDefined: true,
			},
		}
		break

	case "HTTPS":
		if certificateAlias == "" {
			return nil, fmt.Errorf("certificate alias should no be empty for HTTPS service")
		}

		certLibItems, resp, err := client.apiClient.CertificateLibraryApi.QueryCertificateLibrary(ctx,
			1, 128,
			&swaggerClient.CertificateLibraryApiQueryCertificateLibraryOpts{
				Filter: optional.NewString(fmt.Sprintf("alias==%s", certificateAlias)),
			},
		)
		if err != nil {
			return nil, fmt.Errorf("unable to get cert with alias [%s]: resp: [%v]: [%v]",
				certificateAlias, resp, err)
		}
		if len(certLibItems.Values) != 1 {
			return nil, fmt.Errorf("expected 1 cert with alias [%s], obtained [%d]",
				certificateAlias, len(certLibItems.Values))
		}
		virtualServiceConfig = &swaggerClient.EdgeLoadBalancerVirtualService{
			Name:                  virtualServiceName,
			Enabled:               true,
			VirtualIpAddress:      freeIP,
			LoadBalancerPoolRef:   lbPoolRef,
			GatewayRef:            client.gatewayRef,
			ServiceEngineGroupRef: segRef,
			CertificateRef: &swaggerClient.EntityReference{
				Name: certLibItems.Values[0].Alias,
				Id:   certLibItems.Values[0].Id,
			},
			ServicePorts: []swaggerClient.EdgeLoadBalancerServicePort{
				{
					TcpUdpProfile: &swaggerClient.EdgeLoadBalancerTcpUdpProfile{
						Type_: "TCP_PROXY",
					},
					PortStart:  externalPort,
					SslEnabled: true,
				},
			},
			ApplicationProfile: &swaggerClient.EdgeLoadBalancerApplicationProfile{
				Name:          "System-HTTP",
				Type_:         vsType,
				SystemDefined: true,
			},
		}
		break

	default:
		return nil, fmt.Errorf("unhandled virtual service type [%s]", vsType)
	}

	resp, err := client.apiClient.EdgeGatewayLoadBalancerVirtualServicesApi.CreateVirtualService(ctx, *virtualServiceConfig)
	if resp != nil && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf(
			"unable to create virtual service; expected http response [%v], obtained [%v]: resp: [%#v]: [%v]",
			http.StatusAccepted, resp.StatusCode, resp, err)
	} else if err != nil {
		return nil, fmt.Errorf("error while creating virtual service [%s]: [%v]", virtualServiceName, err)
	}

	taskURL := resp.Header.Get("Location")
	task := govcd.NewTask(&client.vcdClient.Client)
	task.Task.HREF = taskURL
	if err = task.WaitTaskCompletion(); err != nil {
		return nil, fmt.Errorf("unable to create virtual service; creation task [%s] did not complete: [%v]",
			taskURL, err)
	}

	// update RDE with freeIp
	if client.ClusterID != "" {
		err = client.addVirtualIpToRDE(ctx, rdeVIP)
		if err != nil {
			klog.Errorf("error when adding virtual IP to RDE: [%v]", err)
		}
	}

	vsSummary, err = client.getVirtualService(ctx, virtualServiceName)
	if err != nil {
		return nil, fmt.Errorf("unable to get summary for freshly created LB VS [%s]: [%v]",
			virtualServiceName, err)
	}
	if vsSummary == nil {
		return nil, fmt.Errorf("unable to get summary of freshly created virtual service [%s]: [%v]",
			virtualServiceName, err)
	}

	virtualServiceRef := &swaggerClient.EntityReference{
		Name: vsSummary.Name,
		Id:   vsSummary.Id,
	}
	klog.Infof("Created virtual service [%v] on gateway [%v]\n", virtualServiceRef, client.gatewayRef.Name)

	return virtualServiceRef, nil
}

func (client *Client) deleteVirtualService(ctx context.Context, virtualServiceName string,
	failIfAbsent bool, rdeVIP string) error {

	if client.gatewayRef == nil {
		return fmt.Errorf("gateway reference should not be nil")
	}

	vsSummary, err := client.getVirtualService(ctx, virtualServiceName)
	if err != nil {
		return fmt.Errorf("unable to get summary for LB Virtual Service [%s]: [%v]",
			virtualServiceName, err)
	}
	if vsSummary == nil {
		if failIfAbsent {
			return fmt.Errorf("Virtual Service [%s] does not exist", virtualServiceName)
		}

		return nil
	}

	resp, err := client.apiClient.EdgeGatewayLoadBalancerVirtualServiceApi.DeleteVirtualService(
		ctx, vsSummary.Id)
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unable to delete virtual service [%s]; expected http response [%v], obtained [%v]",
			vsSummary.Name, http.StatusAccepted, resp.StatusCode)
	}

	taskURL := resp.Header.Get("Location")
	task := govcd.NewTask(&client.vcdClient.Client)
	task.Task.HREF = taskURL
	if err = task.WaitTaskCompletion(); err != nil {
		return fmt.Errorf("unable to delete virtual service; deletion task [%s] did not complete: [%v]",
			taskURL, err)
	}
	klog.Infof("Deleted virtual service [%s]\n", virtualServiceName)

	// remove virtual ip from RDE
	if client.ClusterID != "" {
		err = client.removeVirtualIpFromRDE(ctx, rdeVIP)
		if err != nil {
			return fmt.Errorf("error when removing vip from RDE: [%v]", err)
		}
	}

	return nil
}

// CreateLoadBalancer : create a new load balancer pool and virtual service pointing to it
func (client *Client) CreateLoadBalancer(ctx context.Context, virtualServiceNamePrefix string,
	lbPoolNamePrefix string, ips []string, httpPort int32, httpsPort int32) (string, error) {

	client.rwLock.Lock()
	defer client.rwLock.Unlock()

	if httpPort == 0 && httpsPort == 0 {
		// nothing to do here
		klog.Infof("There is no port specified. Hence nothing to do.")
		return "", fmt.Errorf("nothing to do since http and https ports are not specified")
	}

	if client.gatewayRef == nil {
		return "", fmt.Errorf("gateway reference should not be nil")
	}

	type PortDetails struct {
		portSuffix       string
		serviceType      string
		externalPort     int32
		internalPort     int32
		certificateAlias string
	}
	portDetails := []PortDetails{
		{
			portSuffix:       "http",
			serviceType:      "HTTP",
			externalPort:     client.HTTPPort,
			internalPort:     httpPort,
			certificateAlias: "",
		},
		{
			portSuffix:       "https",
			serviceType:      "HTTPS",
			externalPort:     client.HTTPSPort,
			internalPort:     httpsPort,
			certificateAlias: client.CertificateAlias,
		},
	}

	externalIP, err := client.getUnusedExternalIPAddress(ctx, client.ipamSubnet)
	if err != nil {
		return "", fmt.Errorf("unable to get unused IP address from subnet [%s]: [%v]",
			client.ipamSubnet, err)
	}
	klog.Infof("Using external IP [%s] for virtual service\n", externalIP)

	for _, portDetail := range portDetails {
		if portDetail.internalPort == 0 {
			klog.Infof("No internal port specified for [%s], hence loadbalancer not created\n", portDetail.portSuffix)
			continue
		}

		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portDetail.portSuffix)
		lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portDetail.portSuffix)

		vsSummary, err := client.getVirtualService(ctx, virtualServiceName)
		if err != nil {
			return "", fmt.Errorf("unexpected error while querying for virtual service [%s]: [%v]",
				virtualServiceName, err)
		}
		if vsSummary != nil {
			if vsSummary.LoadBalancerPoolRef.Name != lbPoolName {
				return "", fmt.Errorf("Virtual Service [%s] found with unexpected loadbalancer pool [%s]",
					virtualServiceName, lbPoolName)
			}

			klog.Infof("LoadBalancer Virtual Service [%s] already exists", virtualServiceName)
			continue
		}

		virtualServiceIP := externalIP
		if client.OneArm != nil {
			internalIP, err := client.getUnusedInternalIPAddress(ctx)
			if err != nil {
				return "", fmt.Errorf("unable to get internal IP address for one-arm mode: [%v]", err)
			}

			dnatRuleName := getDNATRuleName(virtualServiceName)
			if err = client.createDNATRule(ctx, dnatRuleName, externalIP, internalIP, portDetail.externalPort); err != nil {
				return "", fmt.Errorf("unable to create dnat rule [%s] => [%s]: [%v]",
					externalIP, internalIP, err)
			}
			// use the internal IP to create virtual service
			virtualServiceIP = internalIP

			// We get an IP address above and try to get-or-create a DNAT rule from external IP => internal IP. If the rule already
			// existed, the old DNAT rule will remain unchanged. Hence we get the old externalIP from the old rule and use it.
			// What happens to the new externalIP that we selected above? It just remains unused and hence does not get
			// allocated and disappears. Since there is no IPAM based resource _acquisition_, the new externalIP can just be
			// forgotten about.
			dnatRuleRef, err := client.getNATRuleRef(ctx, dnatRuleName)
			if err != nil {
				return "", fmt.Errorf("unable to retrieve created dnat rule [%s]: [%v]", dnatRuleName, err)
			}
			if dnatRuleRef == nil {
				return "", fmt.Errorf("retrieved dnat rule ref is nil")
			}
			externalIP = dnatRuleRef.ExternalIP
		}

		segRef, err := client.getLoadBalancerSEG(ctx)
		if err != nil {
			return "", fmt.Errorf("unable to get service engine group from edge [%s]: [%v]",
				client.gatewayRef.Name, err)
		}

		lbPoolRef, err := client.createLoadBalancerPool(ctx, lbPoolName, ips, portDetail.internalPort)
		if err != nil {
			return "", fmt.Errorf("unable to create load balancer pool [%s]: [%v]", lbPoolName, err)
		}

		virtualServiceRef, err := client.createVirtualService(ctx, virtualServiceName, lbPoolRef, segRef,
			virtualServiceIP, externalIP, portDetail.serviceType, portDetail.externalPort, portDetail.certificateAlias)
		if err != nil {
			return "", fmt.Errorf("unable to create virtual service [%s] with address [%s:%d]: [%v]",
				virtualServiceName, virtualServiceIP, portDetail.externalPort, err)
		}
		klog.Infof("Created Load Balancer with virtual service [%v], pool [%v] on gateway [%s]\n",
			virtualServiceRef, lbPoolRef, client.gatewayRef.Name)
	}

	return externalIP, nil
}

func (client *Client) UpdateLoadBalancer(ctx context.Context, lbPoolName string,
	ips []string, internalPort int32) error {

	client.rwLock.Lock()
	defer client.rwLock.Unlock()

	_, err := client.updateLoadBalancerPool(ctx, lbPoolName, ips, internalPort)
	if err != nil {
		return fmt.Errorf("unable to update load balancer pool [%s]: [%v]", lbPoolName, err)
	}

	return nil
}

// DeleteLoadBalancer : create a new load balancer pool and virtual service pointing to it
func (client *Client) DeleteLoadBalancer(ctx context.Context, virtualServiceNamePrefix string,
	lbPoolNamePrefix string) error {

	client.rwLock.Lock()
	defer client.rwLock.Unlock()

	// TODO: try to continue in case of errors
	var err error

	for _, suffix := range []string{"http", "https"} {

		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, suffix)
		lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, suffix)

		// get external IP
		rdeVIP := ""
		dnatRuleName := ""
		if client.OneArm != nil {
			dnatRuleName = getDNATRuleName(virtualServiceName)
			dnatRuleRef, err := client.getNATRuleRef(ctx, dnatRuleName)
			if err != nil {
				return fmt.Errorf("unable to get dnat rule ref for nat rule [%s]: [%v]", dnatRuleName, err)
			}
			rdeVIP = dnatRuleRef.ExternalIP
		} else {
			vsSummary, err := client.getVirtualService(ctx, virtualServiceName)
			if err != nil {
				return fmt.Errorf("unable to get summary for LB Virtual Service [%s]: [%v]",
					virtualServiceName, err)
			}
			if vsSummary == nil {
				return fmt.Errorf("no virtual service with name [%s]", virtualServiceName)
			}
			rdeVIP = vsSummary.VirtualIpAddress
		}

		err = client.deleteVirtualService(ctx, virtualServiceName, false, rdeVIP)
		if err != nil {
			return fmt.Errorf("unable to delete virtual service [%s]: [%v]", virtualServiceName, err)
		}

		err = client.deleteLoadBalancerPool(ctx, lbPoolName, false)
		if err != nil {
			return fmt.Errorf("unable to delete load balancer pool [%s]: [%v]", lbPoolName, err)
		}

		if client.OneArm != nil {
			err = client.deleteDNATRule(ctx, dnatRuleName, false)
			if err != nil {
				return fmt.Errorf("unable to delete dnat rule [%s]: [%v]", dnatRuleName, err)
			}
		}
	}

	return nil
}

// GetLoadBalancer :
func (client *Client) GetLoadBalancer(ctx context.Context, virtualServiceName string) (string, error) {

	vsSummary, err := client.getVirtualService(ctx, virtualServiceName)
	if err != nil {
		return "", fmt.Errorf("unable to get summary for LB Virtual Service [%s]: [%v]",
			virtualServiceName, err)
	}
	if vsSummary == nil {
		return "", nil // this is not an error
	}

	if client.OneArm == nil {
		return vsSummary.VirtualIpAddress, nil
	}

	dnatRuleName := getDNATRuleName(virtualServiceName)
	dnatRuleRef, err := client.getNATRuleRef(ctx, dnatRuleName)
	if err != nil {
		return "", fmt.Errorf("Unable to find dnat rule [%s] for virtual service [%s]: [%v]",
			dnatRuleName, virtualServiceName, err)
	}
	if dnatRuleRef == nil {
		return "", nil // so that a retry creates the DNAT rule
	}

	return dnatRuleRef.ExternalIP, nil
}

// IsNSXTBackedGateway : return true if gateway is backed by NSX-T
func (client *Client) IsNSXTBackedGateway() bool {
	isNSXTBackedGateway :=
		(client.networkBackingType == swaggerClient.NSXT_FIXED_SEGMENT_BackingNetworkType) ||
			(client.networkBackingType == swaggerClient.NSXT_FLEXIBLE_SEGMENT_BackingNetworkType)

	return isNSXTBackedGateway
}

func (client *Client) GetRDEVirtualIps(ctx context.Context) ([]string, string, *swaggerClient.DefinedEntity, error) {
	defEnt, _, etag, err := client.apiClient.DefinedEntityApi.GetDefinedEntity(ctx, client.ClusterID)
	if err != nil {
		return nil, "", nil, fmt.Errorf("error when getting defined entity: [%v]", err)
	}
	statusEntry, ok := defEnt.Entity["status"]
	if !ok {
		return nil, "", nil, fmt.Errorf("could not find 'status' entry in defined entity")
	}
	statusMap, ok := statusEntry.(map[string]interface{})
	if !ok {
		return nil, "", nil, fmt.Errorf("unable to convert [%T] to map", statusEntry)
	}
	virtualIpInterfaces := statusMap["virtual_IPs"]
	if virtualIpInterfaces == nil {
		return make([]string, 0), etag, &defEnt, nil
	}

	virtualIpInterfacesSlice, ok := virtualIpInterfaces.([]interface{})
	if !ok {
		return nil, "", nil, fmt.Errorf("unable to convert [%T] to slice of interface", virtualIpInterfaces)
	}
	virtualIpStrs := make([]string, len(virtualIpInterfacesSlice))
	for ind, ipInterface := range virtualIpInterfacesSlice {
		currIp, ok := ipInterface.(string)
		if !ok {
			return nil, "", nil, fmt.Errorf("unable to convert [%T] to string", ipInterface)
		}
		virtualIpStrs[ind] = currIp
	}
	return virtualIpStrs, etag, &defEnt, nil
}

// This function will modify the passed in defEnt
func (client *Client) updateRDEVirtualIps(ctx context.Context, updatedIps []string, etag string,
	defEnt *swaggerClient.DefinedEntity) (*http.Response, error) {
	statusEntry, ok := defEnt.Entity["status"]
	if !ok {
		return nil, fmt.Errorf("could not find 'status' entry in defined entity")
	}
	statusMap, ok := statusEntry.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unable to convert [%T] to map", statusEntry)
	}

	statusMap["virtual_IPs"] = updatedIps
	// can pass invokeHooks
	_, httpResponse, err := client.apiClient.DefinedEntityApi.UpdateDefinedEntity(ctx, *defEnt, etag, client.ClusterID, nil)
	if err != nil {
		return httpResponse, fmt.Errorf("error when updating defined entity: [%v]", err)
	}
	return httpResponse, nil
}

func (client *Client) addVirtualIpToRDE(ctx context.Context, addIp string) error {
	for i := 0; i < 10; i++ {
		currIps, etag, defEnt, err := client.GetRDEVirtualIps(ctx)
		if err != nil {
			return fmt.Errorf("error for getting current vips: [%v]", err)
		}

		// check if need to update RDE
		foundAddIp := false
		for _, ip := range currIps {
			if ip == addIp {
				foundAddIp = true
				break
			}
		}
		if foundAddIp {
			return nil // no need to update RDE
		}

		updatedIps := append(currIps, addIp)
		httpResponse, err := client.updateRDEVirtualIps(ctx, updatedIps, etag, defEnt)
		if err != nil {
			if httpResponse.StatusCode == http.StatusPreconditionFailed {
				continue
			}
			return fmt.Errorf("error when adding virtual ip to RDE: [%v]", err)
		}
		return nil
	}
	return fmt.Errorf("unable to update rde due to incorrect etag after 10 tries")
}

func (client *Client) removeVirtualIpFromRDE(ctx context.Context, removeIp string) error {
	for i := 0; i < 10; i++ {
		currIps, etag, defEnt, err := client.GetRDEVirtualIps(ctx)
		if err != nil {
			return fmt.Errorf("error for getting current vips: [%v]", err)
		}
		// currIps is guaranteed not to be nil by GetRDEVirtualIps
		if len(currIps) == 0 {
			return fmt.Errorf("no RDE virtual ips to remove")
		}

		// form updated virtual ip list
		foundIdx := -1
		for idx, ip := range currIps {
			if ip == removeIp {
				foundIdx = idx
				break // for inner loop
			}
		}
		if foundIdx == -1 {
			return nil // no need to update RDE
		}
		updatedIps := append(currIps[:foundIdx], currIps[foundIdx+1:]...)

		httpResponse, err := client.updateRDEVirtualIps(ctx, updatedIps, etag, defEnt)
		if err != nil {
			if httpResponse.StatusCode == http.StatusPreconditionFailed {
				continue
			}
			return fmt.Errorf("error when adding virtual ip to RDE: [%v]", err)
		}
		return nil
	}
	return fmt.Errorf("unable to update rde due to incorrect etag after 10 tries")
}
