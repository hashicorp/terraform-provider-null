resource "aws_instance" "green" {
  count         = 3
  ami           = "ami-0dcc1e21636832c5d"
  instance_type = "m5.large"

  # ...
}

resource "aws_instance" "blue" {
  count         = 3
  ami           = "ami-0dcc1e21636832c5d"
  instance_type = "m5.large"

  # ...
}

data "null_data_source" "values" {
  inputs = {
    all_server_ids = concat(
      aws_instance.green[*].id,
      aws_instance.blue[*].id,
    )
    all_server_ips = concat(
      aws_instance.green[*].private_ip,
      aws_instance.blue[*].private_ip,
    )
  }
}

resource "aws_elb" "main" {
  instances = data.null_data_source.values.outputs["all_server_ids"]

  # ...
  listener {
    instance_port     = 8000
    instance_protocol = "http"
    lb_port           = 80
    lb_protocol       = "http"
  }
}

output "all_server_ids" {
  value = data.null_data_source.values.outputs["all_server_ids"]
}

output "all_server_ips" {
  value = data.null_data_source.values.outputs["all_server_ips"]
}
