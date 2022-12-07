package testingsdk

import "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"

func getTestVCDClient(params *VCDAuthParams) (*vcdsdk.Client, error) {
	return vcdsdk.NewVCDClientFromSecrets(
		params.Host,
		params.OrgName,
		params.OvdcName,
		params.UserOrg,
		params.Username,
		"",
		params.RefreshToken,
		true,
		params.GetVdcClient)
}
