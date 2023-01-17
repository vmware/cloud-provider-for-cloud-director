package utils

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/testingsdk"
	"gopkg.in/yaml.v3"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

const (
	servicePollFrequency = 10 * time.Second
	serviceTimeout       = 5 * time.Minute
	clusterUrnPrefix     = "urn:vcloud:entity:vmware:"
	controlPlaneLabel    = "node-role.kubernetes.io/control-plane"

	containerName  = "test-app"
	containerImage = "k8s.gcr.io/e2e-test-images/agnhost:2.36"

	loadbalancerKey = "loadbalancer"
	networkKey      = "network"
	vipSubnetKey    = "vipSubnet"
)

func GetConfigMap(cs kubernetes.Interface, namespace, name string) (*v1.ConfigMap, error) {
	return cs.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func GetIpamSubnetFromConfigMap(cm *v1.ConfigMap) (string, error) {
	data := cm.Data
	ccmYaml, ok := data["vcloud-ccm-config.yaml"]
	if !ok {
		return "", fmt.Errorf("no data present")
	}

	var result map[string]interface{}
	err := yaml.Unmarshal([]byte(ccmYaml), &result)
	if err != nil {
		return "", fmt.Errorf("err occurred: [%v]", err)
	}

	for key, val := range result {
		if key == loadbalancerKey {
			for k, v := range val.(map[string]interface{}) {
				if k == vipSubnetKey {
					return v.(string), nil
				}
			}
		}
	}
	return "", fmt.Errorf("unable to find vipSubnet from ConfigMap [%s]", cm.Name)
}

func GetNetworkNameFromConfigMap(cm *v1.ConfigMap) (string, error) {
	data := cm.Data
	ccmYaml, ok := data["vcloud-ccm-config.yaml"]
	if !ok {
		return "", fmt.Errorf("no data present")
	}

	var result map[string]interface{}
	err := yaml.Unmarshal([]byte(ccmYaml), &result)
	if err != nil {
		return "", fmt.Errorf("err occurred: [%v]", err)
	}

	for key, val := range result {
		if key == loadbalancerKey {
			for k, v := range val.(map[string]interface{}) {
				if k == networkKey {
					return v.(string), nil
				}
			}
		}
	}
	return "", nil
}

/**
CreateDeployment creates a Deployment object with agnhost image which is a E2E testing image.
`netexec` is a command used to run HTTP server on agnhost at certain port specified with --http-port flag.
Reference: https://github.com/kubernetes/kubernetes/issues/90211
*/
func CreateDeployment(ctx context.Context, tc *testingsdk.TestClient, name, namespace string, labels map[string]string) (*appsv1.Deployment, error) {
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
