package business

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_business_service", func(r *config.Resource) {

		r.ShortGroup = "business"
		r.References = config.References{
			"team": {
				TerraformName: "pagerduty_team",
			},
		}
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromId
		r.ExternalName.GetIDFn = c.GetFakeID
		// Deprecated
		if s, ok := r.TerraformResource.Schema["type"]; ok {
			s.Optional = false
			s.Computed = true
		}
	})

	p.AddResourceConfigurator("pagerduty_business_service_subscriber", func(r *config.Resource) {

		r.ShortGroup = "business"
		r.References = config.References{
			"business_service_id": {
				TerraformName: "pagerduty_business_service",
			},
		}

	})
}
