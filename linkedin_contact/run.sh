#!/usr/bin/env sh

python linkedin_contact_to_momentjs.py > part2.html
cat part1.html part2.html part3.html > `date +'%Y-%m-%d-%Hh%Mm%S'`.html
