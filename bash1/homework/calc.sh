#!/bin/bash

function add {
    echo "Result of the sum is $(echo $1+$2 | bc)"
}

function subtract {
    echo "Result of the subtract is $(echo $1-$2 | bc)"
}

function multiply {
    echo "Result of the multiplication is $(echo $1*$2 | bc)"
}

function divide {
    echo "Result of the division is $(echo $1/$2 | bc)"
}

printf "Type in your operation"
read OPER

printf "Type in the first number"
read FIRST_NUM

printf "Type in the second number"
read SECOND_NUM

case $OPER in
    add)
        add $FIRST_NUM $SECOND_NUM
        ;;
    subtract)
        subtract $FIRST_NUM $SECOND_NUM
        ;;
    multiply)
        multiply $FIRST_NUM $SECOND_NUM
        ;;
    divide)
        divide $FIRST_NUM $SECOND_NUM
        ;;
    *)
        echo "No valid operation was entered. You can use either add, subtract, multiply and divide"
        ;;
esac