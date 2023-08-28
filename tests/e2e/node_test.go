package e2e

import (
	"context"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/testingsdk"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/cloud-provider-for-cloud-director/tests/e2e/utils"
	v1 "k8s.io/api/core/v1"
)

// Note: This test requires at least 2 worker nodes as we will be deleting one of them.
var _ = Describe("Node LCM", func() {
	var (
		err                    error
		workerNode             *v1.Node
		tc                     *testingsdk.TestClient
		deleteTestSpecExecuted bool
	)

	tc, err = utils.NewTestClient(host, org, userOrg, ovdcName, username, token, clusterId, true)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(tc).NotTo(BeNil())
	Expect(&tc.Cs).NotTo(BeNil())
	ctx := context.TODO()

	BeforeEach(func() {
		workerNodes, err := tc.GetWorkerNodes(ctx)
		Expect(err).ShouldNot(HaveOccurred(), fmt.Sprintf("error occurred while fetching list of worker nodes: [%v]", err))
		if !deleteTestSpecExecuted && len(workerNodes) < 2 {
			Skip("Skipping Node LCM test case as this cluster does not have 2 or more worker nodes")
		}
	})

	vdcManager, err := vcdsdk.NewVDCManager(tc.VcdClient, org, ovdcName)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(vdcManager).NotTo(BeNil())

	It("should stop a worker VM in VCD", func() {
		By("ensuring that vApp exists")
		clusterVApp, err := tc.VcdClient.VDC.GetVAppByName(tc.ClusterName, true)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(clusterVApp).NotTo(BeNil())
		Expect(clusterVApp.VApp).NotTo(BeNil())

		By("getting at least 1 worker node from our cluster")
		workerNode, err = utils.GetWorkerNode(ctx, tc)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(workerNode).NotTo(BeNil())

		By("verifying that there exists a VM found worker node's name")
		workerVm, err := vdcManager.FindVMByName(clusterVApp.VApp.Name, workerNode.Name)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(workerVm).NotTo(BeNil())

		By("powering off the worker node vm")
		task, err := workerVm.PowerOff()
		Expect(err).ShouldNot(HaveOccurred())
		err = task.WaitTaskCompletion()
		Expect(err).ShouldNot(HaveOccurred())
	})

	// Suspended condition doesn't exist, ref: https://kubernetes.io/docs/concepts/architecture/nodes/#node-status
	It("should expect powered off VM's node status to be NotReady", func() {
		err := tc.WaitForWorkerNodeNotReady(ctx, workerNode)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should start the worker VM that was powered off in VCD", func() {
		clusterVApp, err := tc.VcdClient.VDC.GetVAppByName(tc.ClusterName, true)
		By("ensuring cluster vApp is present")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(clusterVApp).NotTo(BeNil())
		Expect(clusterVApp.VApp).NotTo(BeNil())

		By("finding the worker VM corresponding to our worker node")
		workerVm, err := vdcManager.FindVMByName(clusterVApp.VApp.Name, workerNode.Name)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(workerVm).NotTo(BeNil())

		By("powering off the worker VM")
		task, err := workerVm.PowerOn()
		Expect(err).ShouldNot(HaveOccurred())
		err = task.WaitTaskCompletion()
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("should expect the worker VM's node status is ready", func() {
		err := tc.WaitForWorkerNodeReady(ctx, workerNode)
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("should stop and delete the the worker VM in VCD", func() {
		By("ensuring cluster vApp is present")
		clusterVApp, err := tc.VcdClient.VDC.GetVAppByName(tc.ClusterName, true)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(clusterVApp).NotTo(BeNil())
		Expect(clusterVApp.VApp).NotTo(BeNil())

		By("finding the worker VM corresponding to our worker node")
		workerVm, err := vdcManager.FindVMByName(clusterVApp.VApp.Name, workerNode.Name)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(workerVm).NotTo(BeNil())

		By("delete the VM")
		err = workerVm.Delete()
		Expect(err).ShouldNot(HaveOccurred())

		deleteTestSpecExecuted = true
	})

	It("should no longer contain the worker VM in Kubernetes list of nodes", func() {
		err = utils.WaitForWorkerNodeNotFound(ctx, tc, workerNode)
		Expect(err).ShouldNot(HaveOccurred())
	})
})
