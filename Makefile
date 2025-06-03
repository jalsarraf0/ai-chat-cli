
.PHONY: build test completion man release lint

build:
go build -v -o bin/ai-chat ./cmd/ai-chat

test:
go test -race ./...

completion: build
bin/ai-chat completion bash > dist/ai-chat.bash

man: build
bin/ai-chat man --dir dist/man

release: build completion man

lint:
golangci-lint run

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

