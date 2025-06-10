# Configuration

The CLI stores settings in a YAML file located at:

- `$AI_CHAT_CONFIG` if set
- `$XDG_CONFIG_HOME/ai-chat/ai-chat.yaml` or `$HOME/.config/ai-chat/ai-chat.yaml`
  otherwise

Environment variables with prefix `AICHAT_` override file values. Keys include:

- `openai_api_key` – API token (required, env var `OPENAI_API_KEY`)
- `model` – allowed values `gpt-4`, `gpt-3.5-turbo`

Use `ai-chat config show` to print the config file path and contents, `ai-chat config set <key> <value>` to change a value, or `ai-chat config edit` to open the file in your editor.
