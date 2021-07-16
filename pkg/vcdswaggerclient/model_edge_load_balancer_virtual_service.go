/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// A Virtual Service for an Edge Gateway.
type EdgeLoadBalancerVirtualService struct {
	// Represents current status of the networking object. 
	Status *NetworkingObjectStatusType `json:"status,omitempty"`
	// The identifier of the Virtual Service in URN format
	Id string `json:"id,omitempty"`
	// The name of the Virtual Service. Name is unique across all Virtual Services for an Edge Gateway.
	Name string `json:"name"`
	// The description of the Virtual Service.
	Description string `json:"description,omitempty"`
	// A flag indicating whether Virtual Service is enabled or not.
	Enabled bool `json:"enabled"`
	// The virtual IP Address (VIP) of the Virtual Service. This IP can be an allocated IP to the Gateway from the External Network or it can be an arbitrary internal IP address used for internal load balancing. It it's an internal IP Address, this IP cannot be part of any existing subnet attached to the Edge Gateway or any vDC Group network if the Edge Gateway is scoped accordingly. 
	VirtualIpAddress string `json:"virtualIpAddress"`
	// The Load Balancer Pool associated with this Virtual Service.
	LoadBalancerPoolRef *EntityReference `json:"loadBalancerPoolRef"`
	// The Edge Gateway associated with this Virtual Service.
	GatewayRef *EntityReference `json:"gatewayRef"`
	// The Load Balancer Service Engine Group that is assigned to the Edge Gateway. This Virtual Service will be deployed to this Service Engine Group. 
	ServiceEngineGroupRef *EntityReference `json:"serviceEngineGroupRef"`
	// The certificate used for SSL termination for the Virtual Service. This is required if the service port type is \"HTTPS\" or \"L4_TLS\".
	CertificateRef *EntityReference `json:"certificateRef,omitempty"`
	// A list of service ports supported by this Virtual Service.  Multiple service ports are allowed only with additional licensing. 
	ServicePorts []EdgeLoadBalancerServicePort `json:"servicePorts"`
	// The current health status of the virtual service. Possible values are: <ul> <li> UP - The virtual service is healthy. <li> DOWN - The virtual service is down, inactive, or has failed. <li> DISABLED - The virtual service is disabled. <li> UNAVAILABLE - The virtual service is unavailable. <li> PENDING - The virtual service is being creating or resources are being allocated. <li> UNKNOWN - The virtual service state is unknown. </ul> 
	HealthStatus string `json:"healthStatus,omitempty"`
	// The localized message on the health of the virtual service.
	HealthMessage string `json:"healthMessage,omitempty"`
	// The non-localized detailed message on the health of the virtual service.
	DetailedHealthMessage string `json:"detailedHealthMessage,omitempty"`
	// Specifies the application profile for the virtual service such as whether it's HTTP, HTTPS, or TCP/UDP.
	ApplicationProfile *EdgeLoadBalancerApplicationProfile `json:"applicationProfile"`
}
