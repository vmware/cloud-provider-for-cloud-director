package e2e

import (
	"flag"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	clusterId string
	host      string
	org       string
	userOrg   string
	ovdcName  string
	username  string
	token     string

	ContainerImage string
)

const (
	airgappedImage = "harbor.10.221.134.246.nip.io/airgapped/agnhost:2.36"
	stagingImage   = "projects-stg.registry.vmware.com/vmware-cloud-director/agnhost:2.36"
)

func init() {
	//Inputs needed: VCD site, org, ovdc, username, refreshToken, clusterId
	flag.StringVar(&host, "host", "", "VCD host site to generate client")
	flag.StringVar(&org, "org", "", "Cluster Org to generate client")
	flag.StringVar(&userOrg, "userOrg", "", "User Org to generate client")
	flag.StringVar(&ovdcName, "ovdcName", "", "Ovdc Name to generate client")
	flag.StringVar(&username, "username", "", "Username for login to generate client")
	flag.StringVar(&token, "token", "", "Refresh token of user to generate client")
	flag.StringVar(&clusterId, "clusterId", "", "Cluster ID to fetch cluster RDE")
}

var _ = BeforeSuite(func() {
	// We should validate that all credentials are present for generating a TestClient
	Expect(host).NotTo(BeEmpty(), "Please make sure --host is set correctly.")
	Expect(org).NotTo(BeEmpty(), "Please make sure --org is set correctly.")
	Expect(userOrg).NotTo(BeEmpty(), "Please make sure --userOrg is set correctly.")
	Expect(ovdcName).NotTo(BeEmpty(), "Please make sure --ovdcName is set correctly.")
	Expect(username).NotTo(BeEmpty(), "Please make sure --username is set correctly.")
	Expect(token).NotTo(BeEmpty(), "Please make sure --token is set correctly.")
	Expect(clusterId).NotTo(BeEmpty(), "Please make sure --clusterId is set correctly.")

	// AIRGAP environment variable is expected to be set in jenkins or by user via `export AIRGAP="true"`
	useAirgap := os.Getenv("AIRGAP")
	if useAirgap != "" {
		ContainerImage = airgappedImage
	} else {
		ContainerImage = stagingImage
	}
})

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}
