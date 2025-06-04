<!--
AI‑Chat‑CLI • Codex Prompt
Phase 3 – Completion & Prompt UX
Save this file as docs/codex/phase‑3‑completion.md
Author: Jamal Al‑Sarraf <jalsarraf0@gmail.com>
-->

# Phase 3 Prompt – Shell Completion & Prompt UX ✨ (Self‑Hosted CI/CD)

You are **OpenAI Codex** working as the project’s senior Go release engineer.
Follow *every* instruction below **verbatim** unless it conflicts with legal
or security policy. When conflicts arise, **raise a question** instead of
silently diverging.

> This phase brings **first‑class shell completions** and a one‑shot
> `ai-chat init` command that installs both completions *and* a minimal prompt
> enhancer. All CI jobs remain **self‑hosted only**.

---

## 🎯 Objectives

1. ### Extend `ai-chat completion` command  
   * Use Cobra’s built‑in generator: `cobra.GenBashCompletion`, `GenZshCompletion`, etc. citeturn0search0turn0search3  
   * Support `bash`, `zsh`, `fish`, `powershell`.  
   * Add `--out <file>` flag (default `dist/completion/<shell>`).  
   * Ensure target directory exists (`os.MkdirAll`).  
   * Usage examples in help text.  
   * Unit‑test file creation with `t.TempDir()`.

2. ### Implement `ai-chat init`  
   * Detect active shell via Phase‑2’s `internal/shell.Detect()`.  
   * Internally call `ai-chat completion <shell> --out ~/.ai-chat/completion.<ext>` and append `source …` ( or `. "$HOME/.ai-chat/completion.ps1"` ) to the proper RC/profile file.  
   * Flags: `--dry-run`, `--no-prompt`.  
   * Success message: **✅ Completions installed for <shell>**.  
   * Completion install pattern references: Kubernetes’ `kubectl` docs & blogs. citeturn0search5turn0search7

3. ### Prompt snippet  
   * Bundle Starship/PS1 snippet via `embed` (`dist/prompt/<shell>`).  
   * Shows green `(ai)› ` when directory contains `.ai-chat.yml`.  

4. ### Tests & Coverage  
   * New tests for `cmd/completion`, `cmd/init`; stub FS with `t.TempDir`.  
   * Target **≥ 88 %** coverage. Reference article on CLI testing. citeturn0search12

5. ### Make targets  

   | Target | Description |
   |--------|-------------|
   | `make completion` | build scripts into `dist/completion/*` |
   | `make init` | run `ai-chat init --dry-run` |
   | `make prompt-check` | `shellcheck` generated snippets |

6. ### Self‑Hosted CI adjustments  
   * Linux/macOS/Windows matrix.  
   * Add shellcheck step for Linux job.  
   * Coverage gate raised to 88 %.

7. ### Docs  
   * Update README “Shell Completion” table; add prompt section. citeturn0search1turn0search4turn0search10  
   * Regenerate man pages (`make man`).  

8. ### Git workflow  
   * Branch `phase3/completion-ux` → `dev`.  
   * Signed commit message: **“feat(phase‑3): completions & prompt UX”**.  

---

## ✅ Acceptance Criteria

| Gate | Requirement |
|------|-------------|
| CI   | Self‑hosted pipeline green on all OSes. |
| Coverage | ≥ 88 %. |
| Lint | `golangci-lint`, `shellcheck` zero issues. |
| Race | `go test -race ./...` passes. |
| Merge | PR rebased, signed **Jamal Al‑Sarraf**. |

---

## 🔒 License

MIT © 2025 Jamal Al‑Sarraf <jalsarraf0@gmail.com>
