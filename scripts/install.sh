#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "== ai-chat-cli installer =="

PREFIX=/usr/local
while [ "$#" -gt 0 ]; do
    case "$1" in
        --prefix)
            PREFIX="$2"
            shift 2
            ;;
        *)
            echo "Unknown argument: $1" >&2
            exit 1
            ;;
    esac
done

read -rp "Install prefix [$PREFIX]: " ans || true
if [ -n "$ans" ]; then
    PREFIX="$ans"
fi
read -rp "Proceed installing to $PREFIX? [Y/n]: " ans || true
case "$ans" in
    n|N) echo "Cancelled"; exit 1;;
esac

pkg_install() {
    if command -v apt-get >/dev/null 2>&1; then
        sudo apt-get update -y && sudo apt-get install -y "$@"
    elif command -v dnf >/dev/null 2>&1; then
        sudo dnf install -y "$@"
    else
        echo "Unsupported OS" >&2
        exit 1
    fi
}

ensure() { command -v "$1" >/dev/null 2>&1 || pkg_install "$1"; }

ensure git
ensure curl
ensure go

if ! go version | grep -q "go1.24"; then
    echo "Go 1.24.x required" >&2
    exit 1
fi

OPENAI_API_KEY=${OPENAI_API_KEY:-}
if [ -z "$OPENAI_API_KEY" ]; then
    read -rp "Enter OPENAI_API_KEY (leave blank to edit later): " OPENAI_API_KEY || true
fi

echo "-- building ai-chat..."
go install ./cmd/ai-chat
bin="$(go env GOPATH)/bin/ai-chat"
if [ -x "$bin" ]; then
  target="$PREFIX/bin/ai-chat"
  mkdir -p "$(dirname "$target")"
  if [ -w "$(dirname "$target")" ]; then
    cp "$bin" "$target"
  else
    sudo cp "$bin" "$target"
  fi
fi

config_dir="${XDG_CONFIG_HOME:-$HOME/.config}/ai-chat-cli"
config_file="$config_dir/config.yaml"
mkdir -p "$config_dir"
if [ ! -f "$config_file" ]; then
    cat >"$config_file" <<EOF
openai_api_key: $OPENAI_API_KEY
model: gpt-4
EOF
    echo "Created $config_file. Add your API key if empty." >&2
fi

echo '✅ Installed!  Try:  ai-chat "Hello"'
