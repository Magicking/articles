#!/bin/sh

#uuoc
cat "$1" | awk '{ print $2 }' | xargs -P4 -L1 -I{} /bin/sh -c 'go run gen_business.go --key {} > "bc/business-{}".svg; convert -density 300 bc/business-{}.svg bc/business-{}.png'
