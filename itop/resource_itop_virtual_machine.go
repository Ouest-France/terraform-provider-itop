package itop

import (
	"github.com/Ouest-France/goitop"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"exploitationservice_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "0",
				Description: "Exploitation Service ID",
			},
			"backup": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Backup state",
			},
			"backup_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "0",
				Description: "Backup ID",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description",
			},
		},
	}
}

func resourceVirtualMachineCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	var backup string
	if d.Get("backup").(bool) {
		backup = "yes"
	} else {
		backup = "no"
	}

	id, err := client.CreateVM(
		d.Get("name").(string),
		d.Get("org_id").(string),
		d.Get("env_id").(string),
		d.Get("cluster_id").(string),
		d.Get("exploitationservice_id").(string),
		backup,
		d.Get("backup_id").(string),
		d.Get("description").(string),
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

	fields := map[string]interface{}{
		"name":                   vm.Name,
		"org_id":                 vm.OrgID,
		"env_id":                 vm.EnvID,
		"cluster_id":             vm.ClusterID,
		"exploitationservice_id": vm.ExploitationServiceID,
		"backup_id":              vm.BackupID,
		"description":            vm.Description,
	}

	if vm.Backup == "yes" {
		fields["backup"] = true
	} else {
		fields["backup"] = false
	}

	for field, value := range fields {
		err = d.Set(field, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceVirtualMachineUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	var backup string
	if d.Get("backup").(bool) {
		backup = "yes"
	} else {
		backup = "no"
	}

	err := client.UpdateVM(
		d.Id(),
		d.Get("name").(string),
		d.Get("org_id").(string),
		d.Get("env_id").(string),
		d.Get("cluster_id").(string),
		d.Get("exploitationservice_id").(string),
		backup,
		d.Get("backup_id").(string),
		d.Get("description").(string),
	)
	return err
}

func resourceVirtualMachineDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*goitop.Client)

	return client.DeleteVM(d.Id())
}
