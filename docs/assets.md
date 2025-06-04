# Embedded Assets

The CLI ships with default templates and colour themes compiled into the binary. You can list and export these files using the `ai-chat assets` command. When the `--config` flag specifies a custom templates directory, files there override the embedded versions.

Examples:

```bash
ai-chat assets list
ai-chat assets cat templates/system.tmpl
ai-chat assets export themes/light.json /tmp/theme.json
```
