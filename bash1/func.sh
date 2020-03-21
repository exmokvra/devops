#!/bin/bash

function hello {
    echo "hello $1"
}

function quit {
    echo "bye"
    exit 0
}

printf "Type in something"
read WORD

hello $WORD
quit
echo "This won't be printed"