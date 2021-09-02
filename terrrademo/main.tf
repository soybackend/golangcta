provider "aws" {
    region = "us-east-1"
}

resource "aws_instance" "web" {
    ami = "ami-09e67e426f25ce0d7"
    instance_type = "t2.micro"
    security_groups = [aws_security_group.web_traffic.name]
    user_data = "${file("server-script.sh")}"
    key_name = "ctademo"
    tags = {
        Name = "Web server"
    }
}

resource "aws_eip" "web_ip" {
    instance = aws_instance.web.id
}

variable "ingress_ports" {
    type = list(number)
    default = [80, 22]
}

variable "egress_ports" {
    type = list(number)
    default = [80, 22]
}

resource "aws_security_group" "web_traffic" {
    name = "Allow Web Traffic"

    dynamic "ingress" {
        iterator = port
        for_each = var.ingress_ports
        content {
            from_port = port.value
            to_port = port.value
            protocol = "TCP"
            cidr_blocks = ["0.0.0.0/0"]
        }        
    }

    dynamic "egress" {
        iterator = port
        for_each = var.egress_ports
        content {
            from_port = port.value
            to_port = port.value
            protocol = "TCP"
            cidr_blocks = ["0.0.0.0/0"]
        }        
    }
}

output "PublicIP" {
    value = aws_eip.web_ip.public_ip
}