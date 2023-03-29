package testingsdk

import (
	"context"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	swagger "github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdswaggerclient"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

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

func getRdeById(ctx context.Context, client *vcdsdk.Client, rdeId string) (*swagger.DefinedEntity, error) {
	clusterOrg, err := client.VCDClient.GetOrgByName(client.ClusterOrgName)
	if err != nil {
		return nil, fmt.Errorf("error retrieving org [%s]: [%v]", client.ClusterOrgName, err)
	}
	if clusterOrg == nil || clusterOrg.Org == nil {
		return nil, fmt.Errorf("retrieved org is nil for [%s]", client.ClusterOrgName)
	}
	rde, _, _, err := client.APIClient.DefinedEntityApi.GetDefinedEntity(ctx, rdeId, clusterOrg.Org.ID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving RDE [%s]: [%v]", rdeId, err)
	}
	return &rde, nil
}

func GetVappTemplates(client *vcdsdk.Client) ([]*types.QueryResultVappTemplateType, error) {
	var allVappTemplates []*types.QueryResultVappTemplateType

	clusterOrg, err := client.VCDClient.GetOrgByName(client.ClusterOrgName)
	if err != nil {
		return nil, fmt.Errorf("error retrieving org [%s]: [%v]", client.ClusterOrgName, err)
	}
	catalogRecords, err := clusterOrg.QueryCatalogList()
	if err != nil {
		return nil, fmt.Errorf("error retrieving catalog records: [%v]", err)
	}
	for _, catalogRecord := range catalogRecords {
		catalog, err := clusterOrg.GetCatalogByName(catalogRecord.Name, true)
		if err != nil {
			return nil, fmt.Errorf("error retreiving catalog %s: [%v]", catalogRecord.Name, err)
		}

		vappTemplates, err := catalog.QueryVappTemplateList()
		if err != nil {
			return nil, fmt.Errorf("error retreiving vApp templates from catalog %s: [%v]", catalogRecord.Name, err)
		}

		for _, vappTemplate := range vappTemplates {
			allVappTemplates = append(allVappTemplates, vappTemplate)
		}
	}
	return allVappTemplates, nil
}
