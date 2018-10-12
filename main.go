package main

import (
	"Ouest-France/terraform-provider-itop/itop"

	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return itop.Provider()
		},
	})
}
