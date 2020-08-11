# iTop Provider

The iTop provider is used to interact with the resources supported by [iTop](https://www.combodo.com/itop) API.

## Example Usage

```hcl
# Provider configuration
terraform {
  required_providers {
    itop = {
      source  = "Ouest-France/itop"
    }
  }
}

provider "itop" {
  address  = "https://itop.mydomain.tld"
  user     = "itop_user"
  password = "itop_password"
}

...
```

## Argument Reference

* `address` - (Required) This is the iTop address formatted like `https://itop.mydomain.tld`.

* `user` - (Required) This is the iTop user to access the API.

* `password` - (Required) This is the iTop password to access the API.

## Requirements

* iTop >= 2.0.3