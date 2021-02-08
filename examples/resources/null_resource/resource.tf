resource "aws_instance" "cluster" {
  count = 3

  # ...
}

# The primary use-case for the null resource is as a do-nothing container for
# arbitrary actions taken by a provisioner.
#
# In this example, three EC2 instances are created and then a null_resource instance
# is used to gather data about all three and execute a single action that affects
# them all. Due to the triggers map, the null_resource will be replaced each time
# the instance ids change, and thus the remote-exec provisioner will be re-run.
resource "null_resource" "cluster" {
  # Changes to any instance of the cluster requires re-provisioning
  triggers = {
    cluster_instance_ids = join(",", aws_instance.cluster.*.id)
  }

  # Bootstrap script can run on any instance of the cluster
  # So we just choose the first in this case
  connection {
    host = element(aws_instance.cluster.*.public_ip, 0)
  }

  provisioner "remote-exec" {
    # Bootstrap script called with private_ip of each node in the clutser
    inline = [
      "bootstrap-cluster.sh ${join(" ", aws_instance.cluster.*.private_ip)}",
    ]
  }
}
