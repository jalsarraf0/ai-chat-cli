#!/usr/bin/env bash
###############################################################################
# lint_cli.sh – run golangci-lint entirely via CLI flags (no YAML config)
#
# Why a script?
#   • Avoids YAML parse errors on self-hosted runners.
#   • Makes all settings explicit and easy to tweak in one place.
#
# Exit behaviour: any lint error exits >0 so CI fails appropriately.
###############################################################################
set -Eeuo pipefail

### 0. Environment (adjust here or export before calling) ---------------------
export GO_BIN="${GO_BIN:-$HOME/go/bin}"           # where `go install` puts tools
export COVERAGE_THRESHOLD="${COVERAGE_THRESHOLD:-93}"
export TRIVY_VERSION="${TRIVY_VERSION:-0.63.0}"
export GOLANGCI_VERSION="${GOLANGCI_VERSION:-v1.54.2}"
export GOSEC_VERSION="${GOSEC_VERSION:-v2.19.0}"
export OSV_VERSION="${OSV_VERSION:-v1.7.3}"

echo "🔧  Using golangci-lint ${GOLANGCI_VERSION}"

### 1. Ensure the requested golangci-lint version is installed ----------------
need_install=true
if command -v golangci-lint >/dev/null 2>&1; then
  current="$(golangci-lint --version | awk '/version:/ {print $4}')"
  if [[ "$current" == "${GOLANGCI_VERSION#v}" ]]; then
    need_install=false
  else
    echo "↻  Replacing golangci-lint $current with ${GOLANGCI_VERSION}"
  fi
fi

if $need_install; then
  go install "github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_VERSION}"
fi

# Ensure GO_BIN is on PATH for this run
export PATH="$GO_BIN:$PATH"

### 2. Run the linter via CLI flags ------------------------------------------
# --modules-download-mode=vendor  if you use `go mod vendor`
# --issues-exit-code=1            default (fails build if issues >0)
# --timeout                       same 5 min you use in CI
#
# Example: enable fast linters and disable a few noisy ones
golangci-lint run \
  --timeout 5m \
  --enable govet \
  --enable staticcheck \
  --enable gosec \
  --enable revive \
  --disable funlen \
  "$@"

echo "✅  golangci-lint finished with no (new) issues."
