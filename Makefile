	.PHONY: format lint lint-all static security test docs build man shell-test cross
GOFILES := $(shell git ls-files '*.go')

format:
	gofumpt -l -w $(GOFILES)

lint: ## static analysis
        ./scripts/offline_first.sh
        golangci-lint run ./...

lint-all:
        ./scripts/offline_first.sh
        golangci-lint run ./...

static:
        ./scripts/offline_first.sh
        staticcheck ./...

security:
        ./scripts/offline_first.sh
        gosec ./... && govulncheck ./...

ifeq ($(strip $(GOMAXPROCS)),)
PARFLAG :=
else
PARFLAG := -p $(GOMAXPROCS)
endif

unit: ## unit tests, offline
        ./scripts/offline_first.sh
        go test $(PARFLAG) -race -covermode=atomic -coverprofile=coverage.out -tags unit ./...

test: lint unit
        ./scripts/offline_first.sh
        @$(MAKE) coverage-gate

# ---------------------------------------------------------------------------
# Security analysis
# ---------------------------------------------------------------------------
.PHONY: security-scan

security-scan: ## Run gosec static analysis
	GOFLAGS='-trimpath' gosec ./...

coverage-gate:
	@pct=$$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","" );print $$3}')
	th=93; if echo "$$pct < $$th" | bc -l | grep -q 1; then \
	echo "::error::coverage < $$th% (got $$pct%)"; exit 1; fi

docs:
	./scripts/offline_first.sh
	@git ls-files "*.md" | xargs -r sed -i "s/[ 	]*$$//" && git diff --exit-code || true
	npm install
	npm audit fix --force || true
	npm audit --audit-level=high
	@echo '#!/usr/bin/env bash\nexec mdbook "$@"' > node_modules/.bin/mdbook
	@chmod +x node_modules/.bin/mdbook
	npx mdbook build docs
	


build:
	./scripts/offline_first.sh
	go build -o bin/ai-chat-cli-linux-amd64 .

man:
	cobra-cli man --dir docs/man

shell-test:
	./scripts/offline_first.sh
	go test -run TestShell ./internal/shell/...

cross:
	./scripts/offline_first.sh
	GOOS=windows GOARCH=amd64 $(MAKE) build

tui: ## run terminal UI
	./scripts/offline_first.sh
	go run ./cmd/ai-chat tui --height 20

embed-check: ## verify embedded FS is up to date
	go run scripts/embedgen.go
	@git diff --quiet internal/assets || (echo "::error::embed drift"; exit 1)


prompt:
	@mkdir -p dist/prompt && echo '#!/bin/sh\necho prompt' > dist/prompt/stub.sh
	chmod +x dist/prompt/stub.sh

snapshot:
	./scripts/offline_first.sh
	@command -v goreleaser >/dev/null || GOFLAGS= go install github.com/goreleaser/goreleaser@latest
	goreleaser release --snapshot --clean --skip=publish --skip=docker --skip=sign

release:
	./scripts/offline_first.sh
	@command -v goreleaser >/dev/null || GOFLAGS= go install github.com/goreleaser/goreleaser@latest
	goreleaser release --clean --skip=publish --skip=docker


live-openai-test:
	go test ./pkg/llm/openai -run Live -v
