# Phase 0 Prompt – Repository Bootstrap 🚀

You are **OpenAI Codex** operating as a senior Go release engineer.
Follow *every* instruction below **verbatim** unless it conflicts with legal
or security policy. When conflicts arise, **raise a question** instead of
silently diverging.

---

## 🎯 Objectives

1. **Initialize modules**

   ```bash
   go mod init github.com/jalsarraf0/ai-chat-cli
   ```

2. **Scaffold directories**

   ```
   cmd/
   internal/
   pkg/
   scripts/
   ```

3. **Add development guard-rails**

   - Pre-commit hook running `gofumpt -l -w` and `golangci-lint run`.
   - `.editorconfig` & `.gitattributes` for cross-OS consistency.

4. **Bootstrap docs**

   - `README.md` with CI badge:     `https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml/badge.svg`
   - Short project description (< 40 words).

5. Push all work to branch **`phase0/bootstrap`** with a single signed commit
   titled **“feat(phase-0): repository bootstrap”**.

---

## ✅ Acceptance Criteria

| Gate | Requirement |
|------|-------------|
| CI   | Lone workflow `ci.yml` passes (coverage-gate → gosec → docs). |
| Coverage | ≥ 83 % even for placeholder tests (table-driven, `t.Parallel()`). |
| Conflicts | Branch rebased onto `dev`; **no `<<<<`/`>>>>` markers**. |
| Lint | `golangci-lint run ./...` exits 0. |
| Docs | `make docs` produces no git diff noise. |

---

## 🚦 Ground Rules (verbatim excerpt)

> **Internet is always on.**
> If a dependency or tool is missing, you **MUST**
> `go get -u ./... && go mod tidy -e -v` or
> `go install <module>@latest`.

> **Coverage Gate ≥ 83 %** – the pipeline enforces 80 %; always ship 83 %+.

> **Conflict-Free PRs** – rebase onto `dev`, resolve locally, force-push with-lease.

> **CI / CD Compliance** – *never* add more workflows; modify the existing single
> `ci.yml` only if extending jobs fan-out. Stages order is immutable.

_Check-list before opening a PR_
- [ ] `go vet ./...` clean
- [ ] `go test -race ./...` ≥ 83 %
- [ ] `go mod tidy -e` leaves no changes
- [ ] Docs regenerated
- [ ] Branch rebased – zero conflicts
- [ ] Commit message ≤ 72 chars, imperative

---

## 🛠️ Commands Cheat-Sheet

```bash
# Format + lint + test
make format lint test

# Install hooks
pre-commit install

# Generate docs
make docs
```

---

## 🔒 License

MIT © 2025 Jamal Al-Sarraf
