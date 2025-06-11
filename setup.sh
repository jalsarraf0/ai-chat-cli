#!/usr/bin/env bash
set -euo pipefail
# Wrapper script to run the interactive installer regardless of cwd
SCRIPT_DIR="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "Running ai-chat-cli setup. Press Ctrl+C to cancel." >&2
exec "${SCRIPT_DIR}/scripts/install.sh" "$@"
