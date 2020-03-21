#!/bin/bash

#. Create a bash script that LIST of ENV_VARS in linux, store to a file called env_data.txt 
#located at /backup/conf/$TODAY/

date=$(date +"%y%m%d")
sudo mkdir -p /backup/conf/$date/
printenv > env_data.txt
sudo mv env_data.txt /backup/conf/$date/