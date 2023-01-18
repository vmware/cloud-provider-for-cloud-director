package utils

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/testingsdk"
	v1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"time"
)

const (
	defaultNodeInterval        = 2 * time.Second
	defaultNodeNotFoundTimeout = 20 * time.Minute
)

func GetWorkerNode(ctx context.Context, tc *testingsdk.TestClient) (*v1.Node, error) {
	nodes, err := tc.GetWorkerNodes(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting a list of nodes from cluster")
	}

	if len(nodes) < 2 {
		return nil, fmt.Errorf("error: expected 2 or more worker nodes, but found [%d] worker nodes", len(nodes))
	}

	for _, node := range nodes {
		nodeLabelMap := node.Labels
		if _, exists := nodeLabelMap[testingsdk.ControlPlaneLabel]; !exists {
			return &node, nil
		}
	}
	return nil, fmt.Errorf("could not find a worker node")
}

func WaitForWorkerNodeNotFound(cs kubernetes.Interface, workerNodeName string) error {
	return wait.PollImmediate(defaultNodeInterval, defaultNodeNotFoundTimeout, func() (bool, error) {
		if _, err := cs.CoreV1().Nodes().Get(context.TODO(), workerNodeName, metav1.GetOptions{}); err != nil {
			return apierrs.IsNotFound(err), nil
		}
		return false, nil
	})
}
