/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/ccm"
	"github.com/vmware/cloud-provider-for-cloud-director/version"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/server/healthz"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/cloud-provider/app"
	"k8s.io/cloud-provider/app/config"
	"k8s.io/cloud-provider/options"
	"k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	_ "k8s.io/component-base/metrics/prometheus/clientgo" // load all the prometheus client-go plugins
	_ "k8s.io/component-base/metrics/prometheus/version"  // for version metric registration
	"k8s.io/klog"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	healthz.InstallHandler(http.DefaultServeMux)

	return
}

func main() {

	rand.Seed(time.Now().UnixNano())

	opts, err := options.NewCloudControllerManagerOptions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to construct options: %v\n", err)
		os.Exit(1)
	}
	opts.KubeCloudShared.CloudProvider.Name = ccm.ProviderName
	opts.Authentication.SkipInClusterLookup = true

	command := app.NewCloudControllerManagerCommand(
		opts,
		vcdCCMInitializer,
		app.DefaultInitFuncConstructors,
		flag.NamedFlagSets{},
		wait.NeverStop,
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	klog.Infof("Initializing CCM [%s]: [%s]", ccm.ProviderName, version.Version)

	// This uses os.args[1:] behind the covers
	if err := command.Execute(); err != nil {
		klog.Errorf("unable to execute command: [%v]", err)
		os.Exit(1)
	}

	return
}

func vcdCCMInitializer(cfg *config.CompletedConfig) cloudprovider.Interface {
	cloudConfig := cfg.ComponentConfig.KubeCloudShared.CloudProvider

	// initialize cloud provider with the cloud provider name and config file provided
	cloud, err := cloudprovider.InitCloudProvider(cloudConfig.Name, cloudConfig.CloudConfigFile)
	if err != nil {
		klog.Fatalf("Cloud provider could not be initialized: [%v]", err)
	}
	if cloud == nil {
		klog.Fatalf("Cloud provider is nil")
	}

	return cloud
}
