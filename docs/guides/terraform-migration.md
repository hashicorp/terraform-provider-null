---
page_title: "terraform_data Migration Guide"
description: |-
  Migration guide for moving `null_resource` resources to `terraform_data`
---

# terraform_data Migration Guide

Terraform 1.4 introduced the [`terraform_data` managed resource](https://developer.hashicorp.com/terraform/language/resources/terraform-data) as a built-in replacement for the `null_resource` managed resource.

The built-in `terraform_data` managed resource is designed similar to the `null_resource` managed resource with added benefits:

* The `hashicorp/null` provider is no longer required to be downloaded and installed.
* Resource replacement trigger configuration supports any value type.
* Optional data storage.

Use `terraform_data` instead of `null_resource` with all new Terraform configurations running on Terraform 1.4 and later.

## Migrating existing configurations

-> Migrating from the `null_resource` managed resource to the `terraform_data` managed resource with the `moved` configuration block requires Terraform 1.9 and later.

### Without triggers

Use the [`moved` configuration block](https://developer.hashicorp.com/terraform/language/moved) to migrate from `null_resource` to `terraform_data`.

Given this example configuration with a `null_resource` managed resource:

```terraform
resource "null_resource" "example" {
  provisioner "local-exec" {
    command = "echo 'Hello, World!'"
  }
}
```

Replace this configuration with a `terraform_data` managed resource and `moved` configuration:

```terraform
resource "terraform_data" "example" {
  provisioner "local-exec" {
    command = "echo 'Hello, World!'"
  }
}

moved {
  from = null_resource.example
  to   = terraform_data.example
}
```

Run a plan operation, such as `terraform plan`, to verify that the move will occur as expected.

Example output with no changes:

```console
$ terraform plan
terraform_data.example: Refreshing state... [id=892002337455008838]

Terraform will perform the following actions:

  # null_resource.example has moved to terraform_data.example
    resource "terraform_data" "example" {
        id = "892002337455008838"
    }

Plan: 0 to add, 0 to change, 0 to destroy.
```

Run an apply operation, such as `terraform apply`, to move the resource and complete the migration. Remove the `moved` configuration block at any time afterwards.

### With triggers

Use the [`moved` configuration block](https://developer.hashicorp.com/terraform/language/moved) to migrate from `null_resource` to `terraform_data`. Replace the `null_resource` managed resource `triggers` argument with the `terraform_data` managed resource `triggers_replace` argument when moving.

Given this example configuration with a `null_resource` managed resource that includes the `triggers` argument:

```terraform
resource "null_resource" "example" {
  triggers = {
    examplekey = "examplevalue"
  }

  provisioner "local-exec" {
    command = "echo 'Hello, World!'"
  }
}
```

Replace this configuration with the following `terraform_data` managed resource and `moved` configuration:

```terraform
resource "terraform_data" "example" {
  triggers_replace = {
    examplekey = "examplevalue"
  }

  provisioner "local-exec" {
    command = "echo 'Hello, World!'"
  }
}

moved {
  from = null_resource.example
  to   = terraform_data.example
}
```

Run a plan operation, such as `terraform plan`, to verify that the move will occur as expected.

Example output with no changes:

```console
$ terraform plan
terraform_data.example: Refreshing state... [id=1651348367769440250]

Terraform will perform the following actions:

  # null_resource.example has moved to terraform_data.example
    resource "terraform_data" "example" {
        id               = "1651348367769440250"
        # (1 unchanged attribute hidden)
    }

Plan: 0 to add, 0 to change, 0 to destroy.
```

Run an apply operation, such as `terraform apply`, to move the resource and complete the migration. Remove the `moved` configuration block at any time afterwards.
