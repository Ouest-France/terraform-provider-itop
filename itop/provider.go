package itop

import (
	"github.com/Ouest-France/goitop"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "itop server address",
			},
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "itop username",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "itop password",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"itop_virtual_machine": resourceVirtualMachine(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"itop_environment":         dataSourceEnvironment(),
			"itop_organization":        dataSourceOrganization(),
			"itop_cluster":             dataSourceCluster(),
			"itop_exploitationservice": dataSourceExploitationService(),
			"itop_backup":              dataSourceBackup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return goitop.NewClient(
		d.Get("address").(string),
		d.Get("user").(string),
		d.Get("password").(string),
	), nil
}
