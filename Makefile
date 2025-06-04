.PHONY: format lint test docs build man shell-test cross

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
	cobra-cli man --dir docs/man

shell-test:
	go test -run TestShell ./internal/shell/...

cross:
	GOOS=windows GOARCH=amd64 $(MAKE) build
