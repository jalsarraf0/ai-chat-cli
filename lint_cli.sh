#!/usr/bin/env bash
set -euo pipefail
export PATH="$HOME/go/bin:$PATH"

want="2.1.6"
ver="$(command -v golangci-lint >/dev/null && golangci-lint --version | awk '{print $4}' || echo '')"
if [[ "$ver" != "$want" ]]; then
  echo "🔧 Installing golangci-lint $want"
  GO111MODULE=on go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v${want}
fi

golangci-lint run ./...
