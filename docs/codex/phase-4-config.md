<!--
AI‑Chat‑CLI • Codex Prompt
Phase 4 – Configuration Layer
Save this file as docs/codex/phase‑4‑config.md
Author: Jamal Al‑Sarraf <jalsarraf0@gmail.com>
-->

# Phase 4 Prompt – Configuration Layer 🗄️
*CLI gains a flexible, file‑backed settings system — CI coverage ≥ 88 %.*

---

## 🎯 Objectives

| # | Deliverable | Details |
|---|-------------|---------|
| 1 | **Config package (`pkg/config`)** | Wrap **spf13/viper** with helper functions:<br>`Load()`, `Save()`, `GetString(key)`, `Set(key,val)`. Default path `~/.config/ai-chat/config.yaml` on Unix & `%APPDATA%\ai-chat\config.yaml` on Windows. |
| 2 | **CLI sub‑command `ai-chat config`** | Sub‑commands:<br>`show` (print current YAML), `set <key> <value>`, `edit` (opens `$EDITOR`, fallback to `vi`/`notepad`). |
| 3 | **Environment & flag overrides** | • ENV prefix `AICHAT_` (e.g., `OPENAI_API_KEY`).<br>• Global `--config <file>` flag to override path. |
| 4 | **Validation hooks** | Custom validators (e.g., API key non-empty, model name allowed). Fail fast with helpful error. |
| 5 | **Docs & examples** | Update README + add `docs/config.md` detailing keys. |

---

## 🛠️ Implementation Notes

### Config paths

```go
func defaultPath() string {
if runtime.GOOS == "windows" {
return filepath.Join(os.Getenv("APPDATA"), "ai-chat", "config.yaml")
}
return filepath.Join(xdg.Home, ".config", "ai-chat", "config.yaml")
}
```

Use `os/user` + `os.Getenv("XDG_CONFIG_HOME")` as fallback.

### Command wiring

```go
rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default auto)")
rootCmd.AddCommand(configCmd) // from cmd/config.go
```

---

## 🗂️ CI Workflow delta (`.github/workflows/ci.yml`)

* **Linux** job → `runs-on: [self-hosted, linux]`
* **macOS** job → `runs-on: macos-latest`
* **Windows** job → `runs-on: windows-latest`

(Reuse shared steps from Phase 3; only the coverage gate threshold changes.)

### Coverage gate (Makefile)

```make
coverage-gate:
@pct=$$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","" );print $$3}'); \
if [ $${pct%.*} -lt 88 ]; then \
echo "::error::coverage < 88% (got $${pct}%)"; exit 1; fi
```

---

## 🧪 Tests (target ≥ 90 % in new package, ≥ 88 % overall)

* **pkg/config/config_test.go** — table‑driven tests for `Load`, `Set`, env override.
* **cmd/config_test.go** — uses `cobra` test utilities to run `show`, `set`, `edit --dry-run`.

---

## ✅ Acceptance Criteria

| Gate | Requirement |
|------|-------------|
| CI   | Linux (self‑hosted), macOS (`macos-latest`), Windows (`windows-latest`) pass. |
| Coverage | **≥ 88 % overall** (coverage gate job). |
| Lint | `golangci-lint run ./...` zero issues. |
| Race | `go test -race ./...` passes on Linux & macOS. |
| Docs | README & `docs/config.md` updated, man pages refreshed. |
| Merge | Signed commit by **Jamal Al‑Sarraf <jalsarraf0@gmail.com>**, conflict‑free. |

---

## ✉️ Commit message skeleton

```
feat(phase‑4): Viper‑based configuration layer

* pkg/config with Load/Save/Set + env overrides
* ai-chat config {show,set,edit}
* validators for API key & model
* tests raise total coverage to 88 %
```

---

MIT © 2025 Jamal Al‑Sarraf <jalsarraf0@gmail.com>
