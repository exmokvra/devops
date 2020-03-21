#!/bin/bash

sa="Bash"
sb="Code"

if [ $sa = $sb ]; then
    echo "Strings are equals"
else
    echo "Strings are different"
fi

if [ $sa = $sa ]; then
    echo "Strings are equals"
else   
    echo "Strings are different"
fi