package itop

import (
	"github.com/Ouest-France/goitop"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceClusterRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster name",
			},
		},
	}
}

func dataSourceClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	id, err := client.GetCluster(d.Get("name").(string))
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}
