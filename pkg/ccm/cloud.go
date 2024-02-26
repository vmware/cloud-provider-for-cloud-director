//go:build !testing
// +build !testing

/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package ccm

import (
	"context"
	"errors"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/config"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/cpisdk"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/cloud-provider-for-cloud-director/release"
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
	vcdClient   *vcdsdk.Client
	lb          cloudProvider.LoadBalancer
	instances   cloudProvider.Instances
	instancesV2 cloudProvider.InstancesV2
	zoneMap     *vcdsdk.ZoneMap
}

var _ cloudProvider.Interface = &VCDCloudProvider{}

func init() {
	cloudProvider.RegisterCloudProvider(ProviderName, newVCDCloudProvider)
}

func newVCDCloudProvider(configReader io.Reader) (cloudProvider.Interface, error) {
	var vcdClient *vcdsdk.Client = nil
	var cloudConfig *config.CloudConfig = nil
	cloudConfig, err := config.ParseCloudConfig(configReader)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: [%v]", err)
	}
	for {
		err = config.SetAuthorization(cloudConfig)
		if err != nil {
			klog.Infof("Unable to set authorization in config: [%v]", err)
			time.Sleep(10 * time.Second)
			continue
		}

		err = config.ValidateCloudConfig(cloudConfig)
		if err != nil {
			klog.Infof("error validating config: [%v]", err)
			time.Sleep(10 * time.Second)
			continue
		}

		vcdClient, err = vcdsdk.NewVCDClientFromSecrets(
			cloudConfig.VCD.Host,
			cloudConfig.VCD.Org,
			cloudConfig.VCD.VDC,
			cloudConfig.VCD.UserOrg,
			cloudConfig.VCD.User,
			cloudConfig.VCD.Secret,
			cloudConfig.VCD.RefreshToken,
			true,
			true,
		)
		if err == nil {
			break
		}

		klog.Infof("Error initializing client from secrets: [%v]", err)
		time.Sleep(10 * time.Second)
	}

	rdeManager := vcdsdk.NewRDEManager(vcdClient, cloudConfig.ClusterID, release.CloudControllerManagerName, release.Version)
	cpiRdeManager := cpisdk.NewCPIRDEManager(rdeManager)

	err = cpiRdeManager.AddToEventSet(context.Background(), cpisdk.ClientAuthenticated, cloudConfig.ClusterID, "successfully authenticated into vcdclient from secrets")
	if err != nil {
		klog.Errorf("error adding CPI event [%s] to RDE: [%v]", cpisdk.ClientAuthenticated, err)
	}

	// setup LB only if the gateway is not NSX-T
	var lb cloudProvider.LoadBalancer = nil
	gm, err := vcdsdk.NewGatewayManager(context.Background(), vcdClient, cloudConfig.LB.VDCNetwork, cloudConfig.LB.VIPSubnet, cloudConfig.VCD.VDC)
	if err != nil {
		return nil, fmt.Errorf("failed to create GatewayManager: [%v]", err)
	}
	if !gm.IsNSXTBackedGateway() {
		klog.Infof("Gateway of network [%s] not backed by NSX-T. Hence LB will not be initialized.",
			cloudConfig.LB.VDCNetwork)
	} else {
		var oneArm *vcdsdk.OneArm
		if cloudConfig.LB.OneArm != nil {
			oneArm = &vcdsdk.OneArm{
				StartIP: cloudConfig.LB.OneArm.StartIP,
				EndIP:   cloudConfig.LB.OneArm.EndIP,
			}
		}
		lb = newLoadBalancer(vcdClient, cloudConfig.LB.CertificateAlias, oneArm, cloudConfig.LB.VDCNetwork, cloudConfig.VCD.VDC,
			cloudConfig.LB.VIPSubnet, cloudConfig.ClusterID, cloudConfig.LB.EnableVirtualServiceSharedIP)
	}

	// TODO: upgrade all CAPVCD RDEs here

	err = cpiRdeManager.UpgradeCPIStatusOfExistingRDE(context.Background(), cloudConfig.ClusterID)
	if err != nil {
		klog.Errorf("failed to create CPI status in the RDE [%s]: [%v]", cloudConfig.ClusterID, err)
		err = cpiRdeManager.AddToErrorSet(context.Background(), cpisdk.CPIStatusUpgradeRdeError, cloudConfig.ClusterID, err.Error())
		if err != nil {
			if _, ok := err.(vcdsdk.NonCAPVCDEntityError); !ok {
				msg := fmt.Sprintf("failed to add CPI error [%s] to ErrorSet in RDE [%s], [%v]", cpisdk.CPIStatusUpgradeRdeError, cloudConfig.ClusterID, err)
				klog.Error(msg)
				return nil, errors.New(msg)
			}
		}
	} else {
		klog.Infof("successfully created CPI status in RDE [%s]", cloudConfig.ClusterID)
		err = cpiRdeManager.AddToEventSet(context.Background(), cpisdk.CPIStatusRDEUpgraded, cloudConfig.ClusterID, fmt.Sprintf("CPI status section successfully upgraded from RDE [%s]", cloudConfig.ClusterID))
		if err != nil {
			klog.Errorf("failed to add CPI event [%s] to EventSet in RDE [%s], [%v]", cpisdk.CPIStatusRDEUpgraded, cloudConfig.ClusterID, err)
		}

		err = cpiRdeManager.RDEManager.RemoveErrorByNameOrIdFromErrorSet(context.Background(), vcdsdk.ComponentCPI, cpisdk.CPIStatusUpgradeRdeError, cloudConfig.ClusterID, "")
		if err != nil {
			klog.Errorf("failed to remove CPI error [%s] from ErrorSet in RDE [%s], [%v]", cpisdk.CPIStatusUpgradeRdeError, cloudConfig.ClusterID, err)
		}
	}

	var zm *vcdsdk.ZoneMap = nil
	if cloudConfig.VCD.IsZoneEnabledCluster {
		if zm, err = vcdsdk.NewZoneMap(vcdsdk.ZoneMapConfigFilePath); err != nil {
			return nil, fmt.Errorf("unable to create new zone map from configmap file [%s]: [%v]",
				vcdsdk.ZoneMapConfigFilePath, err)
		}
	}

	// cache for VM Info with an refresh of elements needed after 1 minute
	vmInfoCache := newVmInfoCache(vcdClient, cloudConfig.VAppName, time.Minute, zm)

	// TODO: Do we need to record anything from instances from errors/events aspect?
	return &VCDCloudProvider{
		vcdClient:   vcdClient,
		lb:          lb,
		instances:   newInstances(vmInfoCache),
		instancesV2: newInstancesV2(vmInfoCache, zm),
		zoneMap:     zm,
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

	if vcdCP.lb == nil {
		// lb will be nil if the organization network is not backed by NSX-T
		klog.Infof("service controller will be disabled")
		return nil, false
	}
	return vcdCP.lb, true
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
