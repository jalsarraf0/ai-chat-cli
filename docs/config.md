# Configuration

The CLI stores settings in a YAML file. By default the path is:

- **Unix:** `~/.config/ai-chat/config.yaml`
- **Windows:** `%APPDATA%\ai-chat\config.yaml`

Environment variables with prefix `AICHAT_` override file values. Keys include:

- `openai_api_key` – API token (required)
- `model` – allowed values `gpt-4`, `gpt-3.5-turbo`

Use `ai-chat config show` to print the current configuration, `ai-chat config set <key> <value>` to change a value, or `ai-chat config edit` to open the file in your editor.
