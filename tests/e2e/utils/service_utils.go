package utils

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/testingsdk"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"time"
)

func WaitForExtIP(cs kubernetes.Interface, namespace string, name string) (string, error) {
	var svc *v1.Service
	var err error

	svc, err = waitForServiceExposure(cs, namespace, name)
	if err != nil {
		return "", err
	}

	if svc == nil {
		return "", fmt.Errorf("the service is nil")
	}
	// We can safely return below as we handled the len(IngressList) check in waitServiceExposure()
	return svc.Status.LoadBalancer.Ingress[0].IP, nil
}

func waitForServiceExposure(cs kubernetes.Interface, namespace string, name string) (*v1.Service, error) {
	var svc *v1.Service
	var err error

	err = wait.PollImmediate(servicePollFrequency, serviceTimeout, func() (bool, error) {
		svc, err = cs.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			// If our error is retryable, it's not a major error. So we do not need to return it as an error.
			if testingsdk.IsRetryableError(err) {
				return false, nil
			}
			return false, err
		}

		IngressList := svc.Status.LoadBalancer.Ingress
		if len(IngressList) == 0 {
			// we'll store an error here and continue to retry until timeout, if this was where we time out eventually, we will return an error at the end
			err = fmt.Errorf("cannot find Ingress components after duration: [%d] minutes", serviceTimeout/time.Minute)
			return false, nil
		}

		ip := svc.Status.LoadBalancer.Ingress[0].IP
		return ip != "", nil // Once we have our IP, we can terminate the condition for polling and return the service
	})

	if err != nil {
		return nil, err
	}
	return svc, nil
}
