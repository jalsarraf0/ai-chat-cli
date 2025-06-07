#!/usr/bin/env bash
set -euo pipefail
# Build and run a simple help command to verify binary works
bin="./ai-chat"
go build -o "$bin" ./
$bin --help >/tmp/e2e.out
rm -f "$bin"
cat /tmp/e2e.out
