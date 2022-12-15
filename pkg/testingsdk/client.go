package testingsdk

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	stov1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
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
	Labels          map[string]string
	VolumeParams    VolumeParams
	ContainerParams ContainerParams
}
type VolumeParams struct {
	VolumeName string
	PvcRef     string
	MountPath  string
}

type ContainerParams struct {
	ContainerName  string
	ContainerImage string
	ContainerPort  int32
	Args           []string
}

func (tc *TestClient) CreateNameSpace(ctx context.Context, nsName string) (*apiv1.Namespace, error) {
	ns, err := createNameSpace(ctx, nsName, tc.Cs.(*kubernetes.Clientset))
	if err != nil {
		return nil, fmt.Errorf("error creating NameSpace [%s] for cluster [%s]: [%v]", nsName, tc.ClusterId, err)
	}
	return ns, nil
}

func (tc *TestClient) CreateStorageClass(ctx context.Context, scName string, reclaimPolicy apiv1.PersistentVolumeReclaimPolicy, storageProfile string) (*stov1.StorageClass, error) {
	sc, err := createStorageClass(ctx, tc.Cs.(*kubernetes.Clientset), scName, reclaimPolicy, storageProfile)
	if err != nil {
		return nil, fmt.Errorf("error creating Storage Class [%s] for cluster [%s]: [%v]", scName, tc.ClusterId, err)
	}
	return sc, nil
}

func (tc *TestClient) CreatePV(ctx context.Context, persistentVolumeName string, storageClass string, storageProfile string, storageSize string, reclaimPolicy apiv1.PersistentVolumeReclaimPolicy) (*apiv1.PersistentVolume, error) {
	pv, err := createPV(ctx, tc.Cs.(*kubernetes.Clientset), persistentVolumeName, storageClass, storageProfile, storageSize, reclaimPolicy)
	if err != nil {
		return nil, fmt.Errorf("error creating Persistent Volume [%s] for cluster [%s]: [%v]", persistentVolumeName, tc.ClusterId, err)
	}
	return pv, nil
}

func (tc *TestClient) CreatePVC(ctx context.Context, nameSpace string, pvcName string, storageClass string, storageSize string) (*apiv1.PersistentVolumeClaim, error) {
	pvc, err := createPVC(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace, pvcName, storageClass, storageSize)
	if err != nil {
		return nil, fmt.Errorf("error creating Persistent Volume Claim [%s] for cluster [%s]: [%v]", pvcName, tc.ClusterId, err)
	}
	return pvc, nil
}

func (tc *TestClient) CreateDeployment(ctx context.Context, params *DeployParams, nameSpace string) (*appsv1.Deployment, error) {
	deployment, err := createDeployment(ctx, tc.Cs.(*kubernetes.Clientset), params, nameSpace)
	if err != nil {
		return nil, fmt.Errorf("error creating Deployment [%s] for cluster [%s]: [%v]", params.Name, tc.ClusterId, err)
	}
	return deployment, nil
}

func (tc *TestClient) CreateLoadBalancerService(ctx context.Context, nameSpace string, serviceName string, annotations map[string]string, labels map[string]string, servicePort []apiv1.ServicePort) (*apiv1.Service, error) {
	lbService, err := createLoadBalancerService(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace, serviceName, annotations, labels, servicePort)
	if err != nil {
		return nil, fmt.Errorf("error creating LoadBalancer Service [%s] for cluster [%s]: [%v]", serviceName, tc.ClusterId, err)
	}
	return lbService, nil
}

func (tc *TestClient) DeletePVC(ctx context.Context, nameSpace string, pvcName string) error {
	err := deletePVC(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace, pvcName)
	if err != nil {
		return fmt.Errorf("error deleting Persistent Volume Claim [%s] for cluster [%s]: [%v]", pvcName, tc.ClusterId, err)
	}
	return nil
}

func (tc *TestClient) DeletePV(ctx context.Context, pvName string) error {
	err := deletePV(ctx, tc.Cs.(*kubernetes.Clientset), pvName)
	if err != nil {
		return fmt.Errorf("error deleting Persistent Volume [%s] for cluster [%s]: [%v]", pvName, tc.ClusterId, err)
	}
	return nil
}

func (tc *TestClient) DeleteDeployment(ctx context.Context, nameSpace string, deploymentName string) error {
	err := deleteDeployment(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace, deploymentName)
	if err != nil {
		return fmt.Errorf("error deleting Deployment [%s] for cluster [%s]: [%v]", deploymentName, tc.ClusterId, err)
	}
	return nil

}

func (tc *TestClient) DeleteNameSpace(ctx context.Context, nameSpace string) error {
	err := deleteNameSpace(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace)
	if err != nil {
		return fmt.Errorf("error deleting NameSpace [%s] for cluster [%s]: [%v]", nameSpace, tc.ClusterId, err)
	}
	return nil
}

func (tc *TestClient) DeleteService(ctx context.Context, nameSpace string, serviceName string) error {
	err := deleteService(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace, serviceName)
	if err != nil {
		return fmt.Errorf("error deleting Service [%s] for cluster [%s]: [%v]", serviceName, tc.ClusterId, err)
	}
	return nil
}

func (tc *TestClient) DeleteStorageClass(ctx context.Context, scName string) error {
	err := deleteStorageClass(ctx, tc.Cs.(*kubernetes.Clientset), scName)
	if err != nil {
		return fmt.Errorf("error deleting Persistent Volume[%s] for cluster [%s]: [%v]", scName, tc.ClusterId, err)
	}
	return nil
}

func (tc *TestClient) GetWorkerNodes(ctx context.Context) ([]apiv1.Node, error) {
	wnPool, err := getWorkerNodes(ctx, tc.Cs.(*kubernetes.Clientset))
	if err != nil {
		if err == ResourceNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("error getting Worker Node Pool for cluster [%s]: [%v]", tc.ClusterId, err)
	}
	return wnPool, nil
}

func (tc *TestClient) GetStorageClass(ctx context.Context, scName string) (*stov1.StorageClass, error) {
	sc, err := getStorageClass(ctx, tc.Cs.(*kubernetes.Clientset), scName)
	if err != nil {
		if err == ResourceNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("error getting Storage Class [%s] for cluster [%s]: [%v]", scName, tc.ClusterId, err)
	}
	return sc, nil
}

func (tc *TestClient) GetPV(ctx context.Context, pvName string) (*apiv1.PersistentVolume, error) {
	pv, err := getPV(ctx, tc.Cs.(*kubernetes.Clientset), pvName)
	if err != nil {
		if err == ResourceNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("error getting Persistent Volume [%s] for cluster [%s]: [%v]", pvName, tc.ClusterId, err)
	}
	return pv, nil
}

func (tc *TestClient) GetPVC(ctx context.Context, nameSpace string, pvcName string) (*apiv1.PersistentVolumeClaim, error) {
	pvc, err := getPVC(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace, pvcName)
	if err != nil {
		if err != ResourceNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("error getting Persistent Volume Claim [%s] for cluster [%s]: [%v]", pvcName, tc.ClusterId, err)
	}
	return pvc, nil
}

func (tc *TestClient) GetDeployment(ctx context.Context, nameSpace string, deployName string) (*appsv1.Deployment, error) {
	deployment, err := getDeployment(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace, deployName)
	if err != nil {
		if err == ResourceNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("error getting Deployment [%s] for cluster [%s]: [%v]", deployName, tc.ClusterId, err)
	}
	return deployment, nil
}

func (tc *TestClient) GetService(ctx context.Context, nameSpace string, serviceName string) (*apiv1.Service, error) {
	svc, err := getService(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace, serviceName)
	if err != nil {
		if err == ResourceNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("error getting Service [%s] for cluster [%s]: [%v]", serviceName, tc.ClusterId, err)
	}
	return svc, nil
}

func (tc *TestClient) IsPvcReady(ctx context.Context, nameSpace string, pvcName string) error {
	err := isPvcReady(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace, pvcName)
	if err != nil {
		return fmt.Errorf("error querying PVC [%s] status for cluster [%s]: [%v]", pvcName, tc.ClusterId, err)
	}
	return nil
}

func (tc *TestClient) IsDeploymentReady(ctx context.Context, nameSpace string, deployName string) error {
	err := isDeploymentReady(ctx, tc.Cs.(*kubernetes.Clientset), nameSpace, deployName)
	if err != nil {
		return fmt.Errorf("error querying Deployment [%s] status for cluster [%s]: [%v]", deployName, tc.ClusterId, err)
	}
	return nil
}

func (tc *TestClient) IsWorkerNodeReady(ctx context.Context, workerNode *apiv1.Node) error {
	err := wait.PollImmediate(defaultNodeInterval, defaultNodeReadyTimeout, func() (bool, error) {
		nodes, err := tc.GetWorkerNodes(ctx)
		if err != nil {
			return false, fmt.Errorf("error getting a list of nodes from cluster")
		}

		for _, node := range nodes {
			if node.Name == workerNode.Name {
				for _, condition := range node.Status.Conditions {
					if condition.Type == apiv1.NodeReady && condition.Status == apiv1.ConditionTrue {
						return true, nil
					}
				}
			}
		}
		return false, nil
	})
	if err != nil {
		return fmt.Errorf("error querying node [%s] status for cluster [%s]: [%v]", workerNode.Name, tc.ClusterId, err)
	}
	return nil
}

// IsWorkerNodeNotReady we cannot use negate result from IsWorkerNodeReady() Set different RetryTimeInterval and avoid timeout error
func (tc *TestClient) IsWorkerNodeNotReady(ctx context.Context, workerNode *apiv1.Node) error {
	err := wait.PollImmediate(defaultNodeInterval, defaultNodeNotReadyTimeout, func() (bool, error) {
		nodes, err := tc.GetWorkerNodes(ctx)
		if err != nil {
			return false, fmt.Errorf("error getting a list of nodes from cluster")
		}

		for _, node := range nodes {
			if node.Name == workerNode.Name {
				for _, condition := range node.Status.Conditions {
					if condition.Type == apiv1.NodeReady && condition.Status != apiv1.ConditionTrue {
						return true, nil
					}
				}
			}
		}
		return false, nil
	})
	if err != nil {
		return fmt.Errorf("error querying node [%s] status for cluster [%s]: [%v]", workerNode.Name, tc.ClusterId, err)
	}
	return nil
}
