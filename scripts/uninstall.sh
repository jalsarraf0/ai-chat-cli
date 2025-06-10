#!/usr/bin/env sh
set -e
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
BIN=$(go env GOBIN 2>/dev/null)
if [ -z "$BIN" ]; then
    BIN="$(go env GOPATH 2>/dev/null)/bin"
fi
BIN="$BIN/ai-chat"
CONFIG_DIR="${XDG_CONFIG_HOME:-$HOME/.config}/ai-chat"

if [ "$1" != "--yes" ]; then
    printf 'This will remove %s and %s\n' "$BIN" "$CONFIG_DIR"
    printf 'Continue? [y/N]: '
    read -r ans
    case "$ans" in
        y|Y) ;;
        *) echo "Aborted."; exit 1;;
    esac
fi

rm -f "$BIN"
rm -rf "$CONFIG_DIR"

echo "Uninstalled ai-chat"
