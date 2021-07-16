/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	ccm "github.com/vmware/cloud-provider-for-cloud-director/pkg/ccm"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/component-base/logs"
	"k8s.io/klog"
	"k8s.io/kubernetes/cmd/cloud-controller-manager/app"
	_ "k8s.io/kubernetes/pkg/client/metrics/prometheus" // for client metric registration
	_ "k8s.io/kubernetes/pkg/version/prometheus"        // for version metric registration
)

const (
	version = `0.1.0`
)

func init() {
	healthz.InstallHandler(http.DefaultServeMux)

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewCloudControllerManagerCommand()

	logs.InitLogs()
	defer logs.FlushLogs()

	klog.Infof("Initializing CCM [%s]: [%s]", ccm.ProviderName, version)

	// This uses os.args[1:] behind the covers
	if err := command.Execute(); err != nil {
		panic(fmt.Sprintf("unable to parse command: [%v]\n", err))
	}
}
