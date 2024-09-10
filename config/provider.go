/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	"context"
	_ "embed"
	"fmt"

	upconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	"github.com/equinix/terraform-provider-equinix/equinix"
	framework "github.com/equinix/terraform-provider-equinix/equinix/provider"
	"github.com/equinix/terraform-provider-equinix/version"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/provider-jet-equinix/config/metal"
)

const (
	resourcePrefix = "equinix"
	modulePath     = "github.com/crossplane-contrib/provider-jet-equinix"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

var frameworkResources = []string{
	"equinix_metal_connection",
	"equinix_metal_gateway",
	"equinix_metal_organization",
	"equinix_metal_organization_member",
	"equinix_metal_project",
	"equinix_metal_project_ssh_key",
	"equinix_metal_ssh_key",
	"equinix_metal_vlan",
}

var sdkResources = []string{
	"equinix_fabric_cloud_router",
	"equinix_fabric_connection",
	"equinix_fabric_network",
	"equinix_fabric_routing_protocol",
	"equinix_fabric_service_profile",
	"equinix_metal_bgp_session",
	"equinix_metal_device",
	"equinix_metal_device_network_type",
	"equinix_metal_ip_attachment",
	"equinix_metal_port",
	"equinix_metal_port_vlan_attachment",
	"equinix_metal_project_api_key",
	"equinix_metal_reserved_ip_block",
	"equinix_metal_spot_market_request",
	"equinix_metal_user_api_key",
	"equinix_metal_virtual_circuit",
	"equinix_metal_vrf",
	"equinix_network_acl_template",
	"equinix_network_bgp",
	"equinix_network_device",
	"equinix_network_device_link",
	"equinix_network_file",
	"equinix_network_ssh_key",
	"equinix_network_ssh_user",
}

func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, generationProvider bool) (*upconfig.Provider, error) {
	var p *schema.Provider
	var err error

	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p = equinix.Provider()
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	fwProvider := framework.CreateFrameworkProvider(version.ProviderVersion)

	pc := upconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		upconfig.WithDefaultResourceOptions(
			KnownReferencers(),
			IdentifierAssignedByEquinix(),
			SkipOptCompLateInitialization(),
			LongProvision(), // TODO: use this only for Device and other long-provisioning resources
		),
		upconfig.WithShortName(resourcePrefix),
		upconfig.WithRootGroup("equinix.jet.crossplane.io"),
		// WithIncludeList implies CLI based Terraform reconciliation
		upconfig.WithIncludeList(nil),
		upconfig.WithTerraformPluginFrameworkIncludeList(resourceList(frameworkResources)),
		upconfig.WithTerraformPluginSDKIncludeList(sdkResources),
		upconfig.WithReferenceInjectors([]upconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		upconfig.WithFeaturesPackage("internal/features"),
		upconfig.WithTerraformProvider(p),
		upconfig.WithTerraformPluginFrameworkProvider(fwProvider),
		upconfig.WithSchemaTraversers(&upconfig.SingletonListEmbedder{}),
		// upconfig.WithSkipList([]string{".*"}), // helpful when debugging to minimize the number of resources
		// upconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformSDKIncludeList)),
		// upconfig.WithTerraformPluginFrameworkIncludeList(resourceList(terraformPluginFrameworkExternalNameConfigs)),
	)

	for _, configure := range []func(provider *upconfig.Provider){
		// add custom config functions
		metal.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, err
}

// resourceList returns a list of resources regex patterns that are expected to be included in the provider
func resourceList(t []string) []string {
	l := make([]string, len(t))
	for i, n := range t {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = fmt.Sprintf("^%s$", n)
	}
	return l
}
