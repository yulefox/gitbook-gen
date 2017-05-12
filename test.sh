#!/bin/sh
set -e

export LAMP_CONFIG_PATH=$PWD

echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor) ; do
    go test -v -race -coverprofile=profile.out -covermode=atomic -timeout 90s $d
    if [ -f profile.out ] ; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
