#!/bin/bash
## DeGOps: 0.0.4
set -o errexit
set -o nounset

go build -o $GOBIN/zipf ./cmd/zipf

$GOBIN/zipf -limit=100 -path=$1
