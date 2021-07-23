/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

// +build !testing

package ccm

import (
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/config"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdclient"
	"io"
	"k8s.io/client-go/informers"
	_ "k8s.io/client-go/tools/clientcmd"
	cloudProvider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"time"
)

// ProviderName : name of the cloud provider
const (
	ProviderName = "vmware-cloud-director"
)

// VCDCloudProvider - contains all of the interfaces for our cloud provider
type VCDCloudProvider struct {
	vcdClient *vcdclient.Client
	lb        cloudProvider.LoadBalancer
	instances cloudProvider.Instances
}

var _ cloudProvider.Interface = &VCDCloudProvider{}

func init() {
	cloudProvider.RegisterCloudProvider(ProviderName, newVCDCloudProvider)
}

func newVCDCloudProvider(configReader io.Reader) (cloudProvider.Interface, error) {

	cloudConfig, err := config.ParseCloudConfig(configReader)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: [%v]", err)
	}

	var oneArm *vcdclient.OneArm = nil
	if cloudConfig.LB.OneArm != nil {
		oneArm = &vcdclient.OneArm{
			StartIPAddress: cloudConfig.LB.OneArm.StartIP,
			EndIPAddress:   cloudConfig.LB.OneArm.EndIP,
		}
	}
	var vcdClient *vcdclient.Client
	for {
		vcdClient, err = vcdclient.NewVCDClientFromSecrets(
			cloudConfig.VCD.Host,
			cloudConfig.VCD.Org,
			cloudConfig.VCD.VDC,
			cloudConfig.VCD.VDCNetwork,
			cloudConfig.VCD.VIPSubnet,
			cloudConfig.VCD.User,
			cloudConfig.VCD.Secret,
			true,
			cloudConfig.ClusterID,
			oneArm,
			cloudConfig.LB.Ports.HTTP,
			cloudConfig.LB.Ports.HTTPS,
			cloudConfig.LB.CertificateAlias,
			true,
		)
		if err == nil {
			break
		} else {
			klog.Infof("Error initializing client from secrets: [%v]", err)
			time.Sleep(10 * time.Second)
		}
	}

	// setup LB only if the gateway is not NSX-T
	var lb cloudProvider.LoadBalancer = nil
	if !vcdClient.IsNSXTBackedGateway() {
		klog.Infof("Gateway of network [%s] not backed by NSX-T. Hence LB will not be initialized.",
			cloudConfig.VCD.VDCNetwork)
	} else {
		lb = newLoadBalancer(vcdClient)
	}

	// cache for VM Info with an refresh of elements needed after 1 minute
	vmInfoCache := newVmInfoCache(vcdClient, time.Minute)

	return &VCDCloudProvider{
		vcdClient: vcdClient,
		lb:        lb,
		instances: newInstances(vmInfoCache),
	}, nil
}

// Initialize - starts the cloud-provider controller
func (vcdCP *VCDCloudProvider) Initialize(clientBuilder cloudProvider.ControllerClientBuilder, stop <-chan struct{}) {
	clientSet := clientBuilder.ClientOrDie("do-shared-informers")
	sharedInformer := informers.NewSharedInformerFactory(clientSet, 0)

	sharedInformer.Start(nil)
	sharedInformer.WaitForCacheSync(nil)

	return
}

// ProviderName returns the cloud provider ID.
func (vcdCP *VCDCloudProvider) ProviderName() string {
	return ProviderName
}

// LoadBalancer returns a loadbalancer interface. Also returns true if the interface is supported, false otherwise.
func (vcdCP *VCDCloudProvider) LoadBalancer() (cloudProvider.LoadBalancer, bool) {
	return vcdCP.lb, true
}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (vcdCP *VCDCloudProvider) Instances() (cloudProvider.Instances, bool) {
	return vcdCP.instances, true
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
func (vcdCP *VCDCloudProvider) Zones() (cloudProvider.Zones, bool) {
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (vcdCP *VCDCloudProvider) Clusters() (cloudProvider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface is supported.
func (vcdCP *VCDCloudProvider) Routes() (cloudProvider.Routes, bool) {
	return nil, false
}

// HasClusterID provides an opportunity for cloud-provider-specific code to process DNS settings for pods.
func (vcdCP *VCDCloudProvider) HasClusterID() bool {
	return false
}
