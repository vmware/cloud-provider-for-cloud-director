package kubeconfig_test

import (
	"context"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/testingsdk"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/cloud-provider-for-cloud-director/tests/kubeconfig_test/utils"
	"os"
)

var _ = Describe("Fetch Kubeconfig from RDE", func() {
	var (
		err        error
		vcdClient  *vcdsdk.Client
		kubeConfig string
	)

	vcdClient, err = vcdsdk.NewVCDClientFromSecrets(host, org, ovdcName, userOrg, username, "", token, true, false)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(vcdClient).NotTo(BeNil())

	It("should be able to fetch kubeconfig from RDE ID", func() {
		kubeConfig, err = testingsdk.GetKubeconfigFromRDEId(context.TODO(), vcdClient, clusterId)
		Expect(err).NotTo(HaveOccurred())
		Expect(kubeConfig).NotTo(BeEmpty())

	})

	It("should be able to write a kubeconfig file to output path", func() {
		// Write kubeconfig directly to a file for access.
		kubeconfigFile := []byte(kubeConfig)
		Expect(kubeconfigFile).NotTo(BeEmpty())

		// Trim trailing slashes if there are any passed in, and use that path to create a kubeconfig.
		exportPath := utils.TrimTrailingSlash(kubeconfigExportPath)
		outputPath := fmt.Sprintf("%s/%s.conf", exportPath, "cluster_kubeconfig")
		err = os.WriteFile(outputPath, kubeconfigFile, 0644)
		Expect(err).NotTo(HaveOccurred())

		// Check if the kubeconfig file exists.
		_, err = os.Stat(outputPath)
		Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("error occurred: [%v]", err))
	})
})
