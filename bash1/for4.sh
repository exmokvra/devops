#!/bin/bash

for file in /etc/*
do
    if [ "$file" == "/etc/resolv.conf" ]; then
        countNameServers=$(grep -c nameserver /etc/resolv.conf)
        echo "Total ${countNameServers} nameservers found in ${file}"
        break
    else
        echo "Just another file: $file"
    fi
done