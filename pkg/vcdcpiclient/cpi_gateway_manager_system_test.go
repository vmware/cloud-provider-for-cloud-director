package vcdcpiclient

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	_ "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"path/filepath"
	"testing"
)

var (
	GitRoot string = ""
)

func init() {
	GitRoot = os.Getenv("GITROOT")
	if GitRoot == "" {
		// It is okay to panic here as this will be caught during dev
		panic("GITROOT should be set")
	}
}

func TestLoadBalancerCRUDE(t *testing.T) {

	authFile := filepath.Join(GitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails vcdsdk.AuthorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	cloudConfig, err := vcdsdk.GetTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := vcdsdk.GetTestVCDClient(cloudConfig, map[string]interface{}{
		"user":         authDetails.Username,
		"secret":       authDetails.Password,
		"userOrg":      authDetails.UserOrg,
		"getVdcClient": true,
	})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	ctx := context.Background()

	cgm, err := NewCpiGatewayManager(ctx, vcdClient, cloudConfig.LB.VDCNetwork, cloudConfig.LB.VIPSubnet, cloudConfig.ClusterID)
	assert.NoError(t, err, "gateway manager should be created without error")

	virtualServiceNamePrefix := fmt.Sprintf("test-virtual-service-https-%s", uuid.New().String())
	lbPoolNamePrefix := fmt.Sprintf("test-lb-pool-%s", uuid.New().String())
	certName := cloudConfig.LB.CertificateAlias
	if certName == "" {
		certName = fmt.Sprintf("%s-cert", cloudConfig.ClusterID)
	}
	portDetailsList := []vcdsdk.PortDetails{
		{
			PortSuffix:   `http`,
			ExternalPort: 80,
			InternalPort: 31234,
			Protocol:     "HTTP",
			UseSSL:       false,
		},
		{
			PortSuffix:   `https`,
			ExternalPort: 443,
			InternalPort: 31235,
			Protocol:     "HTTPS",
			UseSSL:       true,
			CertAlias:    certName,
		},
	}
	freeIP := ""
	oneArm := &vcdsdk.OneArm{
		StartIP: cloudConfig.LB.OneArm.StartIP,
		EndIP: cloudConfig.LB.OneArm.EndIP,
	}
	freeIP, err = cgm.CreateLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix, []string{"1.2.3.4", "1.2.3.5"}, portDetailsList, oneArm)
	assert.NoError(t, err, "Load Balancer should be created")
	assert.NotEmpty(t, freeIP, "There should be a non-empty IP returned")

	virtualServiceNameHttp := fmt.Sprintf("%s-http", virtualServiceNamePrefix)
	freeIPObtained, err := cgm.GetLoadBalancer(ctx, virtualServiceNameHttp, oneArm)
	assert.NoError(t, err, "Load Balancer should be found")
	assert.Equal(t, freeIP, freeIPObtained, "The IPs should match")

	virtualServiceNameHttps := fmt.Sprintf("%s-https", virtualServiceNamePrefix)
	freeIPObtained, err = cgm.GetLoadBalancer(ctx, virtualServiceNameHttps, oneArm)
	assert.NoError(t, err, "Load Balancer should be found")
	assert.Equal(t, freeIP, freeIPObtained, "The IPs should match")

	freeIP, err = cgm.CreateLoadBalancer(ctx, virtualServiceNamePrefix,
		lbPoolNamePrefix, []string{"1.2.3.4", "1.2.3.5"}, portDetailsList, oneArm)
	assert.NoError(t, err, "Load Balancer should be created even on second attempt")
	assert.NotEmpty(t, freeIP, "There should be a non-empty IP returned")

	updatedIps := []string{"5.5.5.5"}
	updatedInternalPort := int32(55555)
	// update IPs and internal port
	err = cgm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, updatedInternalPort, 80)
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	err = cgm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, updatedInternalPort, 443)
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	// update external port only
	updatedExternalPortHttp := int32(8080)
	updatedExternalPortHttps := int32(8443)
	err = cgm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, updatedInternalPort, updatedExternalPortHttp)
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	err = cgm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, updatedInternalPort, updatedExternalPortHttps)
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	// No error on repeated update
	err = cgm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, updatedInternalPort, updatedExternalPortHttp)
	assert.NoError(t, err, "HTTP Load Balancer should be updated")

	err = cgm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"-https", updatedIps, updatedInternalPort, updatedExternalPortHttps)
	assert.NoError(t, err, "HTTPS Load Balancer should be updated")

	err = cgm.DeleteLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, oneArm)
	assert.NoError(t, err, "Load Balancer should be deleted")

	freeIPObtained, err = cgm.GetLoadBalancer(ctx, virtualServiceNameHttp, oneArm)
	assert.NoError(t, err, "Load Balancer should not be found")
	assert.Empty(t, freeIPObtained, "The VIP should not be found")

	freeIPObtained, err = cgm.GetLoadBalancer(ctx, virtualServiceNameHttps, oneArm)
	assert.NoError(t, err, "Load Balancer should not be found")
	assert.Empty(t, freeIPObtained, "The VIP should not be found")

	err = cgm.DeleteLoadBalancer(ctx, virtualServiceNamePrefix, lbPoolNamePrefix, portDetailsList, oneArm)
	assert.NoError(t, err, "Repeated deletion of Load Balancer should not fail")

	err = cgm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-http", virtualServiceNamePrefix+"-http", updatedIps, updatedInternalPort, 80)
	assert.Error(t, err, "updating deleted HTTP Load Balancer should be an error")
	err = cgm.UpdateLoadBalancer(ctx, lbPoolNamePrefix+"-https", virtualServiceNamePrefix+"https", updatedIps, updatedInternalPort, 43)
	assert.Error(t, err, "updating deleted HTTPS Load Balancer should be an error")

	return
}
