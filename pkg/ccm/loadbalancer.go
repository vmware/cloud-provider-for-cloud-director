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
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	cloudProvider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"strconv"
	"strings"
)

const (
	sslPortsAnnotation              = `service.beta.kubernetes.io/vcloud-avi-ssl-ports`
	sslCertAliasAnnotation          = `service.beta.kubernetes.io/vcloud-avi-ssl-cert-alias`
	skipAviSSLTerminationAnnotation = `service.beta.kubernetes.io/vcloud-avi-ssl-no-termination`
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
	ovdcName                     string
	ipamSubnet                   string
	clusterID                    string
	EnableVirtualServiceSharedIP bool
}

func newLoadBalancer(vcdClient *vcdsdk.Client, certAlias string, oneArm *vcdsdk.OneArm,
	ovdcNetworkName string, ovdcName string, ipamSubnet string, clusterID string, enableVirtualServiceSharedIP bool) cloudProvider.LoadBalancer {

	return &LBManager{
		vcdClient:                    vcdClient,
		kubeClient:                   GetK8SClient(),
		namespace:                    "default",
		CertificateAlias:             certAlias,
		OneArm:                       oneArm,
		ovdcNetworkName:              ovdcNetworkName,
		ovdcName:                     ovdcName,
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
		// If we can find the worker node from the missing controlPlaneLabel in nodeLabelMap we will take it,
		// but if the node is missing labels, we will just let it go instead of adding it.
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

	// fetch the user specified IP address for the load balancer
	// NOTE: userSpecifiedLBIP cannot be nil as it is a string.
	// if userSpecifiedLBIP is empty, the empty string is passed down to CreateLoadBalancer() which uses an external IP from IP gateway allocations.
	userSpecifiedLBIP := getUserSpecifiedLoadBalancerIP(service)
	klog.Infof("UpdateLoadBalancer called with loadBalancerIP [%s] for service [%s]", userSpecifiedLBIP, service.Name)

	for portName, internalPort := range typeToInternalPortMap {
		lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portName)
		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portName)
		externalPort := typeToExternalPort[portName]
		gm, err := vcdsdk.NewGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet, lb.ovdcName)
		if err != nil {
			return fmt.Errorf("error while creating GatewayManager: [%v]", err)
		}
		klog.Infof("Updating pool [%s] with port [%s:%d]", lbPoolName, portName, internalPort)
		protocol, _ := nameToProtocol[portName]
		resourcesAllocated := &util.AllocatedResourcesMap{}
		vip, err := gm.UpdateLoadBalancer(ctx, lbPoolName, virtualServiceName, nodeIps, userSpecifiedLBIP, internalPort,
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
	gm, err := vcdsdk.NewGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet, lb.ovdcName)
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
				fmt.Errorf("unable to get load balancer information for the service [%s]: [%v]",
					virtualServiceName, err)
		}
		removeErr := cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(ctx, vcdsdk.ComponentCPI, cpisdk.GetLoadbalancerError, "", virtualServiceName)
		if removeErr != nil {
			klog.Errorf("there was an error removing CPI error [%s] from RDE [%s], [%v]", cpisdk.GetLoadbalancerError, lb.clusterID, err)
		}
		portNameToIP[port.Name] = virtualIP
		if virtualIP != "" {
			// There is no check to ensure there is only one virtual IP
			// If a load balancer has its lb IP updated, one vs may have a
			// different IP than the other vs.
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
			// If we have an empty IP for a specific port, there may be resources allocated in other ports so we will do resource check to return for delete.
			// As if there is vcd resources, we should return a nil status, but true for lb exists to the controller to pick up for deletion.

			// In the event we do get an actual vcdResourceCheck error, then hasVcdResource (lbExists) may be false. The controller will return and reconcile
			// as there is an error check first. If there is an error, the controller will return the delete operation immediately and will not continue to indicate lb has been deleted.
			// Only if there is no error then it will check if hasVcdResource (lbExists), and if it exists, then it will ensure the deletion.
			hasVcdResources, vcdResourceCheckErr := lb.VerifyVCDResourcesForApplicationLB(ctx, service)
			if vcdResourceCheckErr != nil {
				klog.Errorf("error occurred while checking vcd resource for application LB in GetLoadBalancer: [%v]", vcdResourceCheckErr)
			}
			return nil, hasVcdResources, vcdResourceCheckErr
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

	gm, err := vcdsdk.NewGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet, lb.ovdcName)
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

func shouldSkipAviSSLTermination(service *v1.Service) bool {
	shouldSkipAviSSLTerminationStr, ok := service.Annotations[skipAviSSLTerminationAnnotation]
	if !ok {
		return false
	}

	return strings.ToLower(shouldSkipAviSSLTerminationStr) == "true"
}

// getUserSpecifiedLoadBalancerIP returns the specified load balancer IP
func getUserSpecifiedLoadBalancerIP(service *v1.Service) string {
	return service.Spec.LoadBalancerIP
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
	gm, err := vcdsdk.NewGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet, lb.ovdcName)
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

	// fetch the user specified IP address for the load balancer
	// NOTE: userSpecifiedLBIP cannot be nil as it is a string.
	// if userSpecifiedLBIP is empty, the empty string is passed down to CreateLoadBalancer() which uses an external IP from IP gateway allocations.
	userSpecifiedLBIP := getUserSpecifiedLoadBalancerIP(service)
	klog.Infof("createLoadBalancer called with loadBalancerIP [%s] for service [%s]", userSpecifiedLBIP, service.Name)

	if lbExists {
		// Update load balancer if there are changes in service properties
		typeToInternalPortMap, typeToExternalPortMap, nameToProtocol := lb.getServicePortMap(service)
		klog.Infof("test1234: typeToExternalPortMap: [%v], %d", typeToExternalPortMap, len(typeToExternalPortMap))
		for portName, internalPort := range typeToInternalPortMap {
			lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portName)
			virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portName)
			klog.Infof("test1234: processing vs: [%s]", virtualServiceName)
			externalPort := typeToExternalPortMap[portName]
			protocol, _ := nameToProtocol[portName]
			klog.Infof("Updating pool [%s] with port [%s:%d:%d]", lbPoolName, portName, internalPort, externalPort)
			resourcesAllocated := &util.AllocatedResourcesMap{}
			vip, err := gm.UpdateLoadBalancer(ctx, lbPoolName, virtualServiceName, nodeIPs, userSpecifiedLBIP, internalPort,
				externalPort, lb.OneArm, lb.EnableVirtualServiceSharedIP, protocol, resourcesAllocated)
			klog.Infof("test1234: UpdateLB returned with vip: [%s]", vip)
			if rdeErr := lb.addLBResourcesToRDE(ctx, resourcesAllocated, vip); rdeErr != nil {
				klog.Infof("test1234: err adding lb resource to rde: [%v]", err)
				return nil, fmt.Errorf("failed to update RDE [%s] with load balancer resources: [%v]", lb.clusterID, err)
			}

			vsSummary, getVsErr := gm.GetVirtualService(ctx, virtualServiceName)
			if getVsErr != nil {
				klog.Infof("test1234: failed to get vs [%s]", virtualServiceName)
				return nil, fmt.Errorf("failed to get virtual service [%s], [%v]", virtualServiceName, getVsErr)
			}
			if vsSummary == nil {
				klog.Infof("test1234: empty vssummary for vs: [%s]", virtualServiceName)
				return nil, fmt.Errorf("virtual service [%s] does not exist", virtualServiceName)
			}

			if err != nil {
				klog.Infof("test1234: error is not nil: [%v]", err)
				addToErrorSetErr := cpiRdeManager.AddToErrorSetWithNameAndId(ctx, cpisdk.UpdateLoadbalancerError, vsSummary.Id, vsSummary.Name, err.Error())
				if addToErrorSetErr != nil {
					klog.Errorf("error adding CPI error [%s] to RDE: [%s], [%v]", cpisdk.UpdateLoadbalancerError, lb.clusterID, addToErrorSetErr)
				}
				return nil, fmt.Errorf("unable to update load balancer [%s] with port [%s:%d:%d] and load balancer IP [%s]: [%v]", lbPoolName, portName,
					internalPort, externalPort, userSpecifiedLBIP, err)
			}

			if userSpecifiedLBIP != "" && vip != userSpecifiedLBIP {
				klog.Infof("test1234: user specified lb ip empty or not equap to vip [%s]: [%v]", vip, err)
				return nil, fmt.Errorf("failed to update loadbalancerIP to [%s] for the service [%s]: expected the load balancer IP to be [%s] but got [%s]",
					userSpecifiedLBIP, service.Name, userSpecifiedLBIP, vip)
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
		// recreate the load balancer status as the load balancer properties may be updated
		lbStatus, _, err = lb.getLoadBalancer(ctx, service)
		if err != nil {
			klog.Infof("test1234: error getting lb: [%v]", err)
			addToErrorSetErr := cpiRdeManager.AddToErrorSetWithNameAndId(ctx, cpisdk.GetLoadbalancerError, "", virtualServiceNamePrefix, err.Error())
			if addToErrorSetErr != nil {
				klog.Errorf("error adding CPI error [%s] to the RDE [%s], [%v]", cpisdk.GetLoadbalancerError, lb.clusterID, addToErrorSetErr)
			}
			return nil, fmt.Errorf("unexpected error while querying for loadbalancer after updating load balancer: [%v]", err)
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

	// allow users to terminate SSL by other means
	skipAviSSLTermination := shouldSkipAviSSLTermination(service)
	klog.Infof("Annotation [%s] set to [%v]", skipAviSSLTerminationAnnotation, skipAviSSLTermination)
	if skipAviSSLTermination {
		certAlias = ""
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
		if port.AppProtocol != nil && *port.AppProtocol != "" && !skipAviSSLTermination {
			switch strings.ToUpper(*port.AppProtocol) {
			// allow override in case of known protocols such as HTTP/HTTPS/TCP which are directly supported in Avi
			case "HTTP", "HTTPS", "TCP":
				portDetailsList[idx].Protocol = strings.ToUpper(*port.AppProtocol)
			}
		}
		if _, ok := portsMap[port.Port]; ok {
			if !skipAviSSLTermination {
				portDetailsList[idx].UseSSL = true
				if certAlias == "" {
					return nil, fmt.Errorf("cert alias empty while port [%d] for SSL is specified", port.Port)
				}
				portDetailsList[idx].CertAlias = certAlias
			}
		}
	}
	klog.Infof("Creating loadbalancer for ports [%#v]\n", portDetailsList)
	// Create using VCD API
	resourcesAllocated := &util.AllocatedResourcesMap{}
	lbIP, err := gm.CreateLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, nodeIPs, portDetailsList,
		lb.OneArm, lb.EnableVirtualServiceSharedIP, portNameToIPMap, userSpecifiedLBIP, resourcesAllocated)
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

func (lb *LBManager) VerifyVCDResourcesForApplicationLB(ctx context.Context, service *v1.Service) (bool, error) {
	virtualServiceNamePrefix := lb.getVirtualServicePrefix(ctx, service)
	lbPoolNamePrefix := lb.getLBPoolNamePrefix(ctx, service)
	klog.Infof("Checking VCD Resources with virtual service [%s] and lb pool [%s]", virtualServiceNamePrefix, lbPoolNamePrefix)

	portDetailsList := make([]vcdsdk.PortDetails, len(service.Spec.Ports))
	for idx, port := range service.Spec.Ports {
		portDetailsList[idx] = vcdsdk.PortDetails{
			PortSuffix:   port.Name,
			ExternalPort: port.Port,
			InternalPort: port.NodePort,
			Protocol:     string(port.Protocol),
		}
	}
	return lb.verifyVCDResourcesForApplicationLB(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, lb.OneArm)
}

/**
In GetVirtualService(), we will always expect 1 virtual service back only. This is due to virtual service names
being unique as there cannot have two of the same virtual service names, and in GetVirtualService() we have a FIQL name==%s filter
to search for a virtual service of %s name.
*/
func (lb *LBManager) verifyVCDResourcesForApplicationLB(ctx context.Context, virtualServiceNamePrefix string,
	lbPoolNamePrefix string, portDetailsList []vcdsdk.PortDetails, oneArm *vcdsdk.OneArm) (bool, error) {

	gatewayMgr, err := vcdsdk.NewGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet, lb.ovdcName)
	if err != nil {
		return false, fmt.Errorf("error creating new gateway manager [%v]", err)
	}

	for _, portDetails := range portDetailsList {
		if portDetails.InternalPort == 0 {
			klog.Infof("No internal port specified for [%s], hence loadbalancer not created", portDetails.PortSuffix)
			continue
		}
		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portDetails.PortSuffix)
		lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portDetails.PortSuffix)

		// GetVirtualService() will always look for 1 virtual service named virtualServiceName, as VCD doesn't allow duplicate VS names.
		// In the event where a virtual service name is not found in GetVirtualService(), it will return nil for both error and vsSummary.
		vsSummary, err := gatewayMgr.GetVirtualService(ctx, virtualServiceName)
		if err != nil {
			return false, fmt.Errorf("error getting virtual service [%s]: [%v] during VCD resources verification", virtualServiceName, err)
		}

		if vsSummary != nil {
			klog.Infof("Virtual Service found: [%s]", virtualServiceName)
			return true, nil
		}

		// GetLoadBalancerPool() will return nil, error if lbPool is not found.
		lbPool, err := gatewayMgr.GetLoadBalancerPool(ctx, lbPoolName)
		if err != nil && err != govcd.ErrorEntityNotFound {
			return false, fmt.Errorf("error getting loadbalancer pool for [%s]: [%v] during VCD resources verification", lbPoolName, err)
		}
		if lbPool != nil {
			klog.Infof("Load balancer pool found: [%s]", lbPoolName)
			return true, nil
		}

		if oneArm != nil {
			dnatRuleName := vcdsdk.GetDNATRuleName(virtualServiceName)
			dnatRuleRef, err := gatewayMgr.GetNATRuleRef(ctx, dnatRuleName)
			if err != nil {
				return false, fmt.Errorf("error getting dnat rule ref for [%s]: [%v] during VCD resources verification", dnatRuleName, err)
			}
			if dnatRuleRef != nil {
				klog.Infof("DNAT Rule Ref found: [%s]", dnatRuleName)
				return true, nil
			}
			appPortProfileName := vcdsdk.GetAppPortProfileName(dnatRuleName)
			org, err := lb.vcdClient.VCDClient.GetOrgByName(lb.vcdClient.ClusterOrgName)
			if err != nil {
				return false, fmt.Errorf("unable to find org [%s] by name: [%v]", lb.vcdClient.ClusterOrgName, err)
			}
			appPortProfile, err := org.GetNsxtAppPortProfileByName(appPortProfileName, types.ApplicationPortProfileScopeTenant)
			// We are doing a string check here because it returns a formatted error instead of govcd.ErrorEntityNotFound
			if err != nil && !strings.Contains(err.Error(), "[ENF] entity not found") {
				return false, fmt.Errorf("unable to get app port profile [%s]: [%v]", appPortProfileName, err)
			}
			if appPortProfile != nil {
				return true, nil
			}
		}
	}
	return false, nil
}
