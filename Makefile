BINARY=ai-chat
BIN_DIR=bin
GOFLAGS ?=

.PHONY: build test lint format clean

build:
	mkdir -p $(BIN_DIR)
	GOFLAGS="$(GOFLAGS)" go build -o $(BIN_DIR)/$(BINARY) ./cmd/ai-chat

test:
	GOFLAGS="$(GOFLAGS)" go test ./...

lint:
	golangci-lint run

format:
	gofmt -w $(shell find . -name *.go)

clean:
	rm -rf $(BIN_DIR)
