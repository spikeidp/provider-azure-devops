/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/spikeidp/provider-azure-devops/config/gitrepository"
	"github.com/spikeidp/provider-azure-devops/config/project"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "azure-devops"
	modulePath     = "github.com/spikeidp/provider-azure-devops"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("spikeidp.cit.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		project.Configure,
		gitrepository.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
