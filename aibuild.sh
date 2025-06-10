#!/usr/bin/env bash
# SPDX-License-Identifier: MIT
#
# Build ai-chat and generate .tar.gz, .deb and .rpm packages on Fedora 42.
# Usage: ./build-packages.sh [version]
#
# Pass a semantic version like v1.2.3 and the script will tag & build that
# release; omit it for a snapshot build.

set -euo pipefail

APP_NAME="ai-chat"      # must match go.mod & main package
GO_VERSION="1.24"
ARCH="amd64"
DIST_DIR="dist"

###############################################################################
# 1. Install Fedora-side prerequisites
###############################################################################
echo "⇒ Installing build prerequisites…"
sudo dnf -y install golang git gcc make gzip tar rpm-build binutils curl util-linux

# Verify Go 1.24.x is really on $PATH
if ! go version | grep -q "go${GO_VERSION}"; then
  echo "❌ Go ${GO_VERSION}.x not found. Install it and re-run." >&2
  exit 1
fi

# Install goreleaser (bundles nfpm) if missing
if ! command -v goreleaser &>/dev/null; then
  echo "⇒ Installing goreleaser…"
  tmp=$(mktemp -d)
  curl -sSL \
    https://github.com/goreleaser/goreleaser/releases/latest/download/goreleaser_Linux_x86_64.tar.gz |
    tar -xz -C "$tmp"
  sudo install -m 0755 "$tmp/goreleaser" /usr/local/bin/
  rm -rf "$tmp"
fi

###############################################################################
# 2. Determine version / tag
###############################################################################
VERSION=${1:-$(git describe --tags --abbrev=0 2>/dev/null || true)}
if [[ -z "$VERSION" ]]; then
  SNAPSHOT="--snapshot"
  echo "⇒ Snapshot build (no Git tag specified)."
else
  SNAPSHOT=""
  echo "⇒ Building version $VERSION"
  if ! git rev-parse "$VERSION" &>/dev/null; then
    git tag -a "$VERSION" -m "Release $VERSION"
  fi
fi

###############################################################################
# 3. Clean & build
###############################################################################
echo "⇒ Cleaning dist/…"
rm -rf "$DIST_DIR"

echo "⇒ Running GoReleaser (packages only, skip Docker)…"
goreleaser release --clean --skip=docker,publish $SNAPSHOT

echo -e "\n✅  Packages ready in $DIST_DIR:\n"
ls -1 "$DIST_DIR"
