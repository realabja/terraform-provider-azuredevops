package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider - The top level Azure DevOps Provider definition.
func Provider() *schema.Provider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"personal_access_token": {
				Type:        schema.TypeString,
				Required: 	 true,	
				DefaultFunc: schema.EnvDefaultFunc("AZDO_PERSONAL_ACCESS_TOKEN", nil),
				Description: "The personal access token which should be used.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{

			// Why is the key/value named the way they are?
			"azuredevops_foo": resourceFoo(),
		},
	}
}