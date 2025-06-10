#!/usr/bin/env bash
set -euo pipefail














echo "== ai-chat-cli installer =="

require() { command -v "$1" >/dev/null 2>&1 || { echo "Error: $1 not found" >&2; exit 1; }; }
require go
require git

if ! go version | grep -q "go1.24"; then
    echo "Go 1.24.x required" >&2
    exit 1
fi

if ! command -v docker >/dev/null 2>&1; then
    echo "Warning: docker not installed; some features may be disabled" >&2

fi

OPENAI_API_KEY=${OPENAI_API_KEY:-}
if [ -z "$OPENAI_API_KEY" ]; then
    read -rp "Enter OPENAI_API_KEY: " OPENAI_API_KEY
fi
[ -z "$OPENAI_API_KEY" ] && { echo "API key required" >&2; exit 1; }

echo "-- building ai-chat..."
go install ./cmd/ai-chat
bin="$(go env GOPATH)/bin/ai-chat"
if [ -x "$bin" ]; then
  if [ -w /usr/local/bin ]; then
    cp "$bin" /usr/local/bin/
  else
    sudo cp "$bin" /usr/local/bin/
  fi
fi


fi

OPENAI_API_KEY=${OPENAI_API_KEY:-}
if [ -z "$OPENAI_API_KEY" ]; then
    read -rp "Enter OPENAI_API_KEY: " OPENAI_API_KEY
fi
[ -z "$OPENAI_API_KEY" ] && { echo "API key required" >&2; exit 1; }

echo "-- building ai-chat..."
go install ./cmd/ai-chat
bin="$(go env GOPATH)/bin/ai-chat"
if [ -x "$bin" ]; then
  if [ -w /usr/local/bin ]; then
    cp "$bin" /usr/local/bin/
  else
    sudo cp "$bin" /usr/local/bin/
  fi
fi


fi

OPENAI_API_KEY=${OPENAI_API_KEY:-}
if [ -z "$OPENAI_API_KEY" ]; then
    read -rp "Enter OPENAI_API_KEY: " OPENAI_API_KEY
fi

fi

OPENAI_API_KEY=${OPENAI_API_KEY:-}
if [ -z "$OPENAI_API_KEY" ]; then
    read -rp "Enter OPENAI_API_KEY: " OPENAI_API_KEY
fi

[ -z "$OPENAI_API_KEY" ] && { echo "API key required" >&2; exit 1; }

echo "-- building ai-chat..."
go install ./cmd/ai-chat



bin="$(go env GOPATH)/bin/ai-chat"
if [ -x "$bin" ]; then
  if [ -w /usr/local/bin ]; then
    cp "$bin" /usr/local/bin/
  else
    sudo cp "$bin" /usr/local/bin/
  fi
fi








config_dir="${XDG_CONFIG_HOME:-$HOME/.config}/ai-chat"
config_file="$config_dir/ai-chat.yaml"
mkdir -p "$config_dir"
if [ ! -f "$config_file" ]; then
cat >"$config_file" <<EOF
openai_api_key: $OPENAI_API_KEY
model: gpt-4o
EOF
fi

if command -v pre-commit >/dev/null 2>&1; then
    read -rp "Install git hooks? [y/N]: " ans
    case "$ans" in
        y|Y) pre-commit install;;
    esac
fi


echo "Done. Try running: ai-chat \"Hello\""


echo "Done. Try running: ai-chat \"Hello\""


echo "Done. Try running: ai-chat \"Hello\""


echo "Done. Try running: ai-chat \"Hello\""

echo "Done. Try: ai-chat \"Hello\""

require(){ command -v "$1" >/dev/null 2>&1 || { echo "$1 required"; exit 1; }; }
require go
require docker
require git

go version | grep -q "go1.24" || { echo "Go 1.24.x required"; exit 1; }

read -rp "OPENAI_API_KEY: " OPENAI_API_KEY
[ -z "$OPENAI_API_KEY" ] && { echo "key required"; exit 1; }

make deps && make build

go install ./cmd/ai-chat

if command -v pre-commit >/dev/null 2>&1; then
    read -rp "Install git hooks? [y/N]: " ans
    case "$ans" in
        y|Y) pre-commit install;;
    esac
fi

echo "Installation complete"





