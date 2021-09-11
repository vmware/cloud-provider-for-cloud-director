# Kubernetes External Cloud Provider for VMware Cloud Director

**cloud-provider-for-cloud-director** contains the source code and build methods to implement and deploy a Kubernetes Cloud Provider with [VMware Cloud Director](https://www.vmware.com/products/cloud-director.html). This Cloud Provider implements the controllers for [Services](https://kubernetes.io/docs/concepts/architecture/cloud-controller/#service-controller) and [Instances](https://kubernetes.io/docs/concepts/architecture/cloud-controller/#node-controller) which would enable users to deploy Ingresses and also manage the lifecycle of Kubernetes nodes, which are [VMware Cloud Director VMs](https://docs.vmware.com/en/VMware-Cloud-Director/10.0/com.vmware.vcloud.tenantportal.doc/GUID-8F806B38-2489-4D36-82FF-B23BAFC3B294.html) in this Cloud Provider.

The version of the VMware Cloud Director API and Installation that are compatible for a given cloud-provider container image are described in the following compatibility matrix:

| CPI Version | VMware Cloud Director API | VMware Cloud Director Installation | Notes |
| :---------: | :-----------------------: | :--------------------------------: | :---: |
| 0.1.0-beta | 35.0+ | 10.3.0+ | Needs NSX-T 3.1.1 with NSX Advanced Load Balancer (Avi) version 20.1.3+ |
| 0.1.0 | 36.1+ | 10.3.1 | Needs NSX-T 3.1.1 with NSX Advanced Load Balancer (Avi) version 20.1.3+ |

This extension is intended to be installed into a Kubernetes cluster installed with [VMware Cloud Director](https://www.vmware.com/products/cloud-director.html) as a Cloud Provider, by a user that has the same rights as a [vApp Author](https://docs.vmware.com/en/VMware-Cloud-Director/9.7/com.vmware.vcloud.admin.doc/GUID-BC504F6B-3D38-4F25-AACF-ED584063754F.html).

**cloud-provider-for-cloud-director** is distributed as a container image hosted at [Distribution Harbor](https://projects.registry.vmware.com) as `projects.registry.vmware.com/vmware-cloud-director/cloud-provider-for-cloud-director:<CPI Version>`.

This cloud-provider is in a preliminary `alpha` state and is not yet ready to be used in production.

## VMware Cloud Director Configuration

In this section, we assume that the Kubernetes cluster is created using the [Container Service Extension](https://github.com/vmware/container-service-extension). However that is not a mandatory requirement. We will later describe the process of enabling a user-created Kubernetes Cluster with CPI.

NSX-T with NSX Advanced Load Balancer is a prerequisite to use LoadBalancers with CPI of VCD. 

The Kubernetes cluster on which the CPI needs to be installed must be accessible to the user whose credentials are used by CPI. Let us call the latter user as the `CPI user`.
The `CPI user` can have access to the Kubernetes cluster in onw of two ways:
1. The `CPI user` itself creates the cluster.
2. The `CPI user` is able to view the vApp of the Kubernetes cluster using VCD sharing methods.

In either case, the `CPI user` will be able to manage the cluster.

### Rights
This `CPI user` needs to be created from a `ClusterAdminRole` with additional rights. The `ClusterAdminRole` is a clone of the [vApp Author Role](https://docs.vmware.com/en/VMware-Cloud-Director/9.7/com.vmware.vcloud.admin.doc/GUID-BC504F6B-3D38-4F25-AACF-ED584063754F.html) with the following additional rights: 
1. Other =>
   1. Full Control: CSE:NATIVECLUSTER
   2. Edit: CSE:NATIVECLUSTER
   3. View: CSE:NATIVECLUSTER
2. Gateway =>
   1. View Gateway
3. Gateway Services =>
   1. NAT Configure (adds NAT View)
   2. LoadBalancer Configure (adds LoadBalancer View)

### Additional Setup Steps for 0.1.0-beta
Since there are dependencies on a new feature coming in a future VCD version (10.3.1), there will be a small workaround to register the `CPI user` credentials into the Kubernetes cluster. Hence, these steps are needed for CPI v0.1.0-beta and VCD v10.3 only.

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
A ServiceEngineGroup needs to be added to the gateway of the OVDC within which the Kubernetes cluster is to be created. The overall steps to achieve that are documented at [Enable Load Balancer on an NSX-T Data Center Edge Gateway](https://docs.vmware.com/en/VMware-Cloud-Director/10.2/VMware-Cloud-Director-Service-Provider-Admin-Portal-Guide/GUID-1784B96B-20F8-4E4D-BF33-86D2264EDBCF.html)

## Documentation

Documentation for the Kubernetes External Cloud Provider for VMware Cloud Director can be obtained here:
* TBD


## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for instructions on how to contribute.


## License

[Apache-2.0](LICENSE.txt)
