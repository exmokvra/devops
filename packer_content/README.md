# Packer

## Classic flow for provisioning:
1 - Terraform creates the infra when needed
2 - Jenkins job starts
3 - Packer builds and creates AMI image
4 - Ansible provision

## Other commends on Packer:
* Amazon AMI = Amazon service to store images
* AMI base for Ubuntu: ami-0c435d654482161c5