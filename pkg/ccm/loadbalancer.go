/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

// +build !testing

package ccm

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/cpisdk"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
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
)

//LBManager -
type LBManager struct {
	gatewayManager   *vcdsdk.GatewayManager
	vcdClient        *vcdsdk.Client
	kubeClient       *kubernetes.Clientset
	namespace        string
	CertificateAlias string
	OneArm           *vcdsdk.OneArm
	ovdcNetworkName  string
	ipamSubnet       string
	clusterID        string
	UseVsSharedIP	 bool
}

func newLoadBalancer(vcdClient *vcdsdk.Client, certAlias string, oneArm *vcdsdk.OneArm,
	ovdcNetworkName string, ipamSubnet string, clusterID string, useVsSharedIP bool) cloudProvider.LoadBalancer {

	return &LBManager{
		vcdClient:        vcdClient,
		kubeClient:       GetK8SClient(),
		namespace:        "default",
		CertificateAlias: certAlias,
		OneArm:           oneArm,
		ovdcNetworkName:  ovdcNetworkName,
		ipamSubnet:       ipamSubnet,
		clusterID:        clusterID,
		UseVsSharedIP: 	  useVsSharedIP,
	}
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
	nodeIPs, err := lb.getNodeIPs(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get nodes in cluster: [%v]", err)
	}
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

func (lb *LBManager) getServicePortMap(service *v1.Service) (map[string]int32, map[string]int32) {
	typeToInternalPort := make(map[string]int32)
	typeToExternalPort := make(map[string]int32)
	for _, port := range service.Spec.Ports {
		typeToInternalPort[strings.ToLower(port.Name)] = port.NodePort
		typeToExternalPort[strings.ToLower(port.Name)] = port.Port
	}
	return typeToInternalPort, typeToExternalPort
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

	nodeIps := lb.getNodeInternalIps(nodes)
	klog.Infof("UpdateLoadBalancer Node Ips: %v", nodeIps)

	lbPoolNamePrefix := lb.getLBPoolNamePrefix(ctx, service)
	virtualServiceNamePrefix := lb.getVirtualServicePrefix(ctx, service)
	typeToInternalPortMap, typeToExternalPort := lb.getServicePortMap(service)
	for portName, internalPort := range typeToInternalPortMap {
		lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portName)
		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portName)
		externalPort := typeToExternalPort[portName]
		cgm, err := cpisdk.NewCpiGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet, lb.clusterID)
		if err != nil {
			return fmt.Errorf("error while creating GatewayManager: [%v]", err)
		}
		klog.Infof("Updating pool [%s] with port [%s:%d]", lbPoolName, portName, internalPort)
		if err := cgm.UpdateLoadBalancer(ctx, lbPoolName, virtualServiceName, nodeIps, internalPort,
			externalPort, lb.OneArm, lb.UseVsSharedIP); err != nil {
			return fmt.Errorf("unable to update pool [%s] with port [%s:%d]: [%v]", lbPoolName, portName,
				internalPort, err)
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
	cgm, err := cpisdk.NewCpiGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet, lb.clusterID)
	if err != nil {
		return nil, nil, fmt.Errorf("error while creating GatewayManager: [%v]", err)
	}
	portNameToIP := make(map[string]string)
	ingressVirtualIP := ""
	for _, port := range service.Spec.Ports {
		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, port.Name)
		virtualIP, err = cgm.GetLoadBalancer(ctx, virtualServiceName, lb.OneArm)
		if err != nil {
			return nil, nil,
				fmt.Errorf("unable to get virtual service summary for [%s]: [%v]",
					virtualServiceName, err)
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

	cgm, err := cpisdk.NewCpiGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet, lb.clusterID)
	if err != nil {
		return fmt.Errorf("error while creating CpiGatewayManager: [%v]", err)
	}
	err = cgm.DeleteLoadBalancer(ctx, virtualServiceName, lbPoolNamePrefix, portDetailsList, lb.OneArm)
	if err != nil {
		return fmt.Errorf("Unable to delete load balancer for virtual-service [%s] and lb pool [%s]: [%v]",
			virtualServiceName, lbPoolNamePrefix, err)
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
	if err != nil {
		return nil, fmt.Errorf("unexpected error while querying for loadbalancer: [%v]", err)
	}
	cgm, err := cpisdk.NewCpiGatewayManager(ctx, lb.vcdClient, lb.ovdcNetworkName, lb.ipamSubnet, lb.clusterID)
	if err != nil {
		return nil, fmt.Errorf("error while creating CpiGatewayManager: [%v]", err)
	}

	// golang doesn't have the set data structure
	lbExists := true
	for _, ip := range portNameToIPMap {
		if ip == "" {
			lbExists = false
		}
	}

	if lbExists {
		// Update load balancer if there are changes in service properties
		typeToInternalPortMap, typeToExternalPortMap := lb.getServicePortMap(service)
		for portName, internalPort := range typeToInternalPortMap {
			lbPoolName := fmt.Sprintf("%s-%s", lbPoolNamePrefix, portName)
			virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, portName)
			externalPort := typeToExternalPortMap[portName]
			klog.Infof("Updating pool [%s] with port [%s:%d:%d]", lbPoolName, portName, internalPort, externalPort)
			if err := cgm.UpdateLoadBalancer(ctx, lbPoolName, virtualServiceName, nodeIPs, internalPort,
				externalPort, lb.OneArm, lb.UseVsSharedIP); err != nil {
				return nil, fmt.Errorf("unable to update pool [%s] with port [%s:%d:%d]: [%v]", lbPoolName, portName,
					internalPort, externalPort, err)
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
	lbIP, err := cgm.CreateLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, nodeIPs, portDetailsList,
		lb.OneArm, lb.UseVsSharedIP, portNameToIPMap)
	if err != nil {
		return nil, fmt.Errorf("unable to create loadbalancer for ports [%#v]: [%v]", portDetailsList, err)
	}
	klog.Infof("Created loadbalancer with external IP [%s], ports [%#v]\n", lbIP, portDetailsList)

	return &v1.LoadBalancerStatus{
		Ingress: []v1.LoadBalancerIngress{
			{
				IP: lbIP,
			},
		},
	}, nil
}
