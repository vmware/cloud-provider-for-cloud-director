package Testingsdk

import (
	"context"
	"errors"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	stov1 "k8s.io/api/storage/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"strings"
	"time"
)

var (
	ResourceExisted   = errors.New("[REX] resource is already existed")
	ResourceNotFound  = errors.New("[RNF] resource is not found")
	ResourceNameNulll = errors.New("[RNN] resource name is null")
)

func isPvcReady(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, pvcName string) (bool, error) {
	retryInterval := 10 * time.Second
	retryTimeout := 60 * time.Second
	err := wait.PollImmediate(retryInterval, retryTimeout, func() (bool, error) {
		ready := false
		pvc, err := getPVC(ctx, k8sClient, nameSpace, pvcName)
		if err != nil {
			if apierrs.IsInternalError(err) || apierrs.IsTimeout(err) || apierrs.IsServerTimeout(err) ||
				apierrs.IsTooManyRequests(err) || utilnet.IsProbableEOF(err) || utilnet.IsConnectionReset(err) {
				return false, nil
			}
			return false, fmt.Errorf("unexpected error occurred while getting pvc [%s]", pvcName)
		}
		if err != nil {
			return false, nil
		}
		if pvc != nil && pvc.Status.Phase == "Bound" {
			ready = true
		}
		if !ready {
			fmt.Printf("pvc %s is not bound\n", pvc.Name)
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return false, fmt.Errorf("error occurred while checking PVC [%s] status", pvcName)
	}
	return true, nil
}

func isDeploymentReady(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, deployName string) (bool, error) {
	retryInterval := 20 * time.Second
	retryTimeout := 140 * time.Second

	err := wait.PollImmediate(retryInterval, retryTimeout, func() (bool, error) {
		pods, err := getPodByPrefix(ctx, k8sClient, nameSpace, deployName)
		if err != nil {
			if apierrs.IsInternalError(err) || apierrs.IsTimeout(err) || apierrs.IsServerTimeout(err) ||
				apierrs.IsTooManyRequests(err) || utilnet.IsProbableEOF(err) || utilnet.IsConnectionReset(err) {
				return false, nil
			}
			return false, fmt.Errorf("unexpected error occurred while getting deployment [%s]", deployName)
		}
		podCount := len(pods)

		ready := 0
		for _, pod := range pods {
			if pod.Status.Phase == apiv1.PodRunning {
				ready++
			}
		}
		if ready < podCount {
			fmt.Printf("running pods: %v < %v\n", ready, podCount)
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func isPVDeleted(ctx context.Context, k8sClient *kubernetes.Clientset, pvName string) (bool, error) {
	retryInterval := 20 * time.Second
	retryTimeout := 140 * time.Second
	err := wait.PollImmediate(retryInterval, retryTimeout, func() (bool, error) {
		_, err := getPV(ctx, k8sClient, pvName)
		if err != nil {
			if err == ResourceNotFound {
				return true, nil
			}
			if apierrs.IsInternalError(err) || apierrs.IsTimeout(err) || apierrs.IsServerTimeout(err) ||
				apierrs.IsTooManyRequests(err) || utilnet.IsProbableEOF(err) || utilnet.IsConnectionReset(err) {
				return false, nil
			}
			return false, fmt.Errorf("unexpected error occurred while getting persistentVolume [%s]", pvName)
		}
		return false, nil
	})
	if err != nil {
		return false, fmt.Errorf("error occurred while checking PV status: %v", err)
	}
	return true, nil
}

func isPVCDeleted(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, pvcName string) (bool, error) {
	retryInterval := 20 * time.Second
	retryTimeout := 140 * time.Second
	err := wait.PollImmediate(retryInterval, retryTimeout, func() (bool, error) {
		_, err := getPVC(ctx, k8sClient, pvcName, nameSpace)
		if err != nil {
			if err == ResourceNotFound {
				return true, nil
			}
			if apierrs.IsInternalError(err) || apierrs.IsTimeout(err) || apierrs.IsServerTimeout(err) ||
				apierrs.IsTooManyRequests(err) || utilnet.IsProbableEOF(err) || utilnet.IsConnectionReset(err) {
				return false, nil
			}
			return false, fmt.Errorf("unexpected error occurred while getting persistentVolumeClaim [%s]", pvcName)
		}
		return false, nil
	})
	if err != nil {
		return false, fmt.Errorf("error occurred while checking PVC status: %v", err)
	}
	return true, nil
}

func isDeploymentDeleted(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, deployName string) (bool, error) {
	retryInterval := 20 * time.Second
	retryTimeout := 140 * time.Second
	err := wait.PollImmediate(retryInterval, retryTimeout, func() (bool, error) {
		_, err := getDeployment(ctx, k8sClient, nameSpace, deployName)
		if err != nil {
			if err == ResourceNotFound {
				return true, nil
			}
			if apierrs.IsInternalError(err) || apierrs.IsTimeout(err) || apierrs.IsServerTimeout(err) ||
				apierrs.IsTooManyRequests(err) || utilnet.IsProbableEOF(err) || utilnet.IsConnectionReset(err) {
				return false, nil
			}
			return false, fmt.Errorf("unexpected error occurred while getting deployment [%s]", deployName)
		}
		return false, nil
	})
	if err != nil {
		return false, fmt.Errorf("error occurred while checking deployment status: %v", err)
	}
	return true, nil
}

func isServiceDeleted(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, serviceName string) (bool, error) {
	retryInterval := 20 * time.Second
	retryTimeout := 140 * time.Second
	err := wait.PollImmediate(retryInterval, retryTimeout, func() (bool, error) {
		_, err := getDeployment(ctx, k8sClient, nameSpace, serviceName)
		if err != nil {
			if err == ResourceNotFound {
				return true, nil
			}
			if apierrs.IsInternalError(err) || apierrs.IsTimeout(err) || apierrs.IsServerTimeout(err) ||
				apierrs.IsTooManyRequests(err) || utilnet.IsProbableEOF(err) || utilnet.IsConnectionReset(err) {
				return false, nil
			}
			return false, fmt.Errorf("unexpected error occurred while getting service [%s]", serviceName)
		}
		return false, nil
	})
	if err != nil {
		return false, fmt.Errorf("error occurred while checking serviceName status: %v", err)
	}
	return true, nil
}

func isNameSpaceDeleted(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string) (bool, error) {
	retryInterval := 60 * time.Second
	retryTimeout := 300 * time.Second
	err := wait.PollImmediate(retryInterval, retryTimeout, func() (bool, error) {
		_, err := k8sClient.CoreV1().Namespaces().Get(ctx, nameSpace, metav1.GetOptions{})
		if err != nil {
			if err.(*apierrs.StatusError).ErrStatus.Reason == "NotFound" {
				return true, nil
			}
			if apierrs.IsInternalError(err) || apierrs.IsTimeout(err) || apierrs.IsServerTimeout(err) ||
				apierrs.IsTooManyRequests(err) || utilnet.IsProbableEOF(err) || utilnet.IsConnectionReset(err) {
				return false, nil
			}
			return false, fmt.Errorf("unexpected error occurred while getting namespace [%s]", nameSpace)
		}
		return false, nil
	})
	if err != nil {
		return false, fmt.Errorf("error occurred while checking namespace status: %v", err)
	}
	return true, nil
}

func isStorageClassDeleted(ctx context.Context, k8sClient *kubernetes.Clientset, scName string) (bool, error) {
	retryInterval := 60 * time.Second
	retryTimeout := 300 * time.Second
	err := wait.PollImmediate(retryInterval, retryTimeout, func() (bool, error) {
		_, err := k8sClient.StorageV1().StorageClasses().Get(ctx, scName, metav1.GetOptions{})
		if err != nil {
			if err.(*apierrs.StatusError).ErrStatus.Reason == "NotFound" {
				return true, nil
			}
			if apierrs.IsInternalError(err) || apierrs.IsTimeout(err) || apierrs.IsServerTimeout(err) ||
				apierrs.IsTooManyRequests(err) || utilnet.IsProbableEOF(err) || utilnet.IsConnectionReset(err) {
				return false, nil
			}
			return false, fmt.Errorf("unexpected error occurred while getting storage class [%s]")
		}
		return false, nil
	})
	if err != nil {
		return false, fmt.Errorf("error occurred while checking namespace status: %v", err)
	}
	return true, nil
}

func getStorageClass(ctx context.Context, k8sClient *kubernetes.Clientset, scName string) (*stov1.StorageClass, error) {
	if scName == "" {
		return nil, ResourceNameNulll
	}
	sc, err := k8sClient.StorageV1().StorageClasses().Get(ctx, scName, metav1.GetOptions{})
	if err != nil {
		if err.(*apierrs.StatusError).ErrStatus.Reason == "NotFound" {
			return nil, ResourceNotFound
		}
		return nil, err
	}
	return sc, nil
}

func getPodByPrefix(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, podNamePrefix string) ([]*apiv1.Pod, error) {
	if podNamePrefix == "" {
		return nil, ResourceNameNulll
	}
	pods, err := k8sClient.CoreV1().Pods(nameSpace).List(ctx, metav1.ListOptions{})
	if err != nil {
		if err.(*apierrs.StatusError).ErrStatus.Reason == "NotFound" {
			return nil, ResourceNotFound
		}
		return nil, err
	}
	podList := []*apiv1.Pod{}
	for _, pod := range pods.Items {
		if strings.HasPrefix(pod.Name, podNamePrefix) {
			podList = append(podList, &pod)
		}
	}
	return podList, nil
}

func getPV(ctx context.Context, k8sClient *kubernetes.Clientset, pvName string) (*apiv1.PersistentVolume, error) {
	if pvName == "" {
		return nil, ResourceNameNulll
	}
	pv, err := k8sClient.CoreV1().PersistentVolumes().Get(ctx, pvName, metav1.GetOptions{})
	if err != nil {
		if err.(*apierrs.StatusError).ErrStatus.Reason == "NotFound" {
			return nil, ResourceNotFound
		}
		return nil, err
	}
	return pv, nil
}

func getPVC(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, pvcName string) (*apiv1.PersistentVolumeClaim, error) {
	if pvcName == "" {
		return nil, ResourceNameNulll
	}
	pvc, err := k8sClient.CoreV1().PersistentVolumeClaims(nameSpace).Get(ctx, pvcName, metav1.GetOptions{})
	if err != nil {
		if err.(*apierrs.StatusError).ErrStatus.Reason == "NotFound" {
			return nil, ResourceNotFound
		}
		return nil, err
	}
	return pvc, nil
}

func getDeployment(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, deployName string) (*appsv1.Deployment, error) {
	if deployName == "" {
		return nil, ResourceNameNulll
	}
	deployment, err := k8sClient.AppsV1().Deployments(nameSpace).Get(ctx, deployName, metav1.GetOptions{})
	if err != nil {
		if err.(*apierrs.StatusError).ErrStatus.Reason == "NotFound" {
			return nil, ResourceNotFound
		}
		return nil, err
	}
	return deployment, nil
}

func getService(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, serviceName string) (*apiv1.Service, error) {
	if serviceName == "" {
		return nil, ResourceNameNulll
	}
	svc, err := k8sClient.CoreV1().Services(nameSpace).Get(ctx, serviceName, metav1.GetOptions{})
	if err != nil {
		if err.(*apierrs.StatusError).ErrStatus.Reason == "NotFound" {
			return nil, ResourceNotFound
		}
		return nil, err
	}
	return svc, nil
}

func createStorageClass(ctx context.Context, k8sClient *kubernetes.Clientset, scName string, reclaimPolicy string, storageProfile string) (*stov1.StorageClass, error) {
	if scName == "" {
		return nil, ResourceNameNulll
	}
	var PersistentVolumeReclaimPolicy = apiv1.PersistentVolumeReclaimPolicy(reclaimPolicy)
	sc := &stov1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: scName,
			Annotations: map[string]string{
				"storageclass.kubernetes.io/is-default-class": "false",
			},
		},
		ReclaimPolicy: &PersistentVolumeReclaimPolicy,
		Provisioner:   "named-disk.csi.cloud-director.vmware.com",
		Parameters: map[string]string{
			"storageProfile": storageProfile,
			"filesystem":     "ext4",
		},
	}
	newSC, err := k8sClient.StorageV1().StorageClasses().Create(ctx, sc, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error occurred while creating new storageclass [%s]: %v", scName, err)
	}
	return newSC, nil
}

func createNameSpace(ctx context.Context, nsName string, k8sClient *kubernetes.Clientset) (*apiv1.Namespace, error) {
	if nsName == "" {
		return nil, ResourceNameNulll
	}
	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: nsName,
		},
	}
	ns, err := k8sClient.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error occurred while creating namespace [%s]", nsName)
	}
	return ns, nil
}

func createPV(ctx context.Context, k8sClient *kubernetes.Clientset, persistentVolumeName string, storageClass string, storageProfile string, storageSize string) (*apiv1.PersistentVolume, error) {
	if persistentVolumeName == "" {
		return nil, ResourceNameNulll
	}
	persistentVolumeFilesystem := apiv1.PersistentVolumeFilesystem
	pv := &apiv1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: persistentVolumeName,
			Annotations: map[string]string{
				"pv.kubernetes.io/provisioned-by": "named-disk.csi.cloud-director.vmware.com",
			},
		},
		Spec: apiv1.PersistentVolumeSpec{
			StorageClassName: storageClass,
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				"ReadWriteOnce",
			},
			PersistentVolumeSource: apiv1.PersistentVolumeSource{
				CSI: &apiv1.CSIPersistentVolumeSource{
					Driver:       "named-disk.csi.cloud-director.vmware.com",
					FSType:       "ext4",
					VolumeHandle: persistentVolumeName,
					VolumeAttributes: map[string]string{
						"busSubType":     "VirtualSCSI",
						"busType":        "SCSI",
						"filesystem":     "ext4",
						"storageProfile": storageProfile,
					},
				},
			},
			Capacity: apiv1.ResourceList{
				"storage": resource.MustParse(storageSize),
			},
			VolumeMode: &persistentVolumeFilesystem,
		},
	}
	_, err := k8sClient.CoreV1().PersistentVolumes().Create(ctx, pv, metav1.CreateOptions{})
	newPV, err := k8sClient.CoreV1().PersistentVolumes().Create(ctx, pv, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error occurred while creating persistent volume [%s]: [%v]", persistentVolumeName, err)
	}
	return newPV, nil
}

func createPVC(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, pvcName string, storageClass string, storageSize string) (*apiv1.PersistentVolumeClaim, error) {
	if pvcName == "" {
		return nil, ResourceNameNulll
	}
	if nameSpace == "" {
		nameSpace = apiv1.NamespaceDefault
	}
	var storageClassName = storageClass
	pvc := &apiv1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pvcName,
			Namespace: nameSpace,
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			StorageClassName: &storageClassName,
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.ReadWriteOnce,
			},
			Resources: apiv1.ResourceRequirements{
				Requests: apiv1.ResourceList{
					apiv1.ResourceStorage: resource.MustParse(storageSize),
				},
			},
		},
	}

	newPVC, err := k8sClient.CoreV1().PersistentVolumeClaims(nameSpace).Create(ctx, pvc, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error occurred while creating persistent volume claim [%s]: [%v]", pvcName, err)
	}
	return newPVC, nil
}

func createDeployment(ctx context.Context, k8sClient *kubernetes.Clientset, params *DeployParams, nameSpace string) (*appsv1.Deployment, error) {
	if params.Name == "" {
		return nil, ResourceNameNulll
	}
	if nameSpace == "" {
		nameSpace = apiv1.NamespaceDefault
	}
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      params.Name,
			Namespace: nameSpace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: params.labels,
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: params.labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Image: params.containerParams.ContainerImage,
							Name:  params.containerParams.ContainerName,
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: params.containerParams.ContainerPort,
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      params.volumeParams.volumeName,
									MountPath: params.volumeParams.mountPath,
								},
							},
						},
					},
					Volumes: []apiv1.Volume{
						{
							Name: params.volumeParams.volumeName,
							VolumeSource: apiv1.VolumeSource{
								PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
									ClaimName: params.volumeParams.pvcRef,
								},
							},
						},
					},
				},
			},
		},
	}
	newDeployment, err := k8sClient.AppsV1().Deployments(nameSpace).Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error occurred while creating deployment [%s]: [%v]", params.Name, err)
	}
	return newDeployment, nil
}

func createLoadBalancerService(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, serviceName string, annotations map[string]string, labels map[string]string, servicePort []apiv1.ServicePort) (*apiv1.Service, error) {
	if serviceName == "" {
		return nil, ResourceNameNulll
	}
	if nameSpace == "" {
		nameSpace = apiv1.NamespaceDefault
	}
	svc := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        serviceName,
			Namespace:   nameSpace,
			Annotations: annotations,
			Labels:      labels,
		},
		Spec: apiv1.ServiceSpec{
			Ports:    servicePort,
			Selector: labels,
			Type:     "LoadBalancer",
		},
	}
	newSVC, err := k8sClient.CoreV1().Services(nameSpace).Create(ctx, svc, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error occurred while creating service [%s]: [%v]", serviceName, err)
	}
	return newSVC, nil
}

func deletePVC(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, pvcName string) error {
	if pvcName == "" {
		return ResourceNameNulll
	}
	pvc, err := getPVC(ctx, k8sClient, nameSpace, pvcName)
	if err != nil {
		if err == ResourceNotFound {
			return fmt.Errorf("the persistentVolumeClaim [%s] does not exist", pvcName)
		}
		return fmt.Errorf("error occurred while deleting persistentVolumeClaim [%s]: [%v]", pvcName, err)
	}
	err = k8sClient.CoreV1().PersistentVolumeClaims(nameSpace).Delete(ctx, pvcName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete persistentVolumeClaim [%s]", pvcName)
	}
	pv, err := getPV(ctx, k8sClient, pvc.Spec.VolumeName)
	if pv.Spec.PersistentVolumeReclaimPolicy == apiv1.PersistentVolumeReclaimDelete {
		pvDeleted, err := isPVDeleted(ctx, k8sClient, pvc.Spec.VolumeName)
		if err != nil {
			return fmt.Errorf("error occurred while deleting persistentVolume [%s]: [%v]", pvc.Spec.VolumeName, err)
		}
		if !pvDeleted {
			return fmt.Errorf("persistentVolume [%s] still exists", pvc.Spec.VolumeName)
		}
	}
	pvcDeleted, err := isPVCDeleted(ctx, k8sClient, nameSpace, pvcName)
	if err != nil {
		return fmt.Errorf("error occurred while deleting persistentVolumeClaim [%s]: [%v]", pvcName, err)
	}
	if !pvcDeleted {
		return fmt.Errorf("persistentVolumeClaim [%s] still exists", pvcName)
	}
	return nil
}

func deletePV(ctx context.Context, k8sClient *kubernetes.Clientset, pvName string) error {
	_, err := getPV(ctx, k8sClient, pvName)
	if err != nil {
		return fmt.Errorf("the persistentVolumeClaim [%s] does not exist", pvName)
	}
	err = k8sClient.CoreV1().PersistentVolumes().Delete(ctx, pvName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete persistentVolume [%s]", pvName)
	}
	pvDeleted, err := isPVDeleted(ctx, k8sClient, pvName)
	if err != nil {
		return fmt.Errorf("error occurred while deleting persistentVolume [%s]: [%v]", pvName, err)
	}
	if !pvDeleted {
		return fmt.Errorf("persistentVolume [%s] still exists", pvName)
	}
	return nil
}

func deleteDeployment(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, deploymentName string) error {
	_, err := getDeployment(ctx, k8sClient, nameSpace, deploymentName)
	if err != nil {
		if err == ResourceNotFound {
			return fmt.Errorf("the deployment [%s] does not exist", deploymentName)
		}
		klog.Info("error occurred while getting deployment [%s]: [%v]", deploymentName, err)
	}
	err = k8sClient.AppsV1().Deployments(nameSpace).Delete(ctx, deploymentName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete deployment [%s]", deploymentName)
	}
	deploymentDeleted, err := isDeploymentDeleted(ctx, k8sClient, nameSpace, deploymentName)
	if err != nil {
		return fmt.Errorf("error occurred while deleting deployment [%s]: [%v]", deploymentName, err)
	}
	if !deploymentDeleted {
		return fmt.Errorf("deployment [%s] still exists", deploymentName)
	}
	return nil
}

func deleteNameSpace(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string) error {
	err := k8sClient.CoreV1().Namespaces().Delete(ctx, nameSpace, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete namespace [%s]", nameSpace)
	}
	namespaceDeleted, err := isNameSpaceDeleted(ctx, k8sClient, nameSpace)
	if err != nil {
		return fmt.Errorf("error occurred while deleting namespace [%s]: [%v]", nameSpace, err)
	}
	if !namespaceDeleted {
		return fmt.Errorf("namespace [%s] still exists", nameSpace)
	}
	return nil
}

func deleteService(ctx context.Context, k8sClient *kubernetes.Clientset, nameSpace string, serviceName string) error {
	_, err := getService(ctx, k8sClient, nameSpace, serviceName)
	if err != nil {
		if err == ResourceNotFound {
			return fmt.Errorf("the service [%s] does not exist", serviceName)
		}
		klog.Info("error occurred while getting service [%s]: [%v]", serviceName, err)
	}
	err = k8sClient.CoreV1().Services(nameSpace).Delete(ctx, serviceName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete service [%s]", serviceName)
	}
	serviceDeleted, err := isServiceDeleted(ctx, k8sClient, nameSpace, serviceName)
	if err != nil {
		return fmt.Errorf("error occurred while deleting service [%s]: [%v]", serviceName, err)
	}
	if !serviceDeleted {
		return fmt.Errorf("service [%s] still exists", serviceName)
	}
	return nil
}

func deleteStorageClass(ctx context.Context, k8sClient *kubernetes.Clientset, scName string) error {
	_, err := getStorageClass(ctx, k8sClient, scName)
	if err != nil {
		if err == ResourceNotFound {
			return fmt.Errorf("the storageClass [%s] does not exist", scName)
		}
		klog.Info("error occurred while getting storageClass [%s]: [%v]", scName, err)
	}
	err = k8sClient.StorageV1().StorageClasses().Delete(ctx, scName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete service [%s]", scName)
	}
	scDeleted, err := isStorageClassDeleted(ctx, k8sClient, scName)
	if err != nil {
		return fmt.Errorf("error occurred while deleting storageClass [%s]: [%v]", scName, err)
	}
	if !scDeleted {
		return fmt.Errorf("storageClass [%s] still exists", scName)
	}
	return nil
}
