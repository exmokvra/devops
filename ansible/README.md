# Ansible

## Ansible sandbox - Orchestration and automation engine
* Use Ansible just when creating and provisioning the machine. After that, remove Ansible

## How to run Ansible:
<code>ANSIBLE-PLAYBOOK <arquivo.yml> --extra-vars “var1=123 var2=456”</code>

## Command examples for the Ansible .yml file:
#### Tasks:
* Name: “shut down CentOS 6 and Debian 7 systems”
* Command: /sbin/shutdown -t now
* When: (ansible_facts[‘distribution’] == ‘CentOS’ and ansible_facts[‘distribution_major_version’] == “6”) or (ansible_facts[‘distribution’] == ‘Debian’ and ansible_facts[‘distribution_major_version’] == “7”)


* Tasks:


* Command: /bin/false
* Register: result
* Ignore_errors: True


* Command: /bin/something
* When: result is failed


* Command: /bin/something_else
* When: result is succeeded


* Command: /bin/still/something_else
* When: result is skipped


* Tasks:

* Shell: echo “I’ve got ‘{{ foo }}’ and am not afraid to use it!”
* When: foo is defined

* Fail: msg=”Bailing out, this play requires ‘bar’”
* When: bar is undefined
