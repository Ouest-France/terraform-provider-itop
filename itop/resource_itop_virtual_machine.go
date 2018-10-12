package itop

import (
	"github.com/Ouest-France/goitop"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualMachineCreate,
		Read:   resourceVirtualMachineRead,
		Update: resourceVirtualMachineUpdate,
		Delete: resourceVirtualMachineDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Virtual Machine name",
			},
			"org_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization ID",
			},
			"env_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Environment ID",
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "0",
				Description: "Cluster ID",
			},
		},
	}
}

func resourceVirtualMachineCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	id, err := client.CreateVM(
		d.Get("name").(string),
		d.Get("org_id").(string),
		d.Get("env_id").(string),
		d.Get("cluster_id").(string),
	)
	if err != nil {
		return err
	}

	d.SetId(id)
	return nil
}

func resourceVirtualMachineRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	vm, err := client.GetVM(d.Id())
	if err != nil {
		return err
	}

	d.SetId(vm.ID)

	d.Set("name", vm.Name)
	d.Set("org_id", vm.OrgID)
	d.Set("env_id", vm.EnvID)
	d.Set("cluster_id", vm.ClusterID)

	return nil
}

func resourceVirtualMachineUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	err := client.UpdateVM(
		d.Id(),
		d.Get("name").(string),
		d.Get("org_id").(string),
		d.Get("env_id").(string),
		d.Get("cluster_id").(string),
	)
	return err
}

func resourceVirtualMachineDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	return client.DeleteVM(d.Id())
}
