#!/bin/bash
# usage: 
# kdb_brightness.sh increase N
# kdb_brightness.sh decrease N

adjust_brightness() {
    if [ "$1" == "increase" ]; then
        sudo light -s sysfs/leds/smc::kbd_backlight -A "$2"
    elif [ "$1" == "decrease" ]; then
        sudo light -s sysfs/leds/smc::kbd_backlight -U "$2"
    else
        echo "Invalid action. Please use 'increase' or 'decrease'."
    fi
}

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 [increase|decrease] [value]" >&2
    exit 1
fi

action=$1
value=$2

adjust_brightness "$action" "$value"
