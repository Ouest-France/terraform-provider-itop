# itop_cluster

`itop_cluster` is a datasource to get get cluster attributes.

## Example Usage

```hcl
data "itop_cluster" "prod" {
  name   = "Production"
}
```

## Argument Reference

* `name` - (Required) The name of the cluster.

## Attribute Reference

* `id` - The id of the cluster.