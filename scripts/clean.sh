#!/bin/bash
set -o errexit
set -o nounset

rm -rf vendor
touch coverage.out && rm coverage.out
touch coverage.html && rm coverage.html
