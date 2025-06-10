# AI‑Chat‑CLI 🤖

<p align="center">

  <!-- CI workflow -->
  <a href="https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci-final.yml">
    <img src="https://img.shields.io/github/actions/workflow/status/jalsarraf0/ai-chat-cli/ci-final.yml?label=CI%20%2B%20Release&logo=githubactions&logoColor=white&style=for-the-badge"
         alt="CI status"/>
  </a>

  <!-- Codecov – shows live percentage; will be bright-green at 93 %+ -->
  <a href="https://codecov.io/gh/jalsarraf0/ai-chat-cli">
    <img src="https://img.shields.io/codecov/c/github/jalsarraf0/ai-chat-cli/dev?label=Coverage&logo=codecov&logoColor=white&style=for-the-badge&color=brightgreen"
         alt="Test coverage"/>
  </a>

  <!-- Go Report Card – “A+” (clean) when everything passes -->
  <a href="https://goreportcard.com/report/github.com/jalsarraf0/ai-chat-cli">
    <img src="https://goreportcard.com/badge/github.com/jalsarraf0/ai-chat-cli?style=for-the-badge"
         alt="Go Report Card"/>
  </a>

  <!-- Latest release – explicitly shows v1.0.1 -->
  <a href="https://github.com/jalsarraf0/ai-chat-cli/releases/tag/v1.0.1">
    <img src="https://img.shields.io/github/v/release/jalsarraf0/ai-chat-cli?include_prereleases&label=v1.0.1&logo=github&logoColor=white&style=for-the-badge"
         alt="Latest release"/>
  </a>

  <!-- MIT licence – static badge so it never “misdetects” -->
  <a href="https://github.com/jalsarraf0/ai-chat-cli/blob/dev/LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-blue?style=for-the-badge"
         alt="License: MIT"/>
  </a>

</p>

> **ai‑chat‑cli** ✨ is a lightweight command‑line tool for GPT‑style models.
> Written in **Go**, it streams answers in real time, remembers chat history and supports plug‑ins — all in a ~6 MiB binary.

## Quick Start 🚀

```bash
curl -fsSL https://raw.githubusercontent.com/jalsarraf0/ai-chat-cli/main/scripts/install.sh | bash
ai-chat "Hello"
```

The installer checks prerequisites, compiles the binary and copies it to `/usr/local/bin`.
It stores your API key in `$XDG_CONFIG_HOME/ai-chat/ai-chat.yaml` so you only enter it once.

---

## 📚 Table of Contents

1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Installation](#installation-)

    - [Packages](#packages)

    - [Build from Source](#build-from-source)

4. [Quick Start](#quick-start-)
5. [Commands](#commands-)
6. [Configuration](#configuration-)
7. [Plug‑ins](#plugins-)
8. [Testing & CI](#testing--ci-)
9. [Development](#development-)
10. [Contributing](#contributing-)
11. [Security](#security-)
12. [License](#license-)
13. [Changelog](#changelog-)
14. [Acknowledgements](#acknowledgements-)

---

## Overview

`ai-chat` wraps the OpenAI, Azure OpenAI and Ollama APIs behind a consistent CLI.
Key features:

| Feature                     | Description                                                |
| --------------------------- | ---------------------------------------------------------- |
| 📨 **Streaming chat**       | Low‑latency updates with ANSI markdown rendering           |
| 🗂 **Persistent history**   | All chats stored in SQLite and searchable                  |
| 🔌 **Plug‑in framework**    | Drop shell scripts in `~/.config/ai-chat/plugins`          |
| 🔒 **Keychain integration** | Secrets stored via `pass`, macOS Keychain or Windows DPAPI |
| 🖇 **Exporters**            | Save replies to clipboard, Markdown, HTML, JSON, PDF       |
| ⚙️ **Embeddable**           | Public Go API (`pkg/chat`) for your own apps               |

---

## Architecture

```text
client (cobra CLI) ─┐
                    ├─► core engine ──► provider interface ──► OpenAI / Azure / Ollama
plugins (shell) ────┘
                             │
                             └──► SQLite history
```

---

## Installation 📦

### Packages

| Package      | Command                                                                                           |
| ------------ | ------------------------------------------------------------------------------------------------- |
| **tar.gz**   | `tar -xzf ai-chat_linux_amd64.tar.gz && sudo mv ai-chat /usr/local/bin`                           |
| **DEB**      | `sudo dpkg -i ai-chat_<ver>_amd64.deb`                                                            |
| **RPM**      | `sudo rpm -Uvh ai-chat_<ver>_amd64.rpm`                                                           |
| **Homebrew** | `brew install jalsarraf0/tap/ai-chat`                                                             |
| **Scoop**    | `scoop bucket add jalsarraf0 https://github.com/jalsarraf0/scoop-bucket && scoop install ai-chat` |

### Build from Source

```bash
git clone https://github.com/jalsarraf0/ai-chat-cli.git
cd ai-chat-cli
make build   # requires Go 1.24.x
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

## Usage 💻

```bash
export OPENAI_API_KEY="sk-..."   # set once
ai-chat                          # start interactive chat
```

![installer](assets/installer-demo.txt)
Use `Ctrl‑K` for the command palette.
The interface adapts to any terminal size and chooses a light or dark theme
based on `$COLORTERM`. Set `NO_COLOR=1` to disable ANSI colours entirely.

## Getting Help

Run `ai-chat --help` for available commands. Use `ai-chat <command> --help` for details.

---

## Commands 🛠

| Command    | Purpose               |
| ---------- | --------------------- |
| _(prompt)_ | One‑off question      |
| `plugins`  | Manage plug‑ins       |
| `history`  | List/search old chats |
| `export`   | Save chats            |
| `config`   | Show or edit config   |
| `version`  | Build info            |

---

## Configuration ⚙

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

## Plug‑ins 🔌

Example hello‑world plug‑in:

```bash
#!/usr/bin/env bash
# ~/.config/ai-chat/plugins/hello
echo "Hello, $* 👋"
```

Any executable placed in the plug‑ins directory becomes a slash‑command: `/hello world`.

---

## Testing & CI ✅

| Job          | Tool               | Gate           |
| ------------ | ------------------ | -------------- |
| **Lint**     | golangci‑lint      | no warnings    |
| **Unit**     | `go test -race`    | 90 %+ coverage |
| **Security** | gosec, govulncheck | zero criticals |

---

## Development 👷

```bash
make              # format, vet, lint, test
make coverage     # HTML coverage report
```

---

## Contributing 🤝

1. Fork & branch: `git checkout -b feat/my-feature`
2. Write tests & code
3. `make` must pass
4. Open PR 🚀

---

## Security 🔐

Please report vulnerabilities via [GitHub Advisories](https://github.com/jalsarraf0/ai-chat-cli/security/advisories).
We follow a 90‑day disclosure window.

---

## License 📝

MIT – see [LICENSE](LICENSE).

---

## Changelog 📜

See [CHANGELOG](CHANGELOG) for recent updates.

See [Releases](https://github.com/jalsarraf0/ai-chat-cli/releases) or [CHANGELOG.md](CHANGELOG.md).

---

## Acknowledgements 🙏

- OpenAI & Azure OpenAI
- Charm Bracelet (Bubble Tea)
- spf13/cobra
- Sigstore **cosign**
