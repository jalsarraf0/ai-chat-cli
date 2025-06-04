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

The CLI auto-detects Bash, Zsh, Fish, PowerShell and Cmd, running seamlessly across shells. Linux jobs run on self-hosted runners while macOS and Windows use GitHub-hosted runners.

### Commands

- `ping` – prints `pong` if the backend is reachable
- `version` – shows semantic version, commit and build date

### Shell Completion

| Shell | Install |
|-------|--------|
| Bash | `ai-chat completion bash --out ~/.ai-chat/completion.bash` |
| Zsh | `ai-chat completion zsh --out ~/.ai-chat/completion.zsh` |
| Fish | `ai-chat completion fish --out ~/.ai-chat/completion.fish` |
| PowerShell | `ai-chat completion powershell --out ~/.ai-chat/completion.ps1` |

### Prompt

Run `ai-chat init` to install completions and a prompt enhancer showing `(ai)›` in green when a `.ai-chat.yml` file is present.
