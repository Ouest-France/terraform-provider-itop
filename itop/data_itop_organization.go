package itop

import (
	"github.com/Ouest-France/goitop"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOrganization() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOrganizationRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization name",
			},
		},
	}
}

func dataSourceOrganizationRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	id, err := client.GetOrganization(d.Get("name").(string))
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}
