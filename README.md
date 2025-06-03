
[![CI](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml)

ai-chat CLI provides a simple command line interface for AI chat interactions.

# AI‑Chat‑CLI

[![CI – AI‑Chat‑Fedora42](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml)

A **cross‑shell, cross‑platform ChatGPT‑style CLI/TUI** written in Go 1.24.  
Runs unmodified on **Linux (bash/zsh/fish)**, **macOS (zsh)**, and **Windows (PowerShell)**.

> **Status:** early bootstrap (`v0.1.0‑alpha`).  Follow the [roadmap](#roadmap).

---

## Features (roadmap snapshot)

| Phase | Milestone                                  | Status |
|-------|--------------------------------------------|--------|
| 0     | Repo bootstrap, CI badge                   | ✅ done |
| 1     | Core CLI (`version`, `ping`, completions)  | ✅ done |
| 2     | Shell detection & command runner           | 🔨 in-progress |
| 6     | Bubble Tea TUI chat window                 | ⏳ pending |
| 12    | v1.0 multi‑platform release                | 🚀 future |

See full details in `docs/roadmap.md` (coming soon).

---

## Quick Start

```bash
# Clone & build
git clone https://github.com/jalsarraf0/ai-chat-cli.git
cd ai-chat-cli
make build

# Run
./bin/ai-chat version
./bin/ai-chat ping --debug
```

> **Prerequisite:** Go 1.24.3+ (already bundled in most dev containers).

---

## Shell Completion

| Shell     | Command to enable |
|-----------|-------------------|
| bash/zsh  | `ai-chat completion bash > /etc/bash_completion.d/ai-chat` |
| fish      | `ai-chat completion fish > ~/.config/fish/completions/ai-chat.fish` |
| PowerShell| ``ai-chat completion ps1 | Out-File $PROFILE`` |

---

## Configuration

```bash
# first-time setup
ai-chat config set api_key $AI_CHAT_API_KEY
ai-chat config show
```

*Config precedence:* `flags` ▶ `env vars` ▶ `~/.config/ai-chat/config.toml`.

---

## Development

```bash
# Format, lint, test
make format lint test

# Generate docs (CI does this automatically)
make man completion
```

The project enforces **≥ 80 % coverage** and runs `gosec` in CI.

---

## Contributing

1. Read `.codex-rules.txt` and `AGENTS.md` to understand the workflow rules.
2. Create a branch `codex/your-feature`.
3. Open a draft PR; ensure CI is green before requesting review.

---

## License

MIT © 2025 Jamal Al‑Sarraf

