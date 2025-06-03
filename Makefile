GO ?= go
.PHONY: test-shell

test-shell:
$(GO) test ./internal/shell -cover -race
