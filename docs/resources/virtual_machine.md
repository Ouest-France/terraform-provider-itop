# itop_virtual_machine

`itop_virtual_machine` is a resource for managing a virtual machine.

## Example Usage

```hcl
data "itop_organization" "org" {
  name   = "MyOrg"
}

data "itop_environment" "env" {
  name   = "Production"
}

resource "itop_virtual_machine" "vm" {
  name   = "terraform_vm_test"
  org_id = data.itop_organization.org.id
  env_id = data.itop_environment.env.id
}
```

## Argument Reference

* `name` - (Required) The name of the virtual machine.
* `org_id` - (Required) The organization id of the virtual machine.
* `env_id` - (Required) The environment id of the virtual machine.
* `cluster_id` - (Optional) The cluster id of the virtual machine. Defaults to `0`.
* `exploitationservice_id` - (Optional) The exploitation service id of the virtual machine. Defaults to `0`.
* `backup` - (Optional) Backup enable status of the virtual machine. Defaults to `false`.
* `backup_id` - (Optional) The backup link id of the virtual machine. Defaults to `0`.

## Attribute Reference

* `id` - The id of the virtual machine.