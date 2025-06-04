<!--
AI‑Chat‑CLI • Codex Prompt
Phase 2 – Cross‑Shell Runtime
Save this file as docs/codex/phase‑2‑cross‑shell.md
Author: Jamal Al‑Sarraf <jalsarraf0@gmail.com>
-->

# Phase 2 Prompt – Cross-Shell Detection & Runtime 🐚 (Self-Hosted CI/CD)

You are **OpenAI Codex** acting as the project’s senior Go release engineer.
Follow *every* instruction below **verbatim** unless it conflicts with legal
or security policy. When conflicts arise, **raise a question** instead of
silently diverging.

---

## 🎯 Objectives

1. **Shell detection helper**

   * Create `internal/shell/detect.go` exposing:

     ```go
     // Detect returns the active shell kind for the current OS/user.
     func Detect() Kind
     ```

   * Inputs: `$SHELL` (Unix), `$ComSpec` (Windows), process parent tree.
   * Supported kinds: **Bash, Zsh, Fish, PowerShell, Cmd**.
   * Unit tests cover Linux/macOS + Windows (`go test -tags windows`).
     Use `t.Setenv` to inject env vars.

2. **Unified command runner**

   * Add `internal/shell/runner.go`:

     ```go
     func Run(ctx context.Context, cmd string, args ...string) (stdout, stderr string, err error)
     ```

   * Internally chooses `exec.CommandContext(shellBin, "-c", …)` or
     `powershell -NoProfile -Command …` on Windows.
   * Return rich `*exec.ExitError` when non-zero.

3. **Cross-compile paths**

   * Use build tags: `_unix.go` vs `_windows.go`.
   * `go test ./...` on self-hosted runners for Linux, macOS, and Windows.
   * Mock `exec.LookPath` for Windows inside tests.

4. **Self-Hosted CI pipeline**

   * Modify `.github/workflows/ci.yml` so every job runs exclusively on self-hosted runners:
     - Linux: `runs-on: [self-hosted, linux]`
     - macOS: `runs-on: [self-hosted, macOS]`
     - Windows: `runs-on: [self-hosted, windows]`
   * Stages `setup → lint → test → coverage-gate → gosec → docs` remain intact.
   * Caches live under `/var/cache/ai-chat-cli/go` on each runner host.

   ```yaml
   jobs:
     test-linux:
       runs-on: [self-hosted, linux]
       steps: *steps
     test-macos:
       runs-on: [self-hosted, macOS]
       steps: *steps
     test-windows:
       runs-on: [self-hosted, windows]
       steps: *steps
   ```

   *(Provide reusable `steps` alias with checkout, Go setup, `make lint test`.)*

5. **Make targets**

   | Target            | Description                                           |
   |-------------------|-------------------------------------------------------|
   | `make shell-test` | runs `go test -run TestShell ./internal/shell/...`    |
   | `make cross`      | shortcut for `GOOS=windows GOARCH=amd64 make build`   |

6. **CLI integration**

   * Root command auto-detects shell on startup and logs it with `--verbose`.
   * `ai-chat ping` prints detected shell in debug mode.

7. **Docs**

   * README “Quick Start” – note cross-shell support and self-hosted CI.
   * Regenerate man pages (`make man`) with new flags.

8. **Git workflow**

   * Branch `phase2/cross-shell` based on `dev`.
   * Single signed commit: “feat(phase-2): cross-shell runtime abstraction”.
   * Rebase onto `dev`; resolve conflicts locally before push.

---

## ✅ Acceptance Criteria

| Gate | Requirement |
|------|-------------|
| CI   | Self-hosted pipeline passes on linux; optional macOS & Windows jobs pass when available. |
| Coverage | ≥ 83 % overall; shell package ≥ 90 %. |
| Lint | `golangci-lint run ./...` exits 0 (`gocritic`, os-specific). |
| Race | `go test -race ./...` passes on Unix runners. |
| Docs | `make man && make docs` produce no git diff noise. |
| UX   | `ai-chat --verbose ping` prints `shell=<detected>` on runner host. |
| Merge | PR rebased; **no conflict markers**. |

---

## 🛠️ Commands Cheat-Sheet

```bash
# Detect shell locally
go run ./cmd/ai-chat --verbose ping

# Run shell tests on Windows mode
go test -tags windows ./internal/shell/...

# Build cross-platform
GOOS=windows GOARCH=amd64 go build -o bin/ai-chat.exe ./cmd/ai-chat
```

---

## 🔒 License

MIT © 2025 Jamal Al-Sarraf <jalsarraf0@gmail.com>
