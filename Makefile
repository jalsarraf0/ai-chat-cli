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
