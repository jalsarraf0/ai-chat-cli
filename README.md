# AI-Chat-CLI

[![CI (Linux · macOS · Windows)](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml)
[![Coverage Status](https://img.shields.io/badge/coverage-92%25-brightgreen)](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/go-1.24.x-blue)](https://golang.org/dl/)
[![Security](https://img.shields.io/badge/Security-%F0%9F%94%92%20Cosign%20(OIDC)-brightgreen)](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml)
[![Security Scan](https://img.shields.io/badge/Security%20Scan-%F0%9F%9B%A1%EF%B8%8F%20100%25%20clean-brightgreen)](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/jalsarraf0/ai-chat-cli)](LICENSE)
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
- [Plugin development tutorial](docs/tutorials/plugin_dev.md)

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
