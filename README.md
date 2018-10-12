# `terraform-provider-itop`

A [Terraform][1] plugin for managing [iTop][2].

## Contents

* [Installation](#installation)
* [`itop` Provider](#provider-configuration)
* [DataSources](#datasources)
  * [`itop_environment`](#itop_environment)
  * [`itop_organization`](#itop_organization)
  * [`itop_cluster`](#itop_cluster)
* [Resources](#resources)
  * [`itop_virtual_machine`](#itop_virtual_machine)
* [Requirements](#requirements)

## Installation

Download and extract the [latest
release](https://github.com/Ouest-France/terraform-provider-itop/releases/latest) to
your [terraform plugin directory][third-party-plugins] (typically `~/.terraform.d/plugins/`)

## Provider Configuration

### Example

Example provider.
```hcl
provider "itop" {
  address  = "https://itop.mydomain.tld"
  user     = "itop_user"
  password = "itop_password"
}
```

| Property            | Description                | Default    |
| ----------------    | -----------------------    | ---------- |
| `address`           | itop server address        | `Required` |
| `user`              | itop username              | `Required` |
| `password`          | itop password              | `Required` |


## DataSources
### `itop_environment`

A datasource to get environment attributes.

#### Example

```hcl
provider "itop" {
  address  = "https://itop.mydomain.tld"
  user     = "itop_user"
  password = "itop_password"
}

data "itop_environment" "env" {
  name   = "Production"
}
```

#### Arguments

| Property             | Description                                    |
| ----------------     | -----------------------                        |
| `name`               | The name of the environment                    |

#### Attributes

| Property             | Description                                    |
| ----------------     | -----------------------                        |
| `id`                 | The id of the environment                      |


### `itop_organization`

A datasource to get organization attributes.

#### Example

```hcl
provider "itop" {
  address  = "https://itop.mydomain.tld"
  user     = "itop_user"
  password = "itop_password"
}

data "itop_organization" "org" {
  name   = "MyOrg"
}
```


### `itop_cluster`

A datasource to get cluster attributes.

#### Example

```hcl
provider "itop" {
  address  = "https://itop.mydomain.tld"
  user     = "itop_user"
  password = "itop_password"
}

data "itop_cluster" "prod" {
  name   = "Production"
}
```


#### Arguments

| Property             | Description                                    |
| ----------------     | -----------------------                        |
| `name`               | The name of the environment                    |

#### Attributes

| Property             | Description                                    |
| ----------------     | -----------------------                        |
| `id`                 | The id of the environment                      |

#### Arguments

| Property             | Description                                    |
| ----------------     | -----------------------                        |
| `name`               | The name of the organization                   |

#### Attributes

| Property             | Description                                    |
| ----------------     | -----------------------                        |
| `id`                 | The id of the organization                     |


## Resources
### `itop_virtual_machine`

A resource for managing virtual machine.

#### Example

```hcl
provider "itop" {
  address  = "https://itop.mydomain.tld"
  user     = "itop_user"
  password = "itop_password"
}

data "itop_organization" "org" {
  name   = "MyOrg"
}

data "itop_environment" "env" {
  name   = "Production"
}

resource "itop_virtual_machine" "vm" {
  name   = "terraform_vm_test"
  org_id = "${data.itop_organization.org.id}"
  env_id = "${data.itop_environment.env.id}"
}
```

#### Arguments

| Property             | Description                                    | Default    |
| ----------------     | -----------------------                        | `Required` |
| `name`               | The name of the virtual machine                | `Required` |
| `org_id`             | The organization id of the virtual machine     | `Required` |
| `env_id`             | The environment id of the virtual machine      | `Required` |
| `cluster_id`         | The cluster id of the virtual machine          | `0`        |

#### Attributes

| Property             | Description                                    |
| ----------------     | -----------------------                        |
| `id`                 | The id of the virtual machine                  |


## Requirements
* iTop >= 2.0.3

[1]: https://www.terraform.io
[2]: https://www.combodo.com/itop
