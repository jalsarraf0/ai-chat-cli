#!/usr/bin/env bash
set -Eeuo pipefail
root=$(git rev-parse --show-toplevel)
cd "$root"

gofumpt -w \
    cmd/ai-chat/main.go \
    internal/hello/hello.go \
    internal/hello/hello_test.go \
    cmd/ai-chat/main_test.go

golangci-lint run ./...

go vet ./...

go test -race ./...
