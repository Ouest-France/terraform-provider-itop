package itop

import (
	"github.com/Ouest-France/goitop"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEnvironmentRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Environment name",
			},
		},
	}
}

func dataSourceEnvironmentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	id, err := client.GetEnvironment(d.Get("name").(string))
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}
