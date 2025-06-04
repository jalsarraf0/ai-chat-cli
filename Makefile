.PHONY: format lint test docs build man shell-test cross

GOFILES := $(shell git ls-files '*.go')

format:
	gofumpt -l -w $(GOFILES)

lint:
        GOTOOLCHAIN=go1.22.4 golangci-lint run ./...

ifeq ($(strip $(GOMAXPROCS)),)
PARFLAG :=
else
PARFLAG := -p $(GOMAXPROCS)
endif

test: lint
	go test $(PARFLAG) -race -covermode=atomic -coverprofile=coverage.out ./...
	@$(MAKE) coverage-gate

coverage-gate:
	@pct=$$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","" );print $$3}'); \
	if [ $${pct%.*} -lt 90 ]; then \
	echo "::error::coverage < 90% (got $${pct}%)"; exit 1; fi

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

prompt:
	@mkdir -p dist/prompt && echo '#!/bin/sh\necho prompt' > dist/prompt/stub.sh
	chmod +x dist/prompt/stub.sh
