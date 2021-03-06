package itop

import (
	"github.com/Ouest-France/goitop"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBackup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBackupRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Backup link name",
			},
		},
	}
}

func dataSourceBackupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	id, err := client.GetBackup(d.Get("name").(string))
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}
