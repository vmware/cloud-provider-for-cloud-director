package utils

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/testingsdk"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
	v1 "k8s.io/api/core/v1"
	"strings"
)

const (
	VCloudZoneConfigMapName      = "vcloud-capvcd-zones"
	VCloudZoneConfigMapNamespace = "kube-system"
)

func NewTestClient(host, org, userOrg, ovdcIdentifier, username, token, clusterId string, getVdcClient bool) (*testingsdk.TestClient, error) {
	vcdAuthParams := &testingsdk.VCDAuthParams{
		Host:           host,
		OrgName:        org,
		UserOrg:        userOrg,
		OvdcIdentifier: ovdcIdentifier,
		Username:       username,
		RefreshToken:   token,
		GetVdcClient:   getVdcClient,
	}
	return testingsdk.NewTestClient(vcdAuthParams, clusterId)
}

func GetLBPoolNamePrefix(service *v1.Service, clusterId string) string {
	return fmt.Sprintf("ingress-pool-%s-%s", service.Name, getTrimmedClusterID(clusterId))
}

func GetVirtualServicePrefix(service *v1.Service, clusterId string) string {
	return fmt.Sprintf("ingress-vs-%s-%s", service.Name, getTrimmedClusterID(clusterId))
}

func GetPortDetailsList(svc *v1.Service) []vcdsdk.PortDetails {
	portDetailsList := make([]vcdsdk.PortDetails, len(svc.Spec.Ports))
	for idx, port := range svc.Spec.Ports {
		portDetailsList[idx] = vcdsdk.PortDetails{
			PortSuffix:   port.Name,
			ExternalPort: port.Port,
			InternalPort: port.NodePort,
			Protocol:     string(port.Protocol),
		}
	}
	return portDetailsList
}

func HasVCDResourcesForApplicationLB(ctx context.Context, testClient *testingsdk.TestClient, gatewayMgr *vcdsdk.GatewayManager, oneArm *vcdsdk.OneArm, virtualServiceNamePrefix string,
	lbPoolNamePrefix string, portDetailsList []vcdsdk.PortDetails) (bool, error) {
	vsFound, lbPoolFound, dnatRuleFound, appPortProfileFound := false, false, false, false
	for _, portDetails := range portDetailsList {
		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portDetails.PortSuffix)
		lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portDetails.PortSuffix)
		vsSummary, err := gatewayMgr.GetVirtualService(ctx, virtualServiceName)
		if err != nil {
			return false, fmt.Errorf("error getting virtual service [%s]: [%v]", virtualServiceName, err)
		}

		if vsSummary != nil {
			vsFound = true
		}

		lbPool, err := gatewayMgr.GetLoadBalancerPool(ctx, lbPoolName)
		if err != nil && err != govcd.ErrorEntityNotFound {
			return false, fmt.Errorf("error getting load balancer pool for [%s]: [%v]", lbPoolName, err)
		}

		if lbPool != nil {
			lbPoolFound = true
		}

		if oneArm != nil {
			dnatRuleName := vcdsdk.GetDNATRuleName(virtualServiceName)
			dnatRuleRef, err := gatewayMgr.GetNATRuleRef(ctx, dnatRuleName)
			if err != nil && err != govcd.ErrorEntityNotFound {
				return false, fmt.Errorf("error getting dnat rule ref for [%s]: [%v]", dnatRuleName, err)
			}
			if dnatRuleRef != nil {
				dnatRuleFound = true
			}

			appPortProfileName := vcdsdk.GetAppPortProfileName(dnatRuleName)
			org, err := testClient.VcdClient.VCDClient.GetOrgByName(testClient.VcdClient.ClusterOrgName)
			if err != nil {
				return false, fmt.Errorf("unable to find org [%s] by name: [%v]", testClient.VcdClient.ClusterOrgName, err)
			}
			appPortProfile, err := org.GetNsxtAppPortProfileByName(appPortProfileName, types.ApplicationPortProfileScopeTenant)
			// We are doing a string check here because it returns a formatted error instead of govcd.ErrorEntityNotFound
			if err != nil && !strings.Contains(err.Error(), "[ENF] entity not found") {
				return false, fmt.Errorf("unable to get app port profile [%s]: [%v]", appPortProfileName, err)
			}
			if appPortProfile != nil {
				appPortProfileFound = true
			}
		}
	}

	// All are present then this is a valid case
	if vsFound && lbPoolFound {
		if oneArm != nil {
			return dnatRuleFound && appPortProfileFound, nil
		}
		return true, nil
	}

	// To handle the case that all resources are deleted, we can if none are found then we can simply return false
	if !vsFound && !lbPoolFound {
		if oneArm != nil && !dnatRuleFound && !appPortProfileFound {
			// We want to return false, if all 4 LB components are not to be found.
			return false, nil
		}
		return false, nil
	}

	// If some are present and some are not, then it's considered invalid and there should be an error as this searches for all resources.
	return false, fmt.Errorf("either one or more resource was present but not all - vs: [%v], lb: [%v], dnat: [%v], appPort: [%v], oneArm: [%v]",
		vsFound, lbPoolFound, dnatRuleFound, appPortProfileFound, oneArm)
}

func getTrimmedClusterID(clusterId string) string {
	return strings.TrimPrefix(clusterId, clusterUrnPrefix)
}

func GetVDCForZone(tc *testingsdk.TestClient, zoneName string) (string, error) {
	zoneConfigMap, err := tc.GetConfigMap(VCloudZoneConfigMapNamespace, VCloudZoneConfigMapName)
	if err != nil {
		return "", fmt.Errorf("failed to get the zone config map: [%v]", err)
	}
	if zoneConfigMap == nil {
		return "", fmt.Errorf("zone config map is nil")
	}
	zoneToOVDC, err := tc.GetZoneMapFromZoneConfigMap(zoneConfigMap)
	if err != nil {
		return "", fmt.Errorf("failed to get zone name to ovdc name mapping: [%v]", err)
	}

	ovdcName, ok := zoneToOVDC[zoneName]
	if !ok {
		return "", fmt.Errorf("zone config map doesn't have an entry for the zone name")
	}
	return ovdcName, nil
}
