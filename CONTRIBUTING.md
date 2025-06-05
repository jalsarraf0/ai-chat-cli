# Contributing to ai-chat-cli

## Prerequisites
* Go 1.24.x _(strictly pinned)_
* GNU Make ≥\_4.3 _(Linux/macOS)_
* Bash 5, Docker (for integration)
* PowerShell 7 _(Windows)_

## Getting Started
```bash
git clone https://github.com/<you>/ai-chat-cli.git
cd ai-chat-cli
make bootstrap   # macOS/Linux
./scripts/bootstrap.ps1  # Windows
```

## Testing
* **Full suite:** `make unit`
* **Sharded:** `CASE=2/4 ./scripts/test-slice.sh`
* **Benchmarks:** `make bench`

## Branch & PR
* Work off `dev` → `feature/* | fix/*`
* Squash-merge with Conventional Commit messages (`feat:` / `fix:` / `ci:` …).

## Release
Tag `vX.Y.Z`; CI triggers GoReleaser.
