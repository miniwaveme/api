#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

MINIWAVEME_ROOT=$(dirname "${BASH_SOURCE}")/..

GO_VERSION=($(go version))

if [[ -z $(echo "${GO_VERSION[2]}" | grep -E 'go1.2|go1.3|go1.4|go1.5') ]]; then
  echo "Unknown go version '${GO_VERSION}', skipping gofmt."
  exit 0
fi

cd "${MINIWAVEME_ROOT}"

find_files() {
  find . -name '*.go'
}

GOFMT="gofmt -s"
bad_files=$(find_files | xargs $GOFMT -l)
if [[ -n "${bad_files}" ]]; then
  echo "!!! '$GOFMT' needs to be run on the following files: "
  echo "${bad_files}"
  exit 1
fi