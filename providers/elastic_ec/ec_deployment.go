package elastic_ec

import (
	"github.com/tamu-edu/TS-Security-Elastic-Terraformer/terraformutils"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi"
)

type EcDeploymentGenerator struct {
	terraformutils.Service
	Client *api.API
}

func (g *EcDeploymentGenerator) InitResources() error {
	res, err := deploymentapi.List(deploymentapi.ListParams{API: g.Client})
	if err != nil {
		return err
	}

	for _, deployment := range res.Deployments {
		resource := terraformutils.NewSimpleResource(
			deployment.ID,
			deployment.Name,
			"ec_deployment",
			"elastic_ec",
			[]string{},
		)
		g.Resources = append(g.Resources, resource)
	}
	return nil
}
