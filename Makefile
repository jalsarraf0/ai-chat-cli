.PHONY: format lint test docs build man shell-test cross completion init prompt-check

GOFILES := $(shell find . -type f -name '*.go' -not -path './vendor/*')

format:
	gofumpt -l -w $(GOFILES)

lint:
	GOTOOLCHAIN=go1.22.4 golangci-lint run ./...

test:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
	pct=$$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","");print $$3}'); \
	if [ $${pct%.*} -lt 88 ]; then echo "::error::coverage < 88%" && exit 1; fi

docs:
	@git ls-files '*.md' | xargs -r sed -i 's/[ \t]*$$//' && git diff --exit-code || true

build:
	go build -o bin/ai-chat ./cmd/ai-chat

man:
	go run ./scripts/genman.go

shell-test:
	go test -run TestShell ./internal/shell/...

cross:
	GOOS=windows GOARCH=amd64 $(MAKE) build

completion:
	go run ./cmd/ai-chat completion bash --out dist/completion/bash
	go run ./cmd/ai-chat completion zsh --out dist/completion/zsh
	go run ./cmd/ai-chat completion fish --out dist/completion/fish
	go run ./cmd/ai-chat completion powershell --out dist/completion/powershell

init:
	go run ./cmd/ai-chat init --dry-run

prompt-check:
	shellcheck dist/prompt/*
