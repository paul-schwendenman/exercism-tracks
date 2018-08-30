#!/usr/bin/env bash

set -o errexit
set -o nounset

main() {
  input="$@"

  echo "Hello, World!"
}

main "$@"
