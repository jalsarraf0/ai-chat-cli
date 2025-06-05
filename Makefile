.PHONY: format lint lint-all static security test docs build man shell-test cross
		
GOFILES := $(shell git ls-files '*.go')

format:
	gofumpt -l -w $(GOFILES)

lint: ## static analysis
	golangci-lint run ./...

lint-all:
	golangci-lint run ./...

static:
	staticcheck ./...

security:
	gosec ./... && govulncheck ./...

ifeq ($(strip $(GOMAXPROCS)),)
PARFLAG :=
else
PARFLAG := -p $(GOMAXPROCS)
endif

unit: ## unit tests, offline
	go test $(PARFLAG) -race -covermode=atomic -coverprofile=coverage.out -tags unit ./...

test: lint unit
	@$(MAKE) coverage-gate

# ---------------------------------------------------------------------------
# Security analysis
# ---------------------------------------------------------------------------
.PHONY: security-scan

security-scan: ## Run gosec static analysis
	GOFLAGS='-trimpath' gosec ./...

coverage-gate:
	@pct=$$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","" );print $$3}'); \
       if [ $${pct%.*} -lt 92 ]; then \
       echo "::error::coverage < 92% (got $${pct}%)"; exit 1; fi

docs:
	@git ls-files '*.md' | xargs -r sed -i 's/[ \t]*$$//' && git diff --exit-code || true

build:
	go build -o bin/ai-chat ./cmd/ai-chat

man:
	cobra-cli man --dir docs/man

shell-test:
	go test -run TestShell ./internal/shell/...

cross:
	GOOS=windows GOARCH=amd64 $(MAKE) build

tui: ## run terminal UI
	go run ./cmd/ai-chat tui --height 20

embed-check: ## verify embedded FS is up to date
	go run scripts/embedgen.go
	@git diff --quiet internal/assets || (echo "::error::embed drift"; exit 1)


prompt:
	@mkdir -p dist/prompt && echo '#!/bin/sh\necho prompt' > dist/prompt/stub.sh
	chmod +x dist/prompt/stub.sh

snapshot:
	@command -v goreleaser >/dev/null || GOFLAGS= go install github.com/goreleaser/goreleaser@latest
	goreleaser release --snapshot --clean --skip=publish --skip=docker --skip=sign

release:
	@command -v goreleaser >/dev/null || GOFLAGS= go install github.com/goreleaser/goreleaser@latest
	goreleaser release --clean --skip=publish --skip=docker


live-openai-test:
	go test ./pkg/llm/openai -run Live -v

.PHONY: bench bench-json

# Run benchmarks once without saving output
bench:
	go test -run=^$ -bench=. ./...

# Run benchmarks three times and save JSON results
bench-json:
	go test -run=^$ -bench=. -benchmem -count=3 -json ./... > bench.json
