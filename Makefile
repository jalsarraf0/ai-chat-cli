.PHONY: format lint lint-all static security test docs build man shell-test cross
GOFILES := $(shell git ls-files '*.go')
export PATH := $(HOME)/.local/bin:$(PATH)

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

coverage:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
.PHONY: coverage

coverage-gate:
        @pct=$$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","" );print $$3}'); \
        th=93; if echo "$$pct < $$th" | bc -l | grep -q 1; then \
            echo "::error::coverage < $$th% (got $$pct%)"; exit 1; fi

install-hugo:
	./scripts/install_hugo.sh

docs: install-hugo
	@echo "📖 Building Hugo site"
	hugo --minify

readme: install-hugo
	@echo "📝 Re-rendering README.md"
	hugo --quiet -D --renderToREADME -d .

docs-all: docs readme

clean:
	rm -f bin/ai-chat-linux-amd64

build:
	GOOS=linux GOARCH=amd64 \
	go build -o bin/ai-chat-linux-amd64 ./cmd/ai-chat

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
	@command -v goreleaser >/dev/null || (\
	curl -sSL https://github.com/goreleaser/goreleaser/releases/download/v2.9.0/goreleaser_Linux_x86_64.tar.gz \
	| tar -xz goreleaser && sudo mv goreleaser /usr/local/bin/)
	goreleaser release --snapshot --clean --skip=publish --skip=docker --skip=sign
	
package:
	@command -v goreleaser >/dev/null || (\
	curl -sSL https://github.com/goreleaser/goreleaser/releases/download/v2.9.0/goreleaser_Linux_x86_64.tar.gz \
	| tar -xz goreleaser && sudo mv goreleaser /usr/local/bin/)
	goreleaser release --snapshot --clean

release:
	       @command -v goreleaser >/dev/null || (\
				               curl -sSL https://github.com/goreleaser/goreleaser/releases/download/v2.9.0/goreleaser_Linux_x86_64.tar.gz \
				               | tar -xz goreleaser && sudo mv goreleaser /usr/local/bin/)
	       goreleaser release --clean --skip=publish --skip=docker


live-openai-test:
	go test ./pkg/llm/openai -run Live -v
