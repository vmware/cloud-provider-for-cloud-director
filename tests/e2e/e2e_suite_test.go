package e2e

import (
	"flag"
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
	Expect(host).NotTo(BeZero(), "Please make sure --host is set correctly.")
	Expect(org).NotTo(BeZero(), "Please make sure --org is set correctly.")
	Expect(userOrg).NotTo(BeZero(), "Please make sure --userOrg is set correctly.")
	Expect(ovdcName).NotTo(BeZero(), "Please make sure --ovdcName is set correctly.")
	Expect(username).NotTo(BeZero(), "Please make sure --username is set correctly.")
	Expect(token).NotTo(BeZero(), "Please make sure --token is set correctly.")
	Expect(clusterId).NotTo(BeZero(), "Please make sure --clusterId is set correctly.")
})

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}
