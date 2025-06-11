#!/usr/bin/env sh
set -e

# Remove installed binaries and configuration, including any API keys.

BIN="$(command -v ai-chat 2>/dev/null || true)"
GOBIN="$(go env GOBIN 2>/dev/null)"
if [ -z "$GOBIN" ]; then
    GOBIN="$(go env GOPATH 2>/dev/null)/bin"

fi
DEFAULT_BIN="$GOBIN/ai-chat"

CONFIG_PATH="${AI_CHAT_CONFIG:-}"
if [ -z "$CONFIG_PATH" ]; then
    base="${XDG_CONFIG_HOME:-$HOME/.config}"
    CONFIG_PATH="$base/ai-chat-cli/config.yaml"
fi

fi
DEFAULT_BIN="$GOBIN/ai-chat"

CONFIG_PATH="${AI_CHAT_CONFIG:-}"
if [ -z "$CONFIG_PATH" ]; then
    base="${XDG_CONFIG_HOME:-$HOME/.config}"
    CONFIG_PATH="$base/ai-chat-cli/config.yaml"
fi

CONFIG_DIR="$(dirname "$CONFIG_PATH")"

if [ "$1" != "--yes" ]; then
    printf 'This will remove %s %s and %s\n' "$DEFAULT_BIN" "$BIN" "$CONFIG_DIR"
    printf 'Continue? [y/N]: '
    read -r ans
    case "$ans" in
        y|Y) ;;
        *) echo "Aborted."; exit 1;;
    esac
fi


# Remove binaries if present, using sudo if necessary.
remove() {
    target="$1"
    if [ -e "$target" ]; then
        if rm -f "$target" 2>/dev/null; then
            :
        else
            sudo rm -f "$target"
        fi
    fi
}

remove "$DEFAULT_BIN"
if [ -n "$BIN" ] && [ "$BIN" != "$DEFAULT_BIN" ]; then
    remove "$BIN"
fi

# Remove binaries if present.
rm -f "$DEFAULT_BIN"
[ -n "$BIN" ] && [ "$BIN" != "$DEFAULT_BIN" ] && rm -f "$BIN"


# Remove configuration directory containing credentials.
rm -rf "$CONFIG_DIR"

echo "Uninstalled ai-chat and removed configuration"
