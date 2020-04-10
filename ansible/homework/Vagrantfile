# -*- mode: ruby -*-
# vi set ft=ruby :

Vagrant.configure(2) do |config|

    config.vm.provider "virtualbox"
    config.vm.provider "virtualbox" do |v|
        v.memory = 1024
        v.cpus = 2
    end

    config.vm.box = "ubuntu/xenial64"
    config.vm.network "private_network", ip: "55.55.55.150"
    config.vm.synced_folder ".", "/home/vagrant/shared"
    config.vm.provision "shell", inline: <<-SHELL

        sudo apt-get update -y
        sudo apt-get install -y wget curl vim git unzip

        # Install PIP
        sudo apt-get install -y python3
        sudo apt-get install -y python3-pip
        sudo pip3 install --upgrade pip
        sudo pip3 install packaging

        # Install Ansible
        sudo apt-get -y install python-jinja2 python-paramiko
        sudo sh -c 'touch /home/vagrant/ansible_hosts'
        sudo sh -c 'echo "[localhost]" > /home/vagrant/ansible_hosts'
        sudo sh -c 'echo "localhost ansible_connection=local" >> /home/vagrant/ansible_hosts'
        sudo sh -c 'echo "export ANSIBLE_INVENTORY=~/ansible_hosts" >> /etc/profile'
        sudo pip3 install ansible

        # Clone and start the app
        sudo git clone https://github.com/guisesterheim/devops/
        # Add command to download repo here
        # Add command to run playbook here
        # Add command to run tests
    SHELL
end