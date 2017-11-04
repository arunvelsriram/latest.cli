#! /bin/bash

set -e

test -n "$TRAVIS_TAG" && curl -sL https://git.io/goreleaser | bash -s -- --config release.yml
