# AI-Chat-CLI

[![CI Status](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml)

A lightweight command-line interface for interacting with AI chat services, providing streamlined workflows for rapid experimentation.

Configuration is stored in a YAML file; see [docs/config.md](docs/config.md) for available keys.

## Quick Start

```bash
go build -o ai-chat
./ai-chat help
./ai-chat version
./ai-chat ping
```

The CLI auto-detects Bash, Zsh, Fish, PowerShell and Cmd, running seamlessly across shells. Linux jobs run on self-hosted runners while macOS and Windows use GitHub-hosted runners.

### Commands

- `ping` – prints `pong` if the backend is reachable
- `version` – shows semantic version, commit and build date
- `config` – manage settings stored in `~/.config/ai-chat/config.yaml` (run `ai-chat config show` to print the path and contents)

## Embedded Assets

Embedded templates and colour themes can be inspected with `ai-chat assets`:

```bash
./ai-chat assets list
./ai-chat assets cat templates/system.tmpl
./ai-chat assets export themes/light.json /tmp/theme.json
```
