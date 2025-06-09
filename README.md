# AI‑Chat‑CLI

[![Coverage](https://img.shields.io/badge/coverage-93%25-brightgreen)](./)
[![Go](https://img.shields.io/badge/go-1.24.x-00ADD8?logo=go)](https://go.dev/doc/go1.24)
[![Security](https://img.shields.io/badge/security-passing-brightgreen)](./)
[![Cosign (OIDC)](https://img.shields.io/badge/cosign%20(OIDC)-verified-brightgreen)](https://github.com/sigstore/cosign)
[![Security Scan](https://img.shields.io/badge/security%20scan-100%25%20clean-brightgreen)](./)
[![Release](https://img.shields.io/badge/release-no%20releases%20or%20repo%20not%20found-lightgrey)](./)
[![Container](https://img.shields.io/badge/container-ghcr.io%2Fjalsarraf0%2Fai--chat--cli-blue)](https://ghcr.io/jalsarraf0/ai-chat-cli)
[![License](https://img.shields.io/badge/license-repo%20not%20found-lightgrey)](./)

> **ai‑chat‑cli** is a lightweight, cross‑platform command‑line interface for interacting with GPT‑style large language models (LLMs).
> Written in pure **Go**, it streams answers in real‑time, keeps your history, and supports an extensible plug‑in system — all in a single ≈ 6 MiB binary.

## Quick Setup
1. Download a release or clone the repo and run `./setup.sh`.
2. `export OPENAI_API_KEY="sk-..."`
3. Run `ai-chat` and start typing.


---

## 📚 Table of Contents
1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Installation](#installation)
   * [Packages](#packages)
   * [Build from Source](#build-from-source)
4. [Quick Start](#quick-start)
5. [Commands](#commands)
6. [Configuration](#configuration)
7. [Plug‑ins](#plug-ins)
8. [Testing & CI](#testing--ci)
9. [Development](#development)
10. [Contributing](#contributing)
11. [Security](#security)
12. [License](#license)
13. [Changelog](#changelog)
14. [Acknowledgements](#acknowledgements)

---

## Overview
`ai-chat` wraps the OpenAI, Azure OpenAI and Ollama APIs behind a consistent CLI.
Key features:

| Feature | Description |
|---------|-------------|
| 📨 **Streaming chat** | Low‑latency updates with ANSI markdown rendering |
| 🗂 **Persistent history** | All chats stored in SQLite and searchable |
| 🔌 **Plug‑in framework** | Drop shell scripts in `~/.config/ai-chat/plugins` |
| 🔒 **Keychain integration** | Secrets stored via `pass`, macOS Keychain or Windows DPAPI |
| 🖇 **Exporters** | Save replies to clipboard, Markdown, HTML, JSON, PDF |
| ⚙️ **Embeddable** | Public Go API (`pkg/chat`) for your own apps |

---

## Architecture

```
client (cobra CLI) ─┐
                    ├─► core engine ──► provider interface ──► OpenAI / Azure / Ollama
plugins (shell) ────┘
                             │
                             └──► SQLite history
```

---

## Installation

### Packages

| Package | Command |
|---------|---------|
| **tar.gz** | `tar -xzf ai-chat_linux_amd64.tar.gz && sudo mv ai-chat /usr/local/bin` |
| **DEB** | `sudo dpkg -i ai-chat_<ver>_amd64.deb` |
| **RPM** | `sudo rpm -Uvh ai-chat_<ver>_amd64.rpm` |
| **Homebrew** | `brew install jalsarraf0/tap/ai-chat` |
| **Scoop** | `scoop bucket add jalsarraf0 https://github.com/jalsarraf0/scoop-bucket && scoop install ai-chat` |

### Build from Source
```bash
git clone https://github.com/jalsarraf0/ai-chat-cli.git
cd ai-chat-cli
make build   # requires Go 1.24.x
```

### Interactive Installer
Run the guided setup:
```bash
./setup.sh
```

### Uninstall
Remove the binary and configuration:
```bash
./uninstall.sh
```

---

## Quick Start
```bash
export OPENAI_API_KEY="sk-..."   # set once
ai-chat                          # start interactive chat
```
![installer](assets/installer-demo.txt)
Use `Ctrl‑K` for the command palette.
The interface adapts to any terminal size and chooses a light or dark theme
based on `$COLORTERM`. Set `NO_COLOR=1` to disable ANSI colours entirely.

## Getting Help
Run `ai-chat --help` to list all commands and global options.
Use `ai-chat <command> --help` for command-specific flags or `man ai-chat` for a full manual page.

---

## Commands

| Command | Purpose |
|---------|---------|
| `chat` | Interactive REPL |
| `ask` | One‑off prompt |
| `plugins` | Manage plug‑ins |
| `history` | List/search old chats |
| `export` | Save chats |
| `config` | Show or edit config |
| `version` | Build info |

---

## Configuration
Default file `~/.config/ai-chat/ai-chat.yaml`:

```yaml
provider: openai
model: gpt-4o
temperature: 0.6
context_window: 16
plugins_dir: ~/.config/ai-chat/plugins
```

Environment variables (`AI_CHAT_MODEL`, etc.) override file values.

---

## Plug‑ins

Example hello‑world plug‑in:

```bash
#!/usr/bin/env bash
# ~/.config/ai-chat/plugins/hello
echo "Hello, $* 👋"
```

Any executable placed in the plug‑ins directory becomes a slash‑command: `/hello world`.

---

## Testing & CI
| Job | Tool | Gate |
|-----|------|------|
| **Lint** | golangci‑lint | no warnings |
| **Unit** | `go test -race` | 90 %+ coverage |
| **Security** | gosec, govulncheck | zero criticals |
| **Release** | GoReleaser v2 | binary, tar.gz, deb, rpm |

---

## Development
```bash
make              # format, vet, lint, test
make coverage     # HTML coverage report
make release      # goreleaser release --clean --skip=docker
```

---

## Contributing
1. Fork & branch: `git checkout -b feat/my-feature`
2. Write tests & code
3. `make` must pass
4. Open PR 🚀

---

## Security
Please report vulnerabilities via [GitHub Advisories](https://github.com/jalsarraf0/ai-chat-cli/security/advisories).
We follow a 90‑day disclosure window.

---

## License
MIT – see [LICENSE](LICENSE).

---

## Changelog
See [Releases](https://github.com/jalsarraf0/ai-chat-cli/releases) or [CHANGELOG.md](CHANGELOG.md).

---

## Acknowledgements
- OpenAI & Azure OpenAI
- Charm Bracelet (Bubble Tea)
- spf13/cobra
- GoReleaser
- Sigstore **cosign**
