# Vagrant
Vagrant provisions a virtual machine locally to isolate from OS and make tests

Check Vagrantfile of example vagrant-provision. Everything inside "script" are BASH commands that linux will execute as soon as is started.

Vagrant is perfect to test application environments. It emulates an AWS machine on your PC to make tests quicker

### Commands
<code>VAGRANT UP = starts a new VM with the instructions found on Vagrantfile of that folder</code>

<code>VAGRANT SSH = access the VM through SSH</code>

<code>logout = leaves the VM</code>

<code>VAGRANT HALT = stops the VM</code>

<code>VAGRANT DESTROY = finishes the machine and its HD</code>

### More Vagrant images:

https://app.vagrantup.com/boxes/search

Amazon uses Amazon Linux 2 on production as default