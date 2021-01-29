---
page_title: "Null Data Source"
description: |-
  A data source that does nothing.
---

# Null Resource

The `null_data_source` data source implements the standard data source
lifecycle but does not interact with any external APIs.

## Example Usage

The primary use-case for the null data source is to gather together
collections of intermediate values to re-use elsewhere in configuration:

```hcl
data "null_data_source" "values" {
  inputs = {
    all_server_ids = "${concat(aws_instance.green.*.id, aws_instance.blue.*.id)}"
    all_server_ips = "${concat(aws_instance.green.*.private_ip, aws_instance.blue.*.private_ip)}"
  }
}

resource "aws_elb" "main" {
  # ...

  instances = "${data.null_data_source.values.outputs["all_server_ids"]}"
}

output "all_server_ids" {
  value = "${data.null_data_source.values.outputs["all_server_ids"]}"
}

output "all_server_ips" {
  value = "${data.null_data_source.values.outputs["all_server_ips"]}"
}

# ... (other uses of the values) ...
```

## Argument Reference

The following arguments are supported:

* `inputs` - (Optional) A map of arbitrary strings that is copied into the
  `outputs` attribute, and accessible directly for interpolation.

* `has_computed_default` - (Optional) If set, its literal value will be
  stored and returned. If not, its value defaults to `"default"`. This
  argument exists primarily for testing and has little practical use.

## Attributes Reference

The following attributes are exported:

* `outputs` - After the data source is "read", a copy of the `inputs` map.

* `random` - A random value. This is primarily for testing and has little
  practical use; prefer
  [the `random` provider](https://registry.terraform.io/providers/hashicorp/random)
  for more practical random number use-cases.
