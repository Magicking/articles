#!/bin/sh

#uuoc
cat "$1" | awk '{ print "go run gen_business.go --key "$2" > bc/business-"$1".svg; convert -density 300 bc/business-"$1".svg bc/business-"$1".png" }' | xargs -P4 -L1 -I{} /bin/sh -c '{}'
