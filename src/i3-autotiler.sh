#!/bin/bash

i3-msg -t subscribe -m '[ "window" ]' | while read -r line; do
    type=$(jq '.change' <<< "$line")
    width=$(jq '.container.window_rect.width' <<< "$line")
    height=$(jq '.container.window_rect.height' <<< "$line")

    if [[ $type == '"focus"' ]]; then
        if [[ $width -gt $height ]]; then
            if [[ $1 == "--debug" ]]; then
                echo "Width $width ,Height $height ,Next split will be horizontal"
            fi
            i3-msg split horizontal &>/dev/null
        else
            if [[ $1 == "--debug" ]]; then
                echo "Width $width ,Height $height ,Next split will be vertical"
            fi
            i3-msg split vertical &>/dev/null
        fi
    fi
done
