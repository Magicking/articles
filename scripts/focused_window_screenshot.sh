#!/bin/sh

screenshot_dir=~/screenshots
d=$(date +'%Y-%m-%d-%H%M%S')
name=$(swaymsg -t get_outputs | jq 'map(select(.focused == true)) | first  | .name' | tr -d '"')

grim -t png -c -o "${name}" "${screenshot_dir}/${d}.png"
