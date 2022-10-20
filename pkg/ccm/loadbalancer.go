//go:build !testing
// +build !testing

/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package ccm

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/cpisdk"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/util"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/cloud-provider-for-cloud-director/release"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	cloudProvider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"strconv"
	"strings"
)

const (
	sslPortsAnnotation     = `service.beta.kubernetes.io/vcloud-avi-ssl-ports`
	sslCertAliasAnnotation = `service.beta.kubernetes.io/vcloud-avi-ssl-cert-alias`
	// TODO: Update controlPlaneLabel to use default K8s constants if available
	controlPlaneLabel = `node-role.kubernetes.io/control-plane`
)

//LBManager -
type LBManager struct {
	gatewayManager               *vcdsdk.GatewayManager
	vcdClient                    *vcdsdk.Client
	kubeClient                   *kubernetes.Clientset
	namespace                    string
	CertificateAlias             string
	OneArm                       *vcdsdk.OneArm
	ovdcNetworkName              string
	ipamSubnet                   string
	clusterID                    string
	EnableVirtualServiceSharedIP bool
}

func newLoadBalancer(vcdClient *vcdsdk.Client, certAlias string, oneArm *vcdsdk.OneArm,
	ovdcNetworkName string, ipamSubnet string, clusterID string, enableVirtualServiceSharedIP bool) cloudProvider.LoadBalancer {

	return &LBManager{
		vcdClient:                    vcdClient,
		kubeClient:                   GetK8SClient(),
		namespace:                    "default",
		CertificateAlias:             certAlias,
		OneArm:                       oneArm,
		ovdcNetworkName:              ovdcNetworkName,
		ipamSubnet:                   ipamSubnet,
		clusterID:                    clusterID,
		EnableVirtualServiceSharedIP: enableVirtualServiceSharedIP,
	}
}

// TODO: Should we add errors from this method to errorSet as it gives a few hard error returns?
func (lb *LBManager) addLBResourcesToRDE(ctx context.Context, resourcesAllocated *util.AllocatedResourcesMap, externalIP string) error {
	rdeManager := vcdsdk.NewRDEManager(lb.vcdClient, lb.clusterID, release.CloudControllerManagerName, release.CpiVersion)
	for _, key := range []string{vcdsdk.VcdResourceDNATRule, vcdsdk.VcdResourceLoadBalancerPool, vcdsdk.VcdResourceAppPortProfile, vcdsdk.VcdResourceVirtualService} {
		if values := resourcesAllocated.Get(key); values != nil {
			for _, value := range values {
				var additionalDetails map[string]interface{}
				var err error
				additionalDetails = nil
				if key == vcdsdk.VcdResourceVirtualService {
					additionalDetails = map[string]interface{}{
						"virtualIP": externalIP,
					}
					cpiRdeManager := cpisdk.NewCPIRDEManager(rdeManager)
					err = cpiRdeManager.AddVIPToVCDResourceSet(ctx, value.Name, value.Id, externalIP)
				} else {
					err = rdeManager.AddToVCDResourceSet(ctx, vcdsdk.ComponentCPI, key,
						value.Name, value.Id, additionalDetails)
				}
				if err != nil {
					return fmt.Errorf(
						"failed to add resource [%s] of type [%s] to VCDResourceSet of RDE [%s]: [%v]",
						value.Name, key, lb.clusterID, err)
				}
			}
		}
	}
	return nil
}

// TODO: Should we add errors from this method to errorSet as it gives a few hard error returns?
func (lb *LBManager) removeLBResourcesFromRDE(ctx context.Context, resourcesDeallocated *util.AllocatedResourcesMap) error {
	rdeManager := vcdsdk.NewRDEManager(lb.vcdClient, lb.clusterID, release.CloudControllerManagerName, release.CpiVersion)
	for _, key := range []string{vcdsdk.VcdResourceDNATRule, vcdsdk.VcdResourceVirtualService,
		vcdsdk.VcdResourceLoadBalancerPool, vcdsdk.VcdResourceAppPortProfile} {
		if values := resourcesDeallocated.Get(key); values != nil {
			for _, value := range values {
				err := rdeManager.RemoveFromVCDResourceSet(ctx, vcdsdk.ComponentCPI, key, value.Name)
				if err != nil {
					return fmt.Errorf(
						"failed to add resource [%s] of type [%s] to VCDResourceSet of RDE [%s]: [%v]",
						value.Name, key, lb.clusterID, err)
				}
			}
		}
	}
	return nil
}

func (lb *LBManager) getNodeIPs(ctx context.Context) ([]string, error) {
	nodes, err := lb.kubeClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("unable to get nodes of cluster: [%v]", err)
	}

	nodeIPs := make([]string, len(nodes.Items))
	for idx, node := range nodes.Items {
		nodeIPs[idx] = node.Status.Addresses[0].Address
	}

	return nodeIPs, nil
}

// EnsureLoadBalancer creates a new load balancer 'name', or updates the existing one.
// Returns the status of the balancer. Implementations must treat the *v1.Service and *v1.Node
// parameters as read-only and not modify them.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (lb *LBManager) EnsureLoadBalancer(ctx context.Context, clusterName string,
	service *v1.Service, nodes []*v1.Node) (lbs *v1.LoadBalancerStatus, err error) {

	if err = lb.vcdClient.RefreshBearerToken(); err != nil {
		return nil, fmt.Errorf("error while obtaining access token: [%v]", err)
	}
	nodeIPs := lb.getWorkerNodeInternalIps(nodes)
	return lb.createLoadBalancer(ctx, service, nodeIPs)
}

func (lb *LBManager) getNodeInternalIps(nodes []*v1.Node) []string {
	nodeIps := make([]string, len(nodes))
	for i, node := range nodes {
		for _, addr := range node.Status.Addresses {
			if addr.Type == v1.NodeInternalIP {
				nodeIps[i] = addr.Address
				break
			}
		}
	}
	return nodeIps
}

func (lb *LBManager) getServicePortMap(service *v1.Service) (map[string]int32, map[string]int32, map[string]string) {
	typeToInternalPort := make(map[string]int32)
	typeToExternalPort := make(map[string]int32)
	nameToProtocol := make(map[string]string)
	for _, port := range service.Spec.Ports {
		typeToInternalPort[strings.ToLower(port.Name)] = port.NodePort
		typeToExternalPort[strings.ToLower(port.Name)] = port.Port
		if port.AppProtocol != nil {
			nameToProtocol[strings.ToLower(port.Name)] = strings.ToUpper(*port.AppProtocol)
		}
	}
	return typeToInternalPort, typeToExternalPort, nameToProtocol
}

func (lb *LBManager) getWorkerNodeInternalIps(nodes []*v1.Node) []string {
	var workerNodeInternalIps []string
	for _, node := range nodes {
		nodeLabelMap := node.ObjectMeta.Labels
		// If we can find it from the nodeLabelMap we will take it, but if it is missing labels, we will just let it go instead of adding it.
		if nodeLabelMap != nil {
			if _, ok := nodeLabelMap[controlPlaneLabel]; !ok {
				for _, addr := range node.Status.Addresses {
					if addr.Type == v1.NodeInternalIP {
						klog.Infof("Worker Node Internal IP found: %s", addr.Address)
						workerNodeInternalIps = append(workerNodeInternalIps, addr.Address)
						break
					}
				}
			}
		}
	}
	return workerNodeInternalIps
}

// UpdateLoadBalancer updates hosts under the specified load balancer.
// Implementations must treat the *v1.Service and *v1.Node
// parameters as read-only and not modify them.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (lb *LBManager) UpdateLoadBalancer(ctx context.Context, clusterName string,
	service *v1.Service, nodes []*v1.Node) (err error) {

	if err = lb.vcdClient.RefreshBearerToken(); err != nil {
		return fmt.Errorf("error while obtaining access token: [%v]", err)
	}

	nodeIps := lb.getWorkerNodeInternalIps(nodes)
	klog.Infof("UpdateLoadBalancer Node Ips: %v", nodeIps)

	lbPoolNamePrefix := lb.getLBPoolNamePrefix(ctx, service)
	virtualServiceNamePrefix := lb.getVirtualServicePrefix(ctx, service)
	typeToInternalPortMap, typeToExternalPort, nameToProtocol := lb.getServicePortMap(service)
	rdeManager := vcdsdk.NewRDEManager(lb.vcdClient, lb.clusterID, release.CloudControllerManagerName, release.CpiVersion)
	cpiRdeManager := cpisdk.NewCPIRDEManager(rdeManager)

	for portName, internalPort := range typeToInternalPortMap {
		lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portName)
		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portName)
		externalPort := typeToExternalPort[portName]
		gm, err := vcdsdk.NewGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet)
		if err != nil {
			return fmt.Errorf("error while creating GatewayManager: [%v]", err)
		}
		klog.Infof("Updating pool [%s] with port [%s:%d]", lbPoolName, portName, internalPort)
		protocol, _ := nameToProtocol[portName]
		resourcesAllocated := &util.AllocatedResourcesMap{}
		vip, err := gm.UpdateLoadBalancer(ctx, lbPoolName, virtualServiceName, nodeIps, internalPort,
			externalPort, lb.OneArm, lb.EnableVirtualServiceSharedIP, protocol, resourcesAllocated)
		// TODO: Should we record this error as well?
		if rdeErr := lb.addLBResourcesToRDE(ctx, resourcesAllocated, vip); rdeErr != nil {
			return fmt.Errorf("failed to add load balancer resources to RDE [%s]: [%v]", lb.clusterID, err)
		}

		vsSummary, getVsErr := gm.GetVirtualService(ctx, virtualServiceName)
		if getVsErr != nil {
			return fmt.Errorf("failed to get virtual service [%s], [%v]", virtualServiceName, getVsErr)
		}

		if vsSummary == nil {
			return fmt.Errorf("virtual service [%s] does not exist", virtualServiceName)
		}

		if err != nil {
			addToErrorSetErr := cpiRdeManager.AddToErrorSetWithNameAndId(ctx, cpisdk.UpdateLoadbalancerError, vsSummary.Id, vsSummary.Name, err.Error())
			if addToErrorSetErr != nil {
				klog.Errorf("error adding CPI error [%s] to RDE: [%s], [%v]", cpisdk.UpdateLoadbalancerError, lb.clusterID, addToErrorSetErr)
			}
			return fmt.Errorf("unable to update pool [%s] with port [%s:%d]: [%v]", lbPoolName, portName,
				internalPort, err)
		}

		// TODO: This may need to be optimized in the future as we are making len(ports) API calls
		err = cpiRdeManager.AddToEventSetWithNameAndId(ctx, cpisdk.UpdatedLoadbalancer, vsSummary.Id, vsSummary.Name, fmt.Sprintf("Successfully updated loadbalancer with virtual service name [%s]: ", vsSummary.Name))
		if err != nil {
			klog.Errorf("error adding CPI event [%s] to RDE: [%s], [%v]", cpisdk.UpdatedLoadbalancer, lb.clusterID, err)
		}

		err = cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, cpisdk.UpdateLoadbalancerError, vsSummary.Id, vsSummary.Name)
		if err != nil {
			klog.Errorf("there was an error removing CPI error [%s] from RDE [%s], [%v]", cpisdk.UpdateLoadbalancerError, lb.clusterID, err)
		}
	}

	return nil
}

// EnsureLoadBalancerDeleted deletes the specified load balancer if it
// exists, returning nil if the load balancer specified either didn't exist or
// was successfully deleted.
// This construction is useful because many cloud providers' load balancers
// have multiple underlying components, meaning a Get could say that the LB
// doesn't exist even if some part of it is still laying around.
// Implementations must treat the *v1.Service parameter as read-only and not modify it.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (lb *LBManager) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string,
	service *v1.Service) error {

	if err := lb.vcdClient.RefreshBearerToken(); err != nil {
		return fmt.Errorf("error while obtaining access token: [%v]", err)
	}
	return lb.deleteLoadBalancer(ctx, service)
}

func (lb *LBManager) getLoadBalancer(ctx context.Context,
	service *v1.Service) (status *v1.LoadBalancerStatus, portNameToIPMap map[string]string, err error) {

	virtualServiceNamePrefix := lb.getLoadBalancerPrefix(ctx, service)
	virtualIP := ""
	gm, err := vcdsdk.NewGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet)
	if err != nil {
		return nil, nil, fmt.Errorf("error while creating GatewayManager: [%v]", err)
	}

	cpiRdeManager := cpisdk.NewCPIRDEManager(vcdsdk.NewRDEManager(
		lb.vcdClient, lb.clusterID, release.CloudControllerManagerName, release.CpiVersion))

	portNameToIP := make(map[string]string)
	ingressVirtualIP := ""
	for _, port := range service.Spec.Ports {
		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, port.Name)
		lbPoolNamePrefix := lb.getLBPoolNamePrefix(ctx, service)
		lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, port.Name)
		virtualIP, _, err = gm.GetLoadBalancer(ctx, virtualServiceName, lbPoolName, lb.OneArm)
		if err != nil {
			addToErrorSetErr := cpiRdeManager.AddToErrorSetWithNameAndId(ctx, cpisdk.GetLoadbalancerError, "", virtualServiceName, err.Error())
			if addToErrorSetErr != nil {
				klog.Errorf("unable to add CPI error [%s] to RDE [%s], [%v]", cpisdk.GetLoadbalancerError, lb.clusterID, addToErrorSetErr)
			}
			return nil, nil,
				fmt.Errorf("unable to get virtual service summary for [%s]: [%v]",
					virtualServiceName, err)
		}
		removeErr := cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, cpisdk.GetLoadbalancerError, "", virtualServiceName)
		if removeErr != nil {
			klog.Errorf("there was an error removing CPI error [%s] from RDE [%s], [%v]", cpisdk.GetLoadbalancerError, lb.clusterID, err)
		}
		portNameToIP[port.Name] = virtualIP
		if virtualIP != "" {
			if ingressVirtualIP != "" && ingressVirtualIP != virtualIP {
				return nil, nil,
					fmt.Errorf("more than one virtual ip found: [%s] and [%s]", virtualIP, ingressVirtualIP)
			}
			ingressVirtualIP = virtualIP
		}
	}
	if len(portNameToIP) == 0 {
		// this implies that no port was specified
		return nil, nil, nil
	}

	return &v1.LoadBalancerStatus{
		Ingress: []v1.LoadBalancerIngress{
			{
				IP: ingressVirtualIP,
			},
		},
	}, portNameToIP, nil

}

// GetLoadBalancer returns whether the specified load balancer exists, and
// if so, what its status is.
// Implementations must treat the *v1.Service parameter as read-only and not modify it.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (lb *LBManager) GetLoadBalancer(ctx context.Context, clusterName string,
	service *v1.Service) (status *v1.LoadBalancerStatus, exists bool, err error) {

	if err = lb.vcdClient.RefreshBearerToken(); err != nil {
		return nil, false, fmt.Errorf("error while obtaining access token: [%v]", err)
	}
	status, portNameToIPMap, err := lb.getLoadBalancer(ctx, service)
	if err != nil {
		return nil, false, fmt.Errorf("error while getting load balancer: [%v]", err)
	}
	if portNameToIPMap == nil {
		return nil, false, nil
	}

	for _, ip := range portNameToIPMap {
		if ip == "" {
			return nil, false, nil // returning false to retry creation
		}
	}
	return status, true, nil
}

// getTrimmedClusterID: this is a mitigation to not overflow VCD name length limits. There is a clearer
// fix needed in the future. Cover all cluster prefixes.
func (lb *LBManager) getTrimmedClusterID() string {
	clusterID := lb.clusterID
	for _, prefix := range []string{
		"urn:vcloud:entity:vmware:",
		"urn:vcloud:entity:cse:nativeCluster:",
	} {
		clusterID = strings.TrimPrefix(clusterID, prefix)
	}
	return clusterID
}

func (lb *LBManager) getLBPoolNamePrefix(_ context.Context, service *v1.Service) string {
	return fmt.Sprintf("ingress-pool-%s-%s", service.Name, lb.getTrimmedClusterID())
}

func (lb *LBManager) getLoadBalancerPrefix(_ context.Context, service *v1.Service) string {
	return fmt.Sprintf("ingress-vs-%s-%s", service.Name, lb.getTrimmedClusterID())
}

func (lb *LBManager) getVirtualServicePrefix(_ context.Context, service *v1.Service) string {
	return fmt.Sprintf("ingress-vs-%s-%s", service.Name, lb.getTrimmedClusterID())
}

// GetLoadBalancerName returns the name of the load balancer. Implementations must treat the
// *v1.Service parameter as read-only and not modify it.
func (lb *LBManager) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
	return lb.getLoadBalancerPrefix(ctx, service)
}

func (lb *LBManager) deleteLoadBalancer(ctx context.Context, service *v1.Service) error {

	virtualServiceName := lb.getVirtualServicePrefix(ctx, service)
	lbPoolNamePrefix := lb.getLBPoolNamePrefix(ctx, service)
	klog.Infof("Deleting virtual service [%s] and lb pool [%s]", virtualServiceName, lbPoolNamePrefix)

	portDetailsList := make([]vcdsdk.PortDetails, len(service.Spec.Ports))
	for idx, port := range service.Spec.Ports {
		portDetailsList[idx] = vcdsdk.PortDetails{
			PortSuffix:   port.Name,
			ExternalPort: port.Port,
			InternalPort: port.NodePort,
			Protocol:     string(port.Protocol),
			// no need to set UseSSL for deletion
		}
	}
	klog.Infof("Deleting loadbalancer for ports [%#v]\n", portDetailsList)

	gm, err := vcdsdk.NewGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet)
	if err != nil {
		return fmt.Errorf("error while creating GatewayManager: [%v]", err)
	}
	resourcesDeallocated := &util.AllocatedResourcesMap{}
	vip, err := gm.DeleteLoadBalancer(ctx, virtualServiceName, lbPoolNamePrefix, portDetailsList, lb.OneArm, resourcesDeallocated)
	if rdeErr := lb.removeLBResourcesFromRDE(ctx, resourcesDeallocated); rdeErr != nil {
		klog.Errorf("failed to remove loadbalancer resources from RDE [%s]: [%v]", lb.clusterID, rdeErr)
		return fmt.Errorf("failed to remove loadbalancer resources from RDE [%s]: [%v]", lb.clusterID, rdeErr)
	}

	cpiRdeManager := cpisdk.NewCPIRDEManager(vcdsdk.NewRDEManager(
		lb.vcdClient, lb.clusterID, release.CloudControllerManagerName, release.CpiVersion))

	if err != nil {
		addToErrorSetErr := cpiRdeManager.AddToErrorSetWithNameAndId(ctx, cpisdk.DeleteLoadbalancerError, "", virtualServiceName, err.Error())
		if addToErrorSetErr != nil {
			klog.Errorf("error adding CPI error [%s] to RDE: [%s], [%v]", cpisdk.DeleteLoadbalancerError, lb.clusterID, addToErrorSetErr)
		}
		return fmt.Errorf("unable to delete load balancer for virtual-service [%s] and lb pool [%s]: [%v]",
			virtualServiceName, lbPoolNamePrefix, err)
	}

	if err := cpiRdeManager.RemoveVirtualIpFromRDE(ctx, vip); err != nil {
		addToErrorSetErr := cpiRdeManager.AddToErrorSet(ctx, cpisdk.RemoveVIPFromRdeError, lb.clusterID, err.Error())
		if addToErrorSetErr != nil {
			klog.Errorf("unable to add CPI error [%s] to RDE: [%s], [%v]", cpisdk.RemoveVIPFromRdeError, lb.clusterID, addToErrorSetErr)
		}
		klog.Errorf("failed to remove virtual IP [%s] from the RDE [%s]: [%v]", vip, lb.clusterID, err)
	}

	err = cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, cpisdk.RemoveVIPFromRdeError, lb.clusterID, "")
	if err != nil {
		klog.Errorf("error removing CPI error [%s] from RDE: [%v]", cpisdk.RemoveVIPFromRdeError, lb.clusterID, err)
	}

	err = cpiRdeManager.AddToEventSetWithNameAndId(ctx, cpisdk.DeletedLoadbalancer, "", virtualServiceName, fmt.Sprintf("Successfully deleted loadbalancer associated with [%s], deleted external IP [%s]", lb.clusterID, vip))
	if err != nil {
		klog.Errorf("error adding CPI event [%s] to RDE: [%v]", cpisdk.DeletedLoadbalancer, err)
	}

	err = cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, cpisdk.DeleteLoadbalancerError, "", virtualServiceName)
	if err != nil {
		klog.Errorf("there was an error removing CPI error [%s] from RDE [%s], [%v]", cpisdk.DeleteLoadbalancerError, lb.clusterID, err)
	}

	return nil
}

func getSSLPorts(service *v1.Service) ([]int32, error) {
	sslPortsString, ok := service.Annotations[sslPortsAnnotation]
	if !ok {
		return nil, nil
	}

	portStrings := strings.Split(sslPortsString, ",")
	ports := make([]int32, len(portStrings))
	for idx, portStr := range portStrings {
		port, err := strconv.ParseInt(portStr, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("unable to convert [%s] to int32: [%v]", portStr, err)
		}
		ports[idx] = int32(port)
	}

	return ports, nil
}

func getSSLCertAlias(service *v1.Service) string {
	sslCertAlias, ok := service.Annotations[sslCertAliasAnnotation]
	if !ok {
		return ""
	}

	return sslCertAlias
}

func (lb *LBManager) createLoadBalancer(ctx context.Context, service *v1.Service,
	nodeIPs []string) (*v1.LoadBalancerStatus, error) {

	lbPoolNamePrefix := lb.getLBPoolNamePrefix(ctx, service)
	virtualServiceNamePrefix := lb.getVirtualServicePrefix(ctx, service)
	lbStatus, portNameToIPMap, err := lb.getLoadBalancer(ctx, service)
	rdeManager := vcdsdk.NewRDEManager(lb.vcdClient, lb.clusterID, release.CloudControllerManagerName, release.CpiVersion)
	cpiRdeManager := cpisdk.NewCPIRDEManager(rdeManager)
	if err != nil {
		addToErrorSetErr := cpiRdeManager.AddToErrorSetWithNameAndId(ctx, cpisdk.GetLoadbalancerError, "", virtualServiceNamePrefix, err.Error())
		if addToErrorSetErr != nil {
			klog.Errorf("error adding CPI error [%s] to the RDE [%s], [%v]", cpisdk.GetLoadbalancerError, lb.clusterID, addToErrorSetErr)
		}
		return nil, fmt.Errorf("unexpected error while querying for loadbalancer: [%v]", err)
	}

	// TODO: is it worth adding loadbalancer query to event?
	removeErr := cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, cpisdk.GetLoadbalancerError, "", virtualServiceNamePrefix)
	if removeErr != nil {
		klog.Errorf("error adding CPI error [%s] to the RDE [%s], [%v]", cpisdk.GetLoadbalancerError, lb.clusterID, removeErr)
	}
	gm, err := vcdsdk.NewGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet)
	if err != nil {
		return nil, fmt.Errorf("error while creating GatewayManager: [%v]", err)
	}

	// golang doesn't have the set data structure
	lbExists := true
	for _, ip := range portNameToIPMap {
		if ip == "" {
			lbExists = false
			break
		}
	}

	if lbExists {
		// Update load balancer if there are changes in service properties
		typeToInternalPortMap, typeToExternalPortMap, nameToProtocol := lb.getServicePortMap(service)
		for portName, internalPort := range typeToInternalPortMap {
			lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portName)
			virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portName)
			externalPort := typeToExternalPortMap[portName]
			protocol, _ := nameToProtocol[portName]
			klog.Infof("Updating pool [%s] with port [%s:%d:%d]", lbPoolName, portName, internalPort, externalPort)
			resourcesAllocated := &util.AllocatedResourcesMap{}
			vip, err := gm.UpdateLoadBalancer(ctx, lbPoolName, virtualServiceName, nodeIPs, internalPort,
				externalPort, lb.OneArm, lb.EnableVirtualServiceSharedIP, protocol, resourcesAllocated)
			if rdeErr := lb.addLBResourcesToRDE(ctx, resourcesAllocated, vip); rdeErr != nil {
				return nil, fmt.Errorf("failed to update RDE [%s] with load balancer resources: [%v]", lb.clusterID, err)
			}

			vsSummary, getVsErr := gm.GetVirtualService(ctx, virtualServiceName)
			if getVsErr != nil {
				return nil, fmt.Errorf("failed to get virtual service [%s], [%v]", virtualServiceName, getVsErr)
			}

			if vsSummary == nil {
				return nil, fmt.Errorf("virtual service [%s] does not exist", virtualServiceName)
			}

			if err != nil {
				addToErrorSetErr := cpiRdeManager.AddToErrorSetWithNameAndId(ctx, cpisdk.UpdateLoadbalancerError, vsSummary.Id, vsSummary.Name, err.Error())
				if addToErrorSetErr != nil {
					klog.Errorf("error adding CPI error [%s] to RDE: [%s], [%v]", cpisdk.UpdateLoadbalancerError, lb.clusterID, addToErrorSetErr)
				}
				return nil, fmt.Errorf("unable to update pool [%s] with port [%s:%d:%d]: [%v]", lbPoolName, portName,
					internalPort, externalPort, err)
			}

			// TODO: This may need to be optimized in the future as we are making len(ports) API calls
			err = cpiRdeManager.AddToEventSetWithNameAndId(ctx, cpisdk.UpdatedLoadbalancer, vsSummary.Id, vsSummary.Name, fmt.Sprintf("Successfully updated loadbalancer with virtual service name [%s]: ", vsSummary.Name))
			if err != nil {
				klog.Errorf("error adding CPI event [%s] to RDE: [%s], [%v]", cpisdk.UpdatedLoadbalancer, lb.clusterID, err)
			}

			err = cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, cpisdk.UpdateLoadbalancerError, vsSummary.Id, vsSummary.Name)
			if err != nil {
				klog.Errorf("there was an error removing CPI error [%s] from RDE [%s], [%v]", cpisdk.UpdateLoadbalancerError, lb.clusterID, err)
			}
		}
		return lbStatus, nil
	}

	// While creating the lb, even if only one of http/https is remaining and the other is completed,
	// ask for both to be created. The already created one will silently pass.

	portDetailsList := make([]vcdsdk.PortDetails, len(service.Spec.Ports))

	ports, err := getSSLPorts(service)
	if err != nil {
		return nil, fmt.Errorf("unable to get ports from service annotation for [%#v]: [%v]",
			service, err)
	}

	certAlias := getSSLCertAlias(service)
	if certAlias == "" {
		certAlias = lb.CertificateAlias
	}

	// golang doesn't have the set data structure
	portsMap := make(map[int32]bool)
	for _, port := range ports {
		portsMap[port] = true
	}
	for idx, port := range service.Spec.Ports {
		portDetailsList[idx] = vcdsdk.PortDetails{
			PortSuffix:   port.Name,
			ExternalPort: port.Port,
			InternalPort: port.NodePort,
			Protocol:     strings.ToUpper(string(port.Protocol)),
		}
		if port.AppProtocol != nil && *port.AppProtocol != "" {
			portDetailsList[idx].Protocol = strings.ToUpper(*port.AppProtocol)
		}
		if _, ok := portsMap[port.Port]; ok {
			portDetailsList[idx].UseSSL = true
			if certAlias == "" {
				return nil, fmt.Errorf("cert alias empty while port [%d] for SSL is specified", port.Port)
			}
			portDetailsList[idx].CertAlias = certAlias
		}
	}
	klog.Infof("Creating loadbalancer for ports [%#v]\n", portDetailsList)
	// Create using VCD API
	resourcesAllocated := &util.AllocatedResourcesMap{}
	lbIP, err := gm.CreateLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, nodeIPs, portDetailsList,
		lb.OneArm, lb.EnableVirtualServiceSharedIP, portNameToIPMap, "", resourcesAllocated)
	if rdeErr := lb.addLBResourcesToRDE(ctx, resourcesAllocated, lbIP); rdeErr != nil {
		return nil, fmt.Errorf("unable to add load balancer pool resources to RDE [%s]: [%v]", lb.clusterID, err)
	}
	if err != nil {
		addToErrorSetErr := cpiRdeManager.AddToErrorSetWithNameAndId(ctx, cpisdk.CreateLoadbalancerError, "", virtualServiceNamePrefix, err.Error())
		if addToErrorSetErr != nil {
			klog.Errorf("error adding CPI error [%s] to RDE: [%s], [%v]", cpisdk.CreateLoadbalancerError, lb.clusterID, addToErrorSetErr)
		}
		return nil, fmt.Errorf("unable to create loadbalancer for ports [%#v]: [%v]", portDetailsList, err)
	}

	if lbIP != "" {
		err = cpiRdeManager.AddVirtualIpToRDE(ctx, lbIP)
		if err != nil {
			addToErrorSetErr := cpiRdeManager.AddToErrorSet(ctx, cpisdk.AddVIPToRdeError, lb.clusterID, err.Error())
			if addToErrorSetErr != nil {
				klog.Errorf("error adding CPI error [%s] to RDE: [%s], [%v]", cpisdk.AddVIPToRdeError, lb.clusterID, addToErrorSetErr)
			}
			klog.Errorf("error when adding virtual IP to RDE: [%v]", err)
		}
	}

	err = cpiRdeManager.AddToEventSetWithNameAndId(ctx, cpisdk.CreatedLoadbalancer, "", virtualServiceNamePrefix, fmt.Sprintf("Created loadbalancer successfully for [%s] with external IP: [%s]", lb.clusterID, lbIP))
	if err != nil {
		klog.Errorf("error adding CPI event [%s] to RDE: [%v]", cpisdk.CreatedLoadbalancer, err)
	}

	// No errors and external IP exists, we can remove the create LB error; empty vcdResourceName as we only have access to the external IP
	err = cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, cpisdk.CreateLoadbalancerError, "", virtualServiceNamePrefix)
	if err != nil {
		klog.Errorf("there was an error removing CPI error [%s] from RDE [%s], [%v]", cpisdk.CreateLoadbalancerError, lb.clusterID, err)
	}

	klog.Infof("Created load balancer with external IP [%s], ports [%#v]\n", lbIP, portDetailsList)

	return &v1.LoadBalancerStatus{
		Ingress: []v1.LoadBalancerIngress{
			{
				IP: lbIP,
			},
		},
	}, nil
}
