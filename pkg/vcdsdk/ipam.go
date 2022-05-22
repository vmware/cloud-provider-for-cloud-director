/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package vcdsdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apparentlymart/go-cidr/cidr"
	swaggerClient "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"k8s.io/klog"
	"net"
)

func getUnusedIPAddressInRange(startIPAddress string, endIPAddress string,
	usedIPAddresses map[string]bool) (string, error) {

	// Note: This is not the best approach and can be optimized further by skipping ranges.
	startIP, _, err := net.ParseCIDR(fmt.Sprintf("%s/32", startIPAddress))
	if err != nil {
		return "", fmt.Errorf("unable to parse start IP CIDR [%s]: [%v]", startIPAddress, err)
	}

	endIP, _, err := net.ParseCIDR(fmt.Sprintf("%s/32", endIPAddress))
	if err != nil {
		return "", fmt.Errorf("unable to parse start IP CIDR [%s]: [%v]", startIPAddress, err)
	}

	freeIP := ""
	for !startIP.Equal(endIP) {
		if !usedIPAddresses[startIP.String()] {
			freeIP = startIP.String()
			break
		}
		startIP = cidr.Inc(startIP)
	}

	// handle the last IP since we skip it above
	if freeIP == "" && startIP.Equal(endIP) {
		if !usedIPAddresses[startIP.String()] {
			freeIP = startIP.String()
		}
	}

	return freeIP, nil
}

// GetUnusedExternalIPAddress returns the first unused IP address in the gateway from an ipamSubnet
// There are races here since there is no 'acquisition' of an IP. However, since k8s retries, it will
// be correct.
func (gm *GatewayManager) GetUnusedExternalIPAddress(ctx context.Context, allowedIPAMSubnetStr string) (string, error) {
	client := gm.Client
	if gm.GatewayRef == nil {
		return "", fmt.Errorf("gateway reference should not be nil")
	}

	// Overall procedure
	// 1. Get all IP ranges in gateway
	// 2. Get all used IP addresses in gateway
	// 3. Loop through allowed ip addresses in allowedIPAMSubnetStr
	// 4. Check if the IP address is unused
	// 5. Verify that the IP address is in at least one of the gateway ranges
	// 6. If so, return it

	// 1. Get all IP ranges in gateway
	edgeGW, resp, err := client.APIClient.EdgeGatewayApi.GetEdgeGateway(ctx, gm.GatewayRef.Id)
	if err != nil {
		return "", fmt.Errorf("unable to retrieve edge gateway details for [%s]: resp [%+v]: [%v]",
			gm.GatewayRef.Name, resp, err)
	}

	ipRangesList := make([]*swaggerClient.IpRanges, 0)
	for _, edgeGWUplink := range edgeGW.EdgeGatewayUplinks {
		for _, subnet := range edgeGWUplink.Subnets.Values {
			ipRangesList = append(ipRangesList, subnet.IpRanges)
		}
	}
	if len(ipRangesList) == 0 {
		return "", fmt.Errorf("unable to get any ipRanges in gateway")
	}

	// 2. Get all used IP addresses in gateway
	usedIPs := make(map[string]bool)
	pageNum := int32(1)
	for {
		gwUsedIPAddresses, resp, err := client.APIClient.EdgeGatewayApi.GetUsedIpAddresses(ctx, pageNum, 25,
			gm.GatewayRef.Id, nil)
		if err != nil {
			return "", fmt.Errorf("unable to get used IP addresses of gateway [%s]: [%+v]: [%v]",
				gm.GatewayRef.Name, resp, err)
		}
		if len(gwUsedIPAddresses.Values) == 0 {
			break
		}

		for _, gwUsedIPAddress := range gwUsedIPAddresses.Values {
			usedIPs[gwUsedIPAddress.IpAddress] = true
		}

		pageNum++
	}

	// 3. Loop through allowed ip addresses in gateway's ipamSubnet
	_, allowedIPAMSubnet, err := net.ParseCIDR(allowedIPAMSubnetStr)
	if err != nil {
		return "", fmt.Errorf("unable to parse CIDR [%s] into a subnet: [%v]", allowedIPAMSubnetStr, err)
	}
	allowedStartIP, allowedEndIP := cidr.AddressRange(allowedIPAMSubnet)

	freeIP := ""
	for !allowedStartIP.Equal(allowedEndIP) {
		// 4. Check if the IP address is unused &&
		// 5. Verify that the IP address is in at least one of the gateway ranges
		if !usedIPs[allowedStartIP.String()] && checkIfIPIsAvailable(allowedStartIP.String(), ipRangesList) {
			freeIP = allowedStartIP.String()
			break
		}

		allowedStartIP = cidr.Inc(allowedStartIP)
	}

	// If no free IP has been obtained so far, handle the last IP the same way
	if freeIP == "" && allowedStartIP.Equal(allowedEndIP) {
		// 4. Check if the IP address is unused &&
		// 5. Verify that the IP address is in at least one of the gateway ranges
		if !usedIPs[allowedStartIP.String()] && checkIfIPIsAvailable(allowedStartIP.String(), ipRangesList) {
			freeIP = allowedStartIP.String()
		}
	}

	if freeIP == "" {
		return "", fmt.Errorf("unable to obtain free IP from gateway [%s]; all are used",
			gm.GatewayRef.Name)
	}
	klog.Infof("Using unused IP [%s] on gateway [%v]\n", freeIP, gm.GatewayRef.Name)

	return freeIP, nil
}

func checkIfIPIsAvailable(ipStr string, ipRangesList []*swaggerClient.IpRanges) bool {
	for _, ipRanges := range ipRangesList {
		for _, ipRange := range ipRanges.Values {
			ip := net.ParseIP(ipStr)
			startIP := net.ParseIP(ipRange.StartAddress)
			endIP := net.ParseIP(ipRange.EndAddress)

			if bytes.Compare(ip, startIP) >= 0 && bytes.Compare(ip, endIP) <= 0 {
				return true
			}
		}
	}

	return false
}

func (gm *GatewayManager) GetUnusedInternalIPAddress(ctx context.Context, oneArm *OneArm) (string, error) {

	if oneArm == nil {
		return "", fmt.Errorf("unable to get unused internal IP address as oneArm is nil")
	}
	client := gm.Client
	if gm.GatewayRef == nil {
		return "", fmt.Errorf("gateway reference should not be nil")
	}

	usedIPAddresses := make(map[string]bool)
	pageNum := int32(1)
	for {
		lbVSSummaries, resp, err := client.APIClient.EdgeGatewayLoadBalancerVirtualServicesApi.GetVirtualServiceSummariesForGateway(
			ctx, pageNum, 25, gm.GatewayRef.Id, nil)
		if err != nil {
			return "", fmt.Errorf("unable to get virtual service summaries for gateway [%s]: resp: [%v]: [%v]",
				gm.GatewayRef.Name, resp, err)
		}
		if len(lbVSSummaries.Values) == 0 {
			break
		}
		for _, lbVSSummary := range lbVSSummaries.Values {
			usedIPAddresses[lbVSSummary.VirtualIpAddress] = true
		}

		pageNum++
	}

	freeIP, err := getUnusedIPAddressInRange(oneArm.StartIP, oneArm.EndIP, usedIPAddresses)
	if err != nil {
		return "", fmt.Errorf("error in finding unused IP address in range [%s-%s]: [%v]",
			oneArm.StartIP, oneArm.EndIP, err)
	}
	if freeIP == "" {
		return "", fmt.Errorf("unable to find unused IP address in range [%s-%s]",
			oneArm.StartIP, oneArm.EndIP)
	}

	return freeIP, nil
}
