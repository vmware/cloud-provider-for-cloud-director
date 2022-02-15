package vcdclient

import "fmt"

type VirtualServicePendingError struct {
	VirtualServiceName string
}

type LoadBalancerPoolBusyError struct {
	LBPoolName string
}

func (vsError *VirtualServicePendingError) Error() string {
	return fmt.Sprintf("virtual service [%s] is in Pending state", vsError.VirtualServiceName)
}

func NewVirtualServicePendingError(virtualServiceName string) *VirtualServicePendingError {
	return &VirtualServicePendingError{
		VirtualServiceName: virtualServiceName,
	}
}

func (lbPoolError *LoadBalancerPoolBusyError) Error() string {
	return fmt.Sprintf("load balancer pool [%s] is busy", lbPoolError.LBPoolName)
}

func NewLBPoolBusyError(lbPoolName string) *LoadBalancerPoolBusyError {
	return &LoadBalancerPoolBusyError{
		LBPoolName: lbPoolName,
	}
}
