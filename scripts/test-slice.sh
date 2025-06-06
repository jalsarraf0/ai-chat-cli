#!/usr/bin/env bash
# Run unit tests slice with coverage
# shellcheck disable=SC2086
set -euo pipefail

GOFLAGS=${GOFLAGS:-}

go test $GOFLAGS -race -covermode=atomic -coverprofile=coverage.out -tags unit ./...
