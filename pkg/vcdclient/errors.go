package vcdclient

import "fmt"

type VirtualServicePendingError struct {
	VirtualServiceName string
}

func (vsError *VirtualServicePendingError) Error() string {
	return fmt.Sprintf("virtual service [%s] is in Pending state", vsError.VirtualServiceName)
}

func NewVirtualServicePendingError(virtualServiceName string) *VirtualServicePendingError {
	return &VirtualServicePendingError{
		VirtualServiceName: virtualServiceName,
	}
}
