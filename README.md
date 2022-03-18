# Kubernetes External Cloud Provider for VMware Cloud Director
This repository contains the [Kubernetes cloud-controller-manager](https://kubernetes.io/docs/concepts/architecture/cloud-controller/) for VMware Cloud Director.

The version of the VMware Cloud Director API and Installation that are compatible for a given cloud-provider container image are described in the following compatibility matrix:

| CPI Version | VMware Cloud Director API | VMware Cloud Director Installation | Notes | Kubernetes Versions |
| :---------: | :-----------------------: | :--------------------------------: | :---: | :------------------ |
| 1.0.0 | 36.0+ | 10.3.1+ <br/>(10.3.1 needs hot-patch to prevent VCD cell crashes in multi-cell environments) | Needs NSX-T 3.1.1 with NSX Advanced Load Balancer (Avi) version 20.1.3+ |<ul><li>1.21</li><li>1.20</li><li>1.19</li></ul>|
| 1.0.1 | 36.0+ | 10.3.1+ <br/>(10.3.1 needs hot-patch to prevent VCD cell crashes in multi-cell environments) | <ul><li>Some resiliency added when VCD cells are restarted (34c1689)</li><li>Added Org ID to cert query so that `system/administrator` can also create https load balancers (44c72ab)</li></ul> |<ul><li>1.21</li><li>1.20</li><li>1.19</li></ul>|
| 1.0.2 | 36.0+ | 10.3.1+ <br/>(10.3.1 needs hot-patch to prevent VCD cell crashes in multi-cell environments) | <ul><li>Added fix to allow multiple http and https ports to be allowed in load-balancer (d67c19b)</li></ul> |<ul><li>1.21</li><li>1.20</li><li>1.19</li></ul>|
| 1.1.0 | 36.0+ | 10.3.1+ <br/>(10.3.1 needs hot-patch to prevent VCD cell crashes in multi-cell environments) | <ul><li>Remove legacy Kubernetes dependencies.</li><li>Application port profiles added to DNAT rules (Fixes #43)</li><li>L4, HTTP and HTTPS services supported using `appProtocol` and annotations (Fixes #44).</li><li>Allow per-service certificates.</li><li>Multiple(>2) service fixes within the same LoadBalancer service</li><li>Support for CAPVCD RDEs.</li><li>Detect and handle `PENDING` Avi LoadBalancer state to allow better controller functionality.</li></ul> |<ul><li>1.21</li><li>1.20</li><li>1.19</li></ul>|
| 1.1.1 | 36.0+ | 10.3.1+ <br/>(10.3.1 needs hot-patch to prevent VCD cell crashes in multi-cell environments) | Needs NSX-T 3.1.1 with NSX Advanced Load Balancer (Avi) version 20.1.3+ |<ul><li>1.21</li><li>1.20</li><li>1.19</li></ul>|

This extension is intended to be installed into a Kubernetes cluster installed with [VMware Cloud Director](https://www.vmware.com/products/cloud-director.html) as a Cloud Provider, by a user that has the rights as described in the sections below.

**cloud-provider-for-cloud-director** is distributed as a container image hosted at [Distribution Harbor](https://projects.registry.vmware.com) as `projects.registry.vmware.com/vmware-cloud-director/cloud-provider-for-cloud-director:<CPI Version>.latest`.

This cloud-provider is in a `GA` state and will be supported in production.

Note: The cloud-provider is not impacted by the Apache Log4j open source component vulnerability.

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

**Note**: In order to create a HTTPS Ingress using the Avi LoadBalancer, a certificate needs to be used. The following steps present an overview **from CPI 1.1.0 onwards**:
1. As a user with OrgAdmin role, upload a certificate into the Trusted Certificates of the Organization using the VCD UI. Let this certificate be called `my-service-cert`.

2. Add the following annotations to the ingress loadbalancer service. Depending on the installation method used (helm etc), the location of addition of these annotations may be different. The annotation mentions the _comma-separated list of ports_ that need SSL and the (single) certificate to be used for it.
```
annotations:
  service.beta.kubernetes.io/vcloud-avi-ssl-ports: "443"
  service.beta.kubernetes.io/vcloud-avi-ssl-cert-alias: "my-service-cert"
```

3. Install the service

This will enable the HTTPS ingresses of the Kubernetes cluster to use the fore-mentioned certificate that has been uploaded here.

**Note:**
From v1.1.0 onwards, certificates can have user-defined names. Each service could use its own certificate and there does not need to be one common certificate used across services.

## CPI v1.1.1 changes
* Fixed refresh token based authentication issue observed when VCD cells are fronted by a load balancer.
* Updates nodePort and port of LoadBalancer services are now supported.

## Contributing
Please see [CONTRIBUTING.md](CONTRIBUTING.md) for instructions on how to contribute.

## License
[Apache-2.0](LICENSE.txt)
