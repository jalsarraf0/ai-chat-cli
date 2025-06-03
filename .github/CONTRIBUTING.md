# Contributing to AI-Chat-CLI

Thank you for your interest in contributing! Follow these steps to submit a pull request that passes CI with at least 82% test coverage.

## Quick Demo

```bash
# Fork the repository and clone your fork
$ git clone https://github.com/<you>/ai-chat-cli.git
$ cd ai-chat-cli
$ git checkout -b codex/my-change

# Hack and test
$ make format lint test

# Push and open a PR
$ git push -u origin HEAD
```

A new contributor should be able to complete the above in under **10 minutes**.

## Branching

Always create branches as `codex/<feature>` from the latest `dev` branch.

## Coding Style

- Go 1.24.3
- Run `make format` before committing. It wraps `gofumpt` and `goimports`.
- Tests must be table-driven and use `t.Parallel()`.
- Target coverage is 82% or higher. Run `make test` to check.

## Commit Messages

Write imperative present-tense messages with a short (<72 chars) subject line.

## Pull Requests

1. Ensure all tests pass and coverage ≥82%.
2. Run `gofumpt`/`goimports` (`make format`).
3. Update documentation if you change a public interface.
4. Fill out every item in the PR checklist.

For more details, read `docs/MAINTAINERS.md` for contact information.
