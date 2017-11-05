#! /bin/bash

set -e

curl -sL https://git.io/goreleaser | bash -s -- --config .release.yml --skip-publish --snapshot --debug
