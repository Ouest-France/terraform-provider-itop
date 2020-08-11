# itop_organization

`itop_organization` is a datasource to get organization attributes.

## Example Usage

```hcl
data "itop_organization" "org" {
  name   = "MyOrg"
}
```

## Argument Reference

* `name` - (Required) The name of the organization.

## Attribute Reference

* `id` - The id of the organization.