# itop_backup

`itop_backup` is a datasource to get get backup link attributes.

## Example Usage

```hcl
data "itop_backup" "rubrik" {
  name   = "rubrik"
}
```

## Argument Reference

* `name` - (Required) The name of the backup link.

## Attribute Reference

* `id` - The id of the backup link.