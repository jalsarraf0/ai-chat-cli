#!/usr/bin/env bash
set -e
# Wrapper script to run the interactive installer
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
exec "$SCRIPT_DIR/scripts/install.sh" "$@"
