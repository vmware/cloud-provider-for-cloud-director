//go:build !testing
// +build !testing

/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package ccm

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	kubeClient *kubernetes.Clientset = nil
)

func getK8SClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to obtain in-cluster config: [%v]", err)
	}

	newConfig, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("unable to create new config from in-cluster config: [%v]", err)
	}

	return newConfig, nil
}

func init() {
	var err error
	if kubeClient, err = getK8SClient(); err != nil {
		panic(fmt.Sprintf("unable to obtain in-cluster client: [%v]", err))
	}

	return
}

// GetK8SClient : Get pre-initialized kube-client
func GetK8SClient() *kubernetes.Clientset {
	return kubeClient
}
