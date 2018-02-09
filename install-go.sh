#!/usr/bin/env bash

set -e
set -u

curl -o /app/go.tgz https://storage.googleapis.com/golang/go1.9.4.linux-amd64.tar.gz
tar xvf /app/go.tgz -C /app
export GOPATH=/app/go PATH=/app/go/bin:$PATH
go get -t -d -v ./...
