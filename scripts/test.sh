#!/bin/bash
set -o errexit
set -o nounset

go test -v -race -cover
