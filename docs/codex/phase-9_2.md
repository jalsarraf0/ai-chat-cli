<!--
AI-Chat-CLI • Codex Prompt
Phase 9.2 – Gosec Pre-flight & Release Modernisation (CI-passing v6)
Save as docs/codex/phase-9_2.md
Author: Jamal Al-Sarraf <jalsarraf0@gmail.com>
-->

# Phase 9.2 – **Deterministic Security Scan + Release Modernisation** 🛡️🚀
_Target: CI green on Linux · macOS · Windows, coverage ≥ 92 %_

## ⚠️ Why the Windows job failed


`go install` expects build flags before the package. The previous
PowerShell script placed `-trimpath` after the package, causing Go to
interpret it as a second module. The corrected line uses
`go install -trimpath $pkg`.

## 🔑 Why CI failed

`make security-scan` was triggered but the **Makefile** lacked that
target. GNU make exited with “No rule to make target 'security-scan'”.

## 1 • Makefile patch

Add this snippet (merge if needed):

```make
# ---------------------------------------------------------------------------
# Security analysis
# ---------------------------------------------------------------------------
.PHONY: security-scan

security-scan: ## Run gosec static analysis
GOFLAGS='-trimpath' gosec ./...
```

## 0 • Immutable rule
**Go 1.24.x** is the only supported tool-chain everywhere.

## 2 • Deliverables

| Item | Description |
|------|-------------|
| **Bootstrap file** | `scripts/bootstrap-tools.sh` installs gofumpt, staticcheck and gosec using Go 1.24. |
| **PowerShell twin** | `scripts/bootstrap.ps1` provides the same on Windows. |
| **CI integration** | Add a Bootstrap step after every `actions/setup-go@v5` call. |
| **Goreleaser clean-up** | Remove deprecated `brews.folder` / `brews.tap`. |
| **Keyless Cosign + SBOM** | The release job installs cosign then signs and attaches an SBOM for every archive. |
| **Docs** | Document the keyless flow in `docs/security.md` and add a README badge. |

## 3 • Reference `scripts/bootstrap-tools.sh`
```bash
$(sed -n '1,25p' scripts/bootstrap-tools.sh)
```

PowerShell uses a matching `bootstrap.ps1`.

## 4 • CI Snippet
```yaml
- uses: actions/setup-go@v5
  with: { go-version: ${{ env.GO_VERSION }} }
- name: Bootstrap critical tools
  run: ./scripts/bootstrap-tools.sh
```

## 5 • Release signing
After `goreleaser release --clean --skip=publish`, the job executes:
```bash
cosign sign --yes dist/*.tar.gz
cosign attach sbom --yes dist/*.tar.gz
```
Permissions `id-token: write` and `contents: read` enable keyless signing.
