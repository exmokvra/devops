#!/bin/bash

printf "Whats your favorite American Footbal team?"
read ANSWER

case $ANSWER in
    ravens)
        echo "Congrats!"
        ;;
    patriots|bucaneers)
        echo "So sad"
        ;;
    *)
        echo "Great! Let's go for soccer"
        ;;
esac