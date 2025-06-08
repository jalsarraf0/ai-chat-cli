# Codex Prompt – Authoritative CI/CD Pipeline (Gated)

*Save this file as `gated-ci-pipeline-codex-prompt.md` and paste its contents into an OpenAI Codex session.  Codex must return a pull‑request that **replaces** the existing workflows with the four‑gate pipeline.*

---

## Required Gate Order  

```
Lint  ─┐
       ↓
Test Matrix ─┐
             ↓
Security ─┐
          ↓
Release (snapshot) ──► artefacts (.tar.gz · .deb · .rpm)
```

If **any** gate fails, subsequent jobs **must not** run and the workflow exits non‑zero.

---

## Key Rules  

| Rule | Requirement |
|------|-------------|
| Runners | `[self-hosted, linux]` **or** `ubuntu-latest` (Codex decides) |
| Toolchain | Go 1.23 / 1.24 / 1.25, latest `golangci-lint v2`, `govulncheck`, `gosec v2`, `osv-scanner`, `Trivy` |
| Coverage | Pipeline fails if total < **93 %** |
| Security | Pipeline fails on any **HIGH / CRITICAL** vuln |
| Docs | `hugo --minify` must still work; no Node/npm |
| Artefacts | Name pattern `ai-chat-cli_<tag>_linux_<arch>.tar.gz` plus `.deb` & `.rpm` for amd64 |

---

## Workflow Skeleton (Codex must implement)

### `.github/workflows/ci.yml`

```yaml
name: CI

on:
  pull_request: {}
  push:
    branches: [ dev, main ]

jobs:

  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: {go-version-file: 'go.mod'}
      - name: Install linter
        run: go install github.com/golangci/golangci-lint/cmd/golangci-lint/v2@latest
      - run: golangci-lint run --timeout 5m

  test-matrix:
    needs: lint
    runs-on: ubuntu-latest
    strategy:
      matrix: {go: ['1.23','1.24','1.25']}
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: {go-version: ${{ matrix.go }}}
      - run: |
          go test -race -covermode=atomic -coverprofile=coverage.out ./...
          pct=$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","");print $3}')
          echo "coverage=${pct}%" 
          if [[ ${pct%.*} -lt 93 ]]; then
             echo "coverage ${pct}% < 93%" && exit 1
          fi

  security:
    needs: test-matrix
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: {go-version-file: 'go.mod'}
      - name: Install scanners
        run: |
          go install golang.org/x/vuln/cmd/govulncheck@latest
          go install github.com/securego/gosec/v2/cmd/gosec@latest
          go install github.com/google/osv-scanner/cmd/osv-scanner@latest
          curl -fsSL https://github.com/aquasecurity/trivy/releases/latest/download/trivy_0_Linux-64bit.tar.gz | tar -xzO trivy > /usr/local/bin/trivy && chmod +x /usr/local/bin/trivy
      - run: govulncheck ./...
      - run: gosec ./...
      - run: osv-scanner .
      - run: trivy fs --exit-code 1 --severity HIGH,CRITICAL .
```

### `.github/workflows/release.yml`

```yaml
name: Release

on:
  push:
    tags: [ 'v*' ]
  workflow_dispatch:

jobs:
  release:
    needs: security
    runs-on: ubuntu-latest
    permissions: {contents: write}
    strategy:
      matrix: {arch: [amd64, arm64]}
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: {go-version-file: 'go.mod'}

      - name: Build binary
        env:
          GOARCH: ${{ matrix.arch }}
        run: |
          mkdir -p dist/linux_${GOARCH}
          GOOS=linux go build -trimpath -ldflags "-s -w" -o dist/linux_${GOARCH}/ai-chat-cli ./cmd/ai-chat-cli

      - name: Package tar.gz
        env:
          TAG: ${{ github.ref_name }}
          ARCH: ${{ matrix.arch }}
        run: tar -C dist/linux_${ARCH} -czf dist/ai-chat-cli_${TAG}_linux_${ARCH}.tar.gz ai-chat-cli

      - name: Goreleaser snapshot
        if: ${{ matrix.arch == 'amd64' }}
        run: |
          curl -sSfL https://install.goreleaser.com/github.com/goreleaser/goreleaser@latest | sh -s -- -b /usr/local/bin
          goreleaser release --clean --snapshot --single-target --id deb --id rpm --dist dist

      - uses: softprops/action-gh-release@v2
        with:
          files: dist/*
```

---

## Codex Deliverables  

* Replace existing workflows with the above, adapting paths & binary names as needed.  
* Update any scripts/Makefile so coverage ≥ 93 %.  
* Commit messages in imperative mood (`ci: add gated pipeline`).  
* PR description must state “All gates pass; ready to merge”.

If Codex cannot make every gate pass, it must refuse with:

```
❌ Unable to satisfy all CI gates.
```

---

### ━━━  End of prompt – Codex, build the gated pipeline  ━━━
