#!/usr/bin/env bash
set -euo pipefail

command -v go >/dev/null || { echo "❌ 'go' not on PATH"; exit 1; }
[[ $(go version) =~ go1\.24 ]] || { echo "❌ Go 1.24.x required"; exit 1; }

export PATH="$PWD/offline-bin:$PATH"
mkdir -p offline-bin

missing=()
for t in gofumpt staticcheck gosec; do
  command -v "$t" >/dev/null || missing+=("$t")
done
((${#missing[@]})) || exit 0

curl -sSfI https://proxy.golang.org >/dev/null || {
  echo "❌ no network, can't install ${missing[*]}" >&2; exit 1; }

echo "🔧 installing: ${missing[*]}"
for pkg in \
  mvdan.cc/gofumpt@latest \
  honnef.co/go/tools/cmd/staticcheck@latest \
  github.com/securego/gosec/v2/cmd/gosec@latest; do
    GOFLAGS='-trimpath' go install "$pkg"
done
