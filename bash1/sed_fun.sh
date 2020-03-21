#!/bin/bash

echo "Welcome to the DevOps class" | sed 's/class/team/'

echo "one
two
three" > x.txt
sed '2d' x.txt
rm -rf x.txt