package elastic_ec

import (
	"errors"
	"net/http"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/auth"
)

type ElasticECProvider struct {
	terraformutils.Provider
	apiKey string
	client *api.API
}

func (p *ElasticECProvider) Init(args []string) error {
	if len(args) < 1 {
		return errors.New("API key required")
	}
	p.apiKey = args[0]
	p.client = api.NewAPI(api.Config{
		Client:     new(http.Client),
		AuthWriter: auth.APIKey(p.apiKey),
	})
	return nil
}

func (p *ElasticECProvider) GetName() string {
	return "elastic_ec"
}

func (p *ElasticECProvider) InitService(serviceName string, verbose bool) (terraformutils.Service, error) {
	switch serviceName {
	case "ec_deployment":
		return &EcDeploymentGenerator{
			Client: p.client,
		}, nil
	default:
		return nil, errors.New("unsupported service: " + serviceName)
	}
}

func (p *ElasticECProvider) GetResourceConnections() map[string][]string {
	return map[string][]string{}
}
