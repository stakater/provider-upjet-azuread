package privileged

import (
	"github.com/crossplane/upjet/pkg/config"
)

const group = "privileged"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_privileged_access_group_assignment_schedule", func(r *config.Resource) {
		r.References["group_id"] = config.Reference{
			TerraformName: "azuread_group",
		}
		r.References["principal_id"] = config.Reference{
			TerraformName: "azuread_user",
		}

		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})

	p.AddResourceConfigurator("azuread_privileged_access_group_eligibility_schedule", func(r *config.Resource) {
		r.References["group_id"] = config.Reference{
			TerraformName: "azuread_group",
		}
		r.References["principal_id"] = config.Reference{
			TerraformName: "azuread_user",
		}

		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
}
