#!/bin/bash
export GOPATH=$PWD
cd src/config-writer/ && \
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $GOPATH/bin/host-local-tools .