// Step 1: terraform init
// Step 2: terraform plan -var aws_region=us-east-1 -var aws_access_key=<access_key> -var aws_secret_key=<secret_key> -out tfout.log
// Step 3: terraform apply tfout.log

provider "aws" {
    region = var.aws_region
    access_key = var.aws_access_key
    secret_key = var.aws_secret_key
}

variable "aws_region" {
    description = "AWS Region to launch the infrastructure"
}

variable "aws_access_key" {
    description = "AWS Access Key"
}

variable "aws_secret_key" {
    description = "AWS Secret KEy"
}

// Check which is the most recent AMI image to use
data "aws_ami" "image" {
  most_recent = true
  owners = ["self"]
  filter {
    name = "tag:Name"
    values = ["GoMicroWithETCD"]
  }
}

// Creates the infrastructure
resource "aws_instance" "app" {
    count = 1
    ami = data.aws_ami.image.id
    instance_type = "t2.micro"
    user_data = file("run_app.sh")
    instance_state = ["running"]

    tags = {
        Name = "GoMicroserviceWithETCD"
    }
}