#!/usr/bin/env bash
set -euo pipefail
# Wrapper script to run the interactive installer regardless of cwd
SCRIPT_DIR="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
exec "${SCRIPT_DIR}/scripts/install.sh" "$@"
