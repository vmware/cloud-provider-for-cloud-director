package utils

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/testingsdk"
	v1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"strings"
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

func WaitForWorkerNodeNotFound(ctx context.Context, tc *testingsdk.TestClient, workerNode *v1.Node) error {
	return wait.PollImmediate(defaultNodeInterval, defaultNodeNotFoundTimeout, func() (bool, error) {
		if _, err := tc.Cs.CoreV1().Nodes().Get(ctx, workerNode.Name, metav1.GetOptions{}); err != nil {
			return apierrs.IsNotFound(err), nil
		}
		return false, nil
	})
}

func GetVAppNameFromNode(clusterName string, machineSet string, ovdcID string) (string, error) {
	parts := strings.Split(ovdcID, ":")
	if len(parts) != 4 {
		// urn:vcloud:ovdc:<uuid>
		return "", fmt.Errorf("invalid URN format for OVDC: [%s]", ovdcID)
	}
	// machine set name will be mdName-abcd
	endIndex := strings.LastIndex(machineSet, "-")
	if endIndex == -1 {
		return "", fmt.Errorf("machine set name [%s] is not in an expected format", machineSet)
	}
	mdName := machineSet[:endIndex]
	// <cluster_name>_<ovdc_id>_<machine_deployment_name>
	return fmt.Sprintf("%s_%s_%s", clusterName, parts[3], mdName), nil
}
