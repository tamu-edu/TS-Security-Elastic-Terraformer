// Copyright 2019 The Terraformer Authors.
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

package ibm

import (
	"os"

	"github.com/tamu-edu/TS-Security-Elastic-Terraformer/terraformutils"

	"github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
)

// DatabasePostgresqlGenerator ...
type DatabasePostgresqlGenerator struct {
	IBMService
}

// loadPostgresqlDB ...
func (g DatabasePostgresqlGenerator) loadPostgresqlDB(dbID string, dbName string) terraformutils.Resource {
	resource := terraformutils.NewSimpleResource(
		dbID,
		normalizeResourceName(dbName, false),
		"ibm_database",
		"ibm",
		[]string{})

	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^node_count$",
		"^members_memory_allocation_mb$",
		"^node_memory_allocation_mb$",
		"^members_disk_allocation_mb$",
		"^members_cpu_allocation_count$",
		"^node_cpu_allocation_count$",
		"^node_disk_allocation_mb$",
	)
	return resource
}

// InitResources ...
func (g *DatabasePostgresqlGenerator) InitResources() error {
	region := g.Args["region"].(string)
	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
		Region:        region,
	}
	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.New(sess)
	if err != nil {
		return err
	}

	controllerClient, err := controllerv2.New(sess)
	if err != nil {
		return err
	}

	serviceID, err := catalogClient.ResourceCatalog().FindByName("databases-for-postgresql", true)
	if err != nil {
		return err
	}
	query := controllerv2.ServiceInstanceQuery{
		ServiceID: serviceID[0].ID,
	}
	postgreSQLInstances, err := controllerClient.ResourceServiceInstanceV2().ListInstances(query)
	if err != nil {
		return err
	}

	for _, db := range postgreSQLInstances {
		if db.RegionID == region {
			g.Resources = append(g.Resources, g.loadPostgresqlDB(db.ID, db.Name))
		}

	}

	return nil
}
