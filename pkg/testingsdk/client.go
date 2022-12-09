package testingsdk

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type TestClient struct {
	VcdClient *vcdsdk.Client
	Cs        kubernetes.Interface
	ClusterId string
}

type VCDAuthParams struct {
	Host         string
	OvdcName     string
	OrgName      string
	Username     string
	RefreshToken string
	UserOrg      string
	GetVdcClient bool // This will need to be set to true as it's needed for CSI, but may not be needed for other use cases
}

type DeployParams struct {
	Name            string
	labels          map[string]string
	volumeParams    VolumeParams
	containerParams ContainerParams
}
type VolumeParams struct {
	volumeName string
	pvcRef     string
	mountPath  string
}

type ContainerParams struct {
	ContainerName  string
	ContainerImage string
	ContainerPort  int32
	Args           []string
}

func NewTestClient(params *VCDAuthParams, clusterId string) (*TestClient, error) {
	client, err := getTestVCDClient(params)
	if err != nil {
		return nil, fmt.Errorf("error occured while generating client using [%s:%s] for cluster [%s]: [%v]", params.Username, params.UserOrg, clusterId, err)
	}

	kubeConfig, err := getKubeconfigFromRDEId(client, clusterId)
	if err != nil {
		return nil, fmt.Errorf("unable to get kubeconfig from RDE [%s]: [%v]", clusterId, err)
	}

	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeConfig))
	if err != nil {
		return nil, fmt.Errorf("unable to create RESTConfig using kubeconfig from RDE [%s]: [%v]", clusterId, err)
	}

	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("unable to create clientset using RESTConfig generated from kubeconfig for cluster [%s]: [%v]", clusterId, err)
	}
	return &TestClient{
		VcdClient: client,
		Cs:        cs,
		ClusterId: clusterId,
	}, nil
}

func (tc *TestClient) GetVCDResourceSet(componentName string) ([]vcdsdk.VCDResource, error) {
	vcdResourceSetMap, err := getVcdResourceSetComponentMapFromRDEId(context.TODO(), tc.VcdClient, componentName, tc.ClusterId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving vcd resource set array from RDE [%s]: [%v]", tc.ClusterId, err)
	}
	return convertVcdResourceSetMapToVcdResourceArr(vcdResourceSetMap)
}

// Returns status.component as map[string]interface{}, this will help us narrow down to specific fields such as nodepools, vcdresources, etc
// Components: vcdKe, projector, csi, cpi, capvcd
func (tc *TestClient) GetComponentMapInStatus(componentName string) (map[string]interface{}, error) {
	return getComponentMapInStatus(context.TODO(), tc.VcdClient, tc.ClusterId, componentName)
}
