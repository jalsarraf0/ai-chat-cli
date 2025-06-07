# Codex Prompt – **Always-Online & SUPREME-Release Add-On**
*(Save or overwrite as `docs/codex/addons/always-online.md` and include it in every change-set or PR.)*

This document **augments** the baseline Codex Rules.
If any statement here conflicts with the base rules, obey the *stricter* instruction or raise a question.

---
## ▍0. Internet Access Policy 💡
1. **Assume continuous Internet connectivity** in all scripts, helper tools and CI jobs.
2. Missing / outdated Go deps → run:
   ```bash
   go get -u ./...
   go mod tidy -e -v
   ```
   Missing CLI tool (e.g. gosec, trivy) → install via:
   ```bash
   go install <module>@latest
   ```
   Fallback: fetch binary from GitHub Releases.

## ▍1. Coverage Gate ≥ 93 % 📈
Pipeline minimum is 90 %; you must deliver ≥ 93 % for a 3-pp buffer.
```bash
go test -race -covermode=atomic -coverprofile=coverage.out ./...
pct=$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","");print $3}')
[[ ${pct%.*} -ge 93 ]] || { echo "::error::coverage < 93 %"; exit 1; }
```
For slice testing, merge first:
```bash
gocovmerge cover*.out > coverage.out
```

## ▍2. Full Security Suite 🛡️
Every PR must pass all checks with zero HIGH / CRITICAL findings.

| Tool | Command |
|------|---------|
| **Gosec** | `gosec ./...` |
| **Govulncheck** | `govulncheck ./...` |
| **Trivy (FS)** | `trivy fs --exit-code 1 --severity HIGH,CRITICAL .` |
| **License Audit** | `addlicense -check $(go list -f '{{.Dir}}' ./...)` |

## ▍3. CI / CD Pipeline Is Immutable 🔒
Do not create, edit, or remove workflow files (.github/workflows/**) or any other CI/CD configuration.
The existing pipeline is authoritative and frozen.
If you believe a pipeline change is essential, escalate to the maintainers instead of committing modifications.

## ▍4. Conflict-Free Pull Requests 🔀
Rebase your branch on dev before each push:
```bash
git fetch origin
git rebase origin/dev
```
Resolve conflicts, rerun full suite, regenerate docs, then:
```bash
git push --force-with-lease
```

## ▍5. Quick Pre-PR Checklist ✅
- [ ] `go vet ./...` passes.
- [ ] `go test -race ./...` shows **≥ 93 %** coverage.
- [ ] Security suite (Gosec + Govulncheck + Trivy) green.
- [ ] `go mod tidy -e` leaves no diff.
- [ ] Docs regenerated (`make docs`) with no noise.
- [ ] `shellcheck scripts/*.sh` passes (ignore **SC2086** where noted).
- [ ] Branch rebased — **no conflicts**.
- [ ] Commit message ≤ 72 chars, imperative; PR labelled **codex-generated**.

> **Always stay online, hit ≥ 93 % coverage, pass every security gate, and ship cleanly — without touching the CI/CD pipeline.**
