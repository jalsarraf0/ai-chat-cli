#!/bin/bash

    # Setup local env for development matching CI pipeline with sudo
    set -euxo pipefail

    # Create writable cache and tmp dirs for Go under your home
    export GOCACHE="$HOME/.cache/go-build"
    export GOTMPDIR="$HOME/.cache/go-tmp"
    sudo mkdir -p "$GOCACHE" "$GOTMPDIR"
    sudo chown -R $(id -u):$(id -g) "$HOME/.cache"

    # Add Go bin dir to PATH
    export PATH="$HOME/go/bin:$PATH"

    # Ensure Go bin directory exists and owned
    mkdir -p "$HOME/go/bin"
    sudo chown -R $(id -u):$(id -g) "$HOME/go"

    # Install required Go tools
    sudo -u $(id -un) go install golang.org/x/vuln/cmd/govulncheck@latest
    sudo -u $(id -un) go install github.com/securego/gosec/v2/cmd/gosec@v2.19.0
    sudo -u $(id -un) go install github.com/google/osv-scanner/cmd/osv-scanner@v1.7.3

    # Install golangci-lint v2.1.6
    sudo -u $(id -un) bash -c '
      curl -sSfL https://github.com/golangci/golangci-lint/releases/download/v2.1.6/golangci-lint-2.1.6-linux-amd64.tar.gz | \
      tar -xz -C "$HOME/go/bin" --strip-components=1 golangci-lint-2.1.6-linux-amd64/golangci-lint
    '

    echo "Sudo local environment setup complete. Please re-login or source your profile."
