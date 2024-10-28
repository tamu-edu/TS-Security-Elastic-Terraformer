// Copyright 2021 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

import (
	"context"

	"github.com/tamu-edu/TS-Security-Elastic-Terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var ssmAllowEmptyValues = []string{"tags."}

type SsmGenerator struct {
	AWSService
}

func (g *SsmGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ssm.NewFromConfig(config)
	p := ssm.NewDescribeParametersPaginator(svc, &ssm.DescribeParametersInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, parameter := range page.Parameters {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				StringValue(parameter.Name),
				StringValue(parameter.Name),
				"aws_ssm_parameter",
				"aws",
				ssmAllowEmptyValues,
			))
		}
	}

	return nil
}
