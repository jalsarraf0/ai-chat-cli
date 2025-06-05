#!/usr/bin/env bash
# install_tools() – Phase 9.2
install_tools() {
  export PATH="$PWD/offline-bin:$PATH"
  mkdir -p offline-bin

  gover=$(go list -f '{{ goVersion }}' -m)
  [[ $gover =~ ^go1\.24 ]] || { echo "❌ Go 1.24.x required, found $gover" >&2; exit 1; }

  local missing=()
  for t in gofumpt staticcheck gosec cosign; do
    command -v "$t" >/dev/null || missing+=("$t")
  done
  ((${#missing[@]})) || return 0

  if ! command -v curl >/dev/null && ! ping -c1 google.com >/dev/null 2>&1; then
    echo "❌ missing tools (${missing[*]}) and no network" >&2
    exit 1
  fi

  echo "🔧 installing: ${missing[*]}"
  for pkg in \
    mvdan.cc/gofumpt@latest \
    honnef.co/go/tools/cmd/staticcheck@latest \
    github.com/securego/gosec/v2/cmd/gosec@latest \
    github.com/sigstore/cosign/v2/cmd/cosign@latest; do
    GOFLAGS='-trimpath' go install "$pkg"
  done
}

if [[ $1 == install_tools ]]; then
  install_tools
fi
