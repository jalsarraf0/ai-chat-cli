# AI-Chat-CLI

[![CI Status](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml)
[![Latest Release](https://img.shields.io/github/v/release/jalsarraf0/ai-chat-cli?label=release)](https://github.com/jalsarraf0/ai-chat-cli/releases/latest)
[![Docker Image](https://img.shields.io/badge/container-ghcr.io%2Fjalsarraf0%2Fai--chat--cli-blue)](https://github.com/jalsarraf0/ai-chat-cli/pkgs/container/ai-chat-cli)

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
- `tui` – interactive terminal UI (see [docs/tui.md](docs/tui.md))
- `ask` – send a prompt to the LLM

## Embedded Assets

Embedded templates and colour themes can be inspected with `ai-chat assets`:

```bash
./ai-chat assets list
./ai-chat assets cat templates/system.tmpl
./ai-chat assets export themes/light.json /tmp/theme.json
```

### Ask

```bash
./ai-chat ask "Hello" --model gpt-3.5-turbo
```

| Flag | Description |
|------|-------------|
| `--model` | model ID |
| `--temperature` | sampling temperature |
| `--max_tokens` | token limit |
