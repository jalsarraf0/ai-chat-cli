#!/usr/bin/env bash
# Run unit tests slice with coverage
# shellcheck disable=SC2086
set -euo pipefail

GOFLAGS=${GOFLAGS:-}

echo "Running tests with coverage..."
go test $GOFLAGS -race -covermode=atomic -coverprofile=coverage.out -tags unit ./...

if [[ ! -s coverage.out ]]; then
  echo "Error: coverage.out is empty or missing" >&2
  exit 1
fi
echo "Successfully generated coverage.out"
