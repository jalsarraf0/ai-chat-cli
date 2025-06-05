<!--
AI‑Chat‑CLI • Codex Prompt
Phase 8 – Go 1.24‑pinned, Lint‑compatible 🚀
Save this file as docs/codex/phase‑8‑release.md
Author: Jamal Al‑Sarraf <jalsarraf0@gmail.com>
-->

# Phase 8 Prompt – Go 1.24‑pinned, Lint‑compatible 🚀
*CI matrix: **Linux [self‑hosted, linux] · Windows (windows‑latest) · macOS (macos‑latest)** — global coverage gate **≥ 92 %***

---

## 🔑 Golden rule
**Go toolchain is permanently pinned to 1.24.x.**
All tooling must compile and run with this version.

---

## 1 • Environment block (every workflow job)
```yaml
env:
  GO_VERSION: "1.24.x"
  # install latest HEAD of golangci-lint because official tags
  # are built with Go ≤ 1.23 and cannot parse 1.24 export data.
  GOLANGCI_INSTALL_VERSION: "latest"
```

## 2 • GolangCI‑Lint installation (HEAD build)
```yaml
- name: Install golangci-lint (Go 1.24 compatible)
  shell: bash
  run: |
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
      | sh -s -- -b "$(go env GOPATH)/bin" "${{ env.GOLANGCI_INSTALL_VERSION }}"
    golangci-lint --version
```
*The `install.sh ... latest` path builds from the master branch with the **current Go toolchain**, ensuring compatibility with 1.24.*

## 3 • x64‑only build matrix (unchanged)
```yaml
builds:
  - id: cli
    goos: [linux, windows, darwin]
    goarch: [amd64]
    env: { CGO_ENABLED: 0 }
```

## 4 • Makefile & coverage gate (still 92 %)
```make
lint:
golangci-lint run ./...

unit:
go test -race -covermode=atomic -coverprofile=coverage.out -tags unit ./...
@$(MAKE) coverage-gate
```

## 5 • README & Credits refresh
* **Badges**: CI ✅ · Coverage 📈 · Go 🐹 1.24 · Security 🔒 · License 📜 · Release 🏷️
* **docs/credits.md** auto‑generated via `scripts/gen_credits.go`.

## 6 • Acceptance checklist
* Workflows run **Go 1.24.x** on Linux, macOS, Windows‑latest.
* `golangci-lint --version` prints a commit hash built with Go 1.24; no `unsupported version: 2` errors.
* All CI gates pass; coverage ≥ 92 %.
* Release pipeline publishes signed x64 artefacts.
* README pretty, credits page added.
* Commit signed **Jamal Al‑Sarraf <jalsarraf0@gmail.com>**.

---

MIT © 2025 Jamal Al‑Sarraf
