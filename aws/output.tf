output "externalip" {
  value = "${join(",", aws_instance.training.*.public_ip)}"
}
