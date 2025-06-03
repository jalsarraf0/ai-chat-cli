# AI-Chat-CLI

[![CI Status](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml)

A lightweight command-line interface for interacting with AI chat services, providing streamlined workflows for rapid experimentation.

## Quick Start

```bash
go build -o ai-chat
./ai-chat help
./ai-chat version
./ai-chat ping
```

The CLI auto-detects Bash, Zsh, Fish, PowerShell and Cmd, running seamlessly across shells. All CI jobs run on self-hosted runners.

### Commands

- `ping` – prints `pong` if the backend is reachable
- `version` – shows semantic version, commit and build date
