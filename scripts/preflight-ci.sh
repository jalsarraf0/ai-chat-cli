#!/usr/bin/env bash
set -euo pipefail

# Fast pre-flight: install Go 1.24.3 and Trivy v0.63.0 if missing
GO_VERSION=1.24.3
TRIVY_VERSION=v0.63.0
CACHE_DIR="${HOME}/.cache/preflight"
mkdir -p "$CACHE_DIR"

install_go() {
  if ! command -v go >/dev/null || ! go version | grep -q "go${GO_VERSION}"; then
    echo "Installing Go ${GO_VERSION}" >&2
    curl -sSfL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" -o "$CACHE_DIR/go.tgz"
    rm -rf "$CACHE_DIR/go" && mkdir -p "$CACHE_DIR/go"
    tar -C "$CACHE_DIR/go" --strip-components=1 -xzf "$CACHE_DIR/go.tgz"
  fi
  export PATH="$CACHE_DIR/go/bin:$PATH"
}

install_trivy() {
  if ! command -v trivy >/dev/null || ! trivy --version 2>/dev/null | grep -q "${TRIVY_VERSION}"; then
    echo "Installing Trivy ${TRIVY_VERSION}" >&2
    curl -sSfL "https://github.com/aquasecurity/trivy/releases/download/${TRIVY_VERSION}/trivy_${TRIVY_VERSION#v}_Linux-64bit.tar.gz" -o "$CACHE_DIR/trivy.tgz"
    rm -rf "$CACHE_DIR/trivy" && mkdir -p "$CACHE_DIR/trivy"
    tar -C "$CACHE_DIR/trivy" -xzf "$CACHE_DIR/trivy.tgz"
    chmod +x "$CACHE_DIR/trivy/trivy"
  fi
  export PATH="$CACHE_DIR/trivy:$PATH"
}

install_gocovmerge() {
  if ! command -v gocovmerge >/dev/null; then
    echo "Installing gocovmerge" >&2
    GOFLAGS='-trimpath' go install github.com/wadey/gocovmerge@latest
  fi
}

install_go
install_trivy
install_gocovmerge

# Display tool versions
go version
trivy --version
