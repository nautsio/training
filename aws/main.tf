resource "aws_instance" "training" {
    count = "${var.count}"
    ami = "${lookup(var.amis, var.region)}"
    instance_type = "m3.large"
    key_name = "${var.key_name}"

    security_groups = [ "${aws_security_group.training.name}" ]
    user_data = "${file("../scripts/cloudinit.sh")}"

    tags {
        Name = "${concat(var.project,"-",count.index)}"
        Project = "${var.project}"
    }
}
resource "aws_security_group" "training" {
    name = "${concat("sg",var.project)}"
    description = "Training traffic"

    ingress {
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
    ingress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        self = true
    }
    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}
