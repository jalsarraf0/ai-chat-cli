#!/usr/bin/env bash
# scripts/offline_first.sh
# ---------------------------------------------------------------------------
# LOCAL‑FIRST provisioner for self‑hosted Linux CI runners.
# Re‑use existing tools if they match required versions; otherwise install
# them on‑the‑fly.  Designed for Ubuntu ≥ 22.04 but works on any Debian family
# image that has apt.
#
# Tools handled
#   • Go                (apt)
#   • golangci-lint     (go install)
#   • gosec             (go install)
#   • govulncheck       (go install)
#   • trivy             (apt)
#   • goreleaser        (curl installer)
#   • addlicense        (go install)
#   • gocovmerge        (go install)
#
# Idempotent, parallel, and fails fast on any install error.
# ---------------------------------------------------------------------------

set -euo pipefail

GO_VERSION=${GO_VERSION:-1.24}
NODE_VERSION=${NODE_VERSION:-20}   # future‑proof; not used yet

# detect whether we need sudo (root inside container doesn't have it)
SUDO=$(command -v sudo || true)

have() { command -v "$1" &>/dev/null; }

require_version() {
  local cmd="$1" want="$2"
  [[ -z "$want" ]] && return 0
  have "$cmd" || return 1
  "$cmd" --version 2>&1 | grep -Eq "$want"
}

need_bin() {
  local cmd="$1" regex="$2"
  if require_version "$cmd" "$regex"; then
    echo "✅  [local] $cmd $( $cmd --version 2>&1 | head -1 )"
    return 1   # satisfied
  fi
  return 0       # missing / outdated
}

# ---------------------------------------------------------------------------
# Install command catalogue
#   Values are eval'd inside install_tool().
#   Keep quoting minimal — $SUDO is expanded in the runtime shell.
# ---------------------------------------------------------------------------
declare -A INSTALLER=(
  [go]="${SUDO:+$SUDO }apt-get install -y golang-${GO_VERSION%.*}"
  [golangci-lint]='GOFLAGS=-buildvcs=false go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest'
  [gosec]='GOFLAGS=-buildvcs=false go install github.com/securego/gosec/v2/cmd/gosec@latest'
  [govulncheck]='GOFLAGS=-buildvcs=false go install golang.org/x/vuln/cmd/govulncheck@latest'
  [trivy]="${SUDO:+$SUDO }apt-get install -y trivy"
  [goreleaser]='curl -sSfL https://git.io/goreleaser | bash -s -- -y'
  [addlicense]='GOFLAGS=-buildvcs=false go install github.com/google/addlicense@latest'
  [gocovmerge]='GOFLAGS=-buildvcs=false go install github.com/wadey/gocovmerge@latest'
)

install_tool() {
  local tool="$1"
  echo "🔧  Installing $tool …"
  # shellcheck disable=SC2086
  eval ${INSTALLER[$tool]}
  echo "✅  Installed $tool"
}

export -f install_tool have require_version need_bin
export SUDO GO_VERSION
export INSTALLER

provision() {
  local missing=()
  while read -r tool regex; do
    need_bin "$tool" "$regex" || missing+=("$tool")
  done <<'EOF_DEPS'
go            ^go'${GO_VERSION%.*}'\.
golangci-lint v
gosec         v
govulncheck   v
trivy         Trivy
goreleaser    v
addlicense
gocovmerge
EOF_DEPS

  (( ${#missing[@]} )) || { echo "🎉  All required tools present"; return 0; }

  echo "📦  Missing tools: ${missing[*]}"

  # One apt-get update speeds up all apt installs
  if printf '%s\n' "${missing[@]}" | grep -qE '^(go|trivy)$'; then
    ${SUDO:+$SUDO }apt-get update -qq
  fi

  # Install everything in parallel (bounded by CPU count)
  printf '%s\n' "${missing[@]}" | xargs -n1 -P"$(nproc)" -I{} bash -c 'install_tool "$@"' _ {}

  echo "✅  Provisioning complete"
}

provision "$@"
