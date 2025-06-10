#!/usr/bin/env bash
set -euo pipefail

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
