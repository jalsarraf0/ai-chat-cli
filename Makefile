.PHONY: format lint test docs

GOFILES := $(shell find . -type f -name '*.go' -not -path './vendor/*')

format:
	gofumpt -l -w $(GOFILES)

lint:
	GOTOOLCHAIN=go1.22.4 golangci-lint run ./...

test:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
	pct=$$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","");print $$3}'); \
	if [ $${pct%.*} -lt 83 ]; then echo "::error::coverage < 83%" && exit 1; fi

docs:
	@git ls-files '*.md' | xargs -r sed -i 's/[ \t]*$$//' && git diff --exit-code || true
