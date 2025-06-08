#!/usr/bin/env bash
# scripts/preflight-v9-hugo.sh -- install Hugo and Trivy if missing
set -euo pipefail

HUGO_VERSION=0.128.1
TRIVY_VERSION=v0.63.0
CACHE_DIR="${HOME}/.cache/preflight-v9"
mkdir -p "$CACHE_DIR"

install_hugo() {
  if ! command -v hugo >/dev/null || ! hugo version 2>/dev/null | grep -q "v${HUGO_VERSION}"; then
    echo "Installing Hugo ${HUGO_VERSION}" >&2
    curl -sSfL "https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_extended_${HUGO_VERSION}_Linux-64bit.tar.gz" -o "$CACHE_DIR/hugo.tgz"
    rm -rf "$CACHE_DIR/hugo" && mkdir -p "$CACHE_DIR/hugo"
    tar -C "$CACHE_DIR/hugo" -xzf "$CACHE_DIR/hugo.tgz"
    chmod +x "$CACHE_DIR/hugo/hugo"
  fi
  export PATH="$CACHE_DIR/hugo:$PATH"
}

install_trivy() {
  if ! command -v trivy >/dev/null || ! trivy --version 2>/dev/null | grep -q "${TRIVY_VERSION}"; then
    echo "Installing Trivy ${TRIVY_VERSION}" >&2
    GOFLAGS='-trimpath' go install github.com/aquasecurity/trivy/cmd/trivy@latest
  fi
}

install_hugo &
install_trivy &
wait

hugo version
trivy --version
