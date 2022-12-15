# Kubernetes External Cloud Provider for VMware Cloud Director
This repository contains the [Kubernetes cloud-controller-manager](https://kubernetes.io/docs/concepts/architecture/cloud-controller/) for VMware Cloud Director.

The version of the VMware Cloud Director API and Installation that are compatible for a given cloud-provider container image are described in the following compatibility matrix:

| CPI Version | VMware Cloud Director API | VMware Cloud Director Installation | Notes | Kubernetes Versions |
| :---------: | :-----------------------: | :--------------------------------: | :---: | :------------------ |
| 1.0.0 | 36.0+ | 10.3.1+ | Needs NSX-T 3.1.1 with NSX Advanced Load Balancer (Avi) version 20.1.3+ |<ul><li>1.21</li><li>1.20</li><li>1.19</li></ul>|

This extension is intended to be installed into a Kubernetes cluster installed with [VMware Cloud Director](https://www.vmware.com/products/cloud-director.html) as a Cloud Provider, by a user that has the rights as described in the sections below.

**cloud-provider-for-cloud-director** is distributed as a container image hosted at [Distribution Harbor](https://projects.registry.vmware.com) as `projects.registry.vmware.com/vmware-cloud-director/cloud-provider-for-cloud-director:<CPI Version>`.

This cloud-provider is in a `GA` state and will be supported in production.

## Terminology
1. VCD: VMware Cloud Director
2. ClusterAdminRole: This is the role that has enough rights to create and administer a Kubernetes Cluster in VCD. This role can be created by cloning the [vApp Author Role](https://docs.vmware.com/en/VMware-Cloud-Director/10.3/VMware-Cloud-Director-Tenant-Portal-Guide/GUID-BC504F6B-3D38-4F25-AACF-ED584063754F.html) and then adding the following rights (details on adding the rights below can be found in the [CSE docs](https://github.com/rocknes/container-service-extension/blob/cse_3_1_docs/docs/cse3_1/RBAC.md#additional-required-rights)):
    1. Full Control: CSE:NATIVECLUSTER
    2. Edit: CSE:NATIVECLUSTER
    3. View: CSE:NATIVECLUSTER
3. ClusterAdminUser: For CPI functionality, there needs to be a set of additional rights added to the `ClusterAdminRole` as described in the "Additional Rights for CPI" section below. The Kubernetes Cluster needs to be **created** by a user belonging to this **enhanced** `ClusterAdminRole`. For convenience, let us term this user as the `ClusterAdminUser`.

## VMware Cloud Director Configuration
In this section, we assume that the Kubernetes cluster is created using the [Container Service Extension](https://github.com/vmware/container-service-extension). However that is not a mandatory requirement.

**Note:** NSX-T with NSX Advanced Load Balancer is a prerequisite to use LoadBalancers with CPI of VCD.

### Additional Rights for CPI
The `ClusterAdminUser` should have view access to the vApp containing the Kubernetes cluster. Since the `ClusterAdminUser` itself creates the cluster, it will have this access by default.
This `ClusterAdminUser` needs to be created from a `ClusterAdminRole` with the following additional rights:

1. Gateway =>
    1. View Gateway
2. Gateway Services =>
    1. NAT Configure (adds NAT View)
    2. LoadBalancer Configure (adds LoadBalancer View)
3. Access Control =>
    1. User => Manage user's own API TOKEN

The `Access Control` right is needed in order to generate refresh tokens for the `ClusterAdminUser`.

### Instances Interface: Node Lifecycle Management (LCM)
There is no particular configuration needed in order to use the Node LCM.

### Services Interface: LoadBalancer Configuration

#### Provider Setup
The LoadBalancers using the CPI of VCD need a preconfigured Avi Controller, NSX-T Cloud and Avi Service Engine Group. This is a provider operation.

The Service Engine Group (SEG) should be created as `Dedicated` and one SEG should be allocated per Edge Gateway in order to ensure that Load Balancers used by Tenants are well-isolated from each other.

The LoadBalancer section of the Edge Gateway for a Tenant should be enabled, and the appropriate Service Engine Group(s) should be configured into the Edge Gateway. This will be used to create Virtual Services when a LoadBalancer request is made from Kubernetes.

#### Tenant Setup
A ServiceEngineGroup needs to be added to the gateway of the OVDC within which the Kubernetes cluster is to be created. The overall steps to achieve that are documented at [Enable Load Balancer on an NSX-T Data Center Edge Gateway](https://docs.vmware.com/en/VMware-Cloud-Director/10.3/VMware-Cloud-Director-Service-Provider-Admin-Portal-Guide/GUID-1784B96B-20F8-4E4D-BF33-86D2264EDBCF.html)

#### Creation of a LoadBalancer using a third-party ingress
Any third party ingress such as Contour could be used with the CPI in order to create an L7 ingress and NSX Advanced Load Balancer with Avi will act as the L4 LoadBalancer.

**Note**: In order to create a HTTPS Ingress using the Avi LoadBalancer, a certificate needs to be used. For Kubernetes clusters created using the Container Service Extension, the certificate needs to be named in a particular way as follows:
1. Get the ID of the Kubernetes cluster using the following command. Let us name it `ClusterID`
```
vcd cse cluster info <kubernetes cluster name>
```
2. Upload a certificate in Certificates Library of the Organization using the VCD UI. Name the certificate with the following format:
```
<cluster ID>-cert
```

This will enable all HTTPS ingresses of the Kubernetes cluster to use the same certificate that has been uploaded here.

## Contributing
Please see [CONTRIBUTING.md](CONTRIBUTING.md) for instructions on how to contribute.

## License
[Apache-2.0](LICENSE.txt)
