# Kubernetes External Cloud Provider for VMware Cloud Director
This repository contains the [Kubernetes cloud-controller-manager](https://kubernetes.io/docs/concepts/architecture/cloud-controller/) for VMware Cloud Director.

The version of the VMware Cloud Director API and Installation that are compatible for a given cloud-provider container image are described in the following compatibility matrix:

| CPI Version | VMware Cloud Director API | VMware Cloud Director Installation | Notes |
| :---------: | :-----------------------: | :--------------------------------: | :---: |
| 0.1.0-beta | 36.0+ | 10.3.0+ | Needs NSX-T 3.1.1 with NSX Advanced Load Balancer (Avi) version 20.1.3+ |

This extension is intended to be installed into a Kubernetes cluster installed with [VMware Cloud Director](https://www.vmware.com/products/cloud-director.html) as a Cloud Provider, by a user that has the rights as described in the sections below.

**cloud-provider-for-cloud-director** is distributed as a container image hosted at [Distribution Harbor](https://projects.registry.vmware.com) as `projects.registry.vmware.com/vmware-cloud-director/cloud-provider-for-cloud-director:<CPI Version>`.

This cloud-provider is in a preliminary `beta` state and is not yet ready to be used in production.

## Terminology
1. VCD: VMware Cloud Director
2. ClusterAdminRole: This is a user who has enough rights to create and administer a Kubernetes Cluster in VCD. This role can be created by cloning the [vApp Author Role](https://docs.vmware.com/en/VMware-Cloud-Director/10.3/VMware-Cloud-Director-Tenant-Portal-Guide/GUID-BC504F6B-3D38-4F25-AACF-ED584063754F.html) and then adding the following rights:
   1. Full Control: CSE:NATIVECLUSTER
   2. Edit: CSE:NATIVECLUSTER
   3. View: CSE:NATIVECLUSTER
3. CPI user: CPI needs to be running in the cluster as a user with a set of rights as described in this section and the Rights section below. For convenience, let us term this user as the `CPI user`.

## VMware Cloud Director Configuration
In this section, we assume that the Kubernetes cluster is created using the [Container Service Extension](https://github.com/vmware/container-service-extension). However that is not a mandatory requirement.

NSX-T with NSX Advanced Load Balancer is a prerequisite to use LoadBalancers with CPI of VCD. 

### Rights
The `CPI user` should have view access to the vApp containing the Kubernetes cluster. If the CPI user itself has created the cluster, it will have this access by default.
This `CPI user` needs to be created from a `ClusterAdminRole` with the following additional rights:
1. Gateway =>
   1. View Gateway
2. Gateway Services =>
   1. NAT Configure (adds NAT View)
   2. LoadBalancer Configure (adds LoadBalancer View)

### Additional Setup Steps for 0.1.0-beta
There is a set of additional steps needed in order to feed the `CPI user` credentials into the Kubernetes cluster. These steps lead to a less secure cluster and are only applicable for the Beta release. The GA release of this product will not need these additional steps and will therefore result in a more secure cluster.

These additional steps are as follows:
1. Get the `KUBECONFIG` file from the cluster created. If the cluster was created using the Container Service Extension, the following command can be used:
```
    vcd cse cluster config <cluster name>  > myk8sclusterkubeconfig
    export KUBECONFIG="<path to myk8sclusterkubeconfig>"
```
2. Create a Kubernetes secret with the username and password of the `CPI user` as follows:
```
VCDUSER=$(echo -n '<cpi user name>' | base64)
PASSWORD=$(echo -n '<cpi user password>' | base64)

cat > vcloud-basic-auth.yaml << END
---
apiVersion: v1
kind: Secret
metadata:
name: vcloud-basic-auth
namespace: kube-system
data:
username: "VCDUSER"
password: "$PASSWORD"
END

kubectl apply  -f vcloud-basic-auth.yaml
```   
This will create a secret and in a while start the CPI cleanly with the right credentials. If you wish, you can monitor it as follows:
```
kubectl get po -A -o wide # <== look for the pod whose name starts with `vmware-cloud-director-ccm-`
kubectl logs -f -n kube-system <pod whose name starts with vmware-cloud-director-ccm>
```

After a while, the CPI initializes the nodes of the cluster, and pods such as `core-dns` will move from `Pending` to `Running` state.


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

Note that in order to create a HTTPS Ingress using the Avi LoadBalancer, a certificate needs to be used. For Kubernetes clusters created using the Container Service Extension, the certificate needs to be named in a particular way as follows:
1. Get the ID of the Kubernetes cluster using the following command. Let us name it `ClusterID`
```
vcd cse cluster info <kubernetes cluster name>
```
2. Upload a certificate into the Trusted Certificates of the Organization using the VCD UI. Name the certificate with the following format:
```
<cluster ID>-cert
```

This will enable all HTTPS ingresses of the Kubernetes cluster to use the same certificate that has been uploaded here.

## Contributing
Please see [CONTRIBUTING.md](CONTRIBUTING.md) for instructions on how to contribute.

## License
[Apache-2.0](LICENSE.txt)
