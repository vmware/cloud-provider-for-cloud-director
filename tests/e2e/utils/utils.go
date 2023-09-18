package utils

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/testingsdk"
	appsv1 "k8s.io/api/apps/v1"
)

const (
	clusterUrnPrefix = "urn:vcloud:entity:vmware:"

	containerName = "test-app"
)

/*
*
CreateDeployment creates a Deployment object with agnhost image which is a E2E testing image.
`netexec` is a command used to run HTTP server on agnhost at certain port specified with --http-port flag.
Reference: https://github.com/kubernetes/kubernetes/issues/90211
*/
func CreateDeployment(ctx context.Context, tc *testingsdk.TestClient, name, namespace, containerImage string, labels map[string]string) (*appsv1.Deployment, error) {
	deploymentParams := &testingsdk.DeployParams{
		Name:   name,
		Labels: labels,
		ContainerParams: testingsdk.ContainerParams{
			ContainerName:  containerName,
			ContainerImage: containerImage,
			ContainerPort:  int32(80), // TODO: In the future, we want to add HTTPS port, we can add it here; ContainerPort would need to be an array instead.
			Args:           []string{"netexec", fmt.Sprintf("--http-port=%d", int32(80))},
		},
	}
	return tc.CreateDeployment(ctx, deploymentParams, namespace)
}
