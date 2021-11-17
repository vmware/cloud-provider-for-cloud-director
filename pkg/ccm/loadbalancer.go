/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

// +build !testing

package ccm

import (
	"context"
	"fmt"
	"strings"

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdclient"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	cloudProvider "k8s.io/cloud-provider"
	"k8s.io/klog"
)

//LBManager -
type LBManager struct {
	vcdClient  *vcdclient.Client
	kubeClient *kubernetes.Clientset
	namespace  string
}

func newLoadBalancer(vcdClient *vcdclient.Client) cloudProvider.LoadBalancer {
	return &LBManager{
		vcdClient:  vcdClient,
		kubeClient: GetK8SClient(),
		namespace:  "default",
	}
}

func (lb *LBManager) getNodeIPs() ([]string, error) {
	nodes, err := lb.kubeClient.CoreV1().Nodes().List(metav1.ListOptions{})
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
	nodeIPs, err := lb.getNodeIPs()
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

func (lb *LBManager) getServicePortMap(service *v1.Service) map[string]int32 {
	typeToInternalPort := make(map[string]int32)
	for _, port := range service.Spec.Ports {
		typeToInternalPort[strings.ToLower(port.Name)] = port.NodePort
	}
	return typeToInternalPort
}

func (lb *LBManager) getLBPoolNamePrefix(serviceName string, clusterID string) string {
	return fmt.Sprintf("ingress-pool-%s-%s", serviceName, clusterID)
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
	lbPoolNamePrefix := lb.getLBPoolNamePrefix(service.Name, lb.vcdClient.ClusterID)
	nodeIps := lb.getNodeInternalIps(nodes)
	klog.Infof("UpdateLoadBalancer Node Ips: %v", nodeIps)
	typeToInternalPortMap := lb.getServicePortMap(service)
	for portType, internalPort := range typeToInternalPortMap {
		if portType != "http" && portType != "https" {
			klog.Infof("Encountered unhandled port type: [%s]", portType)
			continue
		}
		lb.vcdClient.UpdateLoadBalancer(ctx, lbPoolNamePrefix+portType, nodeIps, internalPort)
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
	service *v1.Service) (status *v1.LoadBalancerStatus, exists bool, err error) {

	virtualServiceNamePrefix := lb.getLoadBalancerPrefix(ctx, service)
	virtualIP := ""
	for _, port := range service.Spec.Ports {
		virtualServiceName := fmt.Sprintf("%s-%s", virtualServiceNamePrefix, port.Name)
		virtualIP, err = lb.vcdClient.GetLoadBalancer(ctx, virtualServiceName)
		if err != nil {
			return nil, false,
				fmt.Errorf("unable to get virtual service summary for [%s]: [%v]",
					virtualServiceName, err)
		}
		if virtualIP == "" {
			// if any lb that is expected is not created, return false to retry creation
			return nil, false, nil
		}
	}
	if virtualIP == "" {
		// this implies that no port was specified
		return nil, false, nil
	}

	return &v1.LoadBalancerStatus{
		Ingress: []v1.LoadBalancerIngress{
			{
				IP: virtualIP,
			},
		},
	}, true, nil

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
	return lb.getLoadBalancer(ctx, service)
}

func (lb *LBManager) getLoadBalancerPrefix(_ context.Context, service *v1.Service) string {
	return fmt.Sprintf("ingress-vs-%s-%s", service.Name, lb.vcdClient.ClusterID)
}

// GetLoadBalancerName returns the name of the load balancer. Implementations must treat the
// *v1.Service parameter as read-only and not modify it.
func (lb *LBManager) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
	return lb.getLoadBalancerPrefix(ctx, service)
}

func (lb *LBManager) deleteLoadBalancer(ctx context.Context, service *v1.Service) error {

	virtualServiceName := fmt.Sprintf("ingress-vs-%s-%s", service.Name, lb.vcdClient.ClusterID)
	lbPoolNamePrefix := lb.getLBPoolNamePrefix(service.Name, lb.vcdClient.ClusterID)
	klog.Infof("Deleting virtual service [%s] and lb pool [%s]", virtualServiceName, lbPoolNamePrefix)

	err := lb.vcdClient.DeleteLoadBalancer(ctx, virtualServiceName, lbPoolNamePrefix)
	if err != nil {
		return fmt.Errorf("Unable to delete load balancer for virtual-service [%s] and lb pool [%s]: [%v]",
			virtualServiceName, lbPoolNamePrefix, err)
	}

	return nil
}

func (lb *LBManager) createLoadBalancer(ctx context.Context, service *v1.Service,
	nodeIPs []string) (*v1.LoadBalancerStatus, error) {

	lbStatus, lbExists, err := lb.getLoadBalancer(ctx, service)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while querying for loadbalancer: [%v]", err)
	}
	if lbExists {
		// Return existing LB. No need to update here as separate functions are called for Update.
		return lbStatus, nil
	}

	virtualServiceName := fmt.Sprintf("ingress-vs-%s-%s", service.Name, lb.vcdClient.ClusterID)
	lbPoolNamePrefix := lb.getLBPoolNamePrefix(service.Name, lb.vcdClient.ClusterID)

	// While creating the lb, even if only one of http/https is remaining and the other is completed,
	// ask for both to be created. The already created one will silently pass.
	httpPortSuffix := ""
	httpPort := int32(0)
	httpNodePort := int32(0)

	httpsPortSuffix := ""
	httpsPort := int32(0)
	httpsNodePort := int32(0)
	for _, port := range service.Spec.Ports {
		switch port.Name {
		case `http`:
			httpPortSuffix = port.Name
			httpPort = port.Port
			httpNodePort = port.NodePort

		case `https`:
			httpsPortSuffix = port.Name
			httpsPort = port.Port
			httpsNodePort = port.NodePort

		default:
			klog.Infof("Encountered unhandled port [%#v]\n", port)
		}
	}
	klog.Infof("Creating loadbalancer for ports http:[%d=>%d], https: [%d=>%d]\n",
		httpPort, httpNodePort, httpsPort, httpsNodePort)

	// Create using VCD API
	lbIP, err := lb.vcdClient.CreateLoadBalancer(ctx, virtualServiceName, lbPoolNamePrefix, nodeIPs,
		httpPortSuffix, httpPort, httpNodePort, httpsPortSuffix, httpsPort, httpsNodePort)
	if err != nil {
		return nil, fmt.Errorf("unable to create loadbalancer for http:[%d=>%d], https: [%d=>%d]: [%v]",
			httpPort, httpNodePort, httpsPort, httpsNodePort, err)
	}
	klog.Infof("Created loadbalancer with external IP [%s], ports [%d=>%d], [%d=>%d]\n",
		lbIP, httpPort, httpNodePort, httpsPort, httpsNodePort)

	return &v1.LoadBalancerStatus{
		Ingress: []v1.LoadBalancerIngress{
			{
				IP: lbIP,
			},
		},
	}, nil
}
