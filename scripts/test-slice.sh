#!/usr/bin/env bash
set -euo pipefail
CASE=${CASE:-}

if [[ -n "$CASE" && ! "$CASE" =~ ^[0-9]+/[0-9]+$ ]]; then
  echo "⚠️  Invalid CASE '$CASE' — running full suite"
  CASE=""
fi

if [[ -n "$CASE" ]]; then
  export GOTESTSUM_FORMAT=short-verbose
  gotestsum --subset "$CASE" --packages ./... --coverprofile=coverage.out \
            -- -race -covermode=atomic
else
  go test -race -covermode=atomic -coverprofile=coverage.out ./...
fi

