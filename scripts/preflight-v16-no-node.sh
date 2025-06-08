#!/usr/bin/env bash
set -euo pipefail

# Preflight script for self-hosted CI runners without Node
# Ensures required tools are available and prints their versions.

need() {
  command -v "$1" >/dev/null || { echo "missing $1" >&2; exit 1; }
}

for tool in go golangci-lint gosec govulncheck osv-scanner trivy goreleaser; do
  need "$tool"
  if [ "$tool" = go ]; then
    go version
  else
    "$tool" --version | head -n 1
  fi
done

