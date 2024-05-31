#!/bin/bash

# Function to set the keyboard layout and compose key
set_layout() {
    if [ "$1" == "us" ]; then
        setxkbmap us -option compose:ralt
    elif [ "$1" == "it" ]; then
        setxkbmap it -option
    fi
}

# Check the argument passed to the script
if [ "$#" -eq 1 ]; then
    set_layout $1
else
    echo "Usage: $0 {us|it}"
    exit 1
fi
