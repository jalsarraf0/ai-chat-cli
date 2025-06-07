#!/usr/bin/env bash
set -euo pipefail

export PATH="/usr/local/bin:$HOME/go/bin:$PATH"

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

missing=()
for t in go trivy govulncheck gosec godoc mdbook make; do
  if ! command -v "$t" >/dev/null; then
    echo "MISSING: $t"
    missing+=("$t")
  fi
done

for t in "${missing[@]}"; do
  case "$t" in
    trivy)
      echo "Installing trivy..."
      GOBIN=/usr/local/bin go install github.com/aquasecurity/trivy/cmd/trivy@latest
      ;;
    govulncheck)
      echo "Installing govulncheck..."
      GOBIN=/usr/local/bin go install golang.org/x/vuln/cmd/govulncheck@latest
      ;;
    gosec)
      echo "Installing gosec..."
      GOBIN=/usr/local/bin go install github.com/securego/gosec/v2/cmd/gosec@latest
      ;;
    godoc)
      echo "Installing godoc..."
      GOBIN=/usr/local/bin go install golang.org/x/tools/cmd/godoc@latest
      ;;
    mdbook)
      echo "Installing mdbook..."
      MDBOOK_VERSION="0.4.51"
      wget -q "https://github.com/rust-lang/mdBook/releases/download/v${MDBOOK_VERSION}/mdbook-v${MDBOOK_VERSION}-x86_64-unknown-linux-gnu.tar.gz" -O /tmp/mdbook.tar.gz
      sudo tar -xf /tmp/mdbook.tar.gz -C /usr/local/bin
      sudo chmod +x /usr/local/bin/mdbook
      ;;
    make)
      echo "Installing make..."
      sudo nala install -y make
      ;;
    go)
  echo "Installing Go ${GO_VERSION}" && install_go
      ;;
    *)
      echo "Don't know how to install: $t"
      ;;
  esac
done

install_go
install_trivy
install_gocovmerge

# Persist PATH for subsequent CI steps
{ echo "$CACHE_DIR/go/bin"; echo "$CACHE_DIR/trivy"; echo "$(go env GOPATH)/bin"; } >>"${GITHUB_PATH:-$HOME/.github_path}"

# Display tool versions
go version
trivy --version
