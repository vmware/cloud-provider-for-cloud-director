package vcdsdk

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestComputePolicy(t *testing.T) {

	authFile := filepath.Join(gitRoot, "testdata/auth_test.yaml")
	authFileContent, err := ioutil.ReadFile(authFile)
	assert.NoError(t, err, "There should be no error reading the auth file contents.")

	var authDetails authorizationDetails
	err = yaml.Unmarshal(authFileContent, &authDetails)
	assert.NoError(t, err, "There should be no error parsing auth file content.")

	// get client
	cloudConfig, err := getTestConfig()
	assert.NoError(t, err, "There should be no error opening and parsing cloud config file contents.")

	vcdClient, err := GetTestVCDClient(
		cloudConfig,
		map[string]interface{}{
			"getVdcClient": true,
		})
	assert.NoError(t, err, "Unable to get VCD client")
	require.NotNil(t, vcdClient, "VCD Client should not be nil")

	computePolicyName := "4core4gb"
	orgManager, err := NewOrgManager(vcdClient, vcdClient.ClusterOrgName)
	assert.NoError(t, err, "there should be no error when creating orgManager object")

	computePolicy, err := orgManager.GetComputePolicyDetailsFromName(computePolicyName)
	assert.NoError(t, err, "should be able to get compute policy [%s]", computePolicyName)
	assert.NotNil(t, computePolicy, "should get compute policy for [%s]", computePolicyName)

	computePolicyName = "some-random-policy-name"
	computePolicy, err = orgManager.GetComputePolicyDetailsFromName(computePolicyName)
	assert.Error(t, err, "should NOT be able to get random compute policy [%s]", computePolicyName)
	assert.Nil(t, computePolicy, "should NOT get random compute policy for [%s]", computePolicyName)

	computePolicyName = "cse----native"
	computePolicy, err = orgManager.GetComputePolicyDetailsFromName(computePolicyName)
	assert.NoError(t, err, "should be able to get cse native compute policy [%s]", computePolicyName)
	assert.NotNil(t, computePolicy, "should get compute policy for [%s]", computePolicyName)

	return
}
