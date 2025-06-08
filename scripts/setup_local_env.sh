#!/bin/bash

# Setup local env for development matching CI pipeline

set -euxo pipefail

# Create and export writable cache and tmp dirs for Go
export GOCACHE="$HOME/.cache/go-build"
export GOTMPDIR="$HOME/.cache/go-tmp"
mkdir -p "$GOCACHE" "$GOTMPDIR"

# Add Go bin dir to PATH
export PATH="$HOME/go/bin:$PATH"

# Install required Go tools
go install golang.org/x/vuln/cmd/govulncheck@latest
go install github.com/securego/gosec/v2/cmd/gosec@v2.19.0
go install github.com/google/osv-scanner/cmd/osv-scanner@v1.7.3
curl -sSfL https://github.com/golangci/golangci-lint/releases/download/v2.1.6/golangci-lint-2.1.6-linux-amd64.tar.gz | tar -xz -C $HOME/go/bin --strip-components=1 golangci-lint-2.1.6-linux-amd64/golangci-lint

echo "Local environment setup complete. Remember to source this script before running builds/tests."
