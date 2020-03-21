#!/bin/bash

sudo apt-get install zip unzip -y

date=$(date +"%y%m%d")

sudo mkdir -p -v /backup/data/$date

zip -r zip_file.zip /etc/
sudo mv zip_file.zip /backup/data/$date

#remove the folder after creation
sudo rm -rf /backup/data/$date