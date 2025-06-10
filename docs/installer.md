# Installer Walkthrough

The `install.sh` script downloads and builds **ai-chat-cli**, copies the binary to `/usr/local/bin` and prompts for your
API key. The value is saved to `$XDG_CONFIG_HOME/ai-chat/ai-chat.yaml` along with a default model of `gpt-4o`.

Use `ai-chat config edit` to modify the file later if needed.
