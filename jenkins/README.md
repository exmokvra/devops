# Jenkins

This repository was used to provision a jenkins on localhost docker.

Jenkins' pipeline uses the bake.Jenkinsfile to run

First step builds a new AMI using Packer

Second step launches a new instance on AWS using the just created AMI with Terraform

The AWS instance will have the port :8080/calc/ service running on top of Docker with Docker Compose
