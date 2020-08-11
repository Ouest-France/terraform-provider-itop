# itop_environment

`itop_environment` is a datasource to get environment attributes.

## Example Usage

```hcl
data "itop_environment" "env" {
  name   = "Production"
}
```

## Argument Reference

* `name` - (Required) The name of the environment.

## Attribute Reference

* `id` - The id of the environment.