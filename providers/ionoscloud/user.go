package ionoscloud

import (
	"context"
	"log"

	"github.com/tamu-edu/TS-Security-Elastic-Terraformer/providers/ionoscloud/helpers"
	"github.com/tamu-edu/TS-Security-Elastic-Terraformer/terraformutils"
)

type UserGenerator struct {
	Service
}

func (g *UserGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	resourceType := "ionoscloud_user"

	usersResponse, _, err := cloudAPIClient.UserManagementApi.UmUsersGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if usersResponse.Items == nil {
		log.Printf("[WARNING] expected a response containing users but received 'nil' instead")
		return nil
	}
	for _, user := range *usersResponse.Items {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*user.Id,
			*user.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
