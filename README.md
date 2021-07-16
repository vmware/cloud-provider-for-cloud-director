# Kubernetes External Cloud Provider for VMware Cloud Director

**cloud-provider-for-cloud-director** contains the source code and build methods to implement and deploy a Cloud Provider with [VMware Cloud Director](https://www.vmware.com/products/cloud-director.html) as the Cloud Provider. This Cloud Provider implements the controllers for [Services](https://kubernetes.io/docs/concepts/architecture/cloud-controller/#service-controller) and [Instances](https://kubernetes.io/docs/concepts/architecture/cloud-controller/#node-controller) which would enable users to deploy Ingresses and also manage the lifecycle of Kubernetes nodes, which are [VMware Cloud Director VMs](https://docs.vmware.com/en/VMware-Cloud-Director/10.0/com.vmware.vcloud.tenantportal.doc/GUID-8F806B38-2489-4D36-82FF-B23BAFC3B294.html) in this Cloud Provider.

The version of the VMware Cloud Director API and Installation that are compatible for a given cloud-provider container image are described in the following compatibility matrix:

| CSI Version | VMware Cloud Director API | VMware Cloud Director Installation | Notes |
| :---------: | :-----------------------: | :--------------------------------: | :---: |
| 0.1.0.latest | 35.0+ | 10.2+ | Needs NSX Advanced Load Balancer (Avi) version 20.1.1 |

This extension is intended to be installed into a Kubernetes cluster installed with [VMware Cloud Director](https://www.vmware.com/products/cloud-director.html) as a Cloud Provider, by a user that has the same rights as a [vApp Author](https://docs.vmware.com/en/VMware-Cloud-Director/9.7/com.vmware.vcloud.admin.doc/GUID-BC504F6B-3D38-4F25-AACF-ED584063754F.html).

**cloud-provider-for-cloud-director** is distributed as a container image hosted at [Distribution Harbor](https://projects.registry.vmware.com) as `projects.registry.vmware.com/vmware-cloud-director/cloud-provider-for-cloud-director:<version>`.

This cloud-provider is in a preliminary `alpha` state and is not yet ready to be used in production.

## VMware Cloud Director Configuration

### Instances Interface: Node Lifecycle Management (LCM)
There is no particular configuration needed for a [vApp Author Role](https://docs.vmware.com/en/VMware-Cloud-Director/9.7/com.vmware.vcloud.admin.doc/GUID-BC504F6B-3D38-4F25-AACF-ED584063754F.html) in order to use the Node LCM.

### Services Interface: LoadBalancer Configuration

#### Rights needed by the user
The [vApp Author Role](https://docs.vmware.com/en/VMware-Cloud-Director/9.7/com.vmware.vcloud.admin.doc/GUID-BC504F6B-3D38-4F25-AACF-ED584063754F.html) must have the following additional rights in order to create an NSX-T Load Balancer:
1. Edge Gateway View
2. Nat Configure (which also enables Nat View)
3. LoadBalancer Configure (which also enables LoadBalancer View)

These rights are primarily needed for Avi LoadBalancer creation using the Services interface of the Cloud-Provider. If only the Instances interface is to be used, the vApp Author role does not need these additional rights.

#### Avi ServiceEngineGroup setup
A ServiceEngineGroup needs to be added to the gateway of the OVDC within which the Kubernetes cluster is to be created. The overall steps to achieve that are documented at [Enable Load Balancer on an NSX-T Data Center Edge Gateway](https://docs.vmware.com/en/VMware-Cloud-Director/10.2/VMware-Cloud-Director-Service-Provider-Admin-Portal-Guide/GUID-1784B96B-20F8-4E4D-BF33-86D2264EDBCF.html)


## Documentation

Documentation for the Kubernetes External Cloud Provider for VMware Cloud Director can be obtained here:
* TBD


## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for instructions on how to contribute.


## License

[Apache-2.0](LICENSE.txt)

