/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

package swagger

// Describes Load Balancer Service configuration on an Edge Gateway. 
type EdgeGatewayLoadBalancerConfig struct {
	// A flag indicating whether Load Balancer Service is enabled or not.
	Enabled bool `json:"enabled"`
	// The network definition in Gateway CIDR format which will be used by Load Balancer service on edge. All the load balancer service engines associated with the Service Engine Group will be attached to this network. The subnet prefix length must be 25. If nothing is set, the default is <code>192.168.255.1/25</code>. Default cidr can be configured. This field cannot be updated. 
	ServiceNetworkDefinition string `json:"serviceNetworkDefinition,omitempty"`
	// Reference to the Load Balancer Cloud. This cloud is used by Edge Gateway's Load Balancer Virtual Services and Pools.
	LoadBalancerCloudRef *EntityReference `json:"loadBalancerCloudRef,omitempty"`
	// The license type of the backing Load Balancer Cloud. <ul> <li>BASIC - Basic edition of the NSX Advanced Load Balancer. <li>ENTERPRISE - Full featured edition of the NSX Advanced Load Balancer. </ul> 
	LicenseType string `json:"licenseType,omitempty"`
}
