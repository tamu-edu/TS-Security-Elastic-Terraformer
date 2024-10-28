package elastic_ec

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ElasticECProvider struct {
	terraformutils.Provider
	apiKey string
}

func (p *ElasticECProvider) Init(args []string) error {
	if len(args) < 1 {
		return errors.New("API key required")
	}
	p.apiKey = args[0]
	return nil
}

func (p *ElasticECProvider) GetName() string {
	return "elastic_ec"
}

func (p *ElasticECProvider) InitService(serviceName string, verbose bool) (terraformutils.Service, error) {
	switch serviceName {
	case "ec_deployment":
		return &EcDeploymentGenerator{
			ApiKey: p.apiKey,
		}, nil
	// Add more services as needed
	default:
		return nil, errors.New("unsupported service: " + serviceName)
	}
}

func (p *ElasticECProvider) GetResourceConnections() map[string][]string {
	return map[string][]string{
		// Define connections here if necessary
	}
}
