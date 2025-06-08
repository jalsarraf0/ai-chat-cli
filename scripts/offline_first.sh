#!/usr/bin/env bash
set -euo pipefail

GO_VERSION=${GO_VERSION:-1.24}
NODE_VERSION=${NODE_VERSION:-20}

have() { command -v "$1" >/dev/null 2>&1; }

require_version() {
  local cmd="$1" want="$2"
  [[ -z "$want" ]] && return 0
  have "$cmd" || return 1
  "$cmd" --version 2>&1 | grep -Eq "$want"
}

need_bin() {
  local cmd="$1" regex="$2"
  if require_version "$cmd" "$regex"; then
    echo "✅  Using local $cmd ($( "$cmd" --version | head -1))"
    return 1
  fi
  return 0
}

declare -A INSTALLER=(
  [go]="sudo apt-get update -qq && sudo apt-get install -y golang-${GO_VERSION%.*}"
  [golangci-lint]="GOFLAGS=-buildvcs=false go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
  [goreleaser]="curl -sSfL https://git.io/goreleaser | bash -s -- -y"
  [trivy]="sudo apt-get install -y trivy && trivy plugin install ghcr.io/aquasecurity/trivy-plugin-config@latest"
  [node]="curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION}.x | sudo -E bash - && sudo apt-get install -y nodejs"
)

provision() {
  local missing=()
  while read -r tool regex; do
    need_bin "$tool" "$regex" || continue
    missing+=("$tool")
  done <<EOF_TOOLS
go            ^go${GO_VERSION%.*}\.
golangci-lint v1
goreleaser    v
trivy         Trivy
EOF_TOOLS

  printf '%s\n' "${missing[@]}" | sort -u | xargs -I{} -P"$(nproc)" bash -c "${INSTALLER[{}}]"
}

provision "$@"
