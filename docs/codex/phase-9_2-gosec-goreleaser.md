<!--
AI-Chat-CLI • Codex Prompt
Phase 9.2 – Gosec Pre-flight & Goreleaser Cleanup (Go 1.24-pinned)
Save as docs/codex/phase-9_2-gosec-goreleaser.md
Author: Jamal Al-Sarraf <jalsarraf0@gmail.com>
-->

# Phase 9.2 Prompt – **Deterministic Security Scan + Goreleaser Modernisation** 🛡️🚀
_Target coverage ≥ 92 % (stretch 95 %)_

---

## 0 • Immutable golden rule
> **Go tool-chain is permanently pinned to 1.24.x.**
> All helper scripts *must* detect and enforce this version.

---

## 1 • Objectives
1. **Install _gosec_ before any other step**, both locally and in CI, via a single cross-platform bootstrap.
2. **Replace the legacy `TOOLS_CRITICAL` + `install_tools()`** with the user-supplied implementation, supporting Bash **and** PowerShell.
3. **Modernise `goreleaser.yml`** — remove deprecated keys so snapshot builds are warning-free.
4. Keep a **single GitHub Actions workflow**; only extend, never add new files.

---

## 2 • Updated tooling bootstrap

### 2.1 `scripts/bootstrap.sh` (POSIX)
```bash
$(sed -n '1,30p' scripts/bootstrap.sh)
```

### 2.2 `scripts/bootstrap.ps1` (PowerShell 5+)
```powershell
$(sed -n '1,24p' scripts/bootstrap.ps1)
```

---

## 3 • Goreleaser cleanup (`goreleaser.yml`)
Removed deprecated `brews.folder` and `brews.tap` keys and added a modern Homebrew block with a simple test hook.

### Snapshot test
Run `goreleaser release --snapshot --clean --skip=publish --skip=docker --skip=sign` and verify **no deprecation warnings** appear.

---

## 4 • CI workflow hooks (`.github/workflows/ci.yml`)
The first step in each job bootstraps tools:
```yaml
- name: Bootstrap critical tools
  if: runner.os != 'Windows'
  run: ./scripts/bootstrap.sh install_tools

- name: Bootstrap critical tools (Windows)
  if: runner.os == 'Windows'
  shell: pwsh
  run: ./scripts/bootstrap.ps1
```
The *security-scan* stage now executes `make security-scan`.

---

## 5 • Local developer workflow
```bash
./scripts/bootstrap.sh install_tools
make security-scan            # → gosec ./...

# PowerShell
./scripts/bootstrap.ps1
make security-scan
```

---

## 6 • Acceptance checklist
- [ ] Bootstrap succeeds on Bash (Linux/macOS) and PowerShell (Windows).
- [ ] `gosec ./...` reports **0 critical issues**.
- [ ] `goreleaser --snapshot` completes **without deprecation warnings**.
- [ ] CI green across OS matrix, with no ad-hoc `go install` commands.
- [ ] Scripts exit if Go ≠ 1.24.x.
- [ ] Diff confined to scripts, `goreleaser.yml`, and CI YAML; no new workflow files.
- [ ] Docs updated; commits signed **Jamal Al-Sarraf** and ≤ 72 chars.

---

MIT © 2025 Jamal Al-Sarraf
