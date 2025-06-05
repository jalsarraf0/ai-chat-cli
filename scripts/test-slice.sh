#!/usr/bin/env bash
set -euo pipefail

# Allow CASE="1/4" etc.; empty means run full suite
CASE=${CASE:-}

# Warn & run full suite if malformed
if [[ -n "$CASE" && ! "$CASE" =~ ^[0-9]+/[0-9]+$ ]]; then
  echo "⚠️  Invalid CASE: '$CASE' – running full suite instead"
  CASE=""
fi

if [[ -n "$CASE" ]]; then
  export GOTESTSUM_FORMAT=short-verbose
  gotestsum --subset "${CASE}" --packages ./... --coverprofile=coverage.out \
            -- -race -covermode=atomic -tags unit
else
  go test -race -covermode=atomic -coverprofile=coverage.out -tags unit ./...
fi
