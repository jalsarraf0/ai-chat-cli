# ai-chat-cli

A command line interface for interacting with AI chat services.

## Usage

```sh
$ ai-chat help
```

### Install completions

Bash/Zsh:
```sh
echo 'eval "$(ai-chat completion bash)"' >> ~/.bashrc
```

Fish:
```sh
ai-chat completion fish | tee ~/.config/fish/completions/ai-chat.fish
```

PowerShell:
```powershell
ai-chat completion ps1 | Out-File -Encoding ascii $PROFILE
```

### Generate man pages

```sh
$ ai-chat man --dir dist/man
```
